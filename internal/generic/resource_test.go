// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestPropertyPathToAttributePath(t *testing.T) {
	testCases := []struct {
		TestName      string
		PropertyPath  string
		ExpectedError bool
		ExpectedValue path.Path
	}{
		{
			TestName:      "empty property path",
			ExpectedError: true,
		},
		{
			TestName:      "no seperators",
			PropertyPath:  "test",
			ExpectedError: true,
		},
		{
			TestName:      "just root",
			PropertyPath:  "/properties",
			ExpectedError: true,
		},
		{
			TestName:      "invalid root",
			PropertyPath:  "/definitions/BasicAuthParameters/Password",
			ExpectedError: true,
		},
		{
			TestName:      "one path segment",
			PropertyPath:  "/properties/Tags",
			ExpectedValue: path.Root("tags"),
		},
		{
			TestName:      "two path segments",
			PropertyPath:  "/properties/BasicAuthParameters/Password",
			ExpectedValue: path.Root("basic_auth_parameters").AtName("password"),
		},
		{
			TestName:      "empty segment",
			PropertyPath:  "/properties//Password",
			ExpectedError: true,
		},
		{
			TestName:      "segment with *",
			PropertyPath:  "/properties/Actions/*/AuthenticateOidcConfig/ClientSecret",
			ExpectedError: true,
		},
	}

	rt := genericResource{
		cfToTfNameMap: map[string]string{
			"BasicAuthParameters": "basic_auth_parameters",
			"Password":            "password",
			"Tags":                "tags",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			attributePath, err := rt.propertyPathToAttributePath(testCase.PropertyPath)

			if err == nil && testCase.ExpectedError {
				t.Fatalf("expected error from propertyPathToAttributePath")
			}

			if err != nil && !testCase.ExpectedError {
				t.Fatalf("unexpected error from propertyPathToAttributePath: %s", err)
			}

			if err == nil && !attributePath.Equal(testCase.ExpectedValue) {
				t.Errorf("got: %s, expected: %s", attributePath, testCase.ExpectedValue)
			}
		})
	}
}

// TestMatchUnorderedListsToState verifies that after Create/Update, unknown values inside the
// elements of an unordered list are filled from the matching Cloud Control element, not from the
// element Cloud Control returned at the same index (#3110).
func TestMatchUnorderedListsToState(t *testing.T) {
	ctx := context.Background()

	groupType := tftypes.Object{AttributeTypes: map[string]tftypes.Type{
		"instance_group_name": tftypes.String,
		"current_count":       tftypes.Number,
	}}
	rootType := tftypes.Object{AttributeTypes: map[string]tftypes.Type{
		"instance_groups": tftypes.List{ElementType: groupType},
	}}
	tfSchema := schema.Schema{Attributes: map[string]schema.Attribute{
		"instance_groups": schema.ListNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"instance_group_name": schema.StringAttribute{Optional: true},
					"current_count":       schema.Int64Attribute{Computed: true},
				},
			},
			Optional:      true,
			PlanModifiers: []planmodifier.List{Multiset()},
		},
	}}
	r := genericResource{
		tfSchema:      tfSchema,
		tfToCfNameMap: map[string]string{"instance_groups": "InstanceGroups", "instance_group_name": "InstanceGroupName", "current_count": "CurrentCount"},
		cfToTfNameMap: map[string]string{"InstanceGroups": "instance_groups", "InstanceGroupName": "instance_group_name", "CurrentCount": "current_count"},
	}

	makeState := func(aCount, bCount any) tfsdk.State {
		return tfsdk.State{
			Schema: tfSchema,
			Raw: tftypes.NewValue(rootType, map[string]tftypes.Value{
				"instance_groups": tftypes.NewValue(tftypes.List{ElementType: groupType}, []tftypes.Value{
					tftypes.NewValue(groupType, map[string]tftypes.Value{
						"instance_group_name": tftypes.NewValue(tftypes.String, "worker-a"),
						"current_count":       tftypes.NewValue(tftypes.Number, aCount),
					}),
					tftypes.NewValue(groupType, map[string]tftypes.Value{
						"instance_group_name": tftypes.NewValue(tftypes.String, "worker-b"),
						"current_count":       tftypes.NewValue(tftypes.Number, bCount),
					}),
				}),
			}),
		}
	}

	// Cloud Control returns the groups in the opposite order to the planned state.
	resourceModel := `{"InstanceGroups":[{"InstanceGroupName":"worker-b","CurrentCount":2},{"InstanceGroupName":"worker-a","CurrentCount":1}]}`

	state := makeState(tftypes.UnknownValue, tftypes.UnknownValue)
	unknowns, err := UnknownValuePaths(ctx, state.Raw)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	resourceModel = r.matchUnorderedListsToState(ctx, resourceModel, &state)
	if err := SetUnknownValuesFromResourceModel(ctx, &state, unknowns, resourceModel, r.cfToTfNameMap); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if want := makeState(1, 2); !state.Raw.Equal(want.Raw) {
		t.Errorf("got: %s, expected: %s", state.Raw, want.Raw)
	}
}
