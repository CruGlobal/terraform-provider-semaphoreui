package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
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

type projectRepositoryModel struct {
	ID        types.Int64  `tfsdk:"id"`
	ProjectID types.Int64  `tfsdk:"project_id"`
	Name      types.String `tfsdk:"name"`
	Url       types.String `tfsdk:"url"`
	Branch    types.String `tfsdk:"branch"`
	SSHKeyID  types.Int64  `tfsdk:"ssh_key_id"`
}

func (r *projectRepositoryResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: `Provides a SemaphoreUI Project Repository resource.

SemaphoreUI currently supports only Git repositories, including GitHub HTTP/S protocols and local paths.`,
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				MarkdownDescription: "The repository ID.",
				Computed:            true,
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"project_id": schema.Int64Attribute{
				MarkdownDescription: "The project ID that the repository belongs to.",
				Required:            true,
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The display name of the repository.",
				Required:            true,
			},
			"url": schema.StringAttribute{
				MarkdownDescription: "The URI or path of the Git repository. SemaphoreUI supports `ssh`, `http`, `https`, `file` and `git` URI schemes as well as absolute paths.",
				Required:            true,
			},
			"branch": schema.StringAttribute{
				MarkdownDescription: "The branch of the repository to use. Use an empty string for path based repositories.",
				Required:            true,
			},
			"ssh_key_id": schema.Int64Attribute{
				MarkdownDescription: `The Project Key ID to use for accessing the Git repository.

This attribute is required for all repositories in SemaphoreUI. You should set it to the ID of a Key of type ` + "`none`" + ` if the repository doesn't require credentials.`,
				Required: true,
			},
		},
	}
}

func convertProjectRepositoryModelToRepositoryRequest(repo projectRepositoryModel) *models.RepositoryRequest {
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

func convertRepositoryResponseToProjectRepositoryModel(request *models.Repository) projectRepositoryModel {
	return projectRepositoryModel{
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
	var plan projectRepositoryModel
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
	var state projectRepositoryModel
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
	var plan projectRepositoryModel
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
	var state projectRepositoryModel
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
