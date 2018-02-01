package apiadapter

import (
	"fmt"

	"github.com/bryanl/ksonnet-registry/models"
	"github.com/bryanl/ksonnet-registry/registry"
	"github.com/bryanl/ksonnet-registry/restapi/operations/blobs"
	"github.com/bryanl/ksonnet-registry/restapi/operations/package_operations"
	"github.com/bryanl/ksonnet-registry/store"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// CreatePackage creates a package.
func CreatePackage(s store.Store, params package_operations.CreatePackageParams) middleware.Responder {
	release, err := registry.CreateRelease(
		s,
		params.Namespace,
		params.Package,
		params.Body.Release,
		params.Body.Blob,
	)

	if err != nil {
		m := &models.Error{Message: err.Error()}
		resp := package_operations.NewCreatePackageUnprocessableEntity().
			WithPayload(m)
		return resp
	}

	payload := &models.Package{
		Package:   release.Package(),
		Release:   release.Version,
		CreatedAt: strfmt.DateTime(release.CreatedAt),
	}

	resp := package_operations.NewCreatePackageCreated().
		WithPayload(payload)
	return resp
}

// ShowPackageRelease shows a release in a package.
func ShowPackageRelease(s store.Store, params package_operations.ShowPackageReleaseParams) middleware.Responder {
	release, err := registry.ShowRelease(
		s,
		params.Namespace,
		params.Package,
		params.Release,
	)

	if err != nil {
		m := &models.Error{Message: err.Error()}
		resp := package_operations.NewShowPackageReleaseNotFound().
			WithPayload(m)
		return resp
	}

	manifest := makeManifest(*release)

	resp := package_operations.NewShowPackageReleaseOK().
		WithPayload(manifest)

	return resp
}

func makeManifest(r registry.Release) *models.Manifest {
	var deps models.PartDescriptorDependencies
	for dep, con := range r.Deps {
		md := &models.Dependency{
			Name:       dep,
			Constraint: con,
		}
		deps = append(deps, md)
	}

	manifest := &models.Manifest{
		Package:   r.Package(),
		Release:   r.Version,
		CreatedAt: strfmt.DateTime(r.CreatedAt),
		Content: &models.PartDescriptor{
			Digest:       r.Digest,
			Size:         r.Size,
			Dependencies: deps,
		},
	}

	return manifest
}

// ShowPackageReleases shows releases in a package.
func ShowPackageReleases(s store.Store, params package_operations.ShowPackageReleasesParams) middleware.Responder {
	releases, err := registry.ShowReleases(
		s,
		params.Namespace,
		params.Package,
	)

	if err != nil {
		m := &models.Error{Message: err.Error()}
		resp := package_operations.NewShowPackageReleasesNotFound().
			WithPayload(m)
		return resp
	}

	var manifests models.PackageManifest

	for _, r := range releases {
		manifest := makeManifest(r)
		manifests = append(manifests, manifest)
	}

	resp := package_operations.NewShowPackageReleasesOK().
		WithPayload(manifests)

	return resp
}

// DeletePackageRelease deletes a release from a package.
func DeletePackageRelease(s store.Store, params package_operations.DeletePackageReleaseParams) middleware.Responder {
	if err := registry.DeleteRelease(s, params.Namespace, params.Package, params.Release); err != nil {
		payload := &models.Error{Message: err.Error()}
		return package_operations.NewDeletePackageNotFound().
			WithPayload(payload)
	}

	return package_operations.NewDeletePackageReleaseNoContent()
}

// PullPackage pulls a package.
func PullPackage(s store.Store, params blobs.PullBlobParams) middleware.Responder {
	ns, err := registry.GetNamespace(s, params.Namespace)
	if err != nil {
		return blobs.NewPullBlobNotFound().
			WithPayload(&models.Error{Message: fmt.Sprintf("namespace %q not found", params.Namespace)})
	}

	pkg, err := ns.Package(params.Package)
	if err != nil {
		return blobs.NewPullBlobNotFound().
			WithPayload(&models.Error{Message: fmt.Sprintf("package %q not found", params.Package)})
	}

	f, hdr, err := pkg.Pull(params.Digest)
	if err != nil {
		return blobs.NewPullBlobNotFound().
			WithPayload(&models.Error{Message: fmt.Sprintf("package %q not found", params.Package)})
	}

	payload := models.PullBlobOKBody{
		Data:   f,
		Header: hdr,
	}

	return blobs.NewPullBlobOK().
		WithPayload(payload)
}
