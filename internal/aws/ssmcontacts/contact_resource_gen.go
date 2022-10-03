// Code generated by generators/resource/main.go; DO NOT EDIT.

package ssmcontacts

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceFactory("awscc_ssmcontacts_contact", contactResource)
}

// contactResource returns the Terraform awscc_ssmcontacts_contact resource.
// This Terraform resource corresponds to the CloudFormation AWS::SSMContacts::Contact resource.
func contactResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"alias": {
			// Property: Alias
			// CloudFormation resource type schema:
			// {
			//   "description": "Alias of the contact. String value with 20 to 256 characters. Only alphabetical, numeric characters, dash, or underscore allowed.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "^[a-z0-9_\\-\\.]*$",
			//   "type": "string"
			// }
			Description: "Alias of the contact. String value with 20 to 256 characters. Only alphabetical, numeric characters, dash, or underscore allowed.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 255),
				validate.StringMatch(regexp.MustCompile("^[a-z0-9_\\-\\.]*$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the contact.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the contact.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"display_name": {
			// Property: DisplayName
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of the contact. String value with 3 to 256 characters. Only alphabetical, space, numeric characters, dash, or underscore allowed.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9_\\-\\s]*$",
			//   "type": "string"
			// }
			Description: "Name of the contact. String value with 3 to 256 characters. Only alphabetical, space, numeric characters, dash, or underscore allowed.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 255),
				validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9_\\-\\s]*$"), ""),
			},
		},
		"plan": {
			// Property: Plan
			// CloudFormation resource type schema:
			// {
			//   "description": "The stages that an escalation plan or engagement plan engages contacts and contact methods in.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A set amount of time that an escalation plan or engagement plan engages the specified contacts or contact methods.",
			//     "properties": {
			//       "DurationInMinutes": {
			//         "description": "The time to wait until beginning the next stage.",
			//         "type": "integer"
			//       },
			//       "Targets": {
			//         "description": "The contacts or contact methods that the escalation plan or engagement plan is engaging.",
			//         "items": {
			//           "additionalProperties": false,
			//           "description": "The contacts or contact methods that the escalation plan or engagement plan is engaging.",
			//           "oneOf": [
			//             {
			//               "required": [
			//                 "ChannelTargetInfo"
			//               ]
			//             },
			//             {
			//               "required": [
			//                 "ContactTargetInfo"
			//               ]
			//             }
			//           ],
			//           "properties": {
			//             "ChannelTargetInfo": {
			//               "additionalProperties": false,
			//               "description": "Information about the contact channel that SSM Incident Manager uses to engage the contact.",
			//               "properties": {
			//                 "ChannelId": {
			//                   "description": "The Amazon Resource Name (ARN) of the contact channel.",
			//                   "type": "string"
			//                 },
			//                 "RetryIntervalInMinutes": {
			//                   "description": "The number of minutes to wait to retry sending engagement in the case the engagement initially fails.",
			//                   "type": "integer"
			//                 }
			//               },
			//               "required": [
			//                 "ChannelId",
			//                 "RetryIntervalInMinutes"
			//               ],
			//               "type": "object"
			//             },
			//             "ContactTargetInfo": {
			//               "additionalProperties": false,
			//               "description": "The contact that SSM Incident Manager is engaging during an incident.",
			//               "properties": {
			//                 "ContactId": {
			//                   "description": "The Amazon Resource Name (ARN) of the contact.",
			//                   "type": "string"
			//                 },
			//                 "IsEssential": {
			//                   "description": "A Boolean value determining if the contact's acknowledgement stops the progress of stages in the plan.",
			//                   "type": "boolean"
			//                 }
			//               },
			//               "required": [
			//                 "ContactId",
			//                 "IsEssential"
			//               ],
			//               "type": "object"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "type": "array"
			//       }
			//     },
			//     "required": [
			//       "DurationInMinutes"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "The stages that an escalation plan or engagement plan engages contacts and contact methods in.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"duration_in_minutes": {
						// Property: DurationInMinutes
						Description: "The time to wait until beginning the next stage.",
						Type:        types.Int64Type,
						Required:    true,
					},
					"targets": {
						// Property: Targets
						Description: "The contacts or contact methods that the escalation plan or engagement plan is engaging.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"channel_target_info": {
									// Property: ChannelTargetInfo
									Description: "Information about the contact channel that SSM Incident Manager uses to engage the contact.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"channel_id": {
												// Property: ChannelId
												Description: "The Amazon Resource Name (ARN) of the contact channel.",
												Type:        types.StringType,
												Required:    true,
											},
											"retry_interval_in_minutes": {
												// Property: RetryIntervalInMinutes
												Description: "The number of minutes to wait to retry sending engagement in the case the engagement initially fails.",
												Type:        types.Int64Type,
												Required:    true,
											},
										},
									),
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"contact_target_info": {
									// Property: ContactTargetInfo
									Description: "The contact that SSM Incident Manager is engaging during an incident.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"contact_id": {
												// Property: ContactId
												Description: "The Amazon Resource Name (ARN) of the contact.",
												Type:        types.StringType,
												Required:    true,
											},
											"is_essential": {
												// Property: IsEssential
												Description: "A Boolean value determining if the contact's acknowledgement stops the progress of stages in the plan.",
												Type:        types.BoolType,
												Required:    true,
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
							validate.RequiredAttributes(
								validate.OneOfRequired(
									validate.Required(
										"channel_target_info",
									),
									validate.Required(
										"contact_target_info",
									),
								),
							),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Required: true,
			// Plan is a write-only property.
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			// {
			//   "description": "Contact type, which specify type of contact. Currently supported values: ?PERSONAL?, ?SHARED?, ?OTHER?.",
			//   "enum": [
			//     "PERSONAL",
			//     "CUSTOM",
			//     "SERVICE",
			//     "ESCALATION"
			//   ],
			//   "type": "string"
			// }
			Description: "Contact type, which specify type of contact. Currently supported values: ?PERSONAL?, ?SHARED?, ?OTHER?.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"PERSONAL",
					"CUSTOM",
					"SERVICE",
					"ESCALATION",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			resource.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::SSMContacts::Contact",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SSMContacts::Contact").WithTerraformTypeName("awscc_ssmcontacts_contact")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"alias":                     "Alias",
		"arn":                       "Arn",
		"channel_id":                "ChannelId",
		"channel_target_info":       "ChannelTargetInfo",
		"contact_id":                "ContactId",
		"contact_target_info":       "ContactTargetInfo",
		"display_name":              "DisplayName",
		"duration_in_minutes":       "DurationInMinutes",
		"is_essential":              "IsEssential",
		"plan":                      "Plan",
		"retry_interval_in_minutes": "RetryIntervalInMinutes",
		"targets":                   "Targets",
		"type":                      "Type",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Plan",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
