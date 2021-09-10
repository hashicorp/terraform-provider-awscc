package validate

import (
	"context"
	"fmt"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
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

	val := n.Value

	if !val.IsInt() {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid value",
			"Not an integer",
		)

		return
	}

	var i big.Int
	_, _ = val.Int(&i)

	if i := i.Int64(); i < int64(validator.min) || i > int64(validator.max) {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid value",
			fmt.Sprintf("expected value to be in the range [%d, %d], got %d", validator.min, validator.max, i),
		)

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

	val := n.Value

	if !val.IsInt() {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid value",
			"Not an integer",
		)

		return
	}

	var i big.Int
	_, _ = val.Int(&i)

	if i := i.Int64(); i < int64(validator.min) {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid value",
			fmt.Sprintf("expected value to be at least %d, got %d", validator.min, i),
		)

		return
	}
}

// IntAtLeast returns a new integer value at least validator.
func IntAtLeast(min int) tfsdk.AttributeValidator {
	return intAtLeastValidator{
		min: min,
	}
}

// intAtLeastValidator validates that an integer Attribute's value matches the value of an element in the valid slice.
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

	val := n.Value

	if !val.IsInt() {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid value",
			"Not an integer",
		)

		return
	}

	var i big.Int
	_, _ = val.Int(&i)

	for _, val := range validator.valid {
		if i.Int64() == int64(val) {
			return
		}
	}

	response.Diagnostics.AddAttributeError(
		request.AttributePath,
		"Invalid value",
		fmt.Sprintf("expected value to be one of %v, got %d", validator.valid, i.Int64()),
	)
}

// IntInSlice returns a new integer in slicde validator.
func IntInSlice(valid []int) tfsdk.AttributeValidator {
	return intInSliceValidator{
		valid: valid,
	}
}
