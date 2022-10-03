// Code generated by generators/resource/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceFactory("awscc_apigateway_usage_plan", usagePlanResource)
}

// usagePlanResource returns the Terraform awscc_apigateway_usage_plan resource.
// This Terraform resource corresponds to the CloudFormation AWS::ApiGateway::UsagePlan resource.
func usagePlanResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"api_stages": {
			// Property: ApiStages
			// CloudFormation resource type schema:
			// {
			//   "description": "The API stages to associate with this usage plan.",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "ApiId": {
			//         "description": "The ID of an API that is in the specified Stage property that you want to associate with the usage plan.",
			//         "type": "string"
			//       },
			//       "Stage": {
			//         "description": "The name of the stage to associate with the usage plan.",
			//         "type": "string"
			//       },
			//       "Throttle": {
			//         "additionalProperties": false,
			//         "description": "Map containing method-level throttling information for an API stage in a usage plan. The key for the map is the path and method for which to configure custom throttling, for example, '/pets/GET'. Duplicates are not allowed.",
			//         "patternProperties": {
			//           "": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "BurstLimit": {
			//                 "description": "The maximum API request rate limit over a time ranging from one to a few seconds. The maximum API request rate limit depends on whether the underlying token bucket is at its full capacity.",
			//                 "minimum": 0,
			//                 "type": "integer"
			//               },
			//               "RateLimit": {
			//                 "description": "The API request steady-state rate limit (average requests per second over an extended period of time).",
			//                 "minimum": 0,
			//                 "type": "number"
			//               }
			//             },
			//             "type": "object"
			//           }
			//         },
			//         "type": "object"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "The API stages to associate with this usage plan.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"api_id": {
						// Property: ApiId
						Description: "The ID of an API that is in the specified Stage property that you want to associate with the usage plan.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"stage": {
						// Property: Stage
						Description: "The name of the stage to associate with the usage plan.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"throttle": {
						// Property: Throttle
						Description: "Map containing method-level throttling information for an API stage in a usage plan. The key for the map is the path and method for which to configure custom throttling, for example, '/pets/GET'. Duplicates are not allowed.",
						// Pattern: ""
						Attributes: tfsdk.MapNestedAttributes(
							map[string]tfsdk.Attribute{
								"burst_limit": {
									// Property: BurstLimit
									Description: "The maximum API request rate limit over a time ranging from one to a few seconds. The maximum API request rate limit depends on whether the underlying token bucket is at its full capacity.",
									Type:        types.Int64Type,
									Optional:    true,
									Computed:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.IntAtLeast(0),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"rate_limit": {
									// Property: RateLimit
									Description: "The API request steady-state rate limit (average requests per second over an extended period of time).",
									Type:        types.Float64Type,
									Optional:    true,
									Computed:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.FloatAtLeast(0.000000),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.UniqueItems(),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "A description of the usage plan.",
			//   "type": "string"
			// }
			Description: "A description of the usage plan.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "The provider-assigned unique ID for this managed resource.",
			//   "type": "string"
			// }
			Description: "The provider-assigned unique ID for this managed resource.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"quota": {
			// Property: Quota
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Configures the number of requests that users can make within a given interval.",
			//   "properties": {
			//     "Limit": {
			//       "description": "The maximum number of requests that users can make within the specified time period.",
			//       "minimum": 0,
			//       "type": "integer"
			//     },
			//     "Offset": {
			//       "description": "For the initial time period, the number of requests to subtract from the specified limit. When you first implement a usage plan, the plan might start in the middle of the week or month. With this property, you can decrease the limit for this initial time period.",
			//       "minimum": 0,
			//       "type": "integer"
			//     },
			//     "Period": {
			//       "description": "The time period for which the maximum limit of requests applies, such as DAY or WEEK. For valid values, see the period property for the UsagePlan resource in the Amazon API Gateway REST API Reference.",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Configures the number of requests that users can make within a given interval.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"limit": {
						// Property: Limit
						Description: "The maximum number of requests that users can make within the specified time period.",
						Type:        types.Int64Type,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntAtLeast(0),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"offset": {
						// Property: Offset
						Description: "For the initial time period, the number of requests to subtract from the specified limit. When you first implement a usage plan, the plan might start in the middle of the week or month. With this property, you can decrease the limit for this initial time period.",
						Type:        types.Int64Type,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntAtLeast(0),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"period": {
						// Property: Period
						Description: "The time period for which the maximum limit of requests applies, such as DAY or WEEK. For valid values, see the period property for the UsagePlan resource in the Amazon API Gateway REST API Reference.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of arbitrary tags (key-value pairs) to associate with the usage plan.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "An array of arbitrary tags (key-value pairs) to associate with the usage plan.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
				resource.UseStateForUnknown(),
			},
		},
		"throttle": {
			// Property: Throttle
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Configures the overall request rate (average requests per second) and burst capacity.",
			//   "properties": {
			//     "BurstLimit": {
			//       "description": "The maximum API request rate limit over a time ranging from one to a few seconds. The maximum API request rate limit depends on whether the underlying token bucket is at its full capacity.",
			//       "minimum": 0,
			//       "type": "integer"
			//     },
			//     "RateLimit": {
			//       "description": "The API request steady-state rate limit (average requests per second over an extended period of time).",
			//       "minimum": 0,
			//       "type": "number"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Configures the overall request rate (average requests per second) and burst capacity.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"burst_limit": {
						// Property: BurstLimit
						Description: "The maximum API request rate limit over a time ranging from one to a few seconds. The maximum API request rate limit depends on whether the underlying token bucket is at its full capacity.",
						Type:        types.Int64Type,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntAtLeast(0),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"rate_limit": {
						// Property: RateLimit
						Description: "The API request steady-state rate limit (average requests per second over an extended period of time).",
						Type:        types.Float64Type,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.FloatAtLeast(0.000000),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"usage_plan_name": {
			// Property: UsagePlanName
			// CloudFormation resource type schema:
			// {
			//   "description": "A name for the usage plan.",
			//   "type": "string"
			// }
			Description: "A name for the usage plan.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::ApiGateway::UsagePlan",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::UsagePlan").WithTerraformTypeName("awscc_apigateway_usage_plan")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"api_id":          "ApiId",
		"api_stages":      "ApiStages",
		"burst_limit":     "BurstLimit",
		"description":     "Description",
		"id":              "Id",
		"key":             "Key",
		"limit":           "Limit",
		"offset":          "Offset",
		"period":          "Period",
		"quota":           "Quota",
		"rate_limit":      "RateLimit",
		"stage":           "Stage",
		"tags":            "Tags",
		"throttle":        "Throttle",
		"usage_plan_name": "UsagePlanName",
		"value":           "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
