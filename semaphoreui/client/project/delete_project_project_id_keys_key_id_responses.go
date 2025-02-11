// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteProjectProjectIDKeysKeyIDReader is a Reader for the DeleteProjectProjectIDKeysKeyID structure.
type DeleteProjectProjectIDKeysKeyIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectProjectIDKeysKeyIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteProjectProjectIDKeysKeyIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /project/{project_id}/keys/{key_id}] DeleteProjectProjectIDKeysKeyID", response, response.Code())
	}
}

// NewDeleteProjectProjectIDKeysKeyIDNoContent creates a DeleteProjectProjectIDKeysKeyIDNoContent with default headers values
func NewDeleteProjectProjectIDKeysKeyIDNoContent() *DeleteProjectProjectIDKeysKeyIDNoContent {
	return &DeleteProjectProjectIDKeysKeyIDNoContent{}
}

/*
DeleteProjectProjectIDKeysKeyIDNoContent describes a response with status code 204, with default header values.

access key removed
*/
type DeleteProjectProjectIDKeysKeyIDNoContent struct {
}

// IsSuccess returns true when this delete project project Id keys key Id no content response has a 2xx status code
func (o *DeleteProjectProjectIDKeysKeyIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete project project Id keys key Id no content response has a 3xx status code
func (o *DeleteProjectProjectIDKeysKeyIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project project Id keys key Id no content response has a 4xx status code
func (o *DeleteProjectProjectIDKeysKeyIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete project project Id keys key Id no content response has a 5xx status code
func (o *DeleteProjectProjectIDKeysKeyIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project project Id keys key Id no content response a status code equal to that given
func (o *DeleteProjectProjectIDKeysKeyIDNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete project project Id keys key Id no content response
func (o *DeleteProjectProjectIDKeysKeyIDNoContent) Code() int {
	return 204
}

func (o *DeleteProjectProjectIDKeysKeyIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /project/{project_id}/keys/{key_id}][%d] deleteProjectProjectIdKeysKeyIdNoContent", 204)
}

func (o *DeleteProjectProjectIDKeysKeyIDNoContent) String() string {
	return fmt.Sprintf("[DELETE /project/{project_id}/keys/{key_id}][%d] deleteProjectProjectIdKeysKeyIdNoContent", 204)
}

func (o *DeleteProjectProjectIDKeysKeyIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
