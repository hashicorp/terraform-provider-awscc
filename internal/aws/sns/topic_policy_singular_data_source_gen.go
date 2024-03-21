// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package sns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddDataSourceFactory("awscc_sns_topic_policy", topicPolicyDataSource)
}

// topicPolicyDataSource returns the Terraform awscc_sns_topic_policy data source.
// This Terraform data source corresponds to the CloudFormation AWS::SNS::TopicPolicy resource.
func topicPolicyDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PolicyDocument
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A policy document that contains permissions to add to the specified SNS topics.",
		//	  "type": "string"
		//	}
		"policy_document": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A policy document that contains permissions to add to the specified SNS topics.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Topics
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Names (ARN) of the topics to which you want to add the policy. You can use the ``Ref`` function to specify an ``AWS::SNS::Topic`` resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"topics": schema.ListAttribute{ /*START ATTRIBUTE*/
			CustomType:  cctypes.NewMultisetTypeOf[types.String](ctx),
			Description: "The Amazon Resource Names (ARN) of the topics to which you want to add the policy. You can use the ``Ref`` function to specify an ``AWS::SNS::Topic`` resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::SNS::TopicPolicy",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SNS::TopicPolicy").WithTerraformTypeName("awscc_sns_topic_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"id":              "Id",
		"policy_document": "PolicyDocument",
		"topics":          "Topics",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
