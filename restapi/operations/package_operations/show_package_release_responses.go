// Code generated by go-swagger; DO NOT EDIT.

package package_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bryanl/ksonnet-registry/models"
)

// ShowPackageReleaseOKCode is the HTTP code returned for type ShowPackageReleaseOK
const ShowPackageReleaseOKCode int = 200

/*ShowPackageReleaseOK successful operation

swagger:response showPackageReleaseOK
*/
type ShowPackageReleaseOK struct {

	/*
	  In: Body
	*/
	Payload *models.Manifest `json:"body,omitempty"`
}

// NewShowPackageReleaseOK creates ShowPackageReleaseOK with default headers values
func NewShowPackageReleaseOK() *ShowPackageReleaseOK {
	return &ShowPackageReleaseOK{}
}

// WithPayload adds the payload to the show package release o k response
func (o *ShowPackageReleaseOK) WithPayload(payload *models.Manifest) *ShowPackageReleaseOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the show package release o k response
func (o *ShowPackageReleaseOK) SetPayload(payload *models.Manifest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShowPackageReleaseOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ShowPackageReleaseUnauthorizedCode is the HTTP code returned for type ShowPackageReleaseUnauthorized
const ShowPackageReleaseUnauthorizedCode int = 401

/*ShowPackageReleaseUnauthorized Not authorized to read the package

swagger:response showPackageReleaseUnauthorized
*/
type ShowPackageReleaseUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewShowPackageReleaseUnauthorized creates ShowPackageReleaseUnauthorized with default headers values
func NewShowPackageReleaseUnauthorized() *ShowPackageReleaseUnauthorized {
	return &ShowPackageReleaseUnauthorized{}
}

// WithPayload adds the payload to the show package release unauthorized response
func (o *ShowPackageReleaseUnauthorized) WithPayload(payload *models.Error) *ShowPackageReleaseUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the show package release unauthorized response
func (o *ShowPackageReleaseUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShowPackageReleaseUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ShowPackageReleaseNotFoundCode is the HTTP code returned for type ShowPackageReleaseNotFound
const ShowPackageReleaseNotFoundCode int = 404

/*ShowPackageReleaseNotFound Release not found

swagger:response showPackageReleaseNotFound
*/
type ShowPackageReleaseNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewShowPackageReleaseNotFound creates ShowPackageReleaseNotFound with default headers values
func NewShowPackageReleaseNotFound() *ShowPackageReleaseNotFound {
	return &ShowPackageReleaseNotFound{}
}

// WithPayload adds the payload to the show package release not found response
func (o *ShowPackageReleaseNotFound) WithPayload(payload *models.Error) *ShowPackageReleaseNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the show package release not found response
func (o *ShowPackageReleaseNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShowPackageReleaseNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}