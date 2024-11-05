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

// NewGetProjectProjectIDTasksLastParams creates a new GetProjectProjectIDTasksLastParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProjectProjectIDTasksLastParams() *GetProjectProjectIDTasksLastParams {
	return &GetProjectProjectIDTasksLastParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectProjectIDTasksLastParamsWithTimeout creates a new GetProjectProjectIDTasksLastParams object
// with the ability to set a timeout on a request.
func NewGetProjectProjectIDTasksLastParamsWithTimeout(timeout time.Duration) *GetProjectProjectIDTasksLastParams {
	return &GetProjectProjectIDTasksLastParams{
		timeout: timeout,
	}
}

// NewGetProjectProjectIDTasksLastParamsWithContext creates a new GetProjectProjectIDTasksLastParams object
// with the ability to set a context for a request.
func NewGetProjectProjectIDTasksLastParamsWithContext(ctx context.Context) *GetProjectProjectIDTasksLastParams {
	return &GetProjectProjectIDTasksLastParams{
		Context: ctx,
	}
}

// NewGetProjectProjectIDTasksLastParamsWithHTTPClient creates a new GetProjectProjectIDTasksLastParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProjectProjectIDTasksLastParamsWithHTTPClient(client *http.Client) *GetProjectProjectIDTasksLastParams {
	return &GetProjectProjectIDTasksLastParams{
		HTTPClient: client,
	}
}

/*
GetProjectProjectIDTasksLastParams contains all the parameters to send to the API endpoint

	for the get project project ID tasks last operation.

	Typically these are written to a http.Request.
*/
type GetProjectProjectIDTasksLastParams struct {

	/* ProjectID.

	   Project ID
	*/
	ProjectID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get project project ID tasks last params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectProjectIDTasksLastParams) WithDefaults() *GetProjectProjectIDTasksLastParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get project project ID tasks last params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectProjectIDTasksLastParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get project project ID tasks last params
func (o *GetProjectProjectIDTasksLastParams) WithTimeout(timeout time.Duration) *GetProjectProjectIDTasksLastParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get project project ID tasks last params
func (o *GetProjectProjectIDTasksLastParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get project project ID tasks last params
func (o *GetProjectProjectIDTasksLastParams) WithContext(ctx context.Context) *GetProjectProjectIDTasksLastParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get project project ID tasks last params
func (o *GetProjectProjectIDTasksLastParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get project project ID tasks last params
func (o *GetProjectProjectIDTasksLastParams) WithHTTPClient(client *http.Client) *GetProjectProjectIDTasksLastParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get project project ID tasks last params
func (o *GetProjectProjectIDTasksLastParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the get project project ID tasks last params
func (o *GetProjectProjectIDTasksLastParams) WithProjectID(projectID int64) *GetProjectProjectIDTasksLastParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get project project ID tasks last params
func (o *GetProjectProjectIDTasksLastParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectProjectIDTasksLastParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param project_id
	if err := r.SetPathParam("project_id", swag.FormatInt64(o.ProjectID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
