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

// GetProjectProjectIDInventoryReader is a Reader for the GetProjectProjectIDInventory structure.
type GetProjectProjectIDInventoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectProjectIDInventoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectProjectIDInventoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /project/{project_id}/inventory] GetProjectProjectIDInventory", response, response.Code())
	}
}

// NewGetProjectProjectIDInventoryOK creates a GetProjectProjectIDInventoryOK with default headers values
func NewGetProjectProjectIDInventoryOK() *GetProjectProjectIDInventoryOK {
	return &GetProjectProjectIDInventoryOK{}
}

/*
GetProjectProjectIDInventoryOK describes a response with status code 200, with default header values.

inventory
*/
type GetProjectProjectIDInventoryOK struct {
	Payload []*models.Inventory
}

// IsSuccess returns true when this get project project Id inventory o k response has a 2xx status code
func (o *GetProjectProjectIDInventoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get project project Id inventory o k response has a 3xx status code
func (o *GetProjectProjectIDInventoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project project Id inventory o k response has a 4xx status code
func (o *GetProjectProjectIDInventoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get project project Id inventory o k response has a 5xx status code
func (o *GetProjectProjectIDInventoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get project project Id inventory o k response a status code equal to that given
func (o *GetProjectProjectIDInventoryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get project project Id inventory o k response
func (o *GetProjectProjectIDInventoryOK) Code() int {
	return 200
}

func (o *GetProjectProjectIDInventoryOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /project/{project_id}/inventory][%d] getProjectProjectIdInventoryOK %s", 200, payload)
}

func (o *GetProjectProjectIDInventoryOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /project/{project_id}/inventory][%d] getProjectProjectIdInventoryOK %s", 200, payload)
}

func (o *GetProjectProjectIDInventoryOK) GetPayload() []*models.Inventory {
	return o.Payload
}

func (o *GetProjectProjectIDInventoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
