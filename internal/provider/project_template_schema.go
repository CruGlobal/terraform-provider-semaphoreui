package provider

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	schemaD "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schemaR "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	superschema "github.com/orange-cloudavenue/terraform-plugin-framework-superschema"
	"regexp"
)

type (
	ProjectTemplateModel struct {
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

		Build  *ProjectTemplateTypeBuildModel  `tfsdk:"build"`
		Deploy *ProjectTemplateTypeDeployModel `tfsdk:"deploy"`
	}

	ProjectTemplateTypeBuildModel struct {
		StartVersion types.String `tfsdk:"start_version"`
	}

	ProjectTemplateTypeDeployModel struct {
		BuildTemplateID types.Int64 `tfsdk:"build_template_id"`
		Autorun         types.Bool  `tfsdk:"autorun"`
	}

	ProjectTemplateSurveyVarModel struct {
		Name        types.String      `tfsdk:"name"`
		Title       types.String      `tfsdk:"title"`
		Description types.String      `tfsdk:"description"`
		Required    types.Bool        `tfsdk:"required"`
		Type        types.String      `tfsdk:"type"`
		EnumValues  map[string]string `tfsdk:"enum_values"`
	}

	ProjectTemplateVaultModel struct {
		ID           types.Int64                        `tfsdk:"id"`
		Name         types.String                       `tfsdk:"name"`
		Password     *ProjectTemplateVaultPasswordModel `tfsdk:"password"`
		ClientScript *ProjectTemplateVaultScriptModel   `tfsdk:"client_script"`
	}

	ProjectTemplateVaultPasswordModel struct {
		VaultKeyID types.Int64 `tfsdk:"vault_key_id"`
	}

	ProjectTemplateVaultScriptModel struct {
		Script types.String `tfsdk:"script"`
	}
)

var (
	ProjectTemplateSurveyVarType = types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":        types.StringType,
			"title":       types.StringType,
			"description": types.StringType,
			"required":    types.BoolType,
			"type":        types.StringType,
			"enum_values": types.MapType{
				ElemType: types.StringType,
			},
		},
	}

	ProjectTemplateVaultType = types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":   types.Int64Type,
			"name": types.StringType,
			"password": types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"vault_key_id": types.Int64Type,
				},
			},
			"client_script": types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"script": types.StringType,
				},
			},
		},
	}
)

func ProjectTemplateSchema() superschema.Schema {
	return superschema.Schema{
		Common: superschema.SchemaDetails{
			MarkdownDescription: "The project template",
		},
		Resource: superschema.SchemaDetails{
			MarkdownDescription: "resource allows you to define a task template which tells SemaphoreUI how to run an application task.",
		},
		DataSource: superschema.SchemaDetails{
			MarkdownDescription: "data source allows you to read a template.",
		},
		Attributes: map[string]superschema.Attribute{
			"id": superschema.Int64Attribute{
				Common: &schemaR.Int64Attribute{
					MarkdownDescription: "The template ID.",
				},
				Resource: &schemaR.Int64Attribute{
					Computed:      true,
					PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
				},
				DataSource: &schemaD.Int64Attribute{
					Optional: true,
					Computed: true,
					Validators: []validator.Int64{
						int64validator.ExactlyOneOf(
							path.MatchRoot("id"),
							path.MatchRoot("name"),
						),
					},
				},
			},
			"project_id": superschema.Int64Attribute{
				Common: &schemaR.Int64Attribute{
					MarkdownDescription: "The project ID that the template belongs to.",
					Required:            true,
				},
				Resource: &schemaR.Int64Attribute{
					PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
				},
			},
			"environment_id": superschema.Int64Attribute{
				Common: &schemaR.Int64Attribute{
					MarkdownDescription: "The environment (variable group) ID that the template uses.",
				},
				Resource: &schemaR.Int64Attribute{
					Required: true,
				},
				DataSource: &schemaD.Int64Attribute{
					Computed: true,
				},
			},
			"inventory_id": superschema.Int64Attribute{
				Common: &schemaR.Int64Attribute{
					MarkdownDescription: "The inventory ID that the template uses.",
				},
				Resource: &schemaR.Int64Attribute{
					Required: true,
				},
				DataSource: &schemaD.Int64Attribute{
					Computed: true,
				},
			},
			"repository_id": superschema.Int64Attribute{
				Common: &schemaR.Int64Attribute{
					MarkdownDescription: "The repository ID that the template uses.",
				},
				Resource: &schemaR.Int64Attribute{
					Required: true,
				},
				DataSource: &schemaD.Int64Attribute{
					Computed: true,
				},
			},
			"name": superschema.StringAttribute{
				Common: &schemaR.StringAttribute{
					MarkdownDescription: "The display name of the template.",
				},
				Resource: &schemaR.StringAttribute{
					Required: true,
				},
				DataSource: &schemaD.StringAttribute{
					Optional: true,
					Computed: true,
					Validators: []validator.String{
						stringvalidator.ExactlyOneOf(
							path.MatchRoot("id"),
							path.MatchRoot("name"),
						),
					},
				},
			},
			"app": superschema.StringAttribute{
				Common: &schemaR.StringAttribute{
					MarkdownDescription: "The application name.",
				},
				Resource: &schemaR.StringAttribute{
					MarkdownDescription: "Must be a valid SemaphoreUI application name. Default applications include: `ansible`, `terraform`, `tofu`, `bash`, `powershell` and `python`.",
					Optional:            true,
					Computed:            true,
					Default:             stringdefault.StaticString("ansible"),
				},
				DataSource: &schemaD.StringAttribute{
					Computed: true,
				},
			},
			"playbook": superschema.StringAttribute{
				Common: &schemaR.StringAttribute{
					MarkdownDescription: "The playbook/script filename.",
				},
				Resource: &schemaR.StringAttribute{
					Required: true,
					Validators: []validator.String{
						// Only relative paths are allowed
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^[^/].*$`),
							"must be a relative path (path/to/inventory)",
						),
					},
				},
				DataSource: &schemaD.StringAttribute{
					Computed: true,
				},
			},
			"description": superschema.StringAttribute{
				Common: &schemaR.StringAttribute{
					MarkdownDescription: "The description of the template.",
				},
				Resource: &schemaR.StringAttribute{
					Optional: true,
				},
				DataSource: &schemaD.StringAttribute{
					Computed: true,
				},
			},
			"git_branch": superschema.StringAttribute{
				Common: &schemaR.StringAttribute{
					MarkdownDescription: "Override the git branch defined in the project repository.",
				},
				Resource: &schemaR.StringAttribute{
					Optional: true,
				},
				DataSource: &schemaD.StringAttribute{
					Computed: true,
				},
			},
			"allow_override_args_in_task": superschema.BoolAttribute{
				Common: &schemaR.BoolAttribute{
					MarkdownDescription: "Allow overriding arguments in the task.",
				},
				Resource: &schemaR.BoolAttribute{
					Optional: true,
					Computed: true,
					Default:  booldefault.StaticBool(false),
				},
				DataSource: &schemaD.BoolAttribute{
					Computed: true,
				},
			},
			"suppress_success_alerts": superschema.BoolAttribute{
				Common: &schemaR.BoolAttribute{
					MarkdownDescription: "Suppress success alerts.",
				},
				Resource: &schemaR.BoolAttribute{
					Optional: true,
					Computed: true,
					Default:  booldefault.StaticBool(false),
				},
				DataSource: &schemaD.BoolAttribute{
					Computed: true,
				},
			},
			"arguments": superschema.ListAttribute{
				Common: &schemaR.ListAttribute{
					MarkdownDescription: "Commandline arguments passed to the application.",
					ElementType:         types.StringType,
				},
				Resource: &schemaR.ListAttribute{
					Optional: true,
				},
				DataSource: &schemaD.ListAttribute{
					Computed: true,
				},
			},
			"build": superschema.SingleNestedAttribute{
				Common: &schemaR.SingleNestedAttribute{
					MarkdownDescription: "Specifies a build type template used to create artifacts.",
				},
				Resource: &schemaR.SingleNestedAttribute{
					MarkdownDescription: "SemaphoreUI doesn't support artifacts out-of-box, it only provides task versioning. You should implement the artifact creation yourself.",
					Optional:            true,
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(path.MatchRoot("deploy")),
					},
				},
				DataSource: &schemaD.SingleNestedAttribute{
					Computed: true,
				},
				Attributes: map[string]superschema.Attribute{
					"start_version": superschema.StringAttribute{
						Common: &schemaR.StringAttribute{
							MarkdownDescription: "Defines start version of your artifact.",
						},
						Resource: &schemaR.StringAttribute{
							MarkdownDescription: "Each run increments the artifact version.",
							Optional:            true,
						},
						DataSource: &schemaD.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"deploy": superschema.SingleNestedAttribute{
				Common: &schemaR.SingleNestedAttribute{
					MarkdownDescription: "Specifies a deploy type template used to deploy artifacts. Each `deploy` template is associated with a build template.",
				},
				Resource: &schemaR.SingleNestedAttribute{
					Optional: true,
					Validators: []validator.Object{
						objectvalidator.ConflictsWith(path.MatchRoot("build")),
					},
				},
				DataSource: &schemaD.SingleNestedAttribute{
					Computed: true,
				},
				Attributes: map[string]superschema.Attribute{
					"build_template_id": superschema.Int64Attribute{
						Common: &schemaR.Int64Attribute{
							MarkdownDescription: "The ID of the build template.",
						},
						Resource: &schemaR.Int64Attribute{
							Required: true,
						},
						DataSource: &schemaD.Int64Attribute{
							Computed: true,
						},
					},
					"autorun": superschema.BoolAttribute{
						Common: &schemaR.BoolAttribute{
							MarkdownDescription: "Automatically run the deploy template after the build template.",
						},
						Resource: &schemaR.BoolAttribute{
							Optional: true,
							Computed: true,
							Default:  booldefault.StaticBool(false),
						},
						DataSource: &schemaD.BoolAttribute{
							Computed: true,
						},
					},
				},
			},
			"survey_vars": superschema.ListNestedAttribute{
				Common: &schemaR.ListNestedAttribute{
					MarkdownDescription: "Survey variables.",
				},
				Resource: &schemaR.ListNestedAttribute{
					Optional: true,
				},
				DataSource: &schemaD.ListNestedAttribute{
					Computed: true,
				},
				Attributes: map[string]superschema.Attribute{
					"name": superschema.StringAttribute{
						Common: &schemaR.StringAttribute{
							MarkdownDescription: "The name of the survey variable.",
						},
						Resource: &schemaR.StringAttribute{
							Required: true,
						},
						DataSource: &schemaD.StringAttribute{
							Computed: true,
						},
					},
					"title": superschema.StringAttribute{
						Common: &schemaR.StringAttribute{
							MarkdownDescription: "The title of the survey variable.",
						},
						Resource: &schemaR.StringAttribute{
							Required: true,
						},
						DataSource: &schemaD.StringAttribute{
							Computed: true,
						},
					},
					"description": superschema.StringAttribute{
						Common: &schemaR.StringAttribute{
							MarkdownDescription: "The description of the survey variable.",
						},
						Resource: &schemaR.StringAttribute{
							Optional: true,
						},
						DataSource: &schemaD.StringAttribute{
							Computed: true,
						},
					},
					"required": superschema.BoolAttribute{
						Common: &schemaR.BoolAttribute{
							MarkdownDescription: "Whether the survey variable is required.",
						},
						Resource: &schemaR.BoolAttribute{
							Optional: true,
							Computed: true,
							Default:  booldefault.StaticBool(false),
						},
						DataSource: &schemaD.BoolAttribute{
							Computed: true,
						},
					},
					"type": superschema.StringAttribute{
						Common: &schemaR.StringAttribute{
							MarkdownDescription: "The type of the survey variable.",
						},
						Resource: &schemaR.StringAttribute{
							MarkdownDescription: "Valid types are `string`, `integer`, `secret` and `enum`. When `enum` is used, the `enum_values` attribute must be defined.",
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
						DataSource: &schemaD.StringAttribute{
							Computed: true,
						},
					},
					"enum_values": superschema.MapAttribute{
						Common: &schemaR.MapAttribute{
							MarkdownDescription: "The enum name/values.",
							ElementType:         types.StringType,
						},
						Resource: &schemaR.MapAttribute{
							Optional: true,
							Validators: []validator.Map{
								mapvalidator.SizeAtLeast(1),
								mapvalidator.AlsoRequires(path.Expressions{
									path.MatchRelative().AtParent().AtName("type"),
								}...),
							},
						},
						DataSource: &schemaD.MapAttribute{
							Computed: true,
						},
					},
				},
			},
			"vaults": superschema.ListNestedAttribute{
				Common: &schemaR.ListNestedAttribute{
					MarkdownDescription: "Ansible Vault Passwords.",
				},
				Resource: &schemaR.ListNestedAttribute{
					Optional: true,
				},
				DataSource: &schemaD.ListNestedAttribute{
					Computed: true,
				},
				Attributes: map[string]superschema.Attribute{
					"id": superschema.Int64Attribute{
						Common: &schemaR.Int64Attribute{
							MarkdownDescription: "The vault ID.",
							Computed:            true,
						},
						Resource: &schemaR.Int64Attribute{
							PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
						},
					},
					"name": superschema.StringAttribute{
						Common: &schemaR.StringAttribute{
							MarkdownDescription: "Ansible vault ID name. Must be unique.",
						},
						Resource: &schemaR.StringAttribute{
							Required: true,
						},
						DataSource: &schemaD.StringAttribute{
							Computed: true,
						},
					},
					"password": superschema.SingleNestedAttribute{
						Common: &schemaR.SingleNestedAttribute{
							MarkdownDescription: "Unlock vault using a password.",
						},
						Resource: &schemaR.SingleNestedAttribute{
							Optional: true,
							Validators: []validator.Object{
								objectvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("client_script")),
							},
						},
						DataSource: &schemaD.SingleNestedAttribute{
							Computed: true,
						},
						Attributes: map[string]superschema.Attribute{
							"vault_key_id": superschema.Int64Attribute{
								Common: &schemaR.Int64Attribute{
									MarkdownDescription: "The project key ID to use.",
								},
								Resource: &schemaR.Int64Attribute{
									Required: true,
								},
								DataSource: &schemaD.Int64Attribute{
									Computed: true,
								},
							},
						},
					},
					"client_script": superschema.SingleNestedAttribute{
						Common: &schemaR.SingleNestedAttribute{
							MarkdownDescription: "Unlock vault using an ansible vault password client script.",
						},
						Resource: &schemaR.SingleNestedAttribute{
							Optional: true,
							Validators: []validator.Object{
								objectvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("password")),
							},
						},
						DataSource: &schemaD.SingleNestedAttribute{
							Computed: true,
						},
						Attributes: map[string]superschema.Attribute{
							"script": superschema.StringAttribute{
								Common: &schemaR.StringAttribute{
									MarkdownDescription: "The script path.",
								},
								Resource: &schemaR.StringAttribute{
									MarkdownDescription: "Must end in `-client` with extension. See [Ansible Vault Password Client](https://docs.ansible.com/ansible/latest/vault_guide/vault_managing_passwords.html#storing-passwords-in-third-party-tools-with-vault-password-client-scripts).",
									Required:            true,
								},
								DataSource: &schemaD.StringAttribute{
									Computed: true,
								},
							},
						},
					},
				},
			},
		},
	}
}
