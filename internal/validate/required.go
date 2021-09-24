package validate

import (
	"context"
	"fmt"

	tfdiag "github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-provider-awscc/internal/diags"
	ccdiags "github.com/hashicorp/terraform-provider-awscc/internal/diags"
)

type RequiredAttributesFunc func(names []string) tfdiag.Diagnostics

// Required returns a RequiredAttributesFunc that validates that all required attributes are specfied.
func Required(required ...string) RequiredAttributesFunc {
	return func(names []string) tfdiag.Diagnostics {
		var diags tfdiag.Diagnostics

		for _, r := range required {
			var specified bool

			for _, n := range names {
				if r == n {
					specified = true
					break
				}
			}

			if !specified {
				diags.Append(tfdiag.NewErrorDiagnostic(
					"Attribute not specified",
					fmt.Sprintf("Required attribute (%s) not specified", r),
				))
			}
		}

		return diags
	}
}

// AllOfRequired returns a RequiredAttributesFunc that validates that all of the specified validators pass.
// "To validate against allOf, the given data must be valid against all of the given subschemas."
func AllOfRequired(fs ...RequiredAttributesFunc) RequiredAttributesFunc {
	return func(names []string) tfdiag.Diagnostics {
		var output tfdiag.Diagnostics

		for _, f := range fs {
			output.Append(f(names)...)
		}

		return output
	}
}

// AnyOfRequired returns a RequiredAttributesFunc that validates that any of the specified validators pass.
// "To validate against anyOf, the given data must be valid against any (one or more) of the given subschemas."
func AnyOfRequired(fs ...RequiredAttributesFunc) RequiredAttributesFunc {
	return func(names []string) tfdiag.Diagnostics {
		var output tfdiag.Diagnostics

		for _, f := range fs {
			diags := f(names)

			if diags.HasError() {
				output = append(output, diags...)
			} else {
				return nil
			}
		}

		return output
	}
}

// OneOfRequired returns a RequiredAttributesFunc that validates that exactly one of of the specified validators pass.
// "To validate against oneOf, the given data must be valid against exactly one of the given subschemas."
func OneOfRequired(fs ...RequiredAttributesFunc) RequiredAttributesFunc {
	return func(names []string) tfdiag.Diagnostics {
		var output tfdiag.Diagnostics

		var n int
		for _, f := range fs {
			diags := f(names)

			if diags.HasError() {
				output = append(output, diags...)
			} else {
				n++
			}
		}

		switch n {
		case 0:
		case 1:
			return nil
		default:
			output.Append(tfdiag.NewErrorDiagnostic(
				"Conflicting attributes",
				fmt.Sprintf("%d groups of required attributes match", n),
			))
		}

		return output
	}
}

// requiredAttributesValidator validates that required Attributes are specified.
type requiredAttributesValidator struct {
	tfsdk.AttributeValidator

	fs []RequiredAttributesFunc
}

// Description describes the validation in plain text formatting.
func (validator requiredAttributesValidator) Description(_ context.Context) string {
	return "required Attributes are specified"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator requiredAttributesValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator requiredAttributesValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	// The attribute is either:
	// * Object (SingleNestedAttribute)
	// * List (ListNestedAttribute)
	// * Set (SetNestedAttribute)
	switch v := request.AttributeConfig.(type) {
	case types.Object:
		if v.Null || v.Unknown {
			return
		}

	case types.List:
		if v.Null || v.Unknown {
			return
		}

	case types.Set:
		if v.Null || v.Unknown {
			return
		}

	default:
		response.Diagnostics.Append(ccdiags.NewIncorrectValueTypeAttributeError(
			request.AttributePath,
			v,
		))

		return
	}

	val, err := request.AttributeConfig.ToTerraformValue(ctx)

	if err != nil {
		response.Diagnostics.Append(diags.NewUnableToObtainValueAttributeError(
			request.AttributePath,
			err,
		))

		return
	}

	var diags tfdiag.Diagnostics
	switch v := val.(type) {
	case map[string]tftypes.Value:
		// Ensure that the object is fully known.
		for _, val := range v {
			if !val.IsFullyKnown() {
				return
			}
		}

		diags = evaluateRequiredAttributesFuncs(specifiedAttributes(v), validator.fs...)

	case []tftypes.Value:
		// Ensure that the array is fully known.
		for _, val := range v {
			if !val.IsFullyKnown() {
				return
			}
		}

		for i, val := range v {
			// Each array element must be an Object.
			var vals map[string]tftypes.Value
			if err := val.As(&vals); err != nil {
				response.Diagnostics.Append(ccdiags.NewUnableToConvertValueTypeAttributeError(
					request.AttributePath.WithElementKeyInt(i),
					err,
				))

				return
			}

			diags = evaluateRequiredAttributesFuncs(specifiedAttributes(vals), validator.fs...)

			// Required attributes must be specified for every element.
			if diags.HasError() {
				break
			}
		}
	}

	response.Diagnostics = append(response.Diagnostics, diags...)

	if diags.HasError() {
		return
	}
}

// AttributeRequired returns a new required Attributes validator.
func RequiredAttributes(fs ...RequiredAttributesFunc) tfsdk.AttributeValidator {
	return requiredAttributesValidator{
		fs: fs,
	}
}

// requiredAttributesResourceConfigValidator validates that resource schema-level required Attributes are specified.
type resourceConfigRequiredAttributesValidator struct {
	tfsdk.ResourceConfigValidator

	fs []RequiredAttributesFunc
}

// Description describes the validation in plain text formatting.
func (validator resourceConfigRequiredAttributesValidator) Description(_ context.Context) string {
	return "required Attributes are specified"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator resourceConfigRequiredAttributesValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator resourceConfigRequiredAttributesValidator) Validate(ctx context.Context, request tfsdk.ValidateResourceConfigRequest, response *tfsdk.ValidateResourceConfigResponse) {
	val := request.Config.Raw

	if val.IsNull() || !val.IsFullyKnown() {
		return
	}

	if typ := val.Type(); !typ.Is(tftypes.Object{}) {
		response.Diagnostics.Append(ccdiags.NewIncorrectValueTypeResourceConfigError(typ))

		return
	}

	var vals map[string]tftypes.Value

	if err := val.As(&vals); err != nil {
		response.Diagnostics.Append(ccdiags.NewUnableToConvertValueTypeResourceConfigError(err))

		return
	}

	diags := evaluateRequiredAttributesFuncs(specifiedAttributes(vals), validator.fs...)

	response.Diagnostics = append(response.Diagnostics, diags...)

	if diags.HasError() {
		return
	}
}

// ResourceConfigRequiredAttributes returns a new resource schema-level required Attributes validator.
func ResourceConfigRequiredAttributes(fs ...RequiredAttributesFunc) tfsdk.ResourceConfigValidator {
	return resourceConfigRequiredAttributesValidator{
		fs: fs,
	}
}

func evaluateRequiredAttributesFuncs(names []string, fs ...RequiredAttributesFunc) tfdiag.Diagnostics {
	var diags tfdiag.Diagnostics

	for _, f := range fs {
		diags = append(diags, f(names)...)
	}

	return diags

}

// specifiedAttributes returns the names of the attributes that are set in an object.
// The object is fully known.
func specifiedAttributes(vals map[string]tftypes.Value) []string {
	as := make([]string, 0)

	for a, val := range vals {
		if !val.IsNull() {
			as = append(as, a)
		}
	}

	return as
}
