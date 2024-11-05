// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new projects API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new projects API client with basic auth credentials.
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

// New creates a new projects API client with a bearer token for authentication.
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
Client for projects API
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
	GetProjects(params *GetProjectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProjectsOK, error)

	PostProjects(params *PostProjectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostProjectsCreated, error)

	PostProjectsRestore(params *PostProjectsRestoreParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostProjectsRestoreOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetProjects gets projects
*/
func (a *Client) GetProjects(params *GetProjectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProjectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProjectsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetProjects",
		Method:             "GET",
		PathPattern:        "/projects",
		ProducesMediaTypes: []string{"application/json", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetProjectsReader{formats: a.formats},
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
	success, ok := result.(*GetProjectsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetProjects: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostProjects creates a new project
*/
func (a *Client) PostProjects(params *PostProjectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostProjectsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostProjectsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostProjects",
		Method:             "POST",
		PathPattern:        "/projects",
		ProducesMediaTypes: []string{"application/json", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostProjectsReader{formats: a.formats},
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
	success, ok := result.(*PostProjectsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostProjects: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostProjectsRestore restores project
*/
func (a *Client) PostProjectsRestore(params *PostProjectsRestoreParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostProjectsRestoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostProjectsRestoreParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostProjectsRestore",
		Method:             "POST",
		PathPattern:        "/projects/restore",
		ProducesMediaTypes: []string{"application/json", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostProjectsRestoreReader{formats: a.formats},
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
	success, ok := result.(*PostProjectsRestoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostProjectsRestore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
