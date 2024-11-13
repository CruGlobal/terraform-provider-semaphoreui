package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
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

type projectUserModel struct {
	ProjectID types.Int64  `tfsdk:"project_id"`
	UserID    types.Int64  `tfsdk:"user_id"`
	Username  types.String `tfsdk:"username"`
	Name      types.String `tfsdk:"name"`
	Role      types.String `tfsdk:"role"`
}

func (r *projectUserResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: `
Provides a SemaphoreUI Project User resource.

This resource is used to manages a user's role in a project.`,
		Attributes: map[string]schema.Attribute{
			"project_id": schema.Int64Attribute{
				MarkdownDescription: "ID of the project.",
				Required:            true,
			},
			"user_id": schema.Int64Attribute{
				MarkdownDescription: "ID of the user.",
				Required:            true,
			},
			"role": schema.StringAttribute{
				MarkdownDescription: "Role of the user in the project. Valid values are `\"owner\"`, `\"manager\"`, `\"task_runner\"` and `\"guest\"`.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("owner", "manager", "task_runner", "guest"),
				},
			},
			"username": schema.StringAttribute{
				MarkdownDescription: "Username of the user.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the user.",
				Computed:            true,
			},
		},
	}
}

func (r *projectUserResource) getProjectUserModelFromAPI(projectId types.Int64, userId types.Int64) (*projectUserModel, error) {
	payload, err := r.client.Project.GetProjectProjectIDUsers(&project.GetProjectProjectIDUsersParams{ProjectID: projectId.ValueInt64()}, nil)
	if err != nil {
		return nil, fmt.Errorf("could not read Users for project ID %d: %s", projectId.ValueInt64(), err.Error())
	}

	for _, projectUser := range payload.Payload {
		if projectUser.ID == userId.ValueInt64() {
			return &projectUserModel{
				ProjectID: projectId,
				UserID:    userId,
				Username:  types.StringValue(projectUser.Username),
				Name:      types.StringValue(projectUser.Name),
				Role:      types.StringValue(projectUser.Role),
			}, nil
		}
	}
	return nil, fmt.Errorf("user with ID %d not found in project with ID %d", userId.ValueInt64(), projectId.ValueInt64())
}

func (r *projectUserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan projectUserModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
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
	diags = resp.State.Set(ctx, &user)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *projectUserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state projectUserModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
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
	diags = resp.State.Set(ctx, &user)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *projectUserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan projectUserModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
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
	diags = resp.State.Set(ctx, &user)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *projectUserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state projectUserModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
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

	diags := resp.State.Set(ctx, &user)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
