// Code generated by generators/resource/main.go; DO NOT EDIT.

package iot

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"regexp"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_iot_logging", loggingResource)
}

// loggingResource returns the Terraform awscc_iot_logging resource.
// This Terraform resource corresponds to the CloudFormation AWS::IoT::Logging resource.
func loggingResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccountId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Your 12-digit account ID (used as the primary identifier for the CloudFormation resource).",
		//	  "maxLength": 12,
		//	  "minLength": 12,
		//	  "pattern": "^[0-9]{12}$",
		//	  "type": "string"
		//	}
		"account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Your 12-digit account ID (used as the primary identifier for the CloudFormation resource).",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(12, 12),
				stringvalidator.RegexMatches(regexp.MustCompile("^[0-9]{12}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DefaultLogLevel
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The log level to use. Valid values are: ERROR, WARN, INFO, DEBUG, or DISABLED.",
		//	  "enum": [
		//	    "ERROR",
		//	    "WARN",
		//	    "INFO",
		//	    "DEBUG",
		//	    "DISABLED"
		//	  ],
		//	  "type": "string"
		//	}
		"default_log_level": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The log level to use. Valid values are: ERROR, WARN, INFO, DEBUG, or DISABLED.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"ERROR",
					"WARN",
					"INFO",
					"DEBUG",
					"DISABLED",
				),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the role that allows IoT to write to Cloudwatch logs.",
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the role that allows IoT to write to Cloudwatch logs.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(20, 2048),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Logging Options enable you to configure your IoT V2 logging role and default logging level so that you can monitor progress events logs as it passes from your devices through Iot core service.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::Logging").WithTerraformTypeName("awscc_iot_logging")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_id":        "AccountId",
		"default_log_level": "DefaultLogLevel",
		"role_arn":          "RoleArn",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}