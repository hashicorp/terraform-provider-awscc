package generic

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestTranslateToCloudControl(t *testing.T) {
	testCases := []struct {
		TestName      string
		Plan          tfsdk.Plan
		TfToCfNameMap map[string]string
		ExpectedError bool
		ExpectedState map[string]interface{}
	}{
		{
			TestName:      "simple Plan",
			Plan:          makeSimpleTestPlan(),
			TfToCfNameMap: simpleTfToCfNameMap,
			ExpectedState: map[string]interface{}{
				"Name": "testing",
			},
		},
		{
			TestName:      "simple Plan with Optional",
			Plan:          makeSimpleTestPlanWithOptionalPopulated(),
			TfToCfNameMap: simpleTfToCfNameMap,
			ExpectedState: map[string]interface{}{
				"Name":   "testing",
				"Number": float64(42),
			},
		},
		{
			TestName:      "complex Plan",
			Plan:          makeComplexTestPlan(),
			TfToCfNameMap: complexTfToCfNameMap,
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
		{
			TestName:      "maps Plan",
			Plan:          makeMapsTestPlan(),
			TfToCfNameMap: mapsTfToCfNameMap,
			ExpectedState: map[string]interface{}{
				"Name": "testing",
				"SimpleMap": map[string]interface{}{
					"one": "eno",
					"two": "owt",
				},
				"ComplexMap": map[string]interface{}{
					"x": map[string]interface{}{
						"Id":    float64(1),
						"Flags": []interface{}{true, false},
					},
					"y": map[string]interface{}{
						"Id":    float64(-1),
						"Flags": []interface{}{false, true, true},
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			translator := toCloudControl{tfToCfNameMap: testCase.TfToCfNameMap}
			got, err := translator.AsRaw(context.TODO(), testCase.Plan.Raw)

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

func TestTranslateToTerraform(t *testing.T) {
	testCases := []struct {
		TestName      string
		Schema        tfsdk.Schema
		CfToTfNameMap map[string]string
		ResourceModel map[string]interface{}
		ExpectedError bool
		ExpectedValue tftypes.Value
	}{
		{
			TestName:      "simple State",
			Schema:        testSimpleSchema,
			CfToTfNameMap: simpleCfToTfNameMap,
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
			TestName:      "simple State with JSON string",
			Schema:        testSimpleSchema,
			CfToTfNameMap: simpleCfToTfNameMap,
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
			TestName:      "simple State with extra field",
			Schema:        testSimpleSchema,
			CfToTfNameMap: simpleCfToTfNameMap,
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
			TestName:      "simple State with List",
			Schema:        testSimpleSchemaWithList,
			CfToTfNameMap: simpleCfToTfNameMap,
			ResourceModel: map[string]interface{}{
				"Arn":    "arn:aws:test:::test",
				"Name":   "testing",
				"Number": float64(42),
				"Ports":  []interface{}{float64(8080), float64(8443)},
			},
			ExpectedValue: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"arn":        tftypes.String,
					"identifier": tftypes.String,
					"name":       tftypes.String,
					"number":     tftypes.Number,
					"ports":      tftypes.List{ElementType: tftypes.Number},
				},
			}, map[string]tftypes.Value{
				"arn":        tftypes.NewValue(tftypes.String, "arn:aws:test:::test"),
				"identifier": tftypes.NewValue(tftypes.String, nil),
				"name":       tftypes.NewValue(tftypes.String, "testing"),
				"number":     tftypes.NewValue(tftypes.Number, 42),
				"ports": tftypes.NewValue(tftypes.List{ElementType: tftypes.Number}, []tftypes.Value{
					tftypes.NewValue(tftypes.Number, 8080),
					tftypes.NewValue(tftypes.Number, 8443),
				}),
			}),
		},
		{
			TestName:      "simple State with empty List",
			Schema:        testSimpleSchemaWithList,
			CfToTfNameMap: simpleCfToTfNameMap,
			ResourceModel: map[string]interface{}{
				"Arn":    "arn:aws:test:::test",
				"Name":   "testing",
				"Number": float64(42),
				"Ports":  []interface{}{},
			},
			ExpectedValue: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"arn":        tftypes.String,
					"identifier": tftypes.String,
					"name":       tftypes.String,
					"number":     tftypes.Number,
					"ports":      tftypes.List{ElementType: tftypes.Number},
				},
			}, map[string]tftypes.Value{
				"arn":        tftypes.NewValue(tftypes.String, "arn:aws:test:::test"),
				"identifier": tftypes.NewValue(tftypes.String, nil),
				"name":       tftypes.NewValue(tftypes.String, "testing"),
				"number":     tftypes.NewValue(tftypes.Number, 42),
				"ports":      tftypes.NewValue(tftypes.List{ElementType: tftypes.Number}, nil),
			}),
		},
		{
			TestName:      "simple State with missing List",
			Schema:        testSimpleSchemaWithList,
			CfToTfNameMap: simpleCfToTfNameMap,
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
					"ports":      tftypes.List{ElementType: tftypes.Number},
				},
			}, map[string]tftypes.Value{
				"arn":        tftypes.NewValue(tftypes.String, "arn:aws:test:::test"),
				"identifier": tftypes.NewValue(tftypes.String, nil),
				"name":       tftypes.NewValue(tftypes.String, "testing"),
				"number":     tftypes.NewValue(tftypes.Number, 42),
				"ports":      tftypes.NewValue(tftypes.List{ElementType: tftypes.Number}, nil),
			}),
		},
		{
			TestName:      "complex State",
			Schema:        testComplexSchema,
			CfToTfNameMap: complexCfToTfNameMap,
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
		{
			TestName:      "maps State",
			Schema:        testMapsSchema,
			CfToTfNameMap: mapsCfToTfNameMap,
			ResourceModel: map[string]interface{}{
				"Name": "testing",
				"SimpleMap": map[string]interface{}{
					"one": "eno",
					"two": "owt",
				},
				"ComplexMap": map[string]interface{}{
					"x": map[string]interface{}{
						"Id":    float64(1),
						"Flags": []interface{}{true, false},
					},
					"y": map[string]interface{}{
						"Id":    float64(-1),
						"Flags": []interface{}{false, true, true},
					},
				},
			},
			ExpectedValue: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"name":        tftypes.String,
					"simple_map":  tftypes.Map{ElementType: tftypes.String},
					"complex_map": tftypes.Map{ElementType: videoPortElementType},
				},
			}, map[string]tftypes.Value{
				"name": tftypes.NewValue(tftypes.String, "testing"),
				"simple_map": tftypes.NewValue(tftypes.Map{
					ElementType: tftypes.String,
				}, map[string]tftypes.Value{
					"one": tftypes.NewValue(tftypes.String, "eno"),
					"two": tftypes.NewValue(tftypes.String, "owt"),
				}),

				"complex_map": tftypes.NewValue(tftypes.Map{
					ElementType: videoPortElementType,
				}, map[string]tftypes.Value{
					"x": tftypes.NewValue(videoPortElementType, map[string]tftypes.Value{
						"id": tftypes.NewValue(tftypes.Number, 1),
						"flags": tftypes.NewValue(tftypes.List{
							ElementType: tftypes.Bool,
						}, []tftypes.Value{
							tftypes.NewValue(tftypes.Bool, true),
							tftypes.NewValue(tftypes.Bool, false),
						}),
					}),
					"y": tftypes.NewValue(videoPortElementType, map[string]tftypes.Value{
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
			}),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			translator := toTerraform{cfToTfNameMap: testCase.CfToTfNameMap}
			got, err := translator.FromRaw(context.TODO(), &testCase.Schema, testCase.ResourceModel)

			if err == nil && testCase.ExpectedError {
				t.Fatalf("expected error")
			}

			if err != nil && !testCase.ExpectedError {
				t.Fatalf("unexpected error: %s", err)
			}

			if err == nil {
				if diff := cmp.Diff(got, testCase.ExpectedValue); diff != "" {
					t.Errorf("unexpected diff (+wanted, -got): %s", diff)
				}
			}
		})
	}
}
