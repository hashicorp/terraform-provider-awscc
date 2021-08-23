package validate

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	providertypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

// arrayLengthValidator validates that an array (List/Set) Attribute's length is valid.
type arrayLengthValidator struct {
	tfsdk.AttributeValidator

	minItems, maxItems int
}

// Description describes the validation in plain text formatting.
func (v arrayLengthValidator) Description(_ context.Context) string {
	return fmt.Sprintf("array length must be between %d and %d", v.minItems, v.maxItems)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (v arrayLengthValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// Validate performs the validation.
func (v arrayLengthValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	var l int
	list, ok := request.AttributeConfig.(types.List)

	if ok {
		if list.Unknown {
			return
		}

		if !list.Null {
			l = len(list.Elems)
		}
	} else {
		set, ok := request.AttributeConfig.(providertypes.Set)

		if !ok {
			response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Invalid value type",
				Detail:   fmt.Sprintf("received incorrect value type (%T) at path: %s", request.AttributeConfig, request.AttributePath),
			})
		}

		if set.Unknown {
			return
		}

		if !set.Null {
			l = len(set.Elems)
		}
	}

	if l < v.minItems || l > v.maxItems {
		response.Diagnostics = append(response.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Invalid length",
			Detail:   fmt.Sprintf("invalid array length: %d", l),
		})

		return
	}
}

// ArrayLength returns a new array length validator.
func ArrayLength(minItems, maxItems int) tfsdk.AttributeValidator {
	if minItems < 0 || maxItems < 0 || minItems > maxItems {
		return nil
	}

	return arrayLengthValidator{
		minItems: minItems,
		maxItems: maxItems,
	}
}
