package provider

import (
	"context"
	"fmt"
	apiclient "terraform-provider-semaphoreui/semaphoreui/client"
	"terraform-provider-semaphoreui/semaphoreui/client/project"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource = &projectRepositoryDataSource{}
)

func NewProjectRepositoryDataSource() datasource.DataSource {
	return &projectRepositoryDataSource{}
}

type projectRepositoryDataSource struct {
	client *apiclient.SemaphoreUI
}

func (d *projectRepositoryDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *projectRepositoryDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project_repository"
}

// Schema defines the schema for the data source.
func (d *projectRepositoryDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ProjectRepositorySchema().GetDataSource(ctx)
}

func (d *projectRepositoryDataSource) GetRepositoryByName(projectID int64, name string) (*ProjectRepositoryModel, error) {
	response, err := d.client.Project.GetProjectProjectIDRepositories(&project.GetProjectProjectIDRepositoriesParams{
		ProjectID: projectID,
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("could not read project repositories: %s", err.Error())
	}
	for _, repo := range response.Payload {
		if repo.Name == name {
			model := convertRepositoryResponseToProjectRepositoryModel(repo)
			return &model, nil
		}
	}
	return nil, fmt.Errorf("project repository with name %s not found", name)
}

func (d *projectRepositoryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ProjectRepositoryModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var model ProjectRepositoryModel
	if !config.ID.IsUnknown() && !config.ID.IsNull() {
		response, err := d.client.Project.GetProjectProjectIDRepositoriesRepositoryID(&project.GetProjectProjectIDRepositoriesRepositoryIDParams{
			ProjectID:    config.ProjectID.ValueInt64(),
			RepositoryID: config.ID.ValueInt64(),
		}, nil)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Reading SemaphoreUI Project Repository",
				"Could not read project repository, unexpected error: "+err.Error(),
			)
			return
		}
		model = convertRepositoryResponseToProjectRepositoryModel(response.Payload)
	} else if !config.Name.IsUnknown() && !config.Name.IsNull() {
		repo, err := d.GetRepositoryByName(config.ProjectID.ValueInt64(), config.Name.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Reading SemaphoreUI Project Repository",
				err.Error(),
			)
			return
		}
		model = *repo
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
