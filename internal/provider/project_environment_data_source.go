package provider

import (
	"context"
	apiclient "terraform-provider-semaphoreui/semaphoreui/client"
	"terraform-provider-semaphoreui/semaphoreui/client/project"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource = &projectEnvironmentDataSource{}
)

func NewProjectEnvironmentDataSource() datasource.DataSource {
	return &projectEnvironmentDataSource{}
}

type projectEnvironmentDataSource struct {
	client *apiclient.SemaphoreUI
}

func (d *projectEnvironmentDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*apiclient.SemaphoreUI)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			"Expected *client.SemaphoreUI, got %T. Please report this issue to the provider developers.",
		)
		return
	}
	d.client = client
}

// Metadata returns the data source type name.
func (d *projectEnvironmentDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project_environment"
}

// Schema defines the schema for the data source.
func (d *projectEnvironmentDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ProjectEnvironmentSchema().GetDataSource(ctx)
}

func (d *projectEnvironmentDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ProjectEnvironmentModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := d.client.Project.GetProjectProjectIDEnvironmentEnvironmentID(&project.GetProjectProjectIDEnvironmentEnvironmentIDParams{
		ProjectID:     config.ProjectID.ValueInt64(),
		EnvironmentID: config.ID.ValueInt64(),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Project Environment",
			"Could not read project environment, unexpected error: "+err.Error(),
		)
		return
	}
	model := convertEnvironmentResponseToProjectEnvironmentModel(ctx, response.Payload, &config)

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
