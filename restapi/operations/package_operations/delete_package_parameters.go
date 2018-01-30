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

// NewDeletePackageParams creates a new DeletePackageParams object
// with the default values initialized.
func NewDeletePackageParams() DeletePackageParams {
	var ()
	return DeletePackageParams{}
}

// DeletePackageParams contains all the bound params for the delete package operation
// typically these are obtained from a http.Request
//
// swagger:parameters deletePackage
type DeletePackageParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*content type
	  Required: true
	  In: path
	*/
	MediaType string
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
	/*release name
	  Required: true
	  In: path
	*/
	Release string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *DeletePackageParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rMediaType, rhkMediaType, _ := route.Params.GetOK("media_type")
	if err := o.bindMediaType(rMediaType, rhkMediaType, route.Formats); err != nil {
		res = append(res, err)
	}

	rNamespace, rhkNamespace, _ := route.Params.GetOK("namespace")
	if err := o.bindNamespace(rNamespace, rhkNamespace, route.Formats); err != nil {
		res = append(res, err)
	}

	rPackage, rhkPackage, _ := route.Params.GetOK("package")
	if err := o.bindPackage(rPackage, rhkPackage, route.Formats); err != nil {
		res = append(res, err)
	}

	rRelease, rhkRelease, _ := route.Params.GetOK("release")
	if err := o.bindRelease(rRelease, rhkRelease, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeletePackageParams) bindMediaType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.MediaType = raw

	return nil
}

func (o *DeletePackageParams) bindNamespace(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Namespace = raw

	return nil
}

func (o *DeletePackageParams) bindPackage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Package = raw

	return nil
}

func (o *DeletePackageParams) bindRelease(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Release = raw

	return nil
}
