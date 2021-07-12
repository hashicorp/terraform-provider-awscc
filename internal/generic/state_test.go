package generic

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var testSimpleSchema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"arn": {
			Type:     types.StringType,
			Computed: true,
		},
		"identifier": {
			Type:     types.StringType,
			Computed: true,
		},
		"name": {
			Type:     types.StringType,
			Required: true,
		},
		"number": {
			Type:     types.NumberType,
			Optional: true,
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
				"arn":        tftypes.String,
				"identifier": tftypes.String,
				"name":       tftypes.String,
				"number":     tftypes.Number,
			},
		}, map[string]tftypes.Value{
			"arn":        tftypes.NewValue(tftypes.String, nil),
			"identifier": tftypes.NewValue(tftypes.String, nil),
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
			"name":       tftypes.String,
			"number":     tftypes.Number,
			"identifier": tftypes.String,
		},
	}, map[string]tftypes.Value{
		"arn":        tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
		"name":       tftypes.NewValue(tftypes.String, "testing"),
		"number":     tftypes.NewValue(tftypes.Number, 42),
		"identifier": tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
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

func TestGetCloudFormationResourceModelValue(t *testing.T) {
	testCases := []struct {
		TestName      string
		Schema        schema.Schema
		ResourceModel map[string]interface{}
		ExpectedError bool
		ExpectedValue tftypes.Value
	}{
		{
			TestName: "simple State",
			Schema:   testSimpleSchema,
			ResourceModel: map[string]interface{}{
				"Arn":    "arn:aws:test:::test",
				"Name":   "testing",
				"Number": float64(42),
			},
			ExpectedValue: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"arn":        tftypes.String,
					"identifier": tftypes.String,
					"name":       tftypes.String,
					"number":     tftypes.Number,
				},
			}, map[string]tftypes.Value{
				"arn":        tftypes.NewValue(tftypes.String, "arn:aws:test:::test"),
				"identifier": tftypes.NewValue(tftypes.String, nil),
				"name":       tftypes.NewValue(tftypes.String, "testing"),
				"number":     tftypes.NewValue(tftypes.Number, 42),
			}),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			got, err := GetCloudFormationResourceModelRawValue(context.TODO(), &testCase.Schema, testCase.ResourceModel)

			if err == nil && testCase.ExpectedError {
				t.Fatalf("expected error from GetCloudFormationResourceModelRawValue")
			}

			if err != nil && !testCase.ExpectedError {
				t.Fatalf("unexpected error from GetCloudFormationResourceModelRawValue: %s", err)
			}

			if diff := cmp.Diff(got, testCase.ExpectedValue); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}

func TestGetUnknownValuePaths(t *testing.T) {
	testCases := []struct {
		TestName      string
		Value         tftypes.Value
		ExpectedError bool
		ExpectedPaths []UnknownValuePath
	}{
		{
			TestName: "simple State",
			Value:    makeSimpleValueWithUnknowns(),
			ExpectedPaths: []UnknownValuePath{
				{
					InTerraformState:              tftypes.NewAttributePath().WithAttributeName("arn"),
					InCloudFormationResourceModel: tftypes.NewAttributePath().WithAttributeName("Arn"),
				},
				{
					InTerraformState:              tftypes.NewAttributePath().WithAttributeName("identifier"),
					InCloudFormationResourceModel: tftypes.NewAttributePath().WithAttributeName("Identifier"),
				},
			},
		},
	}

	opts := cmp.Options{
		cmpopts.SortSlices(func(i, j UnknownValuePath) bool {
			return i.InCloudFormationResourceModel.String() < j.InCloudFormationResourceModel.String()
		}),
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			got, err := GetUnknownValuePaths(context.TODO(), testCase.Value)

			if err == nil && testCase.ExpectedError {
				t.Fatalf("expected error from GetUnknownValuePaths")
			}

			if err != nil && !testCase.ExpectedError {
				t.Fatalf("unexpected error from GetUnknownValuePaths: %s", err)
			}

			if diff := cmp.Diff(got, testCase.ExpectedPaths, opts); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}

func TestSetUnknownValuesFromCloudFormationResourceModel(t *testing.T) {
	testCases := []struct {
		TestName      string
		State         tfsdk.State
		ResourceModel map[string]interface{}
		ExpectedError bool
		ExpectedState tfsdk.State
	}{
		{
			TestName: "simple State",
			State: tfsdk.State{
				Raw:    makeSimpleValueWithUnknowns(),
				Schema: testSimpleSchema,
			},
			ResourceModel: map[string]interface{}{
				"Arn": "arn:aws:test:::test",
			},
			ExpectedState: tfsdk.State{
				Raw: tftypes.NewValue(tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"arn":        tftypes.String,
						"identifier": tftypes.String,
						"name":       tftypes.String,
						"number":     tftypes.Number,
					},
				}, map[string]tftypes.Value{
					"arn":        tftypes.NewValue(tftypes.String, "arn:aws:test:::test"),
					"identifier": tftypes.NewValue(tftypes.String, nil),
					"name":       tftypes.NewValue(tftypes.String, "testing"),
					"number":     tftypes.NewValue(tftypes.Number, 42),
				}),
				Schema: testSimpleSchema,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			err := SetUnknownValuesFromCloudFormationResourceModelRaw(context.TODO(), &testCase.State, testCase.ResourceModel)

			if err == nil && testCase.ExpectedError {
				t.Fatalf("expected error from SetUnknownValuesFromCloudFormationResourceModelRaw")
			}

			if err != nil && !testCase.ExpectedError {
				t.Fatalf("unexpected error from SetUnknownValuesFromCloudFormationResourceModelRaw: %s", err)
			}

			if diff := cmp.Diff(testCase.State, testCase.ExpectedState); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}

func makeSimpleTestPlan() tfsdk.Plan {
	return tfsdk.Plan{
		Raw: tftypes.NewValue(tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"name": tftypes.String,
			},
		}, map[string]tftypes.Value{
			"name": tftypes.NewValue(tftypes.String, "testing"),
		}),
		Schema: testComplexSchema,
	}
}

func makeSimpleTestPlanWithOptionalPopulated() tfsdk.Plan {
	return tfsdk.Plan{
		Raw: tftypes.NewValue(tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"name":   tftypes.String,
				"number": tftypes.Number,
			},
		}, map[string]tftypes.Value{
			"name":   tftypes.NewValue(tftypes.String, "testing"),
			"number": tftypes.NewValue(tftypes.Number, 42),
		}),
		Schema: testSimpleSchema,
	}
}

func makeComplexTestPlan() tfsdk.Plan {
	return tfsdk.Plan{
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

func TestPlanGetCloudFormationDesiredState(t *testing.T) {
	testCases := []struct {
		TestName      string
		Plan          tfsdk.Plan
		ExpectedError bool
		ExpectedState map[string]interface{}
	}{
		{
			TestName: "simple Plan",
			Plan:     makeSimpleTestPlan(),
			ExpectedState: map[string]interface{}{
				"Name": "testing",
			},
		},
		{
			TestName: "simple Plan with Optional",
			Plan:     makeSimpleTestPlanWithOptionalPopulated(),
			ExpectedState: map[string]interface{}{
				"Name":   "testing",
				"Number": float64(42),
			},
		},
		{
			TestName: "complex Plan",
			Plan:     makeComplexTestPlan(),
			ExpectedState: map[string]interface{}{
				"Name":        "hello, world",
				"MachineType": "e2-medium",
				"Tags":        []interface{}{"red", "blue", "green"},
				"Disks": []interface{}{
					map[string]interface{}{
						"Id":                 "disk0",
						"DeleteWithInstance": true,
					},
					map[string]interface{}{
						"Id":                 "disk1",
						"DeleteWithInstance": false,
					},
				},
				"BootDisk": map[string]interface{}{
					"Id":                 "bootdisk",
					"DeleteWithInstance": true,
				},
				"ScratchDisk": map[string]interface{}{
					"Interface": "SCSI",
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			got, err := GetCloudFormationDesiredStateRaw(context.TODO(), testCase.Plan.Raw)

			if err == nil && testCase.ExpectedError {
				t.Fatalf("expected error")
			}

			if err != nil && !testCase.ExpectedError {
				t.Fatalf("unexpected error: %s", err)
			}

			if diff := cmp.Diff(got, testCase.ExpectedState); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}
