package provider

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	schemaD "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schemaR "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	superschema "github.com/orange-cloudavenue/terraform-plugin-framework-superschema"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ProjectModel struct {
	ID               types.Int64  `tfsdk:"id"`
	Created          types.String `tfsdk:"created"`
	Name             types.String `tfsdk:"name"`
	Alert            types.Bool   `tfsdk:"alert"`
	AlertChat        types.String `tfsdk:"alert_chat"`
	MaxParallelTasks types.Int64  `tfsdk:"max_parallel_tasks"`
	//Type             types.String `tfsdk:"type"`
}

func ProjectSchema() superschema.Schema {
	return superschema.Schema{
		Common: superschema.SchemaDetails{
			MarkdownDescription: "The project",
		},
		Resource: superschema.SchemaDetails{
			MarkdownDescription: "resource allows you to manage a project in SemaphoreUI.",
		},
		DataSource: superschema.SchemaDetails{
			MarkdownDescription: "data source allows you to read a project in SemaphoreUI.",
		},
		Attributes: map[string]superschema.Attribute{
			"id": superschema.Int64Attribute{
				Common: &schemaR.Int64Attribute{
					MarkdownDescription: "The ID of the project.",
				},
				Resource: &schemaR.Int64Attribute{
					Computed: true,
					PlanModifiers: []planmodifier.Int64{
						int64planmodifier.UseStateForUnknown(),
					},
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
			"name": superschema.StringAttribute{
				Common: &schemaR.StringAttribute{
					MarkdownDescription: "Project name.",
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
			"created": superschema.StringAttribute{
				Common: &schemaR.StringAttribute{
					MarkdownDescription: "Creation date of the project.",
					Computed:            true,
				},
				Resource: &schemaR.StringAttribute{
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
			},
			"alert": superschema.BoolAttribute{
				Common: &schemaR.BoolAttribute{
					MarkdownDescription: "Allow alerts for this project.",
					Computed:            true,
				},
				Resource: &schemaR.BoolAttribute{
					Optional: true,
					Default:  booldefault.StaticBool(false),
				},
			},
			"alert_chat": superschema.StringAttribute{
				Common: &schemaR.StringAttribute{
					MarkdownDescription: "Telegram chat ID.",
				},
				Resource: &schemaR.StringAttribute{
					Optional: true,
				},
				DataSource: &schemaD.StringAttribute{
					Computed: true,
				},
			},
			"max_parallel_tasks": superschema.Int64Attribute{
				Common: &schemaR.Int64Attribute{
					MarkdownDescription: "Maximum number of parallel tasks, `0` for unlimited.",
					Computed:            true,
				},
				Resource: &schemaR.Int64Attribute{
					Optional: true,
					Default:  int64default.StaticInt64(0),
					Validators: []validator.Int64{
						int64validator.AtLeast(0),
					},
				},
			},
			//"type": superschema.StringAttribute{
			//	Common: &schemaR.StringAttribute{
			//		MarkdownDescription: "Project type (currently unused).",
			//	},
			//	Resource: &schemaR.StringAttribute{
			//		Optional: true,
			//	},
			//	DataSource: &schemaD.StringAttribute{
			//		Computed: true,
			//	},
			//},
		},
	}
}
