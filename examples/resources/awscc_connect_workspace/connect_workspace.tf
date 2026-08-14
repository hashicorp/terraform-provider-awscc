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

resource "awscc_connect_workspace" "example" {
  instance_arn = awscc_connect_instance.example.arn
  name         = "example-workspace"
  description  = "Example Connect workspace"

  tags = [{
    key   = "Environment"
    value = "test"
  }]
}
