package provider

import (
	schemaR "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	superschema "github.com/orange-cloudavenue/terraform-plugin-framework-superschema"

	"terraform-provider-semaphoreui/internal/stringvalidator"
)

type (
	ProjectTaskModel struct {
		ID                types.Int64  `tfsdk:"id"`
		ProjectID         types.Int64  `tfsdk:"project_id"`
		TemplateID        types.Int64  `tfsdk:"template_id"`
		InventoryID       types.Int64  `tfsdk:"inventory_id"`
		Playbook          types.String `tfsdk:"playbook"`
		Environment       types.String `tfsdk:"environment"`
		Limit             types.String `tfsdk:"limit"`
		GitBranch         types.String `tfsdk:"git_branch"`
		Message           types.String `tfsdk:"message"`
		Arguments         types.String `tfsdk:"arguments"`
		Debug             types.Bool   `tfsdk:"debug"`
		DryRun            types.Bool   `tfsdk:"dry_run"`
		Diff              types.Bool   `tfsdk:"diff"`
		Triggers          types.Map    `tfsdk:"triggers"`
		WaitForCompletion types.Bool   `tfsdk:"wait_for_completion"`
		Timeout           types.String `tfsdk:"timeout"`
		CaptureOutput     types.Bool   `tfsdk:"capture_output"`
		Status            types.String `tfsdk:"status"`
		Output            types.String `tfsdk:"output"`
	}
)

// ProjectTaskSchema describes semaphoreui_project_task. There is deliberately
// no matching data source: a task is a one-shot job run, and everything worth
// reading back (status, output) is already exposed on the resource.
func ProjectTaskSchema() superschema.Schema {
	return superschema.Schema{
		Common: superschema.SchemaDetails{
			MarkdownDescription: "The project task",
		},
		Resource: superschema.SchemaDetails{
			MarkdownDescription: "resource runs a template as a job in SemaphoreUI.\n\n" +
				"Unlike the other resources in this provider, a task is not a persistent object you edit — it is a\n" +
				"single job run. Creating the resource queues the job; by default Terraform then waits for it to\n" +
				"finish and fails the apply if the job does not end in `success`.\n\n" +
				"Every input is immutable: changing any of them replaces the resource, which queues a **new** job.\n" +
				"Use the `triggers` map to re-run a job when some unrelated value changes (an image tag, a commit\n" +
				"SHA, ...).\n\n" +
				"Destroying the resource stops the job if it is still running, but never deletes the run from\n" +
				"SemaphoreUI's task history.",
		},
		Attributes: map[string]superschema.Attribute{
			"id": superschema.Int64Attribute{
				Resource: &schemaR.Int64Attribute{
					MarkdownDescription: "The task ID assigned by SemaphoreUI.",
					Computed:            true,
					PlanModifiers:       []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
				},
			},
			"project_id": superschema.Int64Attribute{
				Resource: &schemaR.Int64Attribute{
					MarkdownDescription: "The project ID that the task runs in.",
					Required:            true,
					PlanModifiers:       []planmodifier.Int64{int64planmodifier.RequiresReplace()},
				},
			},
			"template_id": superschema.Int64Attribute{
				Resource: &schemaR.Int64Attribute{
					MarkdownDescription: "The template ID to run.",
					Required:            true,
					PlanModifiers:       []planmodifier.Int64{int64planmodifier.RequiresReplace()},
				},
			},
			"inventory_id": superschema.Int64Attribute{
				Resource: &schemaR.Int64Attribute{
					MarkdownDescription: "Override the inventory the template runs against. Defaults to the template's own inventory.",
					Optional:            true,
					PlanModifiers:       []planmodifier.Int64{int64planmodifier.RequiresReplace()},
				},
			},
			"playbook": superschema.StringAttribute{
				Resource: &schemaR.StringAttribute{
					MarkdownDescription: "Override the playbook file the template runs. Defaults to the template's own playbook.",
					Optional:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
				},
			},
			"environment": superschema.StringAttribute{
				Resource: &schemaR.StringAttribute{
					MarkdownDescription: "JSON-encoded object of extra variables passed to the run, e.g. `jsonencode({ app_version = \"1.2.3\" })`.",
					Optional:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
				},
			},
			"limit": superschema.StringAttribute{
				Resource: &schemaR.StringAttribute{
					MarkdownDescription: "Restrict the run to a subset of hosts (Ansible `--limit`).",
					Optional:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
				},
			},
			"git_branch": superschema.StringAttribute{
				Resource: &schemaR.StringAttribute{
					MarkdownDescription: "Override the repository branch checked out for this run.",
					Optional:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
				},
			},
			"message": superschema.StringAttribute{
				Resource: &schemaR.StringAttribute{
					MarkdownDescription: "Free-form message recorded alongside the run in the task history.",
					Optional:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
				},
			},
			"arguments": superschema.StringAttribute{
				Resource: &schemaR.StringAttribute{
					MarkdownDescription: "JSON-encoded array of extra command-line arguments passed to the task runner, e.g. `jsonencode([\"-vvv\"])`.",
					Optional:            true,
					PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
				},
			},
			"debug": superschema.BoolAttribute{
				Resource: &schemaR.BoolAttribute{
					MarkdownDescription: "Run with verbose debug output.",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
					PlanModifiers:       []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
				},
			},
			"dry_run": superschema.BoolAttribute{
				Resource: &schemaR.BoolAttribute{
					MarkdownDescription: "Run in check mode without applying changes (Ansible `--check`).",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
					PlanModifiers:       []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
				},
			},
			"diff": superschema.BoolAttribute{
				Resource: &schemaR.BoolAttribute{
					MarkdownDescription: "Show file diffs for the changes the run makes (Ansible `--diff`).",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
					PlanModifiers:       []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
				},
			},
			"triggers": superschema.MapAttribute{
				Resource: &schemaR.MapAttribute{
					MarkdownDescription: "Arbitrary key/value pairs that force a new task run when they change. Nothing is sent to SemaphoreUI; this exists purely to re-run the job on demand.",
					Optional:            true,
					ElementType:         types.StringType,
					PlanModifiers:       []planmodifier.Map{mapplanmodifier.RequiresReplace()},
				},
			},
			"wait_for_completion": superschema.BoolAttribute{
				Resource: &schemaR.BoolAttribute{
					MarkdownDescription: "Wait for the task to reach a terminal state before the apply returns. When `true` (the default), a task ending in anything other than `success` fails the apply. When `false`, the apply returns as soon as the job is queued and `status` reflects that initial state. Only takes effect when the task is created — flipping it later never re-runs or re-waits on an existing task.",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(true),
				},
			},
			"timeout": superschema.StringAttribute{
				Resource: &schemaR.StringAttribute{
					MarkdownDescription: "How long to wait for the task to finish, as a Go duration string. Ignored when `wait_for_completion` is `false`.",
					Optional:            true,
					Computed:            true,
					Default:             stringdefault.StaticString("30m"),
					Validators: []validator.String{
						stringvalidator.Duration(),
					},
				},
			},
			"capture_output": superschema.BoolAttribute{
				Resource: &schemaR.BoolAttribute{
					MarkdownDescription: "Store the task's console output in the `output` attribute. Job logs can be large and Terraform state is not a good place for them, so this is opt-in. Regardless of this setting, the tail of the output is included in the error message when a job fails.",
					Optional:            true,
					Computed:            true,
					Default:             booldefault.StaticBool(false),
				},
			},
			"status": superschema.StringAttribute{
				Resource: &schemaR.StringAttribute{
					MarkdownDescription: "The task status reported by SemaphoreUI, e.g. `waiting`, `running`, `success`, `error`, `stopped`.",
					Computed:            true,
				},
			},
			"output": superschema.StringAttribute{
				Resource: &schemaR.StringAttribute{
					MarkdownDescription: "The task's console output, newline-joined. Null unless `capture_output` is `true`.",
					Computed:            true,
				},
			},
		},
	}
}
