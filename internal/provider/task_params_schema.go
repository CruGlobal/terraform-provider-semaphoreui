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

// TaskParamsModel is the Terraform shape of task_params, shared by project
// templates and project integrations. Templates send it flat
// (TemplateTaskParams); integrations send it nested (TaskPrams).
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
			MarkdownDescription: "Default task parameters applied when this template or integration runs a task. On a template, SemaphoreUI only uses `ansible.limit`, `ansible.tags`, `ansible.skip_tags` and `terraform.auto_approve`; it stores the other fields but ignores them, so `terraform.plan` does not stop an auto-approved apply.",
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
// generated client's TaskPrams, the shape integrations use (app params nested
// under `params`). Returns nil if the model is unset, so callers can omit
// task_params from requests cleanly.
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
	out.Params.AnsibleTaskParams = ansibleTaskParamsFromModel(ctx, model.Ansible)
	out.Params.TerraformTaskParams = terraformTaskParamsFromModel(model.Terraform)
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
	return &TaskParamsModel{
		Arguments:   stringOrNull(in.Arguments),
		Environment: stringOrNull(in.Environment),
		GitBranch:   stringOrNull(in.GitBranch),
		Message:     stringOrNull(in.Message),
		Ansible:     ansibleTaskParamsToModel(ctx, in.Params.AnsibleTaskParams),
		Terraform:   terraformTaskParamsToModel(in.Params.TerraformTaskParams),
	}
}

// convertTaskParamsModelToTemplateTaskParams converts a Terraform model into
// the generated client's TemplateTaskParams. SemaphoreUI reads a template's
// task_params flat (limit/tags/skip_tags/auto_approve at the top level), unlike
// the nested TaskPrams shape integrations use. Returns nil if the model is
// unset.
func convertTaskParamsModelToTemplateTaskParams(ctx context.Context, model *TaskParamsModel) *models.TemplateTaskParams {
	if model == nil {
		return nil
	}
	return &models.TemplateTaskParams{
		Arguments:           model.Arguments.ValueString(),
		Environment:         model.Environment.ValueString(),
		GitBranch:           model.GitBranch.ValueString(),
		Message:             model.Message.ValueString(),
		AnsibleTaskParams:   ansibleTaskParamsFromModel(ctx, model.Ansible),
		TerraformTaskParams: terraformTaskParamsFromModel(model.Terraform),
	}
}

// convertTemplateTaskParamsToTaskParamsModel converts a template's flat
// task_params back into the Terraform model shape. Returns nil if the input
// is nil. Values an older provider wrote under a nested `params` key are not
// read, so the next plan shows them as a change and the apply rewrites them
// where SemaphoreUI looks.
func convertTemplateTaskParamsToTaskParamsModel(ctx context.Context, in *models.TemplateTaskParams) *TaskParamsModel {
	if in == nil {
		return nil
	}
	return &TaskParamsModel{
		Arguments:   stringOrNull(in.Arguments),
		Environment: stringOrNull(in.Environment),
		GitBranch:   stringOrNull(in.GitBranch),
		Message:     stringOrNull(in.Message),
		Ansible:     ansibleTaskParamsToModel(ctx, in.AnsibleTaskParams),
		Terraform:   terraformTaskParamsToModel(in.TerraformTaskParams),
	}
}

func ansibleTaskParamsFromModel(ctx context.Context, m *AnsibleTaskParamsModel) models.AnsibleTaskParams {
	if m == nil {
		return models.AnsibleTaskParams{}
	}
	return models.AnsibleTaskParams{
		Tags:     stringListToSlice(ctx, m.Tags),
		SkipTags: stringListToSlice(ctx, m.SkipTags),
		Limit:    stringListToSlice(ctx, m.Limit),
		Debug:    m.Debug.ValueBool(),
		Diff:     m.Diff.ValueBool(),
		DryRun:   m.DryRun.ValueBool(),
	}
}

func terraformTaskParamsFromModel(m *TerraformTaskParamsModel) models.TerraformTaskParams {
	if m == nil {
		return models.TerraformTaskParams{}
	}
	return models.TerraformTaskParams{
		AutoApprove: m.AutoApprove.ValueBool(),
		Destroy:     m.Destroy.ValueBool(),
		Plan:        m.Plan.ValueBool(),
		Upgrade:     m.Upgrade.ValueBool(),
	}
}

func ansibleTaskParamsToModel(ctx context.Context, p models.AnsibleTaskParams) *AnsibleTaskParamsModel {
	if ansibleTaskParamsEmpty(p) {
		return nil
	}
	return &AnsibleTaskParamsModel{
		Tags:     sliceToStringList(ctx, p.Tags),
		SkipTags: sliceToStringList(ctx, p.SkipTags),
		Limit:    sliceToStringList(ctx, p.Limit),
		Debug:    types.BoolValue(p.Debug),
		Diff:     types.BoolValue(p.Diff),
		DryRun:   types.BoolValue(p.DryRun),
	}
}

func terraformTaskParamsToModel(p models.TerraformTaskParams) *TerraformTaskParamsModel {
	if terraformTaskParamsEmpty(p) {
		return nil
	}
	return &TerraformTaskParamsModel{
		AutoApprove: types.BoolValue(p.AutoApprove),
		Destroy:     types.BoolValue(p.Destroy),
		Plan:        types.BoolValue(p.Plan),
		Upgrade:     types.BoolValue(p.Upgrade),
	}
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
