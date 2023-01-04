// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ec2_ipam_allocation", iPAMAllocationDataSource)
}

// iPAMAllocationDataSource returns the Terraform awscc_ec2_ipam_allocation data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::IPAMAllocation resource.
func iPAMAllocationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Cidr
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Represents a single IPv4 or IPv6 CIDR",
		//	  "type": "string"
		//	}
		"cidr": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Represents a single IPv4 or IPv6 CIDR",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: IpamPoolAllocationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Id of the allocation.",
		//	  "type": "string"
		//	}
		"ipam_pool_allocation_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Id of the allocation.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IpamPoolId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Id of the IPAM Pool.",
		//	  "type": "string"
		//	}
		"ipam_pool_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Id of the IPAM Pool.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: NetmaskLength
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The desired netmask length of the allocation. If set, IPAM will choose a block of free space with this size and return the CIDR representing it.",
		//	  "type": "integer"
		//	}
		"netmask_length": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The desired netmask length of the allocation. If set, IPAM will choose a block of free space with this size and return the CIDR representing it.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::IPAMAllocation",
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