// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InfoType info type
//
// swagger:model InfoType
type InfoType struct {

	// update
	Update *InfoTypeUpdate `json:"update,omitempty"`

	// update body
	UpdateBody string `json:"updateBody,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this info type
func (m *InfoType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUpdate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InfoType) validateUpdate(formats strfmt.Registry) error {
	if swag.IsZero(m.Update) { // not required
		return nil
	}

	if m.Update != nil {
		if err := m.Update.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("update")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("update")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this info type based on the context it is used
func (m *InfoType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUpdate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InfoType) contextValidateUpdate(ctx context.Context, formats strfmt.Registry) error {

	if m.Update != nil {

		if swag.IsZero(m.Update) { // not required
			return nil
		}

		if err := m.Update.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("update")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("update")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InfoType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InfoType) UnmarshalBinary(b []byte) error {
	var res InfoType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// InfoTypeUpdate info type update
//
// swagger:model InfoTypeUpdate
type InfoTypeUpdate struct {

	// tag name
	TagName string `json:"tag_name,omitempty"`
}

// Validate validates this info type update
func (m *InfoTypeUpdate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this info type update based on context it is used
func (m *InfoTypeUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InfoTypeUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InfoTypeUpdate) UnmarshalBinary(b []byte) error {
	var res InfoTypeUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}