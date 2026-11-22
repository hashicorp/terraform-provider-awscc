// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestTranslateToCloudControl(t *testing.T) {
	testCases := []struct {
		TestName      string
		Plan          tfsdk.Plan
		TfToCfNameMap map[string]string
		ExpectedError bool
		ExpectedState map[string]any
	}{
		{
			TestName:      "simple Plan",
			Plan:          makeSimpleTestPlan(),
			TfToCfNameMap: simpleTfToCfNameMap,
			ExpectedState: map[string]any{
				"Name": "testing",
			},
		},
		{
			TestName:      "simple Plan with Optional",
			Plan:          makeSimpleTestPlanWithOptionalPopulated(),
			TfToCfNameMap: simpleTfToCfNameMap,
			ExpectedState: map[string]any{
				"Name":   "testing",
				"Number": float64(42),
			},
		},
		{
			TestName:      "complex Plan",
			Plan:          makeComplexTestPlan(),
			TfToCfNameMap: complexTfToCfNameMap,
			ExpectedState: map[string]any{
				"Name":        "hello, world",
				"MachineType": "e2-medium",
				"Ports":       []any{float64(80), float64(443)},
				"Tags":        []any{"red", "blue", "green"},
				"Disks": []any{
					map[string]any{
						"Id":                 "disk0",
						"DeleteWithInstance": true,
					},
					map[string]any{
						"Id":                 "disk1",
						"DeleteWithInstance": false,
					},
				},
				"BootDisk": map[string]any{
					"Id":                 "bootdisk",
					"DeleteWithInstance": true,
				},
				"ScratchDisk": map[string]any{
					"Interface": "SCSI",
				},
			},
		},
		{
			TestName:      "maps Plan",
			Plan:          makeMapsTestPlan(),
			TfToCfNameMap: mapsTfToCfNameMap,
			ExpectedState: map[string]any{
				"Name": "testing",
				"SimpleMap": map[string]any{
					"one": "eno",
					"two": "owt",
				},
				"ComplexMap": map[string]any{
					"x": map[string]any{
						"Id":    float64(1),
						"Flags": []any{true, false},
					},
					"y": map[string]any{
						"Id":    float64(-1),
						"Flags": []any{false, true, true},
					},
				},
				"JsonString": map[string]any{
					"Key1": float64(42),
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			translator := toCloudControl{tfToCfNameMap: testCase.TfToCfNameMap}
			got, err := translator.AsRaw(context.TODO(), testCase.Plan.Schema, testCase.Plan.Raw)

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
		Schema        schema.Schema
		CfToTfNameMap map[string]string
		ResourceModel map[string]any
		ExpectedError bool
		ExpectedValue tftypes.Value
	}{
		{
			TestName:      "simple State",
			Schema:        testSimpleSchema,
			CfToTfNameMap: simpleCfToTfNameMap,
			ResourceModel: map[string]any{
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
			ResourceModel: map[string]any{
				"Arn": "arn:aws:test:::test",
				"Name": map[string]any{
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
			ResourceModel: map[string]any{
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
			ResourceModel: map[string]any{
				"Arn":    "arn:aws:test:::test",
				"Name":   "testing",
				"Number": float64(42),
				"Ports":  []any{float64(8080), float64(8443)},
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
			ResourceModel: map[string]any{
				"Arn":    "arn:aws:test:::test",
				"Name":   "testing",
				"Number": float64(42),
				"Ports":  []any{},
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
			ResourceModel: map[string]any{
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
			ResourceModel: map[string]any{
				"Name":        "hello, world",
				"MachineType": "e2-medium",
				"Ports":       []any{float64(80), float64(443)},
				"Tags":        []any{"red", "blue", "green"},
				"Disks": []any{
					map[string]any{
						"Id":                 "disk0",
						"DeleteWithInstance": true,
					},
					map[string]any{
						"Id":                 "disk1",
						"DeleteWithInstance": false,
					},
				},
				"BootDisk": map[string]any{
					"Id":                 "bootdisk",
					"DeleteWithInstance": true,
				},
				"ScratchDisk": map[string]any{
					"Interface": "SCSI",
				},
				"VideoPorts": []any{
					map[string]any{
						"Id":    float64(1),
						"Flags": []any{true, false},
					},
					map[string]any{
						"Id":    float64(-1),
						"Flags": []any{false, true, true},
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
			ResourceModel: map[string]any{
				"Name": "testing",
				"SimpleMap": map[string]any{
					"one": "eno",
					"two": "owt",
				},
				"ComplexMap": map[string]any{
					"x": map[string]any{
						"Id":    float64(1),
						"Flags": []any{true, false},
					},
					"y": map[string]any{
						"Id":    float64(-1),
						"Flags": []any{false, true, true},
					},
				},
				"JsonString": map[string]any{
					"Key1": float64(42),
				},
			},
			ExpectedValue: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"name":        tftypes.String,
					"simple_map":  tftypes.Map{ElementType: tftypes.String},
					"complex_map": tftypes.Map{ElementType: videoPortElementType},
					"json_string": tftypes.String,
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
				"json_string": tftypes.NewValue(tftypes.String, `{"Key1":42}`),
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

func TestReorderKeyValueSliceToMatch(t *testing.T) {
	testCases := []struct {
		name     string
		current  []any
		prior    []any
		expected []any
	}{
		{
			name:     "empty current",
			current:  []any{},
			prior:    []any{map[string]any{"Key": "a", "Value": "1"}},
			expected: []any{},
		},
		{
			name: "reorder to match prior",
			current: []any{
				map[string]any{"Key": "Zebra", "Value": "last"},
				map[string]any{"Key": "Apple", "Value": "first"},
				map[string]any{"Key": "Mango", "Value": "middle"},
			},
			prior: []any{
				map[string]any{"Key": "Apple", "Value": "old-first"},
				map[string]any{"Key": "Mango", "Value": "old-middle"},
				map[string]any{"Key": "Zebra", "Value": "old-last"},
			},
			expected: []any{
				map[string]any{"Key": "Apple", "Value": "first"},
				map[string]any{"Key": "Mango", "Value": "middle"},
				map[string]any{"Key": "Zebra", "Value": "last"},
			},
		},
		{
			name: "new keys appended sorted",
			current: []any{
				map[string]any{"Key": "Zebra", "Value": "z"},
				map[string]any{"Key": "Apple", "Value": "a"},
				map[string]any{"Key": "New", "Value": "n"},
			},
			prior: []any{
				map[string]any{"Key": "Apple", "Value": "old-a"},
				map[string]any{"Key": "Zebra", "Value": "old-z"},
			},
			expected: []any{
				map[string]any{"Key": "Apple", "Value": "a"},
				map[string]any{"Key": "Zebra", "Value": "z"},
				map[string]any{"Key": "New", "Value": "n"},
			},
		},
		{
			name: "lowercase key field",
			current: []any{
				map[string]any{"key": "b", "value": "2"},
				map[string]any{"key": "a", "value": "1"},
			},
			prior: []any{
				map[string]any{"key": "a", "value": "old-1"},
				map[string]any{"key": "b", "value": "old-2"},
			},
			expected: []any{
				map[string]any{"key": "a", "value": "1"},
				map[string]any{"key": "b", "value": "2"},
			},
		},
		{
			name: "multiple new keys appended in sorted order",
			current: []any{
				map[string]any{"Key": "Mango", "Value": "m"},
				map[string]any{"Key": "Apple", "Value": "a"},
				map[string]any{"Key": "Cherry", "Value": "c"},
				map[string]any{"Key": "Banana", "Value": "b"},
			},
			prior: []any{
				map[string]any{"Key": "Apple", "Value": "old-a"},
				map[string]any{"Key": "Mango", "Value": "old-m"},
			},
			expected: []any{
				map[string]any{"Key": "Apple", "Value": "a"},
				map[string]any{"Key": "Mango", "Value": "m"},
				map[string]any{"Key": "Banana", "Value": "b"},
				map[string]any{"Key": "Cherry", "Value": "c"},
			},
		},
		{
			name: "not key-value slice returns nil",
			current: []any{
				map[string]any{"NotKey": "a", "Value": "1"},
			},
			prior:    []any{},
			expected: nil,
		},
		{
			name:     "non-map element returns nil",
			current:  []any{"string"},
			prior:    []any{},
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := reorderKeyValueSliceToMatch(tc.current, tc.prior)
			if diff := cmp.Diff(got, tc.expected); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}

func TestReorderPrimitiveSliceToMatch(t *testing.T) {
	testCases := []struct {
		name     string
		current  []any
		prior    []any
		expected []any
	}{
		{
			name:     "empty current",
			current:  []any{},
			prior:    []any{"a", "b"},
			expected: []any{},
		},
		{
			name:     "reorder strings to match prior",
			current:  []any{"subnet-03", "subnet-01", "subnet-02"},
			prior:    []any{"subnet-01", "subnet-02", "subnet-03"},
			expected: []any{"subnet-01", "subnet-02", "subnet-03"},
		},
		{
			name:     "new strings appended sorted",
			current:  []any{"subnet-03", "subnet-01", "subnet-04"},
			prior:    []any{"subnet-01", "subnet-03"},
			expected: []any{"subnet-01", "subnet-03", "subnet-04"},
		},
		{
			name:     "reorder numbers to match prior",
			current:  []any{float64(443), float64(80), float64(8080)},
			prior:    []any{float64(80), float64(443), float64(8080)},
			expected: []any{float64(80), float64(443), float64(8080)},
		},
		{
			name:     "mixed types returns nil",
			current:  []any{"string", float64(42)},
			prior:    []any{},
			expected: nil,
		},
		{
			name:     "non-primitive returns nil",
			current:  []any{map[string]any{"key": "val"}},
			prior:    []any{},
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := reorderPrimitiveSliceToMatch(tc.current, tc.prior)
			if diff := cmp.Diff(got, tc.expected); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}

func TestSortSliceByKey(t *testing.T) {
	testCases := []struct {
		name     string
		input    []any
		expected []any
		sorted   bool
	}{
		{
			name: "sorts by Key field",
			input: []any{
				map[string]any{"Key": "Zebra", "Value": "z"},
				map[string]any{"Key": "Apple", "Value": "a"},
				map[string]any{"Key": "Mango", "Value": "m"},
			},
			expected: []any{
				map[string]any{"Key": "Apple", "Value": "a"},
				map[string]any{"Key": "Mango", "Value": "m"},
				map[string]any{"Key": "Zebra", "Value": "z"},
			},
			sorted: true,
		},
		{
			name: "sorts by key field lowercase",
			input: []any{
				map[string]any{"key": "z", "value": "last"},
				map[string]any{"key": "a", "value": "first"},
			},
			expected: []any{
				map[string]any{"key": "a", "value": "first"},
				map[string]any{"key": "z", "value": "last"},
			},
			sorted: true,
		},
		{
			name: "non-map returns false",
			input: []any{
				"string",
			},
			expected: []any{
				"string",
			},
			sorted: false,
		},
		{
			name: "missing key field returns false",
			input: []any{
				map[string]any{"NotKey": "a", "Value": "1"},
			},
			expected: []any{
				map[string]any{"NotKey": "a", "Value": "1"},
			},
			sorted: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := sortSliceByKey(tc.input)
			if got != tc.sorted {
				t.Errorf("expected sorted=%v, got %v", tc.sorted, got)
			}
			if diff := cmp.Diff(tc.input, tc.expected); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}

func TestNormalizeKeyValueSlices(t *testing.T) {
	testCases := []struct {
		name     string
		input    any
		expected any
	}{
		{
			name: "sorts tags by key",
			input: map[string]any{
				"Tags": []any{
					map[string]any{"Key": "Zebra", "Value": "z"},
					map[string]any{"Key": "Apple", "Value": "a"},
				},
			},
			expected: map[string]any{
				"Tags": []any{
					map[string]any{"Key": "Apple", "Value": "a"},
					map[string]any{"Key": "Zebra", "Value": "z"},
				},
			},
		},
		{
			name: "recursively sorts nested structures",
			input: map[string]any{
				"VpcConfig": map[string]any{
					"Tags": []any{
						map[string]any{"Key": "z", "Value": "last"},
						map[string]any{"Key": "a", "Value": "first"},
					},
				},
			},
			expected: map[string]any{
				"VpcConfig": map[string]any{
					"Tags": []any{
						map[string]any{"Key": "a", "Value": "first"},
						map[string]any{"Key": "z", "Value": "last"},
					},
				},
			},
		},
		{
			name: "leaves non-key-value slices unchanged",
			input: map[string]any{
				"Ports": []any{float64(443), float64(80)},
			},
			expected: map[string]any{
				"Ports": []any{float64(443), float64(80)},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			normalizeKeyValueSlices(tc.input)
			if diff := cmp.Diff(tc.input, tc.expected); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}

func TestReorderKeyValueSlicesToMatchPrior(t *testing.T) {
	testCases := []struct {
		name     string
		current  map[string]any
		prior    map[string]any
		expected map[string]any
	}{
		{
			name: "reorders tags to match prior",
			current: map[string]any{
				"Tags": []any{
					map[string]any{"Key": "Zebra", "Value": "z"},
					map[string]any{"Key": "Apple", "Value": "a"},
				},
			},
			prior: map[string]any{
				"Tags": []any{
					map[string]any{"Key": "Apple", "Value": "old-a"},
					map[string]any{"Key": "Zebra", "Value": "old-z"},
				},
			},
			expected: map[string]any{
				"Tags": []any{
					map[string]any{"Key": "Apple", "Value": "a"},
					map[string]any{"Key": "Zebra", "Value": "z"},
				},
			},
		},
		{
			name: "reorders primitive slices",
			current: map[string]any{
				"SubnetIds": []any{"subnet-03", "subnet-01", "subnet-02"},
			},
			prior: map[string]any{
				"SubnetIds": []any{"subnet-01", "subnet-02", "subnet-03"},
			},
			expected: map[string]any{
				"SubnetIds": []any{"subnet-01", "subnet-02", "subnet-03"},
			},
		},
		{
			name: "recursively reorders nested structures",
			current: map[string]any{
				"VpcConfig": map[string]any{
					"SubnetIds": []any{"subnet-02", "subnet-01"},
					"Tags": []any{
						map[string]any{"Key": "z", "Value": "last"},
						map[string]any{"Key": "a", "Value": "first"},
					},
				},
			},
			prior: map[string]any{
				"VpcConfig": map[string]any{
					"SubnetIds": []any{"subnet-01", "subnet-02"},
					"Tags": []any{
						map[string]any{"Key": "a", "Value": "old-first"},
						map[string]any{"Key": "z", "Value": "old-last"},
					},
				},
			},
			expected: map[string]any{
				"VpcConfig": map[string]any{
					"SubnetIds": []any{"subnet-01", "subnet-02"},
					"Tags": []any{
						map[string]any{"Key": "a", "Value": "first"},
						map[string]any{"Key": "z", "Value": "last"},
					},
				},
			},
		},
		{
			name: "reorders tags nested inside list elements (map-inside-a-list)",
			current: map[string]any{
				"Listeners": []any{
					map[string]any{
						"Port": "443",
						"Tags": []any{
							map[string]any{"Key": "Env", "Value": "prod"},
							map[string]any{"Key": "App", "Value": "web"},
						},
					},
				},
			},
			prior: map[string]any{
				"Listeners": []any{
					map[string]any{
						"Port": "443",
						"Tags": []any{
							map[string]any{"Key": "App", "Value": "web"},
							map[string]any{"Key": "Env", "Value": "prod"},
						},
					},
				},
			},
			expected: map[string]any{
				"Listeners": []any{
					map[string]any{
						"Port": "443",
						"Tags": []any{
							map[string]any{"Key": "App", "Value": "web"},
							map[string]any{"Key": "Env", "Value": "prod"},
						},
					},
				},
			},
		},
		{
			name: "handles missing prior gracefully",
			current: map[string]any{
				"Tags": []any{
					map[string]any{"Key": "b", "Value": "2"},
					map[string]any{"Key": "a", "Value": "1"},
				},
			},
			prior: map[string]any{},
			expected: map[string]any{
				"Tags": []any{
					map[string]any{"Key": "a", "Value": "1"},
					map[string]any{"Key": "b", "Value": "2"},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reorderKeyValueSlicesToMatchPrior(tc.current, tc.prior)
			if diff := cmp.Diff(tc.current, tc.expected); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}

func TestPrimitiveKind(t *testing.T) {
	testCases := []struct {
		name     string
		input    any
		expected int
	}{
		{
			name:     "string",
			input:    "test",
			expected: primKindString,
		},
		{
			name:     "float64",
			input:    float64(42),
			expected: primKindFloat64,
		},
		{
			name:     "map",
			input:    map[string]any{"key": "val"},
			expected: primKindOther,
		},
		{
			name:     "int",
			input:    42,
			expected: primKindOther,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := primitiveKind(tc.input)
			if got != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, got)
			}
		})
	}
}

func TestKeyFromMap(t *testing.T) {
	testCases := []struct {
		name     string
		input    map[string]any
		expected string
	}{
		{
			name:     "Key field",
			input:    map[string]any{"Key": "test", "Value": "val"},
			expected: "test",
		},
		{
			name:     "key field lowercase",
			input:    map[string]any{"key": "test", "value": "val"},
			expected: "test",
		},
		{
			name:     "no key field",
			input:    map[string]any{"NotKey": "test"},
			expected: "",
		},
		{
			name:     "non-string key",
			input:    map[string]any{"Key": 42},
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := keyFromMap(tc.input)
			if got != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, got)
			}
		})
	}
}

func BenchmarkReorderKeyValueSliceToMatch(b *testing.B) {
	current := make([]any, 50)
	prior := make([]any, 50)
	for i := range 50 {
		key := fmt.Sprintf("Key%02d", i)
		current[49-i] = map[string]any{"Key": key, "Value": fmt.Sprintf("val%d", i)}
		prior[i] = map[string]any{"Key": key, "Value": fmt.Sprintf("old%d", i)}
	}

	for b.Loop() {
		_ = reorderKeyValueSliceToMatch(current, prior)
	}
}

func BenchmarkReorderPrimitiveSliceToMatch(b *testing.B) {
	current := make([]any, 50)
	prior := make([]any, 50)
	for i := range 50 {
		current[49-i] = fmt.Sprintf("subnet-%02d", i)
		prior[i] = fmt.Sprintf("subnet-%02d", i)
	}

	for b.Loop() {
		_ = reorderPrimitiveSliceToMatch(current, prior)
	}
}

func BenchmarkReorderKeyValueSlicesToMatchPrior(b *testing.B) {
	current := map[string]any{
		"Tags": make([]any, 20),
		"VpcConfig": map[string]any{
			"SubnetIds": make([]any, 10),
		},
	}
	prior := map[string]any{
		"Tags": make([]any, 20),
		"VpcConfig": map[string]any{
			"SubnetIds": make([]any, 10),
		},
	}

	for i := range 20 {
		key := fmt.Sprintf("Tag%02d", i)
		current["Tags"].([]any)[19-i] = map[string]any{"Key": key, "Value": "v"}
		prior["Tags"].([]any)[i] = map[string]any{"Key": key, "Value": "v"}
	}
	for i := range 10 {
		current["VpcConfig"].(map[string]any)["SubnetIds"].([]any)[9-i] = fmt.Sprintf("subnet-%02d", i)
		prior["VpcConfig"].(map[string]any)["SubnetIds"].([]any)[i] = fmt.Sprintf("subnet-%02d", i)
	}

	for b.Loop() {
		reorderKeyValueSlicesToMatchPrior(current, prior)
	}
}

// instanceGroupModel returns a SageMaker Cluster InstanceGroups element (Cloud Control
// naming) with the given name and instance type, plus any extra key/value pairs.
func instanceGroupModel(name, instanceType string, extra ...any) map[string]any {
	m := map[string]any{"InstanceGroupName": name, "InstanceType": instanceType}
	for i := 0; i+1 < len(extra); i += 2 {
		m[extra[i].(string)] = extra[i+1]
	}
	return m
}

func TestIdentityField(t *testing.T) {
	group := instanceGroupModel

	testCases := []struct {
		name     string
		current  []any
		prior    []any
		expected string
	}{
		{
			name:     "empty prior",
			current:  []any{group("a", "t3")},
			expected: "",
		},
		{
			name:     "single *Name field",
			current:  []any{group("b", "m5"), group("a", "t3")},
			prior:    []any{group("a", "t3"), group("b", "m5")},
			expected: "InstanceGroupName",
		},
		{
			name:     "*Name preferred over unique *Id",
			current:  []any{group("b", "m5", "ImageId", "ami-2"), group("a", "t3", "ImageId", "ami-1")},
			prior:    []any{group("a", "t3", "ImageId", "ami-1"), group("b", "m5", "ImageId", "ami-2")},
			expected: "InstanceGroupName",
		},
		{
			name:     "exact Name preferred over *Name",
			current:  []any{map[string]any{"Name": "x", "DisplayName": "1"}, map[string]any{"Name": "y", "DisplayName": "2"}},
			prior:    []any{map[string]any{"Name": "y", "DisplayName": "2"}, map[string]any{"Name": "x", "DisplayName": "1"}},
			expected: "Name",
		},
		{
			name:     "Terraform naming",
			current:  []any{map[string]any{"instance_group_name": "b"}, map[string]any{"instance_group_name": "a"}},
			prior:    []any{map[string]any{"instance_group_name": "a"}, map[string]any{"instance_group_name": "b"}},
			expected: "instance_group_name",
		},
		{
			name:     "falls back to unique *Id when *Name is duplicated",
			current:  []any{map[string]any{"GroupName": "g", "RuleId": "r2"}, map[string]any{"GroupName": "g", "RuleId": "r1"}},
			prior:    []any{map[string]any{"GroupName": "g", "RuleId": "r1"}, map[string]any{"GroupName": "g", "RuleId": "r2"}},
			expected: "RuleId",
		},
		{
			name:     "duplicate names, no other candidate",
			current:  []any{group("a", "t3"), group("a", "m5")},
			prior:    []any{group("a", "m5"), group("a", "t3")},
			expected: "",
		},
		{
			name:     "missing on one element",
			current:  []any{group("a", "t3"), map[string]any{"InstanceType": "m5"}},
			prior:    []any{group("a", "t3")},
			expected: "",
		},
		{
			name:     "empty value",
			current:  []any{group("", "t3"), group("b", "m5")},
			prior:    []any{group("b", "m5"), group("c", "t3")},
			expected: "",
		},
		{
			name:     "no shared values",
			current:  []any{group("a", "t3")},
			prior:    []any{group("z", "t3")},
			expected: "",
		},
		{
			name:     "non-scalar candidate ignored",
			current:  []any{map[string]any{"Name": map[string]any{"x": "y"}}},
			prior:    []any{map[string]any{"Name": map[string]any{"x": "y"}}},
			expected: "",
		},
		{
			name:     "no candidate fields",
			current:  []any{map[string]any{"Port": 443.0}},
			prior:    []any{map[string]any{"Port": 443.0}},
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := identityField(tc.current, tc.prior); got != tc.expected {
				t.Errorf("identityField() = %q, want %q", got, tc.expected)
			}
		})
	}
}

func TestSliceMatcherReorder_unorderedObjectLists(t *testing.T) {
	group := instanceGroupModel
	matcher := sliceMatcher{unorderedPaths: map[string]bool{
		"/InstanceGroups":                         true,
		"/InstanceGroups/KubernetesConfig/Taints": true,
		"/InstanceGroups/AutoRollbackAlarms":      true,
		"/InstanceGroups/OnStartDeepHealthChecks": true,
	}}

	testCases := []struct {
		name     string
		current  map[string]any
		prior    map[string]any
		expected map[string]any
	}{
		{
			name: "reorders by identity field (#3110)",
			current: map[string]any{
				"InstanceGroups": []any{group("worker-c", "g5"), group("worker-a", "t3"), group("worker-b", "m5")},
			},
			prior: map[string]any{
				"InstanceGroups": []any{group("worker-a", "t3"), group("worker-b", "m5"), group("worker-c", "g5")},
			},
			expected: map[string]any{
				"InstanceGroups": []any{group("worker-a", "t3"), group("worker-b", "m5"), group("worker-c", "g5")},
			},
		},
		{
			name: "shared ImageId does not prevent matching by name",
			current: map[string]any{
				"InstanceGroups": []any{group("b", "m5", "ImageId", "ami-1"), group("a", "t3", "ImageId", "ami-1")},
			},
			prior: map[string]any{
				"InstanceGroups": []any{group("a", "t3", "ImageId", "ami-1"), group("b", "m5", "ImageId", "ami-1")},
			},
			expected: map[string]any{
				"InstanceGroups": []any{group("a", "t3", "ImageId", "ami-1"), group("b", "m5", "ImageId", "ami-1")},
			},
		},
		{
			name: "values come from current, not prior",
			current: map[string]any{
				"InstanceGroups": []any{group("b", "m5", "InstanceCount", 2.0), group("a", "t3", "InstanceCount", 1.0)},
			},
			prior: map[string]any{
				"InstanceGroups": []any{group("a", "t3", "InstanceCount", 9.0), group("b", "m5", "InstanceCount", 9.0)},
			},
			expected: map[string]any{
				"InstanceGroups": []any{group("a", "t3", "InstanceCount", 1.0), group("b", "m5", "InstanceCount", 2.0)},
			},
		},
		{
			name: "new element appended",
			current: map[string]any{
				"InstanceGroups": []any{group("c", "g5"), group("b", "m5"), group("a", "t3")},
			},
			prior: map[string]any{
				"InstanceGroups": []any{group("a", "t3"), group("b", "m5")},
			},
			expected: map[string]any{
				"InstanceGroups": []any{group("a", "t3"), group("b", "m5"), group("c", "g5")},
			},
		},
		{
			name: "removed element dropped",
			current: map[string]any{
				"InstanceGroups": []any{group("c", "g5"), group("b", "m5")},
			},
			prior: map[string]any{
				"InstanceGroups": []any{group("a", "t3"), group("b", "m5"), group("c", "g5")},
			},
			expected: map[string]any{
				"InstanceGroups": []any{group("b", "m5"), group("c", "g5")},
			},
		},
		{
			name: "duplicate identity values left as is",
			current: map[string]any{
				"InstanceGroups": []any{group("a", "m5"), group("a", "t3")},
			},
			prior: map[string]any{
				"InstanceGroups": []any{group("a", "t3"), group("a", "m5")},
			},
			expected: map[string]any{
				"InstanceGroups": []any{group("a", "m5"), group("a", "t3")},
			},
		},
		{
			name: "ordered (not unordered) object list left as is",
			current: map[string]any{
				"OrderedSteps": []any{map[string]any{"StepName": "b"}, map[string]any{"StepName": "a"}},
			},
			prior: map[string]any{
				"OrderedSteps": []any{map[string]any{"StepName": "a"}, map[string]any{"StepName": "b"}},
			},
			expected: map[string]any{
				"OrderedSteps": []any{map[string]any{"StepName": "b"}, map[string]any{"StepName": "a"}},
			},
		},
		{
			name: "nested lists are matched against the matched element, not the same index",
			current: map[string]any{
				"InstanceGroups": []any{
					group("b", "m5",
						"KubernetesConfig", map[string]any{"Taints": []any{
							map[string]any{"Key": "z", "Effect": "NoSchedule"},
							map[string]any{"Key": "y", "Effect": "NoSchedule"},
						}},
						"AutoRollbackAlarms", []any{
							map[string]any{"AlarmName": "b-2"},
							map[string]any{"AlarmName": "b-1"},
						},
						"OnStartDeepHealthChecks", []any{"InstanceStress", "InstanceConnectivity"},
					),
					group("a", "t3",
						"KubernetesConfig", map[string]any{"Taints": []any{
							map[string]any{"Key": "x", "Effect": "NoSchedule"},
						}},
					),
				},
			},
			prior: map[string]any{
				"InstanceGroups": []any{
					group("a", "t3",
						"KubernetesConfig", map[string]any{"Taints": []any{
							map[string]any{"Key": "x", "Effect": "NoSchedule"},
						}},
					),
					group("b", "m5",
						"KubernetesConfig", map[string]any{"Taints": []any{
							map[string]any{"Key": "y", "Effect": "NoSchedule"},
							map[string]any{"Key": "z", "Effect": "NoSchedule"},
						}},
						"AutoRollbackAlarms", []any{
							map[string]any{"AlarmName": "b-1"},
							map[string]any{"AlarmName": "b-2"},
						},
						"OnStartDeepHealthChecks", []any{"InstanceConnectivity", "InstanceStress"},
					),
				},
			},
			expected: map[string]any{
				"InstanceGroups": []any{
					group("a", "t3",
						"KubernetesConfig", map[string]any{"Taints": []any{
							map[string]any{"Key": "x", "Effect": "NoSchedule"},
						}},
					),
					group("b", "m5",
						"KubernetesConfig", map[string]any{"Taints": []any{
							map[string]any{"Key": "y", "Effect": "NoSchedule"},
							map[string]any{"Key": "z", "Effect": "NoSchedule"},
						}},
						"AutoRollbackAlarms", []any{
							map[string]any{"AlarmName": "b-1"},
							map[string]any{"AlarmName": "b-2"},
						},
						"OnStartDeepHealthChecks", []any{"InstanceConnectivity", "InstanceStress"},
					),
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			matcher.reorder(tc.current, tc.prior, "")
			if diff := cmp.Diff(tc.expected, tc.current); diff != "" {
				t.Errorf("unexpected diff (-wanted, +got): %s", diff)
			}
		})
	}
}

// TestSliceMatcherReorder_identityOnly verifies that, when filling unknown values after
// Create/Update, only unordered object lists are matched; other lists keep Cloud Control order.
func TestSliceMatcherReorder_identityOnly(t *testing.T) {
	group := instanceGroupModel
	matcher := sliceMatcher{unorderedPaths: map[string]bool{"/InstanceGroups": true}, identityOnly: true}

	current := map[string]any{
		"InstanceGroups": []any{group("b", "m5"), group("a", "t3")},
		"Tags":           []any{map[string]any{"Key": "z", "Value": "1"}, map[string]any{"Key": "a", "Value": "2"}},
		"SubnetIds":      []any{"subnet-2", "subnet-1"},
	}
	prior := map[string]any{
		"InstanceGroups": []any{group("a", "t3"), group("b", "m5")},
		"Tags":           []any{map[string]any{"Key": "a", "Value": "2"}, map[string]any{"Key": "z", "Value": "1"}},
		"SubnetIds":      []any{"subnet-1", "subnet-2"},
	}
	expected := map[string]any{
		"InstanceGroups": []any{group("a", "t3"), group("b", "m5")},
		"Tags":           []any{map[string]any{"Key": "z", "Value": "1"}, map[string]any{"Key": "a", "Value": "2"}},
		"SubnetIds":      []any{"subnet-2", "subnet-1"},
	}

	matcher.reorder(current, prior, "")
	if diff := cmp.Diff(expected, current); diff != "" {
		t.Errorf("unexpected diff (-wanted, +got): %s", diff)
	}
}
