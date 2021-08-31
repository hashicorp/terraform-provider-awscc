package validate

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// floatBetweenValidator validates that an float Attribute's value is in a range.
type floatBetweenValidator struct {
	tfsdk.AttributeValidator

	min, max float64
}

// Description describes the validation in plain text formatting.
func (v floatBetweenValidator) Description(_ context.Context) string {
	return fmt.Sprintf("value must be between %f and %f", v.min, v.max)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (v floatBetweenValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// Validate performs the validation.
func (v floatBetweenValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
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

	f, _ := n.Value.Float64()

	if f < v.min || f > v.max {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Invalid length",
			Detail:   fmt.Sprintf("expected %s to be in the range [%f, %f], got %f", request.AttributePath, v.min, v.max, f),
		})

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
func (v floatAtLeastValidator) Description(_ context.Context) string {
	return fmt.Sprintf("value must be at least %f", v.min)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (v floatAtLeastValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// Validate performs the validation.
func (v floatAtLeastValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
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

	f, _ := n.Value.Float64()

	if f < v.min {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Invalid length",
			Detail:   fmt.Sprintf("expected %s to be at least %f, got %f", request.AttributePath, v.min, f),
		})

		return
	}
}

// FloatAtLeast returns a new float value at least validator.
func FloatAtLeast(min float64) tfsdk.AttributeValidator {
	return floatAtLeastValidator{
		min: min,
	}
}
