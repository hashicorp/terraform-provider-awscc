package validate

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/diag"
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
	s, ok := validateString(ctx, request, response)
	if !ok {
		return
	}

	if _, err := time.Parse(time.RFC3339, s); err != nil {
		response.Diagnostics.Append(diag.NewInvalidFormatAttributeError(
			request.Path,
			fmt.Sprintf("expected value to be a valid RFC3339 date, got %s: %+v", s, err),
		))

		return
	}
}

// IsRFC3339Time returns a new string RFC33349Time validator.
func IsRFC3339Time() validator.String {
	return isRFC3339TimeValidator{}
}
