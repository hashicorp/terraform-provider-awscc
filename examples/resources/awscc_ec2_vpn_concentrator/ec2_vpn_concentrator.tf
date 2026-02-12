resource "awscc_ec2_vpn_concentrator" "example" {
  transit_gateway_id = awscc_ec2_transit_gateway.example.transit_gateway_id
  type              = "ipsec.1"

  tags = [
    {
      key   = "Name"
      value = "example-vpn-concentrator"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}

resource "awscc_ec2_transit_gateway" "example" {
  amazon_side_asn                 = 64512
  auto_accept_shared_attachments  = "disable"
  default_route_table_association = "enable"
  default_route_table_propagation = "enable"
  description                     = "Example AWS Transit Gateway"
  dns_support                     = "enable"
  transit_gateway_cidr_blocks     = ["192.0.2.0/24", "2001:db8:abcd::/64"]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
