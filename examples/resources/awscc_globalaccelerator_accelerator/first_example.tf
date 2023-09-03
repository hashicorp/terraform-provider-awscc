resource "awscc_globalaccelerator_accelerator" "example" {
  name            = "Example"
  ip_address_type = "IPV4"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}