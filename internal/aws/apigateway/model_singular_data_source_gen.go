// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_apigateway_model", modelDataSource)
}

// modelDataSource returns the Terraform awscc_apigateway_model data source.
// This Terraform data source corresponds to the CloudFormation AWS::ApiGateway::Model resource.
func modelDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ContentType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The content type for the model.",
		//	  "type": "string"
		//	}
		"content_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The content type for the model.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description that identifies this model.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description that identifies this model.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A name for the model. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the model name.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A name for the model. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the model name.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RestApiId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of a REST API with which to associate this model.",
		//	  "type": "string"
		//	}
		"rest_api_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of a REST API with which to associate this model.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Schema
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The schema to use to transform data to one or more output formats. Specify null ({}) if you don't want to specify a schema.",
		//	  "type": "string"
		//	}
		"schema": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The schema to use to transform data to one or more output formats. Specify null ({}) if you don't want to specify a schema.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ApiGateway::Model",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::Model").WithTerraformTypeName("awscc_apigateway_model")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"content_type": "ContentType",
		"description":  "Description",
		"name":         "Name",
		"rest_api_id":  "RestApiId",
		"schema":       "Schema",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}