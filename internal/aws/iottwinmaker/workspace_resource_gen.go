// Code generated by generators/resource/main.go; DO NOT EDIT.

package iottwinmaker

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
	registry.AddResourceFactory("awscc_iottwinmaker_workspace", workspaceResource)
}

// workspaceResource returns the Terraform awscc_iottwinmaker_workspace resource.
// This Terraform resource corresponds to the CloudFormation AWS::IoTTwinMaker::Workspace resource.
func workspaceResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the workspace.",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "pattern": "arn:((aws)|(aws-cn)|(aws-us-gov)):iottwinmaker:[a-z0-9-]+:[0-9]{12}:[\\/a-zA-Z0-9_\\-\\.:]+",
			//   "type": "string"
			// }
			Description: "The ARN of the workspace.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"creation_date_time": {
			// Property: CreationDateTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The date and time when the workspace was created.",
			//   "format": "date-time",
			//   "type": "string"
			// }
			Description: "The date and time when the workspace was created.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the workspace.",
			//   "maxLength": 512,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "The description of the workspace.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 512),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"role": {
			// Property: Role
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the execution role associated with the workspace.",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "pattern": "arn:((aws)|(aws-cn)|(aws-us-gov)):iam::[0-9]{12}:role/.*",
			//   "type": "string"
			// }
			Description: "The ARN of the execution role associated with the workspace.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(20, 2048),
				validate.StringMatch(regexp.MustCompile("arn:((aws)|(aws-cn)|(aws-us-gov)):iam::[0-9]{12}:role/.*"), ""),
			},
		},
		"s3_location": {
			// Property: S3Location
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the S3 bucket where resources associated with the workspace are stored.",
			//   "type": "string"
			// }
			Description: "The ARN of the S3 bucket where resources associated with the workspace are stored.",
			Type:        types.StringType,
			Required:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A map of key-value pairs to associate with a resource.",
			//   "patternProperties": {
			//     "": {
			//       "maxLength": 256,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "A map of key-value pairs to associate with a resource.",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"update_date_time": {
			// Property: UpdateDateTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The date and time of the current update.",
			//   "format": "date-time",
			//   "type": "string"
			// }
			Description: "The date and time of the current update.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"workspace_id": {
			// Property: WorkspaceId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the workspace.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "[a-zA-Z_0-9][a-zA-Z_\\-0-9]*[a-zA-Z0-9]+",
			//   "type": "string"
			// }
			Description: "The ID of the workspace.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 128),
				validate.StringMatch(regexp.MustCompile("[a-zA-Z_0-9][a-zA-Z_\\-0-9]*[a-zA-Z0-9]+"), ""),
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
		Description: "Resource schema for AWS::IoTTwinMaker::Workspace",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTTwinMaker::Workspace").WithTerraformTypeName("awscc_iottwinmaker_workspace")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                "Arn",
		"creation_date_time": "CreationDateTime",
		"description":        "Description",
		"role":               "Role",
		"s3_location":        "S3Location",
		"tags":               "Tags",
		"update_date_time":   "UpdateDateTime",
		"workspace_id":       "WorkspaceId",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
