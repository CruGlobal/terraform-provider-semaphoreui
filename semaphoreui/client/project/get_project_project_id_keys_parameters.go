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

// NewGetProjectProjectIDKeysParams creates a new GetProjectProjectIDKeysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProjectProjectIDKeysParams() *GetProjectProjectIDKeysParams {
	return &GetProjectProjectIDKeysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectProjectIDKeysParamsWithTimeout creates a new GetProjectProjectIDKeysParams object
// with the ability to set a timeout on a request.
func NewGetProjectProjectIDKeysParamsWithTimeout(timeout time.Duration) *GetProjectProjectIDKeysParams {
	return &GetProjectProjectIDKeysParams{
		timeout: timeout,
	}
}

// NewGetProjectProjectIDKeysParamsWithContext creates a new GetProjectProjectIDKeysParams object
// with the ability to set a context for a request.
func NewGetProjectProjectIDKeysParamsWithContext(ctx context.Context) *GetProjectProjectIDKeysParams {
	return &GetProjectProjectIDKeysParams{
		Context: ctx,
	}
}

// NewGetProjectProjectIDKeysParamsWithHTTPClient creates a new GetProjectProjectIDKeysParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProjectProjectIDKeysParamsWithHTTPClient(client *http.Client) *GetProjectProjectIDKeysParams {
	return &GetProjectProjectIDKeysParams{
		HTTPClient: client,
	}
}

/*
GetProjectProjectIDKeysParams contains all the parameters to send to the API endpoint

	for the get project project ID keys operation.

	Typically these are written to a http.Request.
*/
type GetProjectProjectIDKeysParams struct {

	/* KeyType.

	   Filter by key type
	*/
	KeyType *string

	/* Order.

	   ordering manner
	*/
	Order string

	/* ProjectID.

	   Project ID
	*/
	ProjectID int64

	/* Sort.

	   sorting name
	*/
	Sort string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get project project ID keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectProjectIDKeysParams) WithDefaults() *GetProjectProjectIDKeysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get project project ID keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectProjectIDKeysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get project project ID keys params
func (o *GetProjectProjectIDKeysParams) WithTimeout(timeout time.Duration) *GetProjectProjectIDKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get project project ID keys params
func (o *GetProjectProjectIDKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get project project ID keys params
func (o *GetProjectProjectIDKeysParams) WithContext(ctx context.Context) *GetProjectProjectIDKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get project project ID keys params
func (o *GetProjectProjectIDKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get project project ID keys params
func (o *GetProjectProjectIDKeysParams) WithHTTPClient(client *http.Client) *GetProjectProjectIDKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get project project ID keys params
func (o *GetProjectProjectIDKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKeyType adds the keyType to the get project project ID keys params
func (o *GetProjectProjectIDKeysParams) WithKeyType(keyType *string) *GetProjectProjectIDKeysParams {
	o.SetKeyType(keyType)
	return o
}

// SetKeyType adds the keyType to the get project project ID keys params
func (o *GetProjectProjectIDKeysParams) SetKeyType(keyType *string) {
	o.KeyType = keyType
}

// WithOrder adds the order to the get project project ID keys params
func (o *GetProjectProjectIDKeysParams) WithOrder(order string) *GetProjectProjectIDKeysParams {
	o.SetOrder(order)
	return o
}

// SetOrder adds the order to the get project project ID keys params
func (o *GetProjectProjectIDKeysParams) SetOrder(order string) {
	o.Order = order
}

// WithProjectID adds the projectID to the get project project ID keys params
func (o *GetProjectProjectIDKeysParams) WithProjectID(projectID int64) *GetProjectProjectIDKeysParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get project project ID keys params
func (o *GetProjectProjectIDKeysParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WithSort adds the sort to the get project project ID keys params
func (o *GetProjectProjectIDKeysParams) WithSort(sort string) *GetProjectProjectIDKeysParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get project project ID keys params
func (o *GetProjectProjectIDKeysParams) SetSort(sort string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectProjectIDKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.KeyType != nil {

		// query param Key type
		var qrKeyType string

		if o.KeyType != nil {
			qrKeyType = *o.KeyType
		}
		qKeyType := qrKeyType
		if qKeyType != "" {

			if err := r.SetQueryParam("Key type", qKeyType); err != nil {
				return err
			}
		}
	}

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
