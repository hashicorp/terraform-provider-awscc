// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ssm_resource_data_sync", resourceDataSyncDataSource)
}

// resourceDataSyncDataSource returns the Terraform awscc_ssm_resource_data_sync data source.
// This Terraform data source corresponds to the CloudFormation AWS::SSM::ResourceDataSync resource.
func resourceDataSyncDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: BucketName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: BucketPrefix
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 64,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"bucket_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: BucketRegion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"bucket_region": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: KMSKeyArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 512,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"kms_key_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: S3Destination
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "BucketName": {
		//	      "maxLength": 2048,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "BucketPrefix": {
		//	      "maxLength": 256,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "BucketRegion": {
		//	      "maxLength": 64,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "KMSKeyArn": {
		//	      "maxLength": 512,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "SyncFormat": {
		//	      "maxLength": 1024,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "BucketName",
		//	    "BucketRegion",
		//	    "SyncFormat"
		//	  ],
		//	  "type": "object"
		//	}
		"s3_destination": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: BucketName
				"bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: BucketPrefix
				"bucket_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: BucketRegion
				"bucket_region": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: KMSKeyArn
				"kms_key_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: SyncFormat
				"sync_format": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: SyncFormat
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 1024,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"sync_format": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: SyncName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"sync_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: SyncSource
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "AwsOrganizationsSource": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "OrganizationSourceType": {
		//	          "maxLength": 64,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "OrganizationalUnits": {
		//	          "items": {
		//	            "type": "string"
		//	          },
		//	          "type": "array",
		//	          "uniqueItems": false
		//	        }
		//	      },
		//	      "required": [
		//	        "OrganizationSourceType"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "IncludeFutureRegions": {
		//	      "type": "boolean"
		//	    },
		//	    "SourceRegions": {
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "SourceType": {
		//	      "maxLength": 64,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "SourceType",
		//	    "SourceRegions"
		//	  ],
		//	  "type": "object"
		//	}
		"sync_source": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AwsOrganizationsSource
				"aws_organizations_source": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: OrganizationSourceType
						"organization_source_type": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: OrganizationalUnits
						"organizational_units": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: IncludeFutureRegions
				"include_future_regions": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: SourceRegions
				"source_regions": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SourceType
				"source_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: SyncType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"sync_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::SSM::ResourceDataSync",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SSM::ResourceDataSync").WithTerraformTypeName("awscc_ssm_resource_data_sync")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"aws_organizations_source": "AwsOrganizationsSource",
		"bucket_name":              "BucketName",
		"bucket_prefix":            "BucketPrefix",
		"bucket_region":            "BucketRegion",
		"include_future_regions":   "IncludeFutureRegions",
		"kms_key_arn":              "KMSKeyArn",
		"organization_source_type": "OrganizationSourceType",
		"organizational_units":     "OrganizationalUnits",
		"s3_destination":           "S3Destination",
		"source_regions":           "SourceRegions",
		"source_type":              "SourceType",
		"sync_format":              "SyncFormat",
		"sync_name":                "SyncName",
		"sync_source":              "SyncSource",
		"sync_type":                "SyncType",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
