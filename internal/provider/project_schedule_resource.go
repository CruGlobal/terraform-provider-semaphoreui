package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-semaphoreui/internal/stringvalidator"
	apiclient "terraform-provider-semaphoreui/semaphoreui/client"
	"terraform-provider-semaphoreui/semaphoreui/client/schedule"
	"terraform-provider-semaphoreui/semaphoreui/models"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &projectScheduleResource{}
	_ resource.ResourceWithConfigure   = &projectScheduleResource{}
	_ resource.ResourceWithImportState = &projectScheduleResource{}
)

func NewProjectScheduleResource() resource.Resource {
	return &projectScheduleResource{}
}

type projectScheduleResource struct {
	client *apiclient.SemaphoreUI
}

func (r *projectScheduleResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *projectScheduleResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project_schedule"
}

type projectScheduleModel struct {
	ID         types.Int64  `tfsdk:"id"`
	ProjectID  types.Int64  `tfsdk:"project_id"`
	TemplateID types.Int64  `tfsdk:"template_id"`
	Name       types.String `tfsdk:"name"`
	CronFormat types.String `tfsdk:"cron_format"`
	Enabled    types.Bool   `tfsdk:"enabled"`
}

func (r *projectScheduleResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: `Provides a SemaphoreUI Project Schedule resource.

Allows scheduling the execution of templates in a project.`,
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				MarkdownDescription: "The repository ID.",
				Computed:            true,
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"project_id": schema.Int64Attribute{
				MarkdownDescription: "The project ID that the repository belongs to.",
				Required:            true,
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"template_id": schema.Int64Attribute{
				MarkdownDescription: "The template ID that the schedule executes.",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The display name of the repository.",
				Required:            true,
			},
			"cron_format": schema.StringAttribute{
				MarkdownDescription: "The cron format of the schedule.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.CronFormat(),
				},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether the schedule is enabled.",
				Required:            true,
			},
		},
	}
}

func convertProjectScheduleModelToRepositorySchedule(schedule projectScheduleModel) *models.ScheduleRequest {
	model := models.ScheduleRequest{
		ProjectID:  schedule.ProjectID.ValueInt64(),
		TemplateID: schedule.TemplateID.ValueInt64(),
		Name:       schedule.Name.ValueString(),
		CronFormat: schedule.CronFormat.ValueString(),
		Active:     schedule.Enabled.ValueBool(),
	}
	if !schedule.ID.IsNull() && !schedule.ID.IsUnknown() {
		model.ID = schedule.ID.ValueInt64()
	}
	return &model
}

func convertScheduleResponseToProjectScheduleModel(request *models.Schedule) projectScheduleModel {
	return projectScheduleModel{
		ID:         types.Int64Value(request.ID),
		ProjectID:  types.Int64Value(request.ProjectID),
		TemplateID: types.Int64Value(request.TemplateID),
		Name:       types.StringValue(request.Name),
		CronFormat: types.StringValue(request.CronFormat),
		Enabled:    types.BoolValue(request.Active),
	}
}

func (r *projectScheduleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan projectScheduleModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := r.client.Schedule.PostProjectProjectIDSchedules(&schedule.PostProjectProjectIDSchedulesParams{
		ProjectID: plan.ProjectID.ValueInt64(),
		Schedule:  convertProjectScheduleModelToRepositorySchedule(plan),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating SemaphoreUI Project Schedule",
			"Could not create project schedule, unexpected error: "+err.Error(),
		)
		return
	}
	model := convertScheduleResponseToProjectScheduleModel(response.Payload)

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *projectScheduleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state projectScheduleModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := r.client.Schedule.GetProjectProjectIDSchedulesScheduleID(&schedule.GetProjectProjectIDSchedulesScheduleIDParams{
		ProjectID:  state.ProjectID.ValueInt64(),
		ScheduleID: state.ID.ValueInt64(),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Project Schedule",
			"Could not read project schedule, unexpected error: "+err.Error(),
		)
		return
	}
	model := convertScheduleResponseToProjectScheduleModel(response.Payload)

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *projectScheduleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan projectScheduleModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.Schedule.PutProjectProjectIDSchedulesScheduleID(&schedule.PutProjectProjectIDSchedulesScheduleIDParams{
		ProjectID:  plan.ProjectID.ValueInt64(),
		ScheduleID: plan.ID.ValueInt64(),
		Schedule:   convertProjectScheduleModelToRepositorySchedule(plan),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating SemaphoreUI Project Schedule",
			"Could not update project schedule, unexpected error: "+err.Error(),
		)
		return
	}

	response, err := r.client.Schedule.GetProjectProjectIDSchedulesScheduleID(&schedule.GetProjectProjectIDSchedulesScheduleIDParams{
		ProjectID:  plan.ProjectID.ValueInt64(),
		ScheduleID: plan.ID.ValueInt64(),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Project Schedule",
			"Could not read project schedule, unexpected error: "+err.Error(),
		)
		return
	}
	model := convertScheduleResponseToProjectScheduleModel(response.Payload)

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *projectScheduleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state projectScheduleModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.Schedule.DeleteProjectProjectIDSchedulesScheduleID(&schedule.DeleteProjectProjectIDSchedulesScheduleIDParams{
		ProjectID:  state.ProjectID.ValueInt64(),
		ScheduleID: state.ID.ValueInt64(),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Removing SemaphoreUI Project Schedule",
			"Could not remove project schedule, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *projectScheduleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	fields, err := parseImportFields(req.ID, []string{"project", "schedule"})
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid Project Repository Import ID",
			"Could not parse import ID: "+err.Error(),
		)
		return
	}

	response, err := r.client.Schedule.GetProjectProjectIDSchedulesScheduleID(&schedule.GetProjectProjectIDSchedulesScheduleIDParams{
		ProjectID:  fields["project"],
		ScheduleID: fields["schedule"],
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Project Schedule",
			"Could not read project schedule, unexpected error: "+err.Error(),
		)
		return
	}
	model := convertScheduleResponseToProjectScheduleModel(response.Payload)

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
