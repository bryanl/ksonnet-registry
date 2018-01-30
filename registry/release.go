package registry

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/bryanl/ksonnet-registry/store"
	"github.com/pkg/errors"
)

type Release struct {
	Namespace string
	Package   string
	Version   string
	CreatedAt time.Time
	digest    string
	store     store.Store
}

func NewRelease(s store.Store, ns, pkg, version string, data []byte) (*Release, error) {
	sum := sha256.Sum256(data)
	digest := fmt.Sprintf("%x", sum)

	if err := s.Write(digest, data); err != nil {
		return nil, errors.Wrap(err, "store data")
	}

	return &Release{
		Namespace: ns,
		Package:   pkg,
		Version:   version,
		CreatedAt: time.Now(),
		digest:    digest,
		store:     s,
	}, nil
}

func (r *Release) Digest() string {
	return r.digest
}

func (r *Release) Delete() error {
	return r.store.Delete(r.digest)
}

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
