package generic

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type jsonStringType uint8

const (
	JSONStringType jsonStringType = iota
)

var (
	_ attr.TypeWithValidate = JSONStringType
)

func (t jsonStringType) TerraformType(_ context.Context) tftypes.Type {
	return tftypes.String
}

func (t jsonStringType) ValueFromTerraform(_ context.Context, v tftypes.Value) (attr.Value, error) {
	if !v.IsKnown() {
		return JSONString{Unknown: true}, nil
	}

	if v.IsNull() {
		return JSONString{Null: true}, nil
	}

	var s string
	err := v.As(&s)

	if err != nil {
		return nil, err
	}

	s, err = normalizeJSONString(s)

	if err != nil {
		return nil, err
	}

	return JSONString{Value: s}, nil
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

func (t jsonStringType) Validate(ctx context.Context, v tftypes.Value, path *tftypes.AttributePath) diag.Diagnostics {
	var diags diag.Diagnostics

	if !v.Type().Is(tftypes.String) {
		diags.AddAttributeError(
			path,
			"Duration Type Validation Error",
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
			path,
			"Duration Type Validation Error",
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
			path,
			"JSONString Type Validation Error",
			fmt.Sprintf("Value %q cannot be parsed as a JSON string.", v),
		)

		return diags
	}

	return diags
}

func (t jsonStringType) AttributePlanModifier() tfsdk.AttributePlanModifier {
	return jsonStringAttributePlanModifier{}
}

// A JSONString is a string containing a valid JSON document.
type JSONString struct {
	// Unknown will be true if the value is not yet known.
	Unknown bool

	// Null will be true if the value was not set, or was explicitly set to
	// null.
	Null bool

	// Value contains the set value, as long as Unknown and Null are both
	// false.
	Value string
}

func (s JSONString) Type(_ context.Context) attr.Type {
	return JSONStringType
}

func (s JSONString) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	t := JSONStringType.TerraformType(ctx)

	if s.Null {
		return tftypes.NewValue(t, nil), nil
	}

	if s.Unknown {
		return tftypes.NewValue(t, tftypes.UnknownValue), nil
	}

	if err := tftypes.ValidateValue(tftypes.String, s.Value); err != nil {
		return tftypes.NewValue(t, tftypes.UnknownValue), err
	}

	return tftypes.NewValue(t, s.Value), nil
}

func (s JSONString) Equal(other attr.Value) bool {
	o, ok := other.(JSONString)

	if !ok {
		return false
	}

	if s.Unknown != o.Unknown {
		return false
	}

	if s.Null != o.Null {
		return false
	}

	return s.Value == o.Value
}

type jsonStringAttributePlanModifier struct {
	tfsdk.AttributePlanModifier
}

func (attributePlanModifier jsonStringAttributePlanModifier) Description(_ context.Context) string {
	return "Suppresses semantically insignificant differences."
}

func (attributePlanModifier jsonStringAttributePlanModifier) MarkdownDescription(ctx context.Context) string {
	return attributePlanModifier.Description(ctx)
}

func (attributePlanModifier jsonStringAttributePlanModifier) Modify(ctx context.Context, request tfsdk.ModifyAttributePlanRequest, response *tfsdk.ModifyAttributePlanResponse) {
	if request.AttributeState == nil {
		response.AttributePlan = request.AttributePlan

		return
	}

	// If the current value is semantically equivalent to the planned value
	// then return the current value, else return the planned value.

	var planned types.String
	diags := tfsdk.ValueAs(ctx, request.AttributePlan, &planned)

	if diags.HasError() {
		response.Diagnostics = append(response.Diagnostics, diags...)

		return
	}

	plannedMap, err := expandJSONFromString(planned.Value)

	if err != nil {
		response.Diagnostics.AddError(
			"Invalid JSON string",
			fmt.Sprintf("unable to unmarshal JSON: %s", err.Error()),
		)

		return
	}

	var current types.String
	diags = tfsdk.ValueAs(ctx, request.AttributeState, &current)

	if diags.HasError() {
		response.Diagnostics = append(response.Diagnostics, diags...)

		return
	}

	currentMap, err := expandJSONFromString(current.Value)

	if err != nil {
		response.Diagnostics.AddError(
			"Invalid JSON string",
			fmt.Sprintf("unable to unmarshal JSON: %s", err.Error()),
		)

		return
	}

	if reflect.DeepEqual(plannedMap, currentMap) {
		response.AttributePlan = request.AttributeState
	} else {
		response.AttributePlan = request.AttributePlan
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
