// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestUnknowns(t *testing.T) {
	testCases := []struct {
		TestName      string
		Value         tftypes.Value
		TfToCfNameMap map[string]string
		ExpectedError bool
		ExpectedPaths []*tftypes.AttributePath
	}{
		{
			TestName:      "simple State",
			Value:         makeSimpleValueWithUnknowns(),
			TfToCfNameMap: simpleTfToCfNameMap,
			ExpectedPaths: []*tftypes.AttributePath{
				tftypes.NewAttributePath().WithAttributeName("arn"),
				tftypes.NewAttributePath().WithAttributeName("identifier"),
				tftypes.NewAttributePath().WithAttributeName("ports"),
			},
		},
		{
			TestName:      "complex State",
			Value:         makeComplexValueWithUnknowns(),
			TfToCfNameMap: complexTfToCfNameMap,
			ExpectedPaths: []*tftypes.AttributePath{
				tftypes.NewAttributePath().WithAttributeName("boot_disk").WithAttributeName("delete_with_instance"),
				tftypes.NewAttributePath().WithAttributeName("disks").WithElementKeyInt(1).WithAttributeName("delete_with_instance"),
				tftypes.NewAttributePath().WithAttributeName("identifier"),
				tftypes.NewAttributePath().WithAttributeName("scratch_disk").WithAttributeName("interface"),
				tftypes.NewAttributePath().WithAttributeName("video_ports"),
			},
		},
	}

	opts := cmp.Options{
		cmpopts.SortSlices(func(i, j *tftypes.AttributePath) bool {
			return i.String() < j.String()
		}),
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			got, err := UnknownValuePaths(context.TODO(), testCase.Value)

			if err == nil && testCase.ExpectedError {
				t.Fatalf("expected error")
			}

			if err != nil && !testCase.ExpectedError {
				t.Fatalf("unexpected error: %s", err)
			}

			if err == nil {
				if diff := cmp.Diff(got, testCase.ExpectedPaths, opts); diff != "" {
					t.Errorf("unexpected diff (+wanted, -got): %s", diff)
				}
			}
		})
	}
}

func TestUnknownsSetValue(t *testing.T) {
	testCases := []struct {
		TestName      string
		State         tfsdk.State
		ResourceModel string
		CfToTfNameMap map[string]string
		ExpectedError bool
		ExpectedState tfsdk.State
	}{
		{
			TestName: "simple State",
			State: tfsdk.State{
				Raw:    makeSimpleValueWithUnknowns(),
				Schema: testSimpleSchemaWithList,
			},
			ResourceModel: `{"Arn": "arn:aws:test:::test", "Ports": [8080, 8081]}`,
			CfToTfNameMap: simpleCfToTfNameMap,
			ExpectedState: tfsdk.State{
				Raw: tftypes.NewValue(tftypes.Object{
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
					"ports": tftypes.NewValue(tftypes.List{
						ElementType: tftypes.Number,
					}, []tftypes.Value{
						tftypes.NewValue(tftypes.Number, 8080),
						tftypes.NewValue(tftypes.Number, 8081),
					}),
				}),
				Schema: testSimpleSchemaWithList,
			},
		},
		{
			TestName: "complex State",
			State: tfsdk.State{
				Raw:    makeComplexValueWithUnknowns(),
				Schema: testComplexSchema,
			},
			ResourceModel: `
			{
				"Identifier": "testing",
				"BootDisk": {"DeleteWithInstance": true},
				"ScratchDisk": {"Interface" :"SCSI"},
				"Disks": [
					{"Id": "disk0", "DeleteWithInstance": true},
					{"Id": "disk1", "DeleteWithInstance": false}
				],
				"VideoPorts": [
					{"Id": -1, "Flags": [false, true, true]},
					{"Id": 1, "Flags": [true, false]}
				]
			}`,
			CfToTfNameMap: complexCfToTfNameMap,
			ExpectedState: tfsdk.State{
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
					"identifier": tftypes.NewValue(tftypes.String, "testing"),
				}),
				Schema: testComplexSchema,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			unknowns, err := UnknownValuePaths(context.TODO(), testCase.State.Raw)

			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			err = SetUnknownValuesFromResourceModel(context.TODO(), &testCase.State, unknowns, testCase.ResourceModel, testCase.CfToTfNameMap)

			if err == nil && testCase.ExpectedError {
				t.Fatalf("expected error")
			}

			if err != nil && !testCase.ExpectedError {
				t.Fatalf("unexpected error: %s", err)
			}

			if err == nil {
				if diff := cmp.Diff(testCase.State, testCase.ExpectedState); diff != "" {
					t.Errorf("unexpected diff (+wanted, -got): %s", diff)
				}
			}
		})
	}
}
