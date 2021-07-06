package generic

import (
	"context"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var testSimpleSchema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"name": {
			Type:     types.StringType,
			Required: true,
		},
		"number": {
			Type:     types.NumberType,
			Optional: true,
		},
		"identifier": {
			Type:     types.StringType,
			Computed: true,
		},
	},
}

// Lifted from https://github.com/hashicorp/terraform-plugin-framework/blob/1a7927fec93459115be87f283dd1ee7941b30578/tfsdk/state_test.go.
var testComplexSchema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"name": {
			Type:     types.StringType,
			Required: true,
		},
		"machine_type": {
			Type: types.StringType,
		},
		"tags": {
			Type: types.ListType{
				ElemType: types.StringType,
			},
			Required: true,
		},
		"disks": {
			Attributes: schema.ListNestedAttributes(map[string]schema.Attribute{
				"id": {
					Type:     types.StringType,
					Required: true,
				},
				"delete_with_instance": {
					Type:     types.BoolType,
					Optional: true,
				},
			}, schema.ListNestedAttributesOptions{}),
			Optional: true,
			Computed: true,
		},
		"boot_disk": {
			Attributes: schema.SingleNestedAttributes(map[string]schema.Attribute{
				"id": {
					Type:     types.StringType,
					Required: true,
				},
				"delete_with_instance": {
					Type: types.BoolType,
				},
			}),
		},
		"scratch_disk": {
			Type: types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"interface": types.StringType,
				},
			},
			Optional: true,
		},
		"identifier": {
			Type:     types.StringType,
			Computed: true,
		},
	},
}

// element type for the "disks" attribute, which is a list of disks.
// only used in "disks"
var diskElementType = tftypes.Object{
	AttributeTypes: map[string]tftypes.Type{
		"id":                   tftypes.String,
		"delete_with_instance": tftypes.Bool,
	},
}

func makeSimpleTestState() tfsdk.State {
	return tfsdk.State{
		Raw: tftypes.NewValue(tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"identifier": tftypes.String,
				"name":       tftypes.String,
				"number":     tftypes.Number,
			},
		}, map[string]tftypes.Value{
			"identifier": tftypes.NewValue(tftypes.String, "???"),
			"name":       tftypes.NewValue(tftypes.String, "testing"),
			"number":     tftypes.NewValue(tftypes.Number, 42),
		}),
		Schema: testSimpleSchema,
	}
}

func makeSimpleValueWithUnknowns() tftypes.Value {
	return tftypes.NewValue(tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"arn":        tftypes.String,
			"identifier": tftypes.String,
			"name":       tftypes.String,
			"number":     tftypes.Number,
		},
	}, map[string]tftypes.Value{
		"arn":        tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
		"identifier": tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
		"name":       tftypes.NewValue(tftypes.String, "testing"),
		"number":     tftypes.NewValue(tftypes.Number, 42),
	})
}

// state used for all tests
func makeComplexTestState() tfsdk.State {
	return tfsdk.State{
		Raw: tftypes.NewValue(tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"name":         tftypes.String,
				"machine_type": tftypes.String,
				"tags":         tftypes.List{ElementType: tftypes.String},
				"disks": tftypes.List{
					ElementType: diskElementType,
				},
				"boot_disk": diskElementType,
				"scratch_disk": tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"interface": tftypes.String,
					},
				},
			},
		}, map[string]tftypes.Value{
			"name":         tftypes.NewValue(tftypes.String, "hello, world"),
			"machine_type": tftypes.NewValue(tftypes.String, "e2-medium"),
			"tags": tftypes.NewValue(tftypes.List{
				ElementType: tftypes.String,
			}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "red"),
				tftypes.NewValue(tftypes.String, "blue"),
				tftypes.NewValue(tftypes.String, "green"),
			}),
			"disks": tftypes.NewValue(tftypes.List{
				ElementType: diskElementType,
			}, []tftypes.Value{
				tftypes.NewValue(diskElementType, map[string]tftypes.Value{
					"id":                   tftypes.NewValue(tftypes.String, "disk0"),
					"delete_with_instance": tftypes.NewValue(tftypes.Bool, true),
				}),
				tftypes.NewValue(diskElementType, map[string]tftypes.Value{
					"id":                   tftypes.NewValue(tftypes.String, "disk1"),
					"delete_with_instance": tftypes.NewValue(tftypes.Bool, false),
				}),
			}),
			"boot_disk": tftypes.NewValue(diskElementType, map[string]tftypes.Value{
				"id":                   tftypes.NewValue(tftypes.String, "bootdisk"),
				"delete_with_instance": tftypes.NewValue(tftypes.Bool, true),
			}),
			"scratch_disk": tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"interface": tftypes.String,
				},
			}, map[string]tftypes.Value{
				"interface": tftypes.NewValue(tftypes.String, "SCSI"),
			}),
		}),
		Schema: testComplexSchema,
	}
}

func TestStateGetSetIdentifier(t *testing.T) {
	state := makeSimpleTestState()
	identifier := "TestID"

	err := SetIdentifier(context.TODO(), &state, identifier)

	if err != nil {
		t.Fatalf("SetIdentifier failed: %s", err)
	}

	got, err := GetIdentifier(context.TODO(), &state)

	if err != nil {
		t.Fatalf("GetIdentifier failed: %s", err)
	}

	if got != identifier {
		t.Fatalf("got: %s, expected: %s", got, identifier)
	}
}

func TestStateSetCloudFormationResourceModel(t *testing.T) {
	testCases := []struct {
		TestName      string
		State         tfsdk.State
		Raw           map[string]interface{}
		ExpectedError bool
		ExpectedState tfsdk.State
	}{
		{
			TestName: "simple State",
			State:    makeSimpleTestState(),
			Raw: map[string]interface{}{
				"Identifier": "???",
				"Name":       "testing",
				"Number":     float64(42),
			},
			ExpectedState: makeSimpleTestState(),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			err := SetCloudFormationResourceModelRaw(context.TODO(), &testCase.State, testCase.Raw)

			if err == nil && testCase.ExpectedError {
				t.Fatalf("expected error from SetCloudFormationResourceModelRaw")
			}

			if err != nil && !testCase.ExpectedError {
				t.Fatalf("unexpected error from SetCloudFormationResourceModelRaw: %s", err)
			}

			diffs, err := testCase.State.Raw.Diff(testCase.ExpectedState.Raw)

			if err != nil {
				t.Fatalf("unexpected error from Value.Diff(%s, %s): %s", testCase.State.Raw.Type(), testCase.ExpectedState.Raw.Type(), err)
			}

			if len(diffs) > 0 {
				var b strings.Builder
				for _, diff := range diffs {
					b.WriteString(diff.String())
					b.WriteString("\n")
				}
				t.Errorf("unexpected diff: %s", b.String())
			}
		})
	}
}

func TestGetAttributePathsForUnknownValues(t *testing.T) {
	testCases := []struct {
		TestName      string
		Value         tftypes.Value
		ExpectedError bool
		ExpectedPaths []*tftypes.AttributePath
	}{
		{
			TestName: "simple State",
			Value:    makeSimpleValueWithUnknowns(),
			ExpectedPaths: []*tftypes.AttributePath{
				tftypes.NewAttributePath().WithAttributeName("arn"),
				tftypes.NewAttributePath().WithAttributeName("identifier"),
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			got, err := GetAttributePathsForUnknownValues(context.TODO(), testCase.Value)

			if err == nil && testCase.ExpectedError {
				t.Fatalf("expected error from GetAttributePathsForUnknownValues")
			}

			if err != nil && !testCase.ExpectedError {
				t.Fatalf("unexpected error from GetAttributePathsForUnknownValues: %s", err)
			}

			if diff := cmp.Diff(got, testCase.ExpectedPaths); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}
