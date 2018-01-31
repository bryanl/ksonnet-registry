package registry

import (
	"fmt"
	"mime/multipart"

	"github.com/bryanl/ksonnet-registry/store"
	"github.com/pkg/errors"
)

// Package is a package in a ksonnet registry namespace.
type Package struct {
	Name      string
	Namespace string
	store     store.Store
}

// NewPackage creates an instance of Package.
func NewPackage(s store.Store, ns, name string) *Package {
	return &Package{
		Name:      name,
		Namespace: ns,
		store:     s,
	}
}

// CreateRelease creates a new release version.
func (p *Package) CreateRelease(version string, data []byte) (*Release, error) {
	releases, err := p.store.Releases(p.Namespace, p.Name)
	if err != nil {
		return nil, err
	}

	for _, rm := range releases {
		if version == rm.Version {
			return nil, errors.Errorf("release %q already exists", version)
		}
	}

	rm, err := p.store.CreateRelease(p.Namespace, p.Name, version, data)
	if err != nil {
		return nil, err
	}

	r := NewRelease(p.store, p.Namespace, p.Name, version, rm.Digest, rm.CreatedAt, rm.Size,
		rm.Dependencies.ToMap())
	return r, nil
}

// Release returns a release by version.
func (p *Package) Release(ver string) (*Release, error) {
	releases, err := p.store.Releases(p.Namespace, p.Name)
	if err != nil {
		return nil, err
	}

	for _, rm := range releases {
		if rm.Version == ver {
			rm, err := p.store.Release(p.Namespace, p.Name, ver)
			if err != nil {
				return nil, err
			}

			r := NewRelease(p.store, p.Namespace, p.Name, ver, rm.Digest, rm.CreatedAt, rm.Size,
				rm.Dependencies.ToMap())
			return r, nil
		}
	}

	return nil, errors.Errorf("release %s was not found", ver)
}

// Releases returns all releases.
func (p *Package) Releases() ([]Release, error) {
	versions, err := p.store.Releases(p.Namespace, p.Name)
	if err != nil {
		return nil, err
	}

	var releases []Release

	for _, rm := range versions {
		r := NewRelease(p.store, p.Namespace, p.Name, rm.Version, rm.Digest, rm.CreatedAt, rm.Size,
			rm.Dependencies.ToMap())
		releases = append(releases, *r)
	}

	return releases, nil
}

// Delete deletes a version from a package.
func (p *Package) Delete(ver string) error {
	r, err := p.Release(ver)
	if err != nil {
		return err
	}

	if err := r.Delete(); err != nil {
		return errors.Wrapf(err, "unable to delete version %s", ver)
	}

	return nil
}

// Pull pulls a digest from the package.
func (p *Package) Pull(digest string) (multipart.File, *multipart.FileHeader, error) {
	f, err := p.store.Pull(p.Namespace, p.Name, digest)
	if err != nil {
		return nil, nil, err
	}

	hdr := &multipart.FileHeader{
		Filename: fmt.Sprintf("%s.tar.gz", digest),
	}

	return f, hdr, nil
}
