package validate

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// arnValidator validates that a string is an Amazon Resource Name (ARN).
type arnValidator struct {
	tfsdk.AttributeValidator
}

// Description describes the validation in plain text formatting.
func (validator arnValidator) Description(_ context.Context) string {
	return "string must be an ARN"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator arnValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator arnValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
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

	if !arn.IsARN(s.Value) {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid format",
			"expected value to be an ARN",
		)

		return
	}
}

// ARN returns a new ARN validator.
func ARN() tfsdk.AttributeValidator {
	return arnValidator{}
}
