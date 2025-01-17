package stringvalidator

import (
	"context"
	"fmt"
	"github.com/adhocore/gronx"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = CronFormatValidator{}

type CronFormatValidator struct{}

func (v CronFormatValidator) Description(ctx context.Context) string {
	return ""
}

func (v CronFormatValidator) MarkdownDescription(ctx context.Context) string {
	return "Must be valid [Cron Expression](https://github.com/adhocore/gronx?tab=readme-ov-file#cron-expression)."
}

func (v CronFormatValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	// If the value is unknown or null, there is nothing to validate.
	if req.ConfigValue.IsUnknown() || req.ConfigValue.IsNull() {
		return
	}

	gron := gronx.New()
	if !gron.IsValid(req.ConfigValue.ValueString()) {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid cron format",
			fmt.Sprintf("cron_format must be a valid cron expression, got %s", req.ConfigValue.ValueString()),
		)
		return
	}
}

func CronFormat() CronFormatValidator {
	return CronFormatValidator{}
}
