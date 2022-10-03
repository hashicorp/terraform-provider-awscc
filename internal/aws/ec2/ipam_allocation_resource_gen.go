// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ec2_ipam_allocation", iPAMAllocationResource)
}

// iPAMAllocationResource returns the Terraform awscc_ec2_ipam_allocation resource.
// This Terraform resource corresponds to the CloudFormation AWS::EC2::IPAMAllocation resource.
func iPAMAllocationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"cidr": {
			// Property: Cidr
			// CloudFormation resource type schema:
			// {
			//   "description": "Represents a single IPv4 or IPv6 CIDR",
			//   "type": "string"
			// }
			Description: "Represents a single IPv4 or IPv6 CIDR",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"ipam_pool_allocation_id": {
			// Property: IpamPoolAllocationId
			// CloudFormation resource type schema:
			// {
			//   "description": "Id of the allocation.",
			//   "type": "string"
			// }
			Description: "Id of the allocation.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"ipam_pool_id": {
			// Property: IpamPoolId
			// CloudFormation resource type schema:
			// {
			//   "description": "Id of the IPAM Pool.",
			//   "type": "string"
			// }
			Description: "Id of the IPAM Pool.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"netmask_length": {
			// Property: NetmaskLength
			// CloudFormation resource type schema:
			// {
			//   "description": "The desired netmask length of the allocation. If set, IPAM will choose a block of free space with this size and return the CIDR representing it.",
			//   "type": "integer"
			// }
			Description: "The desired netmask length of the allocation. If set, IPAM will choose a block of free space with this size and return the CIDR representing it.",
			Type:        types.Int64Type,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
			// NetmaskLength is a write-only property.
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
		Description: "Resource Schema of AWS::EC2::IPAMAllocation Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::IPAMAllocation").WithTerraformTypeName("awscc_ec2_ipam_allocation")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cidr":                    "Cidr",
		"description":             "Description",
		"ipam_pool_allocation_id": "IpamPoolAllocationId",
		"ipam_pool_id":            "IpamPoolId",
		"netmask_length":          "NetmaskLength",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/NetmaskLength",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
