// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutProjectProjectIDIntegrationsIntegrationIDReader is a Reader for the PutProjectProjectIDIntegrationsIntegrationID structure.
type PutProjectProjectIDIntegrationsIntegrationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutProjectProjectIDIntegrationsIntegrationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutProjectProjectIDIntegrationsIntegrationIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /project/{project_id}/integrations/{integration_id}] PutProjectProjectIDIntegrationsIntegrationID", response, response.Code())
	}
}

// NewPutProjectProjectIDIntegrationsIntegrationIDNoContent creates a PutProjectProjectIDIntegrationsIntegrationIDNoContent with default headers values
func NewPutProjectProjectIDIntegrationsIntegrationIDNoContent() *PutProjectProjectIDIntegrationsIntegrationIDNoContent {
	return &PutProjectProjectIDIntegrationsIntegrationIDNoContent{}
}

/*
PutProjectProjectIDIntegrationsIntegrationIDNoContent describes a response with status code 204, with default header values.

Integration updated
*/
type PutProjectProjectIDIntegrationsIntegrationIDNoContent struct {
}

// IsSuccess returns true when this put project project Id integrations integration Id no content response has a 2xx status code
func (o *PutProjectProjectIDIntegrationsIntegrationIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put project project Id integrations integration Id no content response has a 3xx status code
func (o *PutProjectProjectIDIntegrationsIntegrationIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put project project Id integrations integration Id no content response has a 4xx status code
func (o *PutProjectProjectIDIntegrationsIntegrationIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put project project Id integrations integration Id no content response has a 5xx status code
func (o *PutProjectProjectIDIntegrationsIntegrationIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put project project Id integrations integration Id no content response a status code equal to that given
func (o *PutProjectProjectIDIntegrationsIntegrationIDNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the put project project Id integrations integration Id no content response
func (o *PutProjectProjectIDIntegrationsIntegrationIDNoContent) Code() int {
	return 204
}

func (o *PutProjectProjectIDIntegrationsIntegrationIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /project/{project_id}/integrations/{integration_id}][%d] putProjectProjectIdIntegrationsIntegrationIdNoContent", 204)
}

func (o *PutProjectProjectIDIntegrationsIntegrationIDNoContent) String() string {
	return fmt.Sprintf("[PUT /project/{project_id}/integrations/{integration_id}][%d] putProjectProjectIdIntegrationsIntegrationIdNoContent", 204)
}

func (o *PutProjectProjectIDIntegrationsIntegrationIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
