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
		Package:   release.Package,
		Release:   release.Version,
		CreatedAt: strfmt.DateTime(release.CreatedAt),
	}

	resp := package_operations.NewCreatePackageCreated().
		WithPayload(payload)
	return resp
}

// ShowPackage shows a package.
func ShowPackage(s store.Store, params package_operations.ShowPackageReleasesParams) middleware.Responder {
	releases, err := registry.ShowReleases(
		s,
		params.Namespace,
		params.Package,
	)

	if err != nil {
		m := &models.Error{Message: err.Error()}
		resp := package_operations.NewShowPackageManifestsNotFound().
			WithPayload(m)
		return resp
	}

	var manifests models.PackageManifest

	for _, r := range releases {
		manifest := &models.Manifest{
			Package:   r.Package,
			Release:   r.Version,
			CreatedAt: strfmt.DateTime(r.CreatedAt),
			Content: &models.OciDescriptor{
				Digest: r.Digest,
				Size:   r.Size,
			},
		}

		manifests = append(manifests, manifest)

	}

	resp := package_operations.NewShowPackageManifestsOK().
		WithPayload(manifests)

	return resp
}

// DeletePackage deletes a package.
func DeletePackage(s store.Store, params package_operations.DeletePackageParams) middleware.Responder {
	if err := registry.DeleteRelease(s, params.Namespace, params.Package, params.Release); err != nil {
		payload := &models.Error{Message: err.Error()}
		return package_operations.NewDeletePackageNotFound().
			WithPayload(payload)
	}

	return package_operations.NewDeletePackageOK()
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
