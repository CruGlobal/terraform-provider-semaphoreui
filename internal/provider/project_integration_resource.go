package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	apiclient "terraform-provider-semaphoreui/semaphoreui/client"
	"terraform-provider-semaphoreui/semaphoreui/client/integration"
	"terraform-provider-semaphoreui/semaphoreui/models"
)

var (
	_ resource.Resource                = &projectIntegrationResource{}
	_ resource.ResourceWithConfigure   = &projectIntegrationResource{}
	_ resource.ResourceWithImportState = &projectIntegrationResource{}
)

func NewProjectIntegrationResource() resource.Resource {
	return &projectIntegrationResource{}
}

type projectIntegrationResource struct {
	client *apiclient.SemaphoreUI
}

func (r *projectIntegrationResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *projectIntegrationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project_integration"
}

func (r *projectIntegrationResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ProjectIntegrationSchema().GetResource(ctx)
}

func convertProjectIntegrationModelToIntegrationRequest(ctx context.Context, model ProjectIntegrationModel) *models.IntegrationRequest {
	req := models.IntegrationRequest{
		ProjectID:    model.ProjectID.ValueInt64(),
		TemplateID:   model.TemplateID.ValueInt64(),
		Name:         model.Name.ValueString(),
		AuthMethod:   model.AuthMethod.ValueString(),
		AuthSecretID: model.AuthSecretID.ValueInt64Pointer(),
		AuthHeader:   model.AuthHeader.ValueString(),
		Searchable:   model.Searchable.ValueBool(),
		TaskParams:   convertTaskParamsModelToTaskPrams(ctx, model.TaskParams),
	}
	if !model.ID.IsNull() && !model.ID.IsUnknown() {
		req.ID = model.ID.ValueInt64()
	}
	return &req
}

func convertIntegrationResponseToProjectIntegrationModel(ctx context.Context, payload *models.Integration) ProjectIntegrationModel {
	return ProjectIntegrationModel{
		ID:           types.Int64Value(payload.ID),
		ProjectID:    types.Int64Value(payload.ProjectID),
		TemplateID:   types.Int64Value(payload.TemplateID),
		Name:         types.StringValue(payload.Name),
		AuthMethod:   types.StringValue(payload.AuthMethod),
		AuthSecretID: types.Int64PointerValue(payload.AuthSecretID),
		AuthHeader:   types.StringValue(payload.AuthHeader),
		Searchable:   types.BoolValue(payload.Searchable),
		TaskParams:   convertTaskPramsToTaskParamsModel(ctx, payload.TaskParams),
	}
}

func (r *projectIntegrationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ProjectIntegrationModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := r.client.Integration.PostProjectProjectIDIntegrations(&integration.PostProjectProjectIDIntegrationsParams{
		ProjectID:   plan.ProjectID.ValueInt64(),
		Integration: convertProjectIntegrationModelToIntegrationRequest(ctx, plan),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating SemaphoreUI Project Integration",
			"Could not create project integration, unexpected error: "+err.Error(),
		)
		return
	}
	model := convertIntegrationResponseToProjectIntegrationModel(ctx, response.Payload)
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *projectIntegrationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ProjectIntegrationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := r.client.Integration.GetProjectProjectIDIntegrationsIntegrationID(&integration.GetProjectProjectIDIntegrationsIntegrationIDParams{
		ProjectID:     state.ProjectID.ValueInt64(),
		IntegrationID: state.ID.ValueInt64(),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Project Integration",
			"Could not read project integration, unexpected error: "+err.Error(),
		)
		return
	}
	model := convertIntegrationResponseToProjectIntegrationModel(ctx, response.Payload)
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *projectIntegrationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ProjectIntegrationModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.Integration.PutProjectProjectIDIntegrationsIntegrationID(&integration.PutProjectProjectIDIntegrationsIntegrationIDParams{
		ProjectID:     plan.ProjectID.ValueInt64(),
		IntegrationID: plan.ID.ValueInt64(),
		Integration:   convertProjectIntegrationModelToIntegrationRequest(ctx, plan),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating SemaphoreUI Project Integration",
			"Could not update project integration, unexpected error: "+err.Error(),
		)
		return
	}

	response, err := r.client.Integration.GetProjectProjectIDIntegrationsIntegrationID(&integration.GetProjectProjectIDIntegrationsIntegrationIDParams{
		ProjectID:     plan.ProjectID.ValueInt64(),
		IntegrationID: plan.ID.ValueInt64(),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Project Integration",
			"Could not read project integration after update, unexpected error: "+err.Error(),
		)
		return
	}
	model := convertIntegrationResponseToProjectIntegrationModel(ctx, response.Payload)
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *projectIntegrationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ProjectIntegrationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.Integration.DeleteProjectProjectIDIntegrationsIntegrationID(&integration.DeleteProjectProjectIDIntegrationsIntegrationIDParams{
		ProjectID:     state.ProjectID.ValueInt64(),
		IntegrationID: state.ID.ValueInt64(),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Removing SemaphoreUI Project Integration",
			"Could not remove project integration, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *projectIntegrationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	fields, err := parseImportFields(req.ID, []string{"project", "integration"})
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid Project Integration Import ID",
			"Could not parse import ID: "+err.Error(),
		)
		return
	}

	response, err := r.client.Integration.GetProjectProjectIDIntegrationsIntegrationID(&integration.GetProjectProjectIDIntegrationsIntegrationIDParams{
		ProjectID:     fields["project"],
		IntegrationID: fields["integration"],
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Project Integration",
			"Could not read project integration, unexpected error: "+err.Error(),
		)
		return
	}
	model := convertIntegrationResponseToProjectIntegrationModel(ctx, response.Payload)
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
