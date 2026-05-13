resource "awscc_ec2_vpc" "example" {
  cidr_block       = "10.0.0.0/16"
  instance_tenancy = "default"
  tags = [{
    key   = "Name"
    value = "example-vpc"
  }, {
    key   = "Environment"
    value = "example"
  }]
}

resource "awscc_directconnect_direct_connect_gateway" "example" {
  direct_connect_gateway_name = "example-direct-connect-gateway"
  amazon_side_asn            = "64512"
  
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

resource "awscc_ec2_virtual_private_gateway" "example" {
  type            = "ipsec.1"
  amazon_side_asn = 65000
  
  tags = [
    {
      key   = "Name"
      value = "example-vpn-gateway"
    },
    {
      key   = "Environment" 
      value = "example"
    }
  ]
}

resource "awscc_ec2_vpn_gateway_attachment" "example" {
  vpn_gateway_id = awscc_ec2_virtual_private_gateway.example.vpn_gateway_id
  vpc_id         = awscc_ec2_vpc.example.vpc_id
}

resource "awscc_directconnect_direct_connect_gateway_association" "example" {
  associated_gateway_id          = awscc_ec2_virtual_private_gateway.example.vpn_gateway_id
  direct_connect_gateway_id      = awscc_directconnect_direct_connect_gateway.example.direct_connect_gateway_id
  allowed_prefixes = [
    {
      cidr = "10.0.0.0/16"
    }
  ]

  tags = [
    {
      key   = "Name"
      value = "example-dc-gateway-association"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}
