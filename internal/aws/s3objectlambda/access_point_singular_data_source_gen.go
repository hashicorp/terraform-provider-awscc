// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package s3objectlambda

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_s3objectlambda_access_point", accessPointDataSource)
}

// accessPointDataSource returns the Terraform awscc_s3objectlambda_access_point data source.
// This Terraform data source corresponds to the CloudFormation AWS::S3ObjectLambda::AccessPoint resource.
func accessPointDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "pattern": "arn:[^:]+:s3-object-lambda:[^:]*:\\d{12}:accesspoint/.*",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"creation_date": {
			// Property: CreationDate
			// CloudFormation resource type schema:
			// {
			//   "description": "The date and time when the Object lambda Access Point was created.",
			//   "type": "string"
			// }
			Description: "The date and time when the Object lambda Access Point was created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name you want to assign to this Object lambda Access Point.",
			//   "maxLength": 45,
			//   "minLength": 3,
			//   "pattern": "^[a-z0-9]([a-z0-9\\-]*[a-z0-9])?$",
			//   "type": "string"
			// }
			Description: "The name you want to assign to this Object lambda Access Point.",
			Type:        types.StringType,
			Computed:    true,
		},
		"object_lambda_configuration": {
			// Property: ObjectLambdaConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The Object lambda Access Point Configuration that configures transformations to be applied on the objects on specified S3 Actions",
			//   "properties": {
			//     "AllowedFeatures": {
			//       "insertionOrder": false,
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array",
			//       "uniqueItems": true
			//     },
			//     "CloudWatchMetricsEnabled": {
			//       "type": "boolean"
			//     },
			//     "SupportingAccessPoint": {
			//       "maxLength": 2048,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "TransformationConfigurations": {
			//       "insertionOrder": false,
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "Configuration to define what content transformation will be applied on which S3 Action.",
			//         "properties": {
			//           "Actions": {
			//             "insertionOrder": false,
			//             "items": {
			//               "type": "string"
			//             },
			//             "type": "array",
			//             "uniqueItems": true
			//           },
			//           "ContentTransformation": {
			//             "properties": {
			//               "AwsLambda": {
			//                 "additionalProperties": false,
			//                 "properties": {
			//                   "FunctionArn": {
			//                     "maxLength": 2048,
			//                     "minLength": 1,
			//                     "type": "string"
			//                   },
			//                   "FunctionPayload": {
			//                     "type": "string"
			//                   }
			//                 },
			//                 "required": [
			//                   "FunctionArn"
			//                 ],
			//                 "type": "object"
			//               }
			//             },
			//             "type": "object"
			//           }
			//         },
			//         "required": [
			//           "Actions",
			//           "ContentTransformation"
			//         ],
			//         "type": "object"
			//       },
			//       "type": "array",
			//       "uniqueItems": true
			//     }
			//   },
			//   "required": [
			//     "SupportingAccessPoint",
			//     "TransformationConfigurations"
			//   ],
			//   "type": "object"
			// }
			Description: "The Object lambda Access Point Configuration that configures transformations to be applied on the objects on specified S3 Actions",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"allowed_features": {
						// Property: AllowedFeatures
						Type:     types.SetType{ElemType: types.StringType},
						Computed: true,
					},
					"cloudwatch_metrics_enabled": {
						// Property: CloudWatchMetricsEnabled
						Type:     types.BoolType,
						Computed: true,
					},
					"supporting_access_point": {
						// Property: SupportingAccessPoint
						Type:     types.StringType,
						Computed: true,
					},
					"transformation_configurations": {
						// Property: TransformationConfigurations
						Attributes: tfsdk.SetNestedAttributes(
							map[string]tfsdk.Attribute{
								"actions": {
									// Property: Actions
									Type:     types.SetType{ElemType: types.StringType},
									Computed: true,
								},
								"content_transformation": {
									// Property: ContentTransformation
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"aws_lambda": {
												// Property: AwsLambda
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"function_arn": {
															// Property: FunctionArn
															Type:     types.StringType,
															Computed: true,
														},
														"function_payload": {
															// Property: FunctionPayload
															Type:     types.StringType,
															Computed: true,
														},
													},
												),
												Computed: true,
											},
										},
									),
									Computed: true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"policy_status": {
			// Property: PolicyStatus
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "IsPublic": {
			//       "description": "Specifies whether the Object lambda Access Point Policy is Public or not. Object lambda Access Points are private by default.",
			//       "type": "boolean"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"is_public": {
						// Property: IsPublic
						Description: "Specifies whether the Object lambda Access Point Policy is Public or not. Object lambda Access Points are private by default.",
						Type:        types.BoolType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"public_access_block_configuration": {
			// Property: PublicAccessBlockConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The PublicAccessBlock configuration that you want to apply to this Access Point. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status 'The Meaning of Public' in the Amazon Simple Storage Service Developer Guide.",
			//   "properties": {
			//     "BlockPublicAcls": {
			//       "description": "Specifies whether Amazon S3 should block public access control lists (ACLs) to this object lambda access point. Setting this element to TRUE causes the following behavior:\n- PUT Bucket acl and PUT Object acl calls fail if the specified ACL is public.\n - PUT Object calls fail if the request includes a public ACL.\n. - PUT Bucket calls fail if the request includes a public ACL.\nEnabling this setting doesn't affect existing policies or ACLs.",
			//       "type": "boolean"
			//     },
			//     "BlockPublicPolicy": {
			//       "description": "Specifies whether Amazon S3 should block public bucket policies for buckets in this account. Setting this element to TRUE causes Amazon S3 to reject calls to PUT Bucket policy if the specified bucket policy allows public access. Enabling this setting doesn't affect existing bucket policies.",
			//       "type": "boolean"
			//     },
			//     "IgnorePublicAcls": {
			//       "description": "Specifies whether Amazon S3 should ignore public ACLs for buckets in this account. Setting this element to TRUE causes Amazon S3 to ignore all public ACLs on buckets in this account and any objects that they contain. Enabling this setting doesn't affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set.",
			//       "type": "boolean"
			//     },
			//     "RestrictPublicBuckets": {
			//       "description": "Specifies whether Amazon S3 should restrict public bucket policies for this bucket. Setting this element to TRUE restricts access to this bucket to only AWS services and authorized users within this account if the bucket has a public policy.\nEnabling this setting doesn't affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked.",
			//       "type": "boolean"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The PublicAccessBlock configuration that you want to apply to this Access Point. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status 'The Meaning of Public' in the Amazon Simple Storage Service Developer Guide.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"block_public_acls": {
						// Property: BlockPublicAcls
						Description: "Specifies whether Amazon S3 should block public access control lists (ACLs) to this object lambda access point. Setting this element to TRUE causes the following behavior:\n- PUT Bucket acl and PUT Object acl calls fail if the specified ACL is public.\n - PUT Object calls fail if the request includes a public ACL.\n. - PUT Bucket calls fail if the request includes a public ACL.\nEnabling this setting doesn't affect existing policies or ACLs.",
						Type:        types.BoolType,
						Computed:    true,
					},
					"block_public_policy": {
						// Property: BlockPublicPolicy
						Description: "Specifies whether Amazon S3 should block public bucket policies for buckets in this account. Setting this element to TRUE causes Amazon S3 to reject calls to PUT Bucket policy if the specified bucket policy allows public access. Enabling this setting doesn't affect existing bucket policies.",
						Type:        types.BoolType,
						Computed:    true,
					},
					"ignore_public_acls": {
						// Property: IgnorePublicAcls
						Description: "Specifies whether Amazon S3 should ignore public ACLs for buckets in this account. Setting this element to TRUE causes Amazon S3 to ignore all public ACLs on buckets in this account and any objects that they contain. Enabling this setting doesn't affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set.",
						Type:        types.BoolType,
						Computed:    true,
					},
					"restrict_public_buckets": {
						// Property: RestrictPublicBuckets
						Description: "Specifies whether Amazon S3 should restrict public bucket policies for this bucket. Setting this element to TRUE restricts access to this bucket to only AWS services and authorized users within this account if the bucket has a public policy.\nEnabling this setting doesn't affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked.",
						Type:        types.BoolType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::S3ObjectLambda::AccessPoint",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::S3ObjectLambda::AccessPoint").WithTerraformTypeName("awscc_s3objectlambda_access_point")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"actions":                           "Actions",
		"allowed_features":                  "AllowedFeatures",
		"arn":                               "Arn",
		"aws_lambda":                        "AwsLambda",
		"block_public_acls":                 "BlockPublicAcls",
		"block_public_policy":               "BlockPublicPolicy",
		"cloudwatch_metrics_enabled":        "CloudWatchMetricsEnabled",
		"content_transformation":            "ContentTransformation",
		"creation_date":                     "CreationDate",
		"function_arn":                      "FunctionArn",
		"function_payload":                  "FunctionPayload",
		"ignore_public_acls":                "IgnorePublicAcls",
		"is_public":                         "IsPublic",
		"name":                              "Name",
		"object_lambda_configuration":       "ObjectLambdaConfiguration",
		"policy_status":                     "PolicyStatus",
		"public_access_block_configuration": "PublicAccessBlockConfiguration",
		"restrict_public_buckets":           "RestrictPublicBuckets",
		"supporting_access_point":           "SupportingAccessPoint",
		"transformation_configurations":     "TransformationConfigurations",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
