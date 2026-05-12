package provider

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
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
)

type ProjectIntegrationModel struct {
	ID           types.Int64      `tfsdk:"id"`
	ProjectID    types.Int64      `tfsdk:"project_id"`
	TemplateID   types.Int64      `tfsdk:"template_id"`
	Name         types.String     `tfsdk:"name"`
	AuthMethod   types.String     `tfsdk:"auth_method"`
	AuthSecretID types.Int64      `tfsdk:"auth_secret_id"`
	AuthHeader   types.String     `tfsdk:"auth_header"`
	Searchable   types.Bool       `tfsdk:"searchable"`
	TaskParams   *TaskParamsModel `tfsdk:"task_params"`
}

func ProjectIntegrationSchema() superschema.Schema {
	return superschema.Schema{
		Common: superschema.SchemaDetails{
			MarkdownDescription: "The project integration",
		},
		Resource: superschema.SchemaDetails{
			MarkdownDescription: "resource allows you to manage an incoming webhook integration that triggers a SemaphoreUI template task.",
		},
		DataSource: superschema.SchemaDetails{
			MarkdownDescription: "data source allows you to read a project integration.",
		},
		Attributes: map[string]superschema.Attribute{
			"id": superschema.Int64Attribute{
				Common: &schemaR.Int64Attribute{
					MarkdownDescription: "The integration ID.",
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
					MarkdownDescription: "The project ID that the integration belongs to.",
					Required:            true,
				},
				Resource: &schemaR.Int64Attribute{
					PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
				},
			},
			"template_id": superschema.Int64Attribute{
				Common: &schemaR.Int64Attribute{
					MarkdownDescription: "The template ID that this integration triggers when invoked.",
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
					MarkdownDescription: "The display name of the integration.",
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
			"auth_method": superschema.StringAttribute{
				Common: &schemaR.StringAttribute{
					MarkdownDescription: "How incoming requests are authenticated. Known values: `none`, `token`, `hmac`, `github`, `gitlab`, `bitbucket`. Defaults to `none`.",
				},
				Resource: &schemaR.StringAttribute{
					Optional: true,
					Computed: true,
					Default:  stringdefault.StaticString("none"),
				},
				DataSource: &schemaD.StringAttribute{
					Computed: true,
				},
			},
			"auth_secret_id": superschema.Int64Attribute{
				Common: &schemaR.Int64Attribute{
					MarkdownDescription: "The project key ID that holds the credential used to verify incoming requests (relevant when `auth_method` is `token` or `hmac`).",
				},
				Resource: &schemaR.Int64Attribute{
					Optional: true,
				},
				DataSource: &schemaD.Int64Attribute{
					Computed: true,
				},
			},
			"auth_header": superschema.StringAttribute{
				Common: &schemaR.StringAttribute{
					MarkdownDescription: "The HTTP header containing the auth token or signature (e.g. `Authorization`, `X-Hub-Signature`).",
				},
				Resource: &schemaR.StringAttribute{
					Optional: true,
					Computed: true,
					Default:  stringdefault.StaticString(""),
				},
				DataSource: &schemaD.StringAttribute{
					Computed: true,
				},
			},
			"searchable": superschema.BoolAttribute{
				Common: &schemaR.BoolAttribute{
					MarkdownDescription: "Whether to index this integration's task history for search.",
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
			"task_params": TaskParamsAttribute(),
		},
	}
}
