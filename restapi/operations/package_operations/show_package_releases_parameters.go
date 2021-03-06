// Code generated by go-swagger; DO NOT EDIT.

package package_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewShowPackageReleasesParams creates a new ShowPackageReleasesParams object
// with the default values initialized.
func NewShowPackageReleasesParams() ShowPackageReleasesParams {
	var ()
	return ShowPackageReleasesParams{}
}

// ShowPackageReleasesParams contains all the bound params for the show package releases operation
// typically these are obtained from a http.Request
//
// swagger:parameters showPackageReleases
type ShowPackageReleasesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*namespace
	  Required: true
	  In: path
	*/
	Namespace string
	/*package name
	  Required: true
	  In: path
	*/
	Package string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *ShowPackageReleasesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rNamespace, rhkNamespace, _ := route.Params.GetOK("namespace")
	if err := o.bindNamespace(rNamespace, rhkNamespace, route.Formats); err != nil {
		res = append(res, err)
	}

	rPackage, rhkPackage, _ := route.Params.GetOK("package")
	if err := o.bindPackage(rPackage, rhkPackage, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ShowPackageReleasesParams) bindNamespace(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Namespace = raw

	return nil
}

func (o *ShowPackageReleasesParams) bindPackage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Package = raw

	return nil
}
