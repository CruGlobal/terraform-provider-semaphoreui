package provider

import (
	"context"
	apiclient "terraform-provider-semaphoreui/semaphoreui/client"
	"terraform-provider-semaphoreui/semaphoreui/client/schedule"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource = &projectScheduleDataSource{}
)

func NewProjectScheduleDataSource() datasource.DataSource {
	return &projectScheduleDataSource{}
}

type projectScheduleDataSource struct {
	client *apiclient.SemaphoreUI
}

func (d *projectScheduleDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *projectScheduleDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project_schedule"
}

// Schema defines the schema for the data source.
func (d *projectScheduleDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ProjectScheduleSchema().GetDataSource(ctx)
}

func (d *projectScheduleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ProjectScheduleModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := d.client.Schedule.GetProjectProjectIDSchedulesScheduleID(&schedule.GetProjectProjectIDSchedulesScheduleIDParams{
		ProjectID:  config.ProjectID.ValueInt64(),
		ScheduleID: config.ID.ValueInt64(),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Project Schedule",
			"Could not read project schedule, unexpected error: "+err.Error(),
		)
		return
	}
	model := convertScheduleResponseToProjectScheduleModel(response.Payload)

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
