// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var testSimpleSchema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"arn": schema.StringAttribute{
			Computed: true,
		},
		"identifier": schema.StringAttribute{
			Computed: true,
		},
		"name": schema.StringAttribute{
			Required: true,
		},
		"number": schema.NumberAttribute{
			Optional: true,
		},
	},
}

var testSimpleSchemaWithList = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"arn": schema.StringAttribute{
			Computed: true,
		},
		"identifier": schema.StringAttribute{
			Computed: true,
		},
		"name": schema.StringAttribute{
			Required: true,
		},
		"number": schema.NumberAttribute{
			Optional: true,
		},
		"ports": schema.ListAttribute{
			ElementType: types.NumberType,
			Optional:    true,
			Computed:    true,
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
var testComplexSchema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"name": schema.StringAttribute{
			Required: true,
		},
		"machine_type": schema.StringAttribute{
			Optional: true,
		},
		"ports": schema.ListAttribute{
			ElementType: types.NumberType,
			Required:    true,
		},
		"tags": schema.SetAttribute{
			ElementType: types.StringType,
			Required:    true,
		},
		"disks": schema.ListNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Required: true,
					},
					"delete_with_instance": schema.BoolAttribute{
						Optional: true,
						Computed: true,
					},
				},
			},
			Optional: true,
			Computed: true,
		},
		"boot_disk": schema.SingleNestedAttribute{
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					Required: true,
				},
				"delete_with_instance": schema.BoolAttribute{
					Optional: true,
					Computed: true,
				},
			},
		},
		"scratch_disk": schema.ObjectAttribute{
			AttributeTypes: map[string]attr.Type{
				"interface": types.StringType,
			},
			Optional: true,
		},
		"video_ports": schema.SetNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"id": schema.NumberAttribute{
						Required: true,
					},
					"flags": schema.ListAttribute{
						ElementType: types.BoolType,
						Optional:    true,
						Computed:    true,
					},
				},
			},
			Optional: true,
		},
		"identifier": schema.StringAttribute{
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
			"ports":      tftypes.List{ElementType: tftypes.Number},
		},
	}, map[string]tftypes.Value{
		"arn":        tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
		"name":       tftypes.NewValue(tftypes.String, "testing"),
		"number":     tftypes.NewValue(tftypes.Number, 42),
		"identifier": tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
		"ports":      tftypes.NewValue(tftypes.List{ElementType: tftypes.Number}, tftypes.UnknownValue),
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
				"delete_with_instance": tftypes.NewValue(tftypes.Bool, tftypes.UnknownValue),
			}),
		}),
		"boot_disk": tftypes.NewValue(diskElementType, map[string]tftypes.Value{
			"id":                   tftypes.NewValue(tftypes.String, "bootdisk"),
			"delete_with_instance": tftypes.NewValue(tftypes.Bool, tftypes.UnknownValue),
		}),
		"scratch_disk": tftypes.NewValue(tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"interface": tftypes.String,
			},
		}, map[string]tftypes.Value{
			"interface": tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
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
				}, tftypes.UnknownValue),
			}),
		}),
		"identifier": tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
	})
}

func TestCopyValueAtPath(t *testing.T) {
	testCases := []struct {
		TestName      string
		SrcState      tfsdk.State
		DstState      tfsdk.State
		Path          path.Path
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
			Path: path.Root("number"),
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
			Path: path.Root("arn"),
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
			Path: path.Root("arn"),
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
			Path:          path.Root("height"),
			ExpectedError: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			diags := copyStateValueAtPath(context.TODO(), &testCase.DstState, &testCase.SrcState, testCase.Path)

			if !diags.HasError() && testCase.ExpectedError {
				t.Fatalf("expected error from CopyValueAtPath")
			}

			if diags.HasError() && !testCase.ExpectedError {
				t.Fatalf("unexpected error from CopyValueAtPath: %s", diags)
			}

			if !diags.HasError() {
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

var testMapsSchema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"name": schema.StringAttribute{
			Required: true,
		},
		"simple_map": schema.MapAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},
		"complex_map": schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"id": schema.NumberAttribute{
						Required: true,
					},
					"flags": schema.ListAttribute{
						ElementType: types.BoolType,
						Optional:    true,
					},
				},
			},
			Optional: true,
		},
		"json_string": schema.StringAttribute{
			CustomType: jsontypes.NormalizedType{},
			Optional:   true,
		},
	},
}

var mapsCfToTfNameMap = map[string]string{
	"Flags":      "flags",
	"Id":         "id",
	"Name":       "name",
	"SimpleMap":  "simple_map",
	"ComplexMap": "complex_map",
	"JsonString": "json_string",
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
		Schema: testMapsSchema,
	}
}

var mapsTfToCfNameMap = map[string]string{
	"flags":       "Flags",
	"id":          "Id",
	"name":        "Name",
	"simple_map":  "SimpleMap",
	"complex_map": "ComplexMap",
	"json_string": "JsonString",
}
