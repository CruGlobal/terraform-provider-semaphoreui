package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/types"
	apiclient "terraform-provider-semaphoreui/semaphoreui/client"
	"terraform-provider-semaphoreui/semaphoreui/client/user"
	"terraform-provider-semaphoreui/semaphoreui/models"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource = &externalUserDataSource{}
)

func NewExternalUserDataSource() datasource.DataSource {
	return &externalUserDataSource{}
}

type externalUserDataSource struct {
	client *apiclient.SemaphoreUI
}

func (d *externalUserDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *externalUserDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_external_user"
}

// Schema defines the schema for the data source.
func (d *externalUserDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ExternalUserSchema().GetDataSource(ctx)
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
	userRequest := models.UserRequest{
		Username: user.Username.ValueString(),
		Name:     user.Name.ValueString(),
		Email:    user.Email.ValueString(),
		External: true,
	}
	if !user.Name.IsUnknown() && !user.Name.IsNull() {
		userRequest.Name = user.Name.ValueString()
	} else {
		userRequest.Name = user.Username.ValueString()
	}
	if !user.Email.IsUnknown() && !user.Email.IsNull() {
		userRequest.Email = user.Email.ValueString()
	} else {
		userRequest.Email = user.Username.ValueString()
	}
	return &userRequest
}

func (r *externalUserDataSource) GetExternalUserByUsername(username string) (*ExternalUserModel, error) {
	response, err := r.client.User.GetUsers(&user.GetUsersParams{}, nil)
	if err != nil {
		return nil, fmt.Errorf("could not get users: %s", err.Error())
	}
	for _, usr := range response.Payload {
		if usr.Username == username {
			if !usr.External {
				return nil, fmt.Errorf("user with username %s is not an external user", username)
			}
			model := convertResponseToExternalUserModel(usr)
			return &model, nil
		}
	}
	return nil, fmt.Errorf("user with username %s not found", username)
}

func (d *externalUserDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ExternalUserModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Lookup user by username
	externalUser, err := d.GetExternalUserByUsername(config.Username.ValueString())
	if err != nil {
		// If user not found, create new user
		if err.Error() == fmt.Sprintf("user with username %s not found", config.Username.ValueString()) {
			response, err := d.client.User.PostUsers(&user.PostUsersParams{
				User: convertExternalUserModelToUserRequest(config),
			}, nil)
			if err != nil {
				resp.Diagnostics.AddError(
					"Error Reading SemaphoreUI User",
					"Could not create user, unexpected error: "+err.Error(),
				)
				return
			}
			usr := convertResponseToExternalUserModel(response.Payload)
			externalUser = &usr
		} else {
			resp.Diagnostics.AddError(
				"Error Reading SemaphoreUI User",
				err.Error(),
			)
			return
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, externalUser)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
