// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ec2_ipam_allocation", iPAMAllocationDataSource)
}

// iPAMAllocationDataSource returns the Terraform awscc_ec2_ipam_allocation data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::IPAMAllocation resource.
func iPAMAllocationDataSource(ctx context.Context) (datasource.DataSource, error) {
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
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
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
			Computed:    true,
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
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::EC2::IPAMAllocation",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::IPAMAllocation").WithTerraformTypeName("awscc_ec2_ipam_allocation")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cidr":                    "Cidr",
		"description":             "Description",
		"ipam_pool_allocation_id": "IpamPoolAllocationId",
		"ipam_pool_id":            "IpamPoolId",
		"netmask_length":          "NetmaskLength",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
