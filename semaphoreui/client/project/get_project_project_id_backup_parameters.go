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

// NewGetProjectProjectIDBackupParams creates a new GetProjectProjectIDBackupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProjectProjectIDBackupParams() *GetProjectProjectIDBackupParams {
	return &GetProjectProjectIDBackupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectProjectIDBackupParamsWithTimeout creates a new GetProjectProjectIDBackupParams object
// with the ability to set a timeout on a request.
func NewGetProjectProjectIDBackupParamsWithTimeout(timeout time.Duration) *GetProjectProjectIDBackupParams {
	return &GetProjectProjectIDBackupParams{
		timeout: timeout,
	}
}

// NewGetProjectProjectIDBackupParamsWithContext creates a new GetProjectProjectIDBackupParams object
// with the ability to set a context for a request.
func NewGetProjectProjectIDBackupParamsWithContext(ctx context.Context) *GetProjectProjectIDBackupParams {
	return &GetProjectProjectIDBackupParams{
		Context: ctx,
	}
}

// NewGetProjectProjectIDBackupParamsWithHTTPClient creates a new GetProjectProjectIDBackupParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProjectProjectIDBackupParamsWithHTTPClient(client *http.Client) *GetProjectProjectIDBackupParams {
	return &GetProjectProjectIDBackupParams{
		HTTPClient: client,
	}
}

/*
GetProjectProjectIDBackupParams contains all the parameters to send to the API endpoint

	for the get project project ID backup operation.

	Typically these are written to a http.Request.
*/
type GetProjectProjectIDBackupParams struct {

	/* ProjectID.

	   Project ID
	*/
	ProjectID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get project project ID backup params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectProjectIDBackupParams) WithDefaults() *GetProjectProjectIDBackupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get project project ID backup params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectProjectIDBackupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get project project ID backup params
func (o *GetProjectProjectIDBackupParams) WithTimeout(timeout time.Duration) *GetProjectProjectIDBackupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get project project ID backup params
func (o *GetProjectProjectIDBackupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get project project ID backup params
func (o *GetProjectProjectIDBackupParams) WithContext(ctx context.Context) *GetProjectProjectIDBackupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get project project ID backup params
func (o *GetProjectProjectIDBackupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get project project ID backup params
func (o *GetProjectProjectIDBackupParams) WithHTTPClient(client *http.Client) *GetProjectProjectIDBackupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get project project ID backup params
func (o *GetProjectProjectIDBackupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the get project project ID backup params
func (o *GetProjectProjectIDBackupParams) WithProjectID(projectID int64) *GetProjectProjectIDBackupParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get project project ID backup params
func (o *GetProjectProjectIDBackupParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectProjectIDBackupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
