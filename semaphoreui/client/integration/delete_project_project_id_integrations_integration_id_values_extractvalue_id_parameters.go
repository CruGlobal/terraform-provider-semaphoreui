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
)

// NewDeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams creates a new DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams() *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams {
	return &DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParamsWithTimeout creates a new DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams object
// with the ability to set a timeout on a request.
func NewDeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParamsWithTimeout(timeout time.Duration) *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams {
	return &DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams{
		timeout: timeout,
	}
}

// NewDeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParamsWithContext creates a new DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams object
// with the ability to set a context for a request.
func NewDeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParamsWithContext(ctx context.Context) *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams {
	return &DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams{
		Context: ctx,
	}
}

// NewDeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParamsWithHTTPClient creates a new DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParamsWithHTTPClient(client *http.Client) *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams {
	return &DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams{
		HTTPClient: client,
	}
}

/*
DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams contains all the parameters to send to the API endpoint

	for the delete project project ID integrations integration ID values extractvalue ID operation.

	Typically these are written to a http.Request.
*/
type DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams struct {

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

// WithDefaults hydrates default values in the delete project project ID integrations integration ID values extractvalue ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams) WithDefaults() *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete project project ID integrations integration ID values extractvalue ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete project project ID integrations integration ID values extractvalue ID params
func (o *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams) WithTimeout(timeout time.Duration) *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete project project ID integrations integration ID values extractvalue ID params
func (o *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete project project ID integrations integration ID values extractvalue ID params
func (o *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams) WithContext(ctx context.Context) *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete project project ID integrations integration ID values extractvalue ID params
func (o *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete project project ID integrations integration ID values extractvalue ID params
func (o *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams) WithHTTPClient(client *http.Client) *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete project project ID integrations integration ID values extractvalue ID params
func (o *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExtractvalueID adds the extractvalueID to the delete project project ID integrations integration ID values extractvalue ID params
func (o *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams) WithExtractvalueID(extractvalueID int64) *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams {
	o.SetExtractvalueID(extractvalueID)
	return o
}

// SetExtractvalueID adds the extractvalueId to the delete project project ID integrations integration ID values extractvalue ID params
func (o *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams) SetExtractvalueID(extractvalueID int64) {
	o.ExtractvalueID = extractvalueID
}

// WithIntegrationID adds the integrationID to the delete project project ID integrations integration ID values extractvalue ID params
func (o *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams) WithIntegrationID(integrationID int64) *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams {
	o.SetIntegrationID(integrationID)
	return o
}

// SetIntegrationID adds the integrationId to the delete project project ID integrations integration ID values extractvalue ID params
func (o *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams) SetIntegrationID(integrationID int64) {
	o.IntegrationID = integrationID
}

// WithProjectID adds the projectID to the delete project project ID integrations integration ID values extractvalue ID params
func (o *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams) WithProjectID(projectID int64) *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete project project ID integrations integration ID values extractvalue ID params
func (o *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
