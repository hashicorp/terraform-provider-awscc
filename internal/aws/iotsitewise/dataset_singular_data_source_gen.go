// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iotsitewise

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iotsitewise_dataset", datasetDataSource)
}

// datasetDataSource returns the Terraform awscc_iotsitewise_dataset data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoTSiteWise::Dataset resource.
func datasetDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DatasetArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the dataset.",
		//	  "type": "string"
		//	}
		"dataset_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the dataset.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DatasetDescription
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description about the dataset, and its functionality.",
		//	  "type": "string"
		//	}
		"dataset_description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description about the dataset, and its functionality.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DatasetId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the dataset.",
		//	  "maxLength": 36,
		//	  "minLength": 36,
		//	  "pattern": "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$",
		//	  "type": "string"
		//	}
		"dataset_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the dataset.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DatasetName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the dataset.",
		//	  "type": "string"
		//	}
		"dataset_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the dataset.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DatasetSource
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The data source for the dataset.",
		//	  "properties": {
		//	    "SourceDetail": {
		//	      "additionalProperties": false,
		//	      "description": "The details of the dataset source associated with the dataset.",
		//	      "properties": {
		//	        "Kendra": {
		//	          "additionalProperties": false,
		//	          "description": "Contains details about the Kendra dataset source.",
		//	          "properties": {
		//	            "KnowledgeBaseArn": {
		//	              "description": "The knowledgeBaseArn details for the Kendra dataset source.",
		//	              "type": "string"
		//	            },
		//	            "RoleArn": {
		//	              "description": "The roleARN details for the Kendra dataset source.",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "KnowledgeBaseArn",
		//	            "RoleArn"
		//	          ],
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "SourceFormat": {
		//	      "description": "The format of the dataset source associated with the dataset.",
		//	      "enum": [
		//	        "KNOWLEDGE_BASE"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "SourceType": {
		//	      "description": "The type of data source for the dataset.",
		//	      "enum": [
		//	        "KENDRA"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "SourceFormat",
		//	    "SourceType"
		//	  ],
		//	  "type": "object"
		//	}
		"dataset_source": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: SourceDetail
				"source_detail": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Kendra
						"kendra": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: KnowledgeBaseArn
								"knowledge_base_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The knowledgeBaseArn details for the Kendra dataset source.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: RoleArn
								"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The roleARN details for the Kendra dataset source.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "Contains details about the Kendra dataset source.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The details of the dataset source associated with the dataset.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SourceFormat
				"source_format": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The format of the dataset source associated with the dataset.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SourceType
				"source_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The type of data source for the dataset.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The data source for the dataset.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
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
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IoTSiteWise::Dataset",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTSiteWise::Dataset").WithTerraformTypeName("awscc_iotsitewise_dataset")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"dataset_arn":         "DatasetArn",
		"dataset_description": "DatasetDescription",
		"dataset_id":          "DatasetId",
		"dataset_name":        "DatasetName",
		"dataset_source":      "DatasetSource",
		"kendra":              "Kendra",
		"key":                 "Key",
		"knowledge_base_arn":  "KnowledgeBaseArn",
		"role_arn":            "RoleArn",
		"source_detail":       "SourceDetail",
		"source_format":       "SourceFormat",
		"source_type":         "SourceType",
		"tags":                "Tags",
		"value":               "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
