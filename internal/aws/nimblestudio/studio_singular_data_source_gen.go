// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package nimblestudio

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_nimblestudio_studio", studioDataSource)
}

// studioDataSource returns the Terraform awscc_nimblestudio_studio data source.
// This Terraform data source corresponds to the CloudFormation AWS::NimbleStudio::Studio resource.
func studioDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"admin_role_arn": {
			// Property: AdminRoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eThe IAM role that Studio Admins will assume when logging in to the Nimble Studio portal.\u003c/p\u003e",
			//   "type": "string"
			// }
			Description: "<p>The IAM role that Studio Admins will assume when logging in to the Nimble Studio portal.</p>",
			Type:        types.StringType,
			Computed:    true,
		},
		"display_name": {
			// Property: DisplayName
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eA friendly name for the studio.\u003c/p\u003e",
			//   "maxLength": 64,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "<p>A friendly name for the studio.</p>",
			Type:        types.StringType,
			Computed:    true,
		},
		"home_region": {
			// Property: HomeRegion
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eThe Amazon Web Services Region where the studio resource is located.\u003c/p\u003e",
			//   "maxLength": 50,
			//   "minLength": 0,
			//   "pattern": "[a-z]{2}-?(iso|gov)?-{1}[a-z]*-{1}[0-9]",
			//   "type": "string"
			// }
			Description: "<p>The Amazon Web Services Region where the studio resource is located.</p>",
			Type:        types.StringType,
			Computed:    true,
		},
		"sso_client_id": {
			// Property: SsoClientId
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eThe Amazon Web Services SSO application client ID used to integrate with Amazon Web Services SSO to enable Amazon Web Services SSO users to log in to Nimble Studio portal.\u003c/p\u003e",
			//   "type": "string"
			// }
			Description: "<p>The Amazon Web Services SSO application client ID used to integrate with Amazon Web Services SSO to enable Amazon Web Services SSO users to log in to Nimble Studio portal.</p>",
			Type:        types.StringType,
			Computed:    true,
		},
		"studio_encryption_configuration": {
			// Property: StudioEncryptionConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "\u003cp\u003eConfiguration of the encryption method that is used for the studio.\u003c/p\u003e",
			//   "properties": {
			//     "KeyArn": {
			//       "description": "\u003cp\u003eThe ARN for a KMS key that is used to encrypt studio data.\u003c/p\u003e",
			//       "minLength": 4,
			//       "pattern": "^arn:.*",
			//       "type": "string"
			//     },
			//     "KeyType": {
			//       "description": "\u003cp\u003eThe type of KMS key that is used to encrypt studio data.\u003c/p\u003e",
			//       "enum": [
			//         "AWS_OWNED_KEY",
			//         "CUSTOMER_MANAGED_KEY"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "KeyType"
			//   ],
			//   "type": "object"
			// }
			Description: "<p>Configuration of the encryption method that is used for the studio.</p>",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"key_arn": {
						// Property: KeyArn
						Description: "<p>The ARN for a KMS key that is used to encrypt studio data.</p>",
						Type:        types.StringType,
						Computed:    true,
					},
					"key_type": {
						// Property: KeyType
						Description: "<p>The type of KMS key that is used to encrypt studio data.</p>",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"studio_id": {
			// Property: StudioId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"studio_name": {
			// Property: StudioName
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eThe studio name that is used in the URL of the Nimble Studio portal when accessed by Nimble Studio users.\u003c/p\u003e",
			//   "maxLength": 64,
			//   "minLength": 3,
			//   "pattern": "^[a-z0-9]*$",
			//   "type": "string"
			// }
			Description: "<p>The studio name that is used in the URL of the Nimble Studio portal when accessed by Nimble Studio users.</p>",
			Type:        types.StringType,
			Computed:    true,
		},
		"studio_url": {
			// Property: StudioUrl
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eThe address of the web page for the studio.\u003c/p\u003e",
			//   "type": "string"
			// }
			Description: "<p>The address of the web page for the studio.</p>",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "",
			//   "patternProperties": {
			//     "": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Computed: true,
		},
		"user_role_arn": {
			// Property: UserRoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eThe IAM role that Studio Users will assume when logging in to the Nimble Studio portal.\u003c/p\u003e",
			//   "type": "string"
			// }
			Description: "<p>The IAM role that Studio Users will assume when logging in to the Nimble Studio portal.</p>",
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
		Description: "Data Source schema for AWS::NimbleStudio::Studio",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::NimbleStudio::Studio").WithTerraformTypeName("awscc_nimblestudio_studio")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"admin_role_arn":                  "AdminRoleArn",
		"display_name":                    "DisplayName",
		"home_region":                     "HomeRegion",
		"key_arn":                         "KeyArn",
		"key_type":                        "KeyType",
		"sso_client_id":                   "SsoClientId",
		"studio_encryption_configuration": "StudioEncryptionConfiguration",
		"studio_id":                       "StudioId",
		"studio_name":                     "StudioName",
		"studio_url":                      "StudioUrl",
		"tags":                            "Tags",
		"user_role_arn":                   "UserRoleArn",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
