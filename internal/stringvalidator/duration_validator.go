package stringvalidator

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = DurationValidator{}

type DurationValidator struct{}

func (v DurationValidator) Description(ctx context.Context) string {
	return ""
}

func (v DurationValidator) MarkdownDescription(ctx context.Context) string {
	return "Must be a valid [Go duration string](https://pkg.go.dev/time#ParseDuration), e.g. `30m`, `1h30m`, `90s`."
}

func (v DurationValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	// If the value is unknown or null, there is nothing to validate.
	if req.ConfigValue.IsUnknown() || req.ConfigValue.IsNull() {
		return
	}

	duration, err := time.ParseDuration(req.ConfigValue.ValueString())
	if err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid duration",
			fmt.Sprintf("must be a valid Go duration string (e.g. \"30m\"), got %q: %s", req.ConfigValue.ValueString(), err.Error()),
		)
		return
	}
	if duration <= 0 {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid duration",
			fmt.Sprintf("must be a positive duration, got %q", req.ConfigValue.ValueString()),
		)
		return
	}
}

func Duration() DurationValidator {
	return DurationValidator{}
}
