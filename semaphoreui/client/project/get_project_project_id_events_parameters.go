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

// NewGetProjectProjectIDEventsParams creates a new GetProjectProjectIDEventsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProjectProjectIDEventsParams() *GetProjectProjectIDEventsParams {
	return &GetProjectProjectIDEventsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectProjectIDEventsParamsWithTimeout creates a new GetProjectProjectIDEventsParams object
// with the ability to set a timeout on a request.
func NewGetProjectProjectIDEventsParamsWithTimeout(timeout time.Duration) *GetProjectProjectIDEventsParams {
	return &GetProjectProjectIDEventsParams{
		timeout: timeout,
	}
}

// NewGetProjectProjectIDEventsParamsWithContext creates a new GetProjectProjectIDEventsParams object
// with the ability to set a context for a request.
func NewGetProjectProjectIDEventsParamsWithContext(ctx context.Context) *GetProjectProjectIDEventsParams {
	return &GetProjectProjectIDEventsParams{
		Context: ctx,
	}
}

// NewGetProjectProjectIDEventsParamsWithHTTPClient creates a new GetProjectProjectIDEventsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProjectProjectIDEventsParamsWithHTTPClient(client *http.Client) *GetProjectProjectIDEventsParams {
	return &GetProjectProjectIDEventsParams{
		HTTPClient: client,
	}
}

/*
GetProjectProjectIDEventsParams contains all the parameters to send to the API endpoint

	for the get project project ID events operation.

	Typically these are written to a http.Request.
*/
type GetProjectProjectIDEventsParams struct {

	/* ProjectID.

	   Project ID
	*/
	ProjectID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get project project ID events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectProjectIDEventsParams) WithDefaults() *GetProjectProjectIDEventsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get project project ID events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectProjectIDEventsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get project project ID events params
func (o *GetProjectProjectIDEventsParams) WithTimeout(timeout time.Duration) *GetProjectProjectIDEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get project project ID events params
func (o *GetProjectProjectIDEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get project project ID events params
func (o *GetProjectProjectIDEventsParams) WithContext(ctx context.Context) *GetProjectProjectIDEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get project project ID events params
func (o *GetProjectProjectIDEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get project project ID events params
func (o *GetProjectProjectIDEventsParams) WithHTTPClient(client *http.Client) *GetProjectProjectIDEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get project project ID events params
func (o *GetProjectProjectIDEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the get project project ID events params
func (o *GetProjectProjectIDEventsParams) WithProjectID(projectID int64) *GetProjectProjectIDEventsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get project project ID events params
func (o *GetProjectProjectIDEventsParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectProjectIDEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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