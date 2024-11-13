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

	"terraform-provider-semaphoreui/semaphoreui/models"
)

// NewPutProjectProjectIDSchedulesScheduleIDParams creates a new PutProjectProjectIDSchedulesScheduleIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutProjectProjectIDSchedulesScheduleIDParams() *PutProjectProjectIDSchedulesScheduleIDParams {
	return &PutProjectProjectIDSchedulesScheduleIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutProjectProjectIDSchedulesScheduleIDParamsWithTimeout creates a new PutProjectProjectIDSchedulesScheduleIDParams object
// with the ability to set a timeout on a request.
func NewPutProjectProjectIDSchedulesScheduleIDParamsWithTimeout(timeout time.Duration) *PutProjectProjectIDSchedulesScheduleIDParams {
	return &PutProjectProjectIDSchedulesScheduleIDParams{
		timeout: timeout,
	}
}

// NewPutProjectProjectIDSchedulesScheduleIDParamsWithContext creates a new PutProjectProjectIDSchedulesScheduleIDParams object
// with the ability to set a context for a request.
func NewPutProjectProjectIDSchedulesScheduleIDParamsWithContext(ctx context.Context) *PutProjectProjectIDSchedulesScheduleIDParams {
	return &PutProjectProjectIDSchedulesScheduleIDParams{
		Context: ctx,
	}
}

// NewPutProjectProjectIDSchedulesScheduleIDParamsWithHTTPClient creates a new PutProjectProjectIDSchedulesScheduleIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutProjectProjectIDSchedulesScheduleIDParamsWithHTTPClient(client *http.Client) *PutProjectProjectIDSchedulesScheduleIDParams {
	return &PutProjectProjectIDSchedulesScheduleIDParams{
		HTTPClient: client,
	}
}

/*
PutProjectProjectIDSchedulesScheduleIDParams contains all the parameters to send to the API endpoint

	for the put project project ID schedules schedule ID operation.

	Typically these are written to a http.Request.
*/
type PutProjectProjectIDSchedulesScheduleIDParams struct {

	/* ProjectID.

	   Project ID
	*/
	ProjectID int64

	// Schedule.
	Schedule *models.ScheduleRequest

	/* ScheduleID.

	   schedule ID
	*/
	ScheduleID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put project project ID schedules schedule ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutProjectProjectIDSchedulesScheduleIDParams) WithDefaults() *PutProjectProjectIDSchedulesScheduleIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put project project ID schedules schedule ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutProjectProjectIDSchedulesScheduleIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put project project ID schedules schedule ID params
func (o *PutProjectProjectIDSchedulesScheduleIDParams) WithTimeout(timeout time.Duration) *PutProjectProjectIDSchedulesScheduleIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put project project ID schedules schedule ID params
func (o *PutProjectProjectIDSchedulesScheduleIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put project project ID schedules schedule ID params
func (o *PutProjectProjectIDSchedulesScheduleIDParams) WithContext(ctx context.Context) *PutProjectProjectIDSchedulesScheduleIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put project project ID schedules schedule ID params
func (o *PutProjectProjectIDSchedulesScheduleIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put project project ID schedules schedule ID params
func (o *PutProjectProjectIDSchedulesScheduleIDParams) WithHTTPClient(client *http.Client) *PutProjectProjectIDSchedulesScheduleIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put project project ID schedules schedule ID params
func (o *PutProjectProjectIDSchedulesScheduleIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the put project project ID schedules schedule ID params
func (o *PutProjectProjectIDSchedulesScheduleIDParams) WithProjectID(projectID int64) *PutProjectProjectIDSchedulesScheduleIDParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the put project project ID schedules schedule ID params
func (o *PutProjectProjectIDSchedulesScheduleIDParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WithSchedule adds the schedule to the put project project ID schedules schedule ID params
func (o *PutProjectProjectIDSchedulesScheduleIDParams) WithSchedule(schedule *models.ScheduleRequest) *PutProjectProjectIDSchedulesScheduleIDParams {
	o.SetSchedule(schedule)
	return o
}

// SetSchedule adds the schedule to the put project project ID schedules schedule ID params
func (o *PutProjectProjectIDSchedulesScheduleIDParams) SetSchedule(schedule *models.ScheduleRequest) {
	o.Schedule = schedule
}

// WithScheduleID adds the scheduleID to the put project project ID schedules schedule ID params
func (o *PutProjectProjectIDSchedulesScheduleIDParams) WithScheduleID(scheduleID int64) *PutProjectProjectIDSchedulesScheduleIDParams {
	o.SetScheduleID(scheduleID)
	return o
}

// SetScheduleID adds the scheduleId to the put project project ID schedules schedule ID params
func (o *PutProjectProjectIDSchedulesScheduleIDParams) SetScheduleID(scheduleID int64) {
	o.ScheduleID = scheduleID
}

// WriteToRequest writes these params to a swagger request
func (o *PutProjectProjectIDSchedulesScheduleIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param project_id
	if err := r.SetPathParam("project_id", swag.FormatInt64(o.ProjectID)); err != nil {
		return err
	}
	if o.Schedule != nil {
		if err := r.SetBodyParam(o.Schedule); err != nil {
			return err
		}
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