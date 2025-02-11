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

// NewGetProjectProjectIDRepositoriesRepositoryIDParams creates a new GetProjectProjectIDRepositoriesRepositoryIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProjectProjectIDRepositoriesRepositoryIDParams() *GetProjectProjectIDRepositoriesRepositoryIDParams {
	return &GetProjectProjectIDRepositoriesRepositoryIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectProjectIDRepositoriesRepositoryIDParamsWithTimeout creates a new GetProjectProjectIDRepositoriesRepositoryIDParams object
// with the ability to set a timeout on a request.
func NewGetProjectProjectIDRepositoriesRepositoryIDParamsWithTimeout(timeout time.Duration) *GetProjectProjectIDRepositoriesRepositoryIDParams {
	return &GetProjectProjectIDRepositoriesRepositoryIDParams{
		timeout: timeout,
	}
}

// NewGetProjectProjectIDRepositoriesRepositoryIDParamsWithContext creates a new GetProjectProjectIDRepositoriesRepositoryIDParams object
// with the ability to set a context for a request.
func NewGetProjectProjectIDRepositoriesRepositoryIDParamsWithContext(ctx context.Context) *GetProjectProjectIDRepositoriesRepositoryIDParams {
	return &GetProjectProjectIDRepositoriesRepositoryIDParams{
		Context: ctx,
	}
}

// NewGetProjectProjectIDRepositoriesRepositoryIDParamsWithHTTPClient creates a new GetProjectProjectIDRepositoriesRepositoryIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProjectProjectIDRepositoriesRepositoryIDParamsWithHTTPClient(client *http.Client) *GetProjectProjectIDRepositoriesRepositoryIDParams {
	return &GetProjectProjectIDRepositoriesRepositoryIDParams{
		HTTPClient: client,
	}
}

/*
GetProjectProjectIDRepositoriesRepositoryIDParams contains all the parameters to send to the API endpoint

	for the get project project ID repositories repository ID operation.

	Typically these are written to a http.Request.
*/
type GetProjectProjectIDRepositoriesRepositoryIDParams struct {

	/* ProjectID.

	   Project ID
	*/
	ProjectID int64

	/* RepositoryID.

	   repository ID
	*/
	RepositoryID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get project project ID repositories repository ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectProjectIDRepositoriesRepositoryIDParams) WithDefaults() *GetProjectProjectIDRepositoriesRepositoryIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get project project ID repositories repository ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectProjectIDRepositoriesRepositoryIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get project project ID repositories repository ID params
func (o *GetProjectProjectIDRepositoriesRepositoryIDParams) WithTimeout(timeout time.Duration) *GetProjectProjectIDRepositoriesRepositoryIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get project project ID repositories repository ID params
func (o *GetProjectProjectIDRepositoriesRepositoryIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get project project ID repositories repository ID params
func (o *GetProjectProjectIDRepositoriesRepositoryIDParams) WithContext(ctx context.Context) *GetProjectProjectIDRepositoriesRepositoryIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get project project ID repositories repository ID params
func (o *GetProjectProjectIDRepositoriesRepositoryIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get project project ID repositories repository ID params
func (o *GetProjectProjectIDRepositoriesRepositoryIDParams) WithHTTPClient(client *http.Client) *GetProjectProjectIDRepositoriesRepositoryIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get project project ID repositories repository ID params
func (o *GetProjectProjectIDRepositoriesRepositoryIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the get project project ID repositories repository ID params
func (o *GetProjectProjectIDRepositoriesRepositoryIDParams) WithProjectID(projectID int64) *GetProjectProjectIDRepositoriesRepositoryIDParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get project project ID repositories repository ID params
func (o *GetProjectProjectIDRepositoriesRepositoryIDParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WithRepositoryID adds the repositoryID to the get project project ID repositories repository ID params
func (o *GetProjectProjectIDRepositoriesRepositoryIDParams) WithRepositoryID(repositoryID int64) *GetProjectProjectIDRepositoriesRepositoryIDParams {
	o.SetRepositoryID(repositoryID)
	return o
}

// SetRepositoryID adds the repositoryId to the get project project ID repositories repository ID params
func (o *GetProjectProjectIDRepositoriesRepositoryIDParams) SetRepositoryID(repositoryID int64) {
	o.RepositoryID = repositoryID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectProjectIDRepositoriesRepositoryIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param project_id
	if err := r.SetPathParam("project_id", swag.FormatInt64(o.ProjectID)); err != nil {
		return err
	}

	// path param repository_id
	if err := r.SetPathParam("repository_id", swag.FormatInt64(o.RepositoryID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
