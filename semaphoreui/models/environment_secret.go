// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EnvironmentSecret environment secret
//
// swagger:model EnvironmentSecret
type EnvironmentSecret struct {

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// type
	// Enum: ["env","var"]
	Type string `json:"type,omitempty"`
}

// Validate validates this environment secret
func (m *EnvironmentSecret) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var environmentSecretTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["env","var"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		environmentSecretTypeTypePropEnum = append(environmentSecretTypeTypePropEnum, v)
	}
}

const (

	// EnvironmentSecretTypeEnv captures enum value "env"
	EnvironmentSecretTypeEnv string = "env"

	// EnvironmentSecretTypeVar captures enum value "var"
	EnvironmentSecretTypeVar string = "var"
)

// prop value enum
func (m *EnvironmentSecret) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, environmentSecretTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EnvironmentSecret) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this environment secret based on context it is used
func (m *EnvironmentSecret) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EnvironmentSecret) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnvironmentSecret) UnmarshalBinary(b []byte) error {
	var res EnvironmentSecret
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}