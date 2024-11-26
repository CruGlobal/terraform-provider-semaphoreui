package provider

import (
	"context"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"regexp"
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

type projectTemplateModel struct {
	ID            types.Int64 `tfsdk:"id"`
	ProjectID     types.Int64 `tfsdk:"project_id"`
	EnvironmentID types.Int64 `tfsdk:"environment_id"`
	InventoryID   types.Int64 `tfsdk:"inventory_id"`
	RepositoryID  types.Int64 `tfsdk:"repository_id"`
	//ViewID        types.Int64 `tfsdk:"view_id"`

	Name                    types.String `tfsdk:"name"`
	Description             types.String `tfsdk:"description"`
	App                     types.String `tfsdk:"app"`
	AllowOverrideArgsInTask types.Bool   `tfsdk:"allow_override_args_in_task"`
	Arguments               types.List   `tfsdk:"arguments"`
	GitBranch               types.String `tfsdk:"git_branch"`
	Playbook                types.String `tfsdk:"playbook"`
	SuppressSuccessAlerts   types.Bool   `tfsdk:"suppress_success_alerts"`
	SurveyVars              types.List   `tfsdk:"survey_vars"`
	Vaults                  types.List   `tfsdk:"vaults"`

	Build  *projectTemplateTypeBuildModel  `tfsdk:"build"`
	Deploy *projectTemplateTypeDeployModel `tfsdk:"deploy"`
}

type projectTemplateTypeBuildModel struct {
	StartVersion types.String `tfsdk:"start_version"`
}

type projectTemplateTypeDeployModel struct {
	BuildTemplateID types.Int64 `tfsdk:"build_template_id"`
	Autorun         types.Bool  `tfsdk:"autorun"`
}

type projectTemplateSurveyVarModel struct {
	Name        types.String      `tfsdk:"name"`
	Title       types.String      `tfsdk:"title"`
	Description types.String      `tfsdk:"description"`
	Required    types.Bool        `tfsdk:"required"`
	Type        types.String      `tfsdk:"type"`
	EnumValues  map[string]string `tfsdk:"enum_values"`
}

type projectTemplateVaultModel struct {
	ID           types.Int64                        `tfsdk:"id"`
	Name         types.String                       `tfsdk:"name"`
	Password     *projectTemplateVaultPasswordModel `tfsdk:"password"`
	ClientScript *projectTemplateVaultScriptModel   `tfsdk:"client_script"`
}

type projectTemplateVaultPasswordModel struct {
	VaultKeyID types.Int64 `tfsdk:"vault_key_id"`
}

type projectTemplateVaultScriptModel struct {
	Script types.String `tfsdk:"script"`
}

func (r *projectTemplateResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: `Provides a SemaphoreUI ProjectTemplate resource.`,
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				MarkdownDescription: "The project template ID.",
				Computed:            true,
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"project_id": schema.Int64Attribute{
				MarkdownDescription: "The project ID.",
				Required:            true,
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"environment_id": schema.Int64Attribute{
				MarkdownDescription: "The project environment ID.",
				Required:            true,
			},
			"inventory_id": schema.Int64Attribute{
				MarkdownDescription: "The project inventory ID.",
				Required:            true,
			},
			"repository_id": schema.Int64Attribute{
				MarkdownDescription: "The project repository ID.",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The display name of the template.",
				Required:            true,
			},
			"app": schema.StringAttribute{
				MarkdownDescription: "The application name. Must be a valid SemaphoreUI application name. Default names include: `ansible`, `terraform`, `tofu`, `bash`, `powershell` and `python`.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("ansible"),
			},
			"playbook": schema.StringAttribute{
				MarkdownDescription: "The playbook/script filename.",
				Required:            true,
				Validators: []validator.String{
					// Only relative paths are allowed
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[^/].*$`),
						"must be a relative path (path/to/inventory)",
					),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the template.",
				Optional:            true,
			},
			"git_branch": schema.StringAttribute{
				MarkdownDescription: "Override the git branch defined in the project repository.",
				Optional:            true,
			},
			"allow_override_args_in_task": schema.BoolAttribute{
				MarkdownDescription: "Allow overriding arguments in the task.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"suppress_success_alerts": schema.BoolAttribute{
				MarkdownDescription: "Suppress success alerts.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"arguments": schema.ListAttribute{
				MarkdownDescription: "Commandline arguments passed to the application.",
				Optional:            true,
				ElementType:         types.StringType,
			},
			"build": schema.SingleNestedAttribute{
				MarkdownDescription: "Specifies a build type template used to create artifacts. SemaphoreUI doesn't support artifacts out-of-box, it only provides task versioning. You should implement the artifact creation yourself.",
				Optional:            true,
				Validators: []validator.Object{
					objectvalidator.ConflictsWith(path.MatchRoot("deploy")),
				},
				Attributes: map[string]schema.Attribute{
					"start_version": schema.StringAttribute{
						MarkdownDescription: "Defines start version of your artifact. Each run increments the artifact version.",
						Optional:            true,
					},
				},
			},
			"deploy": schema.SingleNestedAttribute{
				MarkdownDescription: "Specifies a deploy type template used to deploy artifacts. Each `deploy` template is associated with a build template.",
				Optional:            true,
				Validators: []validator.Object{
					objectvalidator.ConflictsWith(path.MatchRoot("build")),
				},
				Attributes: map[string]schema.Attribute{
					"build_template_id": schema.Int64Attribute{
						MarkdownDescription: "The ID of the build template.",
						Required:            true,
					},
					"autorun": schema.BoolAttribute{
						MarkdownDescription: "Automatically run the deploy template after the build template.",
						Optional:            true,
						Computed:            true,
						Default:             booldefault.StaticBool(false),
					},
				},
			},
			"survey_vars": schema.ListNestedAttribute{
				MarkdownDescription: "Survey variables.",
				Optional:            true,
				NestedObject:        projectTemplateSurveyVarModel{}.GetSchema(),
			},
			"vaults": schema.ListNestedAttribute{
				MarkdownDescription: "Ansible Vaults Passwords.",
				Optional:            true,
				NestedObject:        projectTemplateVaultModel{}.GetSchema(),
			},
		},
	}
}

func (projectTemplateSurveyVarModel) GetSchema() schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the survey variable.",
				Required:            true,
			},
			"title": schema.StringAttribute{
				MarkdownDescription: "The title of the survey variable.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the survey variable.",
				Optional:            true,
			},
			"required": schema.BoolAttribute{
				MarkdownDescription: "Whether the survey variable is required.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "The type of the survey variable. Valid types are `string`, `integer`, `secret` and `enum`. When `enum` is used, the `enum_values` attribute must be defined.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.Any(
						stringvalidator.OneOf("string", "integer", "secret"),
						stringvalidator.All(
							stringvalidator.OneOf("enum"),
							stringvalidator.AlsoRequires(path.Expressions{
								path.MatchRelative().AtParent().AtName("enum_values"),
							}...),
						),
					),
				},
			},
			"enum_values": schema.MapAttribute{
				MarkdownDescription: "The enum name/values.",
				Optional:            true,
				ElementType:         types.StringType,
				Validators: []validator.Map{
					mapvalidator.SizeAtLeast(1),
					mapvalidator.AlsoRequires(path.Expressions{
						path.MatchRelative().AtParent().AtName("type"),
					}...),
				},
			},
		},
	}
}

func (projectTemplateVaultModel) GetSchema() schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				MarkdownDescription: "The vault ID.",
				Computed:            true,
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Ansible vault ID name. Must be unique.",
				Required:            true,
			},
			"password": schema.SingleNestedAttribute{
				MarkdownDescription: "Unlock vault using a password.",
				Optional:            true,
				Validators: []validator.Object{
					objectvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("client_script")),
				},
				Attributes: map[string]schema.Attribute{
					"vault_key_id": schema.Int64Attribute{
						MarkdownDescription: "The project key ID to use.",
						Required:            true,
					},
				},
			},
			"client_script": schema.SingleNestedAttribute{
				MarkdownDescription: "Unlock vault using an ansible vault password client script.",
				Optional:            true,
				Attributes: map[string]schema.Attribute{
					"script": schema.StringAttribute{
						MarkdownDescription: "The script path. Must end in `-client` with extension. See [Ansible Vault Password Client](https://docs.ansible.com/ansible/latest/vault_guide/vault_managing_passwords.html#storing-passwords-in-third-party-tools-with-vault-password-client-scripts).",
						Required:            true,
					},
				},
			},
		},
	}
}

func convertProjectTemplateModelToTemplateRequest(ctx context.Context, template projectTemplateModel) *models.TemplateRequest {
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
		var surveyVars []projectTemplateSurveyVarModel
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
		var vaults []projectTemplateVaultModel
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

func convertTemplateResponseToProjectTemplateModel(ctx context.Context, request *models.Template, prev *projectTemplateModel) projectTemplateModel {
	model := projectTemplateModel{
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
		build := projectTemplateTypeBuildModel{}
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
		model.Deploy = &projectTemplateTypeDeployModel{
			BuildTemplateID: types.Int64Value(request.BuildTemplateID),
			Autorun:         types.BoolValue(request.Autorun),
		}
	}

	if len(request.SurveyVars) == 0 {
		model.SurveyVars = prev.SurveyVars
	} else {
		var surveyVars []projectTemplateSurveyVarModel
		for _, surveyVar := range request.SurveyVars {
			surveyVarModel := projectTemplateSurveyVarModel{
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
		surveyVarsModel, _ := types.ListValueFrom(ctx, projectTemplateSurveyVarModel{}.GetSchema().Type(), &surveyVars)
		model.SurveyVars = surveyVarsModel
	}

	if len(request.Vaults) == 0 {
		model.Vaults = prev.Vaults
	} else {
		sort.Sort(ByVaultID(request.Vaults))

		var vaults []projectTemplateVaultModel
		for _, vault := range request.Vaults {
			vaultModel := projectTemplateVaultModel{
				ID:   types.Int64Value(vault.ID),
				Name: types.StringValue(vault.Name),
			}
			if vault.Type == "password" {
				vaultModel.Password = &projectTemplateVaultPasswordModel{
					VaultKeyID: types.Int64Value(vault.VaultKeyID),
				}
			}
			if vault.Type == "script" {
				vaultModel.ClientScript = &projectTemplateVaultScriptModel{
					Script: types.StringValue(vault.Script),
				}
			}
			vaults = append(vaults, vaultModel)
		}
		typed := projectTemplateVaultModel{}.GetSchema().Type()
		vaultsModel, _ := types.ListValueFrom(ctx, typed, &vaults)
		model.Vaults = vaultsModel
	}

	return model
}

func (r *projectTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan projectTemplateModel
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
	var state projectTemplateModel
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
	var plan projectTemplateModel
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
	var state projectTemplateModel
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
	model := convertTemplateResponseToProjectTemplateModel(ctx, response.Payload, &projectTemplateModel{
		SurveyVars: types.ListNull(projectTemplateSurveyVarModel{}.GetSchema().Type()),
		Vaults:     types.ListNull(projectTemplateVaultModel{}.GetSchema().Type()),
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
