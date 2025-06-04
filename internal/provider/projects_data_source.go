package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/types"
	apiclient "terraform-provider-semaphoreui/semaphoreui/client"
	"terraform-provider-semaphoreui/semaphoreui/client/projects"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource = &projectsDataSource{}
)

func NewProjectsDataSource() datasource.DataSource {
	return &projectsDataSource{}
}

type projectsDataSource struct {
	client *apiclient.SemaphoreUI
}

func (d *projectsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *projectsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_projects"
}

type projectsDataSourceModel struct {
	Projects []ProjectModel `tfsdk:"projects"`
}

// Schema defines the schema for the data source.
func (d *projectsDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	projectAttributes := ProjectSchema().GetDataSource(ctx).Attributes
	projectAttributes["id"] = schema.Int64Attribute{
		MarkdownDescription: "The ID of the project.",
		Computed:            true,
	}
	projectAttributes["name"] = schema.StringAttribute{
		MarkdownDescription: "Project name.",
		Computed:            true,
	}
	resp.Schema = schema.Schema{
		MarkdownDescription: "Provides a List of SemaphoreUI Projects.",
		Attributes: map[string]schema.Attribute{
			"projects": schema.ListNestedAttribute{
				MarkdownDescription: "List of projects.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: projectAttributes,
				},
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *projectsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state projectsDataSourceModel

	response, err := d.client.Projects.GetProjects(&projects.GetProjectsParams{}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Semaphore Projects",
			fmt.Sprintf("Could not read projects: %s", err.Error()),
		)
		return
	}

	for _, project := range response.Payload {
		var maxParallelTasks types.Int64
		if project.MaxParallelTasks == nil {
			maxParallelTasks = types.Int64Value(0)
		} else {
			maxParallelTasks = types.Int64PointerValue(project.MaxParallelTasks)
		}

		state.Projects = append(state.Projects, ProjectModel{
			ID:               types.Int64Value(project.ID),
			Created:          types.StringValue(project.Created),
			Name:             types.StringValue(project.Name),
			Alert:            types.BoolPointerValue(project.Alert),
			AlertChat:        types.StringPointerValue(project.AlertChat),
			MaxParallelTasks: maxParallelTasks,
		})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
