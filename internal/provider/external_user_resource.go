package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	apiclient "terraform-provider-semaphoreui/semaphoreui/client"
	"terraform-provider-semaphoreui/semaphoreui/client/user"
	"terraform-provider-semaphoreui/semaphoreui/models"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &externalUserResource{}
	_ resource.ResourceWithConfigure   = &externalUserResource{}
	_ resource.ResourceWithImportState = &externalUserResource{}
)

func NewExternalUserResource() resource.Resource {
	return &externalUserResource{}
}

type externalUserResource struct {
	client *apiclient.SemaphoreUI
}

func (r *externalUserResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *externalUserResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_external_user"
}

func (r *externalUserResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ExternalUserSchema().GetResource(ctx)
}

func convertResponseToExternalUserModel(user *models.User) ExternalUserModel {
	return ExternalUserModel{
		ID:       types.Int64Value(user.ID),
		Username: types.StringValue(user.Username),
		Name:     types.StringValue(user.Name),
		Email:    types.StringValue(user.Email),
		Admin:    types.BoolValue(user.Admin),
		Alert:    types.BoolValue(user.Alert),
		External: types.BoolValue(user.External),
		Created:  types.StringValue(user.Created),
	}
}

func convertExternalUserModelToUserRequest(user ExternalUserModel) *models.UserRequest {
	return &models.UserRequest{
		Username: user.Username.ValueString(),
		Name:     user.Name.ValueString(),
		Email:    user.Email.ValueString(),
		Admin:    user.Admin.ValueBool(),
		Alert:    user.Alert.ValueBool(),
		External: user.External.ValueBool(),
	}
}

func convertExternalUserModelToUserPutRequest(user ExternalUserModel) *models.UserPutRequest {
	return &models.UserPutRequest{
		Username: user.Username.ValueString(),
		Name:     user.Name.ValueString(),
		Email:    user.Email.ValueString(),
		Admin:    user.Admin.ValueBool(),
		Alert:    user.Alert.ValueBool(),
	}
}

func (r *externalUserResource) GetExternalUserByUsername(username string) (*ExternalUserModel, error) {
	response, err := r.client.User.GetUsers(&user.GetUsersParams{}, nil)
	if err != nil {
		return nil, fmt.Errorf("could not get users: %s", err.Error())
	}
	for _, usr := range response.Payload {
		if usr.Username == username {
			model := convertResponseToExternalUserModel(usr)
			return &model, nil
		}
	}
	return nil, fmt.Errorf("user with username %s not found", username)
}

func (r *externalUserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan ExternalUserModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var id types.Int64

	// Lookup user by username
	externalUser, err := r.GetExternalUserByUsername(plan.Username.ValueString())
	if err != nil {
		// If user not found, create new user
		if err.Error() == fmt.Sprintf("user with username %s not found", plan.Username.ValueString()) {
			response, err := r.client.User.PostUsers(&user.PostUsersParams{
				User: convertExternalUserModelToUserRequest(plan),
			}, nil)
			if err != nil {
				resp.Diagnostics.AddError(
					"Error Creating SemaphoreUI User",
					"Could not create user, unexpected error: "+err.Error(),
				)
				return
			}
			id = convertResponseToExternalUserModel(response.Payload).ID
		} else {
			resp.Diagnostics.AddError(
				"Error Creating SemaphoreUI User",
				"Could not create user, unexpected error: "+err.Error(),
			)
			return
		}
	} else {
		// If user found, update the user
		plan.ID = externalUser.ID
		_, err := r.client.User.PutUsersUserID(&user.PutUsersUserIDParams{
			UserID: plan.ID.ValueInt64(),
			User:   convertExternalUserModelToUserPutRequest(plan),
		}, nil)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Creating Semaphore User",
				"Could not update user, unexpected error: "+err.Error(),
			)
			return
		}
		id = plan.ID
	}

	response, err := r.client.User.GetUsersUserID(&user.GetUsersUserIDParams{UserID: id.ValueInt64()}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Semaphore User",
			"Could not read user, unexpected error: "+err.Error(),
		)
		return
	}
	model := convertResponseToExternalUserModel(response.Payload)

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *externalUserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state ExternalUserModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed value from API
	response, err := r.client.User.GetUsersUserID(&user.GetUsersUserIDParams{UserID: state.ID.ValueInt64()}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Semaphore User",
			"Could not read user, unexpected error: "+err.Error(),
		)
		return
	}

	// Overwrite with refreshed state
	state = convertResponseToExternalUserModel(response.Payload)

	// Set refreshed state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *externalUserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan ExternalUserModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	var payload = convertExternalUserModelToUserPutRequest(plan)

	// Update existing resource
	_, err := r.client.User.PutUsersUserID(&user.PutUsersUserIDParams{UserID: plan.ID.ValueInt64(), User: payload}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Semaphore User",
			"Could not update user, unexpected error: "+err.Error(),
		)
		return
	}

	// Fetch updated values as PutUsersUserIDParams does not return updated user
	response, err := r.client.User.GetUsersUserID(&user.GetUsersUserIDParams{UserID: plan.ID.ValueInt64()}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Semaphore User",
			"Could not read user, unexpected error: "+err.Error(),
		)
		return
	}

	// Update resource state with updated user
	plan = convertResponseToExternalUserModel(response.Payload)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *externalUserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state ExternalUserModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// NOOP: External users are not deleted, just removed from state
}

func (r *externalUserResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	fields, err := parseImportFields(req.ID, []string{"user"})
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid External User Import ID",
			"Could not parse import ID: "+err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), fields["user"])...)
}
