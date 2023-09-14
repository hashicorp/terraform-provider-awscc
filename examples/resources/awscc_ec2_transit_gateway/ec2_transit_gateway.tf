resource "awscc_ec2_transit_gateway" "example_transit_gateway" {
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