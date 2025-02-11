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

// InventoryRequest inventory request
//
// swagger:model InventoryRequest
type InventoryRequest struct {

	// become key id
	// Minimum: 1
	BecomeKeyID int64 `json:"become_key_id,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// inventory
	Inventory string `json:"inventory,omitempty"`

	// name
	// Example: Test
	Name string `json:"name,omitempty"`

	// project id
	// Minimum: 1
	ProjectID int64 `json:"project_id,omitempty"`

	// repository id
	// Minimum: 1
	RepositoryID int64 `json:"repository_id,omitempty"`

	// ssh key id
	// Minimum: 1
	SSHKeyID int64 `json:"ssh_key_id,omitempty"`

	// type
	// Enum: ["static","static-yaml","file","terraform-workspace"]
	Type string `json:"type,omitempty"`
}

// Validate validates this inventory request
func (m *InventoryRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBecomeKeyID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepositoryID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSSHKeyID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InventoryRequest) validateBecomeKeyID(formats strfmt.Registry) error {
	if swag.IsZero(m.BecomeKeyID) { // not required
		return nil
	}

	if err := validate.MinimumInt("become_key_id", "body", m.BecomeKeyID, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *InventoryRequest) validateProjectID(formats strfmt.Registry) error {
	if swag.IsZero(m.ProjectID) { // not required
		return nil
	}

	if err := validate.MinimumInt("project_id", "body", m.ProjectID, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *InventoryRequest) validateRepositoryID(formats strfmt.Registry) error {
	if swag.IsZero(m.RepositoryID) { // not required
		return nil
	}

	if err := validate.MinimumInt("repository_id", "body", m.RepositoryID, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *InventoryRequest) validateSSHKeyID(formats strfmt.Registry) error {
	if swag.IsZero(m.SSHKeyID) { // not required
		return nil
	}

	if err := validate.MinimumInt("ssh_key_id", "body", m.SSHKeyID, 1, false); err != nil {
		return err
	}

	return nil
}

var inventoryRequestTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["static","static-yaml","file","terraform-workspace"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		inventoryRequestTypeTypePropEnum = append(inventoryRequestTypeTypePropEnum, v)
	}
}

const (

	// InventoryRequestTypeStatic captures enum value "static"
	InventoryRequestTypeStatic string = "static"

	// InventoryRequestTypeStaticDashYaml captures enum value "static-yaml"
	InventoryRequestTypeStaticDashYaml string = "static-yaml"

	// InventoryRequestTypeFile captures enum value "file"
	InventoryRequestTypeFile string = "file"

	// InventoryRequestTypeTerraformDashWorkspace captures enum value "terraform-workspace"
	InventoryRequestTypeTerraformDashWorkspace string = "terraform-workspace"
)

// prop value enum
func (m *InventoryRequest) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, inventoryRequestTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InventoryRequest) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this inventory request based on context it is used
func (m *InventoryRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InventoryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryRequest) UnmarshalBinary(b []byte) error {
	var res InventoryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
