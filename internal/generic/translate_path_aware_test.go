// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// TestPathAwareTranslateToCloudControl verifies that the path-aware translator
// correctly maps the same TF attribute name "fsx_lustre_config" to different CF
// property names depending on position in the tree:
//   - InstanceGroups/InstanceStorageConfigs → "FsxLustreConfig" (lowercase s)
//   - RestrictedInstanceGroups/InstanceStorageConfigs → "FsxLustreConfig" (lowercase s)
//   - RestrictedInstanceGroups/EnvironmentConfig → "FSxLustreConfig" (capital S)
//
// This is the core test for the FSxLustreConfig naming collision fix (GitHub issue #3019).
func TestPathAwareTranslateToCloudControl(t *testing.T) {
	plan, testSchema, pathTfToCfMap, _ := makePathAwareTestFixtures()

	expectedState := map[string]any{
		"InstanceGroups": []any{
			map[string]any{
				"InstanceGroupName": "worker-group",
				"InstanceStorageConfigs": []any{
					map[string]any{
						"FsxLustreConfig": map[string]any{
							"DnsName":   "fs-abc123.fsx.us-east-1.amazonaws.com",
							"MountName": "lustre1",
						},
					},
				},
			},
		},
		"RestrictedInstanceGroups": []any{
			map[string]any{
				"InstanceGroupName": "restricted-group",
				"InstanceStorageConfigs": []any{
					map[string]any{
						"FsxLustreConfig": map[string]any{
							"DnsName":   "fs-def456.fsx.us-east-1.amazonaws.com",
							"MountName": "lustre2",
						},
					},
				},
				"EnvironmentConfig": map[string]any{
					"FSxLustreConfig": map[string]any{
						"SizeInGiB":              float64(2400),
						"PerUnitStorageThroughput": float64(250),
					},
				},
			},
		},
	}

	translator := toCloudControl{
		pathAwareNames:    true,
		pathTfToCfNameMap: pathTfToCfMap,
	}

	got, err := translator.AsRaw(context.TODO(), testSchema, plan.Raw)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if diff := cmp.Diff(got, expectedState); diff != "" {
		t.Errorf("unexpected diff (+wanted, -got): %s", diff)
	}
}

// TestPathAwareTranslateToTerraform verifies the reverse direction: CF JSON → TF state.
// Given a Cloud Control API response with both "FSxLustreConfig" (capital S) in
// EnvironmentConfig and "FsxLustreConfig" (lowercase s) in InstanceStorageConfigs,
// verify they all map to "fsx_lustre_config" in the correct nested TF position.
func TestPathAwareTranslateToTerraform(t *testing.T) {
	_, testSchema, _, pathCfToTfMap := makePathAwareTestFixtures()

	// Simulated Cloud Control API response — note different casing
	resourceModel := map[string]any{
		"InstanceGroups": []any{
			map[string]any{
				"InstanceGroupName": "worker-group",
				"InstanceStorageConfigs": []any{
					map[string]any{
						"FsxLustreConfig": map[string]any{
							"DnsName":   "fs-abc123.fsx.us-east-1.amazonaws.com",
							"MountName": "lustre1",
						},
					},
				},
			},
		},
		"RestrictedInstanceGroups": []any{
			map[string]any{
				"InstanceGroupName": "restricted-group",
				"InstanceStorageConfigs": []any{
					map[string]any{
						"FsxLustreConfig": map[string]any{
							"DnsName":   "fs-def456.fsx.us-east-1.amazonaws.com",
							"MountName": "lustre2",
						},
					},
				},
				"EnvironmentConfig": map[string]any{
					"FSxLustreConfig": map[string]any{
						"SizeInGiB":              float64(2400),
						"PerUnitStorageThroughput": float64(250),
					},
				},
			},
		},
	}

	translator := toTerraform{
		pathAwareNames:    true,
		pathCfToTfNameMap: pathCfToTfMap,
	}

	got, err := translator.FromRaw(context.TODO(), testSchema, resourceModel)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	// Verify the result by translating back to CF — roundtrip test
	reverseTranslator := toCloudControl{
		pathAwareNames:    true,
		pathTfToCfNameMap: makePathAwareTfToCfMap(),
	}

	roundTrip, err := reverseTranslator.AsRaw(context.TODO(), testSchema, got)
	if err != nil {
		t.Fatalf("roundtrip unexpected error: %s", err)
	}

	if diff := cmp.Diff(roundTrip, resourceModel); diff != "" {
		t.Errorf("roundtrip mismatch — CF→TF→CF produced different output (+wanted, -got): %s", diff)
	}
}

// makePathAwareTestFixtures creates shared test fixtures that mimic the SageMaker Cluster
// structure with all three fsx_lustre_config locations.
func makePathAwareTestFixtures() (tfsdk.Plan, schema.Schema, map[string]string, map[string]string) {
	// Types
	fsxStorageType := tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"dns_name":   tftypes.String,
			"mount_name": tftypes.String,
		},
	}

	fsxEnvType := tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"size_in_gi_b":              tftypes.Number,
			"per_unit_storage_throughput": tftypes.Number,
		},
	}

	storageConfigType := tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"fsx_lustre_config": fsxStorageType,
		},
	}

	envConfigType := tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"fsx_lustre_config": fsxEnvType,
		},
	}

	instanceGroupType := tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"instance_group_name":      tftypes.String,
			"instance_storage_configs": tftypes.List{ElementType: storageConfigType},
		},
	}

	restrictedGroupType := tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"instance_group_name":      tftypes.String,
			"instance_storage_configs": tftypes.List{ElementType: storageConfigType},
			"environment_config":       envConfigType,
		},
	}

	rootType := tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"instance_groups":            tftypes.List{ElementType: instanceGroupType},
			"restricted_instance_groups": tftypes.List{ElementType: restrictedGroupType},
		},
	}

	// Schema
	testSchema := schema.Schema{
		Attributes: map[string]schema.Attribute{
			"instance_groups": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"instance_group_name": schema.StringAttribute{},
						"instance_storage_configs": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"fsx_lustre_config": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											"dns_name":   schema.StringAttribute{},
											"mount_name": schema.StringAttribute{},
										},
									},
								},
							},
						},
					},
				},
			},
			"restricted_instance_groups": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"instance_group_name": schema.StringAttribute{},
						"instance_storage_configs": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"fsx_lustre_config": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											"dns_name":   schema.StringAttribute{},
											"mount_name": schema.StringAttribute{},
										},
									},
								},
							},
						},
						"environment_config": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"fsx_lustre_config": schema.SingleNestedAttribute{
									Attributes: map[string]schema.Attribute{
										"size_in_gi_b":              schema.Int64Attribute{},
										"per_unit_storage_throughput": schema.Int64Attribute{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	// Plan value — all three fsx_lustre_config populated
	planValue := tftypes.NewValue(rootType, map[string]tftypes.Value{
		"instance_groups": tftypes.NewValue(tftypes.List{ElementType: instanceGroupType}, []tftypes.Value{
			tftypes.NewValue(instanceGroupType, map[string]tftypes.Value{
				"instance_group_name": tftypes.NewValue(tftypes.String, "worker-group"),
				"instance_storage_configs": tftypes.NewValue(tftypes.List{ElementType: storageConfigType}, []tftypes.Value{
					tftypes.NewValue(storageConfigType, map[string]tftypes.Value{
						"fsx_lustre_config": tftypes.NewValue(fsxStorageType, map[string]tftypes.Value{
							"dns_name":   tftypes.NewValue(tftypes.String, "fs-abc123.fsx.us-east-1.amazonaws.com"),
							"mount_name": tftypes.NewValue(tftypes.String, "lustre1"),
						}),
					}),
				}),
			}),
		}),
		"restricted_instance_groups": tftypes.NewValue(tftypes.List{ElementType: restrictedGroupType}, []tftypes.Value{
			tftypes.NewValue(restrictedGroupType, map[string]tftypes.Value{
				"instance_group_name": tftypes.NewValue(tftypes.String, "restricted-group"),
				"instance_storage_configs": tftypes.NewValue(tftypes.List{ElementType: storageConfigType}, []tftypes.Value{
					tftypes.NewValue(storageConfigType, map[string]tftypes.Value{
						"fsx_lustre_config": tftypes.NewValue(fsxStorageType, map[string]tftypes.Value{
							"dns_name":   tftypes.NewValue(tftypes.String, "fs-def456.fsx.us-east-1.amazonaws.com"),
							"mount_name": tftypes.NewValue(tftypes.String, "lustre2"),
						}),
					}),
				}),
				"environment_config": tftypes.NewValue(envConfigType, map[string]tftypes.Value{
					"fsx_lustre_config": tftypes.NewValue(fsxEnvType, map[string]tftypes.Value{
						"size_in_gi_b":              tftypes.NewValue(tftypes.Number, float64(2400)),
						"per_unit_storage_throughput": tftypes.NewValue(tftypes.Number, float64(250)),
					}),
				}),
			}),
		}),
	})

	plan := tfsdk.Plan{Raw: planValue, Schema: testSchema}

	return plan, testSchema, makePathAwareTfToCfMap(), makePathAwareCfToTfMap()
}

func makePathAwareTfToCfMap() map[string]string {
	return map[string]string{
		"instance_groups":                                                                  "InstanceGroups",
		"InstanceGroups/instance_group_name":                                               "InstanceGroupName",
		"InstanceGroups/instance_storage_configs":                                          "InstanceStorageConfigs",
		"InstanceGroups/InstanceStorageConfigs/fsx_lustre_config":                          "FsxLustreConfig",
		"InstanceGroups/InstanceStorageConfigs/FsxLustreConfig/dns_name":                   "DnsName",
		"InstanceGroups/InstanceStorageConfigs/FsxLustreConfig/mount_name":                 "MountName",
		"restricted_instance_groups":                                                       "RestrictedInstanceGroups",
		"RestrictedInstanceGroups/instance_group_name":                                     "InstanceGroupName",
		"RestrictedInstanceGroups/instance_storage_configs":                                "InstanceStorageConfigs",
		"RestrictedInstanceGroups/InstanceStorageConfigs/fsx_lustre_config":                "FsxLustreConfig",
		"RestrictedInstanceGroups/InstanceStorageConfigs/FsxLustreConfig/dns_name":         "DnsName",
		"RestrictedInstanceGroups/InstanceStorageConfigs/FsxLustreConfig/mount_name":       "MountName",
		"RestrictedInstanceGroups/environment_config":                                      "EnvironmentConfig",
		"RestrictedInstanceGroups/EnvironmentConfig/fsx_lustre_config":                     "FSxLustreConfig",
		"RestrictedInstanceGroups/EnvironmentConfig/FSxLustreConfig/size_in_gi_b":          "SizeInGiB",
		"RestrictedInstanceGroups/EnvironmentConfig/FSxLustreConfig/per_unit_storage_throughput": "PerUnitStorageThroughput",
	}
}

func makePathAwareCfToTfMap() map[string]string {
	return map[string]string{
		"InstanceGroups":                                                                   "instance_groups",
		"InstanceGroups/InstanceGroupName":                                                 "instance_group_name",
		"InstanceGroups/InstanceStorageConfigs":                                            "instance_storage_configs",
		"InstanceGroups/InstanceStorageConfigs/FsxLustreConfig":                            "fsx_lustre_config",
		"InstanceGroups/InstanceStorageConfigs/FsxLustreConfig/DnsName":                    "dns_name",
		"InstanceGroups/InstanceStorageConfigs/FsxLustreConfig/MountName":                  "mount_name",
		"RestrictedInstanceGroups":                                                         "restricted_instance_groups",
		"RestrictedInstanceGroups/InstanceGroupName":                                       "instance_group_name",
		"RestrictedInstanceGroups/InstanceStorageConfigs":                                  "instance_storage_configs",
		"RestrictedInstanceGroups/InstanceStorageConfigs/FsxLustreConfig":                  "fsx_lustre_config",
		"RestrictedInstanceGroups/InstanceStorageConfigs/FsxLustreConfig/DnsName":          "dns_name",
		"RestrictedInstanceGroups/InstanceStorageConfigs/FsxLustreConfig/MountName":        "mount_name",
		"RestrictedInstanceGroups/EnvironmentConfig":                                       "environment_config",
		"RestrictedInstanceGroups/EnvironmentConfig/FSxLustreConfig":                       "fsx_lustre_config",
		"RestrictedInstanceGroups/EnvironmentConfig/FSxLustreConfig/SizeInGiB":             "size_in_gi_b",
		"RestrictedInstanceGroups/EnvironmentConfig/FSxLustreConfig/PerUnitStorageThroughput": "per_unit_storage_throughput",
	}
}

// TestPathAwareDataSourceFromString verifies that a path-aware data source can
// correctly parse a Cloud Control API JSON response containing CF properties
// that collide after name conversion (FSxLustreConfig vs FsxLustreConfig).
// This is the same code path that the singular data source's Read() method uses.
func TestPathAwareDataSourceFromString(t *testing.T) {
	_, testSchema, _, pathCfToTfMap := makePathAwareTestFixtures()

	// Simulated Cloud Control API GetResource response — a JSON string with
	// both "FsxLustreConfig" (lowercase s) and "FSxLustreConfig" (capital S)
	// in different parts of the tree.
	apiResponse := `{
		"InstanceGroups": [{
			"InstanceGroupName": "worker-group",
			"InstanceStorageConfigs": [{
				"FsxLustreConfig": {
					"DnsName": "fs-abc123.fsx.us-east-1.amazonaws.com",
					"MountName": "lustre1"
				}
			}]
		}],
		"RestrictedInstanceGroups": [{
			"InstanceGroupName": "restricted-group",
			"InstanceStorageConfigs": [{
				"FsxLustreConfig": {
					"DnsName": "fs-def456.fsx.us-east-1.amazonaws.com",
					"MountName": "lustre2"
				}
			}],
			"EnvironmentConfig": {
				"FSxLustreConfig": {
					"SizeInGiB": 2400,
					"PerUnitStorageThroughput": 250
				}
			}
		}]
	}`

	translator := toTerraform{
		pathAwareNames:    true,
		pathCfToTfNameMap: pathCfToTfMap,
	}

	// FromString is what the singular data source's Read() calls.
	val, err := translator.FromString(context.TODO(), testSchema, apiResponse, nil)
	if err != nil {
		t.Fatalf("FromString failed: %s", err)
	}

	// Verify we can roundtrip back to CF JSON with correct property name casing.
	reverseTranslator := toCloudControl{
		pathAwareNames:    true,
		pathTfToCfNameMap: makePathAwareTfToCfMap(),
	}

	raw, err := reverseTranslator.AsRaw(context.TODO(), testSchema, val)
	if err != nil {
		t.Fatalf("roundtrip AsRaw failed: %s", err)
	}

	// Verify the CF property names have correct casing after roundtrip.
	ig := raw["InstanceGroups"].([]any)[0].(map[string]any)
	isc := ig["InstanceStorageConfigs"].([]any)[0].(map[string]any)
	if _, ok := isc["FsxLustreConfig"]; !ok {
		t.Error("expected FsxLustreConfig (lowercase s) in InstanceStorageConfigs, got keys:", keys(isc))
	}
	if _, ok := isc["FSxLustreConfig"]; ok {
		t.Error("unexpected FSxLustreConfig (capital S) in InstanceStorageConfigs")
	}

	rig := raw["RestrictedInstanceGroups"].([]any)[0].(map[string]any)
	env := rig["EnvironmentConfig"].(map[string]any)
	if _, ok := env["FSxLustreConfig"]; !ok {
		t.Error("expected FSxLustreConfig (capital S) in EnvironmentConfig, got keys:", keys(env))
	}
	if _, ok := env["FsxLustreConfig"]; ok {
		t.Error("unexpected FsxLustreConfig (lowercase s) in EnvironmentConfig")
	}
}

func keys(m map[string]any) []string {
	ks := make([]string, 0, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}
