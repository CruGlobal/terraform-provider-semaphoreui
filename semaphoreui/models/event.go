// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Event event
//
// swagger:model Event
type Event struct {

	// description
	Description string `json:"description,omitempty"`

	// object id
	ObjectID int64 `json:"object_id,omitempty"`

	// object type
	ObjectType string `json:"object_type,omitempty"`

	// project id
	ProjectID int64 `json:"project_id,omitempty"`

	// user id
	UserID int64 `json:"user_id,omitempty"`
}

// Validate validates this event
func (m *Event) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this event based on context it is used
func (m *Event) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Event) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Event) UnmarshalBinary(b []byte) error {
	var res Event
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}