package store

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/bryanl/ksonnet-registry/repository"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	yaml "gopkg.in/yaml.v2"
)

// FileSystemStoreOpt is a configuration option for FileSystemStore.
type FileSystemStoreOpt func(*FileSystemStore)

// FileSystemStoreOptFS configures the afero Fs for the the FileSystemStore.
func FileSystemStoreOptFS(fs afero.Fs) FileSystemStoreOpt {
	return func(s *FileSystemStore) {
		s.fs = fs
	}
}

// FileSystemStoreOptClose sets the close function for FileSystemStore.
func FileSystemStoreOptClose(fn func() error) FileSystemStoreOpt {
	return func(s *FileSystemStore) {
		s.closeFn = fn
	}
}

// FileSystemStoreOptRoot sets the dir for the FileSystemStore.
func FileSystemStoreOptRoot(dir string) FileSystemStoreOpt {
	return func(s *FileSystemStore) {
		s.dir = dir
	}
}

// FileSystemStore stores the registry on the file system.
type FileSystemStore struct {
	dir     string
	fs      afero.Fs
	closeFn func() error
}

var _ Store = (*FileSystemStore)(nil)

// NewFileSystemStore creates an instance of FileSystemStore.
func NewFileSystemStore(opts ...FileSystemStoreOpt) (*FileSystemStore, error) {
	ts := &FileSystemStore{
		dir:     "/data",
		fs:      afero.NewOsFs(),
		closeFn: func() error { return nil },
	}

	for _, opt := range opts {
		opt(ts)
	}

	if _, err := ts.fs.Stat(ts.dir); err != nil {
		return nil, err
	}

	return ts, nil
}

// CreateNamespace creates a namespace.
func (s *FileSystemStore) CreateNamespace(ns string) (NamespaceMetdata, error) {
	return NamespaceMetdata{
		Namespace: ns,
	}, nil
}

// Namespaces returns a list of namespaces in the store.
func (s *FileSystemStore) Namespaces() ([]NamespaceMetdata, error) {
	files, err := afero.ReadDir(s.fs, s.dir)
	if err != nil {
		return nil, err
	}

	var namespaces []NamespaceMetdata

	for _, file := range files {
		if file.IsDir() {
			nm := NamespaceMetdata{Namespace: file.Name()}
			namespaces = append(namespaces, nm)
		}
	}

	return namespaces, nil
}

// CreatePackage creates a package.
func (s *FileSystemStore) CreatePackage(ns, pkg string) (PackageMetadata, error) {

	pkgDir := filepath.Join(s.dir, ns, pkg)

	if fileExists(s.fs, pkgDir) {
		return PackageMetadata{}, errors.Errorf("package %s/%s exists", ns, pkg)
	}

	if err := s.fs.MkdirAll(pkgDir, dirMode); err != nil {
		return PackageMetadata{}, err
	}

	pm := PackageMetadata{
		Namespace: ns,
		Package:   pkg,
		CreatedAt: time.Now().UTC(),
		IsVisible: true,
	}

	b, err := yaml.Marshal(&pm)
	if err != nil {
		return PackageMetadata{}, err
	}

	pmFile := filepath.Join(pkgDir, pkgMetadataName)
	if err := afero.WriteFile(s.fs, pmFile, b, dirMode); err != nil {
		return PackageMetadata{}, err
	}

	return pm, nil
}

// Packages returns packages in a namespace.
func (s *FileSystemStore) Packages(ns string) ([]PackageMetadata, error) {
	dir := filepath.Join(s.dir, ns)

	if _, err := s.fs.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			return make([]PackageMetadata, 0), nil
		}

		return nil, err
	}

	files, err := afero.ReadDir(s.fs, dir)
	if err != nil {
		return nil, err
	}

	var packages []PackageMetadata

	for _, file := range files {
		if file.IsDir() {
			pm, err := s.Package(ns, file.Name())
			if err != nil {
				return nil, err
			}
			packages = append(packages, pm)
		}
	}

	return packages, nil
}

// Package returns a package by name.
func (s *FileSystemStore) Package(ns, pkg string) (PackageMetadata, error) {
	dir := s.pkgDir(ns, pkg)

	if _, err := s.fs.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			return PackageMetadata{}, NewNotFoundError(ns + "/" + pkg)
		}

		return PackageMetadata{}, err
	}

	b, err := afero.ReadFile(s.fs, s.pkgMetadata(ns, pkg))
	if err != nil {
		return PackageMetadata{}, err
	}

	var pm PackageMetadata
	if err := yaml.Unmarshal(b, &pm); err != nil {
		return PackageMetadata{}, err
	}

	return pm, nil
}

// Releases returns releases in a package.
func (s *FileSystemStore) Releases(ns string, pkg string) ([]ReleaseMetadata, error) {
	dir := filepath.Join(s.dir, ns, pkg, "releases")

	if _, err := s.fs.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			return make([]ReleaseMetadata, 0), nil
		}

		return nil, err
	}

	files, err := afero.ReadDir(s.fs, dir)
	if err != nil {
		return nil, err
	}

	var releases []ReleaseMetadata

	for _, file := range files {
		if !file.IsDir() {
			rm, err := s.Release(ns, pkg, file.Name())
			if err != nil {
				return nil, err
			}

			releases = append(releases, rm)
		}
	}

	return releases, nil
}

// CreateRelease creates a release in the store. It returns the release metadata or an error.
func (s *FileSystemStore) CreateRelease(ns, pkg, release string, data []byte) (ReleaseMetadata, error) {
	var rm ReleaseMetadata

	if !fileExists(s.fs, s.pkgDir(ns, pkg)) {
		_, err := s.CreatePackage(ns, pkg)
		if err != nil {
			return ReleaseMetadata{}, errors.Wrapf(err, "create package %s/%s", ns, pkg)
		}
	}

	digest := makeDigest(data)
	digestDir := filepath.Join(s.dir, ns, pkg, "digests", digest)
	releaseDir := filepath.Join(s.dir, ns, pkg, "releases")

	if _, err := s.fs.Stat(digestDir); err == nil {
		return rm, errors.Errorf("digest %q already exists", digest)
	}

	if _, err := s.fs.Stat(filepath.Join(releaseDir, release)); err == nil {
		return rm, errors.Errorf("release %q already exists", release)
	}

	if err := s.fs.MkdirAll(digestDir, dirMode); err != nil {
		return rm, err
	}

	partData := filepath.Join(digestDir, blobName)
	partConfig := filepath.Join(digestDir, configName)
	partDoc := filepath.Join(digestDir, docName)

	if err := s.fs.MkdirAll(releaseDir, dirMode); err != nil {
		return rm, err
	}

	tgz, err := newTarGz(s.fs)
	if err != nil {
		return rm, err
	}
	defer tgz.close()

	r := bytes.NewReader(data)
	if err = tgz.extractTarGz(r); err != nil {
		return rm, errors.Wrap(err, "blob was not a gzip'd tar file")
	}

	if err = s.writeReleaseFile(digestDir, partConfig, tgz.config); err != nil {
		return rm, err
	}

	if err = s.writeReleaseFile(digestDir, partDoc, tgz.readme); err != nil {
		return rm, err
	}

	if err = s.writeReleaseFile(digestDir, partData, func() ([]byte, error) { return data, nil }); err != nil {
		return rm, err
	}

	fi, err := s.fs.Stat(partData)
	if err != nil {
		return rm, err
	}

	if err := afero.WriteFile(s.fs, filepath.Join(releaseDir, release), []byte(digest), fileMode); err != nil {
		return rm, err
	}

	rm.Digest = digest
	rm.CreatedAt = fi.ModTime()
	rm.Size = fi.Size()
	rm.Version = release

	return rm, nil
}

// RemoveRelease removes a release.
func (s *FileSystemStore) RemoveRelease(ns, pkg, ver string) error {
	releaseName := filepath.Join(s.dir, ns, pkg, "releases", ver)

	if _, err := s.fs.Stat(releaseName); err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		return err
	}

	digest, err := afero.ReadFile(s.fs, releaseName)
	if err != nil {
		return err
	}

	if err := s.fs.RemoveAll(releaseName); err != nil {
		return err
	}

	digestPath := filepath.Join(s.dir, ns, pkg, "digests", string(digest))
	if _, err := s.fs.Stat(digestPath); err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		return err
	}

	return s.fs.RemoveAll(digestPath)
}

// Release returns ReleaseMetdata for a release or an error.
func (s *FileSystemStore) Release(ns, pkg, release string) (ReleaseMetadata, error) {
	var rm ReleaseMetadata

	releaseName := filepath.Join(s.dir, ns, pkg, "releases", release)

	if _, err := s.fs.Stat(releaseName); err != nil {
		if os.IsNotExist(err) {
			return rm, errors.Errorf("release %q does not exist", release)
		}

		return rm, err
	}

	b, err := afero.ReadFile(s.fs, releaseName)
	if err != nil {
		return rm, err
	}

	rm.Digest = string(b)

	digestPath := filepath.Join(s.dir, ns, pkg, "digests", string(rm.Digest))
	blobPath := filepath.Join(digestPath, blobName)

	fi, err := s.fs.Stat(blobPath)
	if err != nil {
		return rm, err
	}

	// open parts.yaml
	b, err = afero.ReadFile(s.fs, filepath.Join(digestPath, "parts.yaml"))
	if err != nil {
		return rm, err
	}

	var pc partsConfig
	if err := yaml.Unmarshal(b, &pc); err != nil {
		return rm, err
	}

	rm.Size = fi.Size()
	rm.CreatedAt = fi.ModTime()
	rm.Version = release

	for k, v := range pc.Dependencies {
		dep := repository.Dependency{
			Name:       k,
			Constraint: v,
		}

		rm.Dependencies = append(rm.Dependencies, dep)
	}

	return rm, nil
}

// Pull pulls a digest from a package.
func (s *FileSystemStore) Pull(ns, pkg, digest string) (multipart.File, error) {
	digestDir := filepath.Join(s.dir, ns, pkg, "digests", digest)

	blob := filepath.Join(digestDir, blobName)
	if _, err := s.fs.Stat(blob); err != nil {
		if os.IsNotExist(err) {
			return nil, NewNotFoundError(fmt.Sprintf("digest %s in %s/%s", digest, ns, pkg))
		}

		return nil, err
	}

	return s.fs.Open(blob)
}

// Close closes the TempStore.
func (s *FileSystemStore) Close() error {
	return s.closeFn()
}

func (s *FileSystemStore) copyFile(src, dest string) error {
	from, err := s.fs.Open(src)
	if err != nil {
		return err
	}
	defer from.Close()

	to, err := s.fs.OpenFile(dest, os.O_RDWR|os.O_CREATE, fileMode)
	if err != nil {
		return err
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	return err
}

func fileExists(fs afero.Fs, name string) bool {
	_, err := fs.Stat(name)
	return err == nil
}

func (s *FileSystemStore) nsDir(ns string) string {
	return filepath.Join(s.dir, ns)
}

func (s *FileSystemStore) pkgDir(ns, pkg string) string {
	return filepath.Join(s.nsDir(ns), pkg)
}

func (s *FileSystemStore) pkgMetadata(ns, pkg string) string {
	return filepath.Join(s.pkgDir(ns, pkg), pkgMetadataName)
}

func (s *FileSystemStore) writeReleaseFile(releaseDir, name string, fn func() ([]byte, error)) error {
	data, err := fn()
	if err != nil {
		return err
	}

	return afero.WriteFile(s.fs, name, data, fileMode)
}
