package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/types"
	apiclient "terraform-provider-semaphoreui/semaphoreui/client"
	"terraform-provider-semaphoreui/semaphoreui/client/project"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource = &projectKeyDataSource{}
)

func NewProjectKeyDataSource() datasource.DataSource {
	return &projectKeyDataSource{}
}

type projectKeyDataSource struct {
	client *apiclient.SemaphoreUI
}

func (d *projectKeyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *projectKeyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project_key"
}

// Schema defines the schema for the data source.
func (d *projectKeyDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ProjectKeySchema().GetDataSource(ctx)
}

func (d *projectKeyDataSource) GetKeyByName(projectID int64, name string) (*ProjectKeyModel, error) {
	response, err := d.client.Project.GetProjectProjectIDKeys(&project.GetProjectProjectIDKeysParams{
		ProjectID: projectID,
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("could not read project keys: %s", err.Error())
	}
	for _, key := range response.Payload {
		if key.Name == name {
			model := ProjectKeyModel{
				ProjectID: types.Int64Value(key.ProjectID),
				ID:        types.Int64Value(key.ID),
				Name:      types.StringValue(key.Name),
			}
			switch key.Type {
			case ProjectKeyTypeNone:
				model.None = &ProjectKeyNone{}
			case ProjectKeyTypeLoginPassword:
				model.LoginPassword = &ProjectKeyLoginPassword{
					Password: types.StringValue(""),
				}
			case ProjectKeyTypeSSH:
				model.SSH = &ProjectKeySSH{
					PrivateKey: types.StringValue(""),
				}
			}
			return &model, nil
		}
	}
	return nil, fmt.Errorf("project key with name %s not found", name)
}

func (d *projectKeyDataSource) GetKeyByID(projectID int64, ID int64) (*ProjectKeyModel, error) {
	response, err := d.client.Project.GetProjectProjectIDKeys(&project.GetProjectProjectIDKeysParams{
		ProjectID: projectID,
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("could not read project keys: %s", err.Error())
	}
	for _, key := range response.Payload {
		if key.ID == ID {
			model := ProjectKeyModel{
				ProjectID: types.Int64Value(key.ProjectID),
				ID:        types.Int64Value(key.ID),
				Name:      types.StringValue(key.Name),
			}
			switch key.Type {
			case ProjectKeyTypeNone:
				model.None = &ProjectKeyNone{}
			case ProjectKeyTypeLoginPassword:
				model.LoginPassword = &ProjectKeyLoginPassword{
					Password: types.StringValue(""),
				}
			case ProjectKeyTypeSSH:
				model.SSH = &ProjectKeySSH{
					PrivateKey: types.StringValue(""),
				}
			}
			return &model, nil
		}
	}
	return nil, fmt.Errorf("project key with id %d not found", ID)
}

func (d *projectKeyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ProjectKeyModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var model ProjectKeyModel
	if !config.ID.IsUnknown() && !config.ID.IsNull() {
		key, err := d.GetKeyByID(config.ProjectID.ValueInt64(), config.ID.ValueInt64())
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Reading SemaphoreUI Project Key",
				err.Error(),
			)
			return
		}
		model = *key
	} else if !config.Name.IsUnknown() && !config.Name.IsNull() {
		key, err := d.GetKeyByName(config.ProjectID.ValueInt64(), config.Name.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Reading SemaphoreUI Project Key",
				err.Error(),
			)
			return
		}
		model = *key
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
