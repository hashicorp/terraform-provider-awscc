// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_gamelift_build", buildDataSource)
}

// buildDataSource returns the Terraform awscc_gamelift_build data source.
// This Terraform data source corresponds to the CloudFormation AWS::GameLift::Build resource.
func buildDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: BuildArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift build resource and uniquely identifies it. ARNs are unique across all Regions. In a GameLift build ARN, the resource ID matches the BuildId value.",
		//	  "pattern": "^arn:.*:build\\/build-\\S+",
		//	  "type": "string"
		//	}
		"build_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift build resource and uniquely identifies it. ARNs are unique across all Regions. In a GameLift build ARN, the resource ID matches the BuildId value.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: BuildId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A unique identifier for a build to be deployed on the new fleet. If you are deploying the fleet with a custom game build, you must specify this property. The build must have been successfully uploaded to Amazon GameLift and be in a READY status. This fleet setting cannot be changed once the fleet is created.",
		//	  "type": "string"
		//	}
		"build_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A unique identifier for a build to be deployed on the new fleet. If you are deploying the fleet with a custom game build, you must specify this property. The build must have been successfully uploaded to Amazon GameLift and be in a READY status. This fleet setting cannot be changed once the fleet is created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A descriptive label that is associated with a build. Build names do not need to be unique.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A descriptive label that is associated with a build. Build names do not need to be unique.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: OperatingSystem
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The operating system that the game server binaries are built to run on. This value determines the type of fleet resources that you can use for this build. If your game build contains multiple executables, they all must run on the same operating system. If an operating system is not specified when creating a build, Amazon GameLift uses the default value (WINDOWS_2012). This value cannot be changed later.",
		//	  "enum": [
		//	    "AMAZON_LINUX",
		//	    "AMAZON_LINUX_2",
		//	    "AMAZON_LINUX_2023",
		//	    "WINDOWS_2012",
		//	    "WINDOWS_2016"
		//	  ],
		//	  "type": "string"
		//	}
		"operating_system": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The operating system that the game server binaries are built to run on. This value determines the type of fleet resources that you can use for this build. If your game build contains multiple executables, they all must run on the same operating system. If an operating system is not specified when creating a build, Amazon GameLift uses the default value (WINDOWS_2012). This value cannot be changed later.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ServerSdkVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A server SDK version you used when integrating your game server build with Amazon GameLift. By default Amazon GameLift sets this value to 4.0.2.",
		//	  "type": "string"
		//	}
		"server_sdk_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A server SDK version you used when integrating your game server build with Amazon GameLift. By default Amazon GameLift sets this value to 4.0.2.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: StorageLocation
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "$comment": "Contains object details present in the S3 Bucket",
		//	  "description": "Information indicating where your game build files are stored. Use this parameter only when creating a build with files stored in an Amazon S3 bucket that you own. The storage location must specify an Amazon S3 bucket name and key. The location must also specify a role ARN that you set up to allow Amazon GameLift to access your Amazon S3 bucket. The S3 bucket and your new build must be in the same Region.",
		//	  "properties": {
		//	    "Bucket": {
		//	      "description": "An Amazon S3 bucket identifier. This is the name of the S3 bucket.",
		//	      "type": "string"
		//	    },
		//	    "Key": {
		//	      "description": "The name of the zip file that contains the build files or script files.",
		//	      "type": "string"
		//	    },
		//	    "ObjectVersion": {
		//	      "description": "The version of the file, if object versioning is turned on for the bucket. Amazon GameLift uses this information when retrieving files from your S3 bucket. To retrieve a specific version of the file, provide an object version. To retrieve the latest version of the file, do not set this parameter.",
		//	      "type": "string"
		//	    },
		//	    "RoleArn": {
		//	      "description": "The Amazon Resource Name (ARN) for an IAM role that allows Amazon GameLift to access the S3 bucket.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Bucket",
		//	    "Key",
		//	    "RoleArn"
		//	  ],
		//	  "type": "object"
		//	}
		"storage_location": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Bucket
				"bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "An Amazon S3 bucket identifier. This is the name of the S3 bucket.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Key
				"key": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name of the zip file that contains the build files or script files.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ObjectVersion
				"object_version": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The version of the file, if object versioning is turned on for the bucket. Amazon GameLift uses this information when retrieving files from your S3 bucket. To retrieve a specific version of the file, provide an object version. To retrieve the latest version of the file, do not set this parameter.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: RoleArn
				"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The Amazon Resource Name (ARN) for an IAM role that allows Amazon GameLift to access the S3 bucket.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Information indicating where your game build files are stored. Use this parameter only when creating a build with files stored in an Amazon S3 bucket that you own. The storage location must specify an Amazon S3 bucket name and key. The location must also specify a role ARN that you set up to allow Amazon GameLift to access your Amazon S3 bucket. The S3 bucket and your new build must be in the same Region.",
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
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 200,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Version
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Version information that is associated with this build. Version strings do not need to be unique.",
		//	  "type": "string"
		//	}
		"version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Version information that is associated with this build. Version strings do not need to be unique.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::GameLift::Build",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::GameLift::Build").WithTerraformTypeName("awscc_gamelift_build")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"bucket":             "Bucket",
		"build_arn":          "BuildArn",
		"build_id":           "BuildId",
		"key":                "Key",
		"name":               "Name",
		"object_version":     "ObjectVersion",
		"operating_system":   "OperatingSystem",
		"role_arn":           "RoleArn",
		"server_sdk_version": "ServerSdkVersion",
		"storage_location":   "StorageLocation",
		"tags":               "Tags",
		"value":              "Value",
		"version":            "Version",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
