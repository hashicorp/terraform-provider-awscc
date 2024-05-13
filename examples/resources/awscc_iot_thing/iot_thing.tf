resource "awscc_iot_thing" "example" {
  thing_name = "example"
  attribute_payload = {
    attributes = {
      name = "examplevalue"
    }
  }
}
