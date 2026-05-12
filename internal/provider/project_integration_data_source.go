package provider

import (
	"context"
	"fmt"
	apiclient "terraform-provider-semaphoreui/semaphoreui/client"
	"terraform-provider-semaphoreui/semaphoreui/client/integration"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var (
	_ datasource.DataSource = &projectIntegrationDataSource{}
)

func NewProjectIntegrationDataSource() datasource.DataSource {
	return &projectIntegrationDataSource{}
}

type projectIntegrationDataSource struct {
	client *apiclient.SemaphoreUI
}

func (d *projectIntegrationDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	d.client = client
}

func (d *projectIntegrationDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project_integration"
}

func (d *projectIntegrationDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ProjectIntegrationSchema().GetDataSource(ctx)
}

func (d *projectIntegrationDataSource) GetIntegrationByName(ctx context.Context, projectID int64, name string) (*ProjectIntegrationModel, error) {
	response, err := d.client.Integration.GetProjectProjectIDIntegrations(&integration.GetProjectProjectIDIntegrationsParams{
		ProjectID: projectID,
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("could not read project integrations: %s", err.Error())
	}
	for _, integ := range response.Payload {
		if integ.Name == name {
			model := convertIntegrationResponseToProjectIntegrationModel(ctx, integ)
			return &model, nil
		}
	}
	return nil, fmt.Errorf("project integration with name %s not found", name)
}

func (d *projectIntegrationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ProjectIntegrationModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var model ProjectIntegrationModel
	if !config.ID.IsUnknown() && !config.ID.IsNull() {
		response, err := d.client.Integration.GetProjectProjectIDIntegrationsIntegrationID(&integration.GetProjectProjectIDIntegrationsIntegrationIDParams{
			ProjectID:     config.ProjectID.ValueInt64(),
			IntegrationID: config.ID.ValueInt64(),
		}, nil)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Reading SemaphoreUI Project Integration",
				"Could not read project integration, unexpected error: "+err.Error(),
			)
			return
		}
		model = convertIntegrationResponseToProjectIntegrationModel(ctx, response.Payload)
	} else if !config.Name.IsUnknown() && !config.Name.IsNull() {
		integ, err := d.GetIntegrationByName(ctx, config.ProjectID.ValueInt64(), config.Name.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Reading SemaphoreUI Project Integration",
				err.Error(),
			)
			return
		}
		model = *integ
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
