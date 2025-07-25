// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package s3tables

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_s3tables_table", tableDataSource)
}

// tableDataSource returns the Terraform awscc_s3tables_table data source.
// This Terraform data source corresponds to the CloudFormation AWS::S3Tables::Table resource.
func tableDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Compaction
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Settings governing the Compaction maintenance action. Contains details about the compaction settings for an Iceberg table.",
		//	  "properties": {
		//	    "Status": {
		//	      "description": "Indicates whether the Compaction maintenance action is enabled.",
		//	      "enum": [
		//	        "enabled",
		//	        "disabled"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "TargetFileSizeMB": {
		//	      "description": "The target file size for the table in MB.",
		//	      "minimum": 64,
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"compaction": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Status
				"status": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Indicates whether the Compaction maintenance action is enabled.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: TargetFileSizeMB
				"target_file_size_mb": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The target file size for the table in MB.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Settings governing the Compaction maintenance action. Contains details about the compaction settings for an Iceberg table.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IcebergMetadata
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Contains details about the metadata for an Iceberg table.",
		//	  "properties": {
		//	    "IcebergSchema": {
		//	      "additionalProperties": false,
		//	      "description": "Contains details about the schema for an Iceberg table",
		//	      "properties": {
		//	        "SchemaFieldList": {
		//	          "description": "Contains details about the schema for an Iceberg table",
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "description": "Contains details about the schema for an Iceberg table",
		//	            "properties": {
		//	              "Name": {
		//	                "description": "The name of the field",
		//	                "type": "string"
		//	              },
		//	              "Required": {
		//	                "description": "A Boolean value that specifies whether values are required for each row in this field",
		//	                "type": "boolean"
		//	              },
		//	              "Type": {
		//	                "description": "The field type",
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "Name",
		//	              "Type"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "type": "array"
		//	        }
		//	      },
		//	      "required": [
		//	        "SchemaFieldList"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "IcebergSchema"
		//	  ],
		//	  "type": "object"
		//	}
		"iceberg_metadata": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: IcebergSchema
				"iceberg_schema": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: SchemaFieldList
						"schema_field_list": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
							NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: Name
									"name": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "The name of the field",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: Required
									"required": schema.BoolAttribute{ /*START ATTRIBUTE*/
										Description: "A Boolean value that specifies whether values are required for each row in this field",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: Type
									"type": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "The field type",
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
							}, /*END NESTED OBJECT*/
							Description: "Contains details about the schema for an Iceberg table",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Contains details about the schema for an Iceberg table",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Contains details about the metadata for an Iceberg table.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Namespace
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The namespace that the table belongs to.",
		//	  "type": "string"
		//	}
		"namespace": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The namespace that the table belongs to.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: OpenTableFormat
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Format of the table.",
		//	  "enum": [
		//	    "ICEBERG"
		//	  ],
		//	  "type": "string"
		//	}
		"open_table_format": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Format of the table.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SnapshotManagement
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Contains details about the snapshot management settings for an Iceberg table. A snapshot is expired when it exceeds MinSnapshotsToKeep and MaxSnapshotAgeHours.",
		//	  "properties": {
		//	    "MaxSnapshotAgeHours": {
		//	      "description": "The maximum age of a snapshot before it can be expired.",
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    },
		//	    "MinSnapshotsToKeep": {
		//	      "description": "The minimum number of snapshots to keep.",
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    },
		//	    "Status": {
		//	      "description": "Indicates whether the SnapshotManagement maintenance action is enabled.",
		//	      "enum": [
		//	        "enabled",
		//	        "disabled"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"snapshot_management": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: MaxSnapshotAgeHours
				"max_snapshot_age_hours": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The maximum age of a snapshot before it can be expired.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: MinSnapshotsToKeep
				"min_snapshots_to_keep": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The minimum number of snapshots to keep.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Status
				"status": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Indicates whether the SnapshotManagement maintenance action is enabled.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Contains details about the snapshot management settings for an Iceberg table. A snapshot is expired when it exceeds MinSnapshotsToKeep and MaxSnapshotAgeHours.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TableARN
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the specified table.",
		//	  "examples": [
		//	    "arn:aws:s3tables:us-west-2:123456789012:bucket/mytablebucket/table/813aadd1-a378-4d0f-8467-e3247306f309"
		//	  ],
		//	  "type": "string"
		//	}
		"table_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the specified table.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TableBucketARN
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the specified table bucket.",
		//	  "examples": [
		//	    "arn:aws:s3tables:us-west-2:123456789012:bucket/mytablebucket"
		//	  ],
		//	  "type": "string"
		//	}
		"table_bucket_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the specified table bucket.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TableName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name for the table.",
		//	  "type": "string"
		//	}
		"table_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name for the table.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VersionToken
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The version token of the table",
		//	  "type": "string"
		//	}
		"version_token": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The version token of the table",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: WarehouseLocation
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The warehouse location of the table.",
		//	  "type": "string"
		//	}
		"warehouse_location": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The warehouse location of the table.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: WithoutMetadata
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates that you don't want to specify a schema for the table. This property is mutually exclusive to 'IcebergMetadata', and its only possible value is 'Yes'.",
		//	  "enum": [
		//	    "Yes"
		//	  ],
		//	  "type": "string"
		//	}
		"without_metadata": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates that you don't want to specify a schema for the table. This property is mutually exclusive to 'IcebergMetadata', and its only possible value is 'Yes'.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::S3Tables::Table",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::S3Tables::Table").WithTerraformTypeName("awscc_s3tables_table")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"compaction":             "Compaction",
		"iceberg_metadata":       "IcebergMetadata",
		"iceberg_schema":         "IcebergSchema",
		"max_snapshot_age_hours": "MaxSnapshotAgeHours",
		"min_snapshots_to_keep":  "MinSnapshotsToKeep",
		"name":                   "Name",
		"namespace":              "Namespace",
		"open_table_format":      "OpenTableFormat",
		"required":               "Required",
		"schema_field_list":      "SchemaFieldList",
		"snapshot_management":    "SnapshotManagement",
		"status":                 "Status",
		"table_arn":              "TableARN",
		"table_bucket_arn":       "TableBucketARN",
		"table_name":             "TableName",
		"target_file_size_mb":    "TargetFileSizeMB",
		"type":                   "Type",
		"version_token":          "VersionToken",
		"warehouse_location":     "WarehouseLocation",
		"without_metadata":       "WithoutMetadata",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
