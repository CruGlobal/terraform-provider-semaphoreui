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

// AccessKey access key
//
// swagger:model AccessKey
type AccessKey struct {

	// id
	ID int64 `json:"id,omitempty"`

	// name
	// Example: Test
	Name string `json:"name,omitempty"`

	// project id
	ProjectID int64 `json:"project_id,omitempty"`

	// type
	// Enum: ["none","ssh","login_password"]
	Type string `json:"type,omitempty"`
}

// Validate validates this access key
func (m *AccessKey) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var accessKeyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","ssh","login_password"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accessKeyTypeTypePropEnum = append(accessKeyTypeTypePropEnum, v)
	}
}

const (

	// AccessKeyTypeNone captures enum value "none"
	AccessKeyTypeNone string = "none"

	// AccessKeyTypeSSH captures enum value "ssh"
	AccessKeyTypeSSH string = "ssh"

	// AccessKeyTypeLoginPassword captures enum value "login_password"
	AccessKeyTypeLoginPassword string = "login_password"
)

// prop value enum
func (m *AccessKey) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, accessKeyTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AccessKey) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this access key based on context it is used
func (m *AccessKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccessKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccessKey) UnmarshalBinary(b []byte) error {
	var res AccessKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}