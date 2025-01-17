package provider

import (
	"context"
	"fmt"
	apiclient "terraform-provider-semaphoreui/semaphoreui/client"
	"terraform-provider-semaphoreui/semaphoreui/client/project"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource = &projectInventoryDataSource{}
)

func NewProjectInventoryDataSource() datasource.DataSource {
	return &projectInventoryDataSource{}
}

type projectInventoryDataSource struct {
	client *apiclient.SemaphoreUI
}

func (d *projectInventoryDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

// Metadata returns the data source type name.
func (d *projectInventoryDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project_inventory"
}

// Schema defines the schema for the data source.
func (d *projectInventoryDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ProjectInventorySchema().GetDataSource(ctx)
}

func (d *projectInventoryDataSource) GetInventoryByName(projectID int64, name string) (*ProjectInventoryModel, error) {
	response, err := d.client.Project.GetProjectProjectIDInventory(&project.GetProjectProjectIDInventoryParams{
		ProjectID: projectID,
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("could not read project inventories: %s", err.Error())
	}
	for _, inventory := range response.Payload {
		if inventory.Name == name {
			model := convertInventoryResponseToProjectInventoryModel(inventory)
			return &model, nil
		}
	}
	return nil, fmt.Errorf("project inventory with name %s not found", name)
}

func (d *projectInventoryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ProjectInventoryModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var model ProjectInventoryModel
	if !config.ID.IsUnknown() && !config.ID.IsNull() {
		response, err := d.client.Project.GetProjectProjectIDInventoryInventoryID(&project.GetProjectProjectIDInventoryInventoryIDParams{
			ProjectID:   config.ProjectID.ValueInt64(),
			InventoryID: config.ID.ValueInt64(),
		}, nil)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Reading SemaphoreUI Project Inventory",
				"Could not read project inventory, unexpected error: "+err.Error(),
			)
			return
		}
		model = convertInventoryResponseToProjectInventoryModel(response.Payload)
	} else if !config.Name.IsUnknown() && !config.Name.IsNull() {
		inventory, err := d.GetInventoryByName(config.ProjectID.ValueInt64(), config.Name.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Reading Semaphore Project Inventory",
				err.Error(),
			)
			return
		}
		model = *inventory
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
