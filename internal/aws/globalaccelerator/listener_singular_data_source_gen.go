// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package globalaccelerator

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_globalaccelerator_listener", listenerDataSource)
}

// listenerDataSource returns the Terraform awscc_globalaccelerator_listener data source.
// This Terraform data source corresponds to the CloudFormation AWS::GlobalAccelerator::Listener resource.
func listenerDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AcceleratorArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the accelerator.",
		//	  "type": "string"
		//	}
		"accelerator_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the accelerator.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ClientAffinity
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "NONE",
		//	  "description": "Client affinity lets you direct all requests from a user to the same endpoint.",
		//	  "enum": [
		//	    "NONE",
		//	    "SOURCE_IP"
		//	  ],
		//	  "type": "string"
		//	}
		"client_affinity": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Client affinity lets you direct all requests from a user to the same endpoint.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ListenerArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the listener.",
		//	  "type": "string"
		//	}
		"listener_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the listener.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PortRanges
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A port range to support for connections from  clients to your accelerator.",
		//	    "properties": {
		//	      "FromPort": {
		//	        "description": "A network port number",
		//	        "maximum": 65535,
		//	        "minimum": 0,
		//	        "type": "integer"
		//	      },
		//	      "ToPort": {
		//	        "description": "A network port number",
		//	        "maximum": 65535,
		//	        "minimum": 0,
		//	        "type": "integer"
		//	      }
		//	    },
		//	    "required": [
		//	      "FromPort",
		//	      "ToPort"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"port_ranges": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: FromPort
					"from_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Description: "A network port number",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: ToPort
					"to_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Description: "A network port number",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Protocol
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "TCP",
		//	  "description": "The protocol for the listener.",
		//	  "enum": [
		//	    "TCP",
		//	    "UDP"
		//	  ],
		//	  "type": "string"
		//	}
		"protocol": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The protocol for the listener.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::GlobalAccelerator::Listener",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::GlobalAccelerator::Listener").WithTerraformTypeName("awscc_globalaccelerator_listener")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"accelerator_arn": "AcceleratorArn",
		"client_affinity": "ClientAffinity",
		"from_port":       "FromPort",
		"listener_arn":    "ListenerArn",
		"port_ranges":     "PortRanges",
		"protocol":        "Protocol",
		"to_port":         "ToPort",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}