package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	apiclient "terraform-provider-semaphoreui/semaphoreui/client"
	"terraform-provider-semaphoreui/semaphoreui/client/project"
	"terraform-provider-semaphoreui/semaphoreui/models"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &projectViewResource{}
	_ resource.ResourceWithConfigure   = &projectViewResource{}
	_ resource.ResourceWithImportState = &projectViewResource{}
)

func NewProjectViewResource() resource.Resource {
	return &projectViewResource{}
}

type projectViewResource struct {
	client *apiclient.SemaphoreUI
}

func (r *projectViewResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *projectViewResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project_view"
}

func (r *projectViewResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ProjectViewSchema().GetResource(ctx)
}

func convertProjectViewModelToView(view ProjectViewModel) *models.ViewRequest {
	model := models.ViewRequest{
		ProjectID: view.ProjectID.ValueInt64(),
		Title:     view.Title.ValueString(),
		Position:  view.Position.ValueInt64(),
	}
	//if !view.ID.IsNull() && !view.ID.IsUnknown() {
	//	model.ID = view.ID.ValueInt64()
	//}
	return &model
}

func convertViewResponseToProjectViewModel(request *models.View) ProjectViewModel {
	return ProjectViewModel{
		ID:        types.Int64Value(request.ID),
		ProjectID: types.Int64Value(request.ProjectID),
		Position:  types.Int64Value(request.Position),
		Title:     types.StringValue(request.Title),
	}
}

func (r *projectViewResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ProjectViewModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := r.client.Project.PostProjectProjectIDViews(&project.PostProjectProjectIDViewsParams{
		ProjectID: plan.ProjectID.ValueInt64(),
		View:      convertProjectViewModelToView(plan),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating SemaphoreUI Project View",
			"Could not create project view, unexpected error: "+err.Error(),
		)
		return
	}
	model := convertViewResponseToProjectViewModel(response.Payload)

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *projectViewResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state ProjectViewModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := r.client.Project.GetProjectProjectIDViewsViewID(&project.GetProjectProjectIDViewsViewIDParams{
		ProjectID: state.ProjectID.ValueInt64(),
		ViewID:    state.ID.ValueInt64(),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Project View",
			"Could not read project view, unexpected error: "+err.Error(),
		)
		return
	}
	model := convertViewResponseToProjectViewModel(response.Payload)

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *projectViewResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan ProjectViewModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.Project.PutProjectProjectIDViewsViewID(&project.PutProjectProjectIDViewsViewIDParams{
		ProjectID: plan.ProjectID.ValueInt64(),
		ViewID:    plan.ID.ValueInt64(),
		View: &models.View{
			ID:        plan.ID.ValueInt64(),
			ProjectID: plan.ProjectID.ValueInt64(),
			Title:     plan.Title.ValueString(),
			Position:  plan.Position.ValueInt64(),
		},
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating SemaphoreUI Project View",
			"Could not update project view, unexpected error: "+err.Error(),
		)
		return
	}

	response, err := r.client.Project.GetProjectProjectIDViewsViewID(&project.GetProjectProjectIDViewsViewIDParams{
		ProjectID: plan.ProjectID.ValueInt64(),
		ViewID:    plan.ID.ValueInt64(),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Project View",
			"Could not read project view, unexpected error: "+err.Error(),
		)
		return
	}
	model := convertViewResponseToProjectViewModel(response.Payload)

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *projectViewResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state ProjectViewModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.Project.DeleteProjectProjectIDViewsViewID(&project.DeleteProjectProjectIDViewsViewIDParams{
		ProjectID: state.ProjectID.ValueInt64(),
		ViewID:    state.ID.ValueInt64(),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Removing SemaphoreUI Project View",
			"Could not remove project view, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *projectViewResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	fields, err := parseImportFields(req.ID, []string{"project", "view"})
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid Project View Import ID",
			"Could not parse import ID: "+err.Error(),
		)
		return
	}

	response, err := r.client.Project.GetProjectProjectIDViewsViewID(&project.GetProjectProjectIDViewsViewIDParams{
		ProjectID: fields["project"],
		ViewID:    fields["view"],
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Project View",
			"Could not read project view, unexpected error: "+err.Error(),
		)
		return
	}
	model := convertViewResponseToProjectViewModel(response.Payload)

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
