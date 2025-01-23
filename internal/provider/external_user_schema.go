package provider

import (
	schemaD "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	schemaR "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	superschema "github.com/orange-cloudavenue/terraform-plugin-framework-superschema"
)

type ExternalUserModel struct {
	ID       types.Int64  `tfsdk:"id"`
	Username types.String `tfsdk:"username"`
	Name     types.String `tfsdk:"name"`
	Email    types.String `tfsdk:"email"`
	Admin    types.Bool   `tfsdk:"admin"`
	Alert    types.Bool   `tfsdk:"alert"`
	External types.Bool   `tfsdk:"external"`
	Created  types.String `tfsdk:"created"`
}

func ExternalUserSchema() superschema.Schema {
	return superschema.Schema{
		Resource: superschema.SchemaDetails{
			MarkdownDescription: "The external user resource allows you to lookup an external User in SemaphoreUI. This special data source will create the external user in SemaphoreUI if they do not exist.",
		},
		Attributes: map[string]superschema.Attribute{
			"username": superschema.StringAttribute{
				Common: &schemaR.StringAttribute{
					MarkdownDescription: "Username.",
				},
				DataSource: &schemaD.StringAttribute{
					Required: true,
				},
			},
			"id": superschema.Int64Attribute{
				DataSource: &schemaD.Int64Attribute{
					MarkdownDescription: "The ID of the external user.",
					Computed:            true,
				},
			},
			"name": superschema.StringAttribute{
				DataSource: &schemaD.StringAttribute{
					MarkdownDescription: "Display name. Defaults to the username if not supplied.",
					Optional:            true,
					Computed:            true,
				},
			},
			"email": superschema.StringAttribute{
				DataSource: &schemaD.StringAttribute{
					MarkdownDescription: "Email address. Defaults to the username if not supplied.",
					Optional:            true,
					Computed:            true,
				},
			},
			"admin": superschema.BoolAttribute{
				DataSource: &schemaD.BoolAttribute{
					MarkdownDescription: "Indicates if the user is an admin.",
					Computed:            true,
				},
			},
			"alert": superschema.BoolAttribute{
				DataSource: &schemaD.BoolAttribute{
					MarkdownDescription: "Indicates if alerts should be sent to the user's email.",
					Computed:            true,
				},
			},
			"external": superschema.BoolAttribute{
				DataSource: &schemaD.BoolAttribute{
					MarkdownDescription: "Indicates if the user is linked to an external identity provider.",
					Computed:            true,
				},
			},
			"created": superschema.StringAttribute{
				DataSource: &schemaD.StringAttribute{
					MarkdownDescription: "Creation date of the user.",
					Computed:            true,
				},
			},
		},
	}
}
