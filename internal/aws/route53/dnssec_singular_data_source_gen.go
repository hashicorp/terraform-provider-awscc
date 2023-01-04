// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package route53

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_route53_dnssec", dNSSECDataSource)
}

// dNSSECDataSource returns the Terraform awscc_route53_dnssec data source.
// This Terraform data source corresponds to the CloudFormation AWS::Route53::DNSSEC resource.
func dNSSECDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: HostedZoneId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique string (ID) used to identify a hosted zone.",
		//	  "pattern": "^[A-Z0-9]{1,32}$",
		//	  "type": "string"
		//	}
		"hosted_zone_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique string (ID) used to identify a hosted zone.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Route53::DNSSEC",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53::DNSSEC").WithTerraformTypeName("awscc_route53_dnssec")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"hosted_zone_id": "HostedZoneId",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}