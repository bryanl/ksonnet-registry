// Code generated by go-swagger; DO NOT EDIT.

package package_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// CreatePackageURL generates an URL for the create package operation
type CreatePackageURL struct {
	Namespace string
	Package   string

	Force *bool

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *CreatePackageURL) WithBasePath(bp string) *CreatePackageURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *CreatePackageURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *CreatePackageURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/api/v1/packages/{namespace}/{package}"

	namespace := o.Namespace
	if namespace != "" {
		_path = strings.Replace(_path, "{namespace}", namespace, -1)
	} else {
		return nil, errors.New("Namespace is required on CreatePackageURL")
	}
	packageVar := o.Package
	if packageVar != "" {
		_path = strings.Replace(_path, "{package}", packageVar, -1)
	} else {
		return nil, errors.New("Package is required on CreatePackageURL")
	}
	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var force string
	if o.Force != nil {
		force = swag.FormatBool(*o.Force)
	}
	if force != "" {
		qs.Set("force", force)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *CreatePackageURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *CreatePackageURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *CreatePackageURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on CreatePackageURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on CreatePackageURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *CreatePackageURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
