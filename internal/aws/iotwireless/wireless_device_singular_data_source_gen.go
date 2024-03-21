// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iotwireless

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iotwireless_wireless_device", wirelessDeviceDataSource)
}

// wirelessDeviceDataSource returns the Terraform awscc_iotwireless_wireless_device data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoTWireless::WirelessDevice resource.
func wirelessDeviceDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Wireless device arn. Returned after successful create.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Wireless device arn. Returned after successful create.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Wireless device description",
		//	  "maxLength": 2048,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Wireless device description",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DestinationName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Wireless device destination name",
		//	  "maxLength": 128,
		//	  "type": "string"
		//	}
		"destination_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Wireless device destination name",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Wireless device Id. Returned after successful create.",
		//	  "maxLength": 256,
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Wireless device Id. Returned after successful create.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LastUplinkReceivedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The date and time when the most recent uplink was received.",
		//	  "type": "string"
		//	}
		"last_uplink_received_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The date and time when the most recent uplink was received.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LoRaWAN
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The combination of Package, Station and Model which represents the version of the LoRaWAN Wireless Device.",
		//	  "oneOf": [
		//	    {
		//	      "required": [
		//	        "OtaaV11"
		//	      ]
		//	    },
		//	    {
		//	      "required": [
		//	        "OtaaV10x"
		//	      ]
		//	    },
		//	    {
		//	      "required": [
		//	        "AbpV11"
		//	      ]
		//	    },
		//	    {
		//	      "required": [
		//	        "AbpV10x"
		//	      ]
		//	    }
		//	  ],
		//	  "properties": {
		//	    "AbpV10x": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "DevAddr": {
		//	          "pattern": "[a-fA-F0-9]{8}",
		//	          "type": "string"
		//	        },
		//	        "SessionKeys": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "AppSKey": {
		//	              "pattern": "[a-fA-F0-9]{32}",
		//	              "type": "string"
		//	            },
		//	            "NwkSKey": {
		//	              "pattern": "[a-fA-F0-9]{32}",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "NwkSKey",
		//	            "AppSKey"
		//	          ],
		//	          "type": "object"
		//	        }
		//	      },
		//	      "required": [
		//	        "DevAddr",
		//	        "SessionKeys"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "AbpV11": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "DevAddr": {
		//	          "pattern": "[a-fA-F0-9]{8}",
		//	          "type": "string"
		//	        },
		//	        "SessionKeys": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "AppSKey": {
		//	              "pattern": "[a-fA-F0-9]{32}",
		//	              "type": "string"
		//	            },
		//	            "FNwkSIntKey": {
		//	              "pattern": "[a-fA-F0-9]{32}",
		//	              "type": "string"
		//	            },
		//	            "NwkSEncKey": {
		//	              "pattern": "[a-fA-F0-9]{32}",
		//	              "type": "string"
		//	            },
		//	            "SNwkSIntKey": {
		//	              "pattern": "[a-fA-F0-9]{32}",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "FNwkSIntKey",
		//	            "SNwkSIntKey",
		//	            "NwkSEncKey",
		//	            "AppSKey"
		//	          ],
		//	          "type": "object"
		//	        }
		//	      },
		//	      "required": [
		//	        "DevAddr",
		//	        "SessionKeys"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "DevEui": {
		//	      "pattern": "[a-f0-9]{16}",
		//	      "type": "string"
		//	    },
		//	    "DeviceProfileId": {
		//	      "maxLength": 256,
		//	      "type": "string"
		//	    },
		//	    "FPorts": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "Applications": {
		//	          "description": "A list of optional LoRaWAN application information, which can be used for geolocation.",
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "description": "LoRaWAN application configuration, which can be used to perform geolocation.",
		//	            "properties": {
		//	              "DestinationName": {
		//	                "description": "The name of the position data destination that describes the AWS IoT rule that processes the device's position data for use by AWS IoT Core for LoRaWAN.",
		//	                "maxLength": 128,
		//	                "pattern": "[a-zA-Z0-9-_]+",
		//	                "type": "string"
		//	              },
		//	              "FPort": {
		//	                "description": "The Fport value.",
		//	                "maximum": 223,
		//	                "minimum": 1,
		//	                "type": "integer"
		//	              },
		//	              "Type": {
		//	                "description": "Application type, which can be specified to obtain real-time position information of your LoRaWAN device.",
		//	                "enum": [
		//	                  "SemtechGeolocation"
		//	                ],
		//	                "type": "string"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "OtaaV10x": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "AppEui": {
		//	          "pattern": "[a-fA-F0-9]{16}",
		//	          "type": "string"
		//	        },
		//	        "AppKey": {
		//	          "pattern": "[a-fA-F0-9]{32}",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "AppKey",
		//	        "AppEui"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "OtaaV11": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "AppKey": {
		//	          "pattern": "[a-fA-F0-9]{32}",
		//	          "type": "string"
		//	        },
		//	        "JoinEui": {
		//	          "pattern": "[a-fA-F0-9]{16}",
		//	          "type": "string"
		//	        },
		//	        "NwkKey": {
		//	          "pattern": "[a-fA-F0-9]{32}",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "AppKey",
		//	        "NwkKey",
		//	        "JoinEui"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "ServiceProfileId": {
		//	      "maxLength": 256,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"lo_ra_wan": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AbpV10x
				"abp_v10_x": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: DevAddr
						"dev_addr": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: SessionKeys
						"session_keys": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: AppSKey
								"app_s_key": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: NwkSKey
								"nwk_s_key": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: AbpV11
				"abp_v11": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: DevAddr
						"dev_addr": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: SessionKeys
						"session_keys": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: AppSKey
								"app_s_key": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: FNwkSIntKey
								"f_nwk_s_int_key": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: NwkSEncKey
								"nwk_s_enc_key": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: SNwkSIntKey
								"s_nwk_s_int_key": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: DevEui
				"dev_eui": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: DeviceProfileId
				"device_profile_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: FPorts
				"f_ports": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Applications
						"applications": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
							NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: DestinationName
									"destination_name": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "The name of the position data destination that describes the AWS IoT rule that processes the device's position data for use by AWS IoT Core for LoRaWAN.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: FPort
									"f_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
										Description: "The Fport value.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: Type
									"type": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "Application type, which can be specified to obtain real-time position information of your LoRaWAN device.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
							}, /*END NESTED OBJECT*/
							Description: "A list of optional LoRaWAN application information, which can be used for geolocation.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: OtaaV10x
				"otaa_v10_x": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: AppEui
						"app_eui": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: AppKey
						"app_key": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: OtaaV11
				"otaa_v11": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: AppKey
						"app_key": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: JoinEui
						"join_eui": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: NwkKey
						"nwk_key": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: ServiceProfileId
				"service_profile_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The combination of Package, Station and Model which represents the version of the LoRaWAN Wireless Device.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Wireless device name",
		//	  "maxLength": 256,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Wireless device name",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Positioning
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "FPort values for the GNSS, stream, and ClockSync functions of the positioning information.",
		//	  "enum": [
		//	    "Enabled",
		//	    "Disabled"
		//	  ],
		//	  "type": "string"
		//	}
		"positioning": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "FPort values for the GNSS, stream, and ClockSync functions of the positioning information.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of key-value pairs that contain metadata for the device. Currently not supported, will not create if tags are passed.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "maxItems": 200,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of key-value pairs that contain metadata for the device. Currently not supported, will not create if tags are passed.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ThingArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Thing arn. Passed into update to associate Thing with Wireless device.",
		//	  "type": "string"
		//	}
		"thing_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Thing arn. Passed into update to associate Thing with Wireless device.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ThingName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Thing Arn. If there is a Thing created, this can be returned with a Get call.",
		//	  "type": "string"
		//	}
		"thing_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Thing Arn. If there is a Thing created, this can be returned with a Get call.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Wireless device type, currently only Sidewalk and LoRa",
		//	  "enum": [
		//	    "Sidewalk",
		//	    "LoRaWAN"
		//	  ],
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Wireless device type, currently only Sidewalk and LoRa",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IoTWireless::WirelessDevice",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTWireless::WirelessDevice").WithTerraformTypeName("awscc_iotwireless_wireless_device")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"abp_v10_x":               "AbpV10x",
		"abp_v11":                 "AbpV11",
		"app_eui":                 "AppEui",
		"app_key":                 "AppKey",
		"app_s_key":               "AppSKey",
		"applications":            "Applications",
		"arn":                     "Arn",
		"description":             "Description",
		"destination_name":        "DestinationName",
		"dev_addr":                "DevAddr",
		"dev_eui":                 "DevEui",
		"device_profile_id":       "DeviceProfileId",
		"f_nwk_s_int_key":         "FNwkSIntKey",
		"f_port":                  "FPort",
		"f_ports":                 "FPorts",
		"id":                      "Id",
		"join_eui":                "JoinEui",
		"key":                     "Key",
		"last_uplink_received_at": "LastUplinkReceivedAt",
		"lo_ra_wan":               "LoRaWAN",
		"name":                    "Name",
		"nwk_key":                 "NwkKey",
		"nwk_s_enc_key":           "NwkSEncKey",
		"nwk_s_key":               "NwkSKey",
		"otaa_v10_x":              "OtaaV10x",
		"otaa_v11":                "OtaaV11",
		"positioning":             "Positioning",
		"s_nwk_s_int_key":         "SNwkSIntKey",
		"service_profile_id":      "ServiceProfileId",
		"session_keys":            "SessionKeys",
		"tags":                    "Tags",
		"thing_arn":               "ThingArn",
		"thing_name":              "ThingName",
		"type":                    "Type",
		"value":                   "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
