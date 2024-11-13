package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/resourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"regexp"
	apiclient "terraform-provider-semaphoreui/semaphoreui/client"
	"terraform-provider-semaphoreui/semaphoreui/client/project"
	"terraform-provider-semaphoreui/semaphoreui/models"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                     = &projectInventoryResource{}
	_ resource.ResourceWithConfigure        = &projectInventoryResource{}
	_ resource.ResourceWithImportState      = &projectInventoryResource{}
	_ resource.ResourceWithConfigValidators = &projectInventoryResource{}
)

func NewProjectInventoryResource() resource.Resource {
	return &projectInventoryResource{}
}

type projectInventoryResource struct {
	client *apiclient.SemaphoreUI
}

func (r *projectInventoryResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// Metadata returns the resource type name.
func (r *projectInventoryResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project_inventory"
}

// Schema defines the schema for the resource.
func (r *projectInventoryResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: `Provides a SemaphoreUI Project Inventory resource.

Project Inventory is used to define the Ansible inventory or a Terraform/OpenTofu workspace for a project. Only one of the inventory types (` + "`static`, `static_yaml`, `file` or `terraform_workspace`" + `) can be defined per inventory.`,
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				MarkdownDescription: "The inventory ID.",
				Computed:            true,
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"project_id": schema.Int64Attribute{
				MarkdownDescription: "The project ID that the inventory belongs to.",
				Required:            true,
				PlanModifiers:       []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The display name of the inventory or workspace.",
				Required:            true,
			},
			"ssh_key_id": schema.Int64Attribute{
				MarkdownDescription: `The Project Key ID to use for accessing hosts in the inventory.

This attribute is required for all inventory types in SemaphoreUI. You should set it to the ID of an Key of type ` + "`none`" + ` if the inventory doesn't require credentials, or for Workspace type inventories.`,
				Required: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"static": schema.SingleNestedAttribute{
				MarkdownDescription: "Static Inventory.",
				Attributes: map[string]schema.Attribute{
					"inventory": schema.StringAttribute{
						MarkdownDescription: `Static inventory content in INI format.

Example:
` + "```hcl" + `
inventory = <<-EOT
  mail.example.com

  [webservers]
  foo.example.com
  bar.example.com

  [dbservers]
  one.example.com
  two.example.com
  three.example.com
EOT
` + "```",
						Required: true,
					},
					"become_key_id": schema.Int64Attribute{
						MarkdownDescription: "The Project Key ID to use for privilege escalation (sudo) on hosts in the inventory. Only accepts `password` type Keys.",
						Optional:            true,
					},
				},
				Optional: true,
			},
			"static_yaml": schema.SingleNestedAttribute{
				MarkdownDescription: "Static YAML Inventory.",
				Attributes: map[string]schema.Attribute{
					"inventory": schema.StringAttribute{
						MarkdownDescription: `Static inventory content in YAML format.

Example:
` + "```hcl" + `
inventory = yamlencode({
  ungrouped = {
    hosts = {
      mail.example.com = {}
    }
  }
  webservers = {
    hosts = {
      foo.example.com = {}
      bar.example.com = {}
    }
  }
  dbservers = {
    hosts = {
      one.example.com = {}
      two.example.com = {}
      three.example.com = {}
    }
  }
})
` + "```",
						Required: true,
					},
					"become_key_id": schema.Int64Attribute{
						MarkdownDescription: "The Project Key ID to use for privilege escalation (sudo) on hosts in the inventory. Only accepts `password` type Keys.",
						Optional:            true,
					},
				},
				Optional: true,
			},
			"file": schema.SingleNestedAttribute{
				MarkdownDescription: "Inventory File.",
				Attributes: map[string]schema.Attribute{
					"path": schema.StringAttribute{
						MarkdownDescription: "The path to the inventory file, relative to the Template or custom Repository. Example: `folder/hosts.yml`",
						Required:            true,
						Validators: []validator.String{
							// Only relative paths are allowed
							stringvalidator.RegexMatches(
								regexp.MustCompile(`^[^/].*$`),
								"must be a relative path (path/to/inventory)",
							),
						},
					},
					"repository_id": schema.Int64Attribute{
						MarkdownDescription: "The ID of the Repository that contains the inventory file.",
						Optional:            true,
					},
					"become_key_id": schema.Int64Attribute{
						MarkdownDescription: "The Project Key ID to use for privilege escalation (sudo) on hosts in the inventory. Only accepts `password` type Keys.",
						Optional:            true,
					},
				},
				Optional: true,
			},
			"terraform_workspace": schema.SingleNestedAttribute{
				MarkdownDescription: "Terraform Workspace.",
				Attributes: map[string]schema.Attribute{
					"workspace": schema.StringAttribute{
						MarkdownDescription: "The Terraform workspace name.",
						Required:            true,
					},
				},
				Optional: true,
			},
		},
	}
}

func (r *projectInventoryResource) ConfigValidators(ctx context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{
		resourcevalidator.ExactlyOneOf(
			path.MatchRoot("static"),
			path.MatchRoot("static_yaml"),
			path.MatchRoot("file"),
			path.MatchRoot("terraform_workspace"),
		),
	}
}

const (
	InventoryStatic             string = "static"
	InventoryStaticYaml         string = "static-yaml"
	InventoryFile               string = "file"
	InventoryTerraformWorkspace string = "terraform-workspace"
)

type projectInventoryModel struct {
	ID                 types.Int64                       `tfsdk:"id"`
	ProjectID          types.Int64                       `tfsdk:"project_id"`
	Name               types.String                      `tfsdk:"name"`
	SSHKeyID           types.Int64                       `tfsdk:"ssh_key_id"`
	Static             *inventoryStaticModel             `tfsdk:"static"`
	StaticYaml         *inventoryStaticYamlModel         `tfsdk:"static_yaml"`
	File               *inventoryFileModel               `tfsdk:"file"`
	TerraformWorkspace *inventoryTerraformWorkspaceModel `tfsdk:"terraform_workspace"`
}

type inventoryStaticModel struct {
	Inventory   types.String `tfsdk:"inventory"`
	BecomeKeyID types.Int64  `tfsdk:"become_key_id"`
}

type inventoryStaticYamlModel struct {
	Inventory   types.String `tfsdk:"inventory"`
	BecomeKeyID types.Int64  `tfsdk:"become_key_id"`
}

type inventoryFileModel struct {
	Path         types.String `tfsdk:"path"`
	RepositoryID types.Int64  `tfsdk:"repository_id"`
	BecomeKeyID  types.Int64  `tfsdk:"become_key_id"`
}

type inventoryTerraformWorkspaceModel struct {
	Workspace types.String `tfsdk:"workspace"`
}

func convertProjectInventoryModelToInventoryRequest(inventory projectInventoryModel) *models.InventoryRequest {
	model := models.InventoryRequest{
		ProjectID: inventory.ProjectID.ValueInt64(),
		Name:      inventory.Name.ValueString(),
		SSHKeyID:  inventory.SSHKeyID.ValueInt64(),
	}
	if !inventory.ID.IsNull() && !inventory.ID.IsUnknown() {
		model.ID = inventory.ID.ValueInt64()
	}
	if inventory.Static != nil {
		model.Type = InventoryStatic
		model.Inventory = inventory.Static.Inventory.ValueString()
		model.BecomeKeyID = inventory.Static.BecomeKeyID.ValueInt64()
	} else if inventory.StaticYaml != nil {
		model.Type = InventoryStaticYaml
		model.Inventory = inventory.StaticYaml.Inventory.ValueString()
		model.BecomeKeyID = inventory.StaticYaml.BecomeKeyID.ValueInt64()
	} else if inventory.File != nil {
		model.Type = InventoryFile
		model.Inventory = inventory.File.Path.ValueString()
		model.BecomeKeyID = inventory.File.BecomeKeyID.ValueInt64()
		model.RepositoryID = inventory.File.RepositoryID.ValueInt64()
	} else if inventory.TerraformWorkspace != nil {
		model.Type = InventoryTerraformWorkspace
		model.Inventory = inventory.TerraformWorkspace.Workspace.ValueString()
	}

	return &model
}

func convertInventoryResponseToProjectInventoryModel(inventory *models.Inventory) projectInventoryModel {
	model := projectInventoryModel{
		ID:        types.Int64Value(inventory.ID),
		ProjectID: types.Int64Value(inventory.ProjectID),
		Name:      types.StringValue(inventory.Name),
		SSHKeyID:  types.Int64Value(inventory.SSHKeyID),
	}

	switch inventory.Type {
	case InventoryStatic:
		model.Static = &inventoryStaticModel{
			Inventory: types.StringValue(inventory.Inventory),
		}
		if inventory.BecomeKeyID != 0 {
			model.Static.BecomeKeyID = types.Int64Value(inventory.BecomeKeyID)
		} else {
			model.Static.BecomeKeyID = types.Int64Null()
		}
	case InventoryStaticYaml:
		model.StaticYaml = &inventoryStaticYamlModel{
			Inventory: types.StringValue(inventory.Inventory),
		}
		if inventory.BecomeKeyID != 0 {
			model.StaticYaml.BecomeKeyID = types.Int64Value(inventory.BecomeKeyID)
		} else {
			model.StaticYaml.BecomeKeyID = types.Int64Null()
		}
	case InventoryFile:
		model.File = &inventoryFileModel{
			Path: types.StringValue(inventory.Inventory),
		}
		if inventory.BecomeKeyID != 0 {
			model.File.BecomeKeyID = types.Int64Value(inventory.BecomeKeyID)
		} else {
			model.File.BecomeKeyID = types.Int64Null()
		}
		if inventory.RepositoryID != 0 {
			model.File.RepositoryID = types.Int64Value(inventory.RepositoryID)
		} else {
			model.File.RepositoryID = types.Int64Null()
		}
	case InventoryTerraformWorkspace:
		model.TerraformWorkspace = &inventoryTerraformWorkspaceModel{
			Workspace: types.StringValue(inventory.Inventory),
		}
	}
	return model
}

// Create creates the resource and sets the initial Terraform state.
func (r *projectInventoryResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan projectInventoryModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := r.client.Project.PostProjectProjectIDInventory(&project.PostProjectProjectIDInventoryParams{
		ProjectID: plan.ProjectID.ValueInt64(),
		Inventory: convertProjectInventoryModelToInventoryRequest(plan),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating SemaphoreUI Project Inventory",
			"Could not create project inventory, unexpected error: "+err.Error(),
		)
		return
	}
	plan = convertInventoryResponseToProjectInventoryModel(response.Payload)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *projectInventoryResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state projectInventoryModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := r.client.Project.GetProjectProjectIDInventoryInventoryID(&project.GetProjectProjectIDInventoryInventoryIDParams{
		ProjectID:   state.ProjectID.ValueInt64(),
		InventoryID: state.ID.ValueInt64(),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Project Inventory",
			"Could not read project inventory, unexpected error: "+err.Error(),
		)
		return
	}
	state = convertInventoryResponseToProjectInventoryModel(response.Payload)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *projectInventoryResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan projectInventoryModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.Project.PutProjectProjectIDInventoryInventoryID(&project.PutProjectProjectIDInventoryInventoryIDParams{
		ProjectID:   plan.ProjectID.ValueInt64(),
		InventoryID: plan.ID.ValueInt64(),
		Inventory:   convertProjectInventoryModelToInventoryRequest(plan),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating SemaphoreUI Project Inventory",
			"Could not update project inventory, unexpected error: "+err.Error(),
		)
		return
	}

	// Fetch updated values as PutProjectProjectIDInventoryInventoryID does not return updated project inventory
	response, err := r.client.Project.GetProjectProjectIDInventoryInventoryID(&project.GetProjectProjectIDInventoryInventoryIDParams{
		ProjectID:   plan.ProjectID.ValueInt64(),
		InventoryID: plan.ID.ValueInt64(),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Project Inventory",
			"Could not read project inventory, unexpected error: "+err.Error(),
		)
		return
	}
	plan = convertInventoryResponseToProjectInventoryModel(response.Payload)

	// Update resource state with updated project
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *projectInventoryResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state projectInventoryModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing resource
	_, err := r.client.Project.DeleteProjectProjectIDInventoryInventoryID(&project.DeleteProjectProjectIDInventoryInventoryIDParams{
		ProjectID:   state.ProjectID.ValueInt64(),
		InventoryID: state.ID.ValueInt64(),
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting SemaphoreUI Project Inventory",
			"Could not delete project inventory, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *projectInventoryResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	fields, err := parseImportFields(req.ID, []string{"project", "inventory"})
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid ProjectInventory Import ID",
			"Could not parse import ID: "+err.Error(),
		)
		return
	}

	response, err := r.client.Project.GetProjectProjectIDInventoryInventoryID(&project.GetProjectProjectIDInventoryInventoryIDParams{
		ProjectID:   fields["project"],
		InventoryID: fields["inventory"],
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Project Inventory",
			"Could not read project inventory, unexpected error: "+err.Error(),
		)
		return
	}
	state := convertInventoryResponseToProjectInventoryModel(response.Payload)

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
