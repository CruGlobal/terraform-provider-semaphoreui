// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostProjectProjectIDIntegrationsIntegrationIDMatchersReader is a Reader for the PostProjectProjectIDIntegrationsIntegrationIDMatchers structure.
type PostProjectProjectIDIntegrationsIntegrationIDMatchersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostProjectProjectIDIntegrationsIntegrationIDMatchersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostProjectProjectIDIntegrationsIntegrationIDMatchersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostProjectProjectIDIntegrationsIntegrationIDMatchersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /project/{project_id}/integrations/{integration_id}/matchers] PostProjectProjectIDIntegrationsIntegrationIDMatchers", response, response.Code())
	}
}

// NewPostProjectProjectIDIntegrationsIntegrationIDMatchersOK creates a PostProjectProjectIDIntegrationsIntegrationIDMatchersOK with default headers values
func NewPostProjectProjectIDIntegrationsIntegrationIDMatchersOK() *PostProjectProjectIDIntegrationsIntegrationIDMatchersOK {
	return &PostProjectProjectIDIntegrationsIntegrationIDMatchersOK{}
}

/*
PostProjectProjectIDIntegrationsIntegrationIDMatchersOK describes a response with status code 200, with default header values.

Integration Matcher Created
*/
type PostProjectProjectIDIntegrationsIntegrationIDMatchersOK struct {
}

// IsSuccess returns true when this post project project Id integrations integration Id matchers o k response has a 2xx status code
func (o *PostProjectProjectIDIntegrationsIntegrationIDMatchersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post project project Id integrations integration Id matchers o k response has a 3xx status code
func (o *PostProjectProjectIDIntegrationsIntegrationIDMatchersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post project project Id integrations integration Id matchers o k response has a 4xx status code
func (o *PostProjectProjectIDIntegrationsIntegrationIDMatchersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post project project Id integrations integration Id matchers o k response has a 5xx status code
func (o *PostProjectProjectIDIntegrationsIntegrationIDMatchersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post project project Id integrations integration Id matchers o k response a status code equal to that given
func (o *PostProjectProjectIDIntegrationsIntegrationIDMatchersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post project project Id integrations integration Id matchers o k response
func (o *PostProjectProjectIDIntegrationsIntegrationIDMatchersOK) Code() int {
	return 200
}

func (o *PostProjectProjectIDIntegrationsIntegrationIDMatchersOK) Error() string {
	return fmt.Sprintf("[POST /project/{project_id}/integrations/{integration_id}/matchers][%d] postProjectProjectIdIntegrationsIntegrationIdMatchersOK", 200)
}

func (o *PostProjectProjectIDIntegrationsIntegrationIDMatchersOK) String() string {
	return fmt.Sprintf("[POST /project/{project_id}/integrations/{integration_id}/matchers][%d] postProjectProjectIdIntegrationsIntegrationIdMatchersOK", 200)
}

func (o *PostProjectProjectIDIntegrationsIntegrationIDMatchersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostProjectProjectIDIntegrationsIntegrationIDMatchersBadRequest creates a PostProjectProjectIDIntegrationsIntegrationIDMatchersBadRequest with default headers values
func NewPostProjectProjectIDIntegrationsIntegrationIDMatchersBadRequest() *PostProjectProjectIDIntegrationsIntegrationIDMatchersBadRequest {
	return &PostProjectProjectIDIntegrationsIntegrationIDMatchersBadRequest{}
}

/*
PostProjectProjectIDIntegrationsIntegrationIDMatchersBadRequest describes a response with status code 400, with default header values.

Bad Integration Matcher params
*/
type PostProjectProjectIDIntegrationsIntegrationIDMatchersBadRequest struct {
}

// IsSuccess returns true when this post project project Id integrations integration Id matchers bad request response has a 2xx status code
func (o *PostProjectProjectIDIntegrationsIntegrationIDMatchersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post project project Id integrations integration Id matchers bad request response has a 3xx status code
func (o *PostProjectProjectIDIntegrationsIntegrationIDMatchersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post project project Id integrations integration Id matchers bad request response has a 4xx status code
func (o *PostProjectProjectIDIntegrationsIntegrationIDMatchersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post project project Id integrations integration Id matchers bad request response has a 5xx status code
func (o *PostProjectProjectIDIntegrationsIntegrationIDMatchersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post project project Id integrations integration Id matchers bad request response a status code equal to that given
func (o *PostProjectProjectIDIntegrationsIntegrationIDMatchersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post project project Id integrations integration Id matchers bad request response
func (o *PostProjectProjectIDIntegrationsIntegrationIDMatchersBadRequest) Code() int {
	return 400
}

func (o *PostProjectProjectIDIntegrationsIntegrationIDMatchersBadRequest) Error() string {
	return fmt.Sprintf("[POST /project/{project_id}/integrations/{integration_id}/matchers][%d] postProjectProjectIdIntegrationsIntegrationIdMatchersBadRequest", 400)
}

func (o *PostProjectProjectIDIntegrationsIntegrationIDMatchersBadRequest) String() string {
	return fmt.Sprintf("[POST /project/{project_id}/integrations/{integration_id}/matchers][%d] postProjectProjectIdIntegrationsIntegrationIdMatchersBadRequest", 400)
}

func (o *PostProjectProjectIDIntegrationsIntegrationIDMatchersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
