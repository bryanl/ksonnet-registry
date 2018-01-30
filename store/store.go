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

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
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

// Store manages files
type Store interface {
	Namespaces() ([]string, error)
	Packages(ns string) ([]string, error)
	Releases(ns string, pkg string) ([]string, error)
	CreateRelease(ns, pkg, release string, data []byte) (string, error)
	RemoveRelease(ns, pkg, release string) error
	Digest(ns, pkg, release string) (string, error)
	Pull(ns, pkg, digest string) (multipart.File, error)

	Close() error
}

// TempStoreOpt is a configuration option for TempStore.
type TempStoreOpt func(*TempStore)

// TempStoreOptFS configures the afero Fs for the the TempStore.
func TempStoreOptFS(fs afero.Fs) TempStoreOpt {
	return func(s *TempStore) {
		s.fs = fs
	}
}

// TempStore stores the registry in a temp directory.
type TempStore struct {
	dir string
	fs  afero.Fs
}

var _ Store = (*TempStore)(nil)

// NewTempStore creates an instance of TempStore.
func NewTempStore(opts ...TempStoreOpt) (*TempStore, error) {
	ts := &TempStore{
		fs: afero.NewOsFs(),
	}

	for _, opt := range opts {
		opt(ts)
	}

	dir, err := afero.TempDir(ts.fs, "", "ksonnet-registry")
	if err != nil {
		return nil, err
	}

	ts.dir = dir

	logrus.WithField("root", dir).Info("initialized store")

	return ts, nil
}

// Namespaces returns a list of namespaces in the store.
func (s *TempStore) Namespaces() ([]string, error) {
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
func (s *TempStore) Packages(ns string) ([]string, error) {
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
func (s *TempStore) Releases(ns string, pkg string) ([]string, error) {
	dir := filepath.Join(s.dir, ns, pkg, "releases")

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

	var releases []string

	for _, file := range files {
		if !file.IsDir() {
			releases = append(releases, file.Name())
		}
	}

	return releases, nil
}

// CreateRelease creates a release in the store. It returns the digest or an error.
func (s *TempStore) CreateRelease(ns, pkg, release string, data []byte) (string, error) {
	d := digest(data)
	digestDir := filepath.Join(s.dir, ns, pkg, "digests", d)
	releaseDir := filepath.Join(s.dir, ns, pkg, "releases")

	if _, err := s.fs.Stat(digestDir); err == nil {
		return "", errors.Errorf("digest %q already exists", d)
	}

	if _, err := s.fs.Stat(filepath.Join(releaseDir, release)); err == nil {
		return "", errors.Errorf("release %q already exists", release)
	}

	if err := s.fs.MkdirAll(digestDir, dirMode); err != nil {
		return "", err
	}

	partData := filepath.Join(digestDir, blobName)
	partConfig := filepath.Join(digestDir, configName)
	partDoc := filepath.Join(digestDir, docName)

	tmpDir, err := afero.TempDir(s.fs, "", "extract-part")
	if err != nil {
		return "", err
	}
	defer s.fs.RemoveAll(tmpDir)

	r := bytes.NewReader(data)
	if err := s.extractTarGz(tmpDir, r); err != nil {
		return "", errors.Wrap(err, "blob was not a gzip'd tar file")
	}

	if err := afero.WriteFile(s.fs, partData, data, fileMode); err != nil {
		return "", err
	}

	if err := s.copyFile(filepath.Join(tmpDir, "parts.yaml"), partConfig); err != nil {
		return "", err
	}

	if err := s.copyFile(filepath.Join(tmpDir, "README.md"), partDoc); err != nil {
		return "", err
	}

	if err := s.fs.MkdirAll(releaseDir, dirMode); err != nil {
		return "", err
	}

	if err := afero.WriteFile(s.fs, filepath.Join(releaseDir, release), []byte(d), fileMode); err != nil {
		return "", err
	}

	return d, nil
}

// RemoveRelease removes a release.
func (s *TempStore) RemoveRelease(ns, pkg, ver string) error {
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

// Digest returns the digest for a release.
func (s *TempStore) Digest(ns, pkg, release string) (string, error) {
	releaseName := filepath.Join(s.dir, ns, pkg, "releases", release)

	if _, err := s.fs.Stat(releaseName); err != nil {
		if os.IsNotExist(err) {
			return "", errors.Errorf("release %q does not exist", release)
		}

		return "", err
	}

	b, err := afero.ReadFile(s.fs, releaseName)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// Pull pulls a digest from a package.
func (s *TempStore) Pull(ns, pkg, digest string) (multipart.File, error) {
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
func (s *TempStore) Close() error {
	logrus.Info("removing store")
	return s.fs.RemoveAll(s.dir)
}

func digest(data []byte) string {
	sum := sha256.Sum256(data)
	return fmt.Sprintf("%x", sum)
}

func (s *TempStore) extractTarGz(dest string, r io.Reader) error {
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

func (s *TempStore) copyFile(src, dest string) error {
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
