package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/types"
	apiclient "terraform-provider-semaphoreui/semaphoreui/client"
	"terraform-provider-semaphoreui/semaphoreui/client/project"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource = &projectDataSource{}
)

func NewProjectDataSource() datasource.DataSource {
	return &projectDataSource{}
}

type projectDataSource struct {
	client *apiclient.SemaphoreUI
}

func (d *projectDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *projectDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project"
}

// Schema defines the schema for the data source.
func (d *projectDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Lookup a SemaphoreUI Project.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				MarkdownDescription: "Project ID.",
				Required:            true,
			},
			"created": schema.StringAttribute{
				MarkdownDescription: "Creation date of the project.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Project name.",
				Computed:            true,
			},
			"alert": schema.BoolAttribute{
				MarkdownDescription: "Are alerts enabled for this project.",
				Computed:            true,
			},
			"alert_chat": schema.StringAttribute{
				MarkdownDescription: "Telegram chat ID.",
				Computed:            true,
			},
			"max_parallel_tasks": schema.Int64Attribute{
				MarkdownDescription: "Maximum number of parallel tasks, `0` for unlimited.",
				Computed:            true,
			},
		},
	}
}

func (d *projectDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config projectModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := d.client.Project.GetProjectProjectID(&project.GetProjectProjectIDParams{
		ProjectID: config.ID.ValueInt64(),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Semaphore Project",
			fmt.Sprintf("Could not read project: %s", err.Error()),
		)
		return
	}

	state := projectModel{
		ID:               types.Int64Value(response.Payload.ID),
		Created:          types.StringValue(response.Payload.Created),
		Name:             types.StringValue(response.Payload.Name),
		Alert:            types.BoolValue(response.Payload.Alert),
		AlertChat:        types.StringValue(response.Payload.AlertChat),
		MaxParallelTasks: types.Int64Value(*response.Payload.MaxParallelTasks),
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
