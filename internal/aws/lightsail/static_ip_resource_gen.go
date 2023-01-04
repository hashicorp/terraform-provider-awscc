// Code generated by generators/resource/main.go; DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_lightsail_static_ip", staticIpResource)
}

// staticIpResource returns the Terraform awscc_lightsail_static_ip resource.
// This Terraform resource corresponds to the CloudFormation AWS::Lightsail::StaticIp resource.
func staticIpResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AttachedTo
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The instance where the static IP is attached.",
		//	  "type": "string"
		//	}
		"attached_to": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The instance where the static IP is attached.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IpAddress
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The static IP address.",
		//	  "type": "string"
		//	}
		"ip_address": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The static IP address.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IsAttached
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A Boolean value indicating whether the static IP is attached.",
		//	  "type": "boolean"
		//	}
		"is_attached": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "A Boolean value indicating whether the static IP is attached.",
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StaticIpArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"static_ip_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StaticIpName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the static IP address.",
		//	  "type": "string"
		//	}
		"static_ip_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the static IP address.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
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
		Description: "Resource Type definition for AWS::Lightsail::StaticIp",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lightsail::StaticIp").WithTerraformTypeName("awscc_lightsail_static_ip")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"attached_to":    "AttachedTo",
		"ip_address":     "IpAddress",
		"is_attached":    "IsAttached",
		"static_ip_arn":  "StaticIpArn",
		"static_ip_name": "StaticIpName",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}