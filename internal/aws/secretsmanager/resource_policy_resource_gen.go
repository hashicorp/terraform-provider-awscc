// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package secretsmanager

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_secretsmanager_resource_policy", resourcePolicyResource)
}

// resourcePolicyResource returns the Terraform awscc_secretsmanager_resource_policy resource.
// This Terraform resource corresponds to the CloudFormation AWS::SecretsManager::ResourcePolicy resource.
func resourcePolicyResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: BlockPublicPolicy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies whether to block resource-based policies that allow broad access to the secret.",
		//	  "type": "boolean"
		//	}
		"block_public_policy": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies whether to block resource-based policies that allow broad access to the secret.",
			Optional:    true,
			// BlockPublicPolicy is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Arn of the secret.",
		//	  "type": "string"
		//	}
		"resource_policy_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Arn of the secret.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ResourcePolicy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A JSON-formatted string for an AWS resource-based policy.",
		//	  "type": "string"
		//	}
		"resource_policy": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A JSON-formatted string for an AWS resource-based policy.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: SecretId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN or name of the secret to attach the resource-based policy.",
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"secret_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN or name of the secret to attach the resource-based policy.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 2048),
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
		Description: "Resource Type definition for AWS::SecretsManager::ResourcePolicy",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SecretsManager::ResourcePolicy").WithTerraformTypeName("awscc_secretsmanager_resource_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"block_public_policy": "BlockPublicPolicy",
		"resource_policy":     "ResourcePolicy",
		"resource_policy_id":  "Id",
		"secret_id":           "SecretId",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/BlockPublicPolicy",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
