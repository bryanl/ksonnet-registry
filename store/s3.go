package store

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"time"

	"github.com/bryanl/ksonnet-registry/repository"
	"github.com/bryanl/ksonnet-registry/storage"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

// S3 is a s3 based storage. It uses a datasource to store metadata and s3 for file storage.
type S3 struct {
	client     storage.S3Client
	repository repository.Repository
	fs         afero.Fs
}

var _ Store = (*S3)(nil)

// NewS3 creates an instance of S3.
func NewS3(fs afero.Fs, client storage.S3Client, repo repository.Repository) (*S3, error) {
	if fs == nil {
		return nil, errors.New("fs is nil")
	}

	if client == nil {
		return nil, errors.New("s3 client is nil")
	}

	if repo == nil {
		return nil, errors.New("repository is nil")
	}

	s3 := &S3{
		fs:         fs,
		client:     client,
		repository: repo,
	}

	return s3, nil
}

func (s *S3) nsRepo() repository.NamespaceRepository {
	return s.repository.Namespace()
}

func (s *S3) pkgRepo() repository.PackageRepository {
	return s.repository.Package()
}

func (s *S3) releaseRepo() repository.ReleaseRepository {
	return s.repository.Release()
}

// CreateNamespace creates a namespace.
func (s *S3) CreateNamespace(ns string) (NamespaceMetdata, error) {
	res, err := s.nsRepo().Create(ns)
	if err != nil {
		return NamespaceMetdata{}, err
	}

	return NamespaceMetdata{
		ID:        res.ID,
		Namespace: res.Namespace,
	}, nil
}

// Namespaces returns a list of namespaces in the store.
func (s *S3) Namespaces() ([]NamespaceMetdata, error) {
	nsRepo := s.nsRepo()

	list, err := nsRepo.List()
	if err != nil {
		return nil, err
	}

	var nms []NamespaceMetdata

	for _, ns := range list {
		nm := NamespaceMetdata{
			Namespace: ns.Namespace,
		}

		nms = append(nms, nm)
	}

	return nms, nil
}

// CreatePackage creates a package.
func (s *S3) CreatePackage(ns, pkg string) (PackageMetadata, error) {
	pkgRepo := s.pkgRepo()

	res, err := pkgRepo.Create(ns, pkg)
	if err != nil {
		return PackageMetadata{}, err
	}

	return PackageMetadata{
		ID:        res.ID,
		Namespace: res.Namespace,
		Package:   res.Package,
		CreatedAt: res.CreatedAt,
		IsVisible: res.Visibility,
	}, nil
}

// Packages lists all packages in a namespace.
func (s *S3) Packages(ns string) ([]PackageMetadata, error) {
	pkgRepo := s.pkgRepo()

	ps, err := pkgRepo.List(ns)
	if err != nil {
		return nil, err
	}

	var pms []PackageMetadata
	for _, p := range ps {
		pms = append(pms, convertRepoPackageToMetadata(p))
	}

	return pms, nil
}

// Package retrieves a package from a namespace.
func (s *S3) Package(ns, pkg string) (PackageMetadata, error) {
	pkgRepo := s.pkgRepo()

	res, err := pkgRepo.Retrieve(ns, pkg)
	if err != nil {
		return PackageMetadata{}, err
	}

	return convertRepoPackageToMetadata(res), nil
}

// Releases retrieves all releases from a package.
func (s *S3) Releases(ns string, pkg string) ([]ReleaseMetadata, error) {
	relRepo := s.releaseRepo()

	rs, err := relRepo.List(ns, pkg)
	if err != nil {
		return nil, nil
	}

	var rms []ReleaseMetadata
	for _, r := range rs {
		rms = append(rms, convertRepoReleaseToMetadata(r))
	}

	return rms, nil
}

// CreateRelease creates a release in the store. It returns the release metdata or an error.
func (s *S3) CreateRelease(nsName, pkgName, release string, data []byte) (ReleaseMetadata, error) {
	_, err := s.createOrRetrievePackage(nsName, pkgName)
	if err != nil {
		return ReleaseMetadata{}, err
	}

	digest := makeDigest(data)

	_, err = s.releaseRepo().RetrieveDigest(digest)
	if err == nil {
		return ReleaseMetadata{}, errors.New("release with digest already exists")
	}

	_, err = s.releaseRepo().Retrieve(nsName, pkgName, release)
	if err == nil {
		return ReleaseMetadata{}, errors.New("release already exists")
	}

	tgz, err := newTarGz(s.fs)
	if err != nil {
		return ReleaseMetadata{}, err
	}
	defer tgz.close()

	bReader := bytes.NewReader(data)
	if err = tgz.extractTarGz(bReader); err != nil {
		return ReleaseMetadata{}, errors.Wrap(err, "blob was not a gzip'd tar file")
	}

	n, err := s.storePath(digest, blobName, blobMIMEType, bytes.NewReader(data), int64(-1))
	if err != nil {
		return ReleaseMetadata{}, err
	}

	configData, err := tgz.config()
	if err != nil {
		return ReleaseMetadata{}, err
	}
	_, err = s.storePath(digest, configName, configMIMEType, bytes.NewReader(configData), -1)
	if err != nil {
		return ReleaseMetadata{}, err
	}

	readmeData, err := tgz.readme()
	if err != nil {
		return ReleaseMetadata{}, err
	}
	_, err = s.storePath(digest, docName, docMIMEType, bytes.NewReader(readmeData), -1)
	if err != nil {
		return ReleaseMetadata{}, err
	}

	// associate the digest with release
	r, err := s.releaseRepo().Create(nsName, pkgName, release, digest, n, time.Now().UTC())
	if err != nil {
		return ReleaseMetadata{}, err
	}

	rm := ReleaseMetadata{
		ID:           r.ID,
		Namespace:    r.Namespace,
		Package:      r.Package,
		Digest:       r.Digest,
		Size:         r.Size,
		CreatedAt:    r.CreatedAt,
		Version:      r.Version,
		Dependencies: r.Dependencies,
	}

	return rm, nil
}

func (s *S3) storePath(digest, name, contentType string, r io.Reader, size int64) (int, error) {
	objectName := fmt.Sprintf("%s/%s", digest, name)
	return s.client.PutObject(objectName, contentType, r, size)
}

func (s *S3) createOrRetrievePackage(nsName, pkgName string) (repository.Package, error) {
	pkgRepo := s.pkgRepo()

	pkg, err := pkgRepo.Retrieve(nsName, pkgName)
	if err != nil {
		switch t := err.(type) {
		case repository.Error:
			if t.IsNotExist() {
				return pkgRepo.Create(nsName, pkgName)
			}
		default:
			return repository.Package{}, err
		}
	}

	return pkg, nil
}

// RemoveRelease removes a release from the store.
func (s *S3) RemoveRelease(ns, pkg, version string) error {
	release, err := s.releaseRepo().Retrieve(ns, pkg, version)
	if err != nil {
		return err
	}

	if err = s.releaseRepo().Delete(ns, pkg, version); err != nil {
		return err
	}

	objects := []string{blobName, configName, docName}
	for _, obj := range objects {
		path := fmt.Sprintf("%s/%s", release.Digest, obj)
		if err = s.client.RemoveObject(path); err != nil {
		}
	}

	if err != nil {
		return errors.Wrap(err, "delete objects from store")
	}

	return nil
}

// Release returns a release.
func (s *S3) Release(ns, pkg, version string) (ReleaseMetadata, error) {
	release, err := s.releaseRepo().Retrieve(ns, pkg, version)
	if err != nil {
		return ReleaseMetadata{}, err
	}

	rm := ReleaseMetadata{
		ID:           release.ID,
		Namespace:    release.Namespace,
		Package:      release.Package,
		Digest:       release.Digest,
		Size:         release.Size,
		CreatedAt:    release.CreatedAt,
		Version:      release.Version,
		Dependencies: release.Dependencies,
	}

	return rm, nil
}

// Pull pulls a release from the repository.
func (s *S3) Pull(ns, pkg, digest string) (multipart.File, error) {
	release, err := s.releaseRepo().RetrieveDigest(digest)
	if err != nil {
		return nil, err
	}

	objectName := fmt.Sprintf("%s/%s", release.Digest, blobName)
	o, err := s.client.GetObject(objectName)
	if err != nil {
		return nil, err
	}

	return o, nil
}

func (s *S3) Close() error {
	return nil
}

func convertRepoPackageToMetadata(in repository.Package) PackageMetadata {
	return PackageMetadata{
		ID:        in.ID,
		Namespace: in.Namespace,
		Package:   in.Package,
		CreatedAt: in.CreatedAt,
		IsVisible: in.Visibility,
	}
}

func convertRepoReleaseToMetadata(in repository.Release) ReleaseMetadata {
	return ReleaseMetadata{
		ID:           in.ID,
		Namespace:    in.Namespace,
		Package:      in.Package,
		Digest:       in.Digest,
		Size:         in.Size,
		CreatedAt:    in.CreatedAt,
		Version:      in.Version,
		Dependencies: in.Dependencies,
	}
}
