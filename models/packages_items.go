// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PackagesItems test
// swagger:model packagesItems
type PackagesItems struct {

	// created_at
	//
	// Package creation date
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// default-release
	//
	// Default/latest release version
	Default string `json:"default,omitempty"`

	// package-name
	//
	// Package name
	Name string `json:"name,omitempty"`

	// available-releases
	//
	// All available releases
	Releases []string `json:"releases"`

	// visibility
	//
	// package visibility (public or private)
	Visibility string `json:"visibility,omitempty"`
}

// Validate validates this packages items
func (m *PackagesItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReleases(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PackagesItems) validateReleases(formats strfmt.Registry) error {

	if swag.IsZero(m.Releases) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PackagesItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PackagesItems) UnmarshalBinary(b []byte) error {
	var res PackagesItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
