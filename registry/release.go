package registry

import (
	"encoding/base64"
	"time"

	"github.com/bryanl/ksonnet-registry/store"
)

// Release is a released version of a package.
type Release struct {
	Namespace string
	Package   string
	Version   string
	Digest    string
	CreatedAt time.Time
	store     store.Store
}

// NewRelease creates a new instance of Release.
func NewRelease(s store.Store, ns, pkg, version, digest string) *Release {
	r := &Release{
		Namespace: ns,
		Package:   pkg,
		Version:   version,
		Digest:    digest,
		CreatedAt: time.Now(),
		store:     s,
	}

	return r
}

// Delete removes a release.
func (r *Release) Delete() error {
	return r.store.RemoveRelease(r.Namespace, r.Package, r.Version)
}

// CreateRelease creates a release.
func CreateRelease(s store.Store, nsName, pkgName, ver, blob string) (*Release, error) {
	ns, err := GetNamespace(s, nsName)
	if err != nil {
		return nil, err
	}

	pkg, err := ns.Package(pkgName)
	if err != nil {
		return nil, err
	}

	data, err := base64.StdEncoding.DecodeString(blob)
	if err != nil {
		return nil, err
	}

	release, err := pkg.CreateRelease(ver, data)
	if err != nil {
		return nil, err
	}

	return release, nil
}

// ShowRelease shows a release.
func ShowRelease(s store.Store, nsName, pkgName, ver string) (*Release, error) {
	ns, err := GetNamespace(s, nsName)
	if err != nil {
		return nil, err
	}

	pkg, err := ns.Package(pkgName)
	if err != nil {
		return nil, err
	}

	return pkg.Release(ver)
}

// DeleteRelease removes a release.
func DeleteRelease(s store.Store, nsName, pkgName, ver string) error {
	ns, err := GetNamespace(s, nsName)
	if err != nil {
		return err
	}

	pkg, err := ns.Package(pkgName)
	if err != nil {
		return err
	}

	return pkg.Delete(ver)
}
