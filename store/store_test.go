package store

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStore_Namespaces(t *testing.T) {
	fs := afero.NewMemMapFs()

	s, err := NewTempStore(TempStoreOptFS(fs))
	require.NoError(t, err)

	defer s.Close()

	names, err := s.Namespaces()
	require.NoError(t, err)
	require.Len(t, names, 0)

	err = s.fs.MkdirAll(filepath.Join(s.dir, "ns"), dirMode)
	require.NoError(t, err)

	names, err = s.Namespaces()
	require.NoError(t, err)
	require.Len(t, names, 1)
}

func TestStore_Packages(t *testing.T) {
	fs := afero.NewMemMapFs()

	s, err := NewTempStore(TempStoreOptFS(fs))
	require.NoError(t, err)
	defer s.Close()

	names, err := s.Packages("ns")
	require.NoError(t, err)
	require.Len(t, names, 0)

	err = s.fs.MkdirAll(filepath.Join(s.dir, "ns", "pkg"), dirMode)
	require.NoError(t, err)

	names, err = s.Packages("ns")
	require.NoError(t, err)
	require.Len(t, names, 1)
}

func TestStore_Releases(t *testing.T) {
	fs := afero.NewMemMapFs()

	s, err := NewTempStore(TempStoreOptFS(fs))
	require.NoError(t, err)
	defer s.Close()

	names, err := s.Releases("ns", "pkg")
	require.NoError(t, err)
	require.Len(t, names, 0)

	p := filepath.Join(s.dir, "ns", "pkg", "releases")
	err = s.fs.MkdirAll(p, dirMode)
	require.NoError(t, err)

	digestDir := filepath.Join(s.dir, "ns", "pkg", "digests", "12345")

	writeFile(t, fs, filepath.Join(p, "0.1.0"), []byte("12345"))
	writeFile(t, fs, filepath.Join(digestDir, "part.tar.gz"), []byte("contents"))

	rms, err := s.Releases("ns", "pkg")
	require.NoError(t, err)
	require.Len(t, rms, 1)
}

func TestStore_CreateRelease(t *testing.T) {
	fs := afero.NewMemMapFs()

	s, err := NewTempStore(TempStoreOptFS(fs))
	require.NoError(t, err)
	defer s.Close()

	b, err := ioutil.ReadFile("testdata/node.tar.gz")
	require.NoError(t, err)

	rm, err := s.CreateRelease("ns", "pkg", "0.1.0", b)
	require.NoError(t, err)

	digest := "e2a28469635e14461126bfd0fcdf7d47c9d1516444e2e5ace79d139d1cbd1d48"
	assert.Equal(t, rm.Digest, digest)

	fi, err := fs.Stat(filepath.Join(s.dir, "ns", "pkg", "digests", digest, "part.tar.gz"))
	assert.NoError(t, err)
	assert.Equal(t, rm.CreatedAt, fi.ModTime())
	assert.Equal(t, rm.Size, fi.Size())
	assert.Equal(t, rm.Version, "0.1.0")

	digestFiles := []string{"part.tar.gz", "parts.yaml", "README.md"}
	for _, f := range digestFiles {
		_, err = fs.Stat(filepath.Join(s.dir, "ns", "pkg", "digests", digest, f))
		assert.NoError(t, err)
	}

	releasePath := filepath.Join(s.dir, "ns", "pkg", "releases", "0.1.0")
	_, err = fs.Stat(releasePath)
	assert.NoError(t, err)

	b, err = afero.ReadFile(fs, releasePath)
	assert.NoError(t, err)
	assert.Equal(t, digest, string(b))
}

func TestStore_RemoveRelease(t *testing.T) {
	fs := afero.NewMemMapFs()

	s, err := NewTempStore(TempStoreOptFS(fs))
	require.NoError(t, err)
	defer s.Close()

	digestDir := filepath.Join(s.dir, "ns", "pkg", "digests", "12345")
	releaseDir := filepath.Join(s.dir, "ns", "pkg", "releases")

	for _, dir := range []string{digestDir, releaseDir} {
		err = fs.MkdirAll(dir, dirMode)
		require.NoError(t, err)
	}

	writeFile(t, fs, filepath.Join(releaseDir, "0.1.0"), []byte("12345"))

	err = s.RemoveRelease("ns", "pkg", "0.1.0")
	require.NoError(t, err)

	_, err = fs.Stat(digestDir)
	assert.True(t, os.IsNotExist(err))

	_, err = fs.Stat(filepath.Join(releaseDir, "0.1.0"))
	assert.True(t, os.IsNotExist(err))
}

func TestStore_Release(t *testing.T) {
	cases := []struct {
		name   string
		ver    string
		digest string
		isErr  bool
	}{
		{
			name:   "existing release",
			ver:    "0.1.0",
			digest: "12345",
		},
		{
			name:  "non existent release",
			ver:   "0.2.0",
			isErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fs := afero.NewMemMapFs()

			s, err := NewTempStore(TempStoreOptFS(fs))
			require.NoError(t, err)
			defer s.Close()

			releaseDir := filepath.Join(s.dir, "ns", "pkg", "releases")
			err = fs.MkdirAll(releaseDir, dirMode)
			require.NoError(t, err)

			writeFile(t, fs, filepath.Join(releaseDir, "0.1.0"), []byte("12345"))

			digestDir := filepath.Join(s.dir, "ns", "pkg", "digests", "12345")
			err = fs.MkdirAll(digestDir, dirMode)
			require.NoError(t, err)

			blobName := filepath.Join(digestDir, "part.tar.gz")
			writeFile(t, fs, blobName, []byte("contents"))

			aTime := time.Date(2009, time.January, 20, 0, 0, 0, 0, time.UTC)

			err = fs.Chtimes(blobName, aTime, aTime)
			require.NoError(t, err)

			rm, err := s.Release("ns", "pkg", tc.ver)

			if tc.isErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)

				expected := ReleaseMetadata{
					Digest:    "12345",
					Size:      8,
					CreatedAt: aTime,
					Version:   "0.1.0",
				}
				assert.Equal(t, expected, rm)
			}
		})
	}
}

func TestStore_Pull(t *testing.T) {
	fs := afero.NewMemMapFs()

	s, err := NewTempStore(TempStoreOptFS(fs))
	require.NoError(t, err)
	defer s.Close()

	digestDir := filepath.Join(s.dir, "ns", "pkg", "digests", "12345")
	err = fs.MkdirAll(digestDir, dirMode)
	require.NoError(t, err)

	b, err := ioutil.ReadFile("testdata/node.tar.gz")
	require.NoError(t, err)

	writeFile(t, fs, filepath.Join(digestDir, "part.tar.gz"), b)

	f, err := s.Pull("ns", "pkg", "12345")
	require.NoError(t, err)

	b2, err := ioutil.ReadAll(f)
	require.NoError(t, err)

	assert.Equal(t, b, b2)
}

func writeFile(t *testing.T, fs afero.Fs, name string, contents []byte) {
	err := afero.WriteFile(fs, name, contents, fileMode)
	require.NoError(t, err)
}
