// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TemplateSurveyVar template survey var
//
// swagger:model TemplateSurveyVar
type TemplateSurveyVar struct {

	// description
	Description string `json:"description,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// required
	Required bool `json:"required,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// type
	// Example: String =\u003e \"\", Integer =\u003e \"int\
	Type string `json:"type,omitempty"`
}

// Validate validates this template survey var
func (m *TemplateSurveyVar) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this template survey var based on context it is used
func (m *TemplateSurveyVar) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TemplateSurveyVar) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TemplateSurveyVar) UnmarshalBinary(b []byte) error {
	var res TemplateSurveyVar
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
