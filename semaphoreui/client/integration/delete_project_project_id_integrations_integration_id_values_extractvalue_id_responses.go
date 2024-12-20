// Code generated by go-swagger; DO NOT EDIT.

package integration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDReader is a Reader for the DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueID structure.
type DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /project/{project_id}/integrations/{integration_id}/values/{extractvalue_id}] DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueID", response, response.Code())
	}
}

// NewDeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDNoContent creates a DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDNoContent with default headers values
func NewDeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDNoContent() *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDNoContent {
	return &DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDNoContent{}
}

/*
DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDNoContent describes a response with status code 204, with default header values.

integration extract value removed
*/
type DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDNoContent struct {
}

// IsSuccess returns true when this delete project project Id integrations integration Id values extractvalue Id no content response has a 2xx status code
func (o *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete project project Id integrations integration Id values extractvalue Id no content response has a 3xx status code
func (o *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project project Id integrations integration Id values extractvalue Id no content response has a 4xx status code
func (o *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete project project Id integrations integration Id values extractvalue Id no content response has a 5xx status code
func (o *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project project Id integrations integration Id values extractvalue Id no content response a status code equal to that given
func (o *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete project project Id integrations integration Id values extractvalue Id no content response
func (o *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDNoContent) Code() int {
	return 204
}

func (o *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /project/{project_id}/integrations/{integration_id}/values/{extractvalue_id}][%d] deleteProjectProjectIdIntegrationsIntegrationIdValuesExtractvalueIdNoContent", 204)
}

func (o *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDNoContent) String() string {
	return fmt.Sprintf("[DELETE /project/{project_id}/integrations/{integration_id}/values/{extractvalue_id}][%d] deleteProjectProjectIdIntegrationsIntegrationIdValuesExtractvalueIdNoContent", 204)
}

func (o *DeleteProjectProjectIDIntegrationsIntegrationIDValuesExtractvalueIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
