// Code generated by generators/resource/main.go; DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceFactory("awscc_elasticache_global_replication_group", globalReplicationGroupResource)
}

// globalReplicationGroupResource returns the Terraform awscc_elasticache_global_replication_group resource.
// This Terraform resource corresponds to the CloudFormation AWS::ElastiCache::GlobalReplicationGroup resource.
func globalReplicationGroupResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"automatic_failover_enabled": {
			// Property: AutomaticFailoverEnabled
			// CloudFormation resource type schema:
			// {
			//   "description": "AutomaticFailoverEnabled",
			//   "type": "boolean"
			// }
			Description: "AutomaticFailoverEnabled",
			Type:        types.BoolType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
			// AutomaticFailoverEnabled is a write-only property.
		},
		"cache_node_type": {
			// Property: CacheNodeType
			// CloudFormation resource type schema:
			// {
			//   "description": "The cache node type of the Global Datastore",
			//   "type": "string"
			// }
			Description: "The cache node type of the Global Datastore",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
			// CacheNodeType is a write-only property.
		},
		"cache_parameter_group_name": {
			// Property: CacheParameterGroupName
			// CloudFormation resource type schema:
			// {
			//   "description": "Cache parameter group name to use for the new engine version. This parameter cannot be modified independently.",
			//   "type": "string"
			// }
			Description: "Cache parameter group name to use for the new engine version. This parameter cannot be modified independently.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"engine_version": {
			// Property: EngineVersion
			// CloudFormation resource type schema:
			// {
			//   "description": "The engine version of the Global Datastore.",
			//   "type": "string"
			// }
			Description: "The engine version of the Global Datastore.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
			// EngineVersion is a write-only property.
		},
		"global_node_group_count": {
			// Property: GlobalNodeGroupCount
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates the number of node groups in the Global Datastore.",
			//   "type": "integer"
			// }
			Description: "Indicates the number of node groups in the Global Datastore.",
			Type:        types.Int64Type,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
			// GlobalNodeGroupCount is a write-only property.
		},
		"global_replication_group_description": {
			// Property: GlobalReplicationGroupDescription
			// CloudFormation resource type schema:
			// {
			//   "description": "The optional description of the Global Datastore",
			//   "type": "string"
			// }
			Description: "The optional description of the Global Datastore",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
			// GlobalReplicationGroupDescription is a write-only property.
		},
		"global_replication_group_id": {
			// Property: GlobalReplicationGroupId
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the Global Datastore, it is generated by ElastiCache adding a prefix to GlobalReplicationGroupIdSuffix.",
			//   "type": "string"
			// }
			Description: "The name of the Global Datastore, it is generated by ElastiCache adding a prefix to GlobalReplicationGroupIdSuffix.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"global_replication_group_id_suffix": {
			// Property: GlobalReplicationGroupIdSuffix
			// CloudFormation resource type schema:
			// {
			//   "description": "The suffix name of a Global Datastore. Amazon ElastiCache automatically applies a prefix to the Global Datastore ID when it is created. Each AWS Region has its own prefix. ",
			//   "type": "string"
			// }
			Description: "The suffix name of a Global Datastore. Amazon ElastiCache automatically applies a prefix to the Global Datastore ID when it is created. Each AWS Region has its own prefix. ",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
			// GlobalReplicationGroupIdSuffix is a write-only property.
		},
		"members": {
			// Property: Members
			// CloudFormation resource type schema:
			// {
			//   "description": "The replication groups that comprise the Global Datastore.",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "ReplicationGroupId": {
			//         "description": "Regionally unique identifier for the member i.e. ReplicationGroupId.",
			//         "type": "string"
			//       },
			//       "ReplicationGroupRegion": {
			//         "description": "The AWS region of the Global Datastore member.",
			//         "type": "string"
			//       },
			//       "Role": {
			//         "description": "Indicates the role of the member, primary or secondary.",
			//         "enum": [
			//           "PRIMARY",
			//           "SECONDARY"
			//         ],
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "minItems": 1,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "The replication groups that comprise the Global Datastore.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"replication_group_id": {
						// Property: ReplicationGroupId
						Description: "Regionally unique identifier for the member i.e. ReplicationGroupId.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"replication_group_region": {
						// Property: ReplicationGroupRegion
						Description: "The AWS region of the Global Datastore member.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"role": {
						// Property: Role
						Description: "Indicates the role of the member, primary or secondary.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"PRIMARY",
								"SECONDARY",
							}),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtLeast(1),
				validate.UniqueItems(),
			},
		},
		"regional_configurations": {
			// Property: RegionalConfigurations
			// CloudFormation resource type schema:
			// {
			//   "description": "Describes the replication group IDs, the AWS regions where they are stored and the shard configuration for each that comprise the Global Datastore ",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "ReplicationGroupId": {
			//         "description": "The replication group id of the Global Datastore member.",
			//         "type": "string"
			//       },
			//       "ReplicationGroupRegion": {
			//         "description": "The AWS region of the Global Datastore member.",
			//         "type": "string"
			//       },
			//       "ReshardingConfigurations": {
			//         "description": "A list of PreferredAvailabilityZones objects that specifies the configuration of a node group in the resharded cluster. ",
			//         "items": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "NodeGroupId": {
			//               "description": "Unique identifier for the Node Group. This is either auto-generated by ElastiCache (4-digit id) or a user supplied id.",
			//               "type": "string"
			//             },
			//             "PreferredAvailabilityZones": {
			//               "description": "A list of preferred availability zones for the nodes of new node groups.",
			//               "items": {
			//                 "type": "string"
			//               },
			//               "type": "array",
			//               "uniqueItems": false
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "type": "array",
			//         "uniqueItems": true
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Describes the replication group IDs, the AWS regions where they are stored and the shard configuration for each that comprise the Global Datastore ",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"replication_group_id": {
						// Property: ReplicationGroupId
						Description: "The replication group id of the Global Datastore member.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"replication_group_region": {
						// Property: ReplicationGroupRegion
						Description: "The AWS region of the Global Datastore member.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"resharding_configurations": {
						// Property: ReshardingConfigurations
						Description: "A list of PreferredAvailabilityZones objects that specifies the configuration of a node group in the resharded cluster. ",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"node_group_id": {
									// Property: NodeGroupId
									Description: "Unique identifier for the Node Group. This is either auto-generated by ElastiCache (4-digit id) or a user supplied id.",
									Type:        types.StringType,
									Optional:    true,
									Computed:    true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"preferred_availability_zones": {
									// Property: PreferredAvailabilityZones
									Description: "A list of preferred availability zones for the nodes of new node groups.",
									Type:        types.ListType{ElemType: types.StringType},
									Optional:    true,
									Computed:    true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
							},
						),
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.UniqueItems(),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.UniqueItems(),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
			// RegionalConfigurations is a write-only property.
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "The status of the Global Datastore",
			//   "type": "string"
			// }
			Description: "The status of the Global Datastore",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			resource.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "The AWS::ElastiCache::GlobalReplicationGroup resource creates an Amazon ElastiCache Global Replication Group.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ElastiCache::GlobalReplicationGroup").WithTerraformTypeName("awscc_elasticache_global_replication_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"automatic_failover_enabled":           "AutomaticFailoverEnabled",
		"cache_node_type":                      "CacheNodeType",
		"cache_parameter_group_name":           "CacheParameterGroupName",
		"engine_version":                       "EngineVersion",
		"global_node_group_count":              "GlobalNodeGroupCount",
		"global_replication_group_description": "GlobalReplicationGroupDescription",
		"global_replication_group_id":          "GlobalReplicationGroupId",
		"global_replication_group_id_suffix":   "GlobalReplicationGroupIdSuffix",
		"members":                              "Members",
		"node_group_id":                        "NodeGroupId",
		"preferred_availability_zones":         "PreferredAvailabilityZones",
		"regional_configurations":              "RegionalConfigurations",
		"replication_group_id":                 "ReplicationGroupId",
		"replication_group_region":             "ReplicationGroupRegion",
		"resharding_configurations":            "ReshardingConfigurations",
		"role":                                 "Role",
		"status":                               "Status",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/GlobalReplicationGroupIdSuffix",
		"/properties/AutomaticFailoverEnabled",
		"/properties/CacheNodeType",
		"/properties/EngineVersion",
		"/properties/GlobalNodeGroupCount",
		"/properties/GlobalReplicationGroupDescription",
		"/properties/RegionalConfigurations",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
