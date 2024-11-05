// Code generated by go-swagger; DO NOT EDIT.

package schedule

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

// NewDeleteProjectProjectIDSchedulesScheduleIDParams creates a new DeleteProjectProjectIDSchedulesScheduleIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteProjectProjectIDSchedulesScheduleIDParams() *DeleteProjectProjectIDSchedulesScheduleIDParams {
	return &DeleteProjectProjectIDSchedulesScheduleIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteProjectProjectIDSchedulesScheduleIDParamsWithTimeout creates a new DeleteProjectProjectIDSchedulesScheduleIDParams object
// with the ability to set a timeout on a request.
func NewDeleteProjectProjectIDSchedulesScheduleIDParamsWithTimeout(timeout time.Duration) *DeleteProjectProjectIDSchedulesScheduleIDParams {
	return &DeleteProjectProjectIDSchedulesScheduleIDParams{
		timeout: timeout,
	}
}

// NewDeleteProjectProjectIDSchedulesScheduleIDParamsWithContext creates a new DeleteProjectProjectIDSchedulesScheduleIDParams object
// with the ability to set a context for a request.
func NewDeleteProjectProjectIDSchedulesScheduleIDParamsWithContext(ctx context.Context) *DeleteProjectProjectIDSchedulesScheduleIDParams {
	return &DeleteProjectProjectIDSchedulesScheduleIDParams{
		Context: ctx,
	}
}

// NewDeleteProjectProjectIDSchedulesScheduleIDParamsWithHTTPClient creates a new DeleteProjectProjectIDSchedulesScheduleIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteProjectProjectIDSchedulesScheduleIDParamsWithHTTPClient(client *http.Client) *DeleteProjectProjectIDSchedulesScheduleIDParams {
	return &DeleteProjectProjectIDSchedulesScheduleIDParams{
		HTTPClient: client,
	}
}

/*
DeleteProjectProjectIDSchedulesScheduleIDParams contains all the parameters to send to the API endpoint

	for the delete project project ID schedules schedule ID operation.

	Typically these are written to a http.Request.
*/
type DeleteProjectProjectIDSchedulesScheduleIDParams struct {

	/* ProjectID.

	   Project ID
	*/
	ProjectID int64

	/* ScheduleID.

	   schedule ID
	*/
	ScheduleID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete project project ID schedules schedule ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteProjectProjectIDSchedulesScheduleIDParams) WithDefaults() *DeleteProjectProjectIDSchedulesScheduleIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete project project ID schedules schedule ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteProjectProjectIDSchedulesScheduleIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete project project ID schedules schedule ID params
func (o *DeleteProjectProjectIDSchedulesScheduleIDParams) WithTimeout(timeout time.Duration) *DeleteProjectProjectIDSchedulesScheduleIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete project project ID schedules schedule ID params
func (o *DeleteProjectProjectIDSchedulesScheduleIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete project project ID schedules schedule ID params
func (o *DeleteProjectProjectIDSchedulesScheduleIDParams) WithContext(ctx context.Context) *DeleteProjectProjectIDSchedulesScheduleIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete project project ID schedules schedule ID params
func (o *DeleteProjectProjectIDSchedulesScheduleIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete project project ID schedules schedule ID params
func (o *DeleteProjectProjectIDSchedulesScheduleIDParams) WithHTTPClient(client *http.Client) *DeleteProjectProjectIDSchedulesScheduleIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete project project ID schedules schedule ID params
func (o *DeleteProjectProjectIDSchedulesScheduleIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the delete project project ID schedules schedule ID params
func (o *DeleteProjectProjectIDSchedulesScheduleIDParams) WithProjectID(projectID int64) *DeleteProjectProjectIDSchedulesScheduleIDParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete project project ID schedules schedule ID params
func (o *DeleteProjectProjectIDSchedulesScheduleIDParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WithScheduleID adds the scheduleID to the delete project project ID schedules schedule ID params
func (o *DeleteProjectProjectIDSchedulesScheduleIDParams) WithScheduleID(scheduleID int64) *DeleteProjectProjectIDSchedulesScheduleIDParams {
	o.SetScheduleID(scheduleID)
	return o
}

// SetScheduleID adds the scheduleId to the delete project project ID schedules schedule ID params
func (o *DeleteProjectProjectIDSchedulesScheduleIDParams) SetScheduleID(scheduleID int64) {
	o.ScheduleID = scheduleID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteProjectProjectIDSchedulesScheduleIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param project_id
	if err := r.SetPathParam("project_id", swag.FormatInt64(o.ProjectID)); err != nil {
		return err
	}

	// path param schedule_id
	if err := r.SetPathParam("schedule_id", swag.FormatInt64(o.ScheduleID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
