package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
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

func (r *projectScheduleResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ProjectScheduleSchema().GetResource(ctx)
}

func convertProjectScheduleModelToRepositorySchedule(schedule ProjectScheduleModel) *models.ScheduleRequest {
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

func convertScheduleResponseToProjectScheduleModel(request *models.Schedule) ProjectScheduleModel {
	return ProjectScheduleModel{
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
	var plan ProjectScheduleModel
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
	var state ProjectScheduleModel
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
	var plan ProjectScheduleModel
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
	var state ProjectScheduleModel
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
