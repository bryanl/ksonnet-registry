package registry

import "github.com/bryanl/ksonnet-registry/store"

var (
	namespaces map[string]*Namespace = make(map[string]*Namespace)
)

func GetNamespace(s store.Store, name string) (*Namespace, error) {
	n, ok := namespaces[name]
	if !ok {
		n = NewNamespace(s, name)
		namespaces[name] = n
	}

	return n, nil
}

type Namespace struct {
	Name     string
	packages map[string]*Package
	store    store.Store
}

func NewNamespace(s store.Store, name string) *Namespace {
	return &Namespace{
		Name:     name,
		packages: make(map[string]*Package),
		store:    s,
	}
}

func (n *Namespace) Package(name string) (*Package, error) {
	p, ok := n.packages[name]
	if !ok {
		p = NewPackage(n.store, n.Name, name)
		n.packages[name] = p
	}

	return p, nil
}
