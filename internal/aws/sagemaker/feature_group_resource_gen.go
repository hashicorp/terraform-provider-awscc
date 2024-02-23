// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package sagemaker

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"regexp"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_sagemaker_feature_group", featureGroupResource)
}

// featureGroupResource returns the Terraform awscc_sagemaker_feature_group resource.
// This Terraform resource corresponds to the CloudFormation AWS::SageMaker::FeatureGroup resource.
func featureGroupResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A timestamp of FeatureGroup creation time.",
		//	  "type": "string"
		//	}
		"creation_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A timestamp of FeatureGroup creation time.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Description about the FeatureGroup.",
		//	  "maxLength": 128,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Description about the FeatureGroup.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(128),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EventTimeFeatureName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Event Time Feature Name.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,63}",
		//	  "type": "string"
		//	}
		"event_time_feature_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Event Time Feature Name.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,63}"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FeatureDefinitions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An Array of Feature Definition",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "FeatureName": {
		//	        "maxLength": 64,
		//	        "minLength": 1,
		//	        "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,63}",
		//	        "type": "string"
		//	      },
		//	      "FeatureType": {
		//	        "enum": [
		//	          "Integral",
		//	          "Fractional",
		//	          "String"
		//	        ],
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "FeatureName",
		//	      "FeatureType"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 2500,
		//	  "minItems": 1,
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"feature_definitions": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: FeatureName
					"feature_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 64),
							stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,63}"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: FeatureType
					"feature_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.OneOf(
								"Integral",
								"Fractional",
								"String",
							),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An Array of Feature Definition",
			Required:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(1, 2500),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FeatureGroupName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Name of the FeatureGroup.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,63}",
		//	  "type": "string"
		//	}
		"feature_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Name of the FeatureGroup.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,63}"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FeatureGroupStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The status of the feature group.",
		//	  "type": "string"
		//	}
		"feature_group_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The status of the feature group.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: OfflineStoreConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "DataCatalogConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "Catalog": {
		//	          "maxLength": 255,
		//	          "minLength": 1,
		//	          "pattern": "",
		//	          "type": "string"
		//	        },
		//	        "Database": {
		//	          "maxLength": 255,
		//	          "minLength": 1,
		//	          "pattern": "",
		//	          "type": "string"
		//	        },
		//	        "TableName": {
		//	          "maxLength": 255,
		//	          "minLength": 1,
		//	          "pattern": "",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "TableName",
		//	        "Catalog",
		//	        "Database"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "DisableGlueTableCreation": {
		//	      "type": "boolean"
		//	    },
		//	    "S3StorageConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "KmsKeyId": {
		//	          "maxLength": 2048,
		//	          "type": "string"
		//	        },
		//	        "S3Uri": {
		//	          "maxLength": 1024,
		//	          "pattern": "^(https|s3)://([^/]+)/?(.*)$",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "S3Uri"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "TableFormat": {
		//	      "description": "Format for the offline store feature group. Iceberg is the optimal format for feature groups shared between offline and online stores.",
		//	      "enum": [
		//	        "Iceberg",
		//	        "Glue"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "S3StorageConfig"
		//	  ],
		//	  "type": "object"
		//	}
		"offline_store_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DataCatalogConfig
				"data_catalog_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Catalog
						"catalog": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 255),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: Database
						"database": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 255),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: TableName
						"table_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 255),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: DisableGlueTableCreation
				"disable_glue_table_creation": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: S3StorageConfig
				"s3_storage_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: KmsKeyId
						"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthAtMost(2048),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: S3Uri
						"s3_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthAtMost(1024),
								stringvalidator.RegexMatches(regexp.MustCompile("^(https|s3)://([^/]+)/?(.*)$"), ""),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Required: true,
				}, /*END ATTRIBUTE*/
				// Property: TableFormat
				"table_format": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Format for the offline store feature group. Iceberg is the optimal format for feature groups shared between offline and online stores.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"Iceberg",
							"Glue",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: OnlineStoreConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "EnableOnlineStore": {
		//	      "type": "boolean"
		//	    },
		//	    "SecurityConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "KmsKeyId": {
		//	          "maxLength": 2048,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "StorageType": {
		//	      "enum": [
		//	        "Standard",
		//	        "InMemory"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "TtlDuration": {
		//	      "additionalProperties": false,
		//	      "description": "TTL configuration of the feature group",
		//	      "properties": {
		//	        "Unit": {
		//	          "description": "Unit of ttl configuration",
		//	          "enum": [
		//	            "Seconds",
		//	            "Minutes",
		//	            "Hours",
		//	            "Days",
		//	            "Weeks"
		//	          ],
		//	          "type": "string"
		//	        },
		//	        "Value": {
		//	          "description": "Value of ttl configuration",
		//	          "type": "integer"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"online_store_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: EnableOnlineStore
				"enable_online_store": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
						boolplanmodifier.RequiresReplace(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: SecurityConfig
				"security_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: KmsKeyId
						"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthAtMost(2048),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
						objectplanmodifier.RequiresReplace(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: StorageType
				"storage_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"Standard",
							"InMemory",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
						stringplanmodifier.RequiresReplace(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: TtlDuration
				"ttl_duration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Unit
						"unit": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Unit of ttl configuration",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.OneOf(
									"Seconds",
									"Minutes",
									"Hours",
									"Days",
									"Weeks",
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Value
						"value": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "Value of ttl configuration",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
								int64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "TTL configuration of the feature group",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RecordIdentifierFeatureName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Record Identifier Feature Name.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,63}",
		//	  "type": "string"
		//	}
		"record_identifier_feature_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Record Identifier Feature Name.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,63}"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Role Arn",
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^arn:aws[a-z\\-]*:iam::\\d{12}:role/?[a-zA-Z_0-9+=,.@\\-_/]+$",
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Role Arn",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(20, 2048),
				stringvalidator.RegexMatches(regexp.MustCompile("^arn:aws[a-z\\-]*:iam::\\d{12}:role/?[a-zA-Z_0-9+=,.@\\-_/]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pair to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pair to apply to this resource.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtMost(50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
				listplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ThroughputConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "ProvisionedReadCapacityUnits": {
		//	      "description": "For provisioned feature groups with online store enabled, this indicates the read throughput you are billed for and can consume without throttling.",
		//	      "type": "integer"
		//	    },
		//	    "ProvisionedWriteCapacityUnits": {
		//	      "description": "For provisioned feature groups, this indicates the write throughput you are billed for and can consume without throttling.",
		//	      "type": "integer"
		//	    },
		//	    "ThroughputMode": {
		//	      "description": "Throughput mode configuration of the feature group",
		//	      "enum": [
		//	        "OnDemand",
		//	        "Provisioned"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "ThroughputMode"
		//	  ],
		//	  "type": "object"
		//	}
		"throughput_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ProvisionedReadCapacityUnits
				"provisioned_read_capacity_units": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "For provisioned feature groups with online store enabled, this indicates the read throughput you are billed for and can consume without throttling.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ProvisionedWriteCapacityUnits
				"provisioned_write_capacity_units": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "For provisioned feature groups, this indicates the write throughput you are billed for and can consume without throttling.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ThroughputMode
				"throughput_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Throughput mode configuration of the feature group",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"OnDemand",
							"Provisioned",
						),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::SageMaker::FeatureGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::FeatureGroup").WithTerraformTypeName("awscc_sagemaker_feature_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"catalog":                          "Catalog",
		"creation_time":                    "CreationTime",
		"data_catalog_config":              "DataCatalogConfig",
		"database":                         "Database",
		"description":                      "Description",
		"disable_glue_table_creation":      "DisableGlueTableCreation",
		"enable_online_store":              "EnableOnlineStore",
		"event_time_feature_name":          "EventTimeFeatureName",
		"feature_definitions":              "FeatureDefinitions",
		"feature_group_name":               "FeatureGroupName",
		"feature_group_status":             "FeatureGroupStatus",
		"feature_name":                     "FeatureName",
		"feature_type":                     "FeatureType",
		"key":                              "Key",
		"kms_key_id":                       "KmsKeyId",
		"offline_store_config":             "OfflineStoreConfig",
		"online_store_config":              "OnlineStoreConfig",
		"provisioned_read_capacity_units":  "ProvisionedReadCapacityUnits",
		"provisioned_write_capacity_units": "ProvisionedWriteCapacityUnits",
		"record_identifier_feature_name":   "RecordIdentifierFeatureName",
		"role_arn":                         "RoleArn",
		"s3_storage_config":                "S3StorageConfig",
		"s3_uri":                           "S3Uri",
		"security_config":                  "SecurityConfig",
		"storage_type":                     "StorageType",
		"table_format":                     "TableFormat",
		"table_name":                       "TableName",
		"tags":                             "Tags",
		"throughput_config":                "ThroughputConfig",
		"throughput_mode":                  "ThroughputMode",
		"ttl_duration":                     "TtlDuration",
		"unit":                             "Unit",
		"value":                            "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
