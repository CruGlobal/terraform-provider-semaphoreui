package provider

import (
	"context"
	"fmt"
	apiclient "terraform-provider-semaphoreui/semaphoreui/client"
	"terraform-provider-semaphoreui/semaphoreui/client/project"
	"terraform-provider-semaphoreui/semaphoreui/models"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource = &projectTemplateDataSource{}
)

func NewProjectTemplateDataSource() datasource.DataSource {
	return &projectTemplateDataSource{}
}

type projectTemplateDataSource struct {
	client *apiclient.SemaphoreUI
}

func (d *projectTemplateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *projectTemplateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project_template"
}

// Schema defines the schema for the data source.
func (d *projectTemplateDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ProjectTemplateSchema().GetDataSource(ctx)
}

func (d *projectTemplateDataSource) GetTemplateByName(projectID int64, name string) (*models.Template, error) {
	response, err := d.client.Project.GetProjectProjectIDTemplates(&project.GetProjectProjectIDTemplatesParams{
		ProjectID: projectID,
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("could not read project repositories: %s", err.Error())
	}
	for _, template := range response.Payload {
		if template.Name == name {
			return template, nil
		}
	}
	return nil, fmt.Errorf("project template with name %s not found", name)
}

func (d *projectTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ProjectTemplateModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var model ProjectTemplateModel
	if !config.ID.IsUnknown() && !config.ID.IsNull() {
		response, err := d.client.Project.GetProjectProjectIDTemplatesTemplateID(&project.GetProjectProjectIDTemplatesTemplateIDParams{
			ProjectID:  config.ProjectID.ValueInt64(),
			TemplateID: config.ID.ValueInt64(),
		}, nil)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Reading SemaphoreUI Project Template",
				"Could not read project template, unexpected error: "+err.Error(),
			)
			return
		}
		model = convertTemplateResponseToProjectTemplateModel(ctx, response.Payload, &config)
	} else if !config.Name.IsUnknown() && !config.Name.IsNull() {
		template, err := d.GetTemplateByName(config.ProjectID.ValueInt64(), config.Name.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Reading SemaphoreUI Project Template",
				err.Error(),
			)
			return
		}
		model = convertTemplateResponseToProjectTemplateModel(ctx, template, &config)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
