// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_redshift_endpoint_authorization", endpointAuthorizationDataSource)
}

// endpointAuthorizationDataSource returns the Terraform awscc_redshift_endpoint_authorization data source.
// This Terraform data source corresponds to the CloudFormation AWS::Redshift::EndpointAuthorization resource.
func endpointAuthorizationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"account": {
			// Property: Account
			// CloudFormation resource type schema:
			// {
			//   "description": "The target AWS account ID to grant or revoke access for.",
			//   "pattern": "^\\d{12}$",
			//   "type": "string"
			// }
			Description: "The target AWS account ID to grant or revoke access for.",
			Type:        types.StringType,
			Computed:    true,
		},
		"allowed_all_vp_cs": {
			// Property: AllowedAllVPCs
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates whether all VPCs in the grantee account are allowed access to the cluster.",
			//   "type": "boolean"
			// }
			Description: "Indicates whether all VPCs in the grantee account are allowed access to the cluster.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"allowed_vp_cs": {
			// Property: AllowedVPCs
			// CloudFormation resource type schema:
			// {
			//   "description": "The VPCs allowed access to the cluster.",
			//   "insertionOrder": false,
			//   "items": {
			//     "pattern": "^vpc-[A-Za-z0-9]{1,17}$",
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "The VPCs allowed access to the cluster.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"authorize_time": {
			// Property: AuthorizeTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The time (UTC) when the authorization was created.",
			//   "type": "string"
			// }
			Description: "The time (UTC) when the authorization was created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"cluster_identifier": {
			// Property: ClusterIdentifier
			// CloudFormation resource type schema:
			// {
			//   "description": "The cluster identifier.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The cluster identifier.",
			Type:        types.StringType,
			Computed:    true,
		},
		"cluster_status": {
			// Property: ClusterStatus
			// CloudFormation resource type schema:
			// {
			//   "description": "The status of the cluster.",
			//   "type": "string"
			// }
			Description: "The status of the cluster.",
			Type:        types.StringType,
			Computed:    true,
		},
		"endpoint_count": {
			// Property: EndpointCount
			// CloudFormation resource type schema:
			// {
			//   "description": "The number of Redshift-managed VPC endpoints created for the authorization.",
			//   "type": "integer"
			// }
			Description: "The number of Redshift-managed VPC endpoints created for the authorization.",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"force": {
			// Property: Force
			// CloudFormation resource type schema:
			// {
			//   "description": " Indicates whether to force the revoke action. If true, the Redshift-managed VPC endpoints associated with the endpoint authorization are also deleted.",
			//   "type": "boolean"
			// }
			Description: " Indicates whether to force the revoke action. If true, the Redshift-managed VPC endpoints associated with the endpoint authorization are also deleted.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"grantee": {
			// Property: Grantee
			// CloudFormation resource type schema:
			// {
			//   "description": "The AWS account ID of the grantee of the cluster.",
			//   "pattern": "^\\d{12}$",
			//   "type": "string"
			// }
			Description: "The AWS account ID of the grantee of the cluster.",
			Type:        types.StringType,
			Computed:    true,
		},
		"grantor": {
			// Property: Grantor
			// CloudFormation resource type schema:
			// {
			//   "description": "The AWS account ID of the cluster owner.",
			//   "pattern": "^\\d{12}$",
			//   "type": "string"
			// }
			Description: "The AWS account ID of the cluster owner.",
			Type:        types.StringType,
			Computed:    true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "The status of the authorization action.",
			//   "type": "string"
			// }
			Description: "The status of the authorization action.",
			Type:        types.StringType,
			Computed:    true,
		},
		"vpc_ids": {
			// Property: VpcIds
			// CloudFormation resource type schema:
			// {
			//   "description": "The virtual private cloud (VPC) identifiers to grant or revoke access to.",
			//   "insertionOrder": false,
			//   "items": {
			//     "pattern": "^vpc-[A-Za-z0-9]{1,17}$",
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "The virtual private cloud (VPC) identifiers to grant or revoke access to.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Redshift::EndpointAuthorization",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Redshift::EndpointAuthorization").WithTerraformTypeName("awscc_redshift_endpoint_authorization")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account":            "Account",
		"allowed_all_vp_cs":  "AllowedAllVPCs",
		"allowed_vp_cs":      "AllowedVPCs",
		"authorize_time":     "AuthorizeTime",
		"cluster_identifier": "ClusterIdentifier",
		"cluster_status":     "ClusterStatus",
		"endpoint_count":     "EndpointCount",
		"force":              "Force",
		"grantee":            "Grantee",
		"grantor":            "Grantor",
		"status":             "Status",
		"vpc_ids":            "VpcIds",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
