// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/attr/xattr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type jsonStringType uint8

const (
	JSONStringType jsonStringType = iota
)

var (
	_ xattr.TypeWithValidate = JSONStringType
)

func (t jsonStringType) TerraformType(_ context.Context) tftypes.Type {
	return tftypes.String
}

func (t jsonStringType) ValueFromTerraform(_ context.Context, v tftypes.Value) (attr.Value, error) {
	if !v.IsKnown() {
		return JSONStringUnknown(), nil
	}

	if v.IsNull() {
		return JSONStringNull(), nil
	}

	var s string
	err := v.As(&s)

	if err != nil {
		return nil, err
	}

	// Don't return the normalized string here, else Plan != Config.
	_, err = normalizeJSONString(s)

	if err != nil {
		return nil, err
	}

	return JSONStringValue(s), nil
}

func (t jsonStringType) ValueType(context.Context) attr.Value {
	return JSONString{}
}

func (t jsonStringType) Equal(o attr.Type) bool {
	_, ok := o.(jsonStringType)

	return ok
}

func (t jsonStringType) ApplyTerraform5AttributePathStep(step tftypes.AttributePathStep) (interface{}, error) {
	return nil, fmt.Errorf("cannot apply AttributePathStep %v to %s", step, t.String())
}

func (t jsonStringType) String() string {
	return "JSONStringType"
}

func (t jsonStringType) Validate(ctx context.Context, v tftypes.Value, p path.Path) diag.Diagnostics {
	var diags diag.Diagnostics

	if !v.Type().Is(tftypes.String) {
		diags.AddAttributeError(
			p,
			"JSONString Type Validation Error",
			"An unexpected error was encountered trying to validate an attribute value. This is always an error in the provider. Please report the following to the provider developer:\n\n"+
				fmt.Sprintf("Expected String value, received %T with value: %v", v, v),
		)

		return diags
	}

	if !v.IsKnown() || v.IsNull() {
		return diags
	}

	var s string
	err := v.As(&s)

	if err != nil {
		diags.AddAttributeError(
			p,
			"JSONString Type Validation Error",
			"An unexpected error was encountered trying to validate an attribute value. This is always an error in the provider. Please report the following to the provider developer:\n\n"+
				fmt.Sprintf("Cannot convert value to String: %s", err),
		)

		return diags
	}

	if s == "" {
		return diags
	}

	_, err = expandJSONFromString(s)

	if err != nil {
		diags.AddAttributeError(
			p,
			"JSONString Type Validation Error",
			fmt.Sprintf("Value %q cannot be parsed as a JSON string.", v),
		)

		return diags
	}

	return diags
}

func (t jsonStringType) ValueFromString(_ context.Context, s basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	if s.IsUnknown() {
		return JSONStringUnknown(), diags
	}

	if s.IsNull() {
		return JSONStringNull(), diags
	}

	return JSONStringValue(s.ValueString()), diags
}

func (t jsonStringType) AttributePlanModifier() planmodifier.String {
	return jsonStringAttributePlanModifier{}
}

func JSONStringNull() JSONString {
	return JSONString{
		state: attr.ValueStateNull,
	}
}

func JSONStringUnknown() JSONString {
	return JSONString{
		state: attr.ValueStateUnknown,
	}
}

func JSONStringValue(value string) JSONString {
	return JSONString{
		state: attr.ValueStateKnown,
		value: value,
	}
}

// A JSONString is a string containing a valid JSON document.
type JSONString struct {
	// state represents whether the value is null, unknown, or known. The
	// zero-value is null.
	state attr.ValueState

	// value contains the known value, if not null or unknown.
	value string
}

func (s JSONString) Type(_ context.Context) attr.Type {
	return JSONStringType
}

func (s JSONString) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	t := JSONStringType.TerraformType(ctx)

	switch s.state {
	case attr.ValueStateKnown:
		if err := tftypes.ValidateValue(t, s.value); err != nil {
			return tftypes.NewValue(t, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(t, s.value), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(t, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(t, tftypes.UnknownValue), nil
	default:
		return tftypes.NewValue(t, tftypes.UnknownValue), fmt.Errorf("unhandled ARN state in ToTerraformValue: %s", s.state)
	}
}

func (s JSONString) Equal(other attr.Value) bool {
	o, ok := other.(JSONString)

	if !ok {
		return false
	}

	if s.state != o.state {
		return false
	}

	if s.state != attr.ValueStateKnown {
		return true
	}

	return s.value == o.value
}

// IsNull returns true if the Value is not set, or is explicitly set to null.
func (s JSONString) IsNull() bool {
	return s.state == attr.ValueStateNull
}

// IsUnknown returns true if the Value is not yet known.
func (s JSONString) IsUnknown() bool {
	return s.state == attr.ValueStateUnknown
}

// String returns a summary representation of either the underlying Value,
// or UnknownValueString (`<unknown>`) when IsUnknown() returns true,
// or NullValueString (`<null>`) when IsNull() return true.
//
// This is an intentionally lossy representation, that are best suited for
// logging and error reporting, as they are not protected by
// compatibility guarantees within the framework.
func (s JSONString) String() string {
	if s.IsUnknown() {
		return attr.UnknownValueString
	}

	if s.IsNull() {
		return attr.NullValueString
	}

	return s.value
}

// JSONStringValue returns the known string value. If JSONString is null or unknown, returns "".
func (s JSONString) ValueJSONString() string {
	return s.value
}

func (s JSONString) ToStringValue(context.Context) (basetypes.StringValue, diag.Diagnostics) {
	if s.IsUnknown() {
		return basetypes.NewStringUnknown(), nil
	}

	if s.IsNull() {
		return basetypes.NewStringNull(), nil
	}

	return basetypes.NewStringValue(s.value), nil
}

type jsonStringAttributePlanModifier struct{}

func (attributePlanModifier jsonStringAttributePlanModifier) Description(_ context.Context) string {
	return "Suppresses semantically insignificant differences."
}

func (attributePlanModifier jsonStringAttributePlanModifier) MarkdownDescription(ctx context.Context) string {
	return attributePlanModifier.Description(ctx)
}

func (attributePlanModifier jsonStringAttributePlanModifier) PlanModifyString(ctx context.Context, request planmodifier.StringRequest, response *planmodifier.StringResponse) {
	if request.StateValue.IsNull() {
		response.PlanValue = request.PlanValue

		return
	}

	// If the current value is semantically equivalent to the planned value
	// then return the current value, else return the planned value.

	plannedMap, err := expandJSONFromString(request.PlanValue.ValueString())

	if err != nil {
		response.Diagnostics.AddError(
			"Invalid JSON string (planned)",
			fmt.Sprintf("unable to unmarshal JSON: %s", err.Error()),
		)

		return
	}

	if request.StateValue.IsNull() {
		response.PlanValue = request.PlanValue
	} else {
		currentMap, err := expandJSONFromString(request.StateValue.ValueString())

		if err != nil {
			response.Diagnostics.AddError(
				"Invalid JSON string (current)",
				fmt.Sprintf("unable to unmarshal JSON: %s", err.Error()),
			)

			return
		}

		if reflect.DeepEqual(plannedMap, currentMap) {
			response.PlanValue = request.StateValue
		} else {
			response.PlanValue = request.PlanValue
		}
	}
}

func expandJSONFromString(s string) (map[string]interface{}, error) {
	var v map[string]interface{}

	err := json.Unmarshal([]byte(s), &v)

	return v, err
}

func normalizeJSONString(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	v, err := expandJSONFromString(s)

	if err != nil {
		return "", err
	}

	bytes, err := json.Marshal(v)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
