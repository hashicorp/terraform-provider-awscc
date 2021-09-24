package validate

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-provider-awscc/internal/diags"
)

// arrayLenBetweenValidator validates that an array (List/Set) Attribute's length is in a range.
type arrayLenBetweenValidator struct {
	tfsdk.AttributeValidator

	minItems, maxItems int
}

// Description describes the validation in plain text formatting.
func (validator arrayLenBetweenValidator) Description(_ context.Context) string {
	return fmt.Sprintf("array length must be between %d and %d", validator.minItems, validator.maxItems)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator arrayLenBetweenValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator arrayLenBetweenValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	elems, _, ok := validateArray(request, response)
	if !ok {
		return
	}

	if l := len(elems); l < validator.minItems || l > validator.maxItems {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid length",
			fmt.Sprintf("expected length to be in the range [%d, %d], got %d", validator.minItems, validator.maxItems, l),
		)

		return
	}
}

// ArrayLenBetween returns a new array length between validator.
func ArrayLenBetween(minItems, maxItems int) tfsdk.AttributeValidator {
	if minItems < 0 || maxItems < 0 || minItems > maxItems {
		return nil
	}

	return arrayLenBetweenValidator{
		minItems: minItems,
		maxItems: maxItems,
	}
}

// arrayLenAtLeastValidator validates that an array (List/Set) Attribute's length is at least a certain value.
type arrayLenAtLeastValidator struct {
	tfsdk.AttributeValidator

	minItems int
}

// Description describes the validation in plain text formatting.
func (validator arrayLenAtLeastValidator) Description(_ context.Context) string {
	return fmt.Sprintf("array length must be at least %d", validator.minItems)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator arrayLenAtLeastValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator arrayLenAtLeastValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	elems, _, ok := validateArray(request, response)
	if !ok {
		return
	}

	if l := len(elems); l < validator.minItems {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid length",
			fmt.Sprintf("expected length to be at least %d, got %d", validator.minItems, l),
		)

		return
	}
}

// ArrayLenAtLeast returns a new array length at least validator.
func ArrayLenAtLeast(minItems int) tfsdk.AttributeValidator {
	if minItems < 0 {
		return nil
	}

	return arrayLenAtLeastValidator{
		minItems: minItems,
	}
}

// arrayLenAtMostValidator validates that an array (List/Set) Attribute's length is at most a certain value.
type arrayLenAtMostValidator struct {
	tfsdk.AttributeValidator

	maxItems int
}

// Description describes the validation in plain text formatting.
func (validator arrayLenAtMostValidator) Description(_ context.Context) string {
	return fmt.Sprintf("array length must be at most %d", validator.maxItems)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator arrayLenAtMostValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator arrayLenAtMostValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	elems, _, ok := validateArray(request, response)
	if !ok {
		return
	}

	if l := len(elems); l > validator.maxItems {
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid length",
			fmt.Sprintf("expected length to be at most %d, got %d", validator.maxItems, l),
		)

		return
	}
}

// ArrayLenAtMost returns a new array length at most validator.
func ArrayLenAtMost(maxItems int) tfsdk.AttributeValidator {
	if maxItems < 0 {
		return nil
	}

	return arrayLenAtMostValidator{
		maxItems: maxItems,
	}
}

type arrayKeyer func(context.Context, *tftypes.AttributePath, int, attr.Value) (*tftypes.AttributePath, diag.Diagnostic)

func listKeyer(ctx context.Context, path *tftypes.AttributePath, i int, v attr.Value) (*tftypes.AttributePath, diag.Diagnostic) {
	return path.WithElementKeyInt(i), nil
}

func setKeyer(ctx context.Context, path *tftypes.AttributePath, i int, v attr.Value) (*tftypes.AttributePath, diag.Diagnostic) {
	val, err := v.ToTerraformValue(ctx)
	if err != nil {
		return nil, diags.NewUnableToObtainValueAttributeError(path, err)
	}

	return path.WithElementKeyValue(tftypes.NewValue(v.Type(ctx).TerraformType(ctx), val)), nil
}

func validateArray(request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) ([]attr.Value, arrayKeyer, bool) {
	var elemKeyer arrayKeyer
	var elems []attr.Value
	switch v := request.AttributeConfig.(type) {
	case types.List:
		if v.Null || v.Unknown {
			return elems, elemKeyer, false
		}

		elemKeyer = listKeyer
		elems = v.Elems

	case types.Set:
		if v.Null || v.Unknown {
			return elems, elemKeyer, false
		}

		elemKeyer = setKeyer
		elems = v.Elems

	default:
		response.Diagnostics.Append(diags.NewIncorrectValueTypeAttributeError(
			request.AttributePath,
			v,
		))

		return elems, elemKeyer, false
	}

	return elems, elemKeyer, true
}
