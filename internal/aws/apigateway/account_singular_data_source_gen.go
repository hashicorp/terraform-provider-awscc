// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_apigateway_account", accountDataSource)
}

// accountDataSource returns the Terraform awscc_apigateway_account data source.
// This Terraform data source corresponds to the CloudFormation AWS::ApiGateway::Account resource.
func accountDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"cloudwatch_role_arn": {
			// Property: CloudWatchRoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of an IAM role that has write access to CloudWatch Logs in your account.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of an IAM role that has write access to CloudWatch Logs in your account.",
			Type:        types.StringType,
			Computed:    true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "Primary identifier which is manually generated.",
			//   "type": "string"
			// }
			Description: "Primary identifier which is manually generated.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::ApiGateway::Account",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::Account").WithTerraformTypeName("awscc_apigateway_account")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cloudwatch_role_arn": "CloudWatchRoleArn",
		"id":                  "Id",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
