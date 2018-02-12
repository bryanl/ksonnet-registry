package store

import (
	"crypto/sha256"
	"fmt"
	"mime/multipart"
	"os"
	"time"

	"github.com/bryanl/ksonnet-registry/repository"
	uuid "github.com/satori/go.uuid"
)

const (
	pkgMetadataName = "metadata.yaml"
	blobName        = "part.tar.gz"
	blobMIMEType    = "application/x-gzip"
	configName      = "parts.yaml"
	configMIMEType  = "application/x-yaml"
	docName         = "README.md"
	docMIMEType     = "text/markdown"
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

// NamespaceMetdata is namespace metdata.
type NamespaceMetdata struct {
	ID        uuid.UUID `yaml:"id"`
	Namespace string    `yaml:"namespace"`
}

// PackageMetadata contains package metadata.
type PackageMetadata struct {
	ID        uuid.UUID `yaml:"id"`
	Namespace string    `yaml:"namespace"`
	Package   string    `yaml:"package"`
	CreatedAt time.Time `yaml:"createdAt"`
	IsVisible bool      `yaml:"isVisible"`
}

// ReleaseMetadata contains release metadata.
type ReleaseMetadata struct {
	ID           uuid.UUID `yaml:"id"`
	Namespace    string
	Package      string
	Digest       string
	Size         int64
	CreatedAt    time.Time
	Version      string
	Dependencies repository.Dependencies
}

// Store manages files
type Store interface {
	CreateNamespace(ns string) (NamespaceMetdata, error)
	Namespaces() ([]NamespaceMetdata, error)
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

func makeDigest(data []byte) string {
	sum := sha256.Sum256(data)
	return fmt.Sprintf("%x", sum)
}
