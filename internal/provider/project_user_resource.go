package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	apiclient "terraform-provider-semaphoreui/semaphoreui/client"
	"terraform-provider-semaphoreui/semaphoreui/client/project"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &projectUserResource{}
	_ resource.ResourceWithConfigure   = &projectUserResource{}
	_ resource.ResourceWithImportState = &projectUserResource{}
)

func NewProjectUserResource() resource.Resource {
	return &projectUserResource{}
}

type projectUserResource struct {
	client *apiclient.SemaphoreUI
}

func (r *projectUserResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *projectUserResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project_user"
}

func (r *projectUserResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ProjectUserSchema().GetResource(ctx)
}

func (r *projectUserResource) getProjectUserModelFromAPI(projectId types.Int64, userId types.Int64) (*ProjectUserModel, error) {
	payload, err := r.client.Project.GetProjectProjectIDUsers(&project.GetProjectProjectIDUsersParams{ProjectID: projectId.ValueInt64()}, nil)
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

func (r *projectUserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan ProjectUserModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	//Create new projectUser
	_, err := r.client.Project.PostProjectProjectIDUsers(
		&project.PostProjectProjectIDUsersParams{
			ProjectID: plan.ProjectID.ValueInt64(),
			User: project.PostProjectProjectIDUsersBody{
				Role:   plan.Role.ValueString(),
				UserID: plan.UserID.ValueInt64(),
			},
		}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating SemaphoreUI Project User",
			"Could not create project user, unexpected error: "+err.Error(),
		)
		return
	}

	// Fetch updated values as PostProjectProjectIDUsers does not return updated projectUser
	user, err := r.getProjectUserModelFromAPI(plan.ProjectID, plan.UserID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Semaphore Project Users",
			err.Error(),
		)
		return
	}

	// Set state to fully populated data
	resp.Diagnostics.Append(resp.State.Set(ctx, &user)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *projectUserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state ProjectUserModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed value from API
	user, err := r.getProjectUserModelFromAPI(state.ProjectID, state.UserID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Semaphore Project Users",
			err.Error(),
		)
		return
	}

	// Set refreshed state
	resp.Diagnostics.Append(resp.State.Set(ctx, &user)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *projectUserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan ProjectUserModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update existing resource
	_, err := r.client.Project.PutProjectProjectIDUsersUserID(
		&project.PutProjectProjectIDUsersUserIDParams{
			ProjectID: plan.ProjectID.ValueInt64(),
			UserID:    plan.UserID.ValueInt64(),
			ProjectUser: project.PutProjectProjectIDUsersUserIDBody{
				Role: plan.Role.ValueString(),
			},
		}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Semaphore Project User",
			"Could not update project user, unexpected error: "+err.Error(),
		)
		return
	}

	// Fetch updated values as PutProjectProjectIDUsersUserID does not return updated projectUser
	user, err := r.getProjectUserModelFromAPI(plan.ProjectID, plan.UserID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Semaphore Project Users",
			err.Error(),
		)
		return
	}

	// Update resource state with updated projectUser
	resp.Diagnostics.Append(resp.State.Set(ctx, &user)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *projectUserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state ProjectUserModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing resource
	_, err := r.client.Project.DeleteProjectProjectIDUsersUserID(&project.DeleteProjectProjectIDUsersUserIDParams{
		ProjectID: state.ProjectID.ValueInt64(),
		UserID:    state.UserID.ValueInt64(),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Removing Semaphore Project User",
			"Could not remove project user, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *projectUserResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	fields, err := parseImportFields(req.ID, []string{"project", "user"})
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid ProjectUser Import ID",
			"Could not parse import ID: "+err.Error(),
		)
		return
	}

	user, err := r.getProjectUserModelFromAPI(types.Int64Value(fields["project"]), types.Int64Value(fields["user"]))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Importing Semaphore Project User",
			err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &user)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
