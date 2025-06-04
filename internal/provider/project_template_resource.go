package provider

import (
	"context"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"sort"
	apiclient "terraform-provider-semaphoreui/semaphoreui/client"
	"terraform-provider-semaphoreui/semaphoreui/client/project"
	"terraform-provider-semaphoreui/semaphoreui/models"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &projectTemplateResource{}
	_ resource.ResourceWithConfigure   = &projectTemplateResource{}
	_ resource.ResourceWithImportState = &projectTemplateResource{}
)

func NewProjectTemplateResource() resource.Resource {
	return &projectTemplateResource{}
}

type projectTemplateResource struct {
	client *apiclient.SemaphoreUI
}

func (r *projectTemplateResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *projectTemplateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project_template"
}

func (r *projectTemplateResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ProjectTemplateSchema().GetResource(ctx)
}

func convertProjectTemplateModelToTemplateRequest(ctx context.Context, template ProjectTemplateModel) *models.TemplateRequest {
	model := models.TemplateRequest{
		ProjectID:               template.ProjectID.ValueInt64(),
		EnvironmentID:           template.EnvironmentID.ValueInt64(),
		InventoryID:             template.InventoryID.ValueInt64(),
		RepositoryID:            template.RepositoryID.ValueInt64(),
		App:                     template.App.ValueString(),
		Name:                    template.Name.ValueString(),
		Playbook:                template.Playbook.ValueString(),
		AllowOverrideArgsInTask: template.AllowOverrideArgsInTask.ValueBool(),
		SuppressSuccessAlerts:   template.SuppressSuccessAlerts.ValueBool(),
	}
	if !template.ID.IsNull() && !template.ID.IsUnknown() {
		model.ID = template.ID.ValueInt64()
	}

	if !template.Description.IsNull() && !template.Description.IsUnknown() {
		model.Description = template.Description.ValueString()
	}
	if !template.GitBranch.IsNull() && !template.GitBranch.IsUnknown() {
		model.GitBranch = template.GitBranch.ValueString()
	}
	if !template.ViewID.IsNull() && !template.ViewID.IsUnknown() {
		model.ViewID = template.ViewID.ValueInt64()
	}

	if len(template.Arguments.Elements()) != 0 {
		var arguments []string
		template.Arguments.ElementsAs(ctx, &arguments, false)
		bytes, _ := json.Marshal(arguments)
		model.Arguments = string(bytes)
	} else {
		model.Arguments = "[]"
	}

	if template.Build != nil {
		model.Type = "build"
		if !template.Build.StartVersion.IsNull() && !template.Build.StartVersion.IsUnknown() {
			model.StartVersion = template.Build.StartVersion.ValueString()
		}
	}

	if template.Deploy != nil {
		model.Type = "deploy"
		model.BuildTemplateID = template.Deploy.BuildTemplateID.ValueInt64()
		model.Autorun = template.Deploy.Autorun.ValueBool()
	}

	model.SurveyVars = []*models.TemplateSurveyVar{}
	if !template.SurveyVars.IsNull() && !template.SurveyVars.IsUnknown() {
		var surveyVars []ProjectTemplateSurveyVarModel
		template.SurveyVars.ElementsAs(ctx, &surveyVars, false)
		for _, surveyVar := range surveyVars {
			surveyVarModel := models.TemplateSurveyVar{
				Name:     surveyVar.Name.ValueString(),
				Title:    surveyVar.Title.ValueString(),
				Required: surveyVar.Required.ValueBool(),
				Type:     surveyVar.Type.ValueString(),
			}
			if !surveyVar.Description.IsNull() && !surveyVar.Description.IsUnknown() {
				surveyVarModel.Description = surveyVar.Description.ValueString()
			}
			if surveyVar.Type.ValueString() == "enum" {
				for name, value := range surveyVar.EnumValues {
					surveyVarModel.Values = append(surveyVarModel.Values, &models.TemplateSurveyVarValue{
						Name:  name,
						Value: value,
					})
				}
			}
			model.SurveyVars = append(model.SurveyVars, &surveyVarModel)
		}
	}

	model.Vaults = []*models.TemplateVault{}
	if !template.Vaults.IsNull() || !template.Vaults.IsUnknown() {
		var vaults []ProjectTemplateVaultModel
		template.Vaults.ElementsAs(ctx, &vaults, false)
		for _, vault := range vaults {
			vaultModel := models.TemplateVault{
				Name: vault.Name.ValueString(),
			}
			if !vault.ID.IsNull() && !vault.ID.IsUnknown() {
				vaultModel.ID = vault.ID.ValueInt64()
			}
			if vault.Password != nil {
				vaultModel.Type = "password"
				vaultModel.VaultKeyID = vault.Password.VaultKeyID.ValueInt64()
			}
			if vault.ClientScript != nil {
				vaultModel.Type = "script"
				vaultModel.Script = vault.ClientScript.Script.ValueString()
			}
			model.Vaults = append(model.Vaults, &vaultModel)
		}
	}

	return &model
}

var _ sort.Interface = ByVaultID{}

type ByVaultID []*models.TemplateVault

func (a ByVaultID) Len() int           { return len(a) }
func (a ByVaultID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByVaultID) Less(i, j int) bool { return a[i].ID < a[j].ID }

func convertTemplateResponseToProjectTemplateModel(ctx context.Context, request *models.Template, prev *ProjectTemplateModel) ProjectTemplateModel {
	model := ProjectTemplateModel{
		ID:                      types.Int64Value(request.ID),
		ProjectID:               types.Int64Value(request.ProjectID),
		EnvironmentID:           types.Int64Value(request.EnvironmentID),
		InventoryID:             types.Int64Value(request.InventoryID),
		RepositoryID:            types.Int64Value(request.RepositoryID),
		App:                     types.StringValue(request.App),
		Name:                    types.StringValue(request.Name),
		Playbook:                types.StringValue(request.Playbook),
		AllowOverrideArgsInTask: types.BoolValue(request.AllowOverrideArgsInTask),
		SuppressSuccessAlerts:   types.BoolValue(request.SuppressSuccessAlerts),
	}

	if request.Description != "" {
		model.Description = types.StringValue(request.Description)
	} else {
		model.Description = prev.Description
	}

	if request.GitBranch != "" {
		model.GitBranch = types.StringValue(request.GitBranch)
	} else {
		model.GitBranch = prev.GitBranch
	}

	if request.ViewID != 0 {
		model.ViewID = types.Int64Value(request.ViewID)
	} else {
		model.ViewID = prev.ViewID
	}

	var arguments []string
	if json.Unmarshal([]byte(request.Arguments), &arguments) != nil {
		model.Arguments = types.ListNull(types.StringType)
	} else {
		if len(arguments) == 0 {
			model.Arguments = types.ListNull(types.StringType)
		} else {
			args, _ := types.ListValueFrom(ctx, types.StringType, arguments)
			model.Arguments = args
		}
	}

	if request.Type == "build" {
		build := ProjectTemplateTypeBuildModel{}
		if request.StartVersion != "" {
			build.StartVersion = types.StringValue(request.StartVersion)
		} else {
			if prev.Build != nil {
				build.StartVersion = prev.Build.StartVersion
			} else {
				build.StartVersion = types.StringNull()
			}
		}
		model.Build = &build
	}

	if request.Type == "deploy" {
		model.Deploy = &ProjectTemplateTypeDeployModel{
			BuildTemplateID: types.Int64Value(request.BuildTemplateID),
			Autorun:         types.BoolValue(request.Autorun),
		}
	}

	if len(request.SurveyVars) == 0 {
		model.SurveyVars = prev.SurveyVars
	} else {
		var surveyVars []ProjectTemplateSurveyVarModel
		for _, surveyVar := range request.SurveyVars {
			surveyVarModel := ProjectTemplateSurveyVarModel{
				Name:     types.StringValue(surveyVar.Name),
				Title:    types.StringValue(surveyVar.Title),
				Required: types.BoolValue(surveyVar.Required),
				Type:     types.StringValue(surveyVar.Type),
			}
			if surveyVar.Description != "" {
				surveyVarModel.Description = types.StringValue(surveyVar.Description)
			}
			if surveyVar.Type == "enum" {
				enumValuesMap := map[string]string{}
				for _, value := range surveyVar.Values {
					enumValuesMap[value.Name] = value.Value
				}
				surveyVarModel.EnumValues = enumValuesMap
			}
			surveyVars = append(surveyVars, surveyVarModel)
		}
		surveyVarsModel, _ := types.ListValueFrom(ctx, ProjectTemplateSurveyVarType, &surveyVars)
		model.SurveyVars = surveyVarsModel
	}

	if len(request.Vaults) == 0 {
		model.Vaults = prev.Vaults
	} else {
		sort.Sort(ByVaultID(request.Vaults))

		var vaults []ProjectTemplateVaultModel
		for _, vault := range request.Vaults {
			vaultModel := ProjectTemplateVaultModel{
				ID:   types.Int64Value(vault.ID),
				Name: types.StringValue(vault.Name),
			}
			if vault.Type == "password" {
				vaultModel.Password = &ProjectTemplateVaultPasswordModel{
					VaultKeyID: types.Int64Value(vault.VaultKeyID),
				}
			}
			if vault.Type == "script" {
				vaultModel.ClientScript = &ProjectTemplateVaultScriptModel{
					Script: types.StringValue(vault.Script),
				}
			}
			vaults = append(vaults, vaultModel)
		}
		vaultsModel, _ := types.ListValueFrom(ctx, ProjectTemplateVaultType, &vaults)
		model.Vaults = vaultsModel
	}

	return model
}

func (r *projectTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan ProjectTemplateModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	create, err := r.client.Project.PostProjectProjectIDTemplates(&project.PostProjectProjectIDTemplatesParams{
		ProjectID: plan.ProjectID.ValueInt64(),
		Template:  convertProjectTemplateModelToTemplateRequest(ctx, plan),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating SemaphoreUI Project Template",
			"Could not create project template, unexpected error: "+err.Error(),
		)
		return
	}

	// Create response doesn't fully capture the model, so we need to read it back
	response, err := r.client.Project.GetProjectProjectIDTemplatesTemplateID(&project.GetProjectProjectIDTemplatesTemplateIDParams{
		ProjectID:  plan.ProjectID.ValueInt64(),
		TemplateID: create.Payload.ID,
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Project Template",
			"Could not read project template, unexpected error: "+err.Error(),
		)
		return
	}
	model := convertTemplateResponseToProjectTemplateModel(ctx, response.Payload, &plan)

	// Set state to fully populated data
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *projectTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state ProjectTemplateModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := r.client.Project.GetProjectProjectIDTemplatesTemplateID(&project.GetProjectProjectIDTemplatesTemplateIDParams{
		ProjectID:  state.ProjectID.ValueInt64(),
		TemplateID: state.ID.ValueInt64(),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Project Template",
			"Could not read project template, unexpected error: "+err.Error(),
		)
		return
	}
	model := convertTemplateResponseToProjectTemplateModel(ctx, response.Payload, &state)

	// Set refreshed state
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *projectTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan ProjectTemplateModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.Project.PutProjectProjectIDTemplatesTemplateID(&project.PutProjectProjectIDTemplatesTemplateIDParams{
		ProjectID:  plan.ProjectID.ValueInt64(),
		TemplateID: plan.ID.ValueInt64(),
		Template:   convertProjectTemplateModelToTemplateRequest(ctx, plan),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating SemaphoreUI Project Template",
			"Could not update project template, unexpected error: "+err.Error(),
		)
		return
	}

	response, err := r.client.Project.GetProjectProjectIDTemplatesTemplateID(&project.GetProjectProjectIDTemplatesTemplateIDParams{
		ProjectID:  plan.ProjectID.ValueInt64(),
		TemplateID: plan.ID.ValueInt64(),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Project Template",
			"Could not read project template, unexpected error: "+err.Error(),
		)
		return
	}
	model := convertTemplateResponseToProjectTemplateModel(ctx, response.Payload, &plan)

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *projectTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state ProjectTemplateModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.Project.DeleteProjectProjectIDTemplatesTemplateID(&project.DeleteProjectProjectIDTemplatesTemplateIDParams{
		ProjectID:  state.ProjectID.ValueInt64(),
		TemplateID: state.ID.ValueInt64(),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Removing SemaphoreUI Project Template",
			"Could not delete project template, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *projectTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	fields, err := parseImportFields(req.ID, []string{"project", "template"})
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid Project Template Import ID",
			"Could not parse import ID: "+err.Error(),
		)
		return
	}

	response, err := r.client.Project.GetProjectProjectIDTemplatesTemplateID(&project.GetProjectProjectIDTemplatesTemplateIDParams{
		ProjectID:  fields["project"],
		TemplateID: fields["template"],
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Project Template",
			"Could not read project template, unexpected error: "+err.Error(),
		)
		return
	}
	model := convertTemplateResponseToProjectTemplateModel(ctx, response.Payload, &ProjectTemplateModel{
		SurveyVars: types.ListNull(ProjectTemplateSurveyVarType),
		Vaults:     types.ListNull(ProjectTemplateVaultType),
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
