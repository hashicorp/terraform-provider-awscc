// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

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
	registry.AddDataSourceFactory("awscc_redshift_endpoint_authorizations", endpointAuthorizationsDataSource)
}

// endpointAuthorizationsDataSource returns the Terraform awscc_redshift_endpoint_authorizations data source.
// This Terraform data source corresponds to the CloudFormation AWS::Redshift::EndpointAuthorization resource.
func endpointAuthorizationsDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"id": {
			Description: "Uniquely identifies the data source.",
			Type:        types.StringType,
			Computed:    true,
		},
		"ids": {
			Description: "Set of Resource Identifiers.",
			Type:        types.SetType{ElemType: types.StringType},
			Computed:    true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Plural Data Source schema for AWS::Redshift::EndpointAuthorization",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Redshift::EndpointAuthorization").WithTerraformTypeName("awscc_redshift_endpoint_authorizations")
	opts = opts.WithTerraformSchema(schema)

	v, err := NewPluralDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
