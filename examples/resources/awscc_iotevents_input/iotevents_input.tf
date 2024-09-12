resource "awscc_iotevents_input" "example" {
  input_name = "example"
  input_definition = {
    attributes = [{
      json_path = "payload.transformed_payload.ph_content"
      },
      {
        json_path = "payload.transformed_payload.device_id"
      }
    ]
  }
  input_description = "An example input"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

