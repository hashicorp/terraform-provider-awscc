// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package codeartifact

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_codeartifact_repository", repositoryDataSource)
}

// repositoryDataSource returns the Terraform awscc_codeartifact_repository data source.
// This Terraform data source corresponds to the CloudFormation AWS::CodeArtifact::Repository resource.
func repositoryDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the repository.",
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the repository.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A text description of the repository.",
		//	  "maxLength": 1000,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A text description of the repository.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DomainName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the domain that contains the repository.",
		//	  "maxLength": 50,
		//	  "minLength": 2,
		//	  "pattern": "^([a-z][a-z0-9\\-]{0,48}[a-z0-9])$",
		//	  "type": "string"
		//	}
		"domain_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the domain that contains the repository.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DomainOwner
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The 12-digit account ID of the AWS account that owns the domain.",
		//	  "pattern": "[0-9]{12}",
		//	  "type": "string"
		//	}
		"domain_owner": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The 12-digit account ID of the AWS account that owns the domain.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ExternalConnections
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of external connections associated with the repository.",
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"external_connections": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of external connections associated with the repository.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the repository. This is used for GetAtt",
		//	  "maxLength": 100,
		//	  "minLength": 2,
		//	  "pattern": "^([A-Za-z0-9][A-Za-z0-9._\\-]{1,99})$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the repository. This is used for GetAtt",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PermissionsPolicyDocument
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The access control resource policy on the provided repository.",
		//	  "maxLength": 5120,
		//	  "minLength": 2,
		//	  "type": "object"
		//	}
		"permissions_policy_document": schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The access control resource policy on the provided repository.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RepositoryName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the repository.",
		//	  "maxLength": 100,
		//	  "minLength": 2,
		//	  "pattern": "^([A-Za-z0-9][A-Za-z0-9._\\-]{1,99})$",
		//	  "type": "string"
		//	}
		"repository_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the repository.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
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
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Upstreams
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of upstream repositories associated with the repository.",
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"upstreams": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of upstream repositories associated with the repository.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::CodeArtifact::Repository",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CodeArtifact::Repository").WithTerraformTypeName("awscc_codeartifact_repository")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                         "Arn",
		"description":                 "Description",
		"domain_name":                 "DomainName",
		"domain_owner":                "DomainOwner",
		"external_connections":        "ExternalConnections",
		"key":                         "Key",
		"name":                        "Name",
		"permissions_policy_document": "PermissionsPolicyDocument",
		"repository_name":             "RepositoryName",
		"tags":                        "Tags",
		"upstreams":                   "Upstreams",
		"value":                       "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}