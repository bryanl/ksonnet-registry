// Code generated by go-swagger; DO NOT EDIT.

package channel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bryanl/ksonnet-registry/models"
)

// CreateChannelReleaseOKCode is the HTTP code returned for type CreateChannelReleaseOK
const CreateChannelReleaseOKCode int = 200

/*CreateChannelReleaseOK successful operation

swagger:response createChannelReleaseOK
*/
type CreateChannelReleaseOK struct {

	/*
	  In: Body
	*/
	Payload *models.Channel `json:"body,omitempty"`
}

// NewCreateChannelReleaseOK creates CreateChannelReleaseOK with default headers values
func NewCreateChannelReleaseOK() *CreateChannelReleaseOK {
	return &CreateChannelReleaseOK{}
}

// WithPayload adds the payload to the create channel release o k response
func (o *CreateChannelReleaseOK) WithPayload(payload *models.Channel) *CreateChannelReleaseOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create channel release o k response
func (o *CreateChannelReleaseOK) SetPayload(payload *models.Channel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateChannelReleaseOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateChannelReleaseUnauthorizedCode is the HTTP code returned for type CreateChannelReleaseUnauthorized
const CreateChannelReleaseUnauthorizedCode int = 401

/*CreateChannelReleaseUnauthorized Not authorized to read the package

swagger:response createChannelReleaseUnauthorized
*/
type CreateChannelReleaseUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateChannelReleaseUnauthorized creates CreateChannelReleaseUnauthorized with default headers values
func NewCreateChannelReleaseUnauthorized() *CreateChannelReleaseUnauthorized {
	return &CreateChannelReleaseUnauthorized{}
}

// WithPayload adds the payload to the create channel release unauthorized response
func (o *CreateChannelReleaseUnauthorized) WithPayload(payload *models.Error) *CreateChannelReleaseUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create channel release unauthorized response
func (o *CreateChannelReleaseUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateChannelReleaseUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateChannelReleaseNotFoundCode is the HTTP code returned for type CreateChannelReleaseNotFound
const CreateChannelReleaseNotFoundCode int = 404

/*CreateChannelReleaseNotFound Package not found

swagger:response createChannelReleaseNotFound
*/
type CreateChannelReleaseNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateChannelReleaseNotFound creates CreateChannelReleaseNotFound with default headers values
func NewCreateChannelReleaseNotFound() *CreateChannelReleaseNotFound {
	return &CreateChannelReleaseNotFound{}
}

// WithPayload adds the payload to the create channel release not found response
func (o *CreateChannelReleaseNotFound) WithPayload(payload *models.Error) *CreateChannelReleaseNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create channel release not found response
func (o *CreateChannelReleaseNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateChannelReleaseNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}