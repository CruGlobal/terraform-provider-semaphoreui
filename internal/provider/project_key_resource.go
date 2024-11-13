package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/resourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
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
	_ resource.Resource                     = &projectKeyResource{}
	_ resource.ResourceWithConfigure        = &projectKeyResource{}
	_ resource.ResourceWithImportState      = &projectKeyResource{}
	_ resource.ResourceWithConfigValidators = &projectKeyResource{}
)

func NewProjectKeyResource() resource.Resource {
	return &projectKeyResource{}
}

type projectKeyResource struct {
	client *apiclient.SemaphoreUI
}

func (r *projectKeyResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *projectKeyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project_key"
}

const (
	KeyLoginPassword string = "login_password"
	KeySSH           string = "ssh"
	KeyNone          string = "none"
)

type projectKeyModel struct {
	ID            types.Int64              `tfsdk:"id"`
	ProjectID     types.Int64              `tfsdk:"project_id"`
	Name          types.String             `tfsdk:"name"`
	LoginPassword *projectKeyLoginPassword `tfsdk:"login_password"`
	SSH           *projectKeySSH           `tfsdk:"ssh"`
	None          *projectKeyNone          `tfsdk:"none"`
}

type projectKeyLoginPassword struct {
	Login    types.String `tfsdk:"login"`
	Password types.String `tfsdk:"password"`
}

type projectKeySSH struct {
	Login      types.String `tfsdk:"login"`
	Passphrase types.String `tfsdk:"passphrase"`
	PrivateKey types.String `tfsdk:"private_key"`
}

type projectKeyNone struct{}

func (model *projectKeyModel) Type() types.String {
	if model.LoginPassword != nil {
		return types.StringValue(KeyLoginPassword)
	} else if model.SSH != nil {
		return types.StringValue(KeySSH)
	} else if model.None != nil {
		return types.StringValue(KeyNone)
	}
	return types.StringUnknown()
}

func (r *projectKeyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: `Provides a SemaphoreUI Project Key resource.

A project Key is used to define the credentials used through out a project. Credentials can be Username/Password, SSH key or None. Project keys are used throughout SemaphoreUI, including Inventories, Repositories, Environments and Templates. When a resource doesn't need credentials, the None Key is used.'`,
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				MarkdownDescription: "The Key ID",
				Computed:            true,
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"project_id": schema.Int64Attribute{
				MarkdownDescription: "The project ID that the key belongs to.",
				Required:            true,
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The display name of the key.",
				Required:            true,
			},
			KeyLoginPassword: schema.SingleNestedAttribute{
				MarkdownDescription: "A login password key.",
				Optional:            true,
				Attributes: map[string]schema.Attribute{
					"login": schema.StringAttribute{
						MarkdownDescription: "The login username.",
						Optional:            true,
					},
					"password": schema.StringAttribute{
						MarkdownDescription: "The login password.",
						Required:            true,
						Sensitive:           true,
					},
				},
			},
			KeySSH: schema.SingleNestedAttribute{
				MarkdownDescription: "A SSH key.",
				Optional:            true,
				Attributes: map[string]schema.Attribute{
					"login": schema.StringAttribute{
						MarkdownDescription: "The login username.",
						Optional:            true,
					},
					"passphrase": schema.StringAttribute{
						MarkdownDescription: "The SSH Key passphrase.",
						Optional:            true,
						Sensitive:           true,
					},
					"private_key": schema.StringAttribute{
						MarkdownDescription: "The SSH private key.",
						Required:            true,
						Sensitive:           true,
					},
				},
			},
			KeyNone: schema.SingleNestedAttribute{
				MarkdownDescription: "The special None key.",
				Optional:            true,
			},
		},
	}
}

func (r *projectKeyResource) ConfigValidators(ctx context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{
		resourcevalidator.ExactlyOneOf(
			path.MatchRoot(KeyLoginPassword),
			path.MatchRoot(KeySSH),
			path.MatchRoot(KeyNone),
		),
	}
}

func convertProjectKeyModelToAccessKeyRequest(key projectKeyModel) *models.AccessKeyRequest {
	model := models.AccessKeyRequest{
		ProjectID: key.ProjectID.ValueInt64(),
		Name:      key.Name.ValueString(),
	}
	if !key.ID.IsNull() && !key.ID.IsUnknown() {
		model.ID = key.ID.ValueInt64()
	}
	if key.None != nil {
		model.Type = KeyNone
	} else if key.LoginPassword != nil {
		model.Type = KeyLoginPassword
		model.LoginPassword = &models.AccessKeyRequestLoginPassword{
			Login:    key.LoginPassword.Login.ValueString(),
			Password: key.LoginPassword.Password.ValueString(),
		}
	} else if key.SSH != nil {
		model.Type = KeySSH
		model.SSH = &models.AccessKeyRequestSSH{
			Login:      key.SSH.Login.ValueString(),
			Passphrase: key.SSH.Passphrase.ValueString(),
			PrivateKey: key.SSH.PrivateKey.ValueString(),
		}
	}

	return &model
}

func convertAccessKeyResponseToProjectKeyModel(key *models.AccessKey, prev *projectKeyModel) projectKeyModel {
	model := projectKeyModel{
		ID:        types.Int64Value(key.ID),
		ProjectID: types.Int64Value(key.ProjectID),
		Name:      types.StringValue(key.Name),
	}

	// SemaphoreUI API never returns secret value, so we use the ones from the previous state
	switch key.Type {
	case KeyNone:
		model.None = &projectKeyNone{}
	case KeyLoginPassword:
		model.LoginPassword = prev.LoginPassword
	case KeySSH:
		model.SSH = prev.SSH
	}

	return model
}

func (r *projectKeyResource) getProjectKeyModelFromClient(projectId types.Int64, keyId types.Int64, prev *projectKeyModel) (*projectKeyModel, error) {
	payload, err := r.client.Project.GetProjectProjectIDKeys(&project.GetProjectProjectIDKeysParams{
		ProjectID: projectId.ValueInt64(),
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("could not read Keys for project ID %d: %s", projectId.ValueInt64(), err.Error())
	}

	for _, key := range payload.Payload {
		if key.ID == keyId.ValueInt64() {
			model := projectKeyModel{
				ProjectID: projectId,
				ID:        keyId,
				Name:      types.StringValue(key.Name),
			}
			switch key.Type {
			case KeyNone:
				model.None = &projectKeyNone{}
			case KeyLoginPassword:
				model.LoginPassword = prev.LoginPassword
			case KeySSH:
				model.SSH = prev.SSH
			}
			return &model, nil
		}
	}
	return nil, fmt.Errorf("key with ID %d not found in project with ID %d", keyId.ValueInt64(), projectId.ValueInt64())
}

func (r *projectKeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan projectKeyModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := r.client.Project.PostProjectProjectIDKeys(&project.PostProjectProjectIDKeysParams{
		ProjectID: plan.ProjectID.ValueInt64(),
		AccessKey: convertProjectKeyModelToAccessKeyRequest(plan),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating SemaphoreUI Project Key",
			"Could not create project key, unexpected error: "+err.Error(),
		)
		return
	}
	plan = convertAccessKeyResponseToProjectKeyModel(response.Payload, &plan)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *projectKeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state projectKeyModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	model, err := r.getProjectKeyModelFromClient(state.ProjectID, state.ID, &state)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Semaphore Project Keys",
			err.Error(),
		)
		return
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, &model)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *projectKeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan and state
	var plan, state projectKeyModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create an access key based on the plan
	key := convertProjectKeyModelToAccessKeyRequest(plan)
	// Check if type of key has changed
	if !plan.Type().Equal(state.Type()) {
		// If key type has changed, we must update the secrets
		key.OverrideSecret = true
	} else {
		// Key type has not changed, so we need to check if individual fields have changed
		switch plan.Type().ValueString() {
		case KeyLoginPassword:
			if !plan.LoginPassword.Login.Equal(state.LoginPassword.Login) ||
				!plan.LoginPassword.Password.Equal(state.LoginPassword.Password) ||
				!plan.Name.Equal(state.Name) {
				key.OverrideSecret = true
			} else {
				// Use empty struct when secrets haven't changed
				key.LoginPassword = &models.AccessKeyRequestLoginPassword{}
			}
		case KeySSH:
			if !plan.SSH.Login.Equal(state.SSH.Login) ||
				!plan.SSH.Passphrase.Equal(state.SSH.Passphrase) ||
				!plan.SSH.PrivateKey.Equal(state.SSH.PrivateKey) ||
				!plan.Name.Equal(state.Name) {
				key.OverrideSecret = true
			} else {
				// Use empty struct when secrets haven't changed
				key.SSH = &models.AccessKeyRequestSSH{}
			}
		case KeyNone:
			// type None has no secrets to update, but if Name change, we need to update it
			if !plan.Name.Equal(state.Name) {
				key.OverrideSecret = true
			}
		}
	}
	if key.OverrideSecret {
		// For some reason, the API will only update access keys when all the fields are set
		// So we need to set the other fields to empty strings
		switch key.Type {
		case KeyLoginPassword:
			key.SSH = &models.AccessKeyRequestSSH{
				Login:      "",
				Passphrase: "",
				PrivateKey: "",
			}
		case KeySSH:
			key.LoginPassword = &models.AccessKeyRequestLoginPassword{
				Login:    "",
				Password: "",
			}
		}
	}

	// Update existing resource
	_, err := r.client.Project.PutProjectProjectIDKeysKeyID(&project.PutProjectProjectIDKeysKeyIDParams{
		ProjectID: plan.ProjectID.ValueInt64(),
		KeyID:     plan.ID.ValueInt64(),
		AccessKey: key,
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating SemaphoreUI Project Key",
			"Could not update project key, unexpected error: "+err.Error(),
		)
		return
	}

	// Fetch updated values as PutProjectProjectIDKeysKeyID does not return updated projectKey
	model, err := r.getProjectKeyModelFromClient(state.ProjectID, state.ID, &plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Semaphore Project Keys",
			err.Error(),
		)
		return
	}

	// Update resource state with updated projectKey
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *projectKeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state projectKeyModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing resource
	_, err := r.client.Project.DeleteProjectProjectIDKeysKeyID(&project.DeleteProjectProjectIDKeysKeyIDParams{
		ProjectID: state.ProjectID.ValueInt64(),
		KeyID:     state.ID.ValueInt64(),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Semaphore Project Key",
			fmt.Sprintf("Could not delete project key, unexpected error: %s", err.Error()),
		)
		return
	}
}

func (r *projectKeyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	fields, err := parseImportFields(req.ID, []string{"project", "key"})
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid Project Key Import ID",
			"Could not parse import ID: "+err.Error(),
		)
		return
	}

	// Get the project key from the client filling required secrets with empty strings
	model, err := r.getProjectKeyModelFromClient(types.Int64Value(fields["project"]), types.Int64Value(fields["key"]), &projectKeyModel{
		LoginPassword: &projectKeyLoginPassword{
			Password: types.StringValue(""),
		},
		SSH: &projectKeySSH{
			PrivateKey: types.StringValue(""),
		},
		None: &projectKeyNone{},
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Semaphore Project Keys",
			err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
