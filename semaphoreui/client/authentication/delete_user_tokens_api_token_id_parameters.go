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

// NewDeleteUserTokensAPITokenIDParams creates a new DeleteUserTokensAPITokenIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteUserTokensAPITokenIDParams() *DeleteUserTokensAPITokenIDParams {
	return &DeleteUserTokensAPITokenIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserTokensAPITokenIDParamsWithTimeout creates a new DeleteUserTokensAPITokenIDParams object
// with the ability to set a timeout on a request.
func NewDeleteUserTokensAPITokenIDParamsWithTimeout(timeout time.Duration) *DeleteUserTokensAPITokenIDParams {
	return &DeleteUserTokensAPITokenIDParams{
		timeout: timeout,
	}
}

// NewDeleteUserTokensAPITokenIDParamsWithContext creates a new DeleteUserTokensAPITokenIDParams object
// with the ability to set a context for a request.
func NewDeleteUserTokensAPITokenIDParamsWithContext(ctx context.Context) *DeleteUserTokensAPITokenIDParams {
	return &DeleteUserTokensAPITokenIDParams{
		Context: ctx,
	}
}

// NewDeleteUserTokensAPITokenIDParamsWithHTTPClient creates a new DeleteUserTokensAPITokenIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteUserTokensAPITokenIDParamsWithHTTPClient(client *http.Client) *DeleteUserTokensAPITokenIDParams {
	return &DeleteUserTokensAPITokenIDParams{
		HTTPClient: client,
	}
}

/*
DeleteUserTokensAPITokenIDParams contains all the parameters to send to the API endpoint

	for the delete user tokens API token ID operation.

	Typically these are written to a http.Request.
*/
type DeleteUserTokensAPITokenIDParams struct {

	// APITokenID.
	APITokenID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete user tokens API token ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUserTokensAPITokenIDParams) WithDefaults() *DeleteUserTokensAPITokenIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete user tokens API token ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUserTokensAPITokenIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete user tokens API token ID params
func (o *DeleteUserTokensAPITokenIDParams) WithTimeout(timeout time.Duration) *DeleteUserTokensAPITokenIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user tokens API token ID params
func (o *DeleteUserTokensAPITokenIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user tokens API token ID params
func (o *DeleteUserTokensAPITokenIDParams) WithContext(ctx context.Context) *DeleteUserTokensAPITokenIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user tokens API token ID params
func (o *DeleteUserTokensAPITokenIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user tokens API token ID params
func (o *DeleteUserTokensAPITokenIDParams) WithHTTPClient(client *http.Client) *DeleteUserTokensAPITokenIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user tokens API token ID params
func (o *DeleteUserTokensAPITokenIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPITokenID adds the aPITokenID to the delete user tokens API token ID params
func (o *DeleteUserTokensAPITokenIDParams) WithAPITokenID(aPITokenID string) *DeleteUserTokensAPITokenIDParams {
	o.SetAPITokenID(aPITokenID)
	return o
}

// SetAPITokenID adds the apiTokenId to the delete user tokens API token ID params
func (o *DeleteUserTokensAPITokenIDParams) SetAPITokenID(aPITokenID string) {
	o.APITokenID = aPITokenID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserTokensAPITokenIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param api_token_id
	if err := r.SetPathParam("api_token_id", o.APITokenID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
