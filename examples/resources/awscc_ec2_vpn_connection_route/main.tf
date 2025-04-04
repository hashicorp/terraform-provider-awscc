# VPC
resource "awscc_ec2_vpc" "example" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Name"
    value = "vpn-route-example"
  }]
}

# Virtual Private Gateway
resource "aws_vpn_gateway" "example" {
  vpc_id = awscc_ec2_vpc.example.vpc_id
  tags = {
    Name = "vpn-route-example"
  }
}

# Customer Gateway
resource "awscc_ec2_customer_gateway" "example" {
  bgp_asn    = 65000
  ip_address = "203.0.113.1" # Example IP address
  type       = "ipsec.1"
  tags = [{
    key   = "Name"
    value = "vpn-route-example"
  }]
}

# VPN Connection
resource "aws_vpn_connection" "example" {
  customer_gateway_id = awscc_ec2_customer_gateway.example.id
  type                = "ipsec.1"
  vpn_gateway_id      = aws_vpn_gateway.example.id
  static_routes_only  = true
  tags = {
    Name = "vpn-route-example"
  }
}

# VPN Connection Route
resource "aws_vpn_connection_route" "example" {
  destination_cidr_block = "172.16.0.0/24" # Example customer network CIDR
  vpn_connection_id      = aws_vpn_connection.example.id
}