// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ecr

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"regexp"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ecr_pull_through_cache_rule", pullThroughCacheRuleResource)
}

// pullThroughCacheRuleResource returns the Terraform awscc_ecr_pull_through_cache_rule resource.
// This Terraform resource corresponds to the CloudFormation AWS::ECR::PullThroughCacheRule resource.
func pullThroughCacheRuleResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CredentialArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the AWS Secrets Manager secret that identifies the credentials to authenticate to the upstream registry.",
		//	  "maxLength": 612,
		//	  "minLength": 50,
		//	  "pattern": "^arn:aws:secretsmanager:[a-zA-Z0-9-:]+:secret:ecr\\-pullthroughcache\\/[a-zA-Z0-9\\/_+=.@-]+$",
		//	  "type": "string"
		//	}
		"credential_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the AWS Secrets Manager secret that identifies the credentials to authenticate to the upstream registry.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(50, 612),
				stringvalidator.RegexMatches(regexp.MustCompile("^arn:aws:secretsmanager:[a-zA-Z0-9-:]+:secret:ecr\\-pullthroughcache\\/[a-zA-Z0-9\\/_+=.@-]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// CredentialArn is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: EcrRepositoryPrefix
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ECRRepositoryPrefix is a custom alias for upstream registry url.",
		//	  "maxLength": 30,
		//	  "minLength": 2,
		//	  "pattern": "(?:[a-z0-9]+(?:[._-][a-z0-9]+)*/)*[a-z0-9]+(?:[._-][a-z0-9]+)*",
		//	  "type": "string"
		//	}
		"ecr_repository_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ECRRepositoryPrefix is a custom alias for upstream registry url.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(2, 30),
				stringvalidator.RegexMatches(regexp.MustCompile("(?:[a-z0-9]+(?:[._-][a-z0-9]+)*/)*[a-z0-9]+(?:[._-][a-z0-9]+)*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: UpstreamRegistry
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the upstream registry.",
		//	  "type": "string"
		//	}
		"upstream_registry": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the upstream registry.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// UpstreamRegistry is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: UpstreamRegistryUrl
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The upstreamRegistryUrl is the endpoint of upstream registry url of the public repository to be cached",
		//	  "type": "string"
		//	}
		"upstream_registry_url": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The upstreamRegistryUrl is the endpoint of upstream registry url of the public repository to be cached",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "The AWS::ECR::PullThroughCacheRule resource configures the upstream registry configuration details for an Amazon Elastic Container Registry (Amazon Private ECR) pull-through cache.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ECR::PullThroughCacheRule").WithTerraformTypeName("awscc_ecr_pull_through_cache_rule")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"credential_arn":        "CredentialArn",
		"ecr_repository_prefix": "EcrRepositoryPrefix",
		"upstream_registry":     "UpstreamRegistry",
		"upstream_registry_url": "UpstreamRegistryUrl",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/CredentialArn",
		"/properties/UpstreamRegistry",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
