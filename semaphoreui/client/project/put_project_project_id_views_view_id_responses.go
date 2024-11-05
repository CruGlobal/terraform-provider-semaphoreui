// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutProjectProjectIDViewsViewIDReader is a Reader for the PutProjectProjectIDViewsViewID structure.
type PutProjectProjectIDViewsViewIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutProjectProjectIDViewsViewIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutProjectProjectIDViewsViewIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /project/{project_id}/views/{view_id}] PutProjectProjectIDViewsViewID", response, response.Code())
	}
}

// NewPutProjectProjectIDViewsViewIDNoContent creates a PutProjectProjectIDViewsViewIDNoContent with default headers values
func NewPutProjectProjectIDViewsViewIDNoContent() *PutProjectProjectIDViewsViewIDNoContent {
	return &PutProjectProjectIDViewsViewIDNoContent{}
}

/*
PutProjectProjectIDViewsViewIDNoContent describes a response with status code 204, with default header values.

view updated
*/
type PutProjectProjectIDViewsViewIDNoContent struct {
}

// IsSuccess returns true when this put project project Id views view Id no content response has a 2xx status code
func (o *PutProjectProjectIDViewsViewIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put project project Id views view Id no content response has a 3xx status code
func (o *PutProjectProjectIDViewsViewIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put project project Id views view Id no content response has a 4xx status code
func (o *PutProjectProjectIDViewsViewIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put project project Id views view Id no content response has a 5xx status code
func (o *PutProjectProjectIDViewsViewIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put project project Id views view Id no content response a status code equal to that given
func (o *PutProjectProjectIDViewsViewIDNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the put project project Id views view Id no content response
func (o *PutProjectProjectIDViewsViewIDNoContent) Code() int {
	return 204
}

func (o *PutProjectProjectIDViewsViewIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /project/{project_id}/views/{view_id}][%d] putProjectProjectIdViewsViewIdNoContent", 204)
}

func (o *PutProjectProjectIDViewsViewIDNoContent) String() string {
	return fmt.Sprintf("[PUT /project/{project_id}/views/{view_id}][%d] putProjectProjectIdViewsViewIdNoContent", 204)
}

func (o *PutProjectProjectIDViewsViewIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
