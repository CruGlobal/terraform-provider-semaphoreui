// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteProjectProjectIDRepositoriesRepositoryIDReader is a Reader for the DeleteProjectProjectIDRepositoriesRepositoryID structure.
type DeleteProjectProjectIDRepositoriesRepositoryIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectProjectIDRepositoriesRepositoryIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteProjectProjectIDRepositoriesRepositoryIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /project/{project_id}/repositories/{repository_id}] DeleteProjectProjectIDRepositoriesRepositoryID", response, response.Code())
	}
}

// NewDeleteProjectProjectIDRepositoriesRepositoryIDNoContent creates a DeleteProjectProjectIDRepositoriesRepositoryIDNoContent with default headers values
func NewDeleteProjectProjectIDRepositoriesRepositoryIDNoContent() *DeleteProjectProjectIDRepositoriesRepositoryIDNoContent {
	return &DeleteProjectProjectIDRepositoriesRepositoryIDNoContent{}
}

/*
DeleteProjectProjectIDRepositoriesRepositoryIDNoContent describes a response with status code 204, with default header values.

repository removed
*/
type DeleteProjectProjectIDRepositoriesRepositoryIDNoContent struct {
}

// IsSuccess returns true when this delete project project Id repositories repository Id no content response has a 2xx status code
func (o *DeleteProjectProjectIDRepositoriesRepositoryIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete project project Id repositories repository Id no content response has a 3xx status code
func (o *DeleteProjectProjectIDRepositoriesRepositoryIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project project Id repositories repository Id no content response has a 4xx status code
func (o *DeleteProjectProjectIDRepositoriesRepositoryIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete project project Id repositories repository Id no content response has a 5xx status code
func (o *DeleteProjectProjectIDRepositoriesRepositoryIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project project Id repositories repository Id no content response a status code equal to that given
func (o *DeleteProjectProjectIDRepositoriesRepositoryIDNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete project project Id repositories repository Id no content response
func (o *DeleteProjectProjectIDRepositoriesRepositoryIDNoContent) Code() int {
	return 204
}

func (o *DeleteProjectProjectIDRepositoriesRepositoryIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /project/{project_id}/repositories/{repository_id}][%d] deleteProjectProjectIdRepositoriesRepositoryIdNoContent", 204)
}

func (o *DeleteProjectProjectIDRepositoriesRepositoryIDNoContent) String() string {
	return fmt.Sprintf("[DELETE /project/{project_id}/repositories/{repository_id}][%d] deleteProjectProjectIdRepositoriesRepositoryIdNoContent", 204)
}

func (o *DeleteProjectProjectIDRepositoriesRepositoryIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
