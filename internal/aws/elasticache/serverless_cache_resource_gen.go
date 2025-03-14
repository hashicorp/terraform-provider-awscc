// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package elasticache

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_elasticache_serverless_cache", serverlessCacheResource)
}

// serverlessCacheResource returns the Terraform awscc_elasticache_serverless_cache resource.
// This Terraform resource corresponds to the CloudFormation AWS::ElastiCache::ServerlessCache resource.
func serverlessCacheResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ARN
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the Serverless Cache.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the Serverless Cache.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CacheUsageLimits
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The cache capacity limit of the Serverless Cache.",
		//	  "properties": {
		//	    "DataStorage": {
		//	      "additionalProperties": false,
		//	      "description": "The cached data capacity of the Serverless Cache.",
		//	      "properties": {
		//	        "Maximum": {
		//	          "description": "The maximum cached data capacity of the Serverless Cache.",
		//	          "type": "integer"
		//	        },
		//	        "Minimum": {
		//	          "description": "The minimum cached data capacity of the Serverless Cache.",
		//	          "type": "integer"
		//	        },
		//	        "Unit": {
		//	          "description": "The unit of cached data capacity of the Serverless Cache.",
		//	          "enum": [
		//	            "GB"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "Unit"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "ECPUPerSecond": {
		//	      "additionalProperties": false,
		//	      "description": "The ECPU per second of the Serverless Cache.",
		//	      "properties": {
		//	        "Maximum": {
		//	          "description": "The maximum ECPU per second of the Serverless Cache.",
		//	          "type": "integer"
		//	        },
		//	        "Minimum": {
		//	          "description": "The minimum ECPU per second of the Serverless Cache.",
		//	          "type": "integer"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"cache_usage_limits": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DataStorage
				"data_storage": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Maximum
						"maximum": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "The maximum cached data capacity of the Serverless Cache.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
								int64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Minimum
						"minimum": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "The minimum cached data capacity of the Serverless Cache.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
								int64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Unit
						"unit": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The unit of cached data capacity of the Serverless Cache.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.OneOf(
									"GB",
								),
								fwvalidators.NotNullString(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The cached data capacity of the Serverless Cache.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ECPUPerSecond
				"ecpu_per_second": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Maximum
						"maximum": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "The maximum ECPU per second of the Serverless Cache.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
								int64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Minimum
						"minimum": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "The minimum ECPU per second of the Serverless Cache.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
								int64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The ECPU per second of the Serverless Cache.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The cache capacity limit of the Serverless Cache.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CreateTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The creation time of the Serverless Cache.",
		//	  "type": "string"
		//	}
		"create_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The creation time of the Serverless Cache.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DailySnapshotTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The daily time range (in UTC) during which the service takes automatic snapshot of the Serverless Cache.",
		//	  "type": "string"
		//	}
		"daily_snapshot_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The daily time range (in UTC) during which the service takes automatic snapshot of the Serverless Cache.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the Serverless Cache.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the Serverless Cache.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Endpoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The address and the port.",
		//	  "properties": {
		//	    "Address": {
		//	      "description": "Endpoint address.",
		//	      "type": "string"
		//	    },
		//	    "Port": {
		//	      "description": "Endpoint port.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"endpoint": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Address
				"address": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Endpoint address.",
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Port
				"port": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Endpoint port.",
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The address and the port.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Engine
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The engine name of the Serverless Cache.",
		//	  "type": "string"
		//	}
		"engine": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The engine name of the Serverless Cache.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: FinalSnapshotName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The final snapshot name which is taken before Serverless Cache is deleted.",
		//	  "type": "string"
		//	}
		"final_snapshot_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The final snapshot name which is taken before Serverless Cache is deleted.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// FinalSnapshotName is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: FullEngineVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The full engine version of the Serverless Cache.",
		//	  "type": "string"
		//	}
		"full_engine_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The full engine version of the Serverless Cache.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KmsKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the KMS key used to encrypt the cluster.",
		//	  "type": "string"
		//	}
		"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the KMS key used to encrypt the cluster.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MajorEngineVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The major engine version of the Serverless Cache.",
		//	  "type": "string"
		//	}
		"major_engine_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The major engine version of the Serverless Cache.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ReaderEndpoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The address and the port.",
		//	  "properties": {
		//	    "Address": {
		//	      "description": "Endpoint address.",
		//	      "type": "string"
		//	    },
		//	    "Port": {
		//	      "description": "Endpoint port.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"reader_endpoint": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Address
				"address": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Endpoint address.",
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Port
				"port": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Endpoint port.",
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The address and the port.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SecurityGroupIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "One or more Amazon VPC security groups associated with this Serverless Cache.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"security_group_ids": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "One or more Amazon VPC security groups associated with this Serverless Cache.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ServerlessCacheName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the Serverless Cache. This value must be unique.",
		//	  "type": "string"
		//	}
		"serverless_cache_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the Serverless Cache. This value must be unique.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SnapshotArnsToRestore
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN's of snapshot to restore Serverless Cache.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"snapshot_arns_to_restore": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The ARN's of snapshot to restore Serverless Cache.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
				setplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
			// SnapshotArnsToRestore is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: SnapshotRetentionLimit
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The snapshot retention limit of the Serverless Cache.",
		//	  "type": "integer"
		//	}
		"snapshot_retention_limit": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The snapshot retention limit of the Serverless Cache.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The status of the Serverless Cache.",
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The status of the Serverless Cache.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SubnetIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The subnet id's of the Serverless Cache.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"subnet_ids": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The subnet id's of the Serverless Cache.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
				setplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this Serverless Cache.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with Serverless Cache.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with 'aws:'. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "pattern": "^[a-zA-Z0-9 _\\.\\/=+:\\-@]*$",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with 'aws:'. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
							stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9 _\\.\\/=+:\\-@]*$"), ""),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this Serverless Cache.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: UserGroupId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the user group.",
		//	  "type": "string"
		//	}
		"user_group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the user group.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	// Corresponds to CloudFormation primaryIdentifier.
	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "The AWS::ElastiCache::ServerlessCache resource creates an Amazon ElastiCache Serverless Cache.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ElastiCache::ServerlessCache").WithTerraformTypeName("awscc_elasticache_serverless_cache")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"address":                  "Address",
		"arn":                      "ARN",
		"cache_usage_limits":       "CacheUsageLimits",
		"create_time":              "CreateTime",
		"daily_snapshot_time":      "DailySnapshotTime",
		"data_storage":             "DataStorage",
		"description":              "Description",
		"ecpu_per_second":          "ECPUPerSecond",
		"endpoint":                 "Endpoint",
		"engine":                   "Engine",
		"final_snapshot_name":      "FinalSnapshotName",
		"full_engine_version":      "FullEngineVersion",
		"key":                      "Key",
		"kms_key_id":               "KmsKeyId",
		"major_engine_version":     "MajorEngineVersion",
		"maximum":                  "Maximum",
		"minimum":                  "Minimum",
		"port":                     "Port",
		"reader_endpoint":          "ReaderEndpoint",
		"security_group_ids":       "SecurityGroupIds",
		"serverless_cache_name":    "ServerlessCacheName",
		"snapshot_arns_to_restore": "SnapshotArnsToRestore",
		"snapshot_retention_limit": "SnapshotRetentionLimit",
		"status":                   "Status",
		"subnet_ids":               "SubnetIds",
		"tags":                     "Tags",
		"unit":                     "Unit",
		"user_group_id":            "UserGroupId",
		"value":                    "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/SnapshotArnsToRestore",
		"/properties/FinalSnapshotName",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
