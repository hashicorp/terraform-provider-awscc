// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_apigateway_resource", resourceDataSource)
}

// resourceDataSource returns the Terraform awscc_apigateway_resource data source.
// This Terraform data source corresponds to the CloudFormation AWS::ApiGateway::Resource resource.
func resourceDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"parent_id": {
			// Property: ParentId
			// CloudFormation resource type schema:
			// {
			//   "description": "The parent resource's identifier.",
			//   "type": "string"
			// }
			Description: "The parent resource's identifier.",
			Type:        types.StringType,
			Computed:    true,
		},
		"path_part": {
			// Property: PathPart
			// CloudFormation resource type schema:
			// {
			//   "description": "The last path segment for this resource.",
			//   "type": "string"
			// }
			Description: "The last path segment for this resource.",
			Type:        types.StringType,
			Computed:    true,
		},
		"resource_id": {
			// Property: ResourceId
			// CloudFormation resource type schema:
			// {
			//   "description": "A unique primary identifier for a Resource",
			//   "type": "string"
			// }
			Description: "A unique primary identifier for a Resource",
			Type:        types.StringType,
			Computed:    true,
		},
		"rest_api_id": {
			// Property: RestApiId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the RestApi resource in which you want to create this resource..",
			//   "type": "string"
			// }
			Description: "The ID of the RestApi resource in which you want to create this resource..",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::ApiGateway::Resource",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::Resource").WithTerraformTypeName("awscc_apigateway_resource")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"parent_id":   "ParentId",
		"path_part":   "PathPart",
		"resource_id": "ResourceId",
		"rest_api_id": "RestApiId",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
