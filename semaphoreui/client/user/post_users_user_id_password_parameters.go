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

// NewPostUsersUserIDPasswordParams creates a new PostUsersUserIDPasswordParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostUsersUserIDPasswordParams() *PostUsersUserIDPasswordParams {
	return &PostUsersUserIDPasswordParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostUsersUserIDPasswordParamsWithTimeout creates a new PostUsersUserIDPasswordParams object
// with the ability to set a timeout on a request.
func NewPostUsersUserIDPasswordParamsWithTimeout(timeout time.Duration) *PostUsersUserIDPasswordParams {
	return &PostUsersUserIDPasswordParams{
		timeout: timeout,
	}
}

// NewPostUsersUserIDPasswordParamsWithContext creates a new PostUsersUserIDPasswordParams object
// with the ability to set a context for a request.
func NewPostUsersUserIDPasswordParamsWithContext(ctx context.Context) *PostUsersUserIDPasswordParams {
	return &PostUsersUserIDPasswordParams{
		Context: ctx,
	}
}

// NewPostUsersUserIDPasswordParamsWithHTTPClient creates a new PostUsersUserIDPasswordParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostUsersUserIDPasswordParamsWithHTTPClient(client *http.Client) *PostUsersUserIDPasswordParams {
	return &PostUsersUserIDPasswordParams{
		HTTPClient: client,
	}
}

/*
PostUsersUserIDPasswordParams contains all the parameters to send to the API endpoint

	for the post users user ID password operation.

	Typically these are written to a http.Request.
*/
type PostUsersUserIDPasswordParams struct {

	// Password.
	Password PostUsersUserIDPasswordBody

	/* UserID.

	   User ID
	*/
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post users user ID password params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUsersUserIDPasswordParams) WithDefaults() *PostUsersUserIDPasswordParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post users user ID password params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUsersUserIDPasswordParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post users user ID password params
func (o *PostUsersUserIDPasswordParams) WithTimeout(timeout time.Duration) *PostUsersUserIDPasswordParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post users user ID password params
func (o *PostUsersUserIDPasswordParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post users user ID password params
func (o *PostUsersUserIDPasswordParams) WithContext(ctx context.Context) *PostUsersUserIDPasswordParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post users user ID password params
func (o *PostUsersUserIDPasswordParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post users user ID password params
func (o *PostUsersUserIDPasswordParams) WithHTTPClient(client *http.Client) *PostUsersUserIDPasswordParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post users user ID password params
func (o *PostUsersUserIDPasswordParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPassword adds the password to the post users user ID password params
func (o *PostUsersUserIDPasswordParams) WithPassword(password PostUsersUserIDPasswordBody) *PostUsersUserIDPasswordParams {
	o.SetPassword(password)
	return o
}

// SetPassword adds the password to the post users user ID password params
func (o *PostUsersUserIDPasswordParams) SetPassword(password PostUsersUserIDPasswordBody) {
	o.Password = password
}

// WithUserID adds the userID to the post users user ID password params
func (o *PostUsersUserIDPasswordParams) WithUserID(userID int64) *PostUsersUserIDPasswordParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the post users user ID password params
func (o *PostUsersUserIDPasswordParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PostUsersUserIDPasswordParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Password); err != nil {
		return err
	}

	// path param user_id
	if err := r.SetPathParam("user_id", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
