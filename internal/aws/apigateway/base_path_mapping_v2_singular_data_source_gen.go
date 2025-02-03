// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_apigateway_base_path_mapping_v2", basePathMappingV2DataSource)
}

// basePathMappingV2DataSource returns the Terraform awscc_apigateway_base_path_mapping_v2 data source.
// This Terraform data source corresponds to the CloudFormation AWS::ApiGateway::BasePathMappingV2 resource.
func basePathMappingV2DataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: BasePath
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The base path name that callers of the API must provide in the URL after the domain name.",
		//	  "type": "string"
		//	}
		"base_path": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The base path name that callers of the API must provide in the URL after the domain name.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: BasePathMappingArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Amazon Resource Name (ARN) of the resource.",
		//	  "type": "string"
		//	}
		"base_path_mapping_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Amazon Resource Name (ARN) of the resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DomainNameArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Arn of an AWS::ApiGateway::DomainNameV2 resource.",
		//	  "type": "string"
		//	}
		"domain_name_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Arn of an AWS::ApiGateway::DomainNameV2 resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RestApiId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the API.",
		//	  "type": "string"
		//	}
		"rest_api_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the API.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Stage
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the API's stage.",
		//	  "type": "string"
		//	}
		"stage": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the API's stage.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ApiGateway::BasePathMappingV2",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::BasePathMappingV2").WithTerraformTypeName("awscc_apigateway_base_path_mapping_v2")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"base_path":             "BasePath",
		"base_path_mapping_arn": "BasePathMappingArn",
		"domain_name_arn":       "DomainNameArn",
		"rest_api_id":           "RestApiId",
		"stage":                 "Stage",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
