// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package rds

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_rds_db_proxy_target_group", dBProxyTargetGroupDataSource)
}

// dBProxyTargetGroupDataSource returns the Terraform awscc_rds_db_proxy_target_group data source.
// This Terraform data source corresponds to the CloudFormation AWS::RDS::DBProxyTargetGroup resource.
func dBProxyTargetGroupDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"connection_pool_configuration_info": {
			// Property: ConnectionPoolConfigurationInfo
			// CloudFormation resource type schema:
			// {
			//   "properties": {
			//     "ConnectionBorrowTimeout": {
			//       "description": "The number of seconds for a proxy to wait for a connection to become available in the connection pool.",
			//       "type": "integer"
			//     },
			//     "InitQuery": {
			//       "description": "One or more SQL statements for the proxy to run when opening each new database connection.",
			//       "type": "string"
			//     },
			//     "MaxConnectionsPercent": {
			//       "description": "The maximum size of the connection pool for each target in a target group.",
			//       "maximum": 100,
			//       "minimum": 0,
			//       "type": "integer"
			//     },
			//     "MaxIdleConnectionsPercent": {
			//       "description": "Controls how actively the proxy closes idle database connections in the connection pool.",
			//       "maximum": 100,
			//       "minimum": 0,
			//       "type": "integer"
			//     },
			//     "SessionPinningFilters": {
			//       "description": "Each item in the list represents a class of SQL operations that normally cause all later statements in a session using a proxy to be pinned to the same underlying database connection.",
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"connection_borrow_timeout": {
						// Property: ConnectionBorrowTimeout
						Description: "The number of seconds for a proxy to wait for a connection to become available in the connection pool.",
						Type:        types.Int64Type,
						Computed:    true,
					},
					"init_query": {
						// Property: InitQuery
						Description: "One or more SQL statements for the proxy to run when opening each new database connection.",
						Type:        types.StringType,
						Computed:    true,
					},
					"max_connections_percent": {
						// Property: MaxConnectionsPercent
						Description: "The maximum size of the connection pool for each target in a target group.",
						Type:        types.Int64Type,
						Computed:    true,
					},
					"max_idle_connections_percent": {
						// Property: MaxIdleConnectionsPercent
						Description: "Controls how actively the proxy closes idle database connections in the connection pool.",
						Type:        types.Int64Type,
						Computed:    true,
					},
					"session_pinning_filters": {
						// Property: SessionPinningFilters
						Description: "Each item in the list represents a class of SQL operations that normally cause all later statements in a session using a proxy to be pinned to the same underlying database connection.",
						Type:        types.ListType{ElemType: types.StringType},
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"db_cluster_identifiers": {
			// Property: DBClusterIdentifiers
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Computed: true,
		},
		"db_instance_identifiers": {
			// Property: DBInstanceIdentifiers
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Computed: true,
		},
		"db_proxy_name": {
			// Property: DBProxyName
			// CloudFormation resource type schema:
			// {
			//   "description": "The identifier for the proxy.",
			//   "maxLength": 64,
			//   "pattern": "[A-z][0-z]*",
			//   "type": "string"
			// }
			Description: "The identifier for the proxy.",
			Type:        types.StringType,
			Computed:    true,
		},
		"target_group_arn": {
			// Property: TargetGroupArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) representing the target group.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) representing the target group.",
			Type:        types.StringType,
			Computed:    true,
		},
		"target_group_name": {
			// Property: TargetGroupName
			// CloudFormation resource type schema:
			// {
			//   "description": "The identifier for the DBProxyTargetGroup",
			//   "enum": [
			//     "default"
			//   ],
			//   "type": "string"
			// }
			Description: "The identifier for the DBProxyTargetGroup",
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
		Description: "Data Source schema for AWS::RDS::DBProxyTargetGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::RDS::DBProxyTargetGroup").WithTerraformTypeName("awscc_rds_db_proxy_target_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"connection_borrow_timeout":          "ConnectionBorrowTimeout",
		"connection_pool_configuration_info": "ConnectionPoolConfigurationInfo",
		"db_cluster_identifiers":             "DBClusterIdentifiers",
		"db_instance_identifiers":            "DBInstanceIdentifiers",
		"db_proxy_name":                      "DBProxyName",
		"init_query":                         "InitQuery",
		"max_connections_percent":            "MaxConnectionsPercent",
		"max_idle_connections_percent":       "MaxIdleConnectionsPercent",
		"session_pinning_filters":            "SessionPinningFilters",
		"target_group_arn":                   "TargetGroupArn",
		"target_group_name":                  "TargetGroupName",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
