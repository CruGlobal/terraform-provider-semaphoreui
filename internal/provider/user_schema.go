package provider

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	schemaD "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schemaR "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	superschema "github.com/orange-cloudavenue/terraform-plugin-framework-superschema"
)

type UserModel struct {
	ID       types.Int64  `tfsdk:"id"`
	Created  types.String `tfsdk:"created"`
	Username types.String `tfsdk:"username"`
	Name     types.String `tfsdk:"name"`
	Email    types.String `tfsdk:"email"`
	Admin    types.Bool   `tfsdk:"admin"`
	External types.Bool   `tfsdk:"external"`
	Alert    types.Bool   `tfsdk:"alert"`
	Password types.String `tfsdk:"password"`
}

func userSchema() superschema.Schema {
	return superschema.Schema{
		Common: superschema.SchemaDetails{
			MarkdownDescription: "The user",
		},
		Resource: superschema.SchemaDetails{
			MarkdownDescription: "resource allows you to manage a User in SemaphoreUI.",
		},
		DataSource: superschema.SchemaDetails{
			MarkdownDescription: "data source allows you to read a User in SemaphoreUI.",
		},
		Attributes: map[string]superschema.Attribute{
			"id": superschema.Int64Attribute{
				Common: &schemaR.Int64Attribute{
					MarkdownDescription: "The ID of the user.",
				},
				Resource: &schemaR.Int64Attribute{
					Computed: true,
					PlanModifiers: []planmodifier.Int64{
						int64planmodifier.UseStateForUnknown(),
					},
				},
				DataSource: &schemaD.Int64Attribute{
					Optional: true,
					Computed: true,
					Validators: []validator.Int64{
						int64validator.ExactlyOneOf(
							path.MatchRoot("id"),
							path.MatchRoot("username"),
							path.MatchRoot("email"),
						),
					},
				},
			},
			"created": superschema.StringAttribute{
				Common: &schemaR.StringAttribute{
					MarkdownDescription: "Creation date of the user.",
					Computed:            true,
				},
				Resource: &schemaR.StringAttribute{
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
				DataSource: &schemaD.StringAttribute{},
			},
			"username": superschema.StringAttribute{
				Common: &schemaR.StringAttribute{
					MarkdownDescription: "Username.",
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
							path.MatchRoot("username"),
							path.MatchRoot("email"),
						),
					},
				},
			},
			"name": superschema.StringAttribute{
				Common: &schemaR.StringAttribute{
					MarkdownDescription: "Display name.",
				},
				Resource: &schemaR.StringAttribute{
					Required: true,
				},
				DataSource: &schemaD.StringAttribute{
					Computed: true,
				},
			},
			"email": superschema.StringAttribute{
				Common: &schemaR.StringAttribute{
					MarkdownDescription: "Email address.",
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
							path.MatchRoot("username"),
							path.MatchRoot("email"),
						),
					},
				},
			},
			"password": superschema.StringAttribute{
				Common: &schemaR.StringAttribute{
					Sensitive: true,
				},
				Resource: &schemaR.StringAttribute{
					MarkdownDescription: "Login Password. This value is never returned by the API and will be an empty string after import.",
					Optional:            true,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
				DataSource: &schemaD.StringAttribute{
					MarkdownDescription: "This value is never returned by the API and will be an empty string.",
					Computed:            true,
				},
			},
			"admin": superschema.BoolAttribute{
				Common: &schemaR.BoolAttribute{
					MarkdownDescription: "Indicates if the user is an admin.",
				},
				Resource: &schemaR.BoolAttribute{
					Optional: true,
					Computed: true,
					Default:  booldefault.StaticBool(false),
				},
				DataSource: &schemaD.BoolAttribute{
					Computed: true,
				},
			},
			"alert": superschema.BoolAttribute{
				Common: &schemaR.BoolAttribute{
					MarkdownDescription: "Indicates if alerts should be sent to the user's email.",
				},
				Resource: &schemaR.BoolAttribute{
					Optional: true,
					Computed: true,
					Default:  booldefault.StaticBool(false),
				},
				DataSource: &schemaD.BoolAttribute{
					Computed: true,
				},
			},
			"external": superschema.BoolAttribute{
				Common: &schemaR.BoolAttribute{
					MarkdownDescription: "Indicates if the user is linked to an external identity provider.",
				},
				Resource: &schemaR.BoolAttribute{
					Optional: true,
					Computed: true,
					Default:  booldefault.StaticBool(false),
					PlanModifiers: []planmodifier.Bool{
						boolplanmodifier.RequiresReplace(),
					},
				},
				DataSource: &schemaD.BoolAttribute{
					Computed: true,
				},
			},
		},
	}
}
