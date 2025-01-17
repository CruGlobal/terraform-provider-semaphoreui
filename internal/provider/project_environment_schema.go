package provider

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	schemaD "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	schemaR "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	superschema "github.com/orange-cloudavenue/terraform-plugin-framework-superschema"
)

type (
	ProjectEnvironmentModel struct {
		ID          types.Int64        `tfsdk:"id"`
		ProjectID   types.Int64        `tfsdk:"project_id"`
		Name        types.String       `tfsdk:"name"`
		Variables   *map[string]string `tfsdk:"variables"`
		Environment *map[string]string `tfsdk:"environment"`
		Secrets     types.List         `tfsdk:"secrets"`
	}

	ProjectEnvironmentSecretModel struct {
		ID    types.Int64  `tfsdk:"id"`
		Type  types.String `tfsdk:"type"`
		Name  types.String `tfsdk:"name"`
		Value types.String `tfsdk:"value"`
	}
)

func ProjectEnvironmentSchema() superschema.Schema {
	return superschema.Schema{
		Common: superschema.SchemaDetails{
			MarkdownDescription: "The project environment (variable group)",
		},
		Resource: superschema.SchemaDetails{
			MarkdownDescription: "resource allows you to manage a list of extra and environment variables that can be used in a project's templates.",
		},
		DataSource: superschema.SchemaDetails{
			MarkdownDescription: "data source allows you to read project environment details.",
		},
		Attributes: map[string]superschema.Attribute{
			"id": superschema.Int64Attribute{
				Common: &schemaR.Int64Attribute{
					MarkdownDescription: "The environment ID.",
				},
				Resource: &schemaR.Int64Attribute{
					Computed:      true,
					PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
				},
				DataSource: &schemaD.Int64Attribute{
					Required: true,
				},
			},
			"project_id": superschema.Int64Attribute{
				Common: &schemaR.Int64Attribute{
					MarkdownDescription: "The project ID that the environment belongs to.",
					Required:            true,
				},
				Resource: &schemaR.Int64Attribute{
					PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
				},
			},
			"name": superschema.StringAttribute{
				Common: &schemaR.StringAttribute{
					MarkdownDescription: "The display name of the environment.",
				},
				Resource: &schemaR.StringAttribute{
					Required: true,
				},
				DataSource: &schemaD.StringAttribute{
					Computed: true,
				},
			},
			"variables": superschema.MapAttribute{
				Common: &schemaR.MapAttribute{
					MarkdownDescription: "Extra variables. Passed to Ansible as extra variables (`--extra-vars`) and Terraform/OpenTofu as variables (`-var`).",
					ElementType:         types.StringType,
				},
				Resource: &schemaR.MapAttribute{
					Optional: true,
				},
				DataSource: &schemaD.MapAttribute{
					Computed: true,
				},
			},
			"environment": superschema.MapAttribute{
				Common: &schemaR.MapAttribute{
					MarkdownDescription: "Environment variables.",
					ElementType:         types.StringType,
				},
				Resource: &schemaR.MapAttribute{
					Optional: true,
				},
				DataSource: &schemaD.MapAttribute{
					Computed: true,
				},
			},
			"secrets": superschema.ListNestedAttribute{
				Common: &schemaR.ListNestedAttribute{
					MarkdownDescription: "Secret variables of either `\"var\"` or `\"env\"` type. The `value` is encrypted and will be empty if imported.",
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
							MarkdownDescription: "The variable ID.",
							Computed:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
					},
					"type": superschema.StringAttribute{
						Common: &schemaR.StringAttribute{
							MarkdownDescription: "The variable type.",
						},
						Resource: &schemaR.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOf("env", "var"),
							},
						},
						DataSource: &schemaD.StringAttribute{
							Computed: true,
						},
					},
					"name": superschema.StringAttribute{
						Common: &schemaR.StringAttribute{
							MarkdownDescription: "The variable name.",
						},
						Resource: &schemaR.StringAttribute{
							Required: true,
						},
						DataSource: &schemaD.StringAttribute{
							Computed: true,
						},
					},
					"value": superschema.StringAttribute{
						Common: &schemaR.StringAttribute{
							MarkdownDescription: "The variable value.",
							Sensitive:           true,
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
