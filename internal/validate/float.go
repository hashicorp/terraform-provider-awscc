package validate

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
	n, ok := request.AttributeConfig.(types.Number)

	if !ok {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid value type",
			fmt.Sprintf("received incorrect value type (%T)", request.AttributeConfig),
		)

		return
	}

	if n.Unknown || n.Null {
		return
	}

	f, _ := n.Value.Float64()

	if f < validator.min || f > validator.max {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid value",
			fmt.Sprintf("expected value to be in the range [%f, %f], got %f", validator.min, validator.max, f),
		)

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
	n, ok := request.AttributeConfig.(types.Number)

	if !ok {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid value type",
			fmt.Sprintf("received incorrect value type (%T)", request.AttributeConfig),
		)

		return
	}

	if n.Unknown || n.Null {
		return
	}

	f, _ := n.Value.Float64()

	if f < validator.min {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid value",
			fmt.Sprintf("expected value to be at least %f, got %f", validator.min, f),
		)

		return
	}
}

// FloatAtLeast returns a new float value at least validator.
func FloatAtLeast(min float64) tfsdk.AttributeValidator {
	return floatAtLeastValidator{
		min: min,
	}
}
