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

// NewPostProjectProjectIDViewsParams creates a new PostProjectProjectIDViewsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostProjectProjectIDViewsParams() *PostProjectProjectIDViewsParams {
	return &PostProjectProjectIDViewsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostProjectProjectIDViewsParamsWithTimeout creates a new PostProjectProjectIDViewsParams object
// with the ability to set a timeout on a request.
func NewPostProjectProjectIDViewsParamsWithTimeout(timeout time.Duration) *PostProjectProjectIDViewsParams {
	return &PostProjectProjectIDViewsParams{
		timeout: timeout,
	}
}

// NewPostProjectProjectIDViewsParamsWithContext creates a new PostProjectProjectIDViewsParams object
// with the ability to set a context for a request.
func NewPostProjectProjectIDViewsParamsWithContext(ctx context.Context) *PostProjectProjectIDViewsParams {
	return &PostProjectProjectIDViewsParams{
		Context: ctx,
	}
}

// NewPostProjectProjectIDViewsParamsWithHTTPClient creates a new PostProjectProjectIDViewsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostProjectProjectIDViewsParamsWithHTTPClient(client *http.Client) *PostProjectProjectIDViewsParams {
	return &PostProjectProjectIDViewsParams{
		HTTPClient: client,
	}
}

/*
PostProjectProjectIDViewsParams contains all the parameters to send to the API endpoint

	for the post project project ID views operation.

	Typically these are written to a http.Request.
*/
type PostProjectProjectIDViewsParams struct {

	/* ProjectID.

	   Project ID
	*/
	ProjectID int64

	// View.
	View *models.ViewRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post project project ID views params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostProjectProjectIDViewsParams) WithDefaults() *PostProjectProjectIDViewsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post project project ID views params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostProjectProjectIDViewsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post project project ID views params
func (o *PostProjectProjectIDViewsParams) WithTimeout(timeout time.Duration) *PostProjectProjectIDViewsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post project project ID views params
func (o *PostProjectProjectIDViewsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post project project ID views params
func (o *PostProjectProjectIDViewsParams) WithContext(ctx context.Context) *PostProjectProjectIDViewsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post project project ID views params
func (o *PostProjectProjectIDViewsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post project project ID views params
func (o *PostProjectProjectIDViewsParams) WithHTTPClient(client *http.Client) *PostProjectProjectIDViewsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post project project ID views params
func (o *PostProjectProjectIDViewsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the post project project ID views params
func (o *PostProjectProjectIDViewsParams) WithProjectID(projectID int64) *PostProjectProjectIDViewsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the post project project ID views params
func (o *PostProjectProjectIDViewsParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WithView adds the view to the post project project ID views params
func (o *PostProjectProjectIDViewsParams) WithView(view *models.ViewRequest) *PostProjectProjectIDViewsParams {
	o.SetView(view)
	return o
}

// SetView adds the view to the post project project ID views params
func (o *PostProjectProjectIDViewsParams) SetView(view *models.ViewRequest) {
	o.View = view
}

// WriteToRequest writes these params to a swagger request
func (o *PostProjectProjectIDViewsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param project_id
	if err := r.SetPathParam("project_id", swag.FormatInt64(o.ProjectID)); err != nil {
		return err
	}
	if o.View != nil {
		if err := r.SetBodyParam(o.View); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
