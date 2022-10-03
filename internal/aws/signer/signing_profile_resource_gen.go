// Code generated by generators/resource/main.go; DO NOT EDIT.

package signer

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
	registry.AddResourceFactory("awscc_signer_signing_profile", signingProfileResource)
}

// signingProfileResource returns the Terraform awscc_signer_signing_profile resource.
// This Terraform resource corresponds to the CloudFormation AWS::Signer::SigningProfile resource.
func signingProfileResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the specified signing profile.",
			//   "pattern": "^arn:aws(-(cn|gov))?:[a-z-]+:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the specified signing profile.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"platform_id": {
			// Property: PlatformId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the target signing platform.",
			//   "enum": [
			//     "AWSLambda-SHA384-ECDSA"
			//   ],
			//   "type": "string"
			// }
			Description: "The ID of the target signing platform.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"AWSLambda-SHA384-ECDSA",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"profile_name": {
			// Property: ProfileName
			// CloudFormation resource type schema:
			// {
			//   "description": "A name for the signing profile. AWS CloudFormation generates a unique physical ID and uses that ID for the signing profile name. ",
			//   "type": "string"
			// }
			Description: "A name for the signing profile. AWS CloudFormation generates a unique physical ID and uses that ID for the signing profile name. ",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"profile_version": {
			// Property: ProfileVersion
			// CloudFormation resource type schema:
			// {
			//   "description": "A version for the signing profile. AWS Signer generates a unique version for each profile of the same profile name.",
			//   "pattern": "^[0-9a-zA-Z]{10}$",
			//   "type": "string"
			// }
			Description: "A version for the signing profile. AWS Signer generates a unique version for each profile of the same profile name.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"profile_version_arn": {
			// Property: ProfileVersionArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the specified signing profile version.",
			//   "pattern": "^arn:aws(-(cn|gov))?:[a-z-]+:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the specified signing profile version.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"signature_validity_period": {
			// Property: SignatureValidityPeriod
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Signature validity period of the profile.",
			//   "properties": {
			//     "Type": {
			//       "enum": [
			//         "DAYS",
			//         "MONTHS",
			//         "YEARS"
			//       ],
			//       "type": "string"
			//     },
			//     "Value": {
			//       "type": "integer"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Signature validity period of the profile.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"type": {
						// Property: Type
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"DAYS",
								"MONTHS",
								"YEARS",
							}),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"value": {
						// Property: Value
						Type:     types.Int64Type,
						Optional: true,
						Computed: true,
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
				resource.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of tags associated with the signing profile.",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 127,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "A list of tags associated with the signing profile.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 127),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 255),
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
		Description: "A signing profile is a signing template that can be used to carry out a pre-defined signing job.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Signer::SigningProfile").WithTerraformTypeName("awscc_signer_signing_profile")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                       "Arn",
		"key":                       "Key",
		"platform_id":               "PlatformId",
		"profile_name":              "ProfileName",
		"profile_version":           "ProfileVersion",
		"profile_version_arn":       "ProfileVersionArn",
		"signature_validity_period": "SignatureValidityPeriod",
		"tags":                      "Tags",
		"type":                      "Type",
		"value":                     "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
