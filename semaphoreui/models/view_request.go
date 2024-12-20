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

// ViewRequest view request
//
// swagger:model ViewRequest
type ViewRequest struct {

	// position
	// Minimum: 1
	Position int64 `json:"position,omitempty"`

	// project id
	// Minimum: 1
	ProjectID int64 `json:"project_id,omitempty"`

	// title
	// Example: Test
	Title string `json:"title,omitempty"`
}

// Validate validates this view request
func (m *ViewRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ViewRequest) validatePosition(formats strfmt.Registry) error {
	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if err := validate.MinimumInt("position", "body", m.Position, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *ViewRequest) validateProjectID(formats strfmt.Registry) error {
	if swag.IsZero(m.ProjectID) { // not required
		return nil
	}

	if err := validate.MinimumInt("project_id", "body", m.ProjectID, 1, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this view request based on context it is used
func (m *ViewRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ViewRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ViewRequest) UnmarshalBinary(b []byte) error {
	var res ViewRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
