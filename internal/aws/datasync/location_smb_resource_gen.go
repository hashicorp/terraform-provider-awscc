// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package datasync

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/defaults"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_datasync_location_smb", locationSMBResource)
}

// locationSMBResource returns the Terraform awscc_datasync_location_smb resource.
// This Terraform resource corresponds to the CloudFormation AWS::DataSync::LocationSMB resource.
func locationSMBResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AgentArns
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Names (ARNs) of agents to use for a Simple Message Block (SMB) location.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 128,
		//	    "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\\-0-9]+:[0-9]{12}:agent/agent-[0-9a-z]{17}$",
		//	    "type": "string"
		//	  },
		//	  "maxItems": 4,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"agent_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The Amazon Resource Names (ARNs) of agents to use for a Simple Message Block (SMB) location.",
			Required:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(1, 4),
				listvalidator.ValueStringsAre(
					stringvalidator.LengthAtMost(128),
					stringvalidator.RegexMatches(regexp.MustCompile("^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\\-0-9]+:[0-9]{12}:agent/agent-[0-9a-z]{17}$"), ""),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Domain
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the Windows domain that the SMB server belongs to.",
		//	  "maxLength": 253,
		//	  "pattern": "^([A-Za-z0-9]+[A-Za-z0-9-.]*)*[A-Za-z0-9-]*[A-Za-z0-9]$",
		//	  "type": "string"
		//	}
		"domain": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the Windows domain that the SMB server belongs to.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(253),
				stringvalidator.RegexMatches(regexp.MustCompile("^([A-Za-z0-9]+[A-Za-z0-9-.]*)*[A-Za-z0-9-]*[A-Za-z0-9]$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LocationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the SMB location that is created.",
		//	  "maxLength": 128,
		//	  "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\\-0-9]+:[0-9]{12}:location/loc-[0-9a-z]{17}$",
		//	  "type": "string"
		//	}
		"location_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the SMB location that is created.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LocationUri
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The URL of the SMB location that was described.",
		//	  "maxLength": 4356,
		//	  "pattern": "^(efs|nfs|s3|smb|fsxw)://[a-zA-Z0-9./\\-]+$",
		//	  "type": "string"
		//	}
		"location_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The URL of the SMB location that was described.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MountOptions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "default": {
		//	    "Version": "AUTOMATIC"
		//	  },
		//	  "description": "The mount options used by DataSync to access the SMB server.",
		//	  "properties": {
		//	    "Version": {
		//	      "description": "The specific SMB version that you want DataSync to use to mount your SMB share.",
		//	      "enum": [
		//	        "AUTOMATIC",
		//	        "SMB1",
		//	        "SMB2_0",
		//	        "SMB2",
		//	        "SMB3"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"mount_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Version
				"version": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The specific SMB version that you want DataSync to use to mount your SMB share.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"AUTOMATIC",
							"SMB1",
							"SMB2_0",
							"SMB2",
							"SMB3",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The mount options used by DataSync to access the SMB server.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				defaults.StaticPartialObject(map[string]interface{}{
					"version": "AUTOMATIC",
				}),
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Password
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The password of the user who can mount the share and has the permissions to access files and folders in the SMB share.",
		//	  "maxLength": 104,
		//	  "pattern": "^.{0,104}$",
		//	  "type": "string"
		//	}
		"password": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The password of the user who can mount the share and has the permissions to access files and folders in the SMB share.",
			Optional:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(104),
				stringvalidator.RegexMatches(regexp.MustCompile("^.{0,104}$"), ""),
			}, /*END VALIDATORS*/
			// Password is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: ServerHostname
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the SMB server. This value is the IP address or Domain Name Service (DNS) name of the SMB server.",
		//	  "maxLength": 255,
		//	  "pattern": "^(([a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9\\-]*[A-Za-z0-9])$",
		//	  "type": "string"
		//	}
		"server_hostname": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the SMB server. This value is the IP address or Domain Name Service (DNS) name of the SMB server.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(255),
				stringvalidator.RegexMatches(regexp.MustCompile("^(([a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9\\-]*[A-Za-z0-9])$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
			// ServerHostname is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Subdirectory
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The subdirectory in the SMB file system that is used to read data from the SMB source location or write data to the SMB destination",
		//	  "maxLength": 4096,
		//	  "pattern": "^[a-zA-Z0-9_\\-\\+\\./\\(\\)\\$\\p{Zs}]+$",
		//	  "type": "string"
		//	}
		"subdirectory": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The subdirectory in the SMB file system that is used to read data from the SMB source location or write data to the SMB destination",
			Optional:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(4096),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_\\-\\+\\./\\(\\)\\$\\p{Zs}]+$"), ""),
			}, /*END VALIDATORS*/
			// Subdirectory is a write-only property.
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
		//	        "description": "The key for an AWS resource tag.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "pattern": "^[a-zA-Z0-9\\s+=._:/-]+$",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for an AWS resource tag.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "pattern": "^[a-zA-Z0-9\\s+=._:@/-]+$",
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
						Description: "The key for an AWS resource tag.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 256),
							stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9\\s+=._:/-]+$"), ""),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for an AWS resource tag.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 256),
							stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9\\s+=._:@/-]+$"), ""),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeAtMost(50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: User
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The user who can mount the share, has the permissions to access files and folders in the SMB share.",
		//	  "maxLength": 104,
		//	  "pattern": "^[^\\x5B\\x5D\\\\/:;|=,+*?]{1,104}$",
		//	  "type": "string"
		//	}
		"user": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The user who can mount the share, has the permissions to access files and folders in the SMB share.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(104),
				stringvalidator.RegexMatches(regexp.MustCompile("^[^\\x5B\\x5D\\\\/:;|=,+*?]{1,104}$"), ""),
			}, /*END VALIDATORS*/
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
		Description: "Resource schema for AWS::DataSync::LocationSMB.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataSync::LocationSMB").WithTerraformTypeName("awscc_datasync_location_smb")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"agent_arns":      "AgentArns",
		"domain":          "Domain",
		"key":             "Key",
		"location_arn":    "LocationArn",
		"location_uri":    "LocationUri",
		"mount_options":   "MountOptions",
		"password":        "Password",
		"server_hostname": "ServerHostname",
		"subdirectory":    "Subdirectory",
		"tags":            "Tags",
		"user":            "User",
		"value":           "Value",
		"version":         "Version",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Password",
		"/properties/Subdirectory",
		"/properties/ServerHostname",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
