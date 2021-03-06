// Code generated by go-swagger; DO NOT EDIT.

package package_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bryanl/ksonnet-registry/models"
)

// DeletePackageNoContentCode is the HTTP code returned for type DeletePackageNoContent
const DeletePackageNoContentCode int = 204

/*DeletePackageNoContent successful operation

swagger:response deletePackageNoContent
*/
type DeletePackageNoContent struct {
}

// NewDeletePackageNoContent creates DeletePackageNoContent with default headers values
func NewDeletePackageNoContent() *DeletePackageNoContent {
	return &DeletePackageNoContent{}
}

// WriteResponse to the client
func (o *DeletePackageNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeletePackageUnauthorizedCode is the HTTP code returned for type DeletePackageUnauthorized
const DeletePackageUnauthorizedCode int = 401

/*DeletePackageUnauthorized Not authorized to read the package

swagger:response deletePackageUnauthorized
*/
type DeletePackageUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeletePackageUnauthorized creates DeletePackageUnauthorized with default headers values
func NewDeletePackageUnauthorized() *DeletePackageUnauthorized {
	return &DeletePackageUnauthorized{}
}

// WithPayload adds the payload to the delete package unauthorized response
func (o *DeletePackageUnauthorized) WithPayload(payload *models.Error) *DeletePackageUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete package unauthorized response
func (o *DeletePackageUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePackageUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeletePackageNotFoundCode is the HTTP code returned for type DeletePackageNotFound
const DeletePackageNotFoundCode int = 404

/*DeletePackageNotFound Package not found

swagger:response deletePackageNotFound
*/
type DeletePackageNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeletePackageNotFound creates DeletePackageNotFound with default headers values
func NewDeletePackageNotFound() *DeletePackageNotFound {
	return &DeletePackageNotFound{}
}

// WithPayload adds the payload to the delete package not found response
func (o *DeletePackageNotFound) WithPayload(payload *models.Error) *DeletePackageNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete package not found response
func (o *DeletePackageNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePackageNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
