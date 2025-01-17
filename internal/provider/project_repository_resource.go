package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	apiclient "terraform-provider-semaphoreui/semaphoreui/client"
	"terraform-provider-semaphoreui/semaphoreui/client/project"
	"terraform-provider-semaphoreui/semaphoreui/models"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &projectRepositoryResource{}
	_ resource.ResourceWithConfigure   = &projectRepositoryResource{}
	_ resource.ResourceWithImportState = &projectRepositoryResource{}
)

func NewProjectRepositoryResource() resource.Resource {
	return &projectRepositoryResource{}
}

type projectRepositoryResource struct {
	client *apiclient.SemaphoreUI
}

func (r *projectRepositoryResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.client = client
}

func (r *projectRepositoryResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project_repository"
}

func (r *projectRepositoryResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ProjectRepositorySchema().GetResource(ctx)
}

func convertProjectRepositoryModelToRepositoryRequest(repo ProjectRepositoryModel) *models.RepositoryRequest {
	model := models.RepositoryRequest{
		ProjectID: repo.ProjectID.ValueInt64(),
		Name:      repo.Name.ValueString(),
		GitURL:    repo.Url.ValueString(),
		GitBranch: repo.Branch.ValueString(),
		SSHKeyID:  repo.SSHKeyID.ValueInt64(),
	}
	if !repo.ID.IsNull() && !repo.ID.IsUnknown() {
		model.ID = repo.ID.ValueInt64()
	}
	return &model
}

func convertRepositoryResponseToProjectRepositoryModel(request *models.Repository) ProjectRepositoryModel {
	return ProjectRepositoryModel{
		ID:        types.Int64Value(request.ID),
		ProjectID: types.Int64Value(request.ProjectID),
		Name:      types.StringValue(request.Name),
		Url:       types.StringValue(request.GitURL),
		Branch:    types.StringValue(request.GitBranch),
		SSHKeyID:  types.Int64Value(request.SSHKeyID),
	}
}

func (r *projectRepositoryResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan ProjectRepositoryModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := r.client.Project.PostProjectProjectIDRepositories(&project.PostProjectProjectIDRepositoriesParams{
		ProjectID:  plan.ProjectID.ValueInt64(),
		Repository: convertProjectRepositoryModelToRepositoryRequest(plan),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating SemaphoreUI Project Repository",
			"Could not create project repository, unexpected error: "+err.Error(),
		)
		return
	}
	model := convertRepositoryResponseToProjectRepositoryModel(response.Payload)

	// Set state to fully populated data
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *projectRepositoryResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state ProjectRepositoryModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := r.client.Project.GetProjectProjectIDRepositoriesRepositoryID(&project.GetProjectProjectIDRepositoriesRepositoryIDParams{
		ProjectID:    state.ProjectID.ValueInt64(),
		RepositoryID: state.ID.ValueInt64(),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Project Repository",
			"Could not read project repository, unexpected error: "+err.Error(),
		)
		return
	}
	model := convertRepositoryResponseToProjectRepositoryModel(response.Payload)

	// Set refreshed state
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *projectRepositoryResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan ProjectRepositoryModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.Project.PutProjectProjectIDRepositoriesRepositoryID(&project.PutProjectProjectIDRepositoriesRepositoryIDParams{
		ProjectID:    plan.ProjectID.ValueInt64(),
		RepositoryID: plan.ID.ValueInt64(),
		Repository:   convertProjectRepositoryModelToRepositoryRequest(plan),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating SemaphoreUI Project Repository",
			"Could not update project repository, unexpected error: "+err.Error(),
		)
		return
	}

	response, err := r.client.Project.GetProjectProjectIDRepositoriesRepositoryID(&project.GetProjectProjectIDRepositoriesRepositoryIDParams{
		ProjectID:    plan.ProjectID.ValueInt64(),
		RepositoryID: plan.ID.ValueInt64(),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Project Repository",
			"Could not read project repository, unexpected error: "+err.Error(),
		)
		return
	}
	model := convertRepositoryResponseToProjectRepositoryModel(response.Payload)

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *projectRepositoryResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state ProjectRepositoryModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.Project.DeleteProjectProjectIDRepositoriesRepositoryID(&project.DeleteProjectProjectIDRepositoriesRepositoryIDParams{
		ProjectID:    state.ProjectID.ValueInt64(),
		RepositoryID: state.ID.ValueInt64(),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Removing SemaphoreUI Project Repository",
			"Could not remove project repository, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *projectRepositoryResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	fields, err := parseImportFields(req.ID, []string{"project", "repository"})
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid Project Repository Import ID",
			"Could not parse import ID: "+err.Error(),
		)
		return
	}

	response, err := r.client.Project.GetProjectProjectIDRepositoriesRepositoryID(&project.GetProjectProjectIDRepositoriesRepositoryIDParams{
		ProjectID:    fields["project"],
		RepositoryID: fields["repository"],
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Project Repository",
			"Could not read project repository, unexpected error: "+err.Error(),
		)
		return
	}
	model := convertRepositoryResponseToProjectRepositoryModel(response.Payload)

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
