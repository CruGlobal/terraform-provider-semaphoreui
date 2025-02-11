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

// AccessKeyRequest access key request
//
// swagger:model AccessKeyRequest
type AccessKeyRequest struct {

	// id
	ID int64 `json:"id,omitempty"`

	// login password
	LoginPassword *AccessKeyRequestLoginPassword `json:"login_password,omitempty"`

	// name
	// Example: None
	Name string `json:"name,omitempty"`

	// override secret
	OverrideSecret bool `json:"override_secret,omitempty"`

	// project id
	// Minimum: 1
	ProjectID int64 `json:"project_id,omitempty"`

	// ssh
	SSH *AccessKeyRequestSSH `json:"ssh,omitempty"`

	// type
	// Enum: ["none","ssh","login_password"]
	Type string `json:"type,omitempty"`
}

// Validate validates this access key request
func (m *AccessKeyRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLoginPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSSH(formats); err != nil {
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

func (m *AccessKeyRequest) validateLoginPassword(formats strfmt.Registry) error {
	if swag.IsZero(m.LoginPassword) { // not required
		return nil
	}

	if m.LoginPassword != nil {
		if err := m.LoginPassword.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("login_password")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("login_password")
			}
			return err
		}
	}

	return nil
}

func (m *AccessKeyRequest) validateProjectID(formats strfmt.Registry) error {
	if swag.IsZero(m.ProjectID) { // not required
		return nil
	}

	if err := validate.MinimumInt("project_id", "body", m.ProjectID, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *AccessKeyRequest) validateSSH(formats strfmt.Registry) error {
	if swag.IsZero(m.SSH) { // not required
		return nil
	}

	if m.SSH != nil {
		if err := m.SSH.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ssh")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ssh")
			}
			return err
		}
	}

	return nil
}

var accessKeyRequestTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","ssh","login_password"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accessKeyRequestTypeTypePropEnum = append(accessKeyRequestTypeTypePropEnum, v)
	}
}

const (

	// AccessKeyRequestTypeNone captures enum value "none"
	AccessKeyRequestTypeNone string = "none"

	// AccessKeyRequestTypeSSH captures enum value "ssh"
	AccessKeyRequestTypeSSH string = "ssh"

	// AccessKeyRequestTypeLoginPassword captures enum value "login_password"
	AccessKeyRequestTypeLoginPassword string = "login_password"
)

// prop value enum
func (m *AccessKeyRequest) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, accessKeyRequestTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AccessKeyRequest) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this access key request based on the context it is used
func (m *AccessKeyRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLoginPassword(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSSH(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccessKeyRequest) contextValidateLoginPassword(ctx context.Context, formats strfmt.Registry) error {

	if m.LoginPassword != nil {

		if swag.IsZero(m.LoginPassword) { // not required
			return nil
		}

		if err := m.LoginPassword.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("login_password")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("login_password")
			}
			return err
		}
	}

	return nil
}

func (m *AccessKeyRequest) contextValidateSSH(ctx context.Context, formats strfmt.Registry) error {

	if m.SSH != nil {

		if swag.IsZero(m.SSH) { // not required
			return nil
		}

		if err := m.SSH.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ssh")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ssh")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccessKeyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccessKeyRequest) UnmarshalBinary(b []byte) error {
	var res AccessKeyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AccessKeyRequestLoginPassword access key request login password
//
// swagger:model AccessKeyRequestLoginPassword
type AccessKeyRequestLoginPassword struct {

	// login
	// Example: username
	Login string `json:"login,omitempty"`

	// password
	// Example: password
	Password string `json:"password,omitempty"`
}

// Validate validates this access key request login password
func (m *AccessKeyRequestLoginPassword) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this access key request login password based on context it is used
func (m *AccessKeyRequestLoginPassword) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccessKeyRequestLoginPassword) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccessKeyRequestLoginPassword) UnmarshalBinary(b []byte) error {
	var res AccessKeyRequestLoginPassword
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AccessKeyRequestSSH access key request SSH
//
// swagger:model AccessKeyRequestSSH
type AccessKeyRequestSSH struct {

	// login
	// Example: user
	Login string `json:"login,omitempty"`

	// passphrase
	// Example: passphrase
	Passphrase string `json:"passphrase,omitempty"`

	// private key
	// Example: private key
	PrivateKey string `json:"private_key,omitempty"`
}

// Validate validates this access key request SSH
func (m *AccessKeyRequestSSH) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this access key request SSH based on context it is used
func (m *AccessKeyRequestSSH) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccessKeyRequestSSH) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccessKeyRequestSSH) UnmarshalBinary(b []byte) error {
	var res AccessKeyRequestSSH
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
