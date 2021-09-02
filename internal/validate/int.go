package validate

import (
	"context"
	"fmt"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// intBetweenValidator validates that an integer Attribute's value is in a range.
type intBetweenValidator struct {
	tfsdk.AttributeValidator

	min, max int
}

// Description describes the validation in plain text formatting.
func (v intBetweenValidator) Description(_ context.Context) string {
	return fmt.Sprintf("value must be between %d and %d", v.min, v.max)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (v intBetweenValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// Validate performs the validation.
func (v intBetweenValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	n, ok := request.AttributeConfig.(types.Number)

	if !ok {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Invalid value type",
			Detail:   fmt.Sprintf("received incorrect value type (%T) at path: %s", request.AttributeConfig, request.AttributePath),
		})

		return
	}

	if n.Unknown || n.Null {
		return
	}

	val := n.Value

	if !val.IsInt() {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Invalid value",
			Detail:   fmt.Sprintf("invalid value (%s) at path: %s", request.AttributeConfig, request.AttributePath),
		})

		return
	}

	var i big.Int
	_, _ = val.Int(&i)

	if i := i.Int64(); i < int64(v.min) || i > int64(v.max) {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Invalid length",
			Detail:   fmt.Sprintf("expected %s to be in the range [%d, %d], got %d", request.AttributePath, v.min, v.max, i),
		})

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
func (v intAtLeastValidator) Description(_ context.Context) string {
	return fmt.Sprintf("value must be at least %d", v.min)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (v intAtLeastValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// Validate performs the validation.
func (v intAtLeastValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	n, ok := request.AttributeConfig.(types.Number)

	if !ok {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Invalid value type",
			Detail:   fmt.Sprintf("received incorrect value type (%T) at path: %s", request.AttributeConfig, request.AttributePath),
		})

		return
	}

	if n.Unknown || n.Null {
		return
	}

	val := n.Value

	if !val.IsInt() {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Invalid value",
			Detail:   fmt.Sprintf("invalid value (%s) at path: %s", request.AttributeConfig, request.AttributePath),
		})

		return
	}

	var i big.Int
	_, _ = val.Int(&i)

	if i := i.Int64(); i < int64(v.min) {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Invalid length",
			Detail:   fmt.Sprintf("expected %s to be at least %d, got %d", request.AttributePath, v.min, i),
		})

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
func (v intInSliceValidator) Description(_ context.Context) string {
	return fmt.Sprintf("value must be one of %v", v.valid)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (v intInSliceValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// Validate performs the validation.
func (v intInSliceValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	n, ok := request.AttributeConfig.(types.Number)

	if !ok {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Invalid value type",
			Detail:   fmt.Sprintf("received incorrect value type (%T) at path: %s", request.AttributeConfig, request.AttributePath),
		})

		return
	}

	if n.Unknown || n.Null {
		return
	}

	val := n.Value

	if !val.IsInt() {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Invalid value",
			Detail:   fmt.Sprintf("invalid value (%s) at path: %s", request.AttributeConfig, request.AttributePath),
		})

		return
	}

	var i big.Int
	_, _ = val.Int(&i)

	for _, val := range v.valid {
		if i.Int64() == int64(val) {
			return
		}
	}

	response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
		Severity: tfprotov6.DiagnosticSeverityError,
		Summary:  "Invalid value",
		Detail:   fmt.Sprintf("expected %s to be one of %v, got %d", request.AttributePath, v.valid, i.Int64()),
	})
}

// IntInSlice returns a new integer in slicde validator.
func IntInSlice(valid []int) tfsdk.AttributeValidator {
	return intInSliceValidator{
		valid: valid,
	}
}
