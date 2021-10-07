package generic

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var testSimpleSchema = tfsdk.Schema{
	Attributes: map[string]tfsdk.Attribute{
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

var testSimpleSchemaWithList = tfsdk.Schema{
	Attributes: map[string]tfsdk.Attribute{
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
		"ports": {
			Type: types.ListType{
				ElemType: types.NumberType,
			},
			Optional: true,
		},
	},
}

var simpleCfToTfNameMap = map[string]string{
	"Arn":        "arn",
	"Identifier": "identifier",
	"Name":       "name",
	"Number":     "number",
	"Ports":      "ports",
}

// Adapted from https://github.com/hashicorp/terraform-plugin-framework/blob/1a7927fec93459115be87f283dd1ee7941b30578/tfsdk/state_test.go.
var testComplexSchema = tfsdk.Schema{
	Attributes: map[string]tfsdk.Attribute{
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
			Type: types.SetType{
				ElemType: types.StringType,
			},
			Required: true,
		},
		"disks": {
			Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
				"id": {
					Type:     types.StringType,
					Required: true,
				},
				"delete_with_instance": {
					Type:     types.BoolType,
					Optional: true,
				},
			}, tfsdk.ListNestedAttributesOptions{}),
			Optional: true,
			Computed: true,
		},
		"boot_disk": {
			Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
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
			Attributes: tfsdk.SetNestedAttributes(map[string]tfsdk.Attribute{
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
			}, tfsdk.SetNestedAttributesOptions{}),
			Optional: true,
		},
		"identifier": {
			Type:     types.StringType,
			Computed: true,
		},
	},
}

var complexCfToTfNameMap = map[string]string{
	"BootDisk":           "boot_disk",
	"DeleteWithInstance": "delete_with_instance",
	"Disks":              "disks",
	"Flags":              "flags",
	"Id":                 "id",
	"Identifier":         "identifier",
	"Interface":          "interface",
	"MachineType":        "machine_type",
	"Name":               "name",
	"Ports":              "ports",
	"ScratchDisk":        "scratch_disk",
	"Tags":               "tags",
	"VideoPorts":         "video_ports",
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

var simpleTfToCfNameMap = map[string]string{
	"arn":        "Arn",
	"identifier": "Identifier",
	"name":       "Name",
	"number":     "Number",
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

var complexTfToCfNameMap = map[string]string{
	"boot_disk":            "BootDisk",
	"delete_with_instance": "DeleteWithInstance",
	"disks":                "Disks",
	"id":                   "Id",
	"identifier":           "Identifier",
	"interface":            "Interface",
	"machine_type":         "MachineType",
	"name":                 "Name",
	"ports":                "Ports",
	"scratch_disk":         "ScratchDisk",
	"tags":                 "Tags",
}

var testMapsSchema = tfsdk.Schema{
	Attributes: map[string]tfsdk.Attribute{
		"name": {
			Type:     types.StringType,
			Required: true,
		},
		"simple_map": {
			Type: types.MapType{
				ElemType: types.StringType,
			},
			Optional: true,
		},
		"complex_map": {
			Attributes: tfsdk.MapNestedAttributes(map[string]tfsdk.Attribute{
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
			}, tfsdk.MapNestedAttributesOptions{}),
			Optional: true,
		},
	},
}

var mapsCfToTfNameMap = map[string]string{
	"Flags":      "flags",
	"Id":         "id",
	"Name":       "name",
	"SimpleMap":  "simple_map",
	"ComplexMap": "complex_map",
}

func makeMapsTestPlan() tfsdk.Plan {
	return tfsdk.Plan{
		Raw: tftypes.NewValue(tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"name":       tftypes.String,
				"simple_map": tftypes.Map{ElementType: tftypes.String},
				"complex_map": tftypes.Map{
					ElementType: videoPortElementType,
				},
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
		Schema: testMapsSchema,
	}
}

var mapsTfToCfNameMap = map[string]string{
	"flags":       "Flags",
	"id":          "Id",
	"name":        "Name",
	"simple_map":  "SimpleMap",
	"complex_map": "ComplexMap",
}
