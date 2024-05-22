resource "awscc_globalaccelerator_accelerator" "example" {
  name            = "Example"
  ip_address_type = "IPV4"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_globalaccelerator_listener" "example" {
  accelerator_arn = awscc_globalaccelerator_accelerator.example.id
  protocol        = "TCP"

  port_ranges = [{
    from_port = "80"
    to_port   = "80"
  }]
}

resource "awscc_globalaccelerator_endpoint_group" "example" {
  endpoint_group_region = "eu-west-1"
  listener_arn          = awscc_globalaccelerator_listener.example.id
}
