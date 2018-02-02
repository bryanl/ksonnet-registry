package registry

import (
	"fmt"
	"mime/multipart"
	"time"

	"github.com/bryanl/ksonnet-registry/store"
	"github.com/pkg/errors"
)

// Package is a package in a ksonnet registry namespace.
type Package struct {
	Package   string
	Namespace string
	CreatedAt time.Time
	IsVisible bool
	store     store.Store
}

// NewPackage creates an instance of Package.
func NewPackage(s store.Store, pm store.PackageMetadata) *Package {
	return &Package{
		Package:   pm.Package,
		Namespace: pm.Namespace,
		CreatedAt: pm.CreatedAt,
		IsVisible: pm.IsVisible,
		store:     s,
	}
}

// Name is the fully qualified name of a package. e.g. namespace/package
func (p *Package) Name() string {
	return fmt.Sprintf("%s/%s", p.Namespace, p.Package)
}

// Visibility is the visibility of the package.
func (p *Package) Visibility() string {
	if p.IsVisible {
		return "public"
	}

	return "private"
}

// CreateRelease creates a new release version.
func (p *Package) CreateRelease(version string, data []byte) (*Release, error) {
	releases, err := p.store.Releases(p.Namespace, p.Package)
	if err != nil {
		return nil, err
	}

	for _, rm := range releases {
		if version == rm.Version {
			return nil, errors.Errorf("release %q already exists", version)
		}
	}

	rm, err := p.store.CreateRelease(p.Namespace, p.Package, version, data)
	if err != nil {
		return nil, err
	}

	r := NewRelease(p.store, p.Namespace, p.Package, version, rm.Digest, rm.CreatedAt, rm.Size,
		rm.Dependencies.ToMap())
	return r, nil
}

// Release returns a release by version.
func (p *Package) Release(ver string) (*Release, error) {
	releases, err := p.store.Releases(p.Namespace, p.Package)
	if err != nil {
		return nil, err
	}

	for _, rm := range releases {
		if rm.Version == ver {
			rm, err := p.store.Release(p.Namespace, p.Package, ver)
			if err != nil {
				return nil, err
			}

			r := NewRelease(p.store, p.Namespace, p.Package, ver, rm.Digest, rm.CreatedAt, rm.Size,
				rm.Dependencies.ToMap())
			return r, nil
		}
	}

	return nil, errors.Errorf("release %s in %s/%s was not found", ver, p.Namespace, p.Package)
}

// Releases returns all releases.
func (p *Package) Releases() ([]Release, error) {
	versions, err := p.store.Releases(p.Namespace, p.Package)
	if err != nil {
		return nil, err
	}

	var releases []Release

	for _, rm := range versions {
		r := NewRelease(p.store, p.Namespace, p.Package, rm.Version, rm.Digest, rm.CreatedAt, rm.Size,
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
	f, err := p.store.Pull(p.Namespace, p.Package, digest)
	if err != nil {
		return nil, nil, err
	}

	hdr := &multipart.FileHeader{
		Filename: fmt.Sprintf("%s.tar.gz", digest),
	}

	return f, hdr, nil
}

// ListPackages lists all packages.
func ListPackages(s store.Store) ([]Package, error) {
	namespaces, err := s.Namespaces()
	if err != nil {
		return nil, err
	}

	var packages []Package

	for _, ns := range namespaces {
		pms, err := s.Packages(ns)
		if err != nil {
			return nil, nil
		}

		for _, pm := range pms {
			p := NewPackage(s, pm)
			packages = append(packages, *p)
		}
	}

	return packages, nil
}
