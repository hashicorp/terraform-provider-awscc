resource "random_id" "instance" {
  byte_length = 4
}

resource "awscc_connect_instance" "example" {
  identity_management_type = "CONNECT_MANAGED"
  instance_alias           = "example-connect-instance-${random_id.instance.hex}"

  attributes = {
    inbound_calls  = true
    outbound_calls = true
  }

  tags = [{
    key   = "Environment"
    value = "test"
  }]
}

resource "awscc_connect_data_table" "example" {
  instance_arn     = awscc_connect_instance.example.arn
  name             = "example-data-table"
  description      = "Example data table for testing"
  status           = "PUBLISHED"
  time_zone        = "UTC"
  value_lock_level = "NONE"

  tags = [{
    key   = "Environment"
    value = "test"
  }]
}
