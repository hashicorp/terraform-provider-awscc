// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

package elasticbeanstalk

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_elasticbeanstalk_application_versions", applicationVersionsDataSource)
}

// applicationVersionsDataSource returns the Terraform awscc_elasticbeanstalk_application_versions data source.
// This Terraform data source corresponds to the CloudFormation AWS::ElasticBeanstalk::ApplicationVersion resource.
func applicationVersionsDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Description: "Uniquely identifies the data source.",
			Computed:    true,
		},
		"ids": schema.SetAttribute{
			Description: "Set of Resource Identifiers.",
			ElementType: types.StringType,
			Computed:    true,
		},
	}

	schema := schema.Schema{
		Description: "Plural Data Source schema for AWS::ElasticBeanstalk::ApplicationVersion",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ElasticBeanstalk::ApplicationVersion").WithTerraformTypeName("awscc_elasticbeanstalk_application_versions")
	opts = opts.WithTerraformSchema(schema)

	v, err := NewPluralDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}