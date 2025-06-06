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

// GetProjectProjectIDTemplatesReader is a Reader for the GetProjectProjectIDTemplates structure.
type GetProjectProjectIDTemplatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectProjectIDTemplatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectProjectIDTemplatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /project/{project_id}/templates] GetProjectProjectIDTemplates", response, response.Code())
	}
}

// NewGetProjectProjectIDTemplatesOK creates a GetProjectProjectIDTemplatesOK with default headers values
func NewGetProjectProjectIDTemplatesOK() *GetProjectProjectIDTemplatesOK {
	return &GetProjectProjectIDTemplatesOK{}
}

/*
GetProjectProjectIDTemplatesOK describes a response with status code 200, with default header values.

template
*/
type GetProjectProjectIDTemplatesOK struct {
	Payload []*models.Template
}

// IsSuccess returns true when this get project project Id templates o k response has a 2xx status code
func (o *GetProjectProjectIDTemplatesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get project project Id templates o k response has a 3xx status code
func (o *GetProjectProjectIDTemplatesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project project Id templates o k response has a 4xx status code
func (o *GetProjectProjectIDTemplatesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get project project Id templates o k response has a 5xx status code
func (o *GetProjectProjectIDTemplatesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get project project Id templates o k response a status code equal to that given
func (o *GetProjectProjectIDTemplatesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get project project Id templates o k response
func (o *GetProjectProjectIDTemplatesOK) Code() int {
	return 200
}

func (o *GetProjectProjectIDTemplatesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /project/{project_id}/templates][%d] getProjectProjectIdTemplatesOK %s", 200, payload)
}

func (o *GetProjectProjectIDTemplatesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /project/{project_id}/templates][%d] getProjectProjectIdTemplatesOK %s", 200, payload)
}

func (o *GetProjectProjectIDTemplatesOK) GetPayload() []*models.Template {
	return o.Payload
}

func (o *GetProjectProjectIDTemplatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
