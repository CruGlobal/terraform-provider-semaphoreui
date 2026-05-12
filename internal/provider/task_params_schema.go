package provider

import (
	"context"

	schemaD "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	schemaR "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	superschema "github.com/orange-cloudavenue/terraform-plugin-framework-superschema"

	"terraform-provider-semaphoreui/semaphoreui/models"
)

// TaskParamsModel mirrors the SemaphoreUI TaskPrams JSON object. Used by
// both project templates and project integrations.
type TaskParamsModel struct {
	Arguments   types.String              `tfsdk:"arguments"`
	Environment types.String              `tfsdk:"environment"`
	GitBranch   types.String              `tfsdk:"git_branch"`
	Message     types.String              `tfsdk:"message"`
	Ansible     *AnsibleTaskParamsModel   `tfsdk:"ansible"`
	Terraform   *TerraformTaskParamsModel `tfsdk:"terraform"`
}

type AnsibleTaskParamsModel struct {
	Tags     types.List `tfsdk:"tags"`
	SkipTags types.List `tfsdk:"skip_tags"`
	Limit    types.List `tfsdk:"limit"`
	Debug    types.Bool `tfsdk:"debug"`
	Diff     types.Bool `tfsdk:"diff"`
	DryRun   types.Bool `tfsdk:"dry_run"`
}

type TerraformTaskParamsModel struct {
	AutoApprove types.Bool `tfsdk:"auto_approve"`
	Destroy     types.Bool `tfsdk:"destroy"`
	Plan        types.Bool `tfsdk:"plan"`
	Upgrade     types.Bool `tfsdk:"upgrade"`
}

// TaskParamsAttribute returns the shared `task_params` attribute used by
// project_template and project_integration resources / data sources.
func TaskParamsAttribute() superschema.Attribute {
	return superschema.SingleNestedAttribute{
		Common: &schemaR.SingleNestedAttribute{
			MarkdownDescription: "Default task parameters applied when this template or integration runs a task.",
		},
		Resource: &schemaR.SingleNestedAttribute{
			Optional: true,
		},
		DataSource: &schemaD.SingleNestedAttribute{
			Computed: true,
		},
		Attributes: map[string]superschema.Attribute{
			"arguments": superschema.StringAttribute{
				Common: &schemaR.StringAttribute{
					MarkdownDescription: "JSON-encoded array of extra command-line arguments passed to the task runner (e.g. `\"[\\\"-vvv\\\"]\"`).",
				},
				Resource: &schemaR.StringAttribute{
					Optional: true,
				},
				DataSource: &schemaD.StringAttribute{
					Computed: true,
				},
			},
			"environment": superschema.StringAttribute{
				Common: &schemaR.StringAttribute{
					MarkdownDescription: "JSON-encoded object of environment variables exposed to the task.",
				},
				Resource: &schemaR.StringAttribute{
					Optional: true,
				},
				DataSource: &schemaD.StringAttribute{
					Computed: true,
				},
			},
			"git_branch": superschema.StringAttribute{
				Common: &schemaR.StringAttribute{
					MarkdownDescription: "Override the repository branch checked out for this task.",
				},
				Resource: &schemaR.StringAttribute{
					Optional: true,
				},
				DataSource: &schemaD.StringAttribute{
					Computed: true,
				},
			},
			"message": superschema.StringAttribute{
				Common: &schemaR.StringAttribute{
					MarkdownDescription: "Optional commit-style message recorded with each task run.",
				},
				Resource: &schemaR.StringAttribute{
					Optional: true,
				},
				DataSource: &schemaD.StringAttribute{
					Computed: true,
				},
			},
			"ansible": superschema.SingleNestedAttribute{
				Common: &schemaR.SingleNestedAttribute{
					MarkdownDescription: "Ansible-specific task parameters. Use this when `app` is `ansible`.",
				},
				Resource: &schemaR.SingleNestedAttribute{
					Optional: true,
				},
				DataSource: &schemaD.SingleNestedAttribute{
					Computed: true,
				},
				Attributes: map[string]superschema.Attribute{
					"tags": superschema.ListAttribute{
						Common: &schemaR.ListAttribute{
							MarkdownDescription: "Ansible tags to run (`--tags`).",
							ElementType:         types.StringType,
						},
						Resource: &schemaR.ListAttribute{
							Optional: true,
						},
						DataSource: &schemaD.ListAttribute{
							Computed: true,
						},
					},
					"skip_tags": superschema.ListAttribute{
						Common: &schemaR.ListAttribute{
							MarkdownDescription: "Ansible tags to skip (`--skip-tags`).",
							ElementType:         types.StringType,
						},
						Resource: &schemaR.ListAttribute{
							Optional: true,
						},
						DataSource: &schemaD.ListAttribute{
							Computed: true,
						},
					},
					"limit": superschema.ListAttribute{
						Common: &schemaR.ListAttribute{
							MarkdownDescription: "Ansible hosts to limit the run to (`--limit`).",
							ElementType:         types.StringType,
						},
						Resource: &schemaR.ListAttribute{
							Optional: true,
						},
						DataSource: &schemaD.ListAttribute{
							Computed: true,
						},
					},
					"debug": superschema.BoolAttribute{
						Common: &schemaR.BoolAttribute{
							MarkdownDescription: "Run Ansible with `-vvvv` debug output.",
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
					"diff": superschema.BoolAttribute{
						Common: &schemaR.BoolAttribute{
							MarkdownDescription: "Show file diffs for changes Ansible makes (`--diff`).",
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
					"dry_run": superschema.BoolAttribute{
						Common: &schemaR.BoolAttribute{
							MarkdownDescription: "Run Ansible in check mode (`--check`).",
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
				},
			},
			"terraform": superschema.SingleNestedAttribute{
				Common: &schemaR.SingleNestedAttribute{
					MarkdownDescription: "Terraform / OpenTofu-specific task parameters. Use this when `app` is `terraform` or `tofu`.",
				},
				Resource: &schemaR.SingleNestedAttribute{
					Optional: true,
				},
				DataSource: &schemaD.SingleNestedAttribute{
					Computed: true,
				},
				Attributes: map[string]superschema.Attribute{
					"auto_approve": superschema.BoolAttribute{
						Common: &schemaR.BoolAttribute{
							MarkdownDescription: "Run with `-auto-approve`.",
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
					"destroy": superschema.BoolAttribute{
						Common: &schemaR.BoolAttribute{
							MarkdownDescription: "Run a destroy (`terraform destroy` / `tofu destroy`).",
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
					"plan": superschema.BoolAttribute{
						Common: &schemaR.BoolAttribute{
							MarkdownDescription: "Run plan-only (no apply).",
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
					"upgrade": superschema.BoolAttribute{
						Common: &schemaR.BoolAttribute{
							MarkdownDescription: "Pass `-upgrade` to `terraform init` / `tofu init`.",
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
				},
			},
		},
	}
}

// convertTaskParamsModelToTaskPrams converts a Terraform model into the
// generated client's TaskPrams. Returns nil if the model is unset, so callers
// can omit task_params from requests cleanly.
func convertTaskParamsModelToTaskPrams(ctx context.Context, model *TaskParamsModel) *models.TaskPrams {
	if model == nil {
		return nil
	}
	out := &models.TaskPrams{
		Arguments:   model.Arguments.ValueString(),
		Environment: model.Environment.ValueString(),
		GitBranch:   model.GitBranch.ValueString(),
		Message:     model.Message.ValueString(),
	}
	if model.Ansible != nil {
		out.Params.AnsibleTaskParams = models.AnsibleTaskParams{
			Tags:     stringListToSlice(ctx, model.Ansible.Tags),
			SkipTags: stringListToSlice(ctx, model.Ansible.SkipTags),
			Limit:    stringListToSlice(ctx, model.Ansible.Limit),
			Debug:    model.Ansible.Debug.ValueBool(),
			Diff:     model.Ansible.Diff.ValueBool(),
			DryRun:   model.Ansible.DryRun.ValueBool(),
		}
	}
	if model.Terraform != nil {
		out.Params.TerraformTaskParams = models.TerraformTaskParams{
			AutoApprove: model.Terraform.AutoApprove.ValueBool(),
			Destroy:     model.Terraform.Destroy.ValueBool(),
			Plan:        model.Terraform.Plan.ValueBool(),
			Upgrade:     model.Terraform.Upgrade.ValueBool(),
		}
	}
	return out
}

// convertTaskPramsToTaskParamsModel converts the generated client's TaskPrams
// back into the Terraform model shape. Returns nil if the input is nil.
//
// Note on round-trip: the API omits fields that weren't sent (JSON
// `omitempty`-style behavior), so a TaskPrams returned by GET only reflects
// what was previously written. Empty slices show up as Terraform-null lists,
// and false booleans aren't always echoed. Callers should treat the returned
// model as authoritative for what's stored server-side.
func convertTaskPramsToTaskParamsModel(ctx context.Context, in *models.TaskPrams) *TaskParamsModel {
	if in == nil {
		return nil
	}
	out := &TaskParamsModel{
		Arguments:   stringOrNull(in.Arguments),
		Environment: stringOrNull(in.Environment),
		GitBranch:   stringOrNull(in.GitBranch),
		Message:     stringOrNull(in.Message),
	}
	if !ansibleTaskParamsEmpty(in.Params.AnsibleTaskParams) {
		out.Ansible = &AnsibleTaskParamsModel{
			Tags:     sliceToStringList(ctx, in.Params.Tags),
			SkipTags: sliceToStringList(ctx, in.Params.SkipTags),
			Limit:    sliceToStringList(ctx, in.Params.Limit),
			Debug:    types.BoolValue(in.Params.Debug),
			Diff:     types.BoolValue(in.Params.Diff),
			DryRun:   types.BoolValue(in.Params.DryRun),
		}
	}
	if !terraformTaskParamsEmpty(in.Params.TerraformTaskParams) {
		out.Terraform = &TerraformTaskParamsModel{
			AutoApprove: types.BoolValue(in.Params.AutoApprove),
			Destroy:     types.BoolValue(in.Params.Destroy),
			Plan:        types.BoolValue(in.Params.Plan),
			Upgrade:     types.BoolValue(in.Params.Upgrade),
		}
	}
	return out
}

func ansibleTaskParamsEmpty(p models.AnsibleTaskParams) bool {
	return len(p.Tags) == 0 && len(p.SkipTags) == 0 && len(p.Limit) == 0 &&
		!p.Debug && !p.Diff && !p.DryRun
}

func terraformTaskParamsEmpty(p models.TerraformTaskParams) bool {
	return !p.AutoApprove && !p.Destroy && !p.Plan && !p.Upgrade
}

func stringOrNull(s string) types.String {
	if s == "" {
		return types.StringNull()
	}
	return types.StringValue(s)
}

func stringListToSlice(ctx context.Context, list types.List) []string {
	if list.IsNull() || list.IsUnknown() {
		return nil
	}
	var out []string
	list.ElementsAs(ctx, &out, false)
	return out
}

func sliceToStringList(ctx context.Context, s []string) types.List {
	if len(s) == 0 {
		return types.ListNull(types.StringType)
	}
	out, _ := types.ListValueFrom(ctx, types.StringType, s)
	return out
}
