// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserRequest user request
//
// swagger:model UserRequest
type UserRequest struct {

	// admin
	Admin bool `json:"admin,omitempty"`

	// alert
	Alert bool `json:"alert,omitempty"`

	// email
	// Example: test@ansiblesemaphore.test
	Email string `json:"email,omitempty"`

	// external
	External bool `json:"external,omitempty"`

	// name
	// Example: Integration Test User
	Name string `json:"name,omitempty"`

	// password
	// Format: password
	Password strfmt.Password `json:"password,omitempty"`

	// username
	// Example: test-user
	Username string `json:"username,omitempty"`
}

// Validate validates this user request
func (m *UserRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserRequest) validatePassword(formats strfmt.Registry) error {
	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if err := validate.FormatOf("password", "body", "password", m.Password.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this user request based on context it is used
func (m *UserRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserRequest) UnmarshalBinary(b []byte) error {
	var res UserRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
