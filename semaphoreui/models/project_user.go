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

// ProjectUser project user
//
// swagger:model ProjectUser
type ProjectUser struct {

	// id
	// Minimum: 1
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// role
	// Enum: ["owner","manager","task_runner","guest"]
	Role string `json:"role,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this project user
func (m *ProjectUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectUser) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinimumInt("id", "body", m.ID, 1, false); err != nil {
		return err
	}

	return nil
}

var projectUserTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["owner","manager","task_runner","guest"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		projectUserTypeRolePropEnum = append(projectUserTypeRolePropEnum, v)
	}
}

const (

	// ProjectUserRoleOwner captures enum value "owner"
	ProjectUserRoleOwner string = "owner"

	// ProjectUserRoleManager captures enum value "manager"
	ProjectUserRoleManager string = "manager"

	// ProjectUserRoleTaskRunner captures enum value "task_runner"
	ProjectUserRoleTaskRunner string = "task_runner"

	// ProjectUserRoleGuest captures enum value "guest"
	ProjectUserRoleGuest string = "guest"
)

// prop value enum
func (m *ProjectUser) validateRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, projectUserTypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ProjectUser) validateRole(formats strfmt.Registry) error {
	if swag.IsZero(m.Role) { // not required
		return nil
	}

	// value enum
	if err := m.validateRoleEnum("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this project user based on context it is used
func (m *ProjectUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProjectUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectUser) UnmarshalBinary(b []byte) error {
	var res ProjectUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
