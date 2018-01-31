// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/bryanl/ksonnet-registry/models"
)

// PullPackageJSONReader is a Reader for the PullPackageJSON structure.
type PullPackageJSONReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PullPackageJSONReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPullPackageJSONOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPullPackageJSONUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPullPackageJSONNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPullPackageJSONOK creates a PullPackageJSONOK with default headers values
func NewPullPackageJSONOK() *PullPackageJSONOK {
	return &PullPackageJSONOK{}
}

/*PullPackageJSONOK handles this case with default header values.

successful operation
*/
type PullPackageJSONOK struct {
	Payload *models.PullJSON
}

func (o *PullPackageJSONOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/packages/{namespace}/{package}/{release}/{media_type}/pull/json][%d] pullPackageJsonOK  %+v", 200, o.Payload)
}

func (o *PullPackageJSONOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PullJSON)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPullPackageJSONUnauthorized creates a PullPackageJSONUnauthorized with default headers values
func NewPullPackageJSONUnauthorized() *PullPackageJSONUnauthorized {
	return &PullPackageJSONUnauthorized{}
}

/*PullPackageJSONUnauthorized handles this case with default header values.

Not authorized to read the package
*/
type PullPackageJSONUnauthorized struct {
	Payload *models.Error
}

func (o *PullPackageJSONUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/packages/{namespace}/{package}/{release}/{media_type}/pull/json][%d] pullPackageJsonUnauthorized  %+v", 401, o.Payload)
}

func (o *PullPackageJSONUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPullPackageJSONNotFound creates a PullPackageJSONNotFound with default headers values
func NewPullPackageJSONNotFound() *PullPackageJSONNotFound {
	return &PullPackageJSONNotFound{}
}

/*PullPackageJSONNotFound handles this case with default header values.

Package not found
*/
type PullPackageJSONNotFound struct {
	Payload *models.Error
}

func (o *PullPackageJSONNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/packages/{namespace}/{package}/{release}/{media_type}/pull/json][%d] pullPackageJsonNotFound  %+v", 404, o.Payload)
}

func (o *PullPackageJSONNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
