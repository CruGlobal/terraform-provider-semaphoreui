// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
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
	// Example: int
	// Enum: ["","int","enum","secret"]
	Type string `json:"type,omitempty"`

	// values
	Values []*TemplateSurveyVarValue `json:"values"`
}

// Validate validates this template survey var
func (m *TemplateSurveyVar) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var templateSurveyVarTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["","int","enum","secret"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		templateSurveyVarTypeTypePropEnum = append(templateSurveyVarTypeTypePropEnum, v)
	}
}

const (

	// TemplateSurveyVarTypeEmpty captures enum value ""
	TemplateSurveyVarTypeEmpty string = ""

	// TemplateSurveyVarTypeInt captures enum value "int"
	TemplateSurveyVarTypeInt string = "int"

	// TemplateSurveyVarTypeEnum captures enum value "enum"
	TemplateSurveyVarTypeEnum string = "enum"

	// TemplateSurveyVarTypeSecret captures enum value "secret"
	TemplateSurveyVarTypeSecret string = "secret"
)

// prop value enum
func (m *TemplateSurveyVar) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, templateSurveyVarTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TemplateSurveyVar) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *TemplateSurveyVar) validateValues(formats strfmt.Registry) error {
	if swag.IsZero(m.Values) { // not required
		return nil
	}

	for i := 0; i < len(m.Values); i++ {
		if swag.IsZero(m.Values[i]) { // not required
			continue
		}

		if m.Values[i] != nil {
			if err := m.Values[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("values" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this template survey var based on the context it is used
func (m *TemplateSurveyVar) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateValues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TemplateSurveyVar) contextValidateValues(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Values); i++ {

		if m.Values[i] != nil {

			if swag.IsZero(m.Values[i]) { // not required
				return nil
			}

			if err := m.Values[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("values" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

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
