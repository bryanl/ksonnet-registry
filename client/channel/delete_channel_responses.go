// Code generated by go-swagger; DO NOT EDIT.

package channel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/bryanl/ksonnet-registry/models"
)

// DeleteChannelReader is a Reader for the DeleteChannel structure.
type DeleteChannelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteChannelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteChannelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteChannelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteChannelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteChannelOK creates a DeleteChannelOK with default headers values
func NewDeleteChannelOK() *DeleteChannelOK {
	return &DeleteChannelOK{}
}

/*DeleteChannelOK handles this case with default header values.

successful operation
*/
type DeleteChannelOK struct {
	Payload models.DeleteChannelOKBody
}

func (o *DeleteChannelOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/packages/{namespace}/{package}/channels/{channel}][%d] deleteChannelOK  %+v", 200, o.Payload)
}

func (o *DeleteChannelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteChannelUnauthorized creates a DeleteChannelUnauthorized with default headers values
func NewDeleteChannelUnauthorized() *DeleteChannelUnauthorized {
	return &DeleteChannelUnauthorized{}
}

/*DeleteChannelUnauthorized handles this case with default header values.

Not authorized to read the package
*/
type DeleteChannelUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteChannelUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/packages/{namespace}/{package}/channels/{channel}][%d] deleteChannelUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteChannelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteChannelNotFound creates a DeleteChannelNotFound with default headers values
func NewDeleteChannelNotFound() *DeleteChannelNotFound {
	return &DeleteChannelNotFound{}
}

/*DeleteChannelNotFound handles this case with default header values.

Package not found
*/
type DeleteChannelNotFound struct {
	Payload *models.Error
}

func (o *DeleteChannelNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/packages/{namespace}/{package}/channels/{channel}][%d] deleteChannelNotFound  %+v", 404, o.Payload)
}

func (o *DeleteChannelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}