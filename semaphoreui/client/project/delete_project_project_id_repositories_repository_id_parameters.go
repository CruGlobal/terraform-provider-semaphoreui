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

// NewDeleteProjectProjectIDRepositoriesRepositoryIDParams creates a new DeleteProjectProjectIDRepositoriesRepositoryIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteProjectProjectIDRepositoriesRepositoryIDParams() *DeleteProjectProjectIDRepositoriesRepositoryIDParams {
	return &DeleteProjectProjectIDRepositoriesRepositoryIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteProjectProjectIDRepositoriesRepositoryIDParamsWithTimeout creates a new DeleteProjectProjectIDRepositoriesRepositoryIDParams object
// with the ability to set a timeout on a request.
func NewDeleteProjectProjectIDRepositoriesRepositoryIDParamsWithTimeout(timeout time.Duration) *DeleteProjectProjectIDRepositoriesRepositoryIDParams {
	return &DeleteProjectProjectIDRepositoriesRepositoryIDParams{
		timeout: timeout,
	}
}

// NewDeleteProjectProjectIDRepositoriesRepositoryIDParamsWithContext creates a new DeleteProjectProjectIDRepositoriesRepositoryIDParams object
// with the ability to set a context for a request.
func NewDeleteProjectProjectIDRepositoriesRepositoryIDParamsWithContext(ctx context.Context) *DeleteProjectProjectIDRepositoriesRepositoryIDParams {
	return &DeleteProjectProjectIDRepositoriesRepositoryIDParams{
		Context: ctx,
	}
}

// NewDeleteProjectProjectIDRepositoriesRepositoryIDParamsWithHTTPClient creates a new DeleteProjectProjectIDRepositoriesRepositoryIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteProjectProjectIDRepositoriesRepositoryIDParamsWithHTTPClient(client *http.Client) *DeleteProjectProjectIDRepositoriesRepositoryIDParams {
	return &DeleteProjectProjectIDRepositoriesRepositoryIDParams{
		HTTPClient: client,
	}
}

/*
DeleteProjectProjectIDRepositoriesRepositoryIDParams contains all the parameters to send to the API endpoint

	for the delete project project ID repositories repository ID operation.

	Typically these are written to a http.Request.
*/
type DeleteProjectProjectIDRepositoriesRepositoryIDParams struct {

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

// WithDefaults hydrates default values in the delete project project ID repositories repository ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteProjectProjectIDRepositoriesRepositoryIDParams) WithDefaults() *DeleteProjectProjectIDRepositoriesRepositoryIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete project project ID repositories repository ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteProjectProjectIDRepositoriesRepositoryIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete project project ID repositories repository ID params
func (o *DeleteProjectProjectIDRepositoriesRepositoryIDParams) WithTimeout(timeout time.Duration) *DeleteProjectProjectIDRepositoriesRepositoryIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete project project ID repositories repository ID params
func (o *DeleteProjectProjectIDRepositoriesRepositoryIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete project project ID repositories repository ID params
func (o *DeleteProjectProjectIDRepositoriesRepositoryIDParams) WithContext(ctx context.Context) *DeleteProjectProjectIDRepositoriesRepositoryIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete project project ID repositories repository ID params
func (o *DeleteProjectProjectIDRepositoriesRepositoryIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete project project ID repositories repository ID params
func (o *DeleteProjectProjectIDRepositoriesRepositoryIDParams) WithHTTPClient(client *http.Client) *DeleteProjectProjectIDRepositoriesRepositoryIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete project project ID repositories repository ID params
func (o *DeleteProjectProjectIDRepositoriesRepositoryIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the delete project project ID repositories repository ID params
func (o *DeleteProjectProjectIDRepositoriesRepositoryIDParams) WithProjectID(projectID int64) *DeleteProjectProjectIDRepositoriesRepositoryIDParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete project project ID repositories repository ID params
func (o *DeleteProjectProjectIDRepositoriesRepositoryIDParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WithRepositoryID adds the repositoryID to the delete project project ID repositories repository ID params
func (o *DeleteProjectProjectIDRepositoriesRepositoryIDParams) WithRepositoryID(repositoryID int64) *DeleteProjectProjectIDRepositoriesRepositoryIDParams {
	o.SetRepositoryID(repositoryID)
	return o
}

// SetRepositoryID adds the repositoryId to the delete project project ID repositories repository ID params
func (o *DeleteProjectProjectIDRepositoriesRepositoryIDParams) SetRepositoryID(repositoryID int64) {
	o.RepositoryID = repositoryID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteProjectProjectIDRepositoriesRepositoryIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
