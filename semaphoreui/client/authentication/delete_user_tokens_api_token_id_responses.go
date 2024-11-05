// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteUserTokensAPITokenIDReader is a Reader for the DeleteUserTokensAPITokenID structure.
type DeleteUserTokensAPITokenIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserTokensAPITokenIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteUserTokensAPITokenIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /user/tokens/{api_token_id}] DeleteUserTokensAPITokenID", response, response.Code())
	}
}

// NewDeleteUserTokensAPITokenIDNoContent creates a DeleteUserTokensAPITokenIDNoContent with default headers values
func NewDeleteUserTokensAPITokenIDNoContent() *DeleteUserTokensAPITokenIDNoContent {
	return &DeleteUserTokensAPITokenIDNoContent{}
}

/*
DeleteUserTokensAPITokenIDNoContent describes a response with status code 204, with default header values.

Expired API Token
*/
type DeleteUserTokensAPITokenIDNoContent struct {
}

// IsSuccess returns true when this delete user tokens Api token Id no content response has a 2xx status code
func (o *DeleteUserTokensAPITokenIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete user tokens Api token Id no content response has a 3xx status code
func (o *DeleteUserTokensAPITokenIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user tokens Api token Id no content response has a 4xx status code
func (o *DeleteUserTokensAPITokenIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete user tokens Api token Id no content response has a 5xx status code
func (o *DeleteUserTokensAPITokenIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user tokens Api token Id no content response a status code equal to that given
func (o *DeleteUserTokensAPITokenIDNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete user tokens Api token Id no content response
func (o *DeleteUserTokensAPITokenIDNoContent) Code() int {
	return 204
}

func (o *DeleteUserTokensAPITokenIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /user/tokens/{api_token_id}][%d] deleteUserTokensApiTokenIdNoContent", 204)
}

func (o *DeleteUserTokensAPITokenIDNoContent) String() string {
	return fmt.Sprintf("[DELETE /user/tokens/{api_token_id}][%d] deleteUserTokensApiTokenIdNoContent", 204)
}

func (o *DeleteUserTokensAPITokenIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
