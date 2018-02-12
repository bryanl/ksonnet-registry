package store

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/pkg/errors"
	"github.com/spf13/afero"

	"github.com/bryanl/ksonnet-registry/repository"
	repomocks "github.com/bryanl/ksonnet-registry/repository/mocks"
	storagemocks "github.com/bryanl/ksonnet-registry/storage/mocks"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newID(t *testing.T) uuid.UUID {
	id, err := uuid.NewV4()
	require.NoError(t, err)

	return id
}

func TestS3_CreateNamespace(t *testing.T) {
	withS3(t, func(fs afero.Fs, s *S3, m *mocks) {
		id := newID(t)

		res := repository.Namespace{
			ID:        id,
			Namespace: "ns",
		}

		m.nsRepo.On("Create", "ns").Return(res, nil)

		got, err := s.CreateNamespace("ns")
		require.NoError(t, err)

		expected := NamespaceMetdata{
			ID:        id,
			Namespace: "ns",
		}

		require.Equal(t, expected, got)
	})
}

func TestS3_Namespaces(t *testing.T) {
	withS3(t, func(fs afero.Fs, s *S3, m *mocks) {
		namespaces := []repository.Namespace{}
		m.nsRepo.On("List").Return(namespaces, nil).Once()

		names, err := s.Namespaces()
		require.NoError(t, err)
		require.Len(t, names, 0)

		namespaces = []repository.Namespace{{Namespace: "ns"}}
		m.nsRepo.On("List").Return(namespaces, nil).Once()

		names, err = s.Namespaces()
		require.NoError(t, err)
		require.Len(t, names, 1)
	})
}

func TestS3_CreatePackage(t *testing.T) {
	withS3(t, func(fs afero.Fs, s *S3, m *mocks) {
		id := newID(t)
		createdAt := time.Now().UTC()

		res := repository.Package{
			ID:         id,
			Namespace:  "ns",
			Package:    "pkg",
			CreatedAt:  createdAt,
			Visibility: true,
		}

		m.pkgRepo.On("Create", "ns", "pkg").
			Return(res, nil)

		got, err := s.CreatePackage("ns", "pkg")
		require.NoError(t, err)

		expected := PackageMetadata{
			ID:        id,
			Namespace: "ns",
			Package:   "pkg",
			IsVisible: true,
			CreatedAt: createdAt,
		}

		require.Equal(t, expected, got)
	})
}

func TestS3_Packages(t *testing.T) {
	withS3(t, func(fs afero.Fs, s *S3, m *mocks) {
		id := newID(t)
		createdAt := time.Now().UTC()

		res := []repository.Package{
			{ID: id,
				Namespace:  "ns",
				Package:    "pkg",
				CreatedAt:  createdAt,
				Visibility: true,
			},
		}

		m.pkgRepo.On("List", "ns").
			Return(res, nil)

		got, err := s.Packages("ns")
		require.NoError(t, err)

		expected := []PackageMetadata{
			{
				ID:        id,
				Namespace: "ns",
				Package:   "pkg",
				IsVisible: true,
				CreatedAt: createdAt,
			},
		}

		require.Equal(t, expected, got)

	})
}
func TestS3_Package(t *testing.T) {
	withS3(t, func(fs afero.Fs, s *S3, m *mocks) {
		id := newID(t)
		createdAt := time.Now().UTC()

		res := repository.Package{
			ID:         id,
			Namespace:  "ns",
			Package:    "pkg",
			CreatedAt:  createdAt,
			Visibility: true,
		}

		m.pkgRepo.On("Retrieve", "ns", "pkg").
			Return(res, nil)

		got, err := s.Package("ns", "pkg")
		require.NoError(t, err)

		expected := PackageMetadata{
			ID:        id,
			Namespace: "ns",
			Package:   "pkg",
			IsVisible: true,
			CreatedAt: createdAt,
		}

		require.Equal(t, expected, got)
	})
}

func TestS3_Releases(t *testing.T) {
	withS3(t, func(fs afero.Fs, s *S3, m *mocks) {
		id := newID(t)
		createdAt := time.Now().UTC()

		res := []repository.Release{
			{
				ID:        id,
				Namespace: "ns",
				Package:   "pkg",
				Digest:    "12345",
				Size:      3456,
				CreatedAt: createdAt,
				Version:   "0.1.0",
			},
		}
		m.releaseRepo.On("List", "ns", "pkg").
			Return(res, nil)

		got, err := s.Releases("ns", "pkg")
		require.NoError(t, err)

		expected := []ReleaseMetadata{
			{
				ID:        id,
				Namespace: "ns",
				Package:   "pkg",
				Digest:    "12345",
				Size:      3456,
				CreatedAt: createdAt,
				Version:   "0.1.0",
			},
		}

		require.Equal(t, expected, got)
	})
}

func TestS3_CreateRelease(t *testing.T) {
	withS3(t, func(fs afero.Fs, s *S3, m *mocks) {
		digest := "e2a28469635e14461126bfd0fcdf7d47c9d1516444e2e5ace79d139d1cbd1d48"

		pkg := repository.Package{}
		m.pkgRepo.On("Retrieve", "ns", "pkg").Return(pkg, nil)

		m.releaseRepo.On("RetrieveDigest", digest).Return(repository.Release{}, errors.New("not found"))

		m.releaseRepo.On("Retrieve", "ns", "pkg", "0.1.0").Return(repository.Release{}, errors.New("not found"))

		partPath := fmt.Sprintf("%s/%s", digest, blobName)
		m.s3Client.On("PutObject", partPath, blobMIMEType, mock.AnythingOfType("*bytes.Reader"), mock.AnythingOfType("int64")).
			Return(99, nil)
		configPath := fmt.Sprintf("%s/%s", digest, configName)
		m.s3Client.On("PutObject", configPath, configMIMEType, mock.AnythingOfType("*bytes.Reader"), mock.AnythingOfType("int64")).
			Return(99, nil)
		docPath := fmt.Sprintf("%s/%s", digest, docName)
		m.s3Client.On("PutObject", docPath, docMIMEType, mock.AnythingOfType("*bytes.Reader"), mock.AnythingOfType("int64")).
			Return(99, nil)

		now := time.Now()
		id, err := uuid.NewV4()
		require.NoError(t, err)
		release := repository.Release{
			ID:           id,
			Namespace:    "ns",
			Package:      "pkg",
			Digest:       digest,
			Size:         99,
			CreatedAt:    now,
			Version:      "0.1.0",
			Dependencies: repository.Dependencies{},
		}
		m.releaseRepo.On("Create", "ns", "pkg", "0.1.0", digest, 99, mock.AnythingOfType("time.Time")).
			Return(release, nil)

		b, err := ioutil.ReadFile("testdata/node.tar.gz")
		require.NoError(t, err)

		rm, err := s.CreateRelease("ns", "pkg", "0.1.0", b)
		require.NoError(t, err)

		assert.Equal(t, now, rm.CreatedAt)
		assert.Equal(t, int64(99), rm.Size)
		assert.Equal(t, "0.1.0", rm.Version)
	})
}

func TestS3_RemoveRelease(t *testing.T) {
	withS3(t, func(fs afero.Fs, s *S3, m *mocks) {
		release := repository.Release{
			Digest: "12345",
		}
		m.releaseRepo.On("Retrieve", "ns", "pkg", "0.1.0").Return(release, nil)
		m.releaseRepo.On("Delete", "ns", "pkg", "0.1.0").Return(nil)
		m.s3Client.On("RemoveObject", "12345/part.tar.gz").Return(nil)
		m.s3Client.On("RemoveObject", "12345/parts.yaml").Return(nil)
		m.s3Client.On("RemoveObject", "12345/README.md").Return(nil)

		err := s.RemoveRelease("ns", "pkg", "0.1.0")
		require.NoError(t, err)
	})
}

func TestS3_Release(t *testing.T) {
	withS3(t, func(fs afero.Fs, s *S3, m *mocks) {
		now := time.Now()
		id, err := uuid.NewV4()
		require.NoError(t, err)
		release := repository.Release{
			ID:           id,
			Namespace:    "ns",
			Package:      "pkg",
			Digest:       "12345",
			Size:         99,
			CreatedAt:    now,
			Version:      "0.1.0",
			Dependencies: repository.Dependencies{},
		}

		m.releaseRepo.On("Retrieve", "ns", "pkg", "0.1.0").Return(release, nil)

		rm, err := s.Release("ns", "pkg", "0.1.0")
		require.NoError(t, err)

		assert.Equal(t, release.ID, rm.ID)
		assert.Equal(t, release.Namespace, rm.Namespace)
		assert.Equal(t, release.Package, rm.Package)
		assert.Equal(t, release.Digest, rm.Digest)
		assert.Equal(t, release.Size, rm.Size)
		assert.Equal(t, release.CreatedAt, rm.CreatedAt)
		assert.Equal(t, release.Version, rm.Version)
		assert.Equal(t, release.Dependencies, rm.Dependencies)
	})
}

func TestS3_Pull(t *testing.T) {
	withS3(t, func(fs afero.Fs, s *S3, m *mocks) {
		release := repository.Release{
			Digest: "12345",
		}

		m.releaseRepo.On("RetrieveDigest", "12345").Return(release, nil)

		err := afero.WriteFile(fs, "/file", []byte(`contents`), fileMode)
		require.NoError(t, err)

		f, err := fs.Open("/file")
		require.NoError(t, err)

		m.s3Client.On("GetObject", "12345/part.tar.gz").Return(f, nil)

		object, err := s.Pull("ns", "pkg", "12345")
		require.NoError(t, err)
		defer f.Close()

		b, err := ioutil.ReadAll(object)
		require.NoError(t, err)
		require.Equal(t, "contents", string(b))
	})
}

type mocks struct {
	s3Client    *storagemocks.S3Client
	repo        *repomocks.Repository
	nsRepo      *repomocks.NamespaceRepository
	pkgRepo     *repomocks.PackageRepository
	releaseRepo *repomocks.ReleaseRepository
}

func withS3(t *testing.T, fn func(afero.Fs, *S3, *mocks)) {
	m := &mocks{
		s3Client:    &storagemocks.S3Client{},
		nsRepo:      &repomocks.NamespaceRepository{},
		pkgRepo:     &repomocks.PackageRepository{},
		releaseRepo: &repomocks.ReleaseRepository{},
		repo:        &repomocks.Repository{},
	}

	m.repo.On("Namespace").Return(m.nsRepo)
	m.repo.On("Package").Return(m.pkgRepo)
	m.repo.On("Release").Return(m.releaseRepo)

	fs := afero.NewMemMapFs()

	s, err := NewS3(fs, m.s3Client, m.repo)
	require.NoError(t, err)

	defer s.Close()

	fn(fs, s, m)
}
