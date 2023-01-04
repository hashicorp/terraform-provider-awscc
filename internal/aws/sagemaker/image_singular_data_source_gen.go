// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_sagemaker_image", imageDataSource)
}

// imageDataSource returns the Terraform awscc_sagemaker_image data source.
// This Terraform data source corresponds to the CloudFormation AWS::SageMaker::Image resource.
func imageDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ImageArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the image.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^arn:aws(-[\\w]+)*:sagemaker:[a-z0-9\\-]*:[0-9]{12}:image\\/[a-z0-9]([-.]?[a-z0-9])*$",
		//	  "type": "string"
		//	}
		"image_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the image.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ImageDescription
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description of the image.",
		//	  "maxLength": 512,
		//	  "minLength": 1,
		//	  "pattern": ".+",
		//	  "type": "string"
		//	}
		"image_description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description of the image.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ImageDisplayName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The display name of the image.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "^[A-Za-z0-9 -_]+$",
		//	  "type": "string"
		//	}
		"image_display_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The display name of the image.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ImageName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the image.",
		//	  "maxLength": 63,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9]([-.]?[a-zA-Z0-9])*$",
		//	  "type": "string"
		//	}
		"image_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the image.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ImageRoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on behalf of the customer.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^arn:aws(-[\\w]+)*:iam::[0-9]{12}:role/.*$",
		//	  "type": "string"
		//	}
		"image_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on behalf of the customer.",
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
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "minItems": 1,
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
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::SageMaker::Image",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::Image").WithTerraformTypeName("awscc_sagemaker_image")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"image_arn":          "ImageArn",
		"image_description":  "ImageDescription",
		"image_display_name": "ImageDisplayName",
		"image_name":         "ImageName",
		"image_role_arn":     "ImageRoleArn",
		"key":                "Key",
		"tags":               "Tags",
		"value":              "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}