// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_cloudfront_cache_policy", cachePolicyDataSource)
}

// cachePolicyDataSource returns the Terraform awscc_cloudfront_cache_policy data source.
// This Terraform data source corresponds to the CloudFormation AWS::CloudFront::CachePolicy resource.
func cachePolicyDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CachePolicyConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Comment": {
		//	      "type": "string"
		//	    },
		//	    "DefaultTTL": {
		//	      "minimum": 0,
		//	      "type": "number"
		//	    },
		//	    "MaxTTL": {
		//	      "minimum": 0,
		//	      "type": "number"
		//	    },
		//	    "MinTTL": {
		//	      "minimum": 0,
		//	      "type": "number"
		//	    },
		//	    "Name": {
		//	      "type": "string"
		//	    },
		//	    "ParametersInCacheKeyAndForwardedToOrigin": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "CookiesConfig": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "CookieBehavior": {
		//	              "pattern": "^(none|whitelist|allExcept|all)$",
		//	              "type": "string"
		//	            },
		//	            "Cookies": {
		//	              "items": {
		//	                "type": "string"
		//	              },
		//	              "type": "array",
		//	              "uniqueItems": false
		//	            }
		//	          },
		//	          "required": [
		//	            "CookieBehavior"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "EnableAcceptEncodingBrotli": {
		//	          "type": "boolean"
		//	        },
		//	        "EnableAcceptEncodingGzip": {
		//	          "type": "boolean"
		//	        },
		//	        "HeadersConfig": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "HeaderBehavior": {
		//	              "pattern": "^(none|whitelist)$",
		//	              "type": "string"
		//	            },
		//	            "Headers": {
		//	              "items": {
		//	                "type": "string"
		//	              },
		//	              "type": "array",
		//	              "uniqueItems": false
		//	            }
		//	          },
		//	          "required": [
		//	            "HeaderBehavior"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "QueryStringsConfig": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "QueryStringBehavior": {
		//	              "pattern": "^(none|whitelist|allExcept|all)$",
		//	              "type": "string"
		//	            },
		//	            "QueryStrings": {
		//	              "items": {
		//	                "type": "string"
		//	              },
		//	              "type": "array",
		//	              "uniqueItems": false
		//	            }
		//	          },
		//	          "required": [
		//	            "QueryStringBehavior"
		//	          ],
		//	          "type": "object"
		//	        }
		//	      },
		//	      "required": [
		//	        "EnableAcceptEncodingGzip",
		//	        "HeadersConfig",
		//	        "CookiesConfig",
		//	        "QueryStringsConfig"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "Name",
		//	    "MinTTL",
		//	    "MaxTTL",
		//	    "DefaultTTL",
		//	    "ParametersInCacheKeyAndForwardedToOrigin"
		//	  ],
		//	  "type": "object"
		//	}
		"cache_policy_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Comment
				"comment": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: DefaultTTL
				"default_ttl": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: MaxTTL
				"max_ttl": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: MinTTL
				"min_ttl": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Name
				"name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: ParametersInCacheKeyAndForwardedToOrigin
				"parameters_in_cache_key_and_forwarded_to_origin": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CookiesConfig
						"cookies_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: CookieBehavior
								"cookie_behavior": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: Cookies
								"cookies": schema.ListAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: EnableAcceptEncodingBrotli
						"enable_accept_encoding_brotli": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: EnableAcceptEncodingGzip
						"enable_accept_encoding_gzip": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: HeadersConfig
						"headers_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: HeaderBehavior
								"header_behavior": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: Headers
								"headers": schema.ListAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: QueryStringsConfig
						"query_strings_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: QueryStringBehavior
								"query_string_behavior": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: QueryStrings
								"query_strings": schema.ListAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
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
		// Property: LastModifiedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"last_modified_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::CloudFront::CachePolicy",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFront::CachePolicy").WithTerraformTypeName("awscc_cloudfront_cache_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cache_policy_config":           "CachePolicyConfig",
		"comment":                       "Comment",
		"cookie_behavior":               "CookieBehavior",
		"cookies":                       "Cookies",
		"cookies_config":                "CookiesConfig",
		"default_ttl":                   "DefaultTTL",
		"enable_accept_encoding_brotli": "EnableAcceptEncodingBrotli",
		"enable_accept_encoding_gzip":   "EnableAcceptEncodingGzip",
		"header_behavior":               "HeaderBehavior",
		"headers":                       "Headers",
		"headers_config":                "HeadersConfig",
		"id":                            "Id",
		"last_modified_time":            "LastModifiedTime",
		"max_ttl":                       "MaxTTL",
		"min_ttl":                       "MinTTL",
		"name":                          "Name",
		"parameters_in_cache_key_and_forwarded_to_origin": "ParametersInCacheKeyAndForwardedToOrigin",
		"query_string_behavior":                           "QueryStringBehavior",
		"query_strings":                                   "QueryStrings",
		"query_strings_config":                            "QueryStringsConfig",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}