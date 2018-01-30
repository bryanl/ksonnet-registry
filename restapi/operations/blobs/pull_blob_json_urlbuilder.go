// Code generated by go-swagger; DO NOT EDIT.

package blobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// PullBlobJSONURL generates an URL for the pull blob Json operation
type PullBlobJSONURL struct {
	Digest    string
	Namespace string
	Package   string

	Format *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PullBlobJSONURL) WithBasePath(bp string) *PullBlobJSONURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PullBlobJSONURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *PullBlobJSONURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/api/v1/packages/{namespace}/{package}/blobs/sha256/{digest}/json"

	digest := o.Digest
	if digest != "" {
		_path = strings.Replace(_path, "{digest}", digest, -1)
	} else {
		return nil, errors.New("Digest is required on PullBlobJSONURL")
	}
	namespace := o.Namespace
	if namespace != "" {
		_path = strings.Replace(_path, "{namespace}", namespace, -1)
	} else {
		return nil, errors.New("Namespace is required on PullBlobJSONURL")
	}
	packageVar := o.Package
	if packageVar != "" {
		_path = strings.Replace(_path, "{package}", packageVar, -1)
	} else {
		return nil, errors.New("Package is required on PullBlobJSONURL")
	}
	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var format string
	if o.Format != nil {
		format = *o.Format
	}
	if format != "" {
		qs.Set("format", format)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *PullBlobJSONURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *PullBlobJSONURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *PullBlobJSONURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on PullBlobJSONURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on PullBlobJSONURL")
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
func (o *PullBlobJSONURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
