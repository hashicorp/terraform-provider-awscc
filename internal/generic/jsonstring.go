package generic

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type jsonStringAttributePlanModifier struct {
	tfsdk.AttributePlanModifier
}

// A JSONString is a string containing a valid JSON document.
// This plan modifier suppresses semantically insignificant differences.
func JSONString() tfsdk.AttributePlanModifier {
	return jsonStringAttributePlanModifier{}
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
