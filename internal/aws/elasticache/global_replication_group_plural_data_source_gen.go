// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_elasticache_global_replication_groups", globalReplicationGroupsDataSource)
}

// globalReplicationGroupsDataSource returns the Terraform awscc_elasticache_global_replication_groups data source.
// This Terraform data source corresponds to the CloudFormation AWS::ElastiCache::GlobalReplicationGroup resource.
func globalReplicationGroupsDataSource(ctx context.Context) (datasource.DataSource, error) {
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
		Description: "Plural Data Source schema for AWS::ElastiCache::GlobalReplicationGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ElastiCache::GlobalReplicationGroup").WithTerraformTypeName("awscc_elasticache_global_replication_groups")
	opts = opts.WithTerraformSchema(schema)

	v, err := NewPluralDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
