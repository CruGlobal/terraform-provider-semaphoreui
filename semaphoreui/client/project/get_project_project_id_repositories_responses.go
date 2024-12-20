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

// GetProjectProjectIDRepositoriesReader is a Reader for the GetProjectProjectIDRepositories structure.
type GetProjectProjectIDRepositoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectProjectIDRepositoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectProjectIDRepositoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /project/{project_id}/repositories] GetProjectProjectIDRepositories", response, response.Code())
	}
}

// NewGetProjectProjectIDRepositoriesOK creates a GetProjectProjectIDRepositoriesOK with default headers values
func NewGetProjectProjectIDRepositoriesOK() *GetProjectProjectIDRepositoriesOK {
	return &GetProjectProjectIDRepositoriesOK{}
}

/*
GetProjectProjectIDRepositoriesOK describes a response with status code 200, with default header values.

repositories
*/
type GetProjectProjectIDRepositoriesOK struct {
	Payload []*models.Repository
}

// IsSuccess returns true when this get project project Id repositories o k response has a 2xx status code
func (o *GetProjectProjectIDRepositoriesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get project project Id repositories o k response has a 3xx status code
func (o *GetProjectProjectIDRepositoriesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project project Id repositories o k response has a 4xx status code
func (o *GetProjectProjectIDRepositoriesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get project project Id repositories o k response has a 5xx status code
func (o *GetProjectProjectIDRepositoriesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get project project Id repositories o k response a status code equal to that given
func (o *GetProjectProjectIDRepositoriesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get project project Id repositories o k response
func (o *GetProjectProjectIDRepositoriesOK) Code() int {
	return 200
}

func (o *GetProjectProjectIDRepositoriesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /project/{project_id}/repositories][%d] getProjectProjectIdRepositoriesOK %s", 200, payload)
}

func (o *GetProjectProjectIDRepositoriesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /project/{project_id}/repositories][%d] getProjectProjectIdRepositoriesOK %s", 200, payload)
}

func (o *GetProjectProjectIDRepositoriesOK) GetPayload() []*models.Repository {
	return o.Payload
}

func (o *GetProjectProjectIDRepositoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
