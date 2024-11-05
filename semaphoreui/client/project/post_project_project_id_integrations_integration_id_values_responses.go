// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostProjectProjectIDIntegrationsIntegrationIDValuesReader is a Reader for the PostProjectProjectIDIntegrationsIntegrationIDValues structure.
type PostProjectProjectIDIntegrationsIntegrationIDValuesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostProjectProjectIDIntegrationsIntegrationIDValuesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostProjectProjectIDIntegrationsIntegrationIDValuesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostProjectProjectIDIntegrationsIntegrationIDValuesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /project/{project_id}/integrations/{integration_id}/values] PostProjectProjectIDIntegrationsIntegrationIDValues", response, response.Code())
	}
}

// NewPostProjectProjectIDIntegrationsIntegrationIDValuesCreated creates a PostProjectProjectIDIntegrationsIntegrationIDValuesCreated with default headers values
func NewPostProjectProjectIDIntegrationsIntegrationIDValuesCreated() *PostProjectProjectIDIntegrationsIntegrationIDValuesCreated {
	return &PostProjectProjectIDIntegrationsIntegrationIDValuesCreated{}
}

/*
PostProjectProjectIDIntegrationsIntegrationIDValuesCreated describes a response with status code 201, with default header values.

Integration Extract Value Created
*/
type PostProjectProjectIDIntegrationsIntegrationIDValuesCreated struct {
}

// IsSuccess returns true when this post project project Id integrations integration Id values created response has a 2xx status code
func (o *PostProjectProjectIDIntegrationsIntegrationIDValuesCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post project project Id integrations integration Id values created response has a 3xx status code
func (o *PostProjectProjectIDIntegrationsIntegrationIDValuesCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post project project Id integrations integration Id values created response has a 4xx status code
func (o *PostProjectProjectIDIntegrationsIntegrationIDValuesCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post project project Id integrations integration Id values created response has a 5xx status code
func (o *PostProjectProjectIDIntegrationsIntegrationIDValuesCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post project project Id integrations integration Id values created response a status code equal to that given
func (o *PostProjectProjectIDIntegrationsIntegrationIDValuesCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post project project Id integrations integration Id values created response
func (o *PostProjectProjectIDIntegrationsIntegrationIDValuesCreated) Code() int {
	return 201
}

func (o *PostProjectProjectIDIntegrationsIntegrationIDValuesCreated) Error() string {
	return fmt.Sprintf("[POST /project/{project_id}/integrations/{integration_id}/values][%d] postProjectProjectIdIntegrationsIntegrationIdValuesCreated", 201)
}

func (o *PostProjectProjectIDIntegrationsIntegrationIDValuesCreated) String() string {
	return fmt.Sprintf("[POST /project/{project_id}/integrations/{integration_id}/values][%d] postProjectProjectIdIntegrationsIntegrationIdValuesCreated", 201)
}

func (o *PostProjectProjectIDIntegrationsIntegrationIDValuesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostProjectProjectIDIntegrationsIntegrationIDValuesBadRequest creates a PostProjectProjectIDIntegrationsIntegrationIDValuesBadRequest with default headers values
func NewPostProjectProjectIDIntegrationsIntegrationIDValuesBadRequest() *PostProjectProjectIDIntegrationsIntegrationIDValuesBadRequest {
	return &PostProjectProjectIDIntegrationsIntegrationIDValuesBadRequest{}
}

/*
PostProjectProjectIDIntegrationsIntegrationIDValuesBadRequest describes a response with status code 400, with default header values.

Bad Integration Extract Value params
*/
type PostProjectProjectIDIntegrationsIntegrationIDValuesBadRequest struct {
}

// IsSuccess returns true when this post project project Id integrations integration Id values bad request response has a 2xx status code
func (o *PostProjectProjectIDIntegrationsIntegrationIDValuesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post project project Id integrations integration Id values bad request response has a 3xx status code
func (o *PostProjectProjectIDIntegrationsIntegrationIDValuesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post project project Id integrations integration Id values bad request response has a 4xx status code
func (o *PostProjectProjectIDIntegrationsIntegrationIDValuesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post project project Id integrations integration Id values bad request response has a 5xx status code
func (o *PostProjectProjectIDIntegrationsIntegrationIDValuesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post project project Id integrations integration Id values bad request response a status code equal to that given
func (o *PostProjectProjectIDIntegrationsIntegrationIDValuesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post project project Id integrations integration Id values bad request response
func (o *PostProjectProjectIDIntegrationsIntegrationIDValuesBadRequest) Code() int {
	return 400
}

func (o *PostProjectProjectIDIntegrationsIntegrationIDValuesBadRequest) Error() string {
	return fmt.Sprintf("[POST /project/{project_id}/integrations/{integration_id}/values][%d] postProjectProjectIdIntegrationsIntegrationIdValuesBadRequest", 400)
}

func (o *PostProjectProjectIDIntegrationsIntegrationIDValuesBadRequest) String() string {
	return fmt.Sprintf("[POST /project/{project_id}/integrations/{integration_id}/values][%d] postProjectProjectIdIntegrationsIntegrationIdValuesBadRequest", 400)
}

func (o *PostProjectProjectIDIntegrationsIntegrationIDValuesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
