// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutProjectProjectIDRepositoriesRepositoryIDReader is a Reader for the PutProjectProjectIDRepositoriesRepositoryID structure.
type PutProjectProjectIDRepositoriesRepositoryIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutProjectProjectIDRepositoriesRepositoryIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutProjectProjectIDRepositoriesRepositoryIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutProjectProjectIDRepositoriesRepositoryIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /project/{project_id}/repositories/{repository_id}] PutProjectProjectIDRepositoriesRepositoryID", response, response.Code())
	}
}

// NewPutProjectProjectIDRepositoriesRepositoryIDNoContent creates a PutProjectProjectIDRepositoriesRepositoryIDNoContent with default headers values
func NewPutProjectProjectIDRepositoriesRepositoryIDNoContent() *PutProjectProjectIDRepositoriesRepositoryIDNoContent {
	return &PutProjectProjectIDRepositoriesRepositoryIDNoContent{}
}

/*
PutProjectProjectIDRepositoriesRepositoryIDNoContent describes a response with status code 204, with default header values.

Repository updated
*/
type PutProjectProjectIDRepositoriesRepositoryIDNoContent struct {
}

// IsSuccess returns true when this put project project Id repositories repository Id no content response has a 2xx status code
func (o *PutProjectProjectIDRepositoriesRepositoryIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put project project Id repositories repository Id no content response has a 3xx status code
func (o *PutProjectProjectIDRepositoriesRepositoryIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put project project Id repositories repository Id no content response has a 4xx status code
func (o *PutProjectProjectIDRepositoriesRepositoryIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put project project Id repositories repository Id no content response has a 5xx status code
func (o *PutProjectProjectIDRepositoriesRepositoryIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put project project Id repositories repository Id no content response a status code equal to that given
func (o *PutProjectProjectIDRepositoriesRepositoryIDNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the put project project Id repositories repository Id no content response
func (o *PutProjectProjectIDRepositoriesRepositoryIDNoContent) Code() int {
	return 204
}

func (o *PutProjectProjectIDRepositoriesRepositoryIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /project/{project_id}/repositories/{repository_id}][%d] putProjectProjectIdRepositoriesRepositoryIdNoContent", 204)
}

func (o *PutProjectProjectIDRepositoriesRepositoryIDNoContent) String() string {
	return fmt.Sprintf("[PUT /project/{project_id}/repositories/{repository_id}][%d] putProjectProjectIdRepositoriesRepositoryIdNoContent", 204)
}

func (o *PutProjectProjectIDRepositoriesRepositoryIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutProjectProjectIDRepositoriesRepositoryIDBadRequest creates a PutProjectProjectIDRepositoriesRepositoryIDBadRequest with default headers values
func NewPutProjectProjectIDRepositoriesRepositoryIDBadRequest() *PutProjectProjectIDRepositoriesRepositoryIDBadRequest {
	return &PutProjectProjectIDRepositoriesRepositoryIDBadRequest{}
}

/*
PutProjectProjectIDRepositoriesRepositoryIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PutProjectProjectIDRepositoriesRepositoryIDBadRequest struct {
}

// IsSuccess returns true when this put project project Id repositories repository Id bad request response has a 2xx status code
func (o *PutProjectProjectIDRepositoriesRepositoryIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put project project Id repositories repository Id bad request response has a 3xx status code
func (o *PutProjectProjectIDRepositoriesRepositoryIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put project project Id repositories repository Id bad request response has a 4xx status code
func (o *PutProjectProjectIDRepositoriesRepositoryIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put project project Id repositories repository Id bad request response has a 5xx status code
func (o *PutProjectProjectIDRepositoriesRepositoryIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put project project Id repositories repository Id bad request response a status code equal to that given
func (o *PutProjectProjectIDRepositoriesRepositoryIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put project project Id repositories repository Id bad request response
func (o *PutProjectProjectIDRepositoriesRepositoryIDBadRequest) Code() int {
	return 400
}

func (o *PutProjectProjectIDRepositoriesRepositoryIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /project/{project_id}/repositories/{repository_id}][%d] putProjectProjectIdRepositoriesRepositoryIdBadRequest", 400)
}

func (o *PutProjectProjectIDRepositoriesRepositoryIDBadRequest) String() string {
	return fmt.Sprintf("[PUT /project/{project_id}/repositories/{repository_id}][%d] putProjectProjectIdRepositoriesRepositoryIdBadRequest", 400)
}

func (o *PutProjectProjectIDRepositoriesRepositoryIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
