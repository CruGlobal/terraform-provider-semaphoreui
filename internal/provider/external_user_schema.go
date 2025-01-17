package provider

import (
	schemaR "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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
			MarkdownDescription: "The external user resource allows you to manage an external User in SemaphoreUI. This resource will not error if the user with the given username already exists and will not delete the user from SemaphoreUI when the resource is destroyed. This is useful when needing to manage the same user in multiple Terraform states.",
		},
		Attributes: map[string]superschema.Attribute{
			"id": superschema.Int64Attribute{
				Resource: &schemaR.Int64Attribute{
					MarkdownDescription: "The ID of the user.",
					Computed:            true,
					PlanModifiers: []planmodifier.Int64{
						int64planmodifier.UseStateForUnknown(),
					},
				},
			},
			"username": superschema.StringAttribute{
				Resource: &schemaR.StringAttribute{
					MarkdownDescription: "Username.",
					Required:            true,
				},
			},
			"name": superschema.StringAttribute{
				Resource: &schemaR.StringAttribute{
					MarkdownDescription: "Display name.",
					Required:            true,
				},
			},
			"email": superschema.StringAttribute{
				Resource: &schemaR.StringAttribute{
					MarkdownDescription: "Email address.",
					Required:            true,
				},
			},
			"admin": superschema.BoolAttribute{
				Resource: &schemaR.BoolAttribute{
					MarkdownDescription: "Indicates if the user is an admin.",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
			},
			"alert": superschema.BoolAttribute{
				Resource: &schemaR.BoolAttribute{
					MarkdownDescription: "Indicates if alerts should be sent to the user's email.",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
			},
			"external": superschema.BoolAttribute{
				Resource: &schemaR.BoolAttribute{
					MarkdownDescription: "Indicates if the user is linked to an external identity provider.",
					Computed:            true,
					Default:             booldefault.StaticBool(true),
				},
			},
			"created": superschema.StringAttribute{
				Resource: &schemaR.StringAttribute{
					MarkdownDescription: "Creation date of the user.",
					Computed:            true,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
			},
		},
	}
}
