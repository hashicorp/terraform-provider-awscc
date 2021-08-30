package validate

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	providertypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
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
	var l int
	switch v := request.AttributeConfig.(type) {
	case types.List:
		if v.Null || v.Unknown {
			return
		}

		l = len(v.Elems)

	case providertypes.Set:
		if v.Null || v.Unknown {
			return
		}

		l = len(v.Elems)

	default:
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Invalid value type",
			Detail:   fmt.Sprintf("received incorrect value type (%T) at path: %s", v, request.AttributePath),
		})

		return
	}

	if l < validator.minItems || l > validator.maxItems {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Invalid length",
			Detail:   fmt.Sprintf("expected length of %s to be in the range [%d, %d], got %d", request.AttributePath, validator.minItems, validator.maxItems, l),
		})

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
	var l int
	switch v := request.AttributeConfig.(type) {
	case types.List:
		if v.Null || v.Unknown {
			return
		}

		l = len(v.Elems)

	case providertypes.Set:
		if v.Null || v.Unknown {
			return
		}

		l = len(v.Elems)

	default:
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Invalid value type",
			Detail:   fmt.Sprintf("received incorrect value type (%T) at path: %s", v, request.AttributePath),
		})

		return
	}

	if l < validator.minItems {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Invalid length",
			Detail:   fmt.Sprintf("expected length of %s to be at least %d, got %d", request.AttributePath, validator.minItems, l),
		})

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
