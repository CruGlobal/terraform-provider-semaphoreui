package provider

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	apiclient "terraform-provider-semaphoreui/semaphoreui/client"
	"terraform-provider-semaphoreui/semaphoreui/client/task"
	"terraform-provider-semaphoreui/semaphoreui/models"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &projectTaskResource{}
	_ resource.ResourceWithConfigure   = &projectTaskResource{}
	_ resource.ResourceWithImportState = &projectTaskResource{}
)

const (
	// taskPollInterval is how often a running task is polled while waiting for
	// it to finish.
	taskPollInterval = 3 * time.Second
	// taskFailureOutputLines is how many trailing output lines are quoted in the
	// diagnostic when a task fails.
	taskFailureOutputLines = 50
)

// taskTerminalStatuses are the SemaphoreUI task statuses that mean the job is
// done and will not change again. Anything not listed here is treated as
// still-in-flight, so an unknown status from a newer Semaphore release makes the
// provider keep polling (until timeout) rather than report a bogus result.
var taskTerminalStatuses = map[string]bool{
	"success":  true,
	"error":    true,
	"stopped":  true,
	"rejected": true,
}

// taskSuccessStatus is the only status considered a successful run.
const taskSuccessStatus = "success"

func NewProjectTaskResource() resource.Resource {
	return &projectTaskResource{}
}

type projectTaskResource struct {
	client *apiclient.SemaphoreUI
}

func (r *projectTaskResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *projectTaskResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project_task"
}

func (r *projectTaskResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ProjectTaskSchema().GetResource(ctx)
}

func convertProjectTaskModelToTaskBody(plan ProjectTaskModel) task.PostProjectProjectIDTasksBody {
	return task.PostProjectProjectIDTasksBody{
		TemplateID:  plan.TemplateID.ValueInt64(),
		InventoryID: plan.InventoryID.ValueInt64(),
		Playbook:    plan.Playbook.ValueString(),
		Environment: plan.Environment.ValueString(),
		Limit:       plan.Limit.ValueString(),
		GitBranch:   plan.GitBranch.ValueString(),
		Message:     plan.Message.ValueString(),
		Arguments:   plan.Arguments.ValueString(),
		Debug:       plan.Debug.ValueBool(),
		DryRun:      plan.DryRun.ValueBool(),
		Diff:        plan.Diff.ValueBool(),
	}
}

// apiErrorCode returns the HTTP status carried by a generated-client error, or 0
// if the error is not an API error (transport failure, cancelled context, ...).
func apiErrorCode(err error) int {
	var apiErr *runtime.APIError
	if errors.As(err, &apiErr) {
		return apiErr.Code
	}
	return 0
}

// getTask fetches a single task. A nil task with a nil error means the task no
// longer exists server-side.
func (r *projectTaskResource) getTask(ctx context.Context, projectID, taskID int64) (*models.Task, error) {
	response, err := r.client.Task.GetProjectProjectIDTasksTaskID(&task.GetProjectProjectIDTasksTaskIDParams{
		ProjectID: projectID,
		TaskID:    taskID,
		Context:   ctx,
	}, nil)
	if err != nil {
		if apiErrorCode(err) == 404 {
			return nil, nil
		}
		return nil, err
	}
	return response.Payload, nil
}

// taskOutput returns the task's console output as a single newline-joined
// string. Output is best-effort: a failure to read it is reported to the caller
// but is never fatal on its own.
func (r *projectTaskResource) taskOutput(ctx context.Context, projectID, taskID int64) (string, error) {
	response, err := r.client.Task.GetProjectProjectIDTasksTaskIDOutput(&task.GetProjectProjectIDTasksTaskIDOutputParams{
		ProjectID: projectID,
		TaskID:    taskID,
		Context:   ctx,
	}, nil)
	if err != nil {
		return "", err
	}

	lines := make([]string, 0, len(response.Payload))
	for _, entry := range response.Payload {
		if entry == nil {
			continue
		}
		lines = append(lines, entry.Output)
	}
	return strings.Join(lines, "\n"), nil
}

// outputValue resolves the `output` attribute: null unless capture_output is on.
func (r *projectTaskResource) outputValue(ctx context.Context, capture bool, projectID, taskID int64) (types.String, error) {
	if !capture {
		return types.StringNull(), nil
	}
	output, err := r.taskOutput(ctx, projectID, taskID)
	if err != nil {
		return types.StringNull(), err
	}
	return types.StringValue(output), nil
}

// failureDetail builds the error message for a task that did not succeed,
// quoting the tail of its output so the failure is diagnosable without leaving
// Terraform.
func (r *projectTaskResource) failureDetail(ctx context.Context, projectID, taskID int64, status string) string {
	detail := fmt.Sprintf("Task %d finished with status %q instead of %q.", taskID, status, taskSuccessStatus)

	output, err := r.taskOutput(ctx, projectID, taskID)
	if err != nil {
		return detail + "\n\nThe task output could not be read: " + err.Error()
	}
	if strings.TrimSpace(output) == "" {
		return detail + "\n\nThe task produced no output."
	}

	lines := strings.Split(output, "\n")
	truncated := false
	if len(lines) > taskFailureOutputLines {
		lines = lines[len(lines)-taskFailureOutputLines:]
		truncated = true
	}
	tail := strings.Join(lines, "\n")
	if truncated {
		return fmt.Sprintf("%s\n\nLast %d lines of task output:\n%s", detail, taskFailureOutputLines, tail)
	}
	return detail + "\n\nTask output:\n" + tail
}

// waitForTask polls the task until it reaches a terminal status, the timeout
// elapses, or the context is cancelled.
func (r *projectTaskResource) waitForTask(ctx context.Context, projectID, taskID int64, timeout time.Duration) (*models.Task, error) {
	deadline := time.Now().Add(timeout)

	ticker := time.NewTicker(taskPollInterval)
	defer ticker.Stop()

	for {
		current, err := r.getTask(ctx, projectID, taskID)
		if err != nil {
			return nil, err
		}
		if current == nil {
			return nil, fmt.Errorf("task %d disappeared from project %d while waiting for it to finish", taskID, projectID)
		}
		if taskTerminalStatuses[current.Status] {
			return current, nil
		}

		if time.Now().After(deadline) {
			return current, fmt.Errorf(
				"timed out after %s waiting for task %d to finish (last status: %q)",
				timeout, taskID, current.Status,
			)
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
		}
	}
}

func (r *projectTaskResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan ProjectTaskModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	timeout, err := time.ParseDuration(plan.Timeout.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid SemaphoreUI Project Task Timeout",
			"Could not parse timeout, unexpected error: "+err.Error(),
		)
		return
	}

	projectID := plan.ProjectID.ValueInt64()
	response, err := r.client.Task.PostProjectProjectIDTasks(&task.PostProjectProjectIDTasksParams{
		ProjectID: projectID,
		Task:      convertProjectTaskModelToTaskBody(plan),
		Context:   ctx,
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating SemaphoreUI Project Task",
			"Could not start project task, unexpected error: "+err.Error(),
		)
		return
	}

	created := response.Payload
	taskID := created.ID
	plan.ID = types.Int64Value(taskID)
	plan.Status = types.StringValue(created.Status)

	// From here on the job exists server-side. Every remaining failure path must
	// still write state, otherwise Terraform loses track of a running job.
	if plan.WaitForCompletion.ValueBool() {
		finished, waitErr := r.waitForTask(ctx, projectID, taskID, timeout)
		if finished != nil {
			plan.Status = types.StringValue(finished.Status)
		}
		if waitErr != nil {
			plan.Output = types.StringNull()
			resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
			resp.Diagnostics.AddError(
				"Error Waiting For SemaphoreUI Project Task",
				"Could not wait for project task, unexpected error: "+waitErr.Error(),
			)
			return
		}
		if finished.Status != taskSuccessStatus {
			plan.Output = types.StringNull()
			resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
			resp.Diagnostics.AddError(
				"SemaphoreUI Project Task Failed",
				r.failureDetail(ctx, projectID, taskID, finished.Status),
			)
			return
		}
	} else if queued, getErr := r.getTask(ctx, projectID, taskID); getErr == nil && queued != nil {
		// The POST response does not reliably echo the queued status, so refresh
		// it. This is best-effort: the job is already running either way, and the
		// next Read will correct the status.
		plan.Status = types.StringValue(queued.Status)
	}

	output, err := r.outputValue(ctx, plan.CaptureOutput.ValueBool(), projectID, taskID)
	if err != nil {
		plan.Output = types.StringNull()
		resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Project Task Output",
			"Could not read project task output, unexpected error: "+err.Error(),
		)
		return
	}
	plan.Output = output

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *projectTaskResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state ProjectTaskModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	projectID := state.ProjectID.ValueInt64()
	taskID := state.ID.ValueInt64()

	current, err := r.getTask(ctx, projectID, taskID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Project Task",
			"Could not read project task, unexpected error: "+err.Error(),
		)
		return
	}
	if current == nil {
		// Drift: the task was deleted from the history out-of-band.
		resp.State.RemoveResource(ctx)
		return
	}

	// Only the observed attributes are refreshed. Everything else describes the
	// request that started this run: it is immutable and RequiresReplace, so
	// overwriting it from the response would only manufacture spurious drift.
	state.Status = types.StringValue(current.Status)

	output, err := r.outputValue(ctx, state.CaptureOutput.ValueBool(), projectID, taskID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Project Task Output",
			"Could not read project task output, unexpected error: "+err.Error(),
		)
		return
	}
	state.Output = output

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
//
// Every attribute describing the job itself is RequiresReplace, so Update is
// only ever reached when a client-side knob (wait_for_completion, timeout,
// capture_output) changed. No job is started or re-run here; the computed
// attributes are simply refreshed so they are known after apply.
func (r *projectTaskResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ProjectTaskModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.ID = state.ID
	projectID := plan.ProjectID.ValueInt64()
	taskID := plan.ID.ValueInt64()

	current, err := r.getTask(ctx, projectID, taskID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Project Task",
			"Could not read project task, unexpected error: "+err.Error(),
		)
		return
	}
	if current == nil {
		resp.Diagnostics.AddError(
			"SemaphoreUI Project Task Not Found",
			fmt.Sprintf("Task %d no longer exists in project %d. Remove it from state or taint it to run a new task.", taskID, projectID),
		)
		return
	}
	plan.Status = types.StringValue(current.Status)

	output, err := r.outputValue(ctx, plan.CaptureOutput.ValueBool(), projectID, taskID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Project Task Output",
			"Could not read project task output, unexpected error: "+err.Error(),
		)
		return
	}
	plan.Output = output

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete stops the task if it is still in flight. The run is deliberately left
// in SemaphoreUI's task history: that history is an audit trail, and destroying
// a Terraform resource should not erase the record that a job ran.
func (r *projectTaskResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state ProjectTaskModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	projectID := state.ProjectID.ValueInt64()
	taskID := state.ID.ValueInt64()

	current, err := r.getTask(ctx, projectID, taskID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Project Task",
			"Could not read project task before stopping it, unexpected error: "+err.Error(),
		)
		return
	}
	if current == nil || taskTerminalStatuses[current.Status] {
		// Already gone or already finished: nothing to stop.
		return
	}

	_, err = r.client.Task.PostProjectProjectIDTasksTaskIDStop(&task.PostProjectProjectIDTasksTaskIDStopParams{
		ProjectID: projectID,
		TaskID:    taskID,
		Body:      task.PostProjectProjectIDTasksTaskIDStopBody{Force: false},
		Context:   ctx,
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Stopping SemaphoreUI Project Task",
			"Could not stop project task, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *projectTaskResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	fields, err := parseImportFields(req.ID, []string{"project", "task"})
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid Project Task Import ID",
			"Could not parse import ID: "+err.Error(),
		)
		return
	}

	projectID := fields["project"]
	taskID := fields["task"]

	current, err := r.getTask(ctx, projectID, taskID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SemaphoreUI Project Task",
			"Could not read project task, unexpected error: "+err.Error(),
		)
		return
	}
	if current == nil {
		resp.Diagnostics.AddError(
			"SemaphoreUI Project Task Not Found",
			fmt.Sprintf("No task %d exists in project %d.", taskID, projectID),
		)
		return
	}

	// The import only knows what the API reports about the finished run; the
	// behavioral knobs get their schema defaults.
	model := ProjectTaskModel{
		ID:                types.Int64Value(current.ID),
		ProjectID:         types.Int64Value(projectID),
		TemplateID:        types.Int64Value(current.TemplateID),
		InventoryID:       int64OrNull(current.InventoryID),
		Playbook:          stringOrNull(current.Playbook),
		Environment:       stringOrNull(current.Environment),
		Limit:             stringOrNull(current.Limit),
		GitBranch:         stringOrNull(current.GitBranch),
		Message:           stringOrNull(current.Message),
		Arguments:         stringOrNull(current.Arguments),
		Debug:             types.BoolValue(current.Params.Debug),
		DryRun:            types.BoolValue(current.Params.DryRun),
		Diff:              types.BoolValue(current.Params.Diff),
		Triggers:          types.MapNull(types.StringType),
		WaitForCompletion: types.BoolValue(true),
		Timeout:           types.StringValue("30m"),
		CaptureOutput:     types.BoolValue(false),
		Status:            types.StringValue(current.Status),
		Output:            types.StringNull(),
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func int64OrNull(v int64) types.Int64 {
	if v == 0 {
		return types.Int64Null()
	}
	return types.Int64Value(v)
}
