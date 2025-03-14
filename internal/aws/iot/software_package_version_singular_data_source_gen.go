// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iot

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iot_software_package_version", softwarePackageVersionDataSource)
}

// softwarePackageVersionDataSource returns the Terraform awscc_iot_software_package_version data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoT::SoftwarePackageVersion resource.
func softwarePackageVersionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Artifact
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The artifact location of the package version",
		//	  "properties": {
		//	    "S3Location": {
		//	      "additionalProperties": false,
		//	      "description": "The Amazon S3 location",
		//	      "properties": {
		//	        "Bucket": {
		//	          "description": "The S3 bucket",
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "Key": {
		//	          "description": "The S3 key",
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "Version": {
		//	          "description": "The S3 version",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "Bucket",
		//	        "Key",
		//	        "Version"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "S3Location"
		//	  ],
		//	  "type": "object"
		//	}
		"artifact": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: S3Location
				"s3_location": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Bucket
						"bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The S3 bucket",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: Key
						"key": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The S3 key",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: Version
						"version": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The S3 version",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The Amazon S3 location",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The artifact location of the package version",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Attributes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "patternProperties": {
		//	    "": {
		//	      "minLength": 1,
		//	      "pattern": "^[^\\p{C}]+$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"attributes":        // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 1024,
		//	  "minLength": 0,
		//	  "pattern": "^[^\\p{C}]+$",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ErrorReason
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"error_reason": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PackageName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9-_.]+$",
		//	  "type": "string"
		//	}
		"package_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PackageVersionArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "^arn:[!-~]+$",
		//	  "type": "string"
		//	}
		"package_version_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Recipe
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The inline json job document associated with a software package version",
		//	  "type": "string"
		//	}
		"recipe": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The inline json job document associated with a software package version",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Sbom
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The sbom zip archive location of the package version",
		//	  "properties": {
		//	    "S3Location": {
		//	      "additionalProperties": false,
		//	      "description": "The Amazon S3 location",
		//	      "properties": {
		//	        "Bucket": {
		//	          "description": "The S3 bucket",
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "Key": {
		//	          "description": "The S3 key",
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "Version": {
		//	          "description": "The S3 version",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "Bucket",
		//	        "Key",
		//	        "Version"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "S3Location"
		//	  ],
		//	  "type": "object"
		//	}
		"sbom": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: S3Location
				"s3_location": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Bucket
						"bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The S3 bucket",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: Key
						"key": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The S3 key",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: Version
						"version": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The S3 version",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The Amazon S3 location",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The sbom zip archive location of the package version",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SbomValidationStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The validation status of the Sbom file",
		//	  "enum": [
		//	    "IN_PROGRESS",
		//	    "FAILED",
		//	    "SUCCEEDED",
		//	    ""
		//	  ],
		//	  "type": "string"
		//	}
		"sbom_validation_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The validation status of the Sbom file",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "DRAFT",
		//	    "PUBLISHED",
		//	    "DEPRECATED"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VersionName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9-_.]+$",
		//	  "type": "string"
		//	}
		"version_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IoT::SoftwarePackageVersion",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::SoftwarePackageVersion").WithTerraformTypeName("awscc_iot_software_package_version")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"artifact":               "Artifact",
		"attributes":             "Attributes",
		"bucket":                 "Bucket",
		"description":            "Description",
		"error_reason":           "ErrorReason",
		"key":                    "Key",
		"package_name":           "PackageName",
		"package_version_arn":    "PackageVersionArn",
		"recipe":                 "Recipe",
		"s3_location":            "S3Location",
		"sbom":                   "Sbom",
		"sbom_validation_status": "SbomValidationStatus",
		"status":                 "Status",
		"tags":                   "Tags",
		"value":                  "Value",
		"version":                "Version",
		"version_name":           "VersionName",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
