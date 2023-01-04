// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_cloudfront_cloudfront_origin_access_identity", cloudFrontOriginAccessIdentityDataSource)
}

// cloudFrontOriginAccessIdentityDataSource returns the Terraform awscc_cloudfront_cloudfront_origin_access_identity data source.
// This Terraform data source corresponds to the CloudFormation AWS::CloudFront::CloudFrontOriginAccessIdentity resource.
func cloudFrontOriginAccessIdentityDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CloudFrontOriginAccessIdentityConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Comment": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Comment"
		//	  ],
		//	  "type": "object"
		//	}
		"cloudfront_origin_access_identity_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Comment
				"comment": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: S3CanonicalUserId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"s3_canonical_user_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::CloudFront::CloudFrontOriginAccessIdentity",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFront::CloudFrontOriginAccessIdentity").WithTerraformTypeName("awscc_cloudfront_cloudfront_origin_access_identity")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cloudfront_origin_access_identity_config": "CloudFrontOriginAccessIdentityConfig",
		"comment":              "Comment",
		"id":                   "Id",
		"s3_canonical_user_id": "S3CanonicalUserId",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}