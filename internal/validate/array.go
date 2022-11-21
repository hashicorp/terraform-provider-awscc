package validate

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	ccdiag "github.com/hashicorp/terraform-provider-awscc/internal/diag"
)

// arrayLenBetweenValidator validates that an array (List/Set) Attribute's length is in a range.
type arrayLenBetweenValidator struct {
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
	elems, _, ok := validateArray(ctx, request, response)
	if !ok {
		return
	}

	if l := len(elems); l < validator.minItems || l > validator.maxItems {
		response.Diagnostics.Append(ccdiag.NewInvalidLengthBetweenAttributeError(
			request.AttributePath, validator.minItems, validator.maxItems, l,
		))

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
	elems, _, ok := validateArray(ctx, request, response)
	if !ok {
		return
	}

	if l := len(elems); l < validator.minItems {
		response.Diagnostics.Append(ccdiag.NewInvalidLengthAtLeastAttributeError(
			request.AttributePath, validator.minItems, l,
		))

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
	elems, _, ok := validateArray(ctx, request, response)
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

type arrayKeyer func(context.Context, path.Path, int, attr.Value) (path.Path, diag.Diagnostic)

func listKeyer(ctx context.Context, p path.Path, i int, v attr.Value) (path.Path, diag.Diagnostic) {
	return p.AtListIndex(i), nil
}

func setKeyer(ctx context.Context, p path.Path, i int, v attr.Value) (path.Path, diag.Diagnostic) {
	return p.AtSetValue(v), nil
}

func validateArray(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) ([]attr.Value, arrayKeyer, bool) {
	var elemKeyer arrayKeyer
	var elems []attr.Value

	var v types.List

	diags := tfsdk.ValueAs(ctx, request.AttributeConfig, &v)

	if diags.HasError() {
		var v types.Set

		diags := tfsdk.ValueAs(ctx, request.AttributeConfig, &v)

		if diags.HasError() {
			response.Diagnostics = append(response.Diagnostics, diags...)

			return elems, elemKeyer, false
		} else {
			if v.IsNull() || v.IsUnknown() {
				return elems, elemKeyer, false
			}

			elemKeyer = setKeyer
			elems = v.Elements()
		}
	} else {
		if v.IsNull() || v.IsUnknown() {
			return elems, elemKeyer, false
		}

		elemKeyer = listKeyer
		elems = v.Elements()
	}

	return elems, elemKeyer, true
}
