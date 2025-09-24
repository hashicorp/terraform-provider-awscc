resource "awscc_iot_dimension" "example" {
  string_values = ["device/+/auth"]
  type          = "TOPIC_FILTER"
  name          = "example"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
