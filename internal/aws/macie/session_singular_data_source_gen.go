// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package macie

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_macie_session", sessionDataSource)
}

// sessionDataSource returns the Terraform awscc_macie_session data source.
// This Terraform data source corresponds to the CloudFormation AWS::Macie::Session resource.
func sessionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AwsAccountId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "AWS account ID of customer",
		//	  "type": "string"
		//	}
		"aws_account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "AWS account ID of customer",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FindingPublishingFrequency
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "SIX_HOURS",
		//	  "description": "A enumeration value that specifies how frequently finding updates are published.",
		//	  "enum": [
		//	    "FIFTEEN_MINUTES",
		//	    "ONE_HOUR",
		//	    "SIX_HOURS"
		//	  ],
		//	  "type": "string"
		//	}
		"finding_publishing_frequency": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A enumeration value that specifies how frequently finding updates are published.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ServiceRole
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Service role used by Macie",
		//	  "type": "string"
		//	}
		"service_role": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Service role used by Macie",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "ENABLED",
		//	  "description": "A enumeration value that specifies the status of the Macie Session.",
		//	  "enum": [
		//	    "ENABLED",
		//	    "PAUSED"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A enumeration value that specifies the status of the Macie Session.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Macie::Session",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Macie::Session").WithTerraformTypeName("awscc_macie_session")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"aws_account_id":               "AwsAccountId",
		"finding_publishing_frequency": "FindingPublishingFrequency",
		"service_role":                 "ServiceRole",
		"status":                       "Status",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}