// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostProjectProjectIDTasksTaskIDStopReader is a Reader for the PostProjectProjectIDTasksTaskIDStop structure.
type PostProjectProjectIDTasksTaskIDStopReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostProjectProjectIDTasksTaskIDStopReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostProjectProjectIDTasksTaskIDStopNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /project/{project_id}/tasks/{task_id}/stop] PostProjectProjectIDTasksTaskIDStop", response, response.Code())
	}
}

// NewPostProjectProjectIDTasksTaskIDStopNoContent creates a PostProjectProjectIDTasksTaskIDStopNoContent with default headers values
func NewPostProjectProjectIDTasksTaskIDStopNoContent() *PostProjectProjectIDTasksTaskIDStopNoContent {
	return &PostProjectProjectIDTasksTaskIDStopNoContent{}
}

/*
PostProjectProjectIDTasksTaskIDStopNoContent describes a response with status code 204, with default header values.

Task queued
*/
type PostProjectProjectIDTasksTaskIDStopNoContent struct {
}

// IsSuccess returns true when this post project project Id tasks task Id stop no content response has a 2xx status code
func (o *PostProjectProjectIDTasksTaskIDStopNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post project project Id tasks task Id stop no content response has a 3xx status code
func (o *PostProjectProjectIDTasksTaskIDStopNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post project project Id tasks task Id stop no content response has a 4xx status code
func (o *PostProjectProjectIDTasksTaskIDStopNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this post project project Id tasks task Id stop no content response has a 5xx status code
func (o *PostProjectProjectIDTasksTaskIDStopNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this post project project Id tasks task Id stop no content response a status code equal to that given
func (o *PostProjectProjectIDTasksTaskIDStopNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the post project project Id tasks task Id stop no content response
func (o *PostProjectProjectIDTasksTaskIDStopNoContent) Code() int {
	return 204
}

func (o *PostProjectProjectIDTasksTaskIDStopNoContent) Error() string {
	return fmt.Sprintf("[POST /project/{project_id}/tasks/{task_id}/stop][%d] postProjectProjectIdTasksTaskIdStopNoContent", 204)
}

func (o *PostProjectProjectIDTasksTaskIDStopNoContent) String() string {
	return fmt.Sprintf("[POST /project/{project_id}/tasks/{task_id}/stop][%d] postProjectProjectIdTasksTaskIdStopNoContent", 204)
}

func (o *PostProjectProjectIDTasksTaskIDStopNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
