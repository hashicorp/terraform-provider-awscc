// Code generated by generators/resource/main.go; DO NOT EDIT.

package emr

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceFactory("awscc_emr_studio_session_mapping", studioSessionMappingResource)
}

// studioSessionMappingResource returns the Terraform awscc_emr_studio_session_mapping resource.
// This Terraform resource corresponds to the CloudFormation AWS::EMR::StudioSessionMapping resource.
func studioSessionMappingResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"identity_name": {
			// Property: IdentityName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the user or group. For more information, see UserName and DisplayName in the AWS SSO Identity Store API Reference. Either IdentityName or IdentityId must be specified.",
			//   "type": "string"
			// }
			Description: "The name of the user or group. For more information, see UserName and DisplayName in the AWS SSO Identity Store API Reference. Either IdentityName or IdentityId must be specified.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"identity_type": {
			// Property: IdentityType
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies whether the identity to map to the Studio is a user or a group.",
			//   "enum": [
			//     "USER",
			//     "GROUP"
			//   ],
			//   "type": "string"
			// }
			Description: "Specifies whether the identity to map to the Studio is a user or a group.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"USER",
					"GROUP",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"session_policy_arn": {
			// Property: SessionPolicyArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) for the session policy that will be applied to the user or group. Session policies refine Studio user permissions without the need to use multiple IAM user roles.",
			//   "pattern": "^arn:aws(-(cn|us-gov))?:iam::([0-9]{12})?:policy\\/[^.]+$",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) for the session policy that will be applied to the user or group. Session policies refine Studio user permissions without the need to use multiple IAM user roles.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("^arn:aws(-(cn|us-gov))?:iam::([0-9]{12})?:policy\\/[^.]+$"), ""),
			},
		},
		"studio_id": {
			// Property: StudioId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the Amazon EMR Studio to which the user or group will be mapped.",
			//   "maxLength": 256,
			//   "minLength": 4,
			//   "pattern": "^es-[0-9A-Z]+",
			//   "type": "string"
			// }
			Description: "The ID of the Amazon EMR Studio to which the user or group will be mapped.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(4, 256),
				validate.StringMatch(regexp.MustCompile("^es-[0-9A-Z]+"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
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
		Description: "An example resource schema demonstrating some basic constructs and validation rules.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EMR::StudioSessionMapping").WithTerraformTypeName("awscc_emr_studio_session_mapping")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"identity_name":      "IdentityName",
		"identity_type":      "IdentityType",
		"session_policy_arn": "SessionPolicyArn",
		"studio_id":          "StudioId",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
