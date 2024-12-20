// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewGetUsersUserIDParams creates a new GetUsersUserIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUsersUserIDParams() *GetUsersUserIDParams {
	return &GetUsersUserIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsersUserIDParamsWithTimeout creates a new GetUsersUserIDParams object
// with the ability to set a timeout on a request.
func NewGetUsersUserIDParamsWithTimeout(timeout time.Duration) *GetUsersUserIDParams {
	return &GetUsersUserIDParams{
		timeout: timeout,
	}
}

// NewGetUsersUserIDParamsWithContext creates a new GetUsersUserIDParams object
// with the ability to set a context for a request.
func NewGetUsersUserIDParamsWithContext(ctx context.Context) *GetUsersUserIDParams {
	return &GetUsersUserIDParams{
		Context: ctx,
	}
}

// NewGetUsersUserIDParamsWithHTTPClient creates a new GetUsersUserIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUsersUserIDParamsWithHTTPClient(client *http.Client) *GetUsersUserIDParams {
	return &GetUsersUserIDParams{
		HTTPClient: client,
	}
}

/*
GetUsersUserIDParams contains all the parameters to send to the API endpoint

	for the get users user ID operation.

	Typically these are written to a http.Request.
*/
type GetUsersUserIDParams struct {

	/* UserID.

	   User ID
	*/
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get users user ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsersUserIDParams) WithDefaults() *GetUsersUserIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get users user ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsersUserIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get users user ID params
func (o *GetUsersUserIDParams) WithTimeout(timeout time.Duration) *GetUsersUserIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get users user ID params
func (o *GetUsersUserIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get users user ID params
func (o *GetUsersUserIDParams) WithContext(ctx context.Context) *GetUsersUserIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get users user ID params
func (o *GetUsersUserIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get users user ID params
func (o *GetUsersUserIDParams) WithHTTPClient(client *http.Client) *GetUsersUserIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get users user ID params
func (o *GetUsersUserIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the get users user ID params
func (o *GetUsersUserIDParams) WithUserID(userID int64) *GetUsersUserIDParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get users user ID params
func (o *GetUsersUserIDParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsersUserIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param user_id
	if err := r.SetPathParam("user_id", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
