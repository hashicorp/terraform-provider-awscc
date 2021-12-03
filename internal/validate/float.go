package validate

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/diag"
)

// floatBetweenValidator validates that an float Attribute's value is in a range.
type floatBetweenValidator struct {
	tfsdk.AttributeValidator

	min, max float64
}

// Description describes the validation in plain text formatting.
func (validator floatBetweenValidator) Description(_ context.Context) string {
	return fmt.Sprintf("value must be between %f and %f", validator.min, validator.max)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator floatBetweenValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator floatBetweenValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	f, ok := validateFloat(request, response)
	if !ok {
		return
	}

	if f < validator.min || f > validator.max {
		response.Diagnostics.Append(diag.NewInvalidValueAttributeError(
			request.AttributePath,
			fmt.Sprintf("expected value to be in the range [%f, %f], got %f", validator.min, validator.max, f),
		))

		return
	}
}

// FloatBetween returns a new float value between validator.
func FloatBetween(min, max float64) tfsdk.AttributeValidator {
	if min > max {
		return nil
	}

	return floatBetweenValidator{
		min: min,
		max: max,
	}
}

// floatAtLeastValidator validates that an float Attribute's value is at least a certain value.
type floatAtLeastValidator struct {
	tfsdk.AttributeValidator

	min float64
}

// Description describes the validation in plain text formatting.
func (validator floatAtLeastValidator) Description(_ context.Context) string {
	return fmt.Sprintf("value must be at least %f", validator.min)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator floatAtLeastValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator floatAtLeastValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	f, ok := validateFloat(request, response)
	if !ok {
		return
	}

	if f < validator.min {
		response.Diagnostics.Append(diag.NewInvalidValueAttributeError(
			request.AttributePath,
			fmt.Sprintf("expected value to be at least %f, got %f", validator.min, f),
		))

		return
	}
}

// FloatAtLeast returns a new float value at least validator.
func FloatAtLeast(min float64) tfsdk.AttributeValidator {
	return floatAtLeastValidator{
		min: min,
	}
}

// floatAtMostValidator validates that an float Attribute's value is at most a certain value.
type floatAtMostValidator struct {
	tfsdk.AttributeValidator

	max float64
}

// Description describes the validation in plain text formatting.
func (validator floatAtMostValidator) Description(_ context.Context) string {
	return fmt.Sprintf("value must be at most %f", validator.max)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator floatAtMostValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator floatAtMostValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	f, ok := validateFloat(request, response)
	if !ok {
		return
	}

	if f > validator.max {
		response.Diagnostics.Append(diag.NewInvalidValueAttributeError(
			request.AttributePath,
			fmt.Sprintf("expected value to be at most %f, got %f", validator.max, f),
		))

		return
	}
}

// FloatAtMost returns a new float value at nost validator.
func FloatAtMost(max float64) tfsdk.AttributeValidator {
	return floatAtMostValidator{
		max: max,
	}
}

func validateFloat(request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) (float64, bool) {
	n, ok := request.AttributeConfig.(types.Number)

	if !ok {
		response.Diagnostics.Append(diag.NewIncorrectValueTypeAttributeError(
			request.AttributePath,
			request.AttributeConfig,
		))

		return 0, false
	}

	if n.Unknown || n.Null {
		return 0, false
	}

	f, _ := n.Value.Float64()

	return f, true
}
