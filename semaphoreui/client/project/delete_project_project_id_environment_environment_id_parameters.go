// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeleteProjectProjectIDEnvironmentEnvironmentIDParams creates a new DeleteProjectProjectIDEnvironmentEnvironmentIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteProjectProjectIDEnvironmentEnvironmentIDParams() *DeleteProjectProjectIDEnvironmentEnvironmentIDParams {
	return &DeleteProjectProjectIDEnvironmentEnvironmentIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteProjectProjectIDEnvironmentEnvironmentIDParamsWithTimeout creates a new DeleteProjectProjectIDEnvironmentEnvironmentIDParams object
// with the ability to set a timeout on a request.
func NewDeleteProjectProjectIDEnvironmentEnvironmentIDParamsWithTimeout(timeout time.Duration) *DeleteProjectProjectIDEnvironmentEnvironmentIDParams {
	return &DeleteProjectProjectIDEnvironmentEnvironmentIDParams{
		timeout: timeout,
	}
}

// NewDeleteProjectProjectIDEnvironmentEnvironmentIDParamsWithContext creates a new DeleteProjectProjectIDEnvironmentEnvironmentIDParams object
// with the ability to set a context for a request.
func NewDeleteProjectProjectIDEnvironmentEnvironmentIDParamsWithContext(ctx context.Context) *DeleteProjectProjectIDEnvironmentEnvironmentIDParams {
	return &DeleteProjectProjectIDEnvironmentEnvironmentIDParams{
		Context: ctx,
	}
}

// NewDeleteProjectProjectIDEnvironmentEnvironmentIDParamsWithHTTPClient creates a new DeleteProjectProjectIDEnvironmentEnvironmentIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteProjectProjectIDEnvironmentEnvironmentIDParamsWithHTTPClient(client *http.Client) *DeleteProjectProjectIDEnvironmentEnvironmentIDParams {
	return &DeleteProjectProjectIDEnvironmentEnvironmentIDParams{
		HTTPClient: client,
	}
}

/*
DeleteProjectProjectIDEnvironmentEnvironmentIDParams contains all the parameters to send to the API endpoint

	for the delete project project ID environment environment ID operation.

	Typically these are written to a http.Request.
*/
type DeleteProjectProjectIDEnvironmentEnvironmentIDParams struct {

	/* EnvironmentID.

	   environment ID
	*/
	EnvironmentID int64

	/* ProjectID.

	   Project ID
	*/
	ProjectID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete project project ID environment environment ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteProjectProjectIDEnvironmentEnvironmentIDParams) WithDefaults() *DeleteProjectProjectIDEnvironmentEnvironmentIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete project project ID environment environment ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteProjectProjectIDEnvironmentEnvironmentIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete project project ID environment environment ID params
func (o *DeleteProjectProjectIDEnvironmentEnvironmentIDParams) WithTimeout(timeout time.Duration) *DeleteProjectProjectIDEnvironmentEnvironmentIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete project project ID environment environment ID params
func (o *DeleteProjectProjectIDEnvironmentEnvironmentIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete project project ID environment environment ID params
func (o *DeleteProjectProjectIDEnvironmentEnvironmentIDParams) WithContext(ctx context.Context) *DeleteProjectProjectIDEnvironmentEnvironmentIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete project project ID environment environment ID params
func (o *DeleteProjectProjectIDEnvironmentEnvironmentIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete project project ID environment environment ID params
func (o *DeleteProjectProjectIDEnvironmentEnvironmentIDParams) WithHTTPClient(client *http.Client) *DeleteProjectProjectIDEnvironmentEnvironmentIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete project project ID environment environment ID params
func (o *DeleteProjectProjectIDEnvironmentEnvironmentIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentID adds the environmentID to the delete project project ID environment environment ID params
func (o *DeleteProjectProjectIDEnvironmentEnvironmentIDParams) WithEnvironmentID(environmentID int64) *DeleteProjectProjectIDEnvironmentEnvironmentIDParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the delete project project ID environment environment ID params
func (o *DeleteProjectProjectIDEnvironmentEnvironmentIDParams) SetEnvironmentID(environmentID int64) {
	o.EnvironmentID = environmentID
}

// WithProjectID adds the projectID to the delete project project ID environment environment ID params
func (o *DeleteProjectProjectIDEnvironmentEnvironmentIDParams) WithProjectID(projectID int64) *DeleteProjectProjectIDEnvironmentEnvironmentIDParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete project project ID environment environment ID params
func (o *DeleteProjectProjectIDEnvironmentEnvironmentIDParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteProjectProjectIDEnvironmentEnvironmentIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param environment_id
	if err := r.SetPathParam("environment_id", swag.FormatInt64(o.EnvironmentID)); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", swag.FormatInt64(o.ProjectID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}