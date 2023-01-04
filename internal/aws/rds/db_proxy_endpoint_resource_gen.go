// Code generated by generators/resource/main.go; DO NOT EDIT.

package rds

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"regexp"
)

func init() {
	registry.AddResourceFactory("awscc_rds_db_proxy_endpoint", dBProxyEndpointResource)
}

// dBProxyEndpointResource returns the Terraform awscc_rds_db_proxy_endpoint resource.
// This Terraform resource corresponds to the CloudFormation AWS::RDS::DBProxyEndpoint resource.
func dBProxyEndpointResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DBProxyEndpointArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) for the DB proxy endpoint.",
		//	  "pattern": "arn:aws[A-Za-z0-9-]{0,64}:rds:[A-Za-z0-9-]{1,64}:[0-9]{12}:.*",
		//	  "type": "string"
		//	}
		"db_proxy_endpoint_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) for the DB proxy endpoint.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DBProxyEndpointName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier for the DB proxy endpoint. This name must be unique for all DB proxy endpoints owned by your AWS account in the specified AWS Region.",
		//	  "maxLength": 64,
		//	  "pattern": "[0-z]*",
		//	  "type": "string"
		//	}
		"db_proxy_endpoint_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier for the DB proxy endpoint. This name must be unique for all DB proxy endpoints owned by your AWS account in the specified AWS Region.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("[0-z]*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DBProxyName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region.",
		//	  "maxLength": 64,
		//	  "pattern": "[0-z]*",
		//	  "type": "string"
		//	}
		"db_proxy_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("[0-z]*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Endpoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The endpoint that you can use to connect to the DB proxy. You include the endpoint value in the connection string for a database client application.",
		//	  "maxLength": 256,
		//	  "type": "string"
		//	}
		"endpoint": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The endpoint that you can use to connect to the DB proxy. You include the endpoint value in the connection string for a database client application.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IsDefault
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A value that indicates whether this endpoint is the default endpoint for the associated DB proxy. Default DB proxy endpoints always have read/write capability. Other endpoints that you associate with the DB proxy can be either read/write or read-only.",
		//	  "type": "boolean"
		//	}
		"is_default": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "A value that indicates whether this endpoint is the default endpoint for the associated DB proxy. Default DB proxy endpoints always have read/write capability. Other endpoints that you associate with the DB proxy can be either read/write or read-only.",
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An optional set of key-value pairs to associate arbitrary data of your choosing with the DB proxy endpoint.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "pattern": "(\\w|\\d|\\s|\\\\|-|\\.:=+-)*",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 128,
		//	        "pattern": "(\\w|\\d|\\s|\\\\|-|\\.:=+-)*",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthAtMost(128),
							stringvalidator.RegexMatches(regexp.MustCompile("(\\w|\\d|\\s|\\\\|-|\\.:=+-)*"), ""),
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
							stringvalidator.LengthAtMost(128),
							stringvalidator.RegexMatches(regexp.MustCompile("(\\w|\\d|\\s|\\\\|-|\\.:=+-)*"), ""),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An optional set of key-value pairs to associate arbitrary data of your choosing with the DB proxy endpoint.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TargetRole
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A value that indicates whether the DB proxy endpoint can be used for read/write or read-only operations.",
		//	  "enum": [
		//	    "READ_WRITE",
		//	    "READ_ONLY"
		//	  ],
		//	  "type": "string"
		//	}
		"target_role": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A value that indicates whether the DB proxy endpoint can be used for read/write or read-only operations.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"READ_WRITE",
					"READ_ONLY",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VpcId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "VPC ID to associate with the new DB proxy endpoint.",
		//	  "type": "string"
		//	}
		"vpc_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "VPC ID to associate with the new DB proxy endpoint.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VpcSecurityGroupIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "VPC security group IDs to associate with the new DB proxy endpoint.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"vpc_security_group_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "VPC security group IDs to associate with the new DB proxy endpoint.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtLeast(1),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VpcSubnetIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "VPC subnet IDs to associate with the new DB proxy endpoint.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "minItems": 2,
		//	  "type": "array"
		//	}
		"vpc_subnet_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "VPC subnet IDs to associate with the new DB proxy endpoint.",
			Required:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtLeast(2),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				Multiset(),
				listplanmodifier.RequiresReplace(),
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
		Description: "Resource schema for AWS::RDS::DBProxyEndpoint.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::RDS::DBProxyEndpoint").WithTerraformTypeName("awscc_rds_db_proxy_endpoint")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"db_proxy_endpoint_arn":  "DBProxyEndpointArn",
		"db_proxy_endpoint_name": "DBProxyEndpointName",
		"db_proxy_name":          "DBProxyName",
		"endpoint":               "Endpoint",
		"is_default":             "IsDefault",
		"key":                    "Key",
		"tags":                   "Tags",
		"target_role":            "TargetRole",
		"value":                  "Value",
		"vpc_id":                 "VpcId",
		"vpc_security_group_ids": "VpcSecurityGroupIds",
		"vpc_subnet_ids":         "VpcSubnetIds",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}