package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	apiclient "terraform-provider-semaphoreui/semaphoreui/client"
	"terraform-provider-semaphoreui/semaphoreui/client/project"
	"terraform-provider-semaphoreui/semaphoreui/client/projects"
	"terraform-provider-semaphoreui/semaphoreui/models"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &projectResource{}
	_ resource.ResourceWithConfigure   = &projectResource{}
	_ resource.ResourceWithImportState = &projectResource{}
)

func NewProjectResource() resource.Resource {
	return &projectResource{}
}

type projectResource struct {
	client *apiclient.SemaphoreUI
}

func (r *projectResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *projectResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project"
}

// Schema defines the schema for the resource.
func (r *projectResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = projectSchema().GetResource(ctx)
}

func convertPayloadToProjectModel(payload *models.Project) projectModel {
	var alertChat types.String
	if payload.AlertChat == "" {
		alertChat = types.StringNull()
	} else {
		alertChat = types.StringValue(payload.AlertChat)
	}

	//var projType types.String
	//if payload.Type == "" {
	//	projType = types.StringNull()
	//} else {
	//	projType = types.StringValue(payload.Type)
	//}
	return projectModel{
		ID:               types.Int64Value(payload.ID),
		Name:             types.StringValue(payload.Name),
		Alert:            types.BoolValue(payload.Alert),
		AlertChat:        alertChat,
		MaxParallelTasks: types.Int64Value(*payload.MaxParallelTasks),
		Created:          types.StringValue(payload.Created),
		//Type:             projType,
	}
}

func (r *projectResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan projectModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	var request = models.ProjectRequest{
		Name:             plan.Name.ValueString(),
		Alert:            plan.Alert.ValueBool(),
		AlertChat:        plan.AlertChat.ValueString(),
		MaxParallelTasks: plan.MaxParallelTasks.ValueInt64Pointer(),
	}

	//Create new project
	response, err := r.client.Projects.PostProjects(&projects.PostProjectsParams{Project: &request}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Semaphore Project",
			"Could not create project, unexpected error: "+err.Error(),
		)
		return
	}

	plan = convertPayloadToProjectModel(response.Payload)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *projectResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state projectModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := r.client.Project.GetProjectProjectID(&project.GetProjectProjectIDParams{ProjectID: state.ID.ValueInt64()}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Semaphore Project",
			fmt.Sprintf("Could not read project ID %d: %s", state.ID.ValueInt64(), err.Error()),
		)
		return
	}

	// Overwrite with refreshed state
	state = convertPayloadToProjectModel(response.Payload)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *projectResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan projectModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	var request project.PutProjectProjectIDBody
	request.ID = plan.ID.ValueInt64()
	request.Name = plan.Name.ValueString()
	request.Alert = plan.Alert.ValueBool()
	request.AlertChat = plan.AlertChat.ValueString()
	request.MaxParallelTasks = plan.MaxParallelTasks.ValueInt64Pointer()
	//request.Type = plan.Type.ValueString()

	// Update existing project
	_, err := r.client.Project.PutProjectProjectID(&project.PutProjectProjectIDParams{ProjectID: plan.ID.ValueInt64(), Project: request}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Semaphore Project",
			fmt.Sprintf("Could not update project, unexpected error: %s", err.Error()),
		)
		return
	}

	// Fetch updated project as PutProjectProjectID does not return updated project
	response, err := r.client.Project.GetProjectProjectID(&project.GetProjectProjectIDParams{ProjectID: plan.ID.ValueInt64()}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Semaphore Project",
			fmt.Sprintf("Could not read project ID %d: %s", plan.ID.ValueInt64(), err.Error()),
		)
		return
	}

	// Update resource state with updated project
	plan = convertPayloadToProjectModel(response.Payload)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *projectResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state projectModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing order
	_, err := r.client.Project.DeleteProjectProjectID(&project.DeleteProjectProjectIDParams{ProjectID: state.ID.ValueInt64()}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Semaphore Project",
			fmt.Sprintf("Could not delete project, unexpected error: %s", err.Error()),
		)
		return
	}
}

func (r *projectResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	fields, err := parseImportFields(req.ID, []string{"project"})
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid Project Import ID",
			"Could not parse import ID: "+err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), fields["project"])...)
}
