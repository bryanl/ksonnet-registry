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

func ShowPackage(s store.Store, params package_operations.ShowPackageParams) middleware.Responder {
	release, err := registry.ShowRelease(
		s,
		params.Namespace,
		params.Package,
		params.Release,
	)

	if err != nil {
		m := &models.Error{Message: err.Error()}
		resp := package_operations.NewShowPackageManifestsNotFound().
			WithPayload(m)
		return resp
	}

	manifest := &models.Manifest{
		Package:   release.Package,
		Release:   release.Version,
		CreatedAt: strfmt.DateTime(release.CreatedAt),
		Content: &models.OciDescriptor{
			Digest: release.Digest(),
		},
	}
	payload := models.PackageManifest{manifest}

	resp := package_operations.NewShowPackageManifestsOK().
		WithPayload(payload)

	return resp
}

func DeletePackage(s store.Store, params package_operations.DeletePackageParams) middleware.Responder {
	if err := registry.DeleteRelease(s, params.Namespace, params.Package, params.Release); err != nil {
		payload := &models.Error{Message: err.Error()}
		return package_operations.NewDeletePackageNotFound().
			WithPayload(payload)
	}

	return package_operations.NewDeletePackageOK()
}

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
