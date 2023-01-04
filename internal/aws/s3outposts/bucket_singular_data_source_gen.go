// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package s3outposts

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_s3outposts_bucket", bucketDataSource)
}

// bucketDataSource returns the Terraform awscc_s3outposts_bucket data source.
// This Terraform data source corresponds to the CloudFormation AWS::S3Outposts::Bucket resource.
func bucketDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the specified bucket.",
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^arn:[^:]+:s3-outposts:[a-zA-Z0-9\\-]+:\\d{12}:outpost\\/[^:]+\\/bucket\\/[^:]+$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the specified bucket.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: BucketName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A name for the bucket.",
		//	  "maxLength": 63,
		//	  "minLength": 3,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A name for the bucket.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LifecycleConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Rules that define how Amazon S3Outposts manages objects during their lifetime.",
		//	  "properties": {
		//	    "Rules": {
		//	      "description": "A list of lifecycle rules for individual objects in an Amazon S3Outposts bucket.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "anyOf": [
		//	          {
		//	            "required": [
		//	              "Status",
		//	              "AbortIncompleteMultipartUpload"
		//	            ]
		//	          },
		//	          {
		//	            "required": [
		//	              "Status",
		//	              "ExpirationDate"
		//	            ]
		//	          },
		//	          {
		//	            "required": [
		//	              "Status",
		//	              "ExpirationInDays"
		//	            ]
		//	          }
		//	        ],
		//	        "description": "Specifies lifecycle rules for an Amazon S3Outposts bucket. You must specify at least one of the following: AbortIncompleteMultipartUpload, ExpirationDate, ExpirationInDays.",
		//	        "properties": {
		//	          "AbortIncompleteMultipartUpload": {
		//	            "additionalProperties": false,
		//	            "description": "Specifies a lifecycle rule that stops incomplete multipart uploads to an Amazon S3Outposts bucket.",
		//	            "properties": {
		//	              "DaysAfterInitiation": {
		//	                "description": "Specifies the number of days after which Amazon S3Outposts aborts an incomplete multipart upload.",
		//	                "minimum": 0,
		//	                "type": "integer"
		//	              }
		//	            },
		//	            "required": [
		//	              "DaysAfterInitiation"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "ExpirationDate": {
		//	            "description": "Indicates when objects are deleted from Amazon S3Outposts. The date value must be in ISO 8601 format. The time is always midnight UTC.",
		//	            "pattern": "^([0-2]\\d{3})-(0[0-9]|1[0-2])-([0-2]\\d|3[01])T([01]\\d|2[0-4]):([0-5]\\d):([0-6]\\d)((\\.\\d{3})?)Z$",
		//	            "type": "string"
		//	          },
		//	          "ExpirationInDays": {
		//	            "description": "Indicates the number of days after creation when objects are deleted from Amazon S3Outposts.",
		//	            "minimum": 1,
		//	            "type": "integer"
		//	          },
		//	          "Filter": {
		//	            "additionalProperties": false,
		//	            "description": "The container for the filter of the lifecycle rule.",
		//	            "oneOf": [
		//	              {
		//	                "required": [
		//	                  "Prefix"
		//	                ]
		//	              },
		//	              {
		//	                "required": [
		//	                  "Tag"
		//	                ]
		//	              },
		//	              {
		//	                "required": [
		//	                  "AndOperator"
		//	                ]
		//	              }
		//	            ],
		//	            "properties": {
		//	              "AndOperator": {
		//	                "description": "The container for the AND condition for the lifecycle rule. A combination of Prefix and 1 or more Tags OR a minimum of 2 or more tags.",
		//	                "properties": {
		//	                  "Prefix": {
		//	                    "description": "Prefix identifies one or more objects to which the rule applies.",
		//	                    "type": "string"
		//	                  },
		//	                  "Tags": {
		//	                    "description": "All of these tags must exist in the object's tag set in order for the rule to apply.",
		//	                    "insertionOrder": false,
		//	                    "items": {
		//	                      "additionalProperties": false,
		//	                      "description": "Tag used to identify a subset of objects for an Amazon S3Outposts bucket.",
		//	                      "properties": {
		//	                        "Key": {
		//	                          "maxLength": 1024,
		//	                          "minLength": 1,
		//	                          "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:=+\\/\\-@%]*)$",
		//	                          "type": "string"
		//	                        },
		//	                        "Value": {
		//	                          "maxLength": 1024,
		//	                          "minLength": 1,
		//	                          "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:=+\\/\\-@%]*)$",
		//	                          "type": "string"
		//	                        }
		//	                      },
		//	                      "required": [
		//	                        "Key",
		//	                        "Value"
		//	                      ],
		//	                      "type": "object"
		//	                    },
		//	                    "minItems": 1,
		//	                    "type": "array",
		//	                    "uniqueItems": true
		//	                  }
		//	                },
		//	                "type": "object"
		//	              },
		//	              "Prefix": {
		//	                "description": "Object key prefix that identifies one or more objects to which this rule applies.",
		//	                "type": "string"
		//	              },
		//	              "Tag": {
		//	                "additionalProperties": false,
		//	                "description": "Specifies a tag used to identify a subset of objects for an Amazon S3Outposts bucket.",
		//	                "properties": {
		//	                  "Key": {
		//	                    "maxLength": 1024,
		//	                    "minLength": 1,
		//	                    "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:=+\\/\\-@%]*)$",
		//	                    "type": "string"
		//	                  },
		//	                  "Value": {
		//	                    "maxLength": 1024,
		//	                    "minLength": 1,
		//	                    "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:=+\\/\\-@%]*)$",
		//	                    "type": "string"
		//	                  }
		//	                },
		//	                "required": [
		//	                  "Key",
		//	                  "Value"
		//	                ],
		//	                "type": "object"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "Id": {
		//	            "description": "Unique identifier for the lifecycle rule. The value can't be longer than 255 characters.",
		//	            "maxLength": 255,
		//	            "type": "string"
		//	          },
		//	          "Status": {
		//	            "enum": [
		//	              "Enabled",
		//	              "Disabled"
		//	            ],
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    }
		//	  },
		//	  "required": [
		//	    "Rules"
		//	  ],
		//	  "type": "object"
		//	}
		"lifecycle_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Rules
				"rules": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: AbortIncompleteMultipartUpload
							"abort_incomplete_multipart_upload": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: DaysAfterInitiation
									"days_after_initiation": schema.Int64Attribute{ /*START ATTRIBUTE*/
										Description: "Specifies the number of days after which Amazon S3Outposts aborts an incomplete multipart upload.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "Specifies a lifecycle rule that stops incomplete multipart uploads to an Amazon S3Outposts bucket.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: ExpirationDate
							"expiration_date": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Indicates when objects are deleted from Amazon S3Outposts. The date value must be in ISO 8601 format. The time is always midnight UTC.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: ExpirationInDays
							"expiration_in_days": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Description: "Indicates the number of days after creation when objects are deleted from Amazon S3Outposts.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Filter
							"filter": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: AndOperator
									"and_operator": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: Prefix
											"prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
												Description: "Prefix identifies one or more objects to which the rule applies.",
												Computed:    true,
											}, /*END ATTRIBUTE*/
											// Property: Tags
											"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
												NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
													Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
														// Property: Key
														"key": schema.StringAttribute{ /*START ATTRIBUTE*/
															Computed: true,
														}, /*END ATTRIBUTE*/
														// Property: Value
														"value": schema.StringAttribute{ /*START ATTRIBUTE*/
															Computed: true,
														}, /*END ATTRIBUTE*/
													}, /*END SCHEMA*/
												}, /*END NESTED OBJECT*/
												Description: "All of these tags must exist in the object's tag set in order for the rule to apply.",
												Computed:    true,
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
										Description: "The container for the AND condition for the lifecycle rule. A combination of Prefix and 1 or more Tags OR a minimum of 2 or more tags.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: Prefix
									"prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "Object key prefix that identifies one or more objects to which this rule applies.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: Tag
									"tag": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: Key
											"key": schema.StringAttribute{ /*START ATTRIBUTE*/
												Computed: true,
											}, /*END ATTRIBUTE*/
											// Property: Value
											"value": schema.StringAttribute{ /*START ATTRIBUTE*/
												Computed: true,
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
										Description: "Specifies a tag used to identify a subset of objects for an Amazon S3Outposts bucket.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "The container for the filter of the lifecycle rule.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Id
							"id": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Unique identifier for the lifecycle rule. The value can't be longer than 255 characters.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Status
							"status": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "A list of lifecycle rules for individual objects in an Amazon S3Outposts bucket.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Rules that define how Amazon S3Outposts manages objects during their lifetime.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: OutpostId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The id of the customer outpost on which the bucket resides.",
		//	  "pattern": "^(op-[a-f0-9]{17}|\\d{12}|ec2)$",
		//	  "type": "string"
		//	}
		"outpost_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The id of the customer outpost on which the bucket resides.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An arbitrary set of tags (key-value pairs) for this S3Outposts bucket.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 1024,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 1024,
		//	        "minLength": 1,
		//	        "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:=+\\/\\-@%]*)$",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An arbitrary set of tags (key-value pairs) for this S3Outposts bucket.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::S3Outposts::Bucket",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::S3Outposts::Bucket").WithTerraformTypeName("awscc_s3outposts_bucket")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"abort_incomplete_multipart_upload": "AbortIncompleteMultipartUpload",
		"and_operator":                      "AndOperator",
		"arn":                               "Arn",
		"bucket_name":                       "BucketName",
		"days_after_initiation":             "DaysAfterInitiation",
		"expiration_date":                   "ExpirationDate",
		"expiration_in_days":                "ExpirationInDays",
		"filter":                            "Filter",
		"id":                                "Id",
		"key":                               "Key",
		"lifecycle_configuration":           "LifecycleConfiguration",
		"outpost_id":                        "OutpostId",
		"prefix":                            "Prefix",
		"rules":                             "Rules",
		"status":                            "Status",
		"tag":                               "Tag",
		"tags":                              "Tags",
		"value":                             "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}