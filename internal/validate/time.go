package validate

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// isRFC3339TimeValidator validates that a string Attribute's length is a valid RFC33349Time.
type isRFC3339TimeValidator struct {
	tfsdk.AttributeValidator
}

// Description describes the validation in plain text formatting.
func (validator isRFC3339TimeValidator) Description(_ context.Context) string {
	return "string must be a valid RFC3339 date-time"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator isRFC3339TimeValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator isRFC3339TimeValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	s, ok := request.AttributeConfig.(types.String)

	if !ok {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid value type",
			fmt.Sprintf("received incorrect value type (%T)", request.AttributeConfig),
		)

		return
	}

	if s.Unknown || s.Null {
		return
	}

	if _, err := time.Parse(time.RFC3339, s.Value); err != nil {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid format",
			fmt.Sprintf("expected value to be a valid RFC3339 date, got %s: %+v", s.Value, err),
		)

		return
	}
}

// IsRFC3339Time returns a new string RFC33349Time validator.
func IsRFC3339Time() tfsdk.AttributeValidator {
	return isRFC3339TimeValidator{}
}
