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

// NewPostProjectProjectIDTasksTaskIDStopParams creates a new PostProjectProjectIDTasksTaskIDStopParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostProjectProjectIDTasksTaskIDStopParams() *PostProjectProjectIDTasksTaskIDStopParams {
	return &PostProjectProjectIDTasksTaskIDStopParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostProjectProjectIDTasksTaskIDStopParamsWithTimeout creates a new PostProjectProjectIDTasksTaskIDStopParams object
// with the ability to set a timeout on a request.
func NewPostProjectProjectIDTasksTaskIDStopParamsWithTimeout(timeout time.Duration) *PostProjectProjectIDTasksTaskIDStopParams {
	return &PostProjectProjectIDTasksTaskIDStopParams{
		timeout: timeout,
	}
}

// NewPostProjectProjectIDTasksTaskIDStopParamsWithContext creates a new PostProjectProjectIDTasksTaskIDStopParams object
// with the ability to set a context for a request.
func NewPostProjectProjectIDTasksTaskIDStopParamsWithContext(ctx context.Context) *PostProjectProjectIDTasksTaskIDStopParams {
	return &PostProjectProjectIDTasksTaskIDStopParams{
		Context: ctx,
	}
}

// NewPostProjectProjectIDTasksTaskIDStopParamsWithHTTPClient creates a new PostProjectProjectIDTasksTaskIDStopParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostProjectProjectIDTasksTaskIDStopParamsWithHTTPClient(client *http.Client) *PostProjectProjectIDTasksTaskIDStopParams {
	return &PostProjectProjectIDTasksTaskIDStopParams{
		HTTPClient: client,
	}
}

/*
PostProjectProjectIDTasksTaskIDStopParams contains all the parameters to send to the API endpoint

	for the post project project ID tasks task ID stop operation.

	Typically these are written to a http.Request.
*/
type PostProjectProjectIDTasksTaskIDStopParams struct {

	/* ProjectID.

	   Project ID
	*/
	ProjectID int64

	/* TaskID.

	   task ID
	*/
	TaskID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post project project ID tasks task ID stop params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostProjectProjectIDTasksTaskIDStopParams) WithDefaults() *PostProjectProjectIDTasksTaskIDStopParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post project project ID tasks task ID stop params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostProjectProjectIDTasksTaskIDStopParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post project project ID tasks task ID stop params
func (o *PostProjectProjectIDTasksTaskIDStopParams) WithTimeout(timeout time.Duration) *PostProjectProjectIDTasksTaskIDStopParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post project project ID tasks task ID stop params
func (o *PostProjectProjectIDTasksTaskIDStopParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post project project ID tasks task ID stop params
func (o *PostProjectProjectIDTasksTaskIDStopParams) WithContext(ctx context.Context) *PostProjectProjectIDTasksTaskIDStopParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post project project ID tasks task ID stop params
func (o *PostProjectProjectIDTasksTaskIDStopParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post project project ID tasks task ID stop params
func (o *PostProjectProjectIDTasksTaskIDStopParams) WithHTTPClient(client *http.Client) *PostProjectProjectIDTasksTaskIDStopParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post project project ID tasks task ID stop params
func (o *PostProjectProjectIDTasksTaskIDStopParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the post project project ID tasks task ID stop params
func (o *PostProjectProjectIDTasksTaskIDStopParams) WithProjectID(projectID int64) *PostProjectProjectIDTasksTaskIDStopParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the post project project ID tasks task ID stop params
func (o *PostProjectProjectIDTasksTaskIDStopParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WithTaskID adds the taskID to the post project project ID tasks task ID stop params
func (o *PostProjectProjectIDTasksTaskIDStopParams) WithTaskID(taskID int64) *PostProjectProjectIDTasksTaskIDStopParams {
	o.SetTaskID(taskID)
	return o
}

// SetTaskID adds the taskId to the post project project ID tasks task ID stop params
func (o *PostProjectProjectIDTasksTaskIDStopParams) SetTaskID(taskID int64) {
	o.TaskID = taskID
}

// WriteToRequest writes these params to a swagger request
func (o *PostProjectProjectIDTasksTaskIDStopParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param project_id
	if err := r.SetPathParam("project_id", swag.FormatInt64(o.ProjectID)); err != nil {
		return err
	}

	// path param task_id
	if err := r.SetPathParam("task_id", swag.FormatInt64(o.TaskID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}