// Code generated by go-swagger; DO NOT EDIT.

package integration

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

// NewPutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams creates a new PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams() *PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams {
	return &PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParamsWithTimeout creates a new PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams object
// with the ability to set a timeout on a request.
func NewPutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParamsWithTimeout(timeout time.Duration) *PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams {
	return &PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams{
		timeout: timeout,
	}
}

// NewPutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParamsWithContext creates a new PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams object
// with the ability to set a context for a request.
func NewPutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParamsWithContext(ctx context.Context) *PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams {
	return &PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams{
		Context: ctx,
	}
}

// NewPutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParamsWithHTTPClient creates a new PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParamsWithHTTPClient(client *http.Client) *PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams {
	return &PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams{
		HTTPClient: client,
	}
}

/*
PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams contains all the parameters to send to the API endpoint

	for the put project project ID integrations integration ID values extractvalue ID operation.

	Typically these are written to a http.Request.
*/
type PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams struct {

	// IntegrationExtractValue.
	IntegrationExtractValue *models.IntegrationExtractValueRequest

	/* ExtractvalueID.

	   extractValue ID
	*/
	ExtractvalueID int64

	/* IntegrationID.

	   integration ID
	*/
	IntegrationID int64

	/* ProjectID.

	   Project ID
	*/
	ProjectID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put project project ID integrations integration ID values extractvalue ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams) WithDefaults() *PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put project project ID integrations integration ID values extractvalue ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put project project ID integrations integration ID values extractvalue ID params
func (o *PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams) WithTimeout(timeout time.Duration) *PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put project project ID integrations integration ID values extractvalue ID params
func (o *PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put project project ID integrations integration ID values extractvalue ID params
func (o *PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams) WithContext(ctx context.Context) *PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put project project ID integrations integration ID values extractvalue ID params
func (o *PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put project project ID integrations integration ID values extractvalue ID params
func (o *PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams) WithHTTPClient(client *http.Client) *PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put project project ID integrations integration ID values extractvalue ID params
func (o *PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIntegrationExtractValue adds the integrationExtractValue to the put project project ID integrations integration ID values extractvalue ID params
func (o *PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams) WithIntegrationExtractValue(integrationExtractValue *models.IntegrationExtractValueRequest) *PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams {
	o.SetIntegrationExtractValue(integrationExtractValue)
	return o
}

// SetIntegrationExtractValue adds the integrationExtractValue to the put project project ID integrations integration ID values extractvalue ID params
func (o *PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams) SetIntegrationExtractValue(integrationExtractValue *models.IntegrationExtractValueRequest) {
	o.IntegrationExtractValue = integrationExtractValue
}

// WithExtractvalueID adds the extractvalueID to the put project project ID integrations integration ID values extractvalue ID params
func (o *PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams) WithExtractvalueID(extractvalueID int64) *PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams {
	o.SetExtractvalueID(extractvalueID)
	return o
}

// SetExtractvalueID adds the extractvalueId to the put project project ID integrations integration ID values extractvalue ID params
func (o *PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams) SetExtractvalueID(extractvalueID int64) {
	o.ExtractvalueID = extractvalueID
}

// WithIntegrationID adds the integrationID to the put project project ID integrations integration ID values extractvalue ID params
func (o *PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams) WithIntegrationID(integrationID int64) *PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams {
	o.SetIntegrationID(integrationID)
	return o
}

// SetIntegrationID adds the integrationId to the put project project ID integrations integration ID values extractvalue ID params
func (o *PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams) SetIntegrationID(integrationID int64) {
	o.IntegrationID = integrationID
}

// WithProjectID adds the projectID to the put project project ID integrations integration ID values extractvalue ID params
func (o *PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams) WithProjectID(projectID int64) *PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the put project project ID integrations integration ID values extractvalue ID params
func (o *PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *PutProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.IntegrationExtractValue != nil {
		if err := r.SetBodyParam(o.IntegrationExtractValue); err != nil {
			return err
		}
	}

	// path param extractvalue_id
	if err := r.SetPathParam("extractvalue_id", swag.FormatInt64(o.ExtractvalueID)); err != nil {
		return err
	}

	// path param integration_id
	if err := r.SetPathParam("integration_id", swag.FormatInt64(o.IntegrationID)); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", swag.FormatInt64(o.ProjectID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
