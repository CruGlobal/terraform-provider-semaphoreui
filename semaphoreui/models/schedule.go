// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Schedule schedule
//
// swagger:model Schedule
type Schedule struct {

	// active
	Active bool `json:"active,omitempty"`

	// cron format
	CronFormat string `json:"cron_format,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// project id
	ProjectID int64 `json:"project_id,omitempty"`

	// template id
	TemplateID int64 `json:"template_id,omitempty"`
}

// Validate validates this schedule
func (m *Schedule) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this schedule based on context it is used
func (m *Schedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Schedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Schedule) UnmarshalBinary(b []byte) error {
	var res Schedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
