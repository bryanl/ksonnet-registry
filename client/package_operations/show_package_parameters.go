// Code generated by go-swagger; DO NOT EDIT.

package package_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewShowPackageParams creates a new ShowPackageParams object
// with the default values initialized.
func NewShowPackageParams() *ShowPackageParams {
	var ()
	return &ShowPackageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewShowPackageParamsWithTimeout creates a new ShowPackageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewShowPackageParamsWithTimeout(timeout time.Duration) *ShowPackageParams {
	var ()
	return &ShowPackageParams{

		timeout: timeout,
	}
}

// NewShowPackageParamsWithContext creates a new ShowPackageParams object
// with the default values initialized, and the ability to set a context for a request
func NewShowPackageParamsWithContext(ctx context.Context) *ShowPackageParams {
	var ()
	return &ShowPackageParams{

		Context: ctx,
	}
}

// NewShowPackageParamsWithHTTPClient creates a new ShowPackageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewShowPackageParamsWithHTTPClient(client *http.Client) *ShowPackageParams {
	var ()
	return &ShowPackageParams{
		HTTPClient: client,
	}
}

/*ShowPackageParams contains all the parameters to send to the API endpoint
for the show package operation typically these are written to a http.Request
*/
type ShowPackageParams struct {

	/*MediaType
	  content type

	*/
	MediaType string
	/*Namespace
	  namespace

	*/
	Namespace string
	/*Package
	  package name

	*/
	Package string
	/*Release
	  release name

	*/
	Release string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the show package params
func (o *ShowPackageParams) WithTimeout(timeout time.Duration) *ShowPackageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the show package params
func (o *ShowPackageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the show package params
func (o *ShowPackageParams) WithContext(ctx context.Context) *ShowPackageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the show package params
func (o *ShowPackageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the show package params
func (o *ShowPackageParams) WithHTTPClient(client *http.Client) *ShowPackageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the show package params
func (o *ShowPackageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMediaType adds the mediaType to the show package params
func (o *ShowPackageParams) WithMediaType(mediaType string) *ShowPackageParams {
	o.SetMediaType(mediaType)
	return o
}

// SetMediaType adds the mediaType to the show package params
func (o *ShowPackageParams) SetMediaType(mediaType string) {
	o.MediaType = mediaType
}

// WithNamespace adds the namespace to the show package params
func (o *ShowPackageParams) WithNamespace(namespace string) *ShowPackageParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the show package params
func (o *ShowPackageParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPackage adds the packageVar to the show package params
func (o *ShowPackageParams) WithPackage(packageVar string) *ShowPackageParams {
	o.SetPackage(packageVar)
	return o
}

// SetPackage adds the package to the show package params
func (o *ShowPackageParams) SetPackage(packageVar string) {
	o.Package = packageVar
}

// WithRelease adds the release to the show package params
func (o *ShowPackageParams) WithRelease(release string) *ShowPackageParams {
	o.SetRelease(release)
	return o
}

// SetRelease adds the release to the show package params
func (o *ShowPackageParams) SetRelease(release string) {
	o.Release = release
}

// WriteToRequest writes these params to a swagger request
func (o *ShowPackageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param media_type
	if err := r.SetPathParam("media_type", o.MediaType); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param package
	if err := r.SetPathParam("package", o.Package); err != nil {
		return err
	}

	// path param release
	if err := r.SetPathParam("release", o.Release); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
