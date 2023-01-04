// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_sagemaker_device", deviceDataSource)
}

// deviceDataSource returns the Terraform awscc_sagemaker_device data source.
// This Terraform data source corresponds to the CloudFormation AWS::SageMaker::Device resource.
func deviceDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Device
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The Edge Device you want to register against a device fleet",
		//	  "properties": {
		//	    "Description": {
		//	      "description": "Description of the device",
		//	      "maxLength": 40,
		//	      "minLength": 1,
		//	      "pattern": "[\\S\\s]+",
		//	      "type": "string"
		//	    },
		//	    "DeviceName": {
		//	      "description": "The name of the device",
		//	      "maxLength": 63,
		//	      "minLength": 1,
		//	      "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9])*$",
		//	      "type": "string"
		//	    },
		//	    "IotThingName": {
		//	      "description": "AWS Internet of Things (IoT) object name.",
		//	      "maxLength": 128,
		//	      "pattern": "[a-zA-Z0-9:_-]+",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "DeviceName"
		//	  ],
		//	  "type": "object"
		//	}
		"device": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Description
				"description": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Description of the device",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: DeviceName
				"device_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name of the device",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: IotThingName
				"iot_thing_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "AWS Internet of Things (IoT) object name.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The Edge Device you want to register against a device fleet",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DeviceFleetName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the edge device fleet",
		//	  "maxLength": 63,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9](-*_*[a-zA-Z0-9])*$",
		//	  "type": "string"
		//	}
		"device_fleet_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the edge device fleet",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Associate tags with the resource",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The key value of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key value of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Associate tags with the resource",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::SageMaker::Device",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::Device").WithTerraformTypeName("awscc_sagemaker_device")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"description":       "Description",
		"device":            "Device",
		"device_fleet_name": "DeviceFleetName",
		"device_name":       "DeviceName",
		"iot_thing_name":    "IotThingName",
		"key":               "Key",
		"tags":              "Tags",
		"value":             "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}