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

// NewGetProjectProjectIDEnvironmentParams creates a new GetProjectProjectIDEnvironmentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProjectProjectIDEnvironmentParams() *GetProjectProjectIDEnvironmentParams {
	return &GetProjectProjectIDEnvironmentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectProjectIDEnvironmentParamsWithTimeout creates a new GetProjectProjectIDEnvironmentParams object
// with the ability to set a timeout on a request.
func NewGetProjectProjectIDEnvironmentParamsWithTimeout(timeout time.Duration) *GetProjectProjectIDEnvironmentParams {
	return &GetProjectProjectIDEnvironmentParams{
		timeout: timeout,
	}
}

// NewGetProjectProjectIDEnvironmentParamsWithContext creates a new GetProjectProjectIDEnvironmentParams object
// with the ability to set a context for a request.
func NewGetProjectProjectIDEnvironmentParamsWithContext(ctx context.Context) *GetProjectProjectIDEnvironmentParams {
	return &GetProjectProjectIDEnvironmentParams{
		Context: ctx,
	}
}

// NewGetProjectProjectIDEnvironmentParamsWithHTTPClient creates a new GetProjectProjectIDEnvironmentParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProjectProjectIDEnvironmentParamsWithHTTPClient(client *http.Client) *GetProjectProjectIDEnvironmentParams {
	return &GetProjectProjectIDEnvironmentParams{
		HTTPClient: client,
	}
}

/*
GetProjectProjectIDEnvironmentParams contains all the parameters to send to the API endpoint

	for the get project project ID environment operation.

	Typically these are written to a http.Request.
*/
type GetProjectProjectIDEnvironmentParams struct {

	/* Order.

	   ordering manner

	   Format: asc/desc
	*/
	Order string

	/* ProjectID.

	   Project ID
	*/
	ProjectID int64

	/* Sort.

	   sorting name

	   Format: name
	*/
	Sort string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get project project ID environment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectProjectIDEnvironmentParams) WithDefaults() *GetProjectProjectIDEnvironmentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get project project ID environment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectProjectIDEnvironmentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get project project ID environment params
func (o *GetProjectProjectIDEnvironmentParams) WithTimeout(timeout time.Duration) *GetProjectProjectIDEnvironmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get project project ID environment params
func (o *GetProjectProjectIDEnvironmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get project project ID environment params
func (o *GetProjectProjectIDEnvironmentParams) WithContext(ctx context.Context) *GetProjectProjectIDEnvironmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get project project ID environment params
func (o *GetProjectProjectIDEnvironmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get project project ID environment params
func (o *GetProjectProjectIDEnvironmentParams) WithHTTPClient(client *http.Client) *GetProjectProjectIDEnvironmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get project project ID environment params
func (o *GetProjectProjectIDEnvironmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrder adds the order to the get project project ID environment params
func (o *GetProjectProjectIDEnvironmentParams) WithOrder(order string) *GetProjectProjectIDEnvironmentParams {
	o.SetOrder(order)
	return o
}

// SetOrder adds the order to the get project project ID environment params
func (o *GetProjectProjectIDEnvironmentParams) SetOrder(order string) {
	o.Order = order
}

// WithProjectID adds the projectID to the get project project ID environment params
func (o *GetProjectProjectIDEnvironmentParams) WithProjectID(projectID int64) *GetProjectProjectIDEnvironmentParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get project project ID environment params
func (o *GetProjectProjectIDEnvironmentParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WithSort adds the sort to the get project project ID environment params
func (o *GetProjectProjectIDEnvironmentParams) WithSort(sort string) *GetProjectProjectIDEnvironmentParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get project project ID environment params
func (o *GetProjectProjectIDEnvironmentParams) SetSort(sort string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectProjectIDEnvironmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param order
	qrOrder := o.Order
	qOrder := qrOrder
	if qOrder != "" {

		if err := r.SetQueryParam("order", qOrder); err != nil {
			return err
		}
	}

	// path param project_id
	if err := r.SetPathParam("project_id", swag.FormatInt64(o.ProjectID)); err != nil {
		return err
	}

	// query param sort
	qrSort := o.Sort
	qSort := qrSort
	if qSort != "" {

		if err := r.SetQueryParam("sort", qSort); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
