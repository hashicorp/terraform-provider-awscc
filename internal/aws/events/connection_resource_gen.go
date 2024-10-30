// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package events

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_events_connection", connectionResource)
}

// connectionResource returns the Terraform awscc_events_connection resource.
// This Terraform resource corresponds to the CloudFormation AWS::Events::Connection resource.
func connectionResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The arn of the connection resource.",
		//	  "pattern": "^arn:aws([a-z]|\\-)*:events:([a-z]|\\d|\\-)*:([0-9]{12})?:connection\\/[\\.\\-_A-Za-z0-9]+\\/[\\-A-Za-z0-9]+$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The arn of the connection resource.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AuthParameters
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "oneOf": [
		//	    {
		//	      "required": [
		//	        "BasicAuthParameters"
		//	      ]
		//	    },
		//	    {
		//	      "required": [
		//	        "OAuthParameters"
		//	      ]
		//	    },
		//	    {
		//	      "required": [
		//	        "ApiKeyAuthParameters"
		//	      ]
		//	    }
		//	  ],
		//	  "properties": {
		//	    "ApiKeyAuthParameters": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "ApiKeyName": {
		//	          "pattern": "^[ \\t]*[^\\x00-\\x1F\\x7F]+([ \\t]+[^\\x00-\\x1F\\x7F]+)*[ \\t]*$",
		//	          "type": "string"
		//	        },
		//	        "ApiKeyValue": {
		//	          "pattern": "^[ \\t]*[^\\x00-\\x1F\\x7F]+([ \\t]+[^\\x00-\\x1F\\x7F]+)*[ \\t]*$",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "ApiKeyName",
		//	        "ApiKeyValue"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "BasicAuthParameters": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "Password": {
		//	          "pattern": "^[ \\t]*[^\\x00-\\x1F\\x7F]+([ \\t]+[^\\x00-\\x1F\\x7F]+)*[ \\t]*$",
		//	          "type": "string"
		//	        },
		//	        "Username": {
		//	          "pattern": "^[ \\t]*[^\\x00-\\x1F\\x7F]+([ \\t]+[^\\x00-\\x1F\\x7F]+)*[ \\t]*$",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "Username",
		//	        "Password"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "InvocationHttpParameters": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "BodyParameters": {
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "IsValueSecret": {
		//	                "default": true,
		//	                "type": "boolean"
		//	              },
		//	              "Key": {
		//	                "type": "string"
		//	              },
		//	              "Value": {
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "Key",
		//	              "Value"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "type": "array"
		//	        },
		//	        "HeaderParameters": {
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "IsValueSecret": {
		//	                "default": true,
		//	                "type": "boolean"
		//	              },
		//	              "Key": {
		//	                "type": "string"
		//	              },
		//	              "Value": {
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "Key",
		//	              "Value"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "type": "array"
		//	        },
		//	        "QueryStringParameters": {
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "IsValueSecret": {
		//	                "default": true,
		//	                "type": "boolean"
		//	              },
		//	              "Key": {
		//	                "type": "string"
		//	              },
		//	              "Value": {
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "Key",
		//	              "Value"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "type": "array"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "OAuthParameters": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "AuthorizationEndpoint": {
		//	          "maxLength": 2048,
		//	          "minLength": 1,
		//	          "pattern": "^((%[0-9A-Fa-f]{2}|[-()_.!~*';/?:@\\x26=+$,A-Za-z0-9])+)([).!';/?:,])?$",
		//	          "type": "string"
		//	        },
		//	        "ClientParameters": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "ClientID": {
		//	              "pattern": "^[ \\t]*[^\\x00-\\x1F\\x7F]+([ \\t]+[^\\x00-\\x1F\\x7F]+)*[ \\t]*$",
		//	              "type": "string"
		//	            },
		//	            "ClientSecret": {
		//	              "pattern": "^[ \\t]*[^\\x00-\\x1F\\x7F]+([ \\t]+[^\\x00-\\x1F\\x7F]+)*[ \\t]*$",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "ClientID",
		//	            "ClientSecret"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "HttpMethod": {
		//	          "enum": [
		//	            "GET",
		//	            "POST",
		//	            "PUT"
		//	          ],
		//	          "type": "string"
		//	        },
		//	        "OAuthHttpParameters": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "BodyParameters": {
		//	              "items": {
		//	                "additionalProperties": false,
		//	                "properties": {
		//	                  "IsValueSecret": {
		//	                    "default": true,
		//	                    "type": "boolean"
		//	                  },
		//	                  "Key": {
		//	                    "type": "string"
		//	                  },
		//	                  "Value": {
		//	                    "type": "string"
		//	                  }
		//	                },
		//	                "required": [
		//	                  "Key",
		//	                  "Value"
		//	                ],
		//	                "type": "object"
		//	              },
		//	              "type": "array"
		//	            },
		//	            "HeaderParameters": {
		//	              "items": {
		//	                "additionalProperties": false,
		//	                "properties": {
		//	                  "IsValueSecret": {
		//	                    "default": true,
		//	                    "type": "boolean"
		//	                  },
		//	                  "Key": {
		//	                    "type": "string"
		//	                  },
		//	                  "Value": {
		//	                    "type": "string"
		//	                  }
		//	                },
		//	                "required": [
		//	                  "Key",
		//	                  "Value"
		//	                ],
		//	                "type": "object"
		//	              },
		//	              "type": "array"
		//	            },
		//	            "QueryStringParameters": {
		//	              "items": {
		//	                "additionalProperties": false,
		//	                "properties": {
		//	                  "IsValueSecret": {
		//	                    "default": true,
		//	                    "type": "boolean"
		//	                  },
		//	                  "Key": {
		//	                    "type": "string"
		//	                  },
		//	                  "Value": {
		//	                    "type": "string"
		//	                  }
		//	                },
		//	                "required": [
		//	                  "Key",
		//	                  "Value"
		//	                ],
		//	                "type": "object"
		//	              },
		//	              "type": "array"
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "required": [
		//	        "ClientParameters",
		//	        "AuthorizationEndpoint",
		//	        "HttpMethod"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"auth_parameters": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ApiKeyAuthParameters
				"api_key_auth_parameters": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: ApiKeyName
						"api_key_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.RegexMatches(regexp.MustCompile("^[ \\t]*[^\\x00-\\x1F\\x7F]+([ \\t]+[^\\x00-\\x1F\\x7F]+)*[ \\t]*$"), ""),
								fwvalidators.NotNullString(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: ApiKeyValue
						"api_key_value": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.RegexMatches(regexp.MustCompile("^[ \\t]*[^\\x00-\\x1F\\x7F]+([ \\t]+[^\\x00-\\x1F\\x7F]+)*[ \\t]*$"), ""),
								fwvalidators.NotNullString(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: BasicAuthParameters
				"basic_auth_parameters": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Password
						"password": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.RegexMatches(regexp.MustCompile("^[ \\t]*[^\\x00-\\x1F\\x7F]+([ \\t]+[^\\x00-\\x1F\\x7F]+)*[ \\t]*$"), ""),
								fwvalidators.NotNullString(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Username
						"username": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.RegexMatches(regexp.MustCompile("^[ \\t]*[^\\x00-\\x1F\\x7F]+([ \\t]+[^\\x00-\\x1F\\x7F]+)*[ \\t]*$"), ""),
								fwvalidators.NotNullString(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: InvocationHttpParameters
				"invocation_http_parameters": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: BodyParameters
						"body_parameters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
							NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: IsValueSecret
									"is_value_secret": schema.BoolAttribute{ /*START ATTRIBUTE*/
										Optional: true,
										Computed: true,
										Default:  booldefault.StaticBool(true),
										PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
											boolplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: Key
									"key": schema.StringAttribute{ /*START ATTRIBUTE*/
										Optional: true,
										Computed: true,
										Validators: []validator.String{ /*START VALIDATORS*/
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
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: HeaderParameters
						"header_parameters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
							NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: IsValueSecret
									"is_value_secret": schema.BoolAttribute{ /*START ATTRIBUTE*/
										Optional: true,
										Computed: true,
										Default:  booldefault.StaticBool(true),
										PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
											boolplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: Key
									"key": schema.StringAttribute{ /*START ATTRIBUTE*/
										Optional: true,
										Computed: true,
										Validators: []validator.String{ /*START VALIDATORS*/
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
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: QueryStringParameters
						"query_string_parameters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
							NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: IsValueSecret
									"is_value_secret": schema.BoolAttribute{ /*START ATTRIBUTE*/
										Optional: true,
										Computed: true,
										Default:  booldefault.StaticBool(true),
										PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
											boolplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: Key
									"key": schema.StringAttribute{ /*START ATTRIBUTE*/
										Optional: true,
										Computed: true,
										Validators: []validator.String{ /*START VALIDATORS*/
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
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: OAuthParameters
				"o_auth_parameters": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: AuthorizationEndpoint
						"authorization_endpoint": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 2048),
								stringvalidator.RegexMatches(regexp.MustCompile("^((%[0-9A-Fa-f]{2}|[-()_.!~*';/?:@\\x26=+$,A-Za-z0-9])+)([).!';/?:,])?$"), ""),
								fwvalidators.NotNullString(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: ClientParameters
						"client_parameters": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: ClientID
								"client_id": schema.StringAttribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.RegexMatches(regexp.MustCompile("^[ \\t]*[^\\x00-\\x1F\\x7F]+([ \\t]+[^\\x00-\\x1F\\x7F]+)*[ \\t]*$"), ""),
										fwvalidators.NotNullString(),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: ClientSecret
								"client_secret": schema.StringAttribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.RegexMatches(regexp.MustCompile("^[ \\t]*[^\\x00-\\x1F\\x7F]+([ \\t]+[^\\x00-\\x1F\\x7F]+)*[ \\t]*$"), ""),
										fwvalidators.NotNullString(),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
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
						// Property: HttpMethod
						"http_method": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.OneOf(
									"GET",
									"POST",
									"PUT",
								),
								fwvalidators.NotNullString(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: OAuthHttpParameters
						"o_auth_http_parameters": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: BodyParameters
								"body_parameters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
									NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: IsValueSecret
											"is_value_secret": schema.BoolAttribute{ /*START ATTRIBUTE*/
												Optional: true,
												Computed: true,
												Default:  booldefault.StaticBool(true),
												PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
													boolplanmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
											// Property: Key
											"key": schema.StringAttribute{ /*START ATTRIBUTE*/
												Optional: true,
												Computed: true,
												Validators: []validator.String{ /*START VALIDATORS*/
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
									PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
										listplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: HeaderParameters
								"header_parameters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
									NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: IsValueSecret
											"is_value_secret": schema.BoolAttribute{ /*START ATTRIBUTE*/
												Optional: true,
												Computed: true,
												Default:  booldefault.StaticBool(true),
												PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
													boolplanmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
											// Property: Key
											"key": schema.StringAttribute{ /*START ATTRIBUTE*/
												Optional: true,
												Computed: true,
												Validators: []validator.String{ /*START VALIDATORS*/
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
									PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
										listplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: QueryStringParameters
								"query_string_parameters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
									NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: IsValueSecret
											"is_value_secret": schema.BoolAttribute{ /*START ATTRIBUTE*/
												Optional: true,
												Computed: true,
												Default:  booldefault.StaticBool(true),
												PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
													boolplanmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
											// Property: Key
											"key": schema.StringAttribute{ /*START ATTRIBUTE*/
												Optional: true,
												Computed: true,
												Validators: []validator.String{ /*START VALIDATORS*/
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
									PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
										listplanmodifier.UseStateForUnknown(),
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
			// AuthParameters is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: AuthorizationType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "API_KEY",
		//	    "BASIC",
		//	    "OAUTH_CLIENT_CREDENTIALS"
		//	  ],
		//	  "type": "string"
		//	}
		"authorization_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"API_KEY",
					"BASIC",
					"OAUTH_CLIENT_CREDENTIALS",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Description of the connection.",
		//	  "maxLength": 512,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Description of the connection.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(512),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of the connection.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "[\\.\\-_A-Za-z0-9]+",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of the connection.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 64),
				stringvalidator.RegexMatches(regexp.MustCompile("[\\.\\-_A-Za-z0-9]+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SecretArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The arn of the secrets manager secret created in the customer account.",
		//	  "pattern": "^arn:aws([a-z]|\\-)*:secretsmanager:([a-z]|\\d|\\-)*:([0-9]{12})?:secret:[\\/_+=\\.@\\-A-Za-z0-9]+$",
		//	  "type": "string"
		//	}
		"secret_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The arn of the secrets manager secret created in the customer account.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
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
		Description: "Resource Type definition for AWS::Events::Connection.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Events::Connection").WithTerraformTypeName("awscc_events_connection")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"api_key_auth_parameters":    "ApiKeyAuthParameters",
		"api_key_name":               "ApiKeyName",
		"api_key_value":              "ApiKeyValue",
		"arn":                        "Arn",
		"auth_parameters":            "AuthParameters",
		"authorization_endpoint":     "AuthorizationEndpoint",
		"authorization_type":         "AuthorizationType",
		"basic_auth_parameters":      "BasicAuthParameters",
		"body_parameters":            "BodyParameters",
		"client_id":                  "ClientID",
		"client_parameters":          "ClientParameters",
		"client_secret":              "ClientSecret",
		"description":                "Description",
		"header_parameters":          "HeaderParameters",
		"http_method":                "HttpMethod",
		"invocation_http_parameters": "InvocationHttpParameters",
		"is_value_secret":            "IsValueSecret",
		"key":                        "Key",
		"name":                       "Name",
		"o_auth_http_parameters":     "OAuthHttpParameters",
		"o_auth_parameters":          "OAuthParameters",
		"password":                   "Password",
		"query_string_parameters":    "QueryStringParameters",
		"secret_arn":                 "SecretArn",
		"username":                   "Username",
		"value":                      "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/AuthParameters",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
