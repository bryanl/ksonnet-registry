package registry

import (
	"github.com/bryanl/ksonnet-registry/store"
)

// GetNamespace returns a namespace by name.
func GetNamespace(s store.Store, name string) (*Namespace, error) {
	ns := NewNamespace(s, name)
	return ns, nil
}

// Namespace is a ksonnet registry namespace.
type Namespace struct {
	Name  string
	store store.Store
}

// NewNamespace creates an instance of Namespace.
func NewNamespace(s store.Store, name string) *Namespace {
	return &Namespace{
		Name:  name,
		store: s,
	}
}

// Package returns a package from a namespace.
func (n *Namespace) Package(pm store.PackageMetadata) (*Package, error) {
	p := NewPackage(n.store, pm)
	return p, nil
}

// PackageByName retrieves a package by name.
func (n *Namespace) PackageByName(name string) (*Package, error) {
	pm, err := n.store.Package(n.Name, name)
	if err != nil {
		return nil, err
	}

	return n.Package(pm)
}
