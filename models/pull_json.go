// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PullJSON PackageContent
//
// Package content
// swagger:model PullJson
type PullJSON struct {

	// blob
	//
	// Package blob: a tar.gz in b64-encoded
	Blob string `json:"blob,omitempty"`

	// filename
	//
	// suggested filename
	Filename string `json:"filename,omitempty"`

	// package-name
	//
	// Package name
	Package string `json:"package,omitempty"`

	// package-version
	//
	// Package version
	Release string `json:"release,omitempty"`
}

// Validate validates this pull Json
func (m *PullJSON) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PullJSON) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PullJSON) UnmarshalBinary(b []byte) error {
	var res PullJSON
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
