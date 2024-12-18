// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new authentication API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new authentication API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new authentication API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for authentication API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// This client is generated with a few options you might find useful for your swagger spec.
//
// Feel free to add you own set of options.

// WithAccept allows the client to force the Accept header
// to negotiate a specific Producer from the server.
//
// You may use this option to set arbitrary extensions to your MIME media type.
func WithAccept(mime string) ClientOption {
	return func(r *runtime.ClientOperation) {
		r.ProducesMediaTypes = []string{mime}
	}
}

// WithAcceptApplicationJSON sets the Accept header to "application/json".
func WithAcceptApplicationJSON(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"application/json"}
}

// WithAcceptTextPlainCharsetUTF8 sets the Accept header to "text/plain; charset=utf-8".
func WithAcceptTextPlainCharsetUTF8(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"text/plain; charset=utf-8"}
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteUserTokensAPITokenID(params *DeleteUserTokensAPITokenIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUserTokensAPITokenIDNoContent, error)

	GetAuthLogin(params *GetAuthLoginParams, opts ...ClientOption) (*GetAuthLoginOK, error)

	GetAuthOidcProviderIDLogin(params *GetAuthOidcProviderIDLoginParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) error

	GetAuthOidcProviderIDRedirect(params *GetAuthOidcProviderIDRedirectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) error

	GetUserTokens(params *GetUserTokensParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserTokensOK, error)

	PostAuthLogin(params *PostAuthLoginParams, opts ...ClientOption) (*PostAuthLoginNoContent, error)

	PostAuthLogout(params *PostAuthLogoutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostAuthLogoutNoContent, error)

	PostUserTokens(params *PostUserTokensParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostUserTokensCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteUserTokensAPITokenID expires API token
*/
func (a *Client) DeleteUserTokensAPITokenID(params *DeleteUserTokensAPITokenIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUserTokensAPITokenIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserTokensAPITokenIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteUserTokensAPITokenID",
		Method:             "DELETE",
		PathPattern:        "/user/tokens/{api_token_id}",
		ProducesMediaTypes: []string{"application/json", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteUserTokensAPITokenIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteUserTokensAPITokenIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteUserTokensAPITokenID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAuthLogin fetches login metadata

Fetches metadata for login, such as available OIDC providers
*/
func (a *Client) GetAuthLogin(params *GetAuthLoginParams, opts ...ClientOption) (*GetAuthLoginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAuthLoginParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAuthLogin",
		Method:             "GET",
		PathPattern:        "/auth/login",
		ProducesMediaTypes: []string{"application/json", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAuthLoginReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAuthLoginOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAuthLogin: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAuthOidcProviderIDLogin begins o ID c authentication flow and redirect to o ID c provider

The user agent is redirected to this endpoint when chosing to sign in via OIDC
*/
func (a *Client) GetAuthOidcProviderIDLogin(params *GetAuthOidcProviderIDLoginParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAuthOidcProviderIDLoginParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAuthOidcProviderIDLogin",
		Method:             "GET",
		PathPattern:        "/auth/oidc/{provider_id}/login",
		ProducesMediaTypes: []string{"application/json", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAuthOidcProviderIDLoginReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
GetAuthOidcProviderIDRedirect finishes o ID c authentication flow upon succes you will be logged in

The user agent is redirected here by the OIDC provider to complete authentication
*/
func (a *Client) GetAuthOidcProviderIDRedirect(params *GetAuthOidcProviderIDRedirectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAuthOidcProviderIDRedirectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAuthOidcProviderIDRedirect",
		Method:             "GET",
		PathPattern:        "/auth/oidc/{provider_id}/redirect",
		ProducesMediaTypes: []string{"application/json", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAuthOidcProviderIDRedirectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
GetUserTokens fetches API tokens for user
*/
func (a *Client) GetUserTokens(params *GetUserTokensParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserTokensOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserTokensParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetUserTokens",
		Method:             "GET",
		PathPattern:        "/user/tokens",
		ProducesMediaTypes: []string{"application/json", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetUserTokensReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUserTokensOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetUserTokens: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostAuthLogin performs login

Upon success you will be logged in
*/
func (a *Client) PostAuthLogin(params *PostAuthLoginParams, opts ...ClientOption) (*PostAuthLoginNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAuthLoginParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAuthLogin",
		Method:             "POST",
		PathPattern:        "/auth/login",
		ProducesMediaTypes: []string{"application/json", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostAuthLoginReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostAuthLoginNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAuthLogin: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostAuthLogout destroys current session
*/
func (a *Client) PostAuthLogout(params *PostAuthLogoutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostAuthLogoutNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAuthLogoutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAuthLogout",
		Method:             "POST",
		PathPattern:        "/auth/logout",
		ProducesMediaTypes: []string{"application/json", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostAuthLogoutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostAuthLogoutNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAuthLogout: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostUserTokens creates an API token
*/
func (a *Client) PostUserTokens(params *PostUserTokensParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostUserTokensCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostUserTokensParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostUserTokens",
		Method:             "POST",
		PathPattern:        "/user/tokens",
		ProducesMediaTypes: []string{"application/json", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostUserTokensReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostUserTokensCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostUserTokens: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
