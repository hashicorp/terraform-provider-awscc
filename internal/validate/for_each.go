package validate

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// arrayForEachValidator validates that a List Attribute's contents all satisfy the included validator.
type arrayForEachValidator struct {
	tfsdk.AttributeValidator

	validator tfsdk.AttributeValidator
}

// Description describes the validation in plain text formatting.
func (validator arrayForEachValidator) Description(ctx context.Context) string {
	return fmt.Sprintf("each element must match: %s", validator.validator.Description(ctx))
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator arrayForEachValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator arrayForEachValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	var elemKeyer keyer
	var elems []attr.Value
	switch v := request.AttributeConfig.(type) {
	case types.List:
		if v.Null || v.Unknown {
			return
		}

		elemKeyer = listKeyer
		elems = v.Elems

	case types.Set:
		if v.Null || v.Unknown {
			return
		}

		elemKeyer = setKeyer
		elems = v.Elems

	default:
		response.Diagnostics.AddAttributeError(
			request.AttributePath,
			"Invalid value type",
			fmt.Sprintf("received incorrect value type (%T)", v),
		)

		return
	}

	for i, e := range elems {
		elemPath, diag := elemKeyer(ctx, request.AttributePath, i, e)
		if diag != nil {
			response.Diagnostics.Append(diag)
		}

		elemRequest := tfsdk.ValidateAttributeRequest{
			AttributePath:   elemPath,
			AttributeConfig: e,
			Config:          request.Config,
		}

		var elemResponse tfsdk.ValidateAttributeResponse
		validator.validator.Validate(ctx, elemRequest, &elemResponse)
		response.Diagnostics.Append(elemResponse.Diagnostics...)
	}
}

type keyer func(context.Context, *tftypes.AttributePath, int, attr.Value) (*tftypes.AttributePath, diag.Diagnostic)

func listKeyer(ctx context.Context, path *tftypes.AttributePath, i int, v attr.Value) (*tftypes.AttributePath, diag.Diagnostic) {
	return path.WithElementKeyInt(i), nil
}

func setKeyer(ctx context.Context, path *tftypes.AttributePath, i int, v attr.Value) (*tftypes.AttributePath, diag.Diagnostic) {
	val, err := v.ToTerraformValue(ctx)
	if err != nil {
		return nil, diag.NewAttributeErrorDiagnostic(
			path,
			"No Terraform value",
			"unable to obtain Terraform value:\n\n"+err.Error(),
		)
	}

	return path.WithElementKeyValue(tftypes.NewValue(v.Type(ctx).TerraformType(ctx), val)), nil
}

// ArrayForEach returns a new array for each validator.
func ArrayForEach(validator tfsdk.AttributeValidator) tfsdk.AttributeValidator {
	if validator == nil {
		return nil
	}

	return arrayForEachValidator{
		validator: validator,
	}
}
