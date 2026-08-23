package stringvalidator

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestDurationValidator(t *testing.T) {
	tests := map[string]struct {
		value       types.String
		expectError bool
	}{
		"minutes":             {value: types.StringValue("30m"), expectError: false},
		"hours and minutes":   {value: types.StringValue("1h30m"), expectError: false},
		"seconds":             {value: types.StringValue("90s"), expectError: false},
		"milliseconds":        {value: types.StringValue("500ms"), expectError: false},
		"null":                {value: types.StringNull(), expectError: false},
		"unknown":             {value: types.StringUnknown(), expectError: false},
		"empty":               {value: types.StringValue(""), expectError: true},
		"missing unit":        {value: types.StringValue("30"), expectError: true},
		"nonsense":            {value: types.StringValue("soon"), expectError: true},
		"cron format":         {value: types.StringValue("0 0 * * *"), expectError: true},
		"zero":                {value: types.StringValue("0s"), expectError: true},
		"negative":            {value: types.StringValue("-5m"), expectError: true},
		"trailing whitespace": {value: types.StringValue("30m "), expectError: true},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			request := validator.StringRequest{
				Path:        path.Root("timeout"),
				ConfigValue: test.value,
			}
			response := &validator.StringResponse{}

			Duration().ValidateString(context.Background(), request, response)

			if test.expectError && !response.Diagnostics.HasError() {
				t.Errorf("expected %q to be rejected, but it was accepted", test.value)
			}
			if !test.expectError && response.Diagnostics.HasError() {
				t.Errorf("expected %q to be accepted, but got: %+v", test.value, response.Diagnostics)
			}
		})
	}
}

func TestDurationValidatorDescription(t *testing.T) {
	if Duration().MarkdownDescription(context.Background()) == "" {
		t.Error("MarkdownDescription must not be empty; it is rendered into the generated docs")
	}
}
