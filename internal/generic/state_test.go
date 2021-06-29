package generic

import (
	"context"
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
		"color": {
			Type:     types.StringType,
			Optional: true,
		},
		"arn": {
			Type:     types.StringType,
			Computed: true,
		},
	},
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
				"name":  tftypes.String,
				"color": tftypes.String,
			},
		}, map[string]tftypes.Value{
			"name":  tftypes.NewValue(tftypes.String, "testing"),
			"color": tftypes.NewValue(tftypes.String, "amber"),
		}),
		Schema: testComplexSchema,
	}
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

func TestCloudFormationDesiredState(t *testing.T) {
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
				"Name":  "testing",
				"Color": "amber",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			got, err := cloudFormationDesiredStateRaw(context.TODO(), &testCase.Plan)

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
