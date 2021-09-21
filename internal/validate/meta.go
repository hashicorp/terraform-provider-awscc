package validate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-provider-awscc/internal/strings"
)

// allValidator validates that a string Attribute's length is in a range.
type allValidator struct {
	tfsdk.AttributeValidator

	validators []tfsdk.AttributeValidator
}

// Description describes the validation in plain text formatting.
func (validator allValidator) Description(ctx context.Context) string {
	descriptions := make([]string, len(validator.validators))
	for i, v := range validator.validators {
		descriptions[i] = v.Description(ctx)
	}
	return strings.ProseJoin(descriptions)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator allValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator allValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	for _, v := range validator.validators {
		var elemResponse tfsdk.ValidateAttributeResponse
		v.Validate(ctx, request, &elemResponse)
		response.Diagnostics.Append(elemResponse.Diagnostics...)
	}
}

// All returns a new string length between validator.
func All(validators ...tfsdk.AttributeValidator) tfsdk.AttributeValidator {
	if len(validators) == 0 {
		return nil
	}

	return allValidator{
		validators: validators,
	}
}
