resource "awscc_ec2_vpc_gateway_attachment" "vpn" {
  vpn_gateway_id = awscc_ec2_vpn_gateway.vpn.id
  vpc_id         = awscc_ec2_vpc.vpc.id
}

resource "awscc_ec2_vpn_gateway" "vpn" {
  type = "ipsec.1"
  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}

resource "awscc_ec2_vpc" "vpc" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}
