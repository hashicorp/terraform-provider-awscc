// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package datazone

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_datazone_environment", environmentResource)
}

// environmentResource returns the Terraform awscc_datazone_environment resource.
// This Terraform resource corresponds to the CloudFormation AWS::DataZone::Environment resource.
func environmentResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AwsAccountId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AWS account in which the Amazon DataZone environment is created.",
		//	  "pattern": "^\\d{12}$",
		//	  "type": "string"
		//	}
		"aws_account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AWS account in which the Amazon DataZone environment is created.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AwsAccountRegion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AWS region in which the Amazon DataZone environment is created.",
		//	  "pattern": "^[a-z]{2}-[a-z]{4,10}-\\d$",
		//	  "type": "string"
		//	}
		"aws_account_region": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AWS region in which the Amazon DataZone environment is created.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The timestamp of when the environment was created.",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "The timestamp of when the environment was created.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CreatedBy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon DataZone user who created the environment.",
		//	  "type": "string"
		//	}
		"created_by": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon DataZone user who created the environment.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the Amazon DataZone environment.",
		//	  "maxLength": 2048,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the Amazon DataZone environment.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(2048),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DomainId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the Amazon DataZone domain in which the environment is created.",
		//	  "pattern": "^dzd[-_][a-zA-Z0-9_-]{1,36}$",
		//	  "type": "string"
		//	}
		"domain_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the Amazon DataZone domain in which the environment is created.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DomainIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the Amazon DataZone domain in which the environment would be created.",
		//	  "pattern": "^dzd[-_][a-zA-Z0-9_-]{1,36}$",
		//	  "type": "string"
		//	}
		"domain_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the Amazon DataZone domain in which the environment would be created.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^dzd[-_][a-zA-Z0-9_-]{1,36}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// DomainIdentifier is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: EnvironmentAccountIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AWS account in which the Amazon DataZone environment is created.",
		//	  "pattern": "^\\d{12}$",
		//	  "type": "string"
		//	}
		"environment_account_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AWS account in which the Amazon DataZone environment is created.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^\\d{12}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
			// EnvironmentAccountIdentifier is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: EnvironmentAccountRegion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AWS region in which the Amazon DataZone environment is created.",
		//	  "pattern": "^[a-z]{2}-[a-z]{4,10}-\\d$",
		//	  "type": "string"
		//	}
		"environment_account_region": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AWS region in which the Amazon DataZone environment is created.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-z]{2}-[a-z]{4,10}-\\d$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
			// EnvironmentAccountRegion is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: EnvironmentBlueprintId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the blueprint with which the Amazon DataZone environment was created.",
		//	  "pattern": "^[a-zA-Z0-9_-]{1,36}$",
		//	  "type": "string"
		//	}
		"environment_blueprint_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the blueprint with which the Amazon DataZone environment was created.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EnvironmentProfileId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the environment profile with which the Amazon DataZone environment was created.",
		//	  "pattern": "^[a-zA-Z0-9_-]{0,36}$",
		//	  "type": "string"
		//	}
		"environment_profile_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the environment profile with which the Amazon DataZone environment was created.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EnvironmentProfileIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the environment profile with which the Amazon DataZone environment would be created.",
		//	  "pattern": "^[a-zA-Z0-9_-]{0,36}$",
		//	  "type": "string"
		//	}
		"environment_profile_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the environment profile with which the Amazon DataZone environment would be created.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]{0,36}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
			// EnvironmentProfileIdentifier is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: EnvironmentRoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Environment role arn for custom aws environment permissions",
		//	  "type": "string"
		//	}
		"environment_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Environment role arn for custom aws environment permissions",
			Optional:    true,
			// EnvironmentRoleArn is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: GlossaryTerms
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The glossary terms that can be used in the Amazon DataZone environment.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "pattern": "^[a-zA-Z0-9_-]{1,36}$",
		//	    "type": "string"
		//	  },
		//	  "maxItems": 20,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"glossary_terms": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The glossary terms that can be used in the Amazon DataZone environment.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(1, 20),
				listvalidator.ValueStringsAre(
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]{1,36}$"), ""),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the Amazon DataZone environment.",
		//	  "pattern": "^[a-zA-Z0-9_-]{1,36}$",
		//	  "type": "string"
		//	}
		"environment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the Amazon DataZone environment.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the environment.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[\\w -]+$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the environment.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[\\w -]+$"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: ProjectId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the Amazon DataZone project in which the environment is created.",
		//	  "pattern": "^[a-zA-Z0-9_-]{1,36}$",
		//	  "type": "string"
		//	}
		"project_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the Amazon DataZone project in which the environment is created.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ProjectIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the Amazon DataZone project in which the environment would be created.",
		//	  "pattern": "^[a-zA-Z0-9_-]{1,36}$",
		//	  "type": "string"
		//	}
		"project_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the Amazon DataZone project in which the environment would be created.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]{1,36}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// ProjectIdentifier is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Provider
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The provider of the Amazon DataZone environment.",
		//	  "type": "string"
		//	}
		"provider_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The provider of the Amazon DataZone environment.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The status of the Amazon DataZone environment.",
		//	  "enum": [
		//	    "ACTIVE",
		//	    "CREATING",
		//	    "UPDATING",
		//	    "DELETING",
		//	    "CREATE_FAILED",
		//	    "UPDATE_FAILED",
		//	    "DELETE_FAILED",
		//	    "VALIDATION_FAILED",
		//	    "SUSPENDED",
		//	    "DISABLED",
		//	    "EXPIRED",
		//	    "DELETED",
		//	    "INACCESSIBLE"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The status of the Amazon DataZone environment.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: UpdatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The timestamp of when the environment was updated.",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"updated_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "The timestamp of when the environment was updated.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: UserParameters
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The user parameters of the Amazon DataZone environment.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The parameter details of an environment.",
		//	    "properties": {
		//	      "Name": {
		//	        "description": "The name of an environment parameter.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value of an environment parameter.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"user_parameters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Name
					"name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The name of an environment parameter.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
							stringplanmodifier.RequiresReplaceIfConfigured(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value of an environment parameter.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
							stringplanmodifier.RequiresReplaceIfConfigured(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The user parameters of the Amazon DataZone environment.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
				listplanmodifier.RequiresReplaceIfConfigured(),
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
		Description: "Definition of AWS::DataZone::Environment Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataZone::Environment").WithTerraformTypeName("awscc_datazone_environment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"aws_account_id":                 "AwsAccountId",
		"aws_account_region":             "AwsAccountRegion",
		"created_at":                     "CreatedAt",
		"created_by":                     "CreatedBy",
		"description":                    "Description",
		"domain_id":                      "DomainId",
		"domain_identifier":              "DomainIdentifier",
		"environment_account_identifier": "EnvironmentAccountIdentifier",
		"environment_account_region":     "EnvironmentAccountRegion",
		"environment_blueprint_id":       "EnvironmentBlueprintId",
		"environment_id":                 "Id",
		"environment_profile_id":         "EnvironmentProfileId",
		"environment_profile_identifier": "EnvironmentProfileIdentifier",
		"environment_role_arn":           "EnvironmentRoleArn",
		"glossary_terms":                 "GlossaryTerms",
		"name":                           "Name",
		"project_id":                     "ProjectId",
		"project_identifier":             "ProjectIdentifier",
		"provider_name":                  "Provider",
		"status":                         "Status",
		"updated_at":                     "UpdatedAt",
		"user_parameters":                "UserParameters",
		"value":                          "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/EnvironmentProfileIdentifier",
		"/properties/ProjectIdentifier",
		"/properties/DomainIdentifier",
		"/properties/EnvironmentAccountIdentifier",
		"/properties/EnvironmentAccountRegion",
		"/properties/EnvironmentRoleArn",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
