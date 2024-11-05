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

// GetProjectProjectIDKeysReader is a Reader for the GetProjectProjectIDKeys structure.
type GetProjectProjectIDKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectProjectIDKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectProjectIDKeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /project/{project_id}/keys] GetProjectProjectIDKeys", response, response.Code())
	}
}

// NewGetProjectProjectIDKeysOK creates a GetProjectProjectIDKeysOK with default headers values
func NewGetProjectProjectIDKeysOK() *GetProjectProjectIDKeysOK {
	return &GetProjectProjectIDKeysOK{}
}

/*
GetProjectProjectIDKeysOK describes a response with status code 200, with default header values.

Access Keys
*/
type GetProjectProjectIDKeysOK struct {
	Payload []*models.AccessKey
}

// IsSuccess returns true when this get project project Id keys o k response has a 2xx status code
func (o *GetProjectProjectIDKeysOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get project project Id keys o k response has a 3xx status code
func (o *GetProjectProjectIDKeysOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project project Id keys o k response has a 4xx status code
func (o *GetProjectProjectIDKeysOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get project project Id keys o k response has a 5xx status code
func (o *GetProjectProjectIDKeysOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get project project Id keys o k response a status code equal to that given
func (o *GetProjectProjectIDKeysOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get project project Id keys o k response
func (o *GetProjectProjectIDKeysOK) Code() int {
	return 200
}

func (o *GetProjectProjectIDKeysOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /project/{project_id}/keys][%d] getProjectProjectIdKeysOK %s", 200, payload)
}

func (o *GetProjectProjectIDKeysOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /project/{project_id}/keys][%d] getProjectProjectIdKeysOK %s", 200, payload)
}

func (o *GetProjectProjectIDKeysOK) GetPayload() []*models.AccessKey {
	return o.Payload
}

func (o *GetProjectProjectIDKeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
