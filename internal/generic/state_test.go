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
	providertypes "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/types"
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

// Adapted from https://github.com/hashicorp/terraform-plugin-framework/blob/1a7927fec93459115be87f283dd1ee7941b30578/tfsdk/state_test.go.
var testComplexSchema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"name": {
			Type:     types.StringType,
			Required: true,
		},
		"machine_type": {
			Type:     types.StringType,
			Optional: true,
		},
		"ports": {
			Type: types.ListType{
				ElemType: types.NumberType,
			},
			Required: true,
		},
		"tags": {
			Type: providertypes.SetType{
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
					Type:     types.BoolType,
					Optional: true,
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
		"video_ports": {
			Attributes: providertypes.SetNestedAttributes(map[string]schema.Attribute{
				"id": {
					Type:     types.NumberType,
					Required: true,
				},
				"flags": {
					Type: types.ListType{
						ElemType: types.BoolType,
					},
					Optional: true,
				},
			}, providertypes.SetNestedAttributesOptions{}),
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

var videoPortElementType = tftypes.Object{
	AttributeTypes: map[string]tftypes.Type{
		"id":    tftypes.Number,
		"flags": tftypes.List{ElementType: tftypes.Bool},
	},
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

func makeComplexValueWithUnknowns() tftypes.Value {
	return tftypes.NewValue(tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"name":         tftypes.String,
			"machine_type": tftypes.String,
			"ports":        tftypes.List{ElementType: tftypes.Number},
			"tags":         tftypes.Set{ElementType: tftypes.String},
			"disks": tftypes.List{
				ElementType: diskElementType,
			},
			"boot_disk": diskElementType,
			"scratch_disk": tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"interface": tftypes.String,
				},
			},
			"identifier": tftypes.String,
		},
	}, map[string]tftypes.Value{
		"name":         tftypes.NewValue(tftypes.String, "hello, world"),
		"machine_type": tftypes.NewValue(tftypes.String, "e2-medium"),
		"ports": tftypes.NewValue(tftypes.List{
			ElementType: tftypes.Number,
		}, []tftypes.Value{
			tftypes.NewValue(tftypes.Number, 80),
			tftypes.NewValue(tftypes.Number, 443),
		}),
		"tags": tftypes.NewValue(tftypes.Set{
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
		"identifier": tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
	})
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
		{
			TestName: "simple State with JSON string",
			Schema:   testSimpleSchema,
			ResourceModel: map[string]interface{}{
				"Arn": "arn:aws:test:::test",
				"Name": map[string]interface{}{
					"Value": "testing",
				},
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
				"name":       tftypes.NewValue(tftypes.String, `{"Value":"testing"}`),
				"number":     tftypes.NewValue(tftypes.Number, 42),
			}),
		},
		{
			TestName: "simple State with extra field",
			Schema:   testSimpleSchema,
			ResourceModel: map[string]interface{}{
				"Arn":    "arn:aws:test:::test",
				"Height": float64(1.75),
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
		{
			TestName: "complex State",
			Schema:   testComplexSchema,
			ResourceModel: map[string]interface{}{
				"Name":        "hello, world",
				"MachineType": "e2-medium",
				"Ports":       []interface{}{float64(80), float64(443)},
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
				"VideoPorts": []interface{}{
					map[string]interface{}{
						"Id":    float64(1),
						"Flags": []interface{}{true, false},
					},
					map[string]interface{}{
						"Id":    float64(-1),
						"Flags": []interface{}{false, true, true},
					},
				},
			},
			ExpectedValue: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"name":         tftypes.String,
					"machine_type": tftypes.String,
					"ports":        tftypes.List{ElementType: tftypes.Number},
					"tags":         tftypes.Set{ElementType: tftypes.String},
					"disks": tftypes.List{
						ElementType: diskElementType,
					},
					"boot_disk": diskElementType,
					"scratch_disk": tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"interface": tftypes.String,
						},
					},
					"video_ports": tftypes.Set{
						ElementType: videoPortElementType,
					},
					"identifier": tftypes.String,
				},
			}, map[string]tftypes.Value{
				"name":         tftypes.NewValue(tftypes.String, "hello, world"),
				"machine_type": tftypes.NewValue(tftypes.String, "e2-medium"),
				"ports": tftypes.NewValue(tftypes.List{
					ElementType: tftypes.Number,
				}, []tftypes.Value{
					tftypes.NewValue(tftypes.Number, 80),
					tftypes.NewValue(tftypes.Number, 443),
				}),
				"tags": tftypes.NewValue(tftypes.Set{
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
				"video_ports": tftypes.NewValue(tftypes.Set{
					ElementType: videoPortElementType,
				}, []tftypes.Value{
					tftypes.NewValue(videoPortElementType, map[string]tftypes.Value{
						"id": tftypes.NewValue(tftypes.Number, 1),
						"flags": tftypes.NewValue(tftypes.List{
							ElementType: tftypes.Bool,
						}, []tftypes.Value{
							tftypes.NewValue(tftypes.Bool, true),
							tftypes.NewValue(tftypes.Bool, false),
						}),
					}),
					tftypes.NewValue(videoPortElementType, map[string]tftypes.Value{
						"id": tftypes.NewValue(tftypes.Number, -1),
						"flags": tftypes.NewValue(tftypes.List{
							ElementType: tftypes.Bool,
						}, []tftypes.Value{
							tftypes.NewValue(tftypes.Bool, false),
							tftypes.NewValue(tftypes.Bool, true),
							tftypes.NewValue(tftypes.Bool, true),
						}),
					}),
				}),
				"identifier": tftypes.NewValue(tftypes.String, nil),
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

			if err == nil {
				if diff := cmp.Diff(got, testCase.ExpectedValue); diff != "" {
					t.Errorf("unexpected diff (+wanted, -got): %s", diff)
				}
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
		{
			TestName: "complex State",
			Value:    makeComplexValueWithUnknowns(),
			ExpectedPaths: []UnknownValuePath{
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

			if err == nil {
				if diff := cmp.Diff(got, testCase.ExpectedPaths, opts); diff != "" {
					t.Errorf("unexpected diff (+wanted, -got): %s", diff)
				}
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

			if err == nil {
				if diff := cmp.Diff(testCase.State, testCase.ExpectedState); diff != "" {
					t.Errorf("unexpected diff (+wanted, -got): %s", diff)
				}
			}
		})
	}
}

func TestCopyValueAtPath(t *testing.T) {
	testCases := []struct {
		TestName      string
		SrcState      tfsdk.State
		DstState      tfsdk.State
		Path          *tftypes.AttributePath
		ExpectedError bool
		ExpectedState tfsdk.State
	}{
		{
			TestName: "simple State",
			SrcState: tfsdk.State{
				Raw: tftypes.NewValue(tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"arn":        tftypes.String,
						"name":       tftypes.String,
						"number":     tftypes.Number,
						"identifier": tftypes.String,
					},
				}, map[string]tftypes.Value{
					"arn":        tftypes.NewValue(tftypes.String, "arnsrc"),
					"name":       tftypes.NewValue(tftypes.String, "namesrc"),
					"number":     tftypes.NewValue(tftypes.Number, 42),
					"identifier": tftypes.NewValue(tftypes.String, "idsrc"),
				}),
				Schema: testSimpleSchema,
			},
			DstState: tfsdk.State{
				Raw: tftypes.NewValue(tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"arn":        tftypes.String,
						"name":       tftypes.String,
						"number":     tftypes.Number,
						"identifier": tftypes.String,
					},
				}, map[string]tftypes.Value{
					"arn":        tftypes.NewValue(tftypes.String, "arndest"),
					"name":       tftypes.NewValue(tftypes.String, "namedest"),
					"number":     tftypes.NewValue(tftypes.Number, 0),
					"identifier": tftypes.NewValue(tftypes.String, "iddest"),
				}),
				Schema: testSimpleSchema,
			},
			Path: tftypes.NewAttributePath().WithAttributeName("number"),
			ExpectedState: tfsdk.State{
				Raw: tftypes.NewValue(tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"arn":        tftypes.String,
						"name":       tftypes.String,
						"number":     tftypes.Number,
						"identifier": tftypes.String,
					},
				}, map[string]tftypes.Value{
					"arn":        tftypes.NewValue(tftypes.String, "arndest"),
					"name":       tftypes.NewValue(tftypes.String, "namedest"),
					"number":     tftypes.NewValue(tftypes.Number, 42),
					"identifier": tftypes.NewValue(tftypes.String, "iddest"),
				}),
				Schema: testSimpleSchema,
			},
		},
		{
			TestName: "simple State with Null in Src",
			SrcState: tfsdk.State{
				Raw: tftypes.NewValue(tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"arn":        tftypes.String,
						"name":       tftypes.String,
						"number":     tftypes.Number,
						"identifier": tftypes.String,
					},
				}, map[string]tftypes.Value{
					"arn":        tftypes.NewValue(tftypes.String, nil),
					"name":       tftypes.NewValue(tftypes.String, "namesrc"),
					"number":     tftypes.NewValue(tftypes.Number, 42),
					"identifier": tftypes.NewValue(tftypes.String, "idsrc"),
				}),
				Schema: testSimpleSchema,
			},
			DstState: tfsdk.State{
				Raw: tftypes.NewValue(tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"arn":        tftypes.String,
						"name":       tftypes.String,
						"number":     tftypes.Number,
						"identifier": tftypes.String,
					},
				}, map[string]tftypes.Value{
					"arn":        tftypes.NewValue(tftypes.String, "arndest"),
					"name":       tftypes.NewValue(tftypes.String, "namedest"),
					"number":     tftypes.NewValue(tftypes.Number, 43),
					"identifier": tftypes.NewValue(tftypes.String, "iddest"),
				}),
				Schema: testSimpleSchema,
			},
			Path: tftypes.NewAttributePath().WithAttributeName("arn"),
			ExpectedState: tfsdk.State{
				Raw: tftypes.NewValue(tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"arn":        tftypes.String,
						"name":       tftypes.String,
						"number":     tftypes.Number,
						"identifier": tftypes.String,
					},
				}, map[string]tftypes.Value{
					"arn":        tftypes.NewValue(tftypes.String, nil),
					"name":       tftypes.NewValue(tftypes.String, "namedest"),
					"number":     tftypes.NewValue(tftypes.Number, 43),
					"identifier": tftypes.NewValue(tftypes.String, "iddest"),
				}),
				Schema: testSimpleSchema,
			},
		},
		{
			TestName: "simple State with Null in Dst",
			SrcState: tfsdk.State{
				Raw: tftypes.NewValue(tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"arn":        tftypes.String,
						"name":       tftypes.String,
						"number":     tftypes.Number,
						"identifier": tftypes.String,
					},
				}, map[string]tftypes.Value{
					"arn":        tftypes.NewValue(tftypes.String, "arnsrc"),
					"name":       tftypes.NewValue(tftypes.String, "namesrc"),
					"number":     tftypes.NewValue(tftypes.Number, 42),
					"identifier": tftypes.NewValue(tftypes.String, "idsrc"),
				}),
				Schema: testSimpleSchema,
			},
			DstState: tfsdk.State{
				Raw: tftypes.NewValue(tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"arn":        tftypes.String,
						"name":       tftypes.String,
						"number":     tftypes.Number,
						"identifier": tftypes.String,
					},
				}, map[string]tftypes.Value{
					"arn":        tftypes.NewValue(tftypes.String, nil),
					"name":       tftypes.NewValue(tftypes.String, "namedest"),
					"number":     tftypes.NewValue(tftypes.Number, 43),
					"identifier": tftypes.NewValue(tftypes.String, "iddest"),
				}),
				Schema: testSimpleSchema,
			},
			Path: tftypes.NewAttributePath().WithAttributeName("arn"),
			ExpectedState: tfsdk.State{
				Raw: tftypes.NewValue(tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"arn":        tftypes.String,
						"name":       tftypes.String,
						"number":     tftypes.Number,
						"identifier": tftypes.String,
					},
				}, map[string]tftypes.Value{
					"arn":        tftypes.NewValue(tftypes.String, "arnsrc"),
					"name":       tftypes.NewValue(tftypes.String, "namedest"),
					"number":     tftypes.NewValue(tftypes.Number, 43),
					"identifier": tftypes.NewValue(tftypes.String, "iddest"),
				}),
				Schema: testSimpleSchema,
			},
		},
		{
			TestName: "invalid Path",
			SrcState: tfsdk.State{
				Raw: tftypes.NewValue(tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"arn":        tftypes.String,
						"name":       tftypes.String,
						"number":     tftypes.Number,
						"identifier": tftypes.String,
					},
				}, map[string]tftypes.Value{
					"arn":        tftypes.NewValue(tftypes.String, "arnsrc"),
					"name":       tftypes.NewValue(tftypes.String, "namesrc"),
					"number":     tftypes.NewValue(tftypes.Number, 42),
					"identifier": tftypes.NewValue(tftypes.String, "idsrc"),
				}),
				Schema: testSimpleSchema,
			},
			DstState: tfsdk.State{
				Raw: tftypes.NewValue(tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"arn":        tftypes.String,
						"name":       tftypes.String,
						"number":     tftypes.Number,
						"identifier": tftypes.String,
					},
				}, map[string]tftypes.Value{
					"arn":        tftypes.NewValue(tftypes.String, "arndest"),
					"name":       tftypes.NewValue(tftypes.String, "namedest"),
					"number":     tftypes.NewValue(tftypes.Number, 0),
					"identifier": tftypes.NewValue(tftypes.String, "iddest"),
				}),
				Schema: testSimpleSchema,
			},
			Path:          tftypes.NewAttributePath().WithAttributeName("height"),
			ExpectedError: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			err := CopyValueAtPath(context.TODO(), &testCase.DstState, &testCase.SrcState, testCase.Path)

			if err == nil && testCase.ExpectedError {
				t.Fatalf("expected error from CopyValueAtPath")
			}

			if err != nil && !testCase.ExpectedError {
				t.Fatalf("unexpected error from CopyValueAtPath: %s", err)
			}

			if err == nil {
				if diff := cmp.Diff(testCase.DstState, testCase.ExpectedState); diff != "" {
					t.Errorf("unexpected diff (+wanted, -got): %s", diff)
				}
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
		Schema: testSimpleSchema,
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
				"ports":        tftypes.List{ElementType: tftypes.Number},
				"tags":         tftypes.Set{ElementType: tftypes.String},
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
			"ports": tftypes.NewValue(tftypes.List{
				ElementType: tftypes.Number,
			}, []tftypes.Value{
				tftypes.NewValue(tftypes.Number, 80),
				tftypes.NewValue(tftypes.Number, 443),
			}),
			"tags": tftypes.NewValue(tftypes.Set{
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
				"Ports":       []interface{}{float64(80), float64(443)},
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

			if err == nil {
				if diff := cmp.Diff(got, testCase.ExpectedState); diff != "" {
					t.Errorf("unexpected diff (+wanted, -got): %s", diff)
				}
			}
		})
	}
}
