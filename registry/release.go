package registry

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/bryanl/ksonnet-registry/store"
)

// Release is a released version of a package.
type Release struct {
	Version   string
	Digest    string
	CreatedAt time.Time
	Size      int64
	pkgName   string
	nsName    string
	store     store.Store
}

// NewRelease creates a new instance of Release.
func NewRelease(s store.Store, ns, pkg, version, digest string, createdAt time.Time, size int64) *Release {
	r := &Release{
		nsName:    ns,
		pkgName:   pkg,
		Version:   version,
		Digest:    digest,
		CreatedAt: createdAt,
		Size:      size,
		store:     s,
	}

	return r
}

// Package is the package name for this release. It is the namespace combined with
// the package.
func (r *Release) Package() string {
	return fmt.Sprintf("%s/%s", r.nsName, r.pkgName)
}

// Delete removes a release.
func (r *Release) Delete() error {
	return r.store.RemoveRelease(r.nsName, r.pkgName, r.Version)
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

// ShowReleases shows all releases for a package.
func ShowReleases(s store.Store, nsName, pkgName string) ([]Release, error) {
	ns, err := GetNamespace(s, nsName)
	if err != nil {
		return nil, err
	}

	pkg, err := ns.Package(pkgName)
	if err != nil {
		return nil, err
	}

	return pkg.Releases()
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
