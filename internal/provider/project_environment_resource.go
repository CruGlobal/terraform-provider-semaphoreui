package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"sort"
	apiclient "terraform-provider-semaphoreui/semaphoreui/client"
	"terraform-provider-semaphoreui/semaphoreui/client/project"
	"terraform-provider-semaphoreui/semaphoreui/models"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &projectEnvironmentResource{}
	_ resource.ResourceWithConfigure   = &projectEnvironmentResource{}
	_ resource.ResourceWithImportState = &projectEnvironmentResource{}
)

func NewProjectEnvironmentResource() resource.Resource {
	return &projectEnvironmentResource{}
}

type projectEnvironmentResource struct {
	client *apiclient.SemaphoreUI
}

func (r *projectEnvironmentResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *projectEnvironmentResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project_environment"
}

type projectEnvironmentModel struct {
	ID          types.Int64        `tfsdk:"id"`
	ProjectID   types.Int64        `tfsdk:"project_id"`
	Name        types.String       `tfsdk:"name"`
	Variables   *map[string]string `tfsdk:"variables"`
	Environment *map[string]string `tfsdk:"environment"`
	Secrets     types.List         `tfsdk:"secrets"`
}

func (projectEnvironmentModel) GetSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: `Provides a SemaphoreUI Project Environment resource.

A project environment provides a list of extra and environment variables that can be used in a project's templates.'`,
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				MarkdownDescription: "The environment ID.",
				Computed:            true,
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"project_id": schema.Int64Attribute{
				MarkdownDescription: "The project ID that the environment belongs to.",
				Required:            true,
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The display name of the environment.",
				Required:            true,
			},
			"variables": schema.MapAttribute{
				MarkdownDescription: "Extra variables. Passed to Ansible as extra variables (`--extra-vars`) and Terraform/OpenTofu as variables (`-var`).",
				ElementType:         types.StringType,
				Optional:            true,
			},
			"environment": schema.MapAttribute{
				MarkdownDescription: "Environment variables.",
				ElementType:         types.StringType,
				Optional:            true,
			},
			"secrets": schema.ListNestedAttribute{
				MarkdownDescription: "Secret variables of either `\"var\"` or `\"env\"` type. The `value` is encrypted and will be empty if imported.",
				Optional:            true,
				NestedObject:        projectEnvironmentSecretModel{}.GetSchema(),
			},
		},
	}
}

type projectEnvironmentSecretModel struct {
	ID    types.Int64  `tfsdk:"id"`
	Type  types.String `tfsdk:"type"`
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

func (projectEnvironmentSecretModel) GetSchema() schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		PlanModifiers: []planmodifier.Object{
			objectplanmodifier.UseStateForUnknown(),
		},
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				MarkdownDescription: "The variable ID.",
				Computed:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "The variable type. Either `\"env\"` or `\"var\"`.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("env", "var"),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The variable name.",
				Required:            true,
			},
			"value": schema.StringAttribute{
				MarkdownDescription: "The variable value.",
				Required:            true,
				Sensitive:           true,
			},
		},
	}
}

func (model projectEnvironmentModel) SecretValue(ctx context.Context, name string, varType string) types.String {
	if model.Secrets.IsNull() || model.Secrets.IsUnknown() {
		return types.StringValue("")
	}
	var secrets []projectEnvironmentSecretModel
	diags := model.Secrets.ElementsAs(ctx, &secrets, false)
	if diags.HasError() {
		return types.StringValue("")
	}
	for _, secret := range secrets {
		if secret.Name.Equal(types.StringValue(name)) && secret.Type.Equal(types.StringValue(varType)) {
			return secret.Value
		}
	}
	return types.StringValue("")
}

func (model projectEnvironmentModel) Secret(ctx context.Context, id types.Int64) *projectEnvironmentSecretModel {
	if model.Secrets.IsNull() || model.Secrets.IsUnknown() {
		return nil
	}

	var secrets []projectEnvironmentSecretModel
	diags := model.Secrets.ElementsAs(ctx, &secrets, false)
	if diags.HasError() {
		return nil
	}

	for _, secret := range secrets {
		if secret.ID.Equal(id) {
			return &secret
		}
	}
	return nil
}

func (r *projectEnvironmentResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = projectEnvironmentModel{}.GetSchema()
}

func convertProjectEnvironmentModelToEnvironmentRequest(ctx context.Context, env projectEnvironmentModel, prev *projectEnvironmentModel) *models.EnvironmentRequest {
	model := models.EnvironmentRequest{
		ProjectID: env.ProjectID.ValueInt64(),
		Name:      env.Name.ValueString(),
	}
	if !env.ID.IsNull() && !env.ID.IsUnknown() {
		model.ID = env.ID.ValueInt64()
	}

	if env.Variables == nil {
		model.JSON = "{}"
	} else {
		bytes, _ := json.Marshal(env.Variables)
		model.JSON = string(bytes)
	}

	if env.Environment == nil {
		model.Env = "{}"
	} else {
		bytes, _ := json.Marshal(env.Environment)
		model.Env = string(bytes)
	}

	var secrets []*models.EnvironmentSecretRequest
	var envSecrets, prevSecrets []projectEnvironmentSecretModel
	if env.Secrets.IsNull() || env.Secrets.IsUnknown() {
		envSecrets = []projectEnvironmentSecretModel{}
	} else {
		env.Secrets.ElementsAs(ctx, &envSecrets, false)
	}
	if prev.Secrets.IsUnknown() || prev.Secrets.IsNull() {
		prevSecrets = []projectEnvironmentSecretModel{}
	} else {
		prev.Secrets.ElementsAs(ctx, &prevSecrets, false)
	}

	for _, secret := range envSecrets {
		modelSecret := models.EnvironmentSecretRequest{
			Name: secret.Name.ValueString(),
			Type: secret.Type.ValueString(),
		}
		// Create all secrets from env missing an ID
		if secret.ID.IsUnknown() || secret.ID.IsNull() {
			modelSecret.Operation = "create"
			modelSecret.Secret = secret.Value.ValueString()
		} else {
			modelSecret.ID = secret.ID.ValueInt64()
			// Find the previous secret
			prevSecret := prev.Secret(ctx, secret.ID)
			if prevSecret != nil {
				// Update if any field has changed
				if !secret.Name.Equal(prevSecret.Name) || !secret.Value.Equal(prevSecret.Value) || !secret.Type.Equal(prevSecret.Type) {
					modelSecret.Operation = "update"
					if !secret.Name.Equal(prevSecret.Name) {
						modelSecret.Name = secret.Name.ValueString()
					}
				}
			}
		}
		secrets = append(secrets, &modelSecret)
	}

	// Delete all secrets from prev with an ID missing from env
	for _, prevSecret := range prevSecrets {
		secret := env.Secret(ctx, prevSecret.ID)
		if secret == nil {
			secrets = append(secrets, &models.EnvironmentSecretRequest{
				ID: prevSecret.ID.ValueInt64(),
				// Can't delete a secret without sending the Type
				Type:      prevSecret.Type.ValueString(),
				Operation: "delete",
			})
		}
	}

	model.Secrets = secrets

	return &model
}

var _ sort.Interface = ByEnvironmentID{}

type ByEnvironmentID []*models.EnvironmentSecret

func (a ByEnvironmentID) Len() int           { return len(a) }
func (a ByEnvironmentID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByEnvironmentID) Less(i, j int) bool { return a[i].ID < a[j].ID }

func convertEnvironmentResponseToProjectEnvironmentModel(ctx context.Context, environment *models.Environment, prev *projectEnvironmentModel) projectEnvironmentModel {
	model := projectEnvironmentModel{
		ID:        types.Int64Value(environment.ID),
		ProjectID: types.Int64Value(environment.ProjectID),
		Name:      types.StringValue(environment.Name),
	}

	if json.Unmarshal([]byte(environment.JSON), &model.Variables) != nil {
		model.Variables = &map[string]string{}
	}
	if len(*model.Variables) == 0 && prev.Variables == nil {
		model.Variables = nil
	}

	if json.Unmarshal([]byte(environment.Env), &model.Environment) != nil {
		model.Environment = &map[string]string{}
	}
	if len(*model.Environment) == 0 && prev.Environment == nil {
		model.Environment = nil
	}

	sort.Sort(ByEnvironmentID(environment.Secrets))

	var secrets []projectEnvironmentSecretModel
	for _, secret := range environment.Secrets {
		modelSecret := projectEnvironmentSecretModel{
			ID:   types.Int64Value(secret.ID),
			Type: types.StringValue(secret.Type),
			Name: types.StringValue(secret.Name),
		}
		// Value from previous state since secrets are not returned in the response
		prevSecret := prev.Secret(ctx, modelSecret.ID)
		if prevSecret != nil {
			modelSecret.Value = prevSecret.Value
		} else {
			modelSecret.Value = prev.SecretValue(ctx, secret.Name, secret.Type)
		}
		secrets = append(secrets, modelSecret)
	}
	if len(secrets) == 0 && !prev.Secrets.IsNull() && !prev.Secrets.IsUnknown() {
		prev.Secrets.ElementsAs(ctx, &secrets, false)
	}

	envSecrets, _ := types.ListValueFrom(ctx, projectEnvironmentSecretModel{}.GetSchema().Type(), secrets)

	model.Secrets = envSecrets

	return model
}

func (r *projectEnvironmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan projectEnvironmentModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	//Create new projectEnvironment
	response, err := r.client.Project.PostProjectProjectIDEnvironment(&project.PostProjectProjectIDEnvironmentParams{
		ProjectID:   plan.ProjectID.ValueInt64(),
		Environment: convertProjectEnvironmentModelToEnvironmentRequest(ctx, plan, &projectEnvironmentModel{}),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating SemaphoreUI Project Environment",
			"Could not create project environment, unexpected error: "+err.Error(),
		)
		return
	}

	payload, err := r.client.Project.GetProjectProjectIDEnvironmentEnvironmentID(&project.GetProjectProjectIDEnvironmentEnvironmentIDParams{
		ProjectID:     response.Payload.ProjectID,
		EnvironmentID: response.Payload.ID,
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Project Environment",
			"Could not read project environment, unexpected error: "+err.Error(),
		)
		return
	}
	plan = convertEnvironmentResponseToProjectEnvironmentModel(ctx, payload.Payload, &plan)

	// Set state to fully populated data
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *projectEnvironmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state projectEnvironmentModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := r.client.Project.GetProjectProjectIDEnvironmentEnvironmentID(&project.GetProjectProjectIDEnvironmentEnvironmentIDParams{
		ProjectID:     state.ProjectID.ValueInt64(),
		EnvironmentID: state.ID.ValueInt64(),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Project Environment",
			"Could not read project environment, unexpected error: "+err.Error(),
		)
		return
	}
	model := convertEnvironmentResponseToProjectEnvironmentModel(ctx, response.Payload, &state)

	// Set refreshed state
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *projectEnvironmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan, state projectEnvironmentModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.Project.PutProjectProjectIDEnvironmentEnvironmentID(&project.PutProjectProjectIDEnvironmentEnvironmentIDParams{
		ProjectID:     plan.ProjectID.ValueInt64(),
		EnvironmentID: plan.ID.ValueInt64(),
		Environment:   convertProjectEnvironmentModelToEnvironmentRequest(ctx, plan, &state),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating SemaphoreUI Project Key",
			"Could not update project key, unexpected error: "+err.Error(),
		)
		return
	}

	response, err := r.client.Project.GetProjectProjectIDEnvironmentEnvironmentID(&project.GetProjectProjectIDEnvironmentEnvironmentIDParams{
		ProjectID:     plan.ProjectID.ValueInt64(),
		EnvironmentID: plan.ID.ValueInt64(),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Project Environment",
			"Could not read project environment, unexpected error: "+err.Error(),
		)
		return
	}
	model := convertEnvironmentResponseToProjectEnvironmentModel(ctx, response.Payload, &plan)

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *projectEnvironmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state projectEnvironmentModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing resource
	_, err := r.client.Project.DeleteProjectProjectIDEnvironmentEnvironmentID(&project.DeleteProjectProjectIDEnvironmentEnvironmentIDParams{
		ProjectID:     state.ProjectID.ValueInt64(),
		EnvironmentID: state.ID.ValueInt64(),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Semaphore Project Environment",
			fmt.Sprintf("Could not delete project environment, unexpected error: %s", err.Error()),
		)
		return
	}
}

func (r *projectEnvironmentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	fields, err := parseImportFields(req.ID, []string{"project", "environment"})
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid Project Environment Import ID",
			"Could not parse import ID: "+err.Error(),
		)
		return
	}

	response, err := r.client.Project.GetProjectProjectIDEnvironmentEnvironmentID(&project.GetProjectProjectIDEnvironmentEnvironmentIDParams{
		ProjectID:     fields["project"],
		EnvironmentID: fields["environment"],
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Project Environment",
			"Could not read project environment, unexpected error: "+err.Error(),
		)
		return
	}
	model := convertEnvironmentResponseToProjectEnvironmentModel(ctx, response.Payload, &projectEnvironmentModel{})

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
