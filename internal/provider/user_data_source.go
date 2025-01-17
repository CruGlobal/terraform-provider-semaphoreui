package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/types"
	apiclient "terraform-provider-semaphoreui/semaphoreui/client"
	"terraform-provider-semaphoreui/semaphoreui/client/user"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource = &userDataSource{}
)

func NewUserDataSource() datasource.DataSource {
	return &userDataSource{}
}

type userDataSource struct {
	client *apiclient.SemaphoreUI
}

func (d *userDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *userDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

// Schema defines the schema for the data source.
func (d *userDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = userSchema().GetDataSource(ctx)
}

func (d *userDataSource) GetUserModelByUsername(username types.String) (*UserModel, error) {
	payload, err := d.client.User.GetUsers(&user.GetUsersParams{}, nil)
	if err != nil {
		return nil, fmt.Errorf("could not read Users: %s", err.Error())
	}

	for _, u := range payload.Payload {
		if u.Username == username.ValueString() {
			userModel := convertResponsePayloadToUserModel(u, UserModel{Password: types.StringValue("")})
			return &userModel, nil
		}
	}
	return nil, fmt.Errorf("user with username %s not found", username.ValueString())
}

func (d *userDataSource) GetUserModelByEmail(email types.String) (*UserModel, error) {
	payload, err := d.client.User.GetUsers(&user.GetUsersParams{}, nil)
	if err != nil {
		return nil, fmt.Errorf("could not read Users: %s", err.Error())
	}

	for _, u := range payload.Payload {
		if u.Email == email.ValueString() {
			userModel := convertResponsePayloadToUserModel(u, UserModel{Password: types.StringValue("")})
			return &userModel, nil
		}
	}
	return nil, fmt.Errorf("user with email %s not found", email.ValueString())
}

func (d *userDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config UserModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state UserModel
	if !config.ID.IsNull() && !config.ID.IsUnknown() {
		response, err := d.client.User.GetUsersUserID(&user.GetUsersUserIDParams{
			UserID: config.ID.ValueInt64(),
		}, nil)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Reading Semaphore User",
				fmt.Sprintf("Could not read user: %s", err.Error()),
			)
			return
		}
		state = convertResponsePayloadToUserModel(response.Payload, UserModel{Password: types.StringValue("")})
	} else if !config.Username.IsNull() && !config.Username.IsUnknown() {
		u, err := d.GetUserModelByUsername(config.Username)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Reading Semaphore User",
				err.Error(),
			)
			return
		}
		state = *u
	} else if !config.Email.IsNull() && !config.Email.IsUnknown() {
		u, err := d.GetUserModelByEmail(config.Email)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Reading Semaphore User",
				err.Error(),
			)
			return
		}
		state = *u
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
