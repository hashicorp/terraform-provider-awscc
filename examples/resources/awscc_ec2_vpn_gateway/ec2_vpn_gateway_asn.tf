resource "awscc_ec2_vpn_gateway" "example_vpn_gateway" {
  type            = "ipsec.1"
  amazon_side_asn = 64512
  tags = [
    {
      key   = "Name"
      value = "Example VPN Gateway"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}