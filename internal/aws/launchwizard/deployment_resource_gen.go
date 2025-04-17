// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package launchwizard

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_launchwizard_deployment", deploymentResource)
}

// deploymentResource returns the Terraform awscc_launchwizard_deployment resource.
// This Terraform resource corresponds to the CloudFormation AWS::LaunchWizard::Deployment resource.
func deploymentResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ARN of the LaunchWizard deployment",
		//	  "pattern": "^arn:aws(-cn|-us-gov)?:launchwizard:[a-z0-9-]+:[0-9]{12}:deployment/([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ARN of the LaunchWizard deployment",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Timestamp of LaunchWizard deployment creation",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "Timestamp of LaunchWizard deployment creation",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DeletedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Timestamp of LaunchWizard deployment deletion",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"deleted_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "Timestamp of LaunchWizard deployment deletion",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DeploymentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Deployment ID of the LaunchWizard deployment",
		//	  "maxLength": 128,
		//	  "minLength": 2,
		//	  "pattern": "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$",
		//	  "type": "string"
		//	}
		"deployment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Deployment ID of the LaunchWizard deployment",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DeploymentPatternName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Workload deployment pattern name",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^[A-Za-z0-9][a-zA-Z0-9-]*$",
		//	  "type": "string"
		//	}
		"deployment_pattern_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Workload deployment pattern name",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 256),
				stringvalidator.RegexMatches(regexp.MustCompile("^[A-Za-z0-9][a-zA-Z0-9-]*$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of LaunchWizard deployment",
		//	  "maxLength": 50,
		//	  "minLength": 1,
		//	  "pattern": "^[A-Za-z0-9_\\s\\.-]+$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of LaunchWizard deployment",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 50),
				stringvalidator.RegexMatches(regexp.MustCompile("^[A-Za-z0-9_\\s\\.-]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ResourceGroup
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Resource Group Name created for LaunchWizard deployment",
		//	  "type": "string"
		//	}
		"resource_group": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Resource Group Name created for LaunchWizard deployment",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Specifications
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "LaunchWizard deployment specifications",
		//	  "patternProperties": {
		//	    "": {
		//	      "maxLength": 1500,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"specifications":    // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "LaunchWizard deployment specifications",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// Specifications is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Status of LaunchWizard deployment",
		//	  "enum": [
		//	    "COMPLETED",
		//	    "CREATING",
		//	    "DELETE_IN_PROGRESS",
		//	    "DELETE_INITIATING",
		//	    "DELETE_FAILED",
		//	    "DELETED",
		//	    "FAILED",
		//	    "IN_PROGRESS",
		//	    "VALIDATING"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Status of LaunchWizard deployment",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Tags for LaunchWizard deployment",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "^[a-zA-Z+-=._:/]+$",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
							stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z+-=._:/]+$"), ""),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Tags for LaunchWizard deployment",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: WorkloadName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Workload Name for LaunchWizard deployment",
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "pattern": "^[A-Za-z][a-zA-Z0-9-_]*$",
		//	  "type": "string"
		//	}
		"workload_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Workload Name for LaunchWizard deployment",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 100),
				stringvalidator.RegexMatches(regexp.MustCompile("^[A-Za-z][a-zA-Z0-9-_]*$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
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
		Description: "Definition of AWS::LaunchWizard::Deployment Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::LaunchWizard::Deployment").WithTerraformTypeName("awscc_launchwizard_deployment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                     "Arn",
		"created_at":              "CreatedAt",
		"deleted_at":              "DeletedAt",
		"deployment_id":           "DeploymentId",
		"deployment_pattern_name": "DeploymentPatternName",
		"key":                     "Key",
		"name":                    "Name",
		"resource_group":          "ResourceGroup",
		"specifications":          "Specifications",
		"status":                  "Status",
		"tags":                    "Tags",
		"value":                   "Value",
		"workload_name":           "WorkloadName",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Specifications",
	})
	opts = opts.WithCreateTimeoutInMinutes(370).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
