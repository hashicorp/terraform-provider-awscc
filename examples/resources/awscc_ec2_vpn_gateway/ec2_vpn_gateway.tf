resource "awscc_ec2_vpn_gateway" "example_vpn_gateway" {
  type = "ipsec.1"
  tags = [{
    key   = "Name"
    value = "Example VPN Gateway"
  }]

}