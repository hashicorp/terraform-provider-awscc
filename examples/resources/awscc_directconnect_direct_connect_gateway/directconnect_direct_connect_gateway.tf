resource "awscc_directconnect_direct_connect_gateway" "example" {
  direct_connect_gateway_name = "example-direct-connect-gateway"
  amazon_side_asn             = "64512"

  tags = [
    {
      key   = "Name"
      value = "example-direct-connect-gateway"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}
