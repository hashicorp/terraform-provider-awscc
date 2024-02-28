// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validate

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	ccdiag "github.com/hashicorp/terraform-provider-awscc/internal/errs/diag"
)

// isRFC3339TimeValidator validates that a string Attribute's length is a valid RFC33349Time.
type isRFC3339TimeValidator struct{}

// Description describes the validation in plain text formatting.
func (validator isRFC3339TimeValidator) Description(_ context.Context) string {
	return "string must be a valid RFC3339 date-time"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator isRFC3339TimeValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator isRFC3339TimeValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	value := request.ConfigValue.ValueString()

	if _, err := time.Parse(time.RFC3339, value); err != nil {
		response.Diagnostics.Append(ccdiag.NewInvalidFormatAttributeError(
			request.Path,
			fmt.Sprintf("expected value to be a valid RFC3339 date, got %s: %+v", value, err),
		))

		return
	}
}

// IsRFC3339Time returns a new string RFC33349Time validator.
func IsRFC3339Time() validator.String {
	return isRFC3339TimeValidator{}
}
