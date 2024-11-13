// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"terraform-provider-semaphoreui/semaphoreui/client/authentication"
	"terraform-provider-semaphoreui/semaphoreui/client/integration"
	"terraform-provider-semaphoreui/semaphoreui/client/operations"
	"terraform-provider-semaphoreui/semaphoreui/client/project"
	"terraform-provider-semaphoreui/semaphoreui/client/projects"
	"terraform-provider-semaphoreui/semaphoreui/client/schedule"
	"terraform-provider-semaphoreui/semaphoreui/client/user"
)

// Default semaphore UI HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost:3000"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https", "ws", "wss"}

// NewHTTPClient creates a new semaphore UI HTTP client.
func NewHTTPClient(formats strfmt.Registry) *SemaphoreUI {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new semaphore UI HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *SemaphoreUI {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new semaphore UI client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *SemaphoreUI {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(SemaphoreUI)
	cli.Transport = transport
	cli.Authentication = authentication.New(transport, formats)
	cli.Integration = integration.New(transport, formats)
	cli.Operations = operations.New(transport, formats)
	cli.Project = project.New(transport, formats)
	cli.Projects = projects.New(transport, formats)
	cli.Schedule = schedule.New(transport, formats)
	cli.User = user.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// SemaphoreUI is a client for semaphore UI
type SemaphoreUI struct {
	Authentication authentication.ClientService

	Integration integration.ClientService

	Operations operations.ClientService

	Project project.ClientService

	Projects projects.ClientService

	Schedule schedule.ClientService

	User user.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *SemaphoreUI) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Authentication.SetTransport(transport)
	c.Integration.SetTransport(transport)
	c.Operations.SetTransport(transport)
	c.Project.SetTransport(transport)
	c.Projects.SetTransport(transport)
	c.Schedule.SetTransport(transport)
	c.User.SetTransport(transport)
}