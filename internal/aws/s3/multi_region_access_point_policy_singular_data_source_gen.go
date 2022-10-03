// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package s3

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_s3_multi_region_access_point_policy", multiRegionAccessPointPolicyDataSource)
}

// multiRegionAccessPointPolicyDataSource returns the Terraform awscc_s3_multi_region_access_point_policy data source.
// This Terraform data source corresponds to the CloudFormation AWS::S3::MultiRegionAccessPointPolicy resource.
func multiRegionAccessPointPolicyDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"mrap_name": {
			// Property: MrapName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the Multi Region Access Point to apply policy",
			//   "maxLength": 50,
			//   "minLength": 3,
			//   "pattern": "^[a-z0-9][-a-z0-9]{1,48}[a-z0-9]$",
			//   "type": "string"
			// }
			Description: "The name of the Multi Region Access Point to apply policy",
			Type:        types.StringType,
			Computed:    true,
		},
		"policy": {
			// Property: Policy
			// CloudFormation resource type schema:
			// {
			//   "description": "Policy document to apply to a Multi Region Access Point",
			//   "type": "object"
			// }
			Description: "Policy document to apply to a Multi Region Access Point",
			Type:        types.MapType{ElemType: types.StringType},
			Computed:    true,
		},
		"policy_status": {
			// Property: PolicyStatus
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The Policy Status associated with this Multi Region Access Point",
			//   "properties": {
			//     "IsPublic": {
			//       "description": "Specifies whether the policy is public or not.",
			//       "enum": [
			//         "true",
			//         "false"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "IsPublic"
			//   ],
			//   "type": "object"
			// }
			Description: "The Policy Status associated with this Multi Region Access Point",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"is_public": {
						// Property: IsPublic
						Description: "Specifies whether the policy is public or not.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::S3::MultiRegionAccessPointPolicy",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::S3::MultiRegionAccessPointPolicy").WithTerraformTypeName("awscc_s3_multi_region_access_point_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"is_public":     "IsPublic",
		"mrap_name":     "MrapName",
		"policy":        "Policy",
		"policy_status": "PolicyStatus",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
