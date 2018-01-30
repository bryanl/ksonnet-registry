// Code generated by go-swagger; DO NOT EDIT.

package package_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bryanl/ksonnet-registry/models"
)

// ListPackagesOKCode is the HTTP code returned for type ListPackagesOK
const ListPackagesOKCode int = 200

/*ListPackagesOK successful operation

swagger:response listPackagesOK
*/
type ListPackagesOK struct {

	/*
	  In: Body
	*/
	Payload models.Packages `json:"body,omitempty"`
}

// NewListPackagesOK creates ListPackagesOK with default headers values
func NewListPackagesOK() *ListPackagesOK {
	return &ListPackagesOK{}
}

// WithPayload adds the payload to the list packages o k response
func (o *ListPackagesOK) WithPayload(payload models.Packages) *ListPackagesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list packages o k response
func (o *ListPackagesOK) SetPayload(payload models.Packages) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListPackagesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.Packages, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}