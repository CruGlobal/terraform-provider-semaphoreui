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

// NewPutProjectProjectIDUsersUserIDParams creates a new PutProjectProjectIDUsersUserIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutProjectProjectIDUsersUserIDParams() *PutProjectProjectIDUsersUserIDParams {
	return &PutProjectProjectIDUsersUserIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutProjectProjectIDUsersUserIDParamsWithTimeout creates a new PutProjectProjectIDUsersUserIDParams object
// with the ability to set a timeout on a request.
func NewPutProjectProjectIDUsersUserIDParamsWithTimeout(timeout time.Duration) *PutProjectProjectIDUsersUserIDParams {
	return &PutProjectProjectIDUsersUserIDParams{
		timeout: timeout,
	}
}

// NewPutProjectProjectIDUsersUserIDParamsWithContext creates a new PutProjectProjectIDUsersUserIDParams object
// with the ability to set a context for a request.
func NewPutProjectProjectIDUsersUserIDParamsWithContext(ctx context.Context) *PutProjectProjectIDUsersUserIDParams {
	return &PutProjectProjectIDUsersUserIDParams{
		Context: ctx,
	}
}

// NewPutProjectProjectIDUsersUserIDParamsWithHTTPClient creates a new PutProjectProjectIDUsersUserIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutProjectProjectIDUsersUserIDParamsWithHTTPClient(client *http.Client) *PutProjectProjectIDUsersUserIDParams {
	return &PutProjectProjectIDUsersUserIDParams{
		HTTPClient: client,
	}
}

/*
PutProjectProjectIDUsersUserIDParams contains all the parameters to send to the API endpoint

	for the put project project ID users user ID operation.

	Typically these are written to a http.Request.
*/
type PutProjectProjectIDUsersUserIDParams struct {

	// ProjectUser.
	ProjectUser PutProjectProjectIDUsersUserIDBody

	/* ProjectID.

	   Project ID
	*/
	ProjectID int64

	/* UserID.

	   User ID
	*/
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put project project ID users user ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutProjectProjectIDUsersUserIDParams) WithDefaults() *PutProjectProjectIDUsersUserIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put project project ID users user ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutProjectProjectIDUsersUserIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put project project ID users user ID params
func (o *PutProjectProjectIDUsersUserIDParams) WithTimeout(timeout time.Duration) *PutProjectProjectIDUsersUserIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put project project ID users user ID params
func (o *PutProjectProjectIDUsersUserIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put project project ID users user ID params
func (o *PutProjectProjectIDUsersUserIDParams) WithContext(ctx context.Context) *PutProjectProjectIDUsersUserIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put project project ID users user ID params
func (o *PutProjectProjectIDUsersUserIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put project project ID users user ID params
func (o *PutProjectProjectIDUsersUserIDParams) WithHTTPClient(client *http.Client) *PutProjectProjectIDUsersUserIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put project project ID users user ID params
func (o *PutProjectProjectIDUsersUserIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectUser adds the projectUser to the put project project ID users user ID params
func (o *PutProjectProjectIDUsersUserIDParams) WithProjectUser(projectUser PutProjectProjectIDUsersUserIDBody) *PutProjectProjectIDUsersUserIDParams {
	o.SetProjectUser(projectUser)
	return o
}

// SetProjectUser adds the projectUser to the put project project ID users user ID params
func (o *PutProjectProjectIDUsersUserIDParams) SetProjectUser(projectUser PutProjectProjectIDUsersUserIDBody) {
	o.ProjectUser = projectUser
}

// WithProjectID adds the projectID to the put project project ID users user ID params
func (o *PutProjectProjectIDUsersUserIDParams) WithProjectID(projectID int64) *PutProjectProjectIDUsersUserIDParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the put project project ID users user ID params
func (o *PutProjectProjectIDUsersUserIDParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WithUserID adds the userID to the put project project ID users user ID params
func (o *PutProjectProjectIDUsersUserIDParams) WithUserID(userID int64) *PutProjectProjectIDUsersUserIDParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the put project project ID users user ID params
func (o *PutProjectProjectIDUsersUserIDParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PutProjectProjectIDUsersUserIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.ProjectUser); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", swag.FormatInt64(o.ProjectID)); err != nil {
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
