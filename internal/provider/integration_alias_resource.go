package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	apiclient "terraform-provider-semaphoreui/semaphoreui/client"
	"terraform-provider-semaphoreui/semaphoreui/client/integration"
	"terraform-provider-semaphoreui/semaphoreui/models"
)

var (
	_ resource.Resource                = &integrationAliasResource{}
	_ resource.ResourceWithConfigure   = &integrationAliasResource{}
	_ resource.ResourceWithImportState = &integrationAliasResource{}
)

func NewIntegrationAliasResource() resource.Resource {
	return &integrationAliasResource{}
}

type integrationAliasResource struct {
	client *apiclient.SemaphoreUI
}

func (r *integrationAliasResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *integrationAliasResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_integration_alias"
}

func (r *integrationAliasResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = IntegrationAliasSchema().GetResource(ctx)
}

func (r *integrationAliasResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IntegrationAliasModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var payload *models.IntegrationAlias
	if !plan.IntegrationID.IsNull() && !plan.IntegrationID.IsUnknown() {
		response, err := r.client.Integration.PostProjectProjectIDIntegrationsIntegrationIDAliases(
			&integration.PostProjectProjectIDIntegrationsIntegrationIDAliasesParams{
				ProjectID:     plan.ProjectID.ValueInt64(),
				IntegrationID: plan.IntegrationID.ValueInt64(),
			}, nil)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Creating SemaphoreUI Integration Alias",
				"Could not create integration-scoped alias, unexpected error: "+err.Error(),
			)
			return
		}
		payload = response.Payload
	} else {
		response, err := r.client.Integration.PostProjectProjectIDIntegrationsAliases(
			&integration.PostProjectProjectIDIntegrationsAliasesParams{
				ProjectID: plan.ProjectID.ValueInt64(),
			}, nil)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Creating SemaphoreUI Integration Alias",
				"Could not create project-scoped alias, unexpected error: "+err.Error(),
			)
			return
		}
		payload = response.Payload
	}

	plan.ID = types.Int64Value(payload.ID)
	plan.URL = types.StringValue(payload.URL)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// findAlias looks up an alias by ID in the appropriate scope's list (the
// API doesn't expose a GET-by-id, only list endpoints).
func (r *integrationAliasResource) findAlias(projectID, integrationID, aliasID int64) (*models.IntegrationAlias, error) {
	if integrationID != 0 {
		response, err := r.client.Integration.GetProjectProjectIDIntegrationsIntegrationIDAliases(
			&integration.GetProjectProjectIDIntegrationsIntegrationIDAliasesParams{
				ProjectID:     projectID,
				IntegrationID: integrationID,
			}, nil)
		if err != nil {
			return nil, err
		}
		for _, a := range response.Payload {
			if a.ID == aliasID {
				return a, nil
			}
		}
		return nil, nil
	}

	response, err := r.client.Integration.GetProjectProjectIDIntegrationsAliases(
		&integration.GetProjectProjectIDIntegrationsAliasesParams{
			ProjectID: projectID,
		}, nil)
	if err != nil {
		return nil, err
	}
	for _, a := range response.Payload {
		if a.ID == aliasID {
			return a, nil
		}
	}
	return nil, nil
}

func (r *integrationAliasResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IntegrationAliasModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	alias, err := r.findAlias(state.ProjectID.ValueInt64(), state.IntegrationID.ValueInt64(), state.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Integration Alias",
			"Could not read integration alias, unexpected error: "+err.Error(),
		)
		return
	}
	if alias == nil {
		// Drift: alias deleted out-of-band. Remove from state.
		resp.State.RemoveResource(ctx)
		return
	}

	state.URL = types.StringValue(alias.URL)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *integrationAliasResource) Update(_ context.Context, _ resource.UpdateRequest, resp *resource.UpdateResponse) {
	// All settable attributes are RequiresReplace, so Update should never be
	// called. Defensive error if it is.
	resp.Diagnostics.AddError(
		"Integration Alias Update Not Supported",
		"Integration aliases are immutable. Any change to project_id or integration_id forces resource replacement; Update should not be reachable.",
	)
}

func (r *integrationAliasResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IntegrationAliasModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !state.IntegrationID.IsNull() && state.IntegrationID.ValueInt64() != 0 {
		_, err := r.client.Integration.DeleteProjectProjectIDIntegrationsIntegrationIDAliasesAliasID(
			&integration.DeleteProjectProjectIDIntegrationsIntegrationIDAliasesAliasIDParams{
				ProjectID:     state.ProjectID.ValueInt64(),
				IntegrationID: state.IntegrationID.ValueInt64(),
				AliasID:       state.ID.ValueInt64(),
			}, nil)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Removing SemaphoreUI Integration Alias",
				"Could not remove integration-scoped alias, unexpected error: "+err.Error(),
			)
		}
		return
	}

	_, err := r.client.Integration.DeleteProjectProjectIDIntegrationsAliasesAliasID(
		&integration.DeleteProjectProjectIDIntegrationsAliasesAliasIDParams{
			ProjectID: state.ProjectID.ValueInt64(),
			AliasID:   state.ID.ValueInt64(),
		}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Removing SemaphoreUI Integration Alias",
			"Could not remove project-scoped alias, unexpected error: "+err.Error(),
		)
	}
}

// ImportState supports two ID shapes so users can import either scope:
//
//	project/{project_id}/alias/{alias_id}                                 -> project-scoped
//	project/{project_id}/integration/{integration_id}/alias/{alias_id}    -> integration-scoped
func (r *integrationAliasResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	fields, err := parseImportFields(req.ID, []string{"project", "alias"})
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid Integration Alias Import ID",
			fmt.Sprintf("Could not parse import ID %q: %s. Expected `project/{id}/alias/{id}` or `project/{id}/integration/{id}/alias/{id}`.", req.ID, err.Error()),
		)
		return
	}

	state := IntegrationAliasModel{
		ID:        types.Int64Value(fields["alias"]),
		ProjectID: types.Int64Value(fields["project"]),
	}
	if integrationID, ok := fields["integration"]; ok {
		state.IntegrationID = types.Int64Value(integrationID)
	} else {
		state.IntegrationID = types.Int64Null()
	}

	alias, err := r.findAlias(state.ProjectID.ValueInt64(), state.IntegrationID.ValueInt64(), state.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Integration Alias",
			"Could not read integration alias during import, unexpected error: "+err.Error(),
		)
		return
	}
	if alias == nil {
		resp.Diagnostics.AddError(
			"Integration Alias Not Found",
			fmt.Sprintf("No alias with id=%d found in the specified scope.", state.ID.ValueInt64()),
		)
		return
	}
	state.URL = types.StringValue(alias.URL)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
