// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package cloudformation

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_cloudformation_resource_version", resourceVersionResource)
}

// resourceVersionResource returns the Terraform awscc_cloudformation_resource_version resource.
// This Terraform resource corresponds to the CloudFormation AWS::CloudFormation::ResourceVersion resource.
func resourceVersionResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the type, here the ResourceVersion. This is used to uniquely identify a ResourceVersion resource",
		//	  "pattern": "^arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:([0-9]{12})?:type/resource/.+$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the type, here the ResourceVersion. This is used to uniquely identify a ResourceVersion resource",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ExecutionRoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the IAM execution role to use to register the type. If your resource type calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. CloudFormation then assumes that execution role to provide your resource type with the appropriate credentials.",
		//	  "type": "string"
		//	}
		"execution_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the IAM execution role to use to register the type. If your resource type calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. CloudFormation then assumes that execution role to provide your resource type with the appropriate credentials.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IsDefaultVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates if this type version is the current default version",
		//	  "type": "boolean"
		//	}
		"is_default_version": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates if this type version is the current default version",
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LoggingConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Specifies logging configuration information for a type.",
		//	  "properties": {
		//	    "LogGroupName": {
		//	      "description": "The Amazon CloudWatch log group to which CloudFormation sends error logging information when invoking the type's handlers.",
		//	      "maxLength": 512,
		//	      "minLength": 1,
		//	      "pattern": "^[\\.\\-_/#A-Za-z0-9]+$",
		//	      "type": "string"
		//	    },
		//	    "LogRoleArn": {
		//	      "description": "The ARN of the role that CloudFormation should assume when sending log entries to CloudWatch logs.",
		//	      "maxLength": 256,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"logging_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: LogGroupName
				"log_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The Amazon CloudWatch log group to which CloudFormation sends error logging information when invoking the type's handlers.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 512),
						stringvalidator.RegexMatches(regexp.MustCompile("^[\\.\\-_/#A-Za-z0-9]+$"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
						stringplanmodifier.RequiresReplaceIfConfigured(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: LogRoleArn
				"log_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the role that CloudFormation should assume when sending log entries to CloudWatch logs.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 256),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
						stringplanmodifier.RequiresReplaceIfConfigured(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Specifies logging configuration information for a type.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ProvisioningType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The provisioning behavior of the type. AWS CloudFormation determines the provisioning type during registration, based on the types of handlers in the schema handler package submitted.",
		//	  "enum": [
		//	    "NON_PROVISIONABLE",
		//	    "IMMUTABLE",
		//	    "FULLY_MUTABLE"
		//	  ],
		//	  "type": "string"
		//	}
		"provisioning_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The provisioning behavior of the type. AWS CloudFormation determines the provisioning type during registration, based on the types of handlers in the schema handler package submitted.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SchemaHandlerPackage
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A url to the S3 bucket containing the schema handler package that contains the schema, event handlers, and associated files for the type you want to register.\n\nFor information on generating a schema handler package for the type you want to register, see submit in the CloudFormation CLI User Guide.",
		//	  "type": "string"
		//	}
		"schema_handler_package": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A url to the S3 bucket containing the schema handler package that contains the schema, event handlers, and associated files for the type you want to register.\n\nFor information on generating a schema handler package for the type you want to register, see submit in the CloudFormation CLI User Guide.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// SchemaHandlerPackage is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: TypeArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the type without the versionID.",
		//	  "pattern": "^arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:([0-9]{12})?:type/resource/.+$",
		//	  "type": "string"
		//	}
		"type_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the type without the versionID.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TypeName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the type being registered.\n\nWe recommend that type names adhere to the following pattern: company_or_organization::service::type.",
		//	  "pattern": "^[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}$",
		//	  "type": "string"
		//	}
		"type_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the type being registered.\n\nWe recommend that type names adhere to the following pattern: company_or_organization::service::type.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VersionId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the version of the type represented by this resource instance.",
		//	  "pattern": "^[A-Za-z0-9-]{1,128}$",
		//	  "type": "string"
		//	}
		"version_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the version of the type represented by this resource instance.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Visibility
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The scope at which the type is visible and usable in CloudFormation operations.\n\nValid values include:\n\nPRIVATE: The type is only visible and usable within the account in which it is registered. Currently, AWS CloudFormation marks any types you register as PRIVATE.\n\nPUBLIC: The type is publically visible and usable within any Amazon account.",
		//	  "enum": [
		//	    "PUBLIC",
		//	    "PRIVATE"
		//	  ],
		//	  "type": "string"
		//	}
		"visibility": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The scope at which the type is visible and usable in CloudFormation operations.\n\nValid values include:\n\nPRIVATE: The type is only visible and usable within the account in which it is registered. Currently, AWS CloudFormation marks any types you register as PRIVATE.\n\nPUBLIC: The type is publically visible and usable within any Amazon account.",
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
		Description: "A resource that has been registered in the CloudFormation Registry.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFormation::ResourceVersion").WithTerraformTypeName("awscc_cloudformation_resource_version")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                    "Arn",
		"execution_role_arn":     "ExecutionRoleArn",
		"is_default_version":     "IsDefaultVersion",
		"log_group_name":         "LogGroupName",
		"log_role_arn":           "LogRoleArn",
		"logging_config":         "LoggingConfig",
		"provisioning_type":      "ProvisioningType",
		"schema_handler_package": "SchemaHandlerPackage",
		"type_arn":               "TypeArn",
		"type_name":              "TypeName",
		"version_id":             "VersionId",
		"visibility":             "Visibility",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/SchemaHandlerPackage",
	})
	opts = opts.WithCreateTimeoutInMinutes(2160).WithDeleteTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
