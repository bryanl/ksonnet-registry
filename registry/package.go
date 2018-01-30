package registry

import (
	"io/ioutil"
	"mime/multipart"
	"sync"

	"github.com/bryanl/ksonnet-registry/store"
	"github.com/pkg/errors"
)

type Package struct {
	Name      string
	Namespace string
	releases  map[string]*Release
	store     store.Store

	mu sync.Mutex
}

func NewPackage(s store.Store, ns, name string) *Package {
	return &Package{
		Name:      name,
		Namespace: ns,
		releases:  make(map[string]*Release),
		store:     s,
	}
}

func (p *Package) CreateRelease(release string, data []byte) (*Release, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	_, ok := p.releases[release]
	if ok {
		return nil, errors.Errorf("release %q already exists", release)
	}

	r, err := NewRelease(p.store, p.Namespace, p.Name, release, data)
	if err != nil {
		return nil, errors.Wrapf(err, "create release %q", release)
	}
	p.releases[release] = r

	return r, nil
}

func (p *Package) Release(ver string) (*Release, error) {
	r, ok := p.releases[ver]
	if !ok {
		return nil, errors.Errorf("release %s was not found", ver)
	}

	return r, nil
}

func (p *Package) Delete(ver string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	r, ok := p.releases[ver]
	if !ok {
		return errors.Errorf("release %s was not found", ver)
	}

	if err := r.Delete(); err != nil {
		return errors.Wrapf(err, "unable to delete version %s", ver)
	}

	delete(p.releases, ver)

	return nil
}

func (p *Package) Pull(digest string) (multipart.File, *multipart.FileHeader, error) {
	data, err := p.store.Read(digest)
	if err != nil {
		return nil, nil, err
	}

	// TODO: (bryanl) - when will this file be deleted?
	tmpFile, err := ioutil.TempFile("", digest)
	if err != nil {
		return nil, nil, err
	}

	if _, err := tmpFile.Write(data); err != nil {
		return nil, nil, err
	}

	return tmpFile, nil
}
