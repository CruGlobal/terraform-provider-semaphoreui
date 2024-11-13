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

// NewPutProjectProjectIDRepositoriesRepositoryIDParams creates a new PutProjectProjectIDRepositoriesRepositoryIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutProjectProjectIDRepositoriesRepositoryIDParams() *PutProjectProjectIDRepositoriesRepositoryIDParams {
	return &PutProjectProjectIDRepositoriesRepositoryIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutProjectProjectIDRepositoriesRepositoryIDParamsWithTimeout creates a new PutProjectProjectIDRepositoriesRepositoryIDParams object
// with the ability to set a timeout on a request.
func NewPutProjectProjectIDRepositoriesRepositoryIDParamsWithTimeout(timeout time.Duration) *PutProjectProjectIDRepositoriesRepositoryIDParams {
	return &PutProjectProjectIDRepositoriesRepositoryIDParams{
		timeout: timeout,
	}
}

// NewPutProjectProjectIDRepositoriesRepositoryIDParamsWithContext creates a new PutProjectProjectIDRepositoriesRepositoryIDParams object
// with the ability to set a context for a request.
func NewPutProjectProjectIDRepositoriesRepositoryIDParamsWithContext(ctx context.Context) *PutProjectProjectIDRepositoriesRepositoryIDParams {
	return &PutProjectProjectIDRepositoriesRepositoryIDParams{
		Context: ctx,
	}
}

// NewPutProjectProjectIDRepositoriesRepositoryIDParamsWithHTTPClient creates a new PutProjectProjectIDRepositoriesRepositoryIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutProjectProjectIDRepositoriesRepositoryIDParamsWithHTTPClient(client *http.Client) *PutProjectProjectIDRepositoriesRepositoryIDParams {
	return &PutProjectProjectIDRepositoriesRepositoryIDParams{
		HTTPClient: client,
	}
}

/*
PutProjectProjectIDRepositoriesRepositoryIDParams contains all the parameters to send to the API endpoint

	for the put project project ID repositories repository ID operation.

	Typically these are written to a http.Request.
*/
type PutProjectProjectIDRepositoriesRepositoryIDParams struct {

	// Repository.
	Repository *models.RepositoryRequest

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

// WithDefaults hydrates default values in the put project project ID repositories repository ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutProjectProjectIDRepositoriesRepositoryIDParams) WithDefaults() *PutProjectProjectIDRepositoriesRepositoryIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put project project ID repositories repository ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutProjectProjectIDRepositoriesRepositoryIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put project project ID repositories repository ID params
func (o *PutProjectProjectIDRepositoriesRepositoryIDParams) WithTimeout(timeout time.Duration) *PutProjectProjectIDRepositoriesRepositoryIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put project project ID repositories repository ID params
func (o *PutProjectProjectIDRepositoriesRepositoryIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put project project ID repositories repository ID params
func (o *PutProjectProjectIDRepositoriesRepositoryIDParams) WithContext(ctx context.Context) *PutProjectProjectIDRepositoriesRepositoryIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put project project ID repositories repository ID params
func (o *PutProjectProjectIDRepositoriesRepositoryIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put project project ID repositories repository ID params
func (o *PutProjectProjectIDRepositoriesRepositoryIDParams) WithHTTPClient(client *http.Client) *PutProjectProjectIDRepositoriesRepositoryIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put project project ID repositories repository ID params
func (o *PutProjectProjectIDRepositoriesRepositoryIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepository adds the repository to the put project project ID repositories repository ID params
func (o *PutProjectProjectIDRepositoriesRepositoryIDParams) WithRepository(repository *models.RepositoryRequest) *PutProjectProjectIDRepositoriesRepositoryIDParams {
	o.SetRepository(repository)
	return o
}

// SetRepository adds the repository to the put project project ID repositories repository ID params
func (o *PutProjectProjectIDRepositoriesRepositoryIDParams) SetRepository(repository *models.RepositoryRequest) {
	o.Repository = repository
}

// WithProjectID adds the projectID to the put project project ID repositories repository ID params
func (o *PutProjectProjectIDRepositoriesRepositoryIDParams) WithProjectID(projectID int64) *PutProjectProjectIDRepositoriesRepositoryIDParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the put project project ID repositories repository ID params
func (o *PutProjectProjectIDRepositoriesRepositoryIDParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WithRepositoryID adds the repositoryID to the put project project ID repositories repository ID params
func (o *PutProjectProjectIDRepositoriesRepositoryIDParams) WithRepositoryID(repositoryID int64) *PutProjectProjectIDRepositoriesRepositoryIDParams {
	o.SetRepositoryID(repositoryID)
	return o
}

// SetRepositoryID adds the repositoryId to the put project project ID repositories repository ID params
func (o *PutProjectProjectIDRepositoriesRepositoryIDParams) SetRepositoryID(repositoryID int64) {
	o.RepositoryID = repositoryID
}

// WriteToRequest writes these params to a swagger request
func (o *PutProjectProjectIDRepositoriesRepositoryIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Repository != nil {
		if err := r.SetBodyParam(o.Repository); err != nil {
			return err
		}
	}

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