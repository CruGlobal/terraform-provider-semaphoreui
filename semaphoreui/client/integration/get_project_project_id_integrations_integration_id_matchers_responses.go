// Code generated by go-swagger; DO NOT EDIT.

package integration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"terraform-provider-semaphoreui/semaphoreui/models"
)

// GetProjectProjectIDIntegrationsIntegrationIDMatchersReader is a Reader for the GetProjectProjectIDIntegrationsIntegrationIDMatchers structure.
type GetProjectProjectIDIntegrationsIntegrationIDMatchersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectProjectIDIntegrationsIntegrationIDMatchersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectProjectIDIntegrationsIntegrationIDMatchersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /project/{project_id}/integrations/{integration_id}/matchers] GetProjectProjectIDIntegrationsIntegrationIDMatchers", response, response.Code())
	}
}

// NewGetProjectProjectIDIntegrationsIntegrationIDMatchersOK creates a GetProjectProjectIDIntegrationsIntegrationIDMatchersOK with default headers values
func NewGetProjectProjectIDIntegrationsIntegrationIDMatchersOK() *GetProjectProjectIDIntegrationsIntegrationIDMatchersOK {
	return &GetProjectProjectIDIntegrationsIntegrationIDMatchersOK{}
}

/*
GetProjectProjectIDIntegrationsIntegrationIDMatchersOK describes a response with status code 200, with default header values.

Integration Matcher
*/
type GetProjectProjectIDIntegrationsIntegrationIDMatchersOK struct {
	Payload []*models.IntegrationMatcher
}

// IsSuccess returns true when this get project project Id integrations integration Id matchers o k response has a 2xx status code
func (o *GetProjectProjectIDIntegrationsIntegrationIDMatchersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get project project Id integrations integration Id matchers o k response has a 3xx status code
func (o *GetProjectProjectIDIntegrationsIntegrationIDMatchersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project project Id integrations integration Id matchers o k response has a 4xx status code
func (o *GetProjectProjectIDIntegrationsIntegrationIDMatchersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get project project Id integrations integration Id matchers o k response has a 5xx status code
func (o *GetProjectProjectIDIntegrationsIntegrationIDMatchersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get project project Id integrations integration Id matchers o k response a status code equal to that given
func (o *GetProjectProjectIDIntegrationsIntegrationIDMatchersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get project project Id integrations integration Id matchers o k response
func (o *GetProjectProjectIDIntegrationsIntegrationIDMatchersOK) Code() int {
	return 200
}

func (o *GetProjectProjectIDIntegrationsIntegrationIDMatchersOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /project/{project_id}/integrations/{integration_id}/matchers][%d] getProjectProjectIdIntegrationsIntegrationIdMatchersOK %s", 200, payload)
}

func (o *GetProjectProjectIDIntegrationsIntegrationIDMatchersOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /project/{project_id}/integrations/{integration_id}/matchers][%d] getProjectProjectIdIntegrationsIntegrationIdMatchersOK %s", 200, payload)
}

func (o *GetProjectProjectIDIntegrationsIntegrationIDMatchersOK) GetPayload() []*models.IntegrationMatcher {
	return o.Payload
}

func (o *GetProjectProjectIDIntegrationsIntegrationIDMatchersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}