package provider

import (
	"context"
	"fmt"
	apiclient "terraform-provider-semaphoreui/semaphoreui/client"
	"terraform-provider-semaphoreui/semaphoreui/client/project"
	"terraform-provider-semaphoreui/semaphoreui/client/projects"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
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
func (d *projectDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ProjectSchema().GetDataSource(ctx)
}

func (d *projectDataSource) GetProjectByName(name string) (*ProjectModel, error) {
	response, err := d.client.Projects.GetProjects(&projects.GetProjectsParams{}, nil)
	if err != nil {
		return nil, fmt.Errorf("could not read Projects: %s", err.Error())
	}

	for _, project := range response.Payload {
		if project.Name == name {
			projectModel := convertProjectResponseToProjectModel(project)
			return &projectModel, nil
		}
	}
	return nil, fmt.Errorf("project with name %s not found", name)
}

func (d *projectDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ProjectModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var model ProjectModel
	if !config.ID.IsNull() && !config.ID.IsUnknown() {
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
		model = convertProjectResponseToProjectModel(response.Payload)
	} else if !config.Name.IsUnknown() && !config.Name.IsNull() {
		proj, err := d.GetProjectByName(config.Name.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Reading Semaphore Project",
				fmt.Sprintf("Could not read project: %s", err.Error()),
			)
			return
		}
		model = *proj
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
