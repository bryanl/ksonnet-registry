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

// NewShowPackageReleaseParams creates a new ShowPackageReleaseParams object
// with the default values initialized.
func NewShowPackageReleaseParams() *ShowPackageReleaseParams {
	var ()
	return &ShowPackageReleaseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewShowPackageReleaseParamsWithTimeout creates a new ShowPackageReleaseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewShowPackageReleaseParamsWithTimeout(timeout time.Duration) *ShowPackageReleaseParams {
	var ()
	return &ShowPackageReleaseParams{

		timeout: timeout,
	}
}

// NewShowPackageReleaseParamsWithContext creates a new ShowPackageReleaseParams object
// with the default values initialized, and the ability to set a context for a request
func NewShowPackageReleaseParamsWithContext(ctx context.Context) *ShowPackageReleaseParams {
	var ()
	return &ShowPackageReleaseParams{

		Context: ctx,
	}
}

// NewShowPackageReleaseParamsWithHTTPClient creates a new ShowPackageReleaseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewShowPackageReleaseParamsWithHTTPClient(client *http.Client) *ShowPackageReleaseParams {
	var ()
	return &ShowPackageReleaseParams{
		HTTPClient: client,
	}
}

/*ShowPackageReleaseParams contains all the parameters to send to the API endpoint
for the show package release operation typically these are written to a http.Request
*/
type ShowPackageReleaseParams struct {

	/*Namespace
	  namespace

	*/
	Namespace string
	/*Package
	  package name

	*/
	Package string
	/*Release
	  release

	*/
	Release string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the show package release params
func (o *ShowPackageReleaseParams) WithTimeout(timeout time.Duration) *ShowPackageReleaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the show package release params
func (o *ShowPackageReleaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the show package release params
func (o *ShowPackageReleaseParams) WithContext(ctx context.Context) *ShowPackageReleaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the show package release params
func (o *ShowPackageReleaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the show package release params
func (o *ShowPackageReleaseParams) WithHTTPClient(client *http.Client) *ShowPackageReleaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the show package release params
func (o *ShowPackageReleaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the show package release params
func (o *ShowPackageReleaseParams) WithNamespace(namespace string) *ShowPackageReleaseParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the show package release params
func (o *ShowPackageReleaseParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPackage adds the packageVar to the show package release params
func (o *ShowPackageReleaseParams) WithPackage(packageVar string) *ShowPackageReleaseParams {
	o.SetPackage(packageVar)
	return o
}

// SetPackage adds the package to the show package release params
func (o *ShowPackageReleaseParams) SetPackage(packageVar string) {
	o.Package = packageVar
}

// WithRelease adds the release to the show package release params
func (o *ShowPackageReleaseParams) WithRelease(release string) *ShowPackageReleaseParams {
	o.SetRelease(release)
	return o
}

// SetRelease adds the release to the show package release params
func (o *ShowPackageReleaseParams) SetRelease(release string) {
	o.Release = release
}

// WriteToRequest writes these params to a swagger request
func (o *ShowPackageReleaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
