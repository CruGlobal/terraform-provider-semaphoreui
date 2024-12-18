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

// GetProjectProjectIDViewsReader is a Reader for the GetProjectProjectIDViews structure.
type GetProjectProjectIDViewsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectProjectIDViewsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectProjectIDViewsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /project/{project_id}/views] GetProjectProjectIDViews", response, response.Code())
	}
}

// NewGetProjectProjectIDViewsOK creates a GetProjectProjectIDViewsOK with default headers values
func NewGetProjectProjectIDViewsOK() *GetProjectProjectIDViewsOK {
	return &GetProjectProjectIDViewsOK{}
}

/*
GetProjectProjectIDViewsOK describes a response with status code 200, with default header values.

view
*/
type GetProjectProjectIDViewsOK struct {
	Payload []*models.View
}

// IsSuccess returns true when this get project project Id views o k response has a 2xx status code
func (o *GetProjectProjectIDViewsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get project project Id views o k response has a 3xx status code
func (o *GetProjectProjectIDViewsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project project Id views o k response has a 4xx status code
func (o *GetProjectProjectIDViewsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get project project Id views o k response has a 5xx status code
func (o *GetProjectProjectIDViewsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get project project Id views o k response a status code equal to that given
func (o *GetProjectProjectIDViewsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get project project Id views o k response
func (o *GetProjectProjectIDViewsOK) Code() int {
	return 200
}

func (o *GetProjectProjectIDViewsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /project/{project_id}/views][%d] getProjectProjectIdViewsOK %s", 200, payload)
}

func (o *GetProjectProjectIDViewsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /project/{project_id}/views][%d] getProjectProjectIdViewsOK %s", 200, payload)
}

func (o *GetProjectProjectIDViewsOK) GetPayload() []*models.View {
	return o.Payload
}

func (o *GetProjectProjectIDViewsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
