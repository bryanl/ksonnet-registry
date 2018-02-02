package store

import (
	"fmt"
	"mime/multipart"
	"os"
	"time"
)

const (
	pkgMetadataName = "metadata.yaml"
	blobName        = "part.tar.gz"
	configName      = "parts.yaml"
	docName         = "README.md"
)

var (
	dirMode  os.FileMode = 0755
	fileMode os.FileMode = 0644
)

// NotFoundError is a not found error.
type NotFoundError struct {
	item string
}

// NewNotFoundError creates a NotFoundError.
func NewNotFoundError(item string) *NotFoundError {
	return &NotFoundError{item: item}
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%q was not found", e.item)
}

// PackageMetadata contains package metadata.
type PackageMetadata struct {
	Namespace string    `yaml:"namespace"`
	Package   string    `yaml:"package"`
	CreatedAt time.Time `yaml:"createdAt"`
	IsVisible bool      `yaml:"isVisible"`
}

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
	CreatePackage(ns, pkg string) (PackageMetadata, error)
	Packages(ns string) ([]PackageMetadata, error)
	Package(ns, pkg string) (PackageMetadata, error)
	Releases(ns string, pkg string) ([]ReleaseMetadata, error)
	CreateRelease(ns, pkg, release string, data []byte) (ReleaseMetadata, error)
	RemoveRelease(ns, pkg, release string) error
	Release(ns, pkg, release string) (ReleaseMetadata, error)
	Pull(ns, pkg, digest string) (multipart.File, error)

	Close() error
}
