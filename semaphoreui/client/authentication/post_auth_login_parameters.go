// Code generated by go-swagger; DO NOT EDIT.

package authentication

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

	"terraform-provider-semaphoreui/semaphoreui/models"
)

// NewPostAuthLoginParams creates a new PostAuthLoginParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAuthLoginParams() *PostAuthLoginParams {
	return &PostAuthLoginParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAuthLoginParamsWithTimeout creates a new PostAuthLoginParams object
// with the ability to set a timeout on a request.
func NewPostAuthLoginParamsWithTimeout(timeout time.Duration) *PostAuthLoginParams {
	return &PostAuthLoginParams{
		timeout: timeout,
	}
}

// NewPostAuthLoginParamsWithContext creates a new PostAuthLoginParams object
// with the ability to set a context for a request.
func NewPostAuthLoginParamsWithContext(ctx context.Context) *PostAuthLoginParams {
	return &PostAuthLoginParams{
		Context: ctx,
	}
}

// NewPostAuthLoginParamsWithHTTPClient creates a new PostAuthLoginParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAuthLoginParamsWithHTTPClient(client *http.Client) *PostAuthLoginParams {
	return &PostAuthLoginParams{
		HTTPClient: client,
	}
}

/*
PostAuthLoginParams contains all the parameters to send to the API endpoint

	for the post auth login operation.

	Typically these are written to a http.Request.
*/
type PostAuthLoginParams struct {

	// LoginBody.
	LoginBody *models.Login

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post auth login params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAuthLoginParams) WithDefaults() *PostAuthLoginParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post auth login params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAuthLoginParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post auth login params
func (o *PostAuthLoginParams) WithTimeout(timeout time.Duration) *PostAuthLoginParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post auth login params
func (o *PostAuthLoginParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post auth login params
func (o *PostAuthLoginParams) WithContext(ctx context.Context) *PostAuthLoginParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post auth login params
func (o *PostAuthLoginParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post auth login params
func (o *PostAuthLoginParams) WithHTTPClient(client *http.Client) *PostAuthLoginParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post auth login params
func (o *PostAuthLoginParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLoginBody adds the loginBody to the post auth login params
func (o *PostAuthLoginParams) WithLoginBody(loginBody *models.Login) *PostAuthLoginParams {
	o.SetLoginBody(loginBody)
	return o
}

// SetLoginBody adds the loginBody to the post auth login params
func (o *PostAuthLoginParams) SetLoginBody(loginBody *models.Login) {
	o.LoginBody = loginBody
}

// WriteToRequest writes these params to a swagger request
func (o *PostAuthLoginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.LoginBody != nil {
		if err := r.SetBodyParam(o.LoginBody); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}