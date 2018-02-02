package store

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStore_Namespaces(t *testing.T) {
	withFileSystemStore(t, func(s *FileSystemStore, fs afero.Fs) {
		names, err := s.Namespaces()
		require.NoError(t, err)
		require.Len(t, names, 0)

		mkdirAll(t, fs, filepath.Join(s.dir, "ns"))

		names, err = s.Namespaces()
		require.NoError(t, err)
		require.Len(t, names, 1)
	})
}

func TestStore_CreatePackage(t *testing.T) {
	cases := []struct {
		name         string
		ns           string
		pkg          string
		existingDirs []string
		isErr        bool
	}{
		{
			name: "package does not exist",
			ns:   "ns",
			pkg:  "pkg",
		},
		{
			name:         "package exists",
			ns:           "ns",
			pkg:          "pkg",
			existingDirs: []string{"ns/pkg"},
			isErr:        true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			withFileSystemStore(t, func(s *FileSystemStore, fs afero.Fs) {
				for _, dir := range tc.existingDirs {
					dir = filepath.ToSlash(dir)
					parts := strings.Split(dir, "/")

					parts = append([]string{s.dir}, parts...)
					mkdirAll(t, fs, filepath.Join(parts...))
				}

				pm, err := s.CreatePackage(tc.ns, tc.pkg)

				if tc.isErr {
					require.Error(t, err)
				} else {
					require.NoError(t, err)

					assert.Equal(t, tc.ns, pm.Namespace)
					assert.Equal(t, tc.pkg, pm.Package)
					assert.True(t, pm.IsVisible)

				}
			})
		})
	}

}

func TestStore_Packages(t *testing.T) {
	withFileSystemStore(t, func(s *FileSystemStore, fs afero.Fs) {
		names, err := s.Packages("ns")
		require.NoError(t, err)
		require.Len(t, names, 0)

		mkdirAll(t, fs, filepath.Join(s.dir, "ns", "pkg"))
		writeFile(t, fs, s.pkgMetadata("ns", "pkg"), testdata(t, "pkgMetadata.yaml"))

		names, err = s.Packages("ns")
		require.NoError(t, err)
		require.Len(t, names, 1)
	})
}

func TestStore_Releases(t *testing.T) {
	withFileSystemStore(t, func(s *FileSystemStore, fs afero.Fs) {
		names, err := s.Releases("ns", "pkg")
		require.NoError(t, err)
		require.Len(t, names, 0)

		p := filepath.Join(s.dir, "ns", "pkg", "releases")
		mkdirAll(t, fs, p)

		digestDir := filepath.Join(s.dir, "ns", "pkg", "digests", "12345")

		partsName := filepath.Join(digestDir, "parts.yaml")
		b, err := ioutil.ReadFile("testdata/parts.yaml")
		require.NoError(t, err)
		writeFile(t, fs, partsName, b)
		writeFile(t, fs, filepath.Join(p, "0.1.0"), []byte("12345"))
		writeFile(t, fs, filepath.Join(digestDir, "part.tar.gz"), []byte("contents"))

		rms, err := s.Releases("ns", "pkg")
		require.NoError(t, err)
		require.Len(t, rms, 1)
	})
}

func TestStore_CreateRelease(t *testing.T) {
	withFileSystemStore(t, func(s *FileSystemStore, fs afero.Fs) {
		b, err := ioutil.ReadFile("testdata/node.tar.gz")
		require.NoError(t, err)

		rm, err := s.CreateRelease("ns", "pkg", "0.1.0", b)
		require.NoError(t, err)

		digest := "e2a28469635e14461126bfd0fcdf7d47c9d1516444e2e5ace79d139d1cbd1d48"
		assert.Equal(t, rm.Digest, digest)

		digestDir := filepath.Join(s.dir, "ns", "pkg", "digests", digest)
		partsName := filepath.Join(digestDir, "parts.yaml")
		b, err = ioutil.ReadFile("testdata/parts.yaml")
		require.NoError(t, err)
		writeFile(t, fs, partsName, b)

		fi, err := fs.Stat(filepath.Join(digestDir, "part.tar.gz"))
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
	})
}

func TestStore_RemoveRelease(t *testing.T) {
	withFileSystemStore(t, func(s *FileSystemStore, fs afero.Fs) {
		digestDir := filepath.Join(s.dir, "ns", "pkg", "digests", "12345")
		releaseDir := filepath.Join(s.dir, "ns", "pkg", "releases")

		for _, dir := range []string{digestDir, releaseDir} {
			mkdirAll(t, fs, dir)
		}

		writeFile(t, fs, filepath.Join(releaseDir, "0.1.0"), []byte("12345"))

		err := s.RemoveRelease("ns", "pkg", "0.1.0")
		require.NoError(t, err)

		_, err = fs.Stat(digestDir)
		assert.True(t, os.IsNotExist(err))

		_, err = fs.Stat(filepath.Join(releaseDir, "0.1.0"))
		assert.True(t, os.IsNotExist(err))
	})
}

func TestStore_Release(t *testing.T) {
	cases := []struct {
		name   string
		ver    string
		digest string
		isErr  bool
		depLen int
	}{
		{
			name:   "existing release",
			ver:    "0.1.0",
			digest: "12345",
			depLen: 11,
		},
		{
			name:  "non existent release",
			ver:   "0.2.0",
			isErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			withFileSystemStore(t, func(s *FileSystemStore, fs afero.Fs) {
				releaseDir := filepath.Join(s.dir, "ns", "pkg", "releases")
				mkdirAll(t, fs, releaseDir)

				writeFile(t, fs, filepath.Join(releaseDir, "0.1.0"), []byte("12345"))

				digestDir := filepath.Join(s.dir, "ns", "pkg", "digests", "12345")
				mkdirAll(t, fs, digestDir)

				blobName := filepath.Join(digestDir, "part.tar.gz")
				writeFile(t, fs, blobName, []byte("contents"))

				partsName := filepath.Join(digestDir, "parts.yaml")
				b, err := ioutil.ReadFile("testdata/parts.yaml")
				require.NoError(t, err)
				writeFile(t, fs, partsName, b)

				aTime := time.Date(2009, time.January, 20, 0, 0, 0, 0, time.UTC)

				err = fs.Chtimes(blobName, aTime, aTime)
				require.NoError(t, err)

				rm, err := s.Release("ns", "pkg", tc.ver)

				if tc.isErr {
					require.Error(t, err)
				} else {
					require.NoError(t, err)

					assert.Equal(t, "12345", rm.Digest)
					assert.Equal(t, int64(8), rm.Size)
					assert.Equal(t, aTime, rm.CreatedAt)
					assert.Equal(t, "0.1.0", rm.Version)
					assert.Len(t, rm.Dependencies, tc.depLen)
				}
			})
		})
	}
}

func TestStore_Pull(t *testing.T) {
	withFileSystemStore(t, func(s *FileSystemStore, fs afero.Fs) {
		digestDir := filepath.Join(s.dir, "ns", "pkg", "digests", "12345")
		mkdirAll(t, fs, digestDir)

		b, err := ioutil.ReadFile("testdata/node.tar.gz")
		require.NoError(t, err)

		writeFile(t, fs, filepath.Join(digestDir, "part.tar.gz"), b)

		f, err := s.Pull("ns", "pkg", "12345")
		require.NoError(t, err)

		b2, err := ioutil.ReadAll(f)
		require.NoError(t, err)

		assert.Equal(t, b, b2)
	})
}

func writeFile(t *testing.T, fs afero.Fs, name string, contents []byte) {
	err := afero.WriteFile(fs, name, contents, fileMode)
	require.NoError(t, err)
}

func mkdirAll(t *testing.T, fs afero.Fs, dir string) {
	err := fs.MkdirAll(dir, dirMode)
	require.NoError(t, err)
}

func testdata(t *testing.T, name string) []byte {
	path := filepath.Join("testdata", name)
	b, err := ioutil.ReadFile(path)
	require.NoError(t, err)

	return b
}

func withFileSystemStore(t *testing.T, fn func(*FileSystemStore, afero.Fs)) {
	fs := afero.NewMemMapFs()
	closeFn := func() error {
		return os.RemoveAll("/data")
	}

	fs.Mkdir("/data", dirMode)

	s, err := NewFileSystemStore(
		FileSystemStoreOptFS(fs),
		FileSystemStoreOptRoot("/data"),
		FileSystemStoreOptClose(closeFn))

	require.NoError(t, err)

	defer s.Close()

	fn(s, fs)
}
