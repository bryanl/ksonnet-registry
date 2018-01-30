// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bryanl/ksonnet-registry/models"
)

// PullPackageOKCode is the HTTP code returned for type PullPackageOK
const PullPackageOKCode int = 200

/*PullPackageOK successful operation

swagger:response pullPackageOK
*/
type PullPackageOK struct {

	/*
	  In: Body
	*/
	Payload models.PullPackageOKBody `json:"body,omitempty"`
}

// NewPullPackageOK creates PullPackageOK with default headers values
func NewPullPackageOK() *PullPackageOK {
	return &PullPackageOK{}
}

// WithPayload adds the payload to the pull package o k response
func (o *PullPackageOK) WithPayload(payload models.PullPackageOKBody) *PullPackageOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the pull package o k response
func (o *PullPackageOK) SetPayload(payload models.PullPackageOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PullPackageOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// PullPackageUnauthorizedCode is the HTTP code returned for type PullPackageUnauthorized
const PullPackageUnauthorizedCode int = 401

/*PullPackageUnauthorized Not authorized to read the package

swagger:response pullPackageUnauthorized
*/
type PullPackageUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPullPackageUnauthorized creates PullPackageUnauthorized with default headers values
func NewPullPackageUnauthorized() *PullPackageUnauthorized {
	return &PullPackageUnauthorized{}
}

// WithPayload adds the payload to the pull package unauthorized response
func (o *PullPackageUnauthorized) WithPayload(payload *models.Error) *PullPackageUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the pull package unauthorized response
func (o *PullPackageUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PullPackageUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PullPackageNotFoundCode is the HTTP code returned for type PullPackageNotFound
const PullPackageNotFoundCode int = 404

/*PullPackageNotFound Package not found

swagger:response pullPackageNotFound
*/
type PullPackageNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPullPackageNotFound creates PullPackageNotFound with default headers values
func NewPullPackageNotFound() *PullPackageNotFound {
	return &PullPackageNotFound{}
}

// WithPayload adds the payload to the pull package not found response
func (o *PullPackageNotFound) WithPayload(payload *models.Error) *PullPackageNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the pull package not found response
func (o *PullPackageNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PullPackageNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}