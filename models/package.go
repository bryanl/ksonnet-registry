// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Package Package
//
// Package object
// swagger:model Package
type Package struct {

	// content
	Content *Manifest `json:"content,omitempty"`

	// created_at
	//
	// Package creation date
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// package-name
	//
	// Package name
	Package string `json:"package,omitempty"`

	// package-release
	//
	// Package release
	Release string `json:"release,omitempty"`
}

// Validate validates this package
func (m *Package) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Package) validateContent(formats strfmt.Registry) error {

	if swag.IsZero(m.Content) { // not required
		return nil
	}

	if m.Content != nil {

		if err := m.Content.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Package) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Package) UnmarshalBinary(b []byte) error {
	var res Package
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
