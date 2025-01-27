package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	apiclient "terraform-provider-semaphoreui/semaphoreui/client"
	"terraform-provider-semaphoreui/semaphoreui/client/project"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource = &projectViewDataSource{}
)

func NewProjectViewDataSource() datasource.DataSource {
	return &projectViewDataSource{}
}

type projectViewDataSource struct {
	client *apiclient.SemaphoreUI
}

func (d *projectViewDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *projectViewDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project_view"
}

// Schema defines the schema for the data source.
func (d *projectViewDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ProjectViewSchema().GetDataSource(ctx)
}

func (d *projectViewDataSource) GetViewModelByTitle(projectID int64, title types.String) (*ProjectViewModel, error) {
	payload, err := d.client.Project.GetProjectProjectIDViews(&project.GetProjectProjectIDViewsParams{
		ProjectID: projectID,
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("could not read Views: %s", err.Error())
	}

	for _, view := range payload.Payload {
		if view.Title == title.ValueString() {
			viewModel := convertViewResponseToProjectViewModel(view)
			return &viewModel, nil
		}
	}
	return nil, fmt.Errorf("view with title %s not found", title.ValueString())
}

func (d *projectViewDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ProjectViewModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var model ProjectViewModel
	if !config.ID.IsNull() && !config.ID.IsUnknown() {
		response, err := d.client.Project.GetProjectProjectIDViewsViewID(&project.GetProjectProjectIDViewsViewIDParams{
			ProjectID: config.ProjectID.ValueInt64(),
			ViewID:    config.ID.ValueInt64(),
		}, nil)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Reading SemaphoreUI Project View",
				"Could not read project view, unexpected error: "+err.Error(),
			)
			return
		}
		model = convertViewResponseToProjectViewModel(response.Payload)
	} else if !config.Title.IsNull() && !config.Title.IsUnknown() {
		view, err := d.GetViewModelByTitle(config.ProjectID.ValueInt64(), config.Title)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Reading SemaphoreUI Project View",
				err.Error(),
			)
			return
		}
		model = *view
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
