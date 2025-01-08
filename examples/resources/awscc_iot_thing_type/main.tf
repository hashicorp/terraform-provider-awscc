# IoT Thing Type resource
resource "awscc_iot_thing_type" "example" {
  thing_type_name = "example_thing_type"
  thing_type_properties = {
    thing_type_description = "Example IoT Thing Type created with AWSCC provider"
    searchable_attributes  = ["attribute1", "attribute2"]
    mqtt_5_configuration = {
      propagating_attributes = [
        {
          connection_attribute = "iot:ClientId"
          user_property_key    = "example.key"
        }
      ]
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}