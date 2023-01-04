// Code generated by generators/resource/main.go; DO NOT EDIT.

package logs

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
	registry.AddResourceFactory("awscc_logs_query_definition", queryDefinitionResource)
}

// queryDefinitionResource returns the Terraform awscc_logs_query_definition resource.
// This Terraform resource corresponds to the CloudFormation AWS::Logs::QueryDefinition resource.
func queryDefinitionResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: LogGroupNames
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Optionally define specific log groups as part of your query definition",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "description": "LogGroup name",
		//	    "maxLength": 512,
		//	    "minLength": 1,
		//	    "pattern": "[\\.\\-_/#A-Za-z0-9]+",
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"log_group_names": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "Optionally define specific log groups as part of your query definition",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.ValueStringsAre(
					stringvalidator.LengthBetween(1, 512),
					stringvalidator.RegexMatches(regexp.MustCompile("[\\.\\-_/#A-Za-z0-9]+"), ""),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A name for the saved query definition",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "^([^:*\\/]+\\/?)*[^:*\\/]+$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A name for the saved query definition",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 255),
				stringvalidator.RegexMatches(regexp.MustCompile("^([^:*\\/]+\\/?)*[^:*\\/]+$"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: QueryDefinitionId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Unique identifier of a query definition",
		//	  "maxLength": 256,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"query_definition_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Unique identifier of a query definition",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: QueryString
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The query string to use for this definition",
		//	  "maxLength": 10000,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"query_string": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The query string to use for this definition",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 10000),
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
		Description: "The resource schema for AWSLogs QueryDefinition",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Logs::QueryDefinition").WithTerraformTypeName("awscc_logs_query_definition")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"log_group_names":     "LogGroupNames",
		"name":                "Name",
		"query_definition_id": "QueryDefinitionId",
		"query_string":        "QueryString",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}