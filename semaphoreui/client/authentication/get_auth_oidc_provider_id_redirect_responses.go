// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetAuthOidcProviderIDRedirectReader is a Reader for the GetAuthOidcProviderIDRedirect structure.
type GetAuthOidcProviderIDRedirectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthOidcProviderIDRedirectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 302:
		result := NewGetAuthOidcProviderIDRedirectFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /auth/oidc/{provider_id}/redirect] GetAuthOidcProviderIDRedirect", response, response.Code())
	}
}

// NewGetAuthOidcProviderIDRedirectFound creates a GetAuthOidcProviderIDRedirectFound with default headers values
func NewGetAuthOidcProviderIDRedirectFound() *GetAuthOidcProviderIDRedirectFound {
	return &GetAuthOidcProviderIDRedirectFound{}
}

/*
GetAuthOidcProviderIDRedirectFound describes a response with status code 302, with default header values.

Redirection to the Semaphore root URL on success, or to the login page on error
*/
type GetAuthOidcProviderIDRedirectFound struct {
}

// IsSuccess returns true when this get auth oidc provider Id redirect found response has a 2xx status code
func (o *GetAuthOidcProviderIDRedirectFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get auth oidc provider Id redirect found response has a 3xx status code
func (o *GetAuthOidcProviderIDRedirectFound) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get auth oidc provider Id redirect found response has a 4xx status code
func (o *GetAuthOidcProviderIDRedirectFound) IsClientError() bool {
	return false
}

// IsServerError returns true when this get auth oidc provider Id redirect found response has a 5xx status code
func (o *GetAuthOidcProviderIDRedirectFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get auth oidc provider Id redirect found response a status code equal to that given
func (o *GetAuthOidcProviderIDRedirectFound) IsCode(code int) bool {
	return code == 302
}

// Code gets the status code for the get auth oidc provider Id redirect found response
func (o *GetAuthOidcProviderIDRedirectFound) Code() int {
	return 302
}

func (o *GetAuthOidcProviderIDRedirectFound) Error() string {
	return fmt.Sprintf("[GET /auth/oidc/{provider_id}/redirect][%d] getAuthOidcProviderIdRedirectFound", 302)
}

func (o *GetAuthOidcProviderIDRedirectFound) String() string {
	return fmt.Sprintf("[GET /auth/oidc/{provider_id}/redirect][%d] getAuthOidcProviderIdRedirectFound", 302)
}

func (o *GetAuthOidcProviderIDRedirectFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}