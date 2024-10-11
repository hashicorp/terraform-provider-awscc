// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package supportapp

import (
	"context"
	"regexp"

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
	registry.AddResourceFactory("awscc_supportapp_slack_workspace_configuration", slackWorkspaceConfigurationResource)
}

// slackWorkspaceConfigurationResource returns the Terraform awscc_supportapp_slack_workspace_configuration resource.
// This Terraform resource corresponds to the CloudFormation AWS::SupportApp::SlackWorkspaceConfiguration resource.
func slackWorkspaceConfigurationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: TeamId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The team ID in Slack, which uniquely identifies a workspace.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^\\S+$",
		//	  "type": "string"
		//	}
		"team_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The team ID in Slack, which uniquely identifies a workspace.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 256),
				stringvalidator.RegexMatches(regexp.MustCompile("^\\S+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VersionId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An identifier used to update an existing Slack workspace configuration in AWS CloudFormation.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^[0-9]+$",
		//	  "type": "string"
		//	}
		"version_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An identifier used to update an existing Slack workspace configuration in AWS CloudFormation.",
			Optional:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 256),
				stringvalidator.RegexMatches(regexp.MustCompile("^[0-9]+$"), ""),
			}, /*END VALIDATORS*/
			// VersionId is a write-only property.
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
		Description: "An AWS Support App resource that creates, updates, lists, and deletes Slack workspace configurations.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SupportApp::SlackWorkspaceConfiguration").WithTerraformTypeName("awscc_supportapp_slack_workspace_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"team_id":    "TeamId",
		"version_id": "VersionId",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/VersionId",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
