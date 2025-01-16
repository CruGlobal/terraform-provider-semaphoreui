package provider

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	schemaD "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schemaR "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	superschema "github.com/orange-cloudavenue/terraform-plugin-framework-superschema"
	"regexp"
)

type (
	ProjectInventoryModel struct {
		ID                 types.Int64                              `tfsdk:"id"`
		ProjectID          types.Int64                              `tfsdk:"project_id"`
		Name               types.String                             `tfsdk:"name"`
		SSHKeyID           types.Int64                              `tfsdk:"ssh_key_id"`
		Static             *ProjectInventoryStaticModel             `tfsdk:"static"`
		StaticYaml         *ProjectInventoryStaticYamlModel         `tfsdk:"static_yaml"`
		File               *ProjectInventoryFileModel               `tfsdk:"file"`
		TerraformWorkspace *ProjectInventoryTerraformWorkspaceModel `tfsdk:"terraform_workspace"`
	}

	ProjectInventoryStaticModel struct {
		Inventory   types.String `tfsdk:"inventory"`
		BecomeKeyID types.Int64  `tfsdk:"become_key_id"`
	}

	ProjectInventoryStaticYamlModel struct {
		Inventory   types.String `tfsdk:"inventory"`
		BecomeKeyID types.Int64  `tfsdk:"become_key_id"`
	}

	ProjectInventoryFileModel struct {
		Path         types.String `tfsdk:"path"`
		RepositoryID types.Int64  `tfsdk:"repository_id"`
		BecomeKeyID  types.Int64  `tfsdk:"become_key_id"`
	}

	ProjectInventoryTerraformWorkspaceModel struct {
		Workspace types.String `tfsdk:"workspace"`
	}
)

const (
	ProjectInventoryStatic             string = "static"
	ProjectInventoryStaticYaml         string = "static-yaml"
	ProjectInventoryFile               string = "file"
	ProjectInventoryTerraformWorkspace string = "terraform-workspace"
)

func ProjectInventorySchema() superschema.Schema {
	return superschema.Schema{
		Common: superschema.SchemaDetails{
			MarkdownDescription: "The project inventory",
		},
		Resource: superschema.SchemaDetails{
			MarkdownDescription: "resource allows you to define the Ansible inventory or a Terraform/OpenTofu workspace for a project.  Only one of the inventory types (`static`, `static_yaml`, `file` or `terraform_workspace`) can be defined per inventory.",
		},
		DataSource: superschema.SchemaDetails{
			MarkdownDescription: "data source allows you to read the Ansible inventory or a Terraform/OpenTofu workspace for a project.",
		},
		Attributes: map[string]superschema.Attribute{
			"id": superschema.Int64Attribute{
				Common: &schemaR.Int64Attribute{
					MarkdownDescription: "The inventory ID.",
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
					MarkdownDescription: "The project ID that the inventory belongs to.",
					Required:            true,
				},
				Resource: &schemaR.Int64Attribute{
					PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
				},
			},
			"name": superschema.StringAttribute{
				Common: &schemaR.StringAttribute{
					MarkdownDescription: "The display name of the inventory or workspace.",
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
			"ssh_key_id": superschema.Int64Attribute{
				Common: &schemaR.Int64Attribute{
					MarkdownDescription: "The Project Key ID to use for accessing hosts in the inventory. This attribute is required for all inventory types in SemaphoreUI. You should set it to the ID of a Key of type `none` if the inventory doesn't require credentials, or for Workspace type inventories.",
				},
				Resource: &schemaR.Int64Attribute{
					Required: true,
				},
				DataSource: &schemaD.Int64Attribute{
					Computed: true,
				},
			},
			"static": superschema.SingleNestedAttribute{
				Common: &schemaR.SingleNestedAttribute{
					MarkdownDescription: "Static Inventory.",
				},
				Resource: &schemaR.SingleNestedAttribute{
					Optional: true,
				},
				DataSource: &schemaD.SingleNestedAttribute{
					Computed: true,
				},
				Attributes: map[string]superschema.Attribute{
					"inventory": superschema.StringAttribute{
						Common: &schemaR.StringAttribute{
							MarkdownDescription: "Static inventory content in INI format.",
						},
						Resource: &schemaR.StringAttribute{
							MarkdownDescription: "See examples above for format.",
							Required:            true,
						},
						DataSource: &schemaD.StringAttribute{
							Computed: true,
						},
					},
					"become_key_id": superschema.Int64Attribute{
						Common: &schemaR.Int64Attribute{
							MarkdownDescription: "The Project Key ID to use for privilege escalation (sudo) on hosts in the inventory. Only accepts `password` type Keys.",
						},
						Resource: &schemaR.Int64Attribute{
							Optional: true,
						},
						DataSource: &schemaD.Int64Attribute{
							Computed: true,
						},
					},
				},
			},
			"static_yaml": superschema.SingleNestedAttribute{
				Common: &schemaR.SingleNestedAttribute{
					MarkdownDescription: "Static YAML Inventory.",
				},
				Resource: &schemaR.SingleNestedAttribute{
					Optional: true,
				},
				DataSource: &schemaD.SingleNestedAttribute{
					Computed: true,
				},
				Attributes: map[string]superschema.Attribute{
					"inventory": superschema.StringAttribute{
						Common: &schemaR.StringAttribute{
							MarkdownDescription: "Static inventory content in YAML format.",
						},
						Resource: &schemaR.StringAttribute{
							MarkdownDescription: "See examples above for format.",
							Required:            true,
						},
						DataSource: &schemaD.StringAttribute{
							Computed: true,
						},
					},
					"become_key_id": superschema.Int64Attribute{
						Common: &schemaR.Int64Attribute{
							MarkdownDescription: "The Project Key ID to use for privilege escalation (sudo) on hosts in the inventory. Only accepts `password` type Keys.",
						},
						Resource: &schemaR.Int64Attribute{
							Optional: true,
						},
						DataSource: &schemaD.Int64Attribute{
							Computed: true,
						},
					},
				},
			},
			"file": superschema.SingleNestedAttribute{
				Common: &schemaR.SingleNestedAttribute{
					MarkdownDescription: "Inventory File.",
				},
				Resource: &schemaR.SingleNestedAttribute{
					Optional: true,
				},
				DataSource: &schemaD.SingleNestedAttribute{
					Computed: true,
				},
				Attributes: map[string]superschema.Attribute{
					"path": superschema.StringAttribute{
						Common: &schemaR.StringAttribute{
							MarkdownDescription: "The path to the inventory file, relative to the Template or custom Repository. Example: `folder/hosts.yml`.",
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
					"repository_id": superschema.Int64Attribute{
						Common: &schemaR.Int64Attribute{
							MarkdownDescription: "The ID of the Repository that contains the inventory file.",
						},
						Resource: &schemaR.Int64Attribute{
							Optional: true,
						},
						DataSource: &schemaD.Int64Attribute{
							Computed: true,
						},
					},
					"become_key_id": superschema.Int64Attribute{
						Common: &schemaR.Int64Attribute{
							MarkdownDescription: "The Project Key ID to use for privilege escalation (sudo) on hosts in the inventory. Only accepts `password` type Keys.",
						},
						Resource: &schemaR.Int64Attribute{
							Optional: true,
						},
						DataSource: &schemaD.Int64Attribute{
							Computed: true,
						},
					},
				},
			},
			"terraform_workspace": superschema.SingleNestedAttribute{
				Common: &schemaR.SingleNestedAttribute{
					MarkdownDescription: "Terraform Workspace.",
				},
				Resource: &schemaR.SingleNestedAttribute{
					Optional: true,
				},
				DataSource: &schemaD.SingleNestedAttribute{
					Computed: true,
				},
				Attributes: map[string]superschema.Attribute{
					"workspace": superschema.StringAttribute{
						Common: &schemaR.StringAttribute{
							MarkdownDescription: "The Terraform workspace name.",
						},
						Resource: &schemaR.StringAttribute{
							Required: true,
						},
						DataSource: &schemaD.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}
