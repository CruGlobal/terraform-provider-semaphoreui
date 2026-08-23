package provider

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/go-openapi/runtime"
	fwresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-semaphoreui/semaphoreui/client/task"
)

// TestProjectTaskSchema catches schema misconfiguration (conflicting
// Optional/Computed/Required flags, defaults on non-computed attributes, bad
// attribute names) without needing a live API.
func TestProjectTaskSchema(t *testing.T) {
	ctx := context.Background()

	schemaResponse := &fwresource.SchemaResponse{}
	NewProjectTaskResource().Schema(ctx, fwresource.SchemaRequest{}, schemaResponse)

	if schemaResponse.Diagnostics.HasError() {
		t.Fatalf("schema method diagnostics: %+v", schemaResponse.Diagnostics)
	}

	if diags := schemaResponse.Schema.ValidateImplementation(ctx); diags.HasError() {
		t.Fatalf("schema validation diagnostics: %+v", diags)
	}
}

// TestProjectTaskSchemaModelParity guards against the schema and
// ProjectTaskModel drifting apart, which surfaces at runtime as an opaque
// "mismatch between struct and object type" error.
func TestProjectTaskSchemaModelParity(t *testing.T) {
	ctx := context.Background()

	schemaResponse := &fwresource.SchemaResponse{}
	NewProjectTaskResource().Schema(ctx, fwresource.SchemaRequest{}, schemaResponse)

	attributes := schemaResponse.Schema.GetAttributes()
	expected := []string{
		"id", "project_id", "template_id", "inventory_id", "playbook", "environment",
		"limit", "git_branch", "message", "arguments", "debug", "dry_run", "diff",
		"triggers", "wait_for_completion", "timeout", "capture_output", "status", "output",
	}

	for _, name := range expected {
		if _, ok := attributes[name]; !ok {
			t.Errorf("schema is missing attribute %q", name)
		}
	}
	if len(attributes) != len(expected) {
		t.Errorf("schema has %d attributes, expected %d; update ProjectTaskModel and this test together", len(attributes), len(expected))
	}
}

// TestProjectTaskSchemaForceNew pins the immutability contract: every attribute
// describing the job must force replacement, so that changing one queues a new
// run instead of silently mutating a finished one. Only the client-side knobs
// may change in place.
func TestProjectTaskSchemaForceNew(t *testing.T) {
	ctx := context.Background()

	schemaResponse := &fwresource.SchemaResponse{}
	NewProjectTaskResource().Schema(ctx, fwresource.SchemaRequest{}, schemaResponse)

	inPlace := map[string]bool{
		"wait_for_completion": true,
		"timeout":             true,
		"capture_output":      true,
		// Computed-only, never set from configuration.
		"id": true, "status": true, "output": true,
	}

	for name, attribute := range schemaResponse.Schema.GetAttributes() {
		// superschema renders a "(ForceNew)" marker into the description for
		// attributes carrying a RequiresReplace plan modifier.
		forceNew := strings.Contains(attribute.GetMarkdownDescription(), "ForceNew")
		if inPlace[name] && forceNew {
			t.Errorf("attribute %q is ForceNew but should be updatable in place", name)
		}
		if !inPlace[name] && !forceNew {
			t.Errorf("attribute %q describes the job and must be ForceNew", name)
		}
	}
}

func TestConvertProjectTaskModelToTaskBody(t *testing.T) {
	tests := map[string]struct {
		model    ProjectTaskModel
		expected task.PostProjectProjectIDTasksBody
	}{
		"only required attributes": {
			model: ProjectTaskModel{
				TemplateID:  types.Int64Value(7),
				InventoryID: types.Int64Null(),
				Playbook:    types.StringNull(),
				Environment: types.StringNull(),
				Limit:       types.StringNull(),
				GitBranch:   types.StringNull(),
				Message:     types.StringNull(),
				Arguments:   types.StringNull(),
				Debug:       types.BoolValue(false),
				DryRun:      types.BoolValue(false),
				Diff:        types.BoolValue(false),
			},
			// Null optionals must collapse to Go zero values so the generated
			// client's `omitempty` tags drop them from the request body.
			expected: task.PostProjectProjectIDTasksBody{TemplateID: 7},
		},
		"all attributes set": {
			model: ProjectTaskModel{
				TemplateID:  types.Int64Value(7),
				InventoryID: types.Int64Value(3),
				Playbook:    types.StringValue("deploy.yml"),
				Environment: types.StringValue(`{"app_version":"1.2.3"}`),
				Limit:       types.StringValue("web"),
				GitBranch:   types.StringValue("main"),
				Message:     types.StringValue("from terraform"),
				Arguments:   types.StringValue(`["-vvv"]`),
				Debug:       types.BoolValue(true),
				DryRun:      types.BoolValue(true),
				Diff:        types.BoolValue(true),
			},
			expected: task.PostProjectProjectIDTasksBody{
				TemplateID:  7,
				InventoryID: 3,
				Playbook:    "deploy.yml",
				Environment: `{"app_version":"1.2.3"}`,
				Limit:       "web",
				GitBranch:   "main",
				Message:     "from terraform",
				Arguments:   `["-vvv"]`,
				Debug:       true,
				DryRun:      true,
				Diff:        true,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := convertProjectTaskModelToTaskBody(test.model)
			if got != test.expected {
				t.Errorf("got %+v, expected %+v", got, test.expected)
			}
		})
	}
}

func TestApiErrorCode(t *testing.T) {
	tests := map[string]struct {
		err      error
		expected int
	}{
		"nil error":              {err: nil, expected: 0},
		"plain error":            {err: fmt.Errorf("connection refused"), expected: 0},
		"api error 404":          {err: runtime.NewAPIError("op", nil, 404), expected: 404},
		"api error 500":          {err: runtime.NewAPIError("op", nil, 500), expected: 500},
		"wrapped api error 404":  {err: fmt.Errorf("reading task: %w", runtime.NewAPIError("op", nil, 404)), expected: 404},
		"doubly wrapped api 403": {err: fmt.Errorf("outer: %w", fmt.Errorf("inner: %w", runtime.NewAPIError("op", nil, 403))), expected: 403},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if got := apiErrorCode(test.err); got != test.expected {
				t.Errorf("got %d, expected %d", got, test.expected)
			}
		})
	}
}

// TestTaskTerminalStatuses pins the polling contract. An unrecognised status
// must NOT be treated as terminal: a newer Semaphore release adding a status is
// then reported as a timeout rather than as a spurious success.
func TestTaskTerminalStatuses(t *testing.T) {
	terminal := []string{"success", "error", "stopped", "rejected"}
	inFlight := []string{"waiting", "starting", "running", "stopping", "waiting_confirmation", "confirmed", "", "SUCCESS", "some_future_status"}

	for _, status := range terminal {
		if !taskTerminalStatuses[status] {
			t.Errorf("status %q should be terminal", status)
		}
	}
	for _, status := range inFlight {
		if taskTerminalStatuses[status] {
			t.Errorf("status %q should not be terminal", status)
		}
	}

	if !taskTerminalStatuses[taskSuccessStatus] {
		t.Errorf("the success status %q must be terminal", taskSuccessStatus)
	}
}

func TestInt64OrNull(t *testing.T) {
	if got := int64OrNull(0); !got.IsNull() {
		t.Errorf("int64OrNull(0) = %v, expected null", got)
	}
	if got := int64OrNull(5); got.ValueInt64() != 5 {
		t.Errorf("int64OrNull(5) = %v, expected 5", got)
	}
}

func TestStringOrNull(t *testing.T) {
	if got := stringOrNull(""); !got.IsNull() {
		t.Errorf(`stringOrNull("") = %v, expected null`, got)
	}
	if got := stringOrNull("x"); got.ValueString() != "x" {
		t.Errorf(`stringOrNull("x") = %v, expected "x"`, got)
	}
}
