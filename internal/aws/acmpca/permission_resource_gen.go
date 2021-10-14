// Code generated by generators/resource/main.go; DO NOT EDIT.

package acmpca

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_acmpca_permission", permissionResourceType)
}

// permissionResourceType returns the Terraform awscc_acmpca_permission resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ACMPCA::Permission resource type.
func permissionResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"actions": {
			// Property: Actions
			// CloudFormation resource type schema:
			// {
			//   "description": "The actions that the specified AWS service principal can use. Actions IssueCertificate, GetCertificate and ListPermissions must be provided.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "The actions that the specified AWS service principal can use. Actions IssueCertificate, GetCertificate and ListPermissions must be provided.",
			Type:        types.ListType{ElemType: types.StringType},
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
				tfsdk.RequiresReplace(),
			},
		},
		"certificate_authority_arn": {
			// Property: CertificateAuthorityArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the Private Certificate Authority that grants the permission.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the Private Certificate Authority that grants the permission.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"principal": {
			// Property: Principal
			// CloudFormation resource type schema:
			// {
			//   "description": "The AWS service or identity that receives the permission. At this time, the only valid principal is acm.amazonaws.com.",
			//   "type": "string"
			// }
			Description: "The AWS service or identity that receives the permission. At this time, the only valid principal is acm.amazonaws.com.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"source_account": {
			// Property: SourceAccount
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the calling account.",
			//   "type": "string"
			// }
			Description: "The ID of the calling account.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				ComputedOptionalForceNew(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Permission set on private certificate authority",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ACMPCA::Permission").WithTerraformTypeName("awscc_acmpca_permission")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"actions":                   "Actions",
		"certificate_authority_arn": "CertificateAuthorityArn",
		"principal":                 "Principal",
		"source_account":            "SourceAccount",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
