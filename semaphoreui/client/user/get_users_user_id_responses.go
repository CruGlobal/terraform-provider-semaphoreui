// Code generated by go-swagger; DO NOT EDIT.

package user

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

// GetUsersUserIDReader is a Reader for the GetUsersUserID structure.
type GetUsersUserIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersUserIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsersUserIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /users/{user_id}/] GetUsersUserID", response, response.Code())
	}
}

// NewGetUsersUserIDOK creates a GetUsersUserIDOK with default headers values
func NewGetUsersUserIDOK() *GetUsersUserIDOK {
	return &GetUsersUserIDOK{}
}

/*
GetUsersUserIDOK describes a response with status code 200, with default header values.

User profile
*/
type GetUsersUserIDOK struct {
	Payload *models.User
}

// IsSuccess returns true when this get users user Id o k response has a 2xx status code
func (o *GetUsersUserIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get users user Id o k response has a 3xx status code
func (o *GetUsersUserIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users user Id o k response has a 4xx status code
func (o *GetUsersUserIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get users user Id o k response has a 5xx status code
func (o *GetUsersUserIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get users user Id o k response a status code equal to that given
func (o *GetUsersUserIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get users user Id o k response
func (o *GetUsersUserIDOK) Code() int {
	return 200
}

func (o *GetUsersUserIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/{user_id}/][%d] getUsersUserIdOK %s", 200, payload)
}

func (o *GetUsersUserIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/{user_id}/][%d] getUsersUserIdOK %s", 200, payload)
}

func (o *GetUsersUserIDOK) GetPayload() *models.User {
	return o.Payload
}

func (o *GetUsersUserIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}