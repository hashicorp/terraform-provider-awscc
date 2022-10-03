// Code generated by generators/resource/main.go; DO NOT EDIT.

package ecr

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceFactory("awscc_ecr_public_repository", publicRepositoryResource)
}

// publicRepositoryResource returns the Terraform awscc_ecr_public_repository resource.
// This Terraform resource corresponds to the CloudFormation AWS::ECR::PublicRepository resource.
func publicRepositoryResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"repository_catalog_data": {
			// Property: RepositoryCatalogData
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The CatalogData property type specifies Catalog data for ECR Public Repository. For information about Catalog Data, see \u003clink\u003e",
			//   "properties": {
			//     "AboutText": {
			//       "description": "Provide a detailed description of the repository. Identify what is included in the repository, any licensing details, or other relevant information.",
			//       "maxLength": 10240,
			//       "type": "string"
			//     },
			//     "Architectures": {
			//       "description": "Select the system architectures that the images in your repository are compatible with.",
			//       "insertionOrder": false,
			//       "items": {
			//         "description": "The name of the architecture.",
			//         "maxLength": 50,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "maxItems": 50,
			//       "type": "array",
			//       "uniqueItems": true
			//     },
			//     "OperatingSystems": {
			//       "description": "Select the operating systems that the images in your repository are compatible with.",
			//       "insertionOrder": false,
			//       "items": {
			//         "description": "The name of the operating system.",
			//         "maxLength": 50,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "maxItems": 50,
			//       "type": "array",
			//       "uniqueItems": true
			//     },
			//     "RepositoryDescription": {
			//       "description": "The description of the public repository.",
			//       "maxLength": 1024,
			//       "type": "string"
			//     },
			//     "UsageText": {
			//       "description": "Provide detailed information about how to use the images in the repository. This provides context, support information, and additional usage details for users of the repository.",
			//       "maxLength": 10240,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The CatalogData property type specifies Catalog data for ECR Public Repository. For information about Catalog Data, see <link>",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"about_text": {
						// Property: AboutText
						Description: "Provide a detailed description of the repository. Identify what is included in the repository, any licensing details, or other relevant information.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(10240),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"architectures": {
						// Property: Architectures
						Description: "Select the system architectures that the images in your repository are compatible with.",
						Type:        types.SetType{ElemType: types.StringType},
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenAtMost(50),
							validate.ArrayForEach(validate.StringLenBetween(1, 50)),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"operating_systems": {
						// Property: OperatingSystems
						Description: "Select the operating systems that the images in your repository are compatible with.",
						Type:        types.SetType{ElemType: types.StringType},
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenAtMost(50),
							validate.ArrayForEach(validate.StringLenBetween(1, 50)),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"repository_description": {
						// Property: RepositoryDescription
						Description: "The description of the public repository.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(1024),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"usage_text": {
						// Property: UsageText
						Description: "Provide detailed information about how to use the images in the repository. This provides context, support information, and additional usage details for users of the repository.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(10240),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"repository_name": {
			// Property: RepositoryName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name to use for the repository. The repository name may be specified on its own (such as nginx-web-app) or it can be prepended with a namespace to group the repository into a category (such as project-a/nginx-web-app). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the repository name. For more information, see https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html.",
			//   "maxLength": 256,
			//   "minLength": 2,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name to use for the repository. The repository name may be specified on its own (such as nginx-web-app) or it can be prepended with a namespace to group the repository into a category (such as project-a/nginx-web-app). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the repository name. For more information, see https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(2, 256),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"repository_policy_text": {
			// Property: RepositoryPolicyText
			// CloudFormation resource type schema:
			// {
			//   "description": "The JSON repository policy text to apply to the repository. For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/RepositoryPolicyExamples.html in the Amazon Elastic Container Registry User Guide. ",
			//   "type": "string"
			// }
			Description: "The JSON repository policy text to apply to the repository. For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/RepositoryPolicyExamples.html in the Amazon Elastic Container Registry User Guide. ",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 127,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 127),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 255),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(50),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			resource.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "The AWS::ECR::PublicRepository resource specifies an Amazon Elastic Container Public Registry (Amazon Public ECR) repository, where users can push and pull Docker images. For more information, see https://docs.aws.amazon.com/AmazonECR",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ECR::PublicRepository").WithTerraformTypeName("awscc_ecr_public_repository")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
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

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
