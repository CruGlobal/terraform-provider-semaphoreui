package provider

import (
	"context"
	"fmt"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"os"
	"strconv"
	apiclient "terraform-provider-semaphoreui/semaphoreui/client"
)

var _ provider.Provider = &SemaphoreUIProvider{}
var _ provider.ProviderWithFunctions = &SemaphoreUIProvider{}

// SemaphoreUIProvider defines the provider implementation.
type SemaphoreUIProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// SemaphoreUIProviderModel describes the provider data model.
type SemaphoreUIProviderModel struct {
	Hostname types.String `tfsdk:"hostname"`
	Port     types.Int32  `tfsdk:"port"`
	Path     types.String `tfsdk:"path"`
	Protocol types.String `tfsdk:"protocol"`
	ApiToken types.String `tfsdk:"api_token"`
}

func (p *SemaphoreUIProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "semaphoreui"
	resp.Version = p.version
}

func (p *SemaphoreUIProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: `
Use the SemaphoreUI provider to interact with Semaphore UI via the API. You must configure the provider with the proper credentials before you can use it.

## API Token
You can generate a Semaphore API token by logging into Semaphore, opening the browser Developer Tools console, and running the following command:
` + "```javascript" + `
fetch("/api/user/tokens", {
  method: "POST",
  headers: {'Content-Type': 'application/json'}, 
  body: JSON.stringify({})
}).then(res => res.json()).then(data => console.log("api_token = " + data.id));
` + "```" + `
The token will be printed in the console. This token will grant the same level of access as the logged in user. Copy the token value and use it to configure the provider. The token is sensitive and should be treated as a secret. It is recommended to use the ` + "`SEMAPHOREUI_API_TOKEN`" + ` environment variable to configure the provider.
`,
		Attributes: map[string]schema.Attribute{
			"hostname": schema.StringAttribute{
				MarkdownDescription: "SemaphoreUI API hostname. This can also be defined by the `SEMAPHOREUI_HOSTNAME` environment variable. Example: `example.com`.",
				Optional:            true,
			},
			"port": schema.Int32Attribute{
				MarkdownDescription: "SemaphoreUI API port. This can also be defined by the `SEMAPHOREUI_PORT` environment variable. Default: `3000`.",
				Optional:            true,
			},
			"path": schema.StringAttribute{
				MarkdownDescription: "SemaphoreUI API base path. This can also be defined by the `SEMAPHOREUI_PATH` environment variable. Default: `/api`.",
				Optional:            true,
			},
			"protocol": schema.StringAttribute{
				MarkdownDescription: "SemaphoreUI API protocol. This can also be defined by the `SEMAPHOREUI_PROTOCOL` environment variable. Must be one of `http` or `https`. Default: `https`.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("http", "https"),
				},
			},
			"api_token": schema.StringAttribute{
				MarkdownDescription: "SemaphoreUI API token. This can also be defined by the `SEMAPHOREUI_API_TOKEN` environment variable.",
				Sensitive:           true,
				Optional:            true,
			},
		},
	}
}

func (p *SemaphoreUIProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config SemaphoreUIProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if config.Hostname.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("hostname"),
			"Unknown SemaphoreUI API Hostname",
			"The provider cannot create the SemaphoreUI API client as there is an unknown configuration value for the SemaphoreUI API hostname. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the SEMAPHOREUI_HOSTNAME environment variable.",
		)
	}

	if config.Port.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("port"),
			"Unknown SemaphoreUI API Port",
			"The provider cannot create the SemaphoreUI API client as there is an unknown configuration value for the SemaphoreUI API port. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the SEMAPHOREUI_PORT environment variable.",
		)
	}

	if config.Path.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("path"),
			"Unknown SemaphoreUI API Path",
			"The provider cannot create the SemaphoreUI API client as there is an unknown configuration value for the SemaphoreUI API path. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the SEMAPHOREUI_PATH environment variable.",
		)
	}

	if config.Protocol.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("protocol"),
			"Unknown SemaphoreUI API Protocol",
			"The provider cannot create the SemaphoreUI API client as there is an unknown configuration value for the SemaphoreUI API protocol. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the SEMAPHOREUI_PROTOCOL environment variable.",
		)
	}

	if config.ApiToken.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_token"),
			"Unknown SemaphoreUI API Token",
			"The provider cannot create the SemaphoreUI API client as there is an unknown configuration value for the SemaphoreUI API token. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the SEMAPHOREUI_API_TOKEN environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.
	hostname := os.Getenv("SEMAPHOREUI_HOSTNAME")
	port := os.Getenv("SEMAPHOREUI_PORT")
	basePath := os.Getenv("SEMAPHOREUI_PATH")
	protocol := os.Getenv("SEMAPHOREUI_PROTOCOL")
	apiToken := os.Getenv("SEMAPHOREUI_API_TOKEN")

	if !config.Hostname.IsNull() {
		hostname = config.Hostname.ValueString()
	}
	if !config.Port.IsNull() {
		port = strconv.Itoa(int(config.Port.ValueInt32()))
	}
	if !config.Path.IsNull() {
		basePath = config.Path.ValueString()
	}
	if !config.Protocol.IsNull() {
		protocol = config.Protocol.ValueString()
	}
	if !config.ApiToken.IsNull() {
		apiToken = config.ApiToken.ValueString()
	}

	// If any of the expected configurations are missing, use defaults or return
	// errors with provider-specific guidance.
	if hostname == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("hostname"),
			"Missing SemaphoreUI API Hostname",
			"Set the host value in the configuration or use the SEMAPHOREUI_HOSTNAME environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if apiToken == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_token"),
			"Missing SemaphoreUI API Token",
			"Set the API Token value in the configuration or use the SEMAPHOREUI_API_TOKEN environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if port == "" {
		port = "3000" // Default
	}
	if protocol == "" {
		protocol = "https" // Default
	}
	if basePath == "" {
		basePath = "/api" // Default
	}

	if resp.Diagnostics.HasError() {
		return
	}

	r := httptransport.New(fmt.Sprintf("%s:%s", hostname, port), basePath, []string{protocol})
	r.DefaultAuthentication = httptransport.BearerToken(apiToken)

	client := apiclient.New(r, strfmt.Default)
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *SemaphoreUIProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewProjectEnvironmentResource,
		NewProjectInventoryResource,
		NewProjectKeyResource,
		NewProjectRepositoryResource,
		NewProjectResource,
		NewProjectScheduleResource,
		NewProjectTemplateResource,
		NewProjectUserResource,
		NewUserResource,
	}
}

func (p *SemaphoreUIProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewProjectDataSource,
		NewProjectsDataSource,
		NewUserDataSource,
	}
}

func (p *SemaphoreUIProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &SemaphoreUIProvider{
			version: version,
		}
	}
}
