package provider

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	schemaD "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schemaR "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	superschema "github.com/orange-cloudavenue/terraform-plugin-framework-superschema"
)

type (
	ProjectKeyModel struct {
		ID            types.Int64              `tfsdk:"id"`
		ProjectID     types.Int64              `tfsdk:"project_id"`
		Name          types.String             `tfsdk:"name"`
		LoginPassword *ProjectKeyLoginPassword `tfsdk:"login_password"`
		SSH           *ProjectKeySSH           `tfsdk:"ssh"`
		None          *ProjectKeyNone          `tfsdk:"none"`
	}

	ProjectKeyLoginPassword struct {
		Login    types.String `tfsdk:"login"`
		Password types.String `tfsdk:"password"`
	}

	ProjectKeySSH struct {
		Login      types.String `tfsdk:"login"`
		Passphrase types.String `tfsdk:"passphrase"`
		PrivateKey types.String `tfsdk:"private_key"`
	}

	ProjectKeyNone struct{}
)

const (
	ProjectKeyTypeLoginPassword string = "login_password"
	ProjectKeyTypeSSH           string = "ssh"
	ProjectKeyTypeNone          string = "none"
)

func (model *ProjectKeyModel) Type() types.String {
	if model.LoginPassword != nil {
		return types.StringValue(ProjectKeyTypeLoginPassword)
	} else if model.SSH != nil {
		return types.StringValue(ProjectKeyTypeSSH)
	} else if model.None != nil {
		return types.StringValue(ProjectKeyTypeNone)
	}
	return types.StringUnknown()
}

func ProjectKeySchema() superschema.Schema {
	return superschema.Schema{
		Common: superschema.SchemaDetails{
			MarkdownDescription: "The project key",
		},
		Resource: superschema.SchemaDetails{
			MarkdownDescription: "resource allows you to define the credentials used throughout a project. Credentials can be Username/Password, SSH key or None. Project keys are used throughout SemaphoreUI, including Inventories, Repositories, Environments and Templates. When a resource doesn't need credentials, the None Key is used.",
		},
		DataSource: superschema.SchemaDetails{
			MarkdownDescription: "data source allows you to read the the credentials used throughout a project.",
		},
		Attributes: map[string]superschema.Attribute{
			"id": superschema.Int64Attribute{
				Common: &schemaR.Int64Attribute{
					MarkdownDescription: "The key ID.",
				},
				Resource: &schemaR.Int64Attribute{
					Computed:      true,
					PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
				},
				DataSource: &schemaD.Int64Attribute{
					Optional: true,
					Computed: true,
					Validators: []validator.Int64{
						int64validator.ExactlyOneOf(
							path.MatchRoot("id"),
							path.MatchRoot("name"),
						),
					},
				},
			},
			"project_id": superschema.Int64Attribute{
				Common: &schemaR.Int64Attribute{
					MarkdownDescription: "The project ID that the key belongs to.",
					Required:            true,
				},
				Resource: &schemaR.Int64Attribute{
					PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
				},
			},
			"name": superschema.StringAttribute{
				Common: &schemaR.StringAttribute{
					MarkdownDescription: "The display name of the key.",
				},
				Resource: &schemaR.StringAttribute{
					Required: true,
				},
				DataSource: &schemaD.StringAttribute{
					Optional: true,
					Computed: true,
					Validators: []validator.String{
						stringvalidator.ExactlyOneOf(
							path.MatchRoot("id"),
							path.MatchRoot("name"),
						),
					},
				},
			},
			ProjectKeyTypeLoginPassword: superschema.SingleNestedAttribute{
				Common: &schemaR.SingleNestedAttribute{
					MarkdownDescription: "A login password key.",
				},
				Resource: &schemaR.SingleNestedAttribute{
					Optional: true,
				},
				DataSource: &schemaD.SingleNestedAttribute{
					Computed: true,
				},
				Attributes: map[string]superschema.Attribute{
					"login": superschema.StringAttribute{
						Common: &schemaR.StringAttribute{
							MarkdownDescription: "The login username.",
						},
						Resource: &schemaR.StringAttribute{
							Optional: true,
						},
						DataSource: &schemaD.StringAttribute{
							Computed: true,
						},
					},
					"password": superschema.StringAttribute{
						Common: &schemaR.StringAttribute{
							MarkdownDescription: "The login password.",
							Sensitive:           true,
						},
						Resource: &schemaR.StringAttribute{
							Required: true,
						},
						DataSource: &schemaD.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			ProjectKeyTypeSSH: superschema.SingleNestedAttribute{
				Common: &schemaR.SingleNestedAttribute{
					MarkdownDescription: "A SSH key.",
				},
				Resource: &schemaR.SingleNestedAttribute{
					Optional: true,
				},
				DataSource: &schemaD.SingleNestedAttribute{
					Computed: true,
				},
				Attributes: map[string]superschema.Attribute{
					"login": superschema.StringAttribute{
						Common: &schemaR.StringAttribute{
							MarkdownDescription: "The login username.",
						},
						Resource: &schemaR.StringAttribute{
							Optional: true,
						},
						DataSource: &schemaD.StringAttribute{
							Computed: true,
						},
					},
					"passphrase": superschema.StringAttribute{
						Common: &schemaR.StringAttribute{
							MarkdownDescription: "The SSH Key passphrase.",
							Sensitive:           true,
						},
						Resource: &schemaR.StringAttribute{
							Optional: true,
						},
						DataSource: &schemaD.StringAttribute{
							Computed: true,
						},
					},
					"private_key": superschema.StringAttribute{
						Common: &schemaR.StringAttribute{
							MarkdownDescription: "The SSH private key.",
							Sensitive:           true,
						},
						Resource: &schemaR.StringAttribute{
							Optional: true,
						},
						DataSource: &schemaD.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			ProjectKeyTypeNone: superschema.SingleNestedAttribute{
				Common: &schemaR.SingleNestedAttribute{
					MarkdownDescription: "The special None key.",
				},
				Resource: &schemaR.SingleNestedAttribute{
					Optional: true,
				},
				DataSource: &schemaD.SingleNestedAttribute{
					Computed: true,
				},
			},
		},
	}
}
