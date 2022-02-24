package validate

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	ccdiag "github.com/hashicorp/terraform-provider-awscc/internal/diag"
)

// intBetweenValidator validates that an integer Attribute's value is in a range.
type intBetweenValidator struct {
	tfsdk.AttributeValidator

	min, max int
}

// Description describes the validation in plain text formatting.
func (validator intBetweenValidator) Description(_ context.Context) string {
	return fmt.Sprintf("value must be between %d and %d", validator.min, validator.max)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator intBetweenValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator intBetweenValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	i, ok := validateInt(ctx, request, response)
	if !ok {
		return
	}

	if i < int64(validator.min) || i > int64(validator.max) {
		response.Diagnostics.Append(ccdiag.NewInvalidValueAttributeError(
			request.AttributePath,
			fmt.Sprintf("expected value to be in the range [%d, %d], got %d", validator.min, validator.max, i),
		))

		return
	}
}

// IntBetween returns a new integer value between validator.
func IntBetween(min, max int) tfsdk.AttributeValidator {
	if min > max {
		return nil
	}

	return intBetweenValidator{
		min: min,
		max: max,
	}
}

// intAtLeastValidator validates that an integer Attribute's value is at least a certain value.
type intAtLeastValidator struct {
	tfsdk.AttributeValidator

	min int
}

// Description describes the validation in plain text formatting.
func (validator intAtLeastValidator) Description(_ context.Context) string {
	return fmt.Sprintf("value must be at least %d", validator.min)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator intAtLeastValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator intAtLeastValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	i, ok := validateInt(ctx, request, response)
	if !ok {
		return
	}

	if i < int64(validator.min) {
		response.Diagnostics.Append(ccdiag.NewInvalidValueAttributeError(
			request.AttributePath,
			fmt.Sprintf("expected value to be at least %d, got %d", validator.min, i),
		))

		return
	}
}

// IntAtLeast returns a new integer value at least validator.
func IntAtLeast(min int) tfsdk.AttributeValidator {
	return intAtLeastValidator{
		min: min,
	}
}

// intAtMostValidator validates that an integer Attribute's value is at most a certain value.
type intAtMostValidator struct {
	tfsdk.AttributeValidator

	max int
}

// Description describes the validation in plain text formatting.
func (validator intAtMostValidator) Description(_ context.Context) string {
	return fmt.Sprintf("value must be at most %d", validator.max)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator intAtMostValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator intAtMostValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	i, ok := validateInt(ctx, request, response)
	if !ok {
		return
	}

	if i > int64(validator.max) {
		response.Diagnostics.Append(ccdiag.NewInvalidValueAttributeError(
			request.AttributePath,
			fmt.Sprintf("expected value to be at most %d, got %d", validator.max, i),
		))

		return
	}
}

// IntAtMost returns a new integer value at most validator.
func IntAtMost(max int) tfsdk.AttributeValidator {
	return intAtMostValidator{
		max: max,
	}
}

// intInSliceValidator validates that an integer Attribute's value matches the value of an element in the valid slice.
type intInSliceValidator struct {
	tfsdk.AttributeValidator

	valid []int
}

// Description describes the validation in plain text formatting.
func (validator intInSliceValidator) Description(_ context.Context) string {
	return fmt.Sprintf("value must be one of %v", validator.valid)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator intInSliceValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator intInSliceValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	i, ok := validateInt(ctx, request, response)
	if !ok {
		return
	}

	for _, val := range validator.valid {
		if i == int64(val) {
			return
		}
	}

	response.Diagnostics.Append(newIntNotInSliceError(
		request.AttributePath,
		validator.valid,
		i,
	))

}

func newIntNotInSliceError(path *tftypes.AttributePath, valid []int, value int64) diag.Diagnostic {
	return ccdiag.NewInvalidValueAttributeError(
		path,
		fmt.Sprintf("expected value to be one of %v, got %d", valid, value),
	)
}

// IntInSlice returns a new integer in slicde validator.
func IntInSlice(valid []int) tfsdk.AttributeValidator {
	return intInSliceValidator{
		valid: valid,
	}
}

func newNotAnIntegerValueError(path *tftypes.AttributePath) diag.Diagnostic {
	return ccdiag.NewInvalidValueAttributeError(path, "Not an integer")
}

func validateInt(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) (int64, bool) {
	var n types.Int64

	diags := tfsdk.ValueAs(ctx, request.AttributeConfig, &n)

	if diags.HasError() {
		var n types.Number

		diags := tfsdk.ValueAs(ctx, request.AttributeConfig, &n)

		if diags.HasError() {
			response.Diagnostics = append(response.Diagnostics, diags...)

			return 0, false
		} else {
			if n.Unknown || n.Null {
				return 0, false
			}

			val := n.Value

			if !val.IsInt() {
				response.Diagnostics.Append(newNotAnIntegerValueError(
					request.AttributePath,
				))

				return 0, false
			}

			i, _ := val.Int64()

			return i, true
		}
	} else {
		if n.Unknown || n.Null {
			return 0, false
		}

		return n.Value, true
	}
}
