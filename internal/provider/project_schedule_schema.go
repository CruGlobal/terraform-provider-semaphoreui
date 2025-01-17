package provider

import (
	schemaD "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	schemaR "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	superschema "github.com/orange-cloudavenue/terraform-plugin-framework-superschema"
	"terraform-provider-semaphoreui/internal/stringvalidator"
)

type (
	ProjectScheduleModel struct {
		ID         types.Int64  `tfsdk:"id"`
		ProjectID  types.Int64  `tfsdk:"project_id"`
		TemplateID types.Int64  `tfsdk:"template_id"`
		Name       types.String `tfsdk:"name"`
		CronFormat types.String `tfsdk:"cron_format"`
		Enabled    types.Bool   `tfsdk:"enabled"`
	}
)

func ProjectScheduleSchema() superschema.Schema {
	return superschema.Schema{
		Common: superschema.SchemaDetails{
			MarkdownDescription: "The project schedule",
		},
		Resource: superschema.SchemaDetails{
			MarkdownDescription: "resource allows you to schedule the execution of templates in a project.",
		},
		DataSource: superschema.SchemaDetails{
			MarkdownDescription: "data source allows you to read a project schedule",
		},
		Attributes: map[string]superschema.Attribute{
			"id": superschema.Int64Attribute{
				Common: &schemaR.Int64Attribute{
					MarkdownDescription: "The schedule ID.",
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
					MarkdownDescription: "The project ID that the schedule belongs to.",
					Required:            true,
				},
				Resource: &schemaR.Int64Attribute{
					PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
				},
			},
			"template_id": superschema.Int64Attribute{
				Common: &schemaR.Int64Attribute{
					MarkdownDescription: "The template ID that the schedule executes.",
				},
				Resource: &schemaR.Int64Attribute{
					Required:      true,
					PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
				},
				DataSource: &schemaD.Int64Attribute{
					Computed: true,
				},
			},
			"name": superschema.StringAttribute{
				Common: &schemaR.StringAttribute{
					MarkdownDescription: "The display name of the schedule.",
				},
				Resource: &schemaR.StringAttribute{
					Required: true,
				},
				DataSource: &schemaD.StringAttribute{
					Computed: true,
				},
			},
			"cron_format": superschema.StringAttribute{
				Common: &schemaR.StringAttribute{
					MarkdownDescription: "The cron format of the schedule.",
				},
				Resource: &schemaR.StringAttribute{
					Required: true,
					Validators: []validator.String{
						stringvalidator.CronFormat(),
					},
				},
				DataSource: &schemaD.StringAttribute{
					Computed: true,
				},
			},
			"enabled": superschema.BoolAttribute{
				Common: &schemaR.BoolAttribute{
					MarkdownDescription: "Whether the schedule is enabled.",
				},
				Resource: &schemaR.BoolAttribute{
					Required: true,
				},
				DataSource: &schemaD.BoolAttribute{
					Computed: true,
				},
			},
		},
	}
}
