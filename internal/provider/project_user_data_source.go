package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/types"
	apiclient "terraform-provider-semaphoreui/semaphoreui/client"
	"terraform-provider-semaphoreui/semaphoreui/client/project"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource = &projectUserDataSource{}
)

func NewProjectUserDataSource() datasource.DataSource {
	return &projectUserDataSource{}
}

type projectUserDataSource struct {
	client *apiclient.SemaphoreUI
}

func (d *projectUserDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *projectUserDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project_user"
}

// Schema defines the schema for the data source.
func (d *projectUserDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ProjectUserSchema().GetDataSource(ctx)
}

func (d *projectUserDataSource) getProjectUserModelFromAPI(projectId types.Int64, userId types.Int64) (*ProjectUserModel, error) {
	payload, err := d.client.Project.GetProjectProjectIDUsers(&project.GetProjectProjectIDUsersParams{ProjectID: projectId.ValueInt64()}, nil)
	if err != nil {
		return nil, fmt.Errorf("could not read Users for project ID %d: %s", projectId.ValueInt64(), err.Error())
	}

	for _, projectUser := range payload.Payload {
		if projectUser.ID == userId.ValueInt64() {
			return &ProjectUserModel{
				ProjectID: projectId,
				UserID:    userId,
				Role:      types.StringValue(projectUser.Role),
				Username:  types.StringValue(projectUser.Username),
				Name:      types.StringValue(projectUser.Name),
			}, nil
		}
	}
	return nil, fmt.Errorf("user with ID %d not found in project with ID %d", userId.ValueInt64(), projectId.ValueInt64())
}

func (d *projectUserDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ProjectUserModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	state, err := d.getProjectUserModelFromAPI(config.ProjectID, config.UserID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Semaphore Project Users",
			err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
