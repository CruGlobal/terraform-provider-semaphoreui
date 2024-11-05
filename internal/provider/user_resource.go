package provider

import (
	"context"
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	apiclient "terraform-provider-semaphoreui/semaphoreui/client"
	"terraform-provider-semaphoreui/semaphoreui/client/user"
	"terraform-provider-semaphoreui/semaphoreui/models"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &userResource{}
	_ resource.ResourceWithConfigure   = &userResource{}
	_ resource.ResourceWithImportState = &userResource{}
)

func NewUserResource() resource.Resource {
	return &userResource{}
}

type userResource struct {
	client *apiclient.SemaphoreUI
}

func (r *userResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *userResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

type userModel struct {
	ID       types.Int64  `tfsdk:"id"`
	Created  types.String `tfsdk:"created"`
	Username types.String `tfsdk:"username"`
	Name     types.String `tfsdk:"name"`
	Email    types.String `tfsdk:"email"`
	Password types.String `tfsdk:"password"`
	Admin    types.Bool   `tfsdk:"admin"`
	External types.Bool   `tfsdk:"external"`
	Alert    types.Bool   `tfsdk:"alert"`
}

func (r *userResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: `
Provides a Semaphoreui User resource.`,
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				MarkdownDescription: "User ID.",
				Computed:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"created": schema.StringAttribute{
				MarkdownDescription: "Creation date of the user.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"username": schema.StringAttribute{
				MarkdownDescription: "Username.",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Display name.",
				Required:            true,
			},
			"email": schema.StringAttribute{
				MarkdownDescription: "Email address.",
				Required:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Login Password.",
				Optional:            true,
				Sensitive:           true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"admin": schema.BoolAttribute{
				MarkdownDescription: "Is the user an admin? Default `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"alert": schema.BoolAttribute{
				MarkdownDescription: "Send alerts to the user's email? Default `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"external": schema.BoolAttribute{
				MarkdownDescription: "Is the user linked to an external identity provider? Default `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func convertResponsePayloadToUserModel(user *models.User, password types.String) userModel {
	return userModel{
		ID:       types.Int64Value(user.ID),
		Created:  types.StringValue(user.Created),
		Username: types.StringValue(user.Username),
		Name:     types.StringValue(user.Name),
		Email:    types.StringValue(user.Email),
		Admin:    types.BoolValue(user.Admin),
		External: types.BoolValue(user.External),
		Alert:    types.BoolValue(user.Alert),
		// Password is not returned by the API so we use previously set password
		Password: password,
	}
}

func convertUserModelToUserRequest(user userModel) *models.UserRequest {
	return &models.UserRequest{
		Username: user.Username.ValueString(),
		Name:     user.Name.ValueString(),
		Email:    user.Email.ValueString(),
		Password: strfmt.Password(user.Password.ValueString()),
		Admin:    user.Admin.ValueBool(),
		Alert:    user.Alert.ValueBool(),
		External: user.External.ValueBool(),
	}
}

func convertUserModelToUserPutRequest(user userModel) *models.UserPutRequest {
	return &models.UserPutRequest{
		Username: user.Username.ValueString(),
		Name:     user.Name.ValueString(),
		Email:    user.Email.ValueString(),
		Admin:    user.Admin.ValueBool(),
		Alert:    user.Alert.ValueBool(),
	}
}

func (r *userResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan userModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	var payload = convertUserModelToUserRequest(plan)

	//Create new user
	response, err := r.client.User.PostUsers(&user.PostUsersParams{User: payload}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Semaphoreui User",
			"Could not create user, unexpected error: "+err.Error(),
		)
		return
	}

	// Map response body to schema and populate Computed attribute values
	plan = convertResponsePayloadToUserModel(response.Payload, plan.Password)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *userResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state userModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
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
	state = convertResponsePayloadToUserModel(response.Payload, state.Password)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *userResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan userModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	var payload = convertUserModelToUserPutRequest(plan)

	// Update existing resource
	_, err := r.client.User.PutUsersUserID(&user.PutUsersUserIDParams{UserID: plan.ID.ValueInt64(), User: payload}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Semaphore User",
			"Could not update user, unexpected error: "+err.Error(),
		)
		return
	}

	// Update password if it's changed
	var prevPassword types.String
	req.State.GetAttribute(ctx, path.Root("password"), &prevPassword)
	if plan.Password != prevPassword {
		_, err := r.client.User.PostUsersUserIDPassword(&user.PostUsersUserIDPasswordParams{UserID: plan.ID.ValueInt64(), Password: user.PostUsersUserIDPasswordBody{Password: strfmt.Password(plan.Password.ValueString())}}, nil)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Updating Semaphore User Password",
				"Could not update user password, unexpected error: "+err.Error(),
			)
		}
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
	plan = convertResponsePayloadToUserModel(response.Payload, plan.Password)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *userResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state userModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing resource
	_, err := r.client.User.DeleteUsersUserID(&user.DeleteUsersUserIDParams{UserID: state.ID.ValueInt64()}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Semaphore User",
			fmt.Sprintf("Could not delete user, unexpected error: %s", err.Error()),
		)
		return
	}

}

func (r *userResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	fields, err := parseImportFields(req.ID, []string{"user"})
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid User Import ID",
			"Could not parse import ID: "+err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), fields["user"])...)
}
