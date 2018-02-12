package repository

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Error interface {
	IsNotExist() bool
}

type Repository interface {
	Namespace() NamespaceRepository
	Package() PackageRepository
	Release() ReleaseRepository
}

type NamespaceRepository interface {
	Create(ns string) (Namespace, error)
	List() ([]Namespace, error)
}

type PackageRepository interface {
	Create(ns, pkg string) (Package, error)
	List(ns string) ([]Package, error)
	Retrieve(ns, pkg string) (Package, error)
}

type ReleaseRepository interface {
	Create(ns, pkg, release, digest string, size int, createdAt time.Time) (Release, error)
	Delete(ns, pkg, release string) error
	List(ns, pkg string) ([]Release, error)
	Retrieve(ns, pkg, release string) (Release, error)
	RetrieveDigest(digest string) (Release, error)
}

type Namespace struct {
	ID        uuid.UUID
	Namespace string
}

type Package struct {
	ID         uuid.UUID
	Namespace  string
	Package    string
	CreatedAt  time.Time
	Visibility bool
}

type Release struct {
	ID           uuid.UUID
	Namespace    string
	Package      string
	Digest       string
	Size         int64
	CreatedAt    time.Time
	Version      string
	Dependencies Dependencies
}

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
