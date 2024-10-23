// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package neptunegraph

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_neptunegraph_graph", graphResource)
}

// graphResource returns the Terraform awscc_neptunegraph_graph resource.
// This Terraform resource corresponds to the CloudFormation AWS::NeptuneGraph::Graph resource.
func graphResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DeletionProtection
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Value that indicates whether the Graph has deletion protection enabled. The graph can't be deleted when deletion protection is enabled.\n\n_Default_: If not specified, the default value is true.",
		//	  "type": "boolean"
		//	}
		"deletion_protection": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Value that indicates whether the Graph has deletion protection enabled. The graph can't be deleted when deletion protection is enabled.\n\n_Default_: If not specified, the default value is true.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Endpoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The connection endpoint for the graph. For example: `g-12a3bcdef4.us-east-1.neptune-graph.amazonaws.com`",
		//	  "type": "string"
		//	}
		"endpoint": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The connection endpoint for the graph. For example: `g-12a3bcdef4.us-east-1.neptune-graph.amazonaws.com`",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: GraphArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Graph resource ARN",
		//	  "type": "string"
		//	}
		"graph_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Graph resource ARN",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: GraphId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The auto-generated id assigned by the service.",
		//	  "type": "string"
		//	}
		"graph_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The auto-generated id assigned by the service.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: GraphName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Contains a user-supplied name for the Graph. \n\nIf you don't specify a name, we generate a unique Graph Name using a combination of Stack Name and a UUID comprising of 4 characters.\n\n_Important_: If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.",
		//	  "maxLength": 63,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-z][a-zA-Z0-9]*(-[a-zA-Z0-9]+)*$",
		//	  "type": "string"
		//	}
		"graph_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Contains a user-supplied name for the Graph. \n\nIf you don't specify a name, we generate a unique Graph Name using a combination of Stack Name and a UUID comprising of 4 characters.\n\n_Important_: If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 63),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-z][a-zA-Z0-9]*(-[a-zA-Z0-9]+)*$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ProvisionedMemory
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Memory for the Graph.",
		//	  "type": "integer"
		//	}
		"provisioned_memory": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Memory for the Graph.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: PublicConnectivity
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies whether the Graph can be reached over the internet. Access to all graphs requires IAM authentication.\n\nWhen the Graph is publicly reachable, its Domain Name System (DNS) endpoint resolves to the public IP address from the internet.\n\nWhen the Graph isn't publicly reachable, you need to create a PrivateGraphEndpoint in a given VPC to ensure the DNS name resolves to a private IP address that is reachable from the VPC.\n\n_Default_: If not specified, the default value is false.",
		//	  "type": "boolean"
		//	}
		"public_connectivity": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies whether the Graph can be reached over the internet. Access to all graphs requires IAM authentication.\n\nWhen the Graph is publicly reachable, its Domain Name System (DNS) endpoint resolves to the public IP address from the internet.\n\nWhen the Graph isn't publicly reachable, you need to create a PrivateGraphEndpoint in a given VPC to ensure the DNS name resolves to a private IP address that is reachable from the VPC.\n\n_Default_: If not specified, the default value is false.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ReplicaCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the number of replicas you want when finished. All replicas will be provisioned in different availability zones.\n\nReplica Count should always be less than or equal to 2.\n\n_Default_: If not specified, the default value is 1.",
		//	  "type": "integer"
		//	}
		"replica_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Specifies the number of replicas you want when finished. All replicas will be provisioned in different availability zones.\n\nReplica Count should always be less than or equal to 2.\n\n_Default_: If not specified, the default value is 1.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
				int64planmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags associated with this graph.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key"
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
						Optional:    true,
						Computed:    true,
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
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The tags associated with this graph.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeAtMost(50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VectorSearchConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Vector Search Configuration",
		//	  "properties": {
		//	    "VectorSearchDimension": {
		//	      "description": "The vector search dimension",
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "required": [
		//	    "VectorSearchDimension"
		//	  ],
		//	  "type": "object"
		//	}
		"vector_search_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: VectorSearchDimension
				"vector_search_dimension": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The vector search dimension",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						fwvalidators.NotNullInt64(),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
						int64planmodifier.RequiresReplaceIfConfigured(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Vector Search Configuration",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplaceIfConfigured(),
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
		Description: "The AWS::NeptuneGraph::Graph resource creates an Amazon NeptuneGraph Graph.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::NeptuneGraph::Graph").WithTerraformTypeName("awscc_neptunegraph_graph")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"deletion_protection":         "DeletionProtection",
		"endpoint":                    "Endpoint",
		"graph_arn":                   "GraphArn",
		"graph_id":                    "GraphId",
		"graph_name":                  "GraphName",
		"key":                         "Key",
		"provisioned_memory":          "ProvisionedMemory",
		"public_connectivity":         "PublicConnectivity",
		"replica_count":               "ReplicaCount",
		"tags":                        "Tags",
		"value":                       "Value",
		"vector_search_configuration": "VectorSearchConfiguration",
		"vector_search_dimension":     "VectorSearchDimension",
	})

	opts = opts.WithCreateTimeoutInMinutes(2160).WithDeleteTimeoutInMinutes(2160)

	opts = opts.WithUpdateTimeoutInMinutes(2160)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
