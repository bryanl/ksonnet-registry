// Code generated by go-swagger; DO NOT EDIT.

package blobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new blobs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for blobs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PullBlob pulls a package blob by digest
*/
func (a *Client) PullBlob(params *PullBlobParams) (*PullBlobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPullBlobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pullBlob",
		Method:             "GET",
		PathPattern:        "/api/v1/packages/{namespace}/{package}/blobs/sha256/{digest}",
		ProducesMediaTypes: []string{"application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PullBlobReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PullBlobOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
