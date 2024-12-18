// Code generated by go-swagger; DO NOT EDIT.

package project

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

// GetProjectProjectIDTasksLastReader is a Reader for the GetProjectProjectIDTasksLast structure.
type GetProjectProjectIDTasksLastReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectProjectIDTasksLastReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectProjectIDTasksLastOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /project/{project_id}/tasks/last] GetProjectProjectIDTasksLast", response, response.Code())
	}
}

// NewGetProjectProjectIDTasksLastOK creates a GetProjectProjectIDTasksLastOK with default headers values
func NewGetProjectProjectIDTasksLastOK() *GetProjectProjectIDTasksLastOK {
	return &GetProjectProjectIDTasksLastOK{}
}

/*
GetProjectProjectIDTasksLastOK describes a response with status code 200, with default header values.

Array of tasks in chronological order
*/
type GetProjectProjectIDTasksLastOK struct {
	Payload []*models.Task
}

// IsSuccess returns true when this get project project Id tasks last o k response has a 2xx status code
func (o *GetProjectProjectIDTasksLastOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get project project Id tasks last o k response has a 3xx status code
func (o *GetProjectProjectIDTasksLastOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project project Id tasks last o k response has a 4xx status code
func (o *GetProjectProjectIDTasksLastOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get project project Id tasks last o k response has a 5xx status code
func (o *GetProjectProjectIDTasksLastOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get project project Id tasks last o k response a status code equal to that given
func (o *GetProjectProjectIDTasksLastOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get project project Id tasks last o k response
func (o *GetProjectProjectIDTasksLastOK) Code() int {
	return 200
}

func (o *GetProjectProjectIDTasksLastOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /project/{project_id}/tasks/last][%d] getProjectProjectIdTasksLastOK %s", 200, payload)
}

func (o *GetProjectProjectIDTasksLastOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /project/{project_id}/tasks/last][%d] getProjectProjectIdTasksLastOK %s", 200, payload)
}

func (o *GetProjectProjectIDTasksLastOK) GetPayload() []*models.Task {
	return o.Payload
}

func (o *GetProjectProjectIDTasksLastOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
