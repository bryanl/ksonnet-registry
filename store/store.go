package store

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
	yaml "gopkg.in/yaml.v2"
)

const (
	blobName   = "part.tar.gz"
	configName = "parts.yaml"
	docName    = "README.md"
)

var (
	dirMode  os.FileMode = 0755
	fileMode os.FileMode = 0644

	// ErrNotFound is a not found error.
	ErrNotFound = errors.New("not found")
)

// ReleaseMetadata contains release metadata.
type ReleaseMetadata struct {
	Digest       string
	Size         int64
	CreatedAt    time.Time
	Version      string
	Dependencies Dependencies
}

// Dependency specifies a release dependency.
type Dependency struct {
	Name       string
	Constraint string
}

// Dependencies are a slice of Dependency.
type Dependencies []Dependency

// ToMap converts the Dependencies to a map.
func (ds Dependencies) ToMap() map[string]string {
	m := make(map[string]string)

	for _, d := range ds {
		m[d.Name] = d.Constraint
	}

	return m
}

// Store manages files
type Store interface {
	Namespaces() ([]string, error)
	Packages(ns string) ([]string, error)
	Releases(ns string, pkg string) ([]ReleaseMetadata, error)
	CreateRelease(ns, pkg, release string, data []byte) (ReleaseMetadata, error)
	RemoveRelease(ns, pkg, release string) error
	Release(ns, pkg, release string) (ReleaseMetadata, error)
	Pull(ns, pkg, digest string) (multipart.File, error)

	Close() error
}

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

// Namespaces returns a list of namespaces in the store.
func (s *FileSystemStore) Namespaces() ([]string, error) {
	files, err := afero.ReadDir(s.fs, s.dir)
	if err != nil {
		return nil, err
	}

	var namespaces []string

	for _, file := range files {
		if file.IsDir() {
			namespaces = append(namespaces, file.Name())
		}
	}

	return namespaces, nil
}

// Packages returns packages in a namespace.
func (s *FileSystemStore) Packages(ns string) ([]string, error) {
	dir := filepath.Join(s.dir, ns)

	if _, err := s.fs.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			return make([]string, 0), nil
		}

		return nil, err
	}

	files, err := afero.ReadDir(s.fs, dir)
	if err != nil {
		return nil, err
	}

	var packages []string

	for _, file := range files {
		if file.IsDir() {
			packages = append(packages, file.Name())
		}
	}

	return packages, nil
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

	d := digest(data)
	digestDir := filepath.Join(s.dir, ns, pkg, "digests", d)
	releaseDir := filepath.Join(s.dir, ns, pkg, "releases")

	if _, err := s.fs.Stat(digestDir); err == nil {
		return rm, errors.Errorf("digest %q already exists", d)
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

	tmpDir, err := afero.TempDir(s.fs, "", "extract-part")
	if err != nil {
		return rm, err
	}
	defer s.fs.RemoveAll(tmpDir)

	r := bytes.NewReader(data)
	if err = s.extractTarGz(tmpDir, r); err != nil {
		return rm, errors.Wrap(err, "blob was not a gzip'd tar file")
	}

	if err = afero.WriteFile(s.fs, partData, data, fileMode); err != nil {
		return rm, err
	}

	if err = s.copyFile(filepath.Join(tmpDir, "parts.yaml"), partConfig); err != nil {
		return rm, errors.New("part is missing parts.yaml")
	}

	if err = s.copyFile(filepath.Join(tmpDir, "README.md"), partDoc); err != nil {
		return rm, errors.New("part is missing README.md")
	}

	if err = s.fs.MkdirAll(releaseDir, dirMode); err != nil {
		return rm, err
	}

	if err = afero.WriteFile(s.fs, filepath.Join(releaseDir, release), []byte(d), fileMode); err != nil {
		return rm, err
	}

	fi, err := s.fs.Stat(partData)
	if err != nil {
		return rm, err
	}

	rm.Digest = d
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
		dep := Dependency{
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
			return nil, ErrNotFound
		}

		return nil, err
	}

	return s.fs.Open(blob)
}

// Close closes the TempStore.
func (s *FileSystemStore) Close() error {
	return s.closeFn()
}

func digest(data []byte) string {
	sum := sha256.Sum256(data)
	return fmt.Sprintf("%x", sum)
}

func (s *FileSystemStore) extractTarGz(dest string, r io.Reader) error {
	gzr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}

	defer gzr.Close()

	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()

		switch {

		// no more files
		case err == io.EOF:
			return nil

		// unknown error
		case err != nil:
			return err

		// no header
		case header == nil:
			continue
		}

		target := filepath.Join(dest, header.Name)

		switch header.Typeflag {

		// ensure dir exists
		case tar.TypeDir:
			if _, err := s.fs.Stat(target); err != nil {
				if err := s.fs.MkdirAll(target, dirMode); err != nil {
					return err
				}
			}

		case tar.TypeReg:
			f, err := s.fs.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			defer f.Close()

			if _, err := io.Copy(f, tr); err != nil {
				return err
			}
		}
	}
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
