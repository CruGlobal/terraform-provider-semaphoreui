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

// Template template
//
// swagger:model Template
type Template struct {

	// allow override args in task
	// Example: false
	AllowOverrideArgsInTask bool `json:"allow_override_args_in_task,omitempty"`

	// app
	App string `json:"app,omitempty"`

	// arguments
	// Example: []
	Arguments string `json:"arguments,omitempty"`

	// autorun
	Autorun bool `json:"autorun,omitempty"`

	// build template id
	BuildTemplateID int64 `json:"build_template_id,omitempty"`

	// description
	// Example: Hello, World!
	Description string `json:"description,omitempty"`

	// environment id
	// Minimum: 1
	EnvironmentID int64 `json:"environment_id,omitempty"`

	// git branch
	// Example: main
	GitBranch string `json:"git_branch,omitempty"`

	// id
	// Minimum: 1
	ID int64 `json:"id,omitempty"`

	// inventory id
	// Minimum: 1
	InventoryID int64 `json:"inventory_id,omitempty"`

	// name
	// Example: Test
	Name string `json:"name,omitempty"`

	// playbook
	// Example: test.yml
	Playbook string `json:"playbook,omitempty"`

	// project id
	// Minimum: 1
	ProjectID int64 `json:"project_id,omitempty"`

	// repository id
	RepositoryID int64 `json:"repository_id,omitempty"`

	// start version
	StartVersion string `json:"start_version,omitempty"`

	// suppress success alerts
	SuppressSuccessAlerts bool `json:"suppress_success_alerts,omitempty"`

	// survey vars
	SurveyVars []*TemplateSurveyVar `json:"survey_vars"`

	// type
	// Enum: ["","build","deploy"]
	Type string `json:"type,omitempty"`

	// vaults
	Vaults []*TemplateVault `json:"vaults"`

	// view id
	// Minimum: 1
	ViewID *int64 `json:"view_id,omitempty"`
}

// Validate validates this template
func (m *Template) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironmentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInventoryID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSurveyVars(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVaults(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateViewID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Template) validateEnvironmentID(formats strfmt.Registry) error {
	if swag.IsZero(m.EnvironmentID) { // not required
		return nil
	}

	if err := validate.MinimumInt("environment_id", "body", m.EnvironmentID, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *Template) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinimumInt("id", "body", m.ID, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *Template) validateInventoryID(formats strfmt.Registry) error {
	if swag.IsZero(m.InventoryID) { // not required
		return nil
	}

	if err := validate.MinimumInt("inventory_id", "body", m.InventoryID, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *Template) validateProjectID(formats strfmt.Registry) error {
	if swag.IsZero(m.ProjectID) { // not required
		return nil
	}

	if err := validate.MinimumInt("project_id", "body", m.ProjectID, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *Template) validateSurveyVars(formats strfmt.Registry) error {
	if swag.IsZero(m.SurveyVars) { // not required
		return nil
	}

	for i := 0; i < len(m.SurveyVars); i++ {
		if swag.IsZero(m.SurveyVars[i]) { // not required
			continue
		}

		if m.SurveyVars[i] != nil {
			if err := m.SurveyVars[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("survey_vars" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("survey_vars" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var templateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["","build","deploy"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		templateTypeTypePropEnum = append(templateTypeTypePropEnum, v)
	}
}

const (

	// TemplateTypeEmpty captures enum value ""
	TemplateTypeEmpty string = ""

	// TemplateTypeBuild captures enum value "build"
	TemplateTypeBuild string = "build"

	// TemplateTypeDeploy captures enum value "deploy"
	TemplateTypeDeploy string = "deploy"
)

// prop value enum
func (m *Template) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, templateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Template) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *Template) validateVaults(formats strfmt.Registry) error {
	if swag.IsZero(m.Vaults) { // not required
		return nil
	}

	for i := 0; i < len(m.Vaults); i++ {
		if swag.IsZero(m.Vaults[i]) { // not required
			continue
		}

		if m.Vaults[i] != nil {
			if err := m.Vaults[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vaults" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vaults" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Template) validateViewID(formats strfmt.Registry) error {
	if swag.IsZero(m.ViewID) { // not required
		return nil
	}

	if err := validate.MinimumInt("view_id", "body", *m.ViewID, 1, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this template based on the context it is used
func (m *Template) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSurveyVars(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVaults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Template) contextValidateSurveyVars(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SurveyVars); i++ {

		if m.SurveyVars[i] != nil {

			if swag.IsZero(m.SurveyVars[i]) { // not required
				return nil
			}

			if err := m.SurveyVars[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("survey_vars" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("survey_vars" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Template) contextValidateVaults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Vaults); i++ {

		if m.Vaults[i] != nil {

			if swag.IsZero(m.Vaults[i]) { // not required
				return nil
			}

			if err := m.Vaults[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vaults" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vaults" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Template) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Template) UnmarshalBinary(b []byte) error {
	var res Template
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
