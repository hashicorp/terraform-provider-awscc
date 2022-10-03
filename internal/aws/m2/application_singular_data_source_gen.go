// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package m2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_m2_application", applicationDataSource)
}

// applicationDataSource returns the Terraform awscc_m2_application data source.
// This Terraform data source corresponds to the CloudFormation AWS::M2::Application resource.
func applicationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"application_arn": {
			// Property: ApplicationArn
			// CloudFormation resource type schema:
			// {
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"application_id": {
			// Property: ApplicationId
			// CloudFormation resource type schema:
			// {
			//   "pattern": "^\\S{1,80}$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"definition": {
			// Property: Definition
			// CloudFormation resource type schema:
			// {
			//   "properties": {
			//     "Content": {
			//       "maxLength": 65000,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "S3Location": {
			//       "pattern": "",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"content": {
						// Property: Content
						Type:     types.StringType,
						Computed: true,
					},
					"s3_location": {
						// Property: S3Location
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 500,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"engine_type": {
			// Property: EngineType
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "microfocus",
			//     "bluage"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "pattern": "^[A-Za-z0-9][A-Za-z0-9_\\-]{1,59}$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "patternProperties": {
			//     "": {
			//       "maxLength": 256,
			//       "minLength": 0,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::M2::Application",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::M2::Application").WithTerraformTypeName("awscc_m2_application")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"application_arn": "ApplicationArn",
		"application_id":  "ApplicationId",
		"content":         "Content",
		"definition":      "Definition",
		"description":     "Description",
		"engine_type":     "EngineType",
		"name":            "Name",
		"s3_location":     "S3Location",
		"tags":            "Tags",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
