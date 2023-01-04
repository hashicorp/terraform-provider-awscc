// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ecr

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ecr_public_repository", publicRepositoryDataSource)
}

// publicRepositoryDataSource returns the Terraform awscc_ecr_public_repository data source.
// This Terraform data source corresponds to the CloudFormation AWS::ECR::PublicRepository resource.
func publicRepositoryDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: RepositoryCatalogData
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The CatalogData property type specifies Catalog data for ECR Public Repository. For information about Catalog Data, see \u003clink\u003e",
		//	  "properties": {
		//	    "AboutText": {
		//	      "description": "Provide a detailed description of the repository. Identify what is included in the repository, any licensing details, or other relevant information.",
		//	      "maxLength": 10240,
		//	      "type": "string"
		//	    },
		//	    "Architectures": {
		//	      "description": "Select the system architectures that the images in your repository are compatible with.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "description": "The name of the architecture.",
		//	        "maxLength": 50,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "maxItems": 50,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "OperatingSystems": {
		//	      "description": "Select the operating systems that the images in your repository are compatible with.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "description": "The name of the operating system.",
		//	        "maxLength": 50,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "maxItems": 50,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "RepositoryDescription": {
		//	      "description": "The description of the public repository.",
		//	      "maxLength": 1024,
		//	      "type": "string"
		//	    },
		//	    "UsageText": {
		//	      "description": "Provide detailed information about how to use the images in the repository. This provides context, support information, and additional usage details for users of the repository.",
		//	      "maxLength": 10240,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"repository_catalog_data": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AboutText
				"about_text": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Provide a detailed description of the repository. Identify what is included in the repository, any licensing details, or other relevant information.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Architectures
				"architectures": schema.SetAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "Select the system architectures that the images in your repository are compatible with.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: OperatingSystems
				"operating_systems": schema.SetAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "Select the operating systems that the images in your repository are compatible with.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: RepositoryDescription
				"repository_description": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The description of the public repository.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: UsageText
				"usage_text": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Provide detailed information about how to use the images in the repository. This provides context, support information, and additional usage details for users of the repository.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The CatalogData property type specifies Catalog data for ECR Public Repository. For information about Catalog Data, see <link>",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RepositoryName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name to use for the repository. The repository name may be specified on its own (such as nginx-web-app) or it can be prepended with a namespace to group the repository into a category (such as project-a/nginx-web-app). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the repository name. For more information, see https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html.",
		//	  "maxLength": 256,
		//	  "minLength": 2,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"repository_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name to use for the repository. The repository name may be specified on its own (such as nginx-web-app) or it can be prepended with a namespace to group the repository into a category (such as project-a/nginx-web-app). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the repository name. For more information, see https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RepositoryPolicyText
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The JSON repository policy text to apply to the repository. For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/RepositoryPolicyExamples.html in the Amazon Elastic Container Registry User Guide. ",
		//	  "type": "string"
		//	}
		"repository_policy_text": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The JSON repository policy text to apply to the repository. For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/RepositoryPolicyExamples.html in the Amazon Elastic Container Registry User Guide. ",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 127,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 255,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
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
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ECR::PublicRepository",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ECR::PublicRepository").WithTerraformTypeName("awscc_ecr_public_repository")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"about_text":              "AboutText",
		"architectures":           "Architectures",
		"arn":                     "Arn",
		"key":                     "Key",
		"operating_systems":       "OperatingSystems",
		"repository_catalog_data": "RepositoryCatalogData",
		"repository_description":  "RepositoryDescription",
		"repository_name":         "RepositoryName",
		"repository_policy_text":  "RepositoryPolicyText",
		"tags":                    "Tags",
		"usage_text":              "UsageText",
		"value":                   "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}