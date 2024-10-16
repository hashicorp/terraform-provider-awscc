// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package b2bi

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_b2bi_transformer", transformerDataSource)
}

// transformerDataSource returns the Terraform awscc_b2bi_transformer data source.
// This Terraform data source corresponds to the CloudFormation AWS::B2BI::Transformer resource.
func transformerDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType: timetypes.RFC3339Type{},
			Computed:   true,
		}, /*END ATTRIBUTE*/
		// Property: EdiType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "properties": {
		//	    "X12Details": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "TransactionSet": {
		//	          "enum": [
		//	            "X12_110",
		//	            "X12_180",
		//	            "X12_204",
		//	            "X12_210",
		//	            "X12_211",
		//	            "X12_214",
		//	            "X12_215",
		//	            "X12_259",
		//	            "X12_260",
		//	            "X12_266",
		//	            "X12_269",
		//	            "X12_270",
		//	            "X12_271",
		//	            "X12_274",
		//	            "X12_275",
		//	            "X12_276",
		//	            "X12_277",
		//	            "X12_278",
		//	            "X12_310",
		//	            "X12_315",
		//	            "X12_322",
		//	            "X12_404",
		//	            "X12_410",
		//	            "X12_417",
		//	            "X12_421",
		//	            "X12_426",
		//	            "X12_810",
		//	            "X12_820",
		//	            "X12_824",
		//	            "X12_830",
		//	            "X12_832",
		//	            "X12_834",
		//	            "X12_835",
		//	            "X12_837",
		//	            "X12_844",
		//	            "X12_846",
		//	            "X12_849",
		//	            "X12_850",
		//	            "X12_852",
		//	            "X12_855",
		//	            "X12_856",
		//	            "X12_860",
		//	            "X12_861",
		//	            "X12_864",
		//	            "X12_865",
		//	            "X12_869",
		//	            "X12_870",
		//	            "X12_940",
		//	            "X12_945",
		//	            "X12_990",
		//	            "X12_997",
		//	            "X12_999",
		//	            "X12_270_X279",
		//	            "X12_271_X279",
		//	            "X12_275_X210",
		//	            "X12_275_X211",
		//	            "X12_276_X212",
		//	            "X12_277_X212",
		//	            "X12_277_X214",
		//	            "X12_277_X364",
		//	            "X12_278_X217",
		//	            "X12_820_X218",
		//	            "X12_820_X306",
		//	            "X12_824_X186",
		//	            "X12_834_X220",
		//	            "X12_834_X307",
		//	            "X12_834_X318",
		//	            "X12_835_X221",
		//	            "X12_837_X222",
		//	            "X12_837_X223",
		//	            "X12_837_X224",
		//	            "X12_837_X291",
		//	            "X12_837_X292",
		//	            "X12_837_X298",
		//	            "X12_999_X231"
		//	          ],
		//	          "type": "string"
		//	        },
		//	        "Version": {
		//	          "enum": [
		//	            "VERSION_4010",
		//	            "VERSION_4030",
		//	            "VERSION_5010",
		//	            "VERSION_5010_HIPAA"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"edi_type": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: X12Details
				"x12_details": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: TransactionSet
						"transaction_set": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: Version
						"version": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: FileFormat
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "XML",
		//	    "JSON",
		//	    "NOT_USED"
		//	  ],
		//	  "type": "string"
		//	}
		"file_format": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: InputConversion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "FormatOptions": {
		//	      "properties": {
		//	        "X12": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "TransactionSet": {
		//	              "enum": [
		//	                "X12_110",
		//	                "X12_180",
		//	                "X12_204",
		//	                "X12_210",
		//	                "X12_211",
		//	                "X12_214",
		//	                "X12_215",
		//	                "X12_259",
		//	                "X12_260",
		//	                "X12_266",
		//	                "X12_269",
		//	                "X12_270",
		//	                "X12_271",
		//	                "X12_274",
		//	                "X12_275",
		//	                "X12_276",
		//	                "X12_277",
		//	                "X12_278",
		//	                "X12_310",
		//	                "X12_315",
		//	                "X12_322",
		//	                "X12_404",
		//	                "X12_410",
		//	                "X12_417",
		//	                "X12_421",
		//	                "X12_426",
		//	                "X12_810",
		//	                "X12_820",
		//	                "X12_824",
		//	                "X12_830",
		//	                "X12_832",
		//	                "X12_834",
		//	                "X12_835",
		//	                "X12_837",
		//	                "X12_844",
		//	                "X12_846",
		//	                "X12_849",
		//	                "X12_850",
		//	                "X12_852",
		//	                "X12_855",
		//	                "X12_856",
		//	                "X12_860",
		//	                "X12_861",
		//	                "X12_864",
		//	                "X12_865",
		//	                "X12_869",
		//	                "X12_870",
		//	                "X12_940",
		//	                "X12_945",
		//	                "X12_990",
		//	                "X12_997",
		//	                "X12_999",
		//	                "X12_270_X279",
		//	                "X12_271_X279",
		//	                "X12_275_X210",
		//	                "X12_275_X211",
		//	                "X12_276_X212",
		//	                "X12_277_X212",
		//	                "X12_277_X214",
		//	                "X12_277_X364",
		//	                "X12_278_X217",
		//	                "X12_820_X218",
		//	                "X12_820_X306",
		//	                "X12_824_X186",
		//	                "X12_834_X220",
		//	                "X12_834_X307",
		//	                "X12_834_X318",
		//	                "X12_835_X221",
		//	                "X12_837_X222",
		//	                "X12_837_X223",
		//	                "X12_837_X224",
		//	                "X12_837_X291",
		//	                "X12_837_X292",
		//	                "X12_837_X298",
		//	                "X12_999_X231"
		//	              ],
		//	              "type": "string"
		//	            },
		//	            "Version": {
		//	              "enum": [
		//	                "VERSION_4010",
		//	                "VERSION_4030",
		//	                "VERSION_5010",
		//	                "VERSION_5010_HIPAA"
		//	              ],
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "FromFormat": {
		//	      "enum": [
		//	        "X12"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "FromFormat"
		//	  ],
		//	  "type": "object"
		//	}
		"input_conversion": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: FormatOptions
				"format_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: X12
						"x12": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: TransactionSet
								"transaction_set": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: Version
								"version": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: FromFormat
				"from_format": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Mapping
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Template": {
		//	      "maxLength": 350000,
		//	      "minLength": 0,
		//	      "type": "string"
		//	    },
		//	    "TemplateLanguage": {
		//	      "enum": [
		//	        "XSLT",
		//	        "JSONATA"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "TemplateLanguage"
		//	  ],
		//	  "type": "object"
		//	}
		"mapping": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Template
				"template": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: TemplateLanguage
				"template_language": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: MappingTemplate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "This shape is deprecated: This is a legacy trait. Please use input-conversion or output-conversion.",
		//	  "maxLength": 350000,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"mapping_template": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "This shape is deprecated: This is a legacy trait. Please use input-conversion or output-conversion.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ModifiedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"modified_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType: timetypes.RFC3339Type{},
			Computed:   true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 254,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9_-]{1,512}$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: OutputConversion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "FormatOptions": {
		//	      "properties": {
		//	        "X12": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "TransactionSet": {
		//	              "enum": [
		//	                "X12_110",
		//	                "X12_180",
		//	                "X12_204",
		//	                "X12_210",
		//	                "X12_211",
		//	                "X12_214",
		//	                "X12_215",
		//	                "X12_259",
		//	                "X12_260",
		//	                "X12_266",
		//	                "X12_269",
		//	                "X12_270",
		//	                "X12_271",
		//	                "X12_274",
		//	                "X12_275",
		//	                "X12_276",
		//	                "X12_277",
		//	                "X12_278",
		//	                "X12_310",
		//	                "X12_315",
		//	                "X12_322",
		//	                "X12_404",
		//	                "X12_410",
		//	                "X12_417",
		//	                "X12_421",
		//	                "X12_426",
		//	                "X12_810",
		//	                "X12_820",
		//	                "X12_824",
		//	                "X12_830",
		//	                "X12_832",
		//	                "X12_834",
		//	                "X12_835",
		//	                "X12_837",
		//	                "X12_844",
		//	                "X12_846",
		//	                "X12_849",
		//	                "X12_850",
		//	                "X12_852",
		//	                "X12_855",
		//	                "X12_856",
		//	                "X12_860",
		//	                "X12_861",
		//	                "X12_864",
		//	                "X12_865",
		//	                "X12_869",
		//	                "X12_870",
		//	                "X12_940",
		//	                "X12_945",
		//	                "X12_990",
		//	                "X12_997",
		//	                "X12_999",
		//	                "X12_270_X279",
		//	                "X12_271_X279",
		//	                "X12_275_X210",
		//	                "X12_275_X211",
		//	                "X12_276_X212",
		//	                "X12_277_X212",
		//	                "X12_277_X214",
		//	                "X12_277_X364",
		//	                "X12_278_X217",
		//	                "X12_820_X218",
		//	                "X12_820_X306",
		//	                "X12_824_X186",
		//	                "X12_834_X220",
		//	                "X12_834_X307",
		//	                "X12_834_X318",
		//	                "X12_835_X221",
		//	                "X12_837_X222",
		//	                "X12_837_X223",
		//	                "X12_837_X224",
		//	                "X12_837_X291",
		//	                "X12_837_X292",
		//	                "X12_837_X298",
		//	                "X12_999_X231"
		//	              ],
		//	              "type": "string"
		//	            },
		//	            "Version": {
		//	              "enum": [
		//	                "VERSION_4010",
		//	                "VERSION_4030",
		//	                "VERSION_5010",
		//	                "VERSION_5010_HIPAA"
		//	              ],
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "ToFormat": {
		//	      "enum": [
		//	        "X12"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "ToFormat"
		//	  ],
		//	  "type": "object"
		//	}
		"output_conversion": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: FormatOptions
				"format_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: X12
						"x12": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: TransactionSet
								"transaction_set": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: Version
								"version": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: ToFormat
				"to_format": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: SampleDocument
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "This shape is deprecated: This is a legacy trait. Please use input-conversion or output-conversion.",
		//	  "maxLength": 1024,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"sample_document": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "This shape is deprecated: This is a legacy trait. Please use input-conversion or output-conversion.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SampleDocuments
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "BucketName": {
		//	      "maxLength": 63,
		//	      "minLength": 3,
		//	      "type": "string"
		//	    },
		//	    "Keys": {
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Input": {
		//	            "maxLength": 1024,
		//	            "minLength": 0,
		//	            "type": "string"
		//	          },
		//	          "Output": {
		//	            "maxLength": 1024,
		//	            "minLength": 0,
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "required": [
		//	    "BucketName",
		//	    "Keys"
		//	  ],
		//	  "type": "object"
		//	}
		"sample_documents": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: BucketName
				"bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Keys
				"keys": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Input
							"input": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Output
							"output": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "active",
		//	    "inactive"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 200,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
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
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: TransformerArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"transformer_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: TransformerId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9_-]+$",
		//	  "type": "string"
		//	}
		"transformer_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::B2BI::Transformer",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::B2BI::Transformer").WithTerraformTypeName("awscc_b2bi_transformer")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"bucket_name":       "BucketName",
		"created_at":        "CreatedAt",
		"edi_type":          "EdiType",
		"file_format":       "FileFormat",
		"format_options":    "FormatOptions",
		"from_format":       "FromFormat",
		"input":             "Input",
		"input_conversion":  "InputConversion",
		"key":               "Key",
		"keys":              "Keys",
		"mapping":           "Mapping",
		"mapping_template":  "MappingTemplate",
		"modified_at":       "ModifiedAt",
		"name":              "Name",
		"output":            "Output",
		"output_conversion": "OutputConversion",
		"sample_document":   "SampleDocument",
		"sample_documents":  "SampleDocuments",
		"status":            "Status",
		"tags":              "Tags",
		"template":          "Template",
		"template_language": "TemplateLanguage",
		"to_format":         "ToFormat",
		"transaction_set":   "TransactionSet",
		"transformer_arn":   "TransformerArn",
		"transformer_id":    "TransformerId",
		"value":             "Value",
		"version":           "Version",
		"x12":               "X12",
		"x12_details":       "X12Details",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
