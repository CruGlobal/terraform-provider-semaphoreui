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
	ProjectRepositoryModel struct {
		ID        types.Int64  `tfsdk:"id"`
		ProjectID types.Int64  `tfsdk:"project_id"`
		Name      types.String `tfsdk:"name"`
		Url       types.String `tfsdk:"url"`
		Branch    types.String `tfsdk:"branch"`
		SSHKeyID  types.Int64  `tfsdk:"ssh_key_id"`
	}
)

func ProjectRepositorySchema() superschema.Schema {
	return superschema.Schema{
		Common: superschema.SchemaDetails{
			MarkdownDescription: "The project repository",
		},
		Resource: superschema.SchemaDetails{
			MarkdownDescription: "resource allows you to define the repositories used throughout a project. SemaphoreUI currently supports only Git repositories, including GitHub HTTP/S protocols and local paths.",
		},
		DataSource: superschema.SchemaDetails{
			MarkdownDescription: "data source allows you to read a repository.",
		},
		Attributes: map[string]superschema.Attribute{
			"id": superschema.Int64Attribute{
				Common: &schemaR.Int64Attribute{
					MarkdownDescription: "The repository ID.",
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
					MarkdownDescription: "The project ID that the repository belongs to.",
					Required:            true,
				},
				Resource: &schemaR.Int64Attribute{
					PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
				},
			},
			"name": superschema.StringAttribute{
				Common: &schemaR.StringAttribute{
					MarkdownDescription: "The display name of the repository.",
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
			"url": superschema.StringAttribute{
				Common: &schemaR.StringAttribute{
					MarkdownDescription: "The URI or path of the Git repository.",
				},
				Resource: &schemaR.StringAttribute{
					MarkdownDescription: "SemaphoreUI supports `ssh`, `http`, `https`, `file` and `git` URI schemes as well as absolute paths.",
					Required:            true,
				},
				DataSource: &schemaD.StringAttribute{
					Computed: true,
				},
			},
			"branch": superschema.StringAttribute{
				Common: &schemaR.StringAttribute{
					MarkdownDescription: "The branch of the repository to use.",
				},
				Resource: &schemaR.StringAttribute{
					MarkdownDescription: "Use an empty string for path based repositories.",
					Required:            true,
				},
				DataSource: &schemaD.StringAttribute{
					Computed: true,
				},
			},
			"ssh_key_id": superschema.Int64Attribute{
				Common: &schemaR.Int64Attribute{
					MarkdownDescription: "The Project Key ID to use for accessing the Git repository.",
				},
				Resource: &schemaR.Int64Attribute{
					MarkdownDescription: "This attribute is required for all repositories in SemaphoreUI. You should set it to the ID of a Key of type \"`none`\" if the repository doesn't require credentials.",
					Required:            true,
				},
				DataSource: &schemaD.Int64Attribute{
					Computed: true,
				},
			},
		},
	}
}
