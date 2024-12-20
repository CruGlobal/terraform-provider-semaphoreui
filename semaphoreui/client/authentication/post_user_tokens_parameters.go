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
)

// NewPostUserTokensParams creates a new PostUserTokensParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostUserTokensParams() *PostUserTokensParams {
	return &PostUserTokensParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostUserTokensParamsWithTimeout creates a new PostUserTokensParams object
// with the ability to set a timeout on a request.
func NewPostUserTokensParamsWithTimeout(timeout time.Duration) *PostUserTokensParams {
	return &PostUserTokensParams{
		timeout: timeout,
	}
}

// NewPostUserTokensParamsWithContext creates a new PostUserTokensParams object
// with the ability to set a context for a request.
func NewPostUserTokensParamsWithContext(ctx context.Context) *PostUserTokensParams {
	return &PostUserTokensParams{
		Context: ctx,
	}
}

// NewPostUserTokensParamsWithHTTPClient creates a new PostUserTokensParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostUserTokensParamsWithHTTPClient(client *http.Client) *PostUserTokensParams {
	return &PostUserTokensParams{
		HTTPClient: client,
	}
}

/*
PostUserTokensParams contains all the parameters to send to the API endpoint

	for the post user tokens operation.

	Typically these are written to a http.Request.
*/
type PostUserTokensParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post user tokens params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUserTokensParams) WithDefaults() *PostUserTokensParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post user tokens params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUserTokensParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post user tokens params
func (o *PostUserTokensParams) WithTimeout(timeout time.Duration) *PostUserTokensParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post user tokens params
func (o *PostUserTokensParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post user tokens params
func (o *PostUserTokensParams) WithContext(ctx context.Context) *PostUserTokensParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post user tokens params
func (o *PostUserTokensParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post user tokens params
func (o *PostUserTokensParams) WithHTTPClient(client *http.Client) *PostUserTokensParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post user tokens params
func (o *PostUserTokensParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostUserTokensParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
