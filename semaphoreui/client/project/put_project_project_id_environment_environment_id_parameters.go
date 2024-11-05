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

	"terraform-provider-semaphoreui/semaphoreui/models"
)

// NewPutProjectProjectIDEnvironmentEnvironmentIDParams creates a new PutProjectProjectIDEnvironmentEnvironmentIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutProjectProjectIDEnvironmentEnvironmentIDParams() *PutProjectProjectIDEnvironmentEnvironmentIDParams {
	return &PutProjectProjectIDEnvironmentEnvironmentIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutProjectProjectIDEnvironmentEnvironmentIDParamsWithTimeout creates a new PutProjectProjectIDEnvironmentEnvironmentIDParams object
// with the ability to set a timeout on a request.
func NewPutProjectProjectIDEnvironmentEnvironmentIDParamsWithTimeout(timeout time.Duration) *PutProjectProjectIDEnvironmentEnvironmentIDParams {
	return &PutProjectProjectIDEnvironmentEnvironmentIDParams{
		timeout: timeout,
	}
}

// NewPutProjectProjectIDEnvironmentEnvironmentIDParamsWithContext creates a new PutProjectProjectIDEnvironmentEnvironmentIDParams object
// with the ability to set a context for a request.
func NewPutProjectProjectIDEnvironmentEnvironmentIDParamsWithContext(ctx context.Context) *PutProjectProjectIDEnvironmentEnvironmentIDParams {
	return &PutProjectProjectIDEnvironmentEnvironmentIDParams{
		Context: ctx,
	}
}

// NewPutProjectProjectIDEnvironmentEnvironmentIDParamsWithHTTPClient creates a new PutProjectProjectIDEnvironmentEnvironmentIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutProjectProjectIDEnvironmentEnvironmentIDParamsWithHTTPClient(client *http.Client) *PutProjectProjectIDEnvironmentEnvironmentIDParams {
	return &PutProjectProjectIDEnvironmentEnvironmentIDParams{
		HTTPClient: client,
	}
}

/*
PutProjectProjectIDEnvironmentEnvironmentIDParams contains all the parameters to send to the API endpoint

	for the put project project ID environment environment ID operation.

	Typically these are written to a http.Request.
*/
type PutProjectProjectIDEnvironmentEnvironmentIDParams struct {

	// Environment.
	Environment *models.EnvironmentRequest

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

// WithDefaults hydrates default values in the put project project ID environment environment ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutProjectProjectIDEnvironmentEnvironmentIDParams) WithDefaults() *PutProjectProjectIDEnvironmentEnvironmentIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put project project ID environment environment ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutProjectProjectIDEnvironmentEnvironmentIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put project project ID environment environment ID params
func (o *PutProjectProjectIDEnvironmentEnvironmentIDParams) WithTimeout(timeout time.Duration) *PutProjectProjectIDEnvironmentEnvironmentIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put project project ID environment environment ID params
func (o *PutProjectProjectIDEnvironmentEnvironmentIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put project project ID environment environment ID params
func (o *PutProjectProjectIDEnvironmentEnvironmentIDParams) WithContext(ctx context.Context) *PutProjectProjectIDEnvironmentEnvironmentIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put project project ID environment environment ID params
func (o *PutProjectProjectIDEnvironmentEnvironmentIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put project project ID environment environment ID params
func (o *PutProjectProjectIDEnvironmentEnvironmentIDParams) WithHTTPClient(client *http.Client) *PutProjectProjectIDEnvironmentEnvironmentIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put project project ID environment environment ID params
func (o *PutProjectProjectIDEnvironmentEnvironmentIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironment adds the environment to the put project project ID environment environment ID params
func (o *PutProjectProjectIDEnvironmentEnvironmentIDParams) WithEnvironment(environment *models.EnvironmentRequest) *PutProjectProjectIDEnvironmentEnvironmentIDParams {
	o.SetEnvironment(environment)
	return o
}

// SetEnvironment adds the environment to the put project project ID environment environment ID params
func (o *PutProjectProjectIDEnvironmentEnvironmentIDParams) SetEnvironment(environment *models.EnvironmentRequest) {
	o.Environment = environment
}

// WithEnvironmentID adds the environmentID to the put project project ID environment environment ID params
func (o *PutProjectProjectIDEnvironmentEnvironmentIDParams) WithEnvironmentID(environmentID int64) *PutProjectProjectIDEnvironmentEnvironmentIDParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the put project project ID environment environment ID params
func (o *PutProjectProjectIDEnvironmentEnvironmentIDParams) SetEnvironmentID(environmentID int64) {
	o.EnvironmentID = environmentID
}

// WithProjectID adds the projectID to the put project project ID environment environment ID params
func (o *PutProjectProjectIDEnvironmentEnvironmentIDParams) WithProjectID(projectID int64) *PutProjectProjectIDEnvironmentEnvironmentIDParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the put project project ID environment environment ID params
func (o *PutProjectProjectIDEnvironmentEnvironmentIDParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *PutProjectProjectIDEnvironmentEnvironmentIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Environment != nil {
		if err := r.SetBodyParam(o.Environment); err != nil {
			return err
		}
	}

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
