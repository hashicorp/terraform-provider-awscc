// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotfleetwise

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_iotfleetwise_vehicle", vehicleResource)
}

// vehicleResource returns the Terraform awscc_iotfleetwise_vehicle resource.
// This Terraform resource corresponds to the CloudFormation AWS::IoTFleetWise::Vehicle resource.
func vehicleResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AssociationBehavior
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "CreateIotThing",
		//	    "ValidateIotThingExists"
		//	  ],
		//	  "type": "string"
		//	}
		"association_behavior": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"CreateIotThing",
					"ValidateIotThingExists",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// AssociationBehavior is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Attributes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "patternProperties": {
		//	    "": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"attributes":        // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"creation_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType: timetypes.RFC3339Type{},
			Computed:   true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DecoderManifestArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"decoder_manifest_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
		}, /*END ATTRIBUTE*/
		// Property: LastModificationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"last_modification_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType: timetypes.RFC3339Type{},
			Computed:   true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ModelManifestArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"model_manifest_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z\\d\\-_:]+$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 100),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d\\-_:]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StateTemplates
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Identifier": {
		//	        "maxLength": 100,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "StateTemplateUpdateStrategy": {
		//	        "properties": {
		//	          "OnChange": {
		//	            "additionalProperties": false,
		//	            "type": "object"
		//	          },
		//	          "Periodic": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "StateTemplateUpdateRate": {
		//	                "additionalProperties": false,
		//	                "properties": {
		//	                  "Unit": {
		//	                    "enum": [
		//	                      "MILLISECOND",
		//	                      "SECOND",
		//	                      "MINUTE",
		//	                      "HOUR"
		//	                    ],
		//	                    "type": "string"
		//	                  },
		//	                  "Value": {
		//	                    "minimum": 1,
		//	                    "type": "number"
		//	                  }
		//	                },
		//	                "required": [
		//	                  "Unit",
		//	                  "Value"
		//	                ],
		//	                "type": "object"
		//	              }
		//	            },
		//	            "required": [
		//	              "StateTemplateUpdateRate"
		//	            ],
		//	            "type": "object"
		//	          }
		//	        },
		//	        "type": "object"
		//	      }
		//	    },
		//	    "required": [
		//	      "Identifier",
		//	      "StateTemplateUpdateStrategy"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 20,
		//	  "minItems": 0,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"state_templates": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Identifier
					"identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 100),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: StateTemplateUpdateStrategy
					"state_template_update_strategy": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: OnChange
							"on_change": schema.StringAttribute{ /*START ATTRIBUTE*/
								CustomType: jsontypes.NormalizedType{},
								Optional:   true,
								Computed:   true,
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Periodic
							"periodic": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: StateTemplateUpdateRate
									"state_template_update_rate": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: Unit
											"unit": schema.StringAttribute{ /*START ATTRIBUTE*/
												Optional: true,
												Computed: true,
												Validators: []validator.String{ /*START VALIDATORS*/
													stringvalidator.OneOf(
														"MILLISECOND",
														"SECOND",
														"MINUTE",
														"HOUR",
													),
													fwvalidators.NotNullString(),
												}, /*END VALIDATORS*/
												PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
													stringplanmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
											// Property: Value
											"value": schema.Float64Attribute{ /*START ATTRIBUTE*/
												Optional: true,
												Computed: true,
												Validators: []validator.Float64{ /*START VALIDATORS*/
													float64validator.AtLeast(1.000000),
													fwvalidators.NotNullFloat64(),
												}, /*END VALIDATORS*/
												PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
													float64planmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
										Optional: true,
										Computed: true,
										Validators: []validator.Object{ /*START VALIDATORS*/
											fwvalidators.NotNullObject(),
										}, /*END VALIDATORS*/
										PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
											objectplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
									objectplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Optional: true,
						Computed: true,
						Validators: []validator.Object{ /*START VALIDATORS*/
							fwvalidators.NotNullObject(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeBetween(0, 20),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
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
		//	  "maxItems": 50,
		//	  "minItems": 0,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeBetween(0, 50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	// Corresponds to CloudFormation primaryIdentifier.
	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Definition of AWS::IoTFleetWise::Vehicle Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTFleetWise::Vehicle").WithTerraformTypeName("awscc_iotfleetwise_vehicle")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                            "Arn",
		"association_behavior":           "AssociationBehavior",
		"attributes":                     "Attributes",
		"creation_time":                  "CreationTime",
		"decoder_manifest_arn":           "DecoderManifestArn",
		"identifier":                     "Identifier",
		"key":                            "Key",
		"last_modification_time":         "LastModificationTime",
		"model_manifest_arn":             "ModelManifestArn",
		"name":                           "Name",
		"on_change":                      "OnChange",
		"periodic":                       "Periodic",
		"state_template_update_rate":     "StateTemplateUpdateRate",
		"state_template_update_strategy": "StateTemplateUpdateStrategy",
		"state_templates":                "StateTemplates",
		"tags":                           "Tags",
		"unit":                           "Unit",
		"value":                          "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/AssociationBehavior",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
