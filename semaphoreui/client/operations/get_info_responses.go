// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// GetInfoReader is a Reader for the GetInfo structure.
type GetInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /info] GetInfo", response, response.Code())
	}
}

// NewGetInfoOK creates a GetInfoOK with default headers values
func NewGetInfoOK() *GetInfoOK {
	return &GetInfoOK{}
}

/*
GetInfoOK describes a response with status code 200, with default header values.

ok
*/
type GetInfoOK struct {
	Payload *models.InfoType
}

// IsSuccess returns true when this get info o k response has a 2xx status code
func (o *GetInfoOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get info o k response has a 3xx status code
func (o *GetInfoOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get info o k response has a 4xx status code
func (o *GetInfoOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get info o k response has a 5xx status code
func (o *GetInfoOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get info o k response a status code equal to that given
func (o *GetInfoOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get info o k response
func (o *GetInfoOK) Code() int {
	return 200
}

func (o *GetInfoOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /info][%d] getInfoOK %s", 200, payload)
}

func (o *GetInfoOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /info][%d] getInfoOK %s", 200, payload)
}

func (o *GetInfoOK) GetPayload() *models.InfoType {
	return o.Payload
}

func (o *GetInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfoType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
