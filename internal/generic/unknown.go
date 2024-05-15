// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	ccdiag "github.com/hashicorp/terraform-provider-awscc/internal/errs/diag"
)

type unknownsVisitor struct {
	paths []*tftypes.AttributePath
}

func (w *unknownsVisitor) visit(path *tftypes.AttributePath, val tftypes.Value) (bool, error) {
	if !val.IsKnown() {
		w.paths = append(w.paths, path)

		return true, nil
	}

	if typ := val.Type(); typ.Is(tftypes.Set{}) {
		var vals []tftypes.Value
		if err := val.As(&vals); err != nil {
			return false, path.NewError(err)
		}

		for _, val := range vals {
			// If any of the set's elements has unknown values, mark the whole set as unknown.
			// This prevents problems in copyStateValueAtPath where a match between set elements cannot be made.
			if !val.IsFullyKnown() {
				w.paths = append(w.paths, path)

				return false, nil
			}
		}
	}

	return true, nil
}

// UnknownValuePaths returns all the paths to all the unknown values in the specified Terraform Value.
func UnknownValuePaths(ctx context.Context, val tftypes.Value) ([]*tftypes.AttributePath, error) {
	walker := &unknownsVisitor{}

	if err := tftypes.Walk(val, walker.visit); err != nil {
		return nil, err
	}

	return walker.paths, nil
}

// SetUnknownValuesFromResourceModel fills any unknown State values from a Cloud Control ResourceModel (string).
// The unknown value paths are obtained from the State via a previous call to UnknownValuePaths.
// Functionality is split between these 2 functions, rather than calling UnknownValuePaths from within this function,
// so as to avoid unnecessary Cloud Control API calls to obtain the current ResourceModel.
func SetUnknownValuesFromResourceModel(ctx context.Context, state *tfsdk.State, unknowns []*tftypes.AttributePath, resourceModel string, cfToTfNameMap map[string]string) error {
	// Get the Terraform Value of the ResourceModel.
	translator := toTerraform{cfToTfNameMap: cfToTfNameMap}
	schema := state.Schema
	val, err := translator.FromString(ctx, schema, resourceModel)

	if err != nil {
		return err
	}

	src := tfsdk.State{
		Schema: schema,
		Raw:    val,
	}

	// Copy all unknown values from the ResourceModel to destination State.
	for _, path := range unknowns {
		path, diags := attributePath(ctx, path, schema)
		if diags.HasError() {
			return ccdiag.DiagnosticsError(diags)
		}

		diags.Append(copyStateValueAtPath(ctx, state, &src, path)...)
		if diags.HasError() {
			return ccdiag.DiagnosticsError(diags)
		}
	}

	return nil
}

type typeAtTerraformPather interface {
	TypeAtTerraformPath(context.Context, *tftypes.AttributePath) (attr.Type, error)
}

// Lifted from github.com/hashicorp/terraform-plugin-framework/tfsdk/tftypes_attribute_path.go.
// attributePath returns the path.Path equivalent of a *tftypes.AttributePath.
func attributePath(ctx context.Context, tfType *tftypes.AttributePath, schema typeAtTerraformPather) (path.Path, diag.Diagnostics) {
	fwPath := path.Empty()

	for tfTypeStepIndex, tfTypeStep := range tfType.Steps() {
		currentTfTypeSteps := tfType.Steps()[:tfTypeStepIndex+1]
		currentTfTypePath := tftypes.NewAttributePathWithSteps(currentTfTypeSteps)
		attrType, err := schema.TypeAtTerraformPath(ctx, currentTfTypePath)

		if err != nil {
			return path.Empty(), diag.Diagnostics{
				diag.NewErrorDiagnostic(
					"Unable to Convert Attribute Path",
					"An unexpected error occurred while trying to convert an attribute path. "+
						"This is an error in terraform-plugin-framework used by the provider. "+
						"Please report the following to the provider developers.\n\n"+
						// Since this is an error with the attribute path
						// conversion, we cannot return a protocol path-based
						// diagnostic. Returning a framework human-readable
						// representation seems like the next best thing to do.
						fmt.Sprintf("Attribute Path: %s\n", currentTfTypePath.String())+
						fmt.Sprintf("Original Error: %s", err),
				),
			}
		}

		fwStep, err := attributePathStep(ctx, tfTypeStep, attrType)

		if err != nil {
			return path.Empty(), diag.Diagnostics{
				diag.NewErrorDiagnostic(
					"Unable to Convert Attribute Path",
					"An unexpected error occurred while trying to convert an attribute path. "+
						"This is either an error in terraform-plugin-framework or a custom attribute type used by the provider. "+
						"Please report the following to the provider developers.\n\n"+
						// Since this is an error with the attribute path
						// conversion, we cannot return a protocol path-based
						// diagnostic. Returning a framework human-readable
						// representation seems like the next best thing to do.
						fmt.Sprintf("Attribute Path: %s\n", currentTfTypePath.String())+
						fmt.Sprintf("Original Error: %s", err),
				),
			}
		}

		// In lieu of creating a path.NewPathFromSteps function, this path
		// building logic is inlined to not expand the path package API.
		switch fwStep := fwStep.(type) {
		case path.PathStepAttributeName:
			fwPath = fwPath.AtName(string(fwStep))
		case path.PathStepElementKeyInt:
			fwPath = fwPath.AtListIndex(int(fwStep))
		case path.PathStepElementKeyString:
			fwPath = fwPath.AtMapKey(string(fwStep))
		case path.PathStepElementKeyValue:
			fwPath = fwPath.AtSetValue(fwStep.Value)
		default:
			return fwPath, diag.Diagnostics{
				diag.NewErrorDiagnostic(
					"Unable to Convert Attribute Path",
					"An unexpected error occurred while trying to convert an attribute path. "+
						"This is an error in terraform-plugin-framework used by the provider. "+
						"Please report the following to the provider developers.\n\n"+
						// Since this is an error with the attribute path
						// conversion, we cannot return a protocol path-based
						// diagnostic. Returning a framework human-readable
						// representation seems like the next best thing to do.
						fmt.Sprintf("Attribute Path: %s\n", currentTfTypePath.String())+
						fmt.Sprintf("Original Error: unknown path.PathStep type: %#v", fwStep),
				),
			}
		}
	}

	return fwPath, nil
}

// Lifted from github.com/hashicorp/terraform-plugin-framework/internal/fromtftypes/attribute_path_step.go.
func attributePathStep(ctx context.Context, tfType tftypes.AttributePathStep, attrType attr.Type) (path.PathStep, error) {
	switch tfType := tfType.(type) {
	case tftypes.AttributeName:
		return path.PathStepAttributeName(string(tfType)), nil
	case tftypes.ElementKeyInt:
		return path.PathStepElementKeyInt(int64(tfType)), nil
	case tftypes.ElementKeyString:
		return path.PathStepElementKeyString(string(tfType)), nil
	case tftypes.ElementKeyValue:
		attrValue, err := value(ctx, tftypes.Value(tfType), attrType)

		if err != nil {
			return nil, fmt.Errorf("unable to create PathStepElementKeyValue from tftypes.Value: %w", err)
		}

		return path.PathStepElementKeyValue{Value: attrValue}, nil
	default:
		return nil, fmt.Errorf("unknown tftypes.AttributePathStep: %#v", tfType)
	}
}

// Lifted from github.com/hashicorp/terraform-plugin-framework/internal/fromtftypes/value.go.
func value(ctx context.Context, tfType tftypes.Value, attrType attr.Type) (attr.Value, error) {
	if attrType == nil {
		return nil, fmt.Errorf("unable to convert tftypes.Value (%s) to attr.Value: missing attr.Type", tfType.String())
	}

	attrValue, err := attrType.ValueFromTerraform(ctx, tfType)

	if err != nil {
		return nil, fmt.Errorf("unable to convert tftypes.Value (%s) to attr.Value: %w", tfType.String(), err)
	}

	return attrValue, nil
}
