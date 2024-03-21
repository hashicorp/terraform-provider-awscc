// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package connect

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddDataSourceFactory("awscc_connect_user", userDataSource)
}

// userDataSource returns the Terraform awscc_connect_user data source.
// This Terraform data source corresponds to the CloudFormation AWS::Connect::User resource.
func userDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DirectoryUserId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the user account in the directory used for identity management.",
		//	  "type": "string"
		//	}
		"directory_user_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the user account in the directory used for identity management.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: HierarchyGroupArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the hierarchy group for the user.",
		//	  "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/agent-group/[-a-zA-Z0-9]*$",
		//	  "type": "string"
		//	}
		"hierarchy_group_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the hierarchy group for the user.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IdentityInfo
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The information about the identity of the user.",
		//	  "properties": {
		//	    "Email": {
		//	      "description": "The email address. If you are using SAML for identity management and include this parameter, an error is returned.",
		//	      "type": "string"
		//	    },
		//	    "FirstName": {
		//	      "description": "The first name. This is required if you are using Amazon Connect or SAML for identity management.",
		//	      "type": "string"
		//	    },
		//	    "LastName": {
		//	      "description": "The last name. This is required if you are using Amazon Connect or SAML for identity management.",
		//	      "type": "string"
		//	    },
		//	    "Mobile": {
		//	      "description": "The mobile phone number.",
		//	      "pattern": "^\\+[1-9]\\d{1,14}$",
		//	      "type": "string"
		//	    },
		//	    "SecondaryEmail": {
		//	      "description": "The secondary email address. If you provide a secondary email, the user receives email notifications -- other than password reset notifications -- to this email address instead of to their primary email address.",
		//	      "pattern": "",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"identity_info": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Email
				"email": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The email address. If you are using SAML for identity management and include this parameter, an error is returned.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: FirstName
				"first_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The first name. This is required if you are using Amazon Connect or SAML for identity management.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: LastName
				"last_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The last name. This is required if you are using Amazon Connect or SAML for identity management.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Mobile
				"mobile": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The mobile phone number.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SecondaryEmail
				"secondary_email": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The secondary email address. If you provide a secondary email, the user receives email notifications -- other than password reset notifications -- to this email address instead of to their primary email address.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The information about the identity of the user.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InstanceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the Amazon Connect instance.",
		//	  "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$",
		//	  "type": "string"
		//	}
		"instance_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the Amazon Connect instance.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Password
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The password for the user account. A password is required if you are using Amazon Connect for identity management. Otherwise, it is an error to include a password.",
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"password": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The password for the user account. A password is required if you are using Amazon Connect for identity management. Otherwise, it is an error to include a password.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PhoneConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The phone settings for the user.",
		//	  "properties": {
		//	    "AfterContactWorkTimeLimit": {
		//	      "description": "The After Call Work (ACW) timeout setting, in seconds.",
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    },
		//	    "AutoAccept": {
		//	      "description": "The Auto accept setting.",
		//	      "type": "boolean"
		//	    },
		//	    "DeskPhoneNumber": {
		//	      "description": "The phone number for the user's desk phone.",
		//	      "type": "string"
		//	    },
		//	    "PhoneType": {
		//	      "description": "The phone type.",
		//	      "enum": [
		//	        "SOFT_PHONE",
		//	        "DESK_PHONE"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "PhoneType"
		//	  ],
		//	  "type": "object"
		//	}
		"phone_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AfterContactWorkTimeLimit
				"after_contact_work_time_limit": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The After Call Work (ACW) timeout setting, in seconds.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: AutoAccept
				"auto_accept": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "The Auto accept setting.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: DeskPhoneNumber
				"desk_phone_number": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The phone number for the user's desk phone.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: PhoneType
				"phone_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The phone type.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The phone settings for the user.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RoutingProfileArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the routing profile for the user.",
		//	  "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/routing-profile/[-a-zA-Z0-9]*$",
		//	  "type": "string"
		//	}
		"routing_profile_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the routing profile for the user.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SecurityProfileArns
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "One or more security profile arns for the user",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "description": "The identifier of the security profile for the user.",
		//	    "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/security-profile/[-a-zA-Z0-9]*$",
		//	    "type": "string"
		//	  },
		//	  "maxItems": 10,
		//	  "minItems": 1,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"security_profile_arns": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "One or more security profile arns for the user",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "One or more tags.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 256,
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
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "One or more tags.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: UserArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) for the user.",
		//	  "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/agent/[-a-zA-Z0-9]*$",
		//	  "type": "string"
		//	}
		"user_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) for the user.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: UserProficiencies
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "One or more predefined attributes assigned to a user, with a level that indicates how skilled they are.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Proficiency of a user.",
		//	    "properties": {
		//	      "AttributeName": {
		//	        "description": "The name of user's proficiency. You must use name of predefined attribute present in the Amazon Connect instance.",
		//	        "maxLength": 64,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "AttributeValue": {
		//	        "description": "The value of user's proficiency. You must use value of predefined attribute present in the Amazon Connect instance.",
		//	        "maxLength": 64,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Level": {
		//	        "description": "The level of the proficiency. The valid values are 1, 2, 3, 4 and 5.",
		//	        "maximum": 5.0,
		//	        "minimum": 1.0,
		//	        "type": "number"
		//	      }
		//	    },
		//	    "required": [
		//	      "AttributeName",
		//	      "AttributeValue",
		//	      "Level"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"user_proficiencies": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: AttributeName
					"attribute_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The name of user's proficiency. You must use name of predefined attribute present in the Amazon Connect instance.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: AttributeValue
					"attribute_value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value of user's proficiency. You must use value of predefined attribute present in the Amazon Connect instance.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Level
					"level": schema.Float64Attribute{ /*START ATTRIBUTE*/
						Description: "The level of the proficiency. The valid values are 1, 2, 3, 4 and 5.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType:  cctypes.NewMultisetTypeOf[types.Object](ctx),
			Description: "One or more predefined attributes assigned to a user, with a level that indicates how skilled they are.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Username
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The user name for the account.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z0-9\\_\\-\\.\\@]+",
		//	  "type": "string"
		//	}
		"username": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The user name for the account.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Connect::User",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Connect::User").WithTerraformTypeName("awscc_connect_user")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"after_contact_work_time_limit": "AfterContactWorkTimeLimit",
		"attribute_name":                "AttributeName",
		"attribute_value":               "AttributeValue",
		"auto_accept":                   "AutoAccept",
		"desk_phone_number":             "DeskPhoneNumber",
		"directory_user_id":             "DirectoryUserId",
		"email":                         "Email",
		"first_name":                    "FirstName",
		"hierarchy_group_arn":           "HierarchyGroupArn",
		"identity_info":                 "IdentityInfo",
		"instance_arn":                  "InstanceArn",
		"key":                           "Key",
		"last_name":                     "LastName",
		"level":                         "Level",
		"mobile":                        "Mobile",
		"password":                      "Password",
		"phone_config":                  "PhoneConfig",
		"phone_type":                    "PhoneType",
		"routing_profile_arn":           "RoutingProfileArn",
		"secondary_email":               "SecondaryEmail",
		"security_profile_arns":         "SecurityProfileArns",
		"tags":                          "Tags",
		"user_arn":                      "UserArn",
		"user_proficiencies":            "UserProficiencies",
		"username":                      "Username",
		"value":                         "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
