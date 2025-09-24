resource "awscc_ec2_transit_gateway" "example" {}

resource "awscc_ec2_customer_gateway" "example" {
  bgp_asn    = 65000
  ip_address = "198.51.100.1"
  type       = "ipsec.1"
}

resource "awscc_ec2_vpn_connection" "example" {
  customer_gateway_id = awscc_ec2_customer_gateway.example.id
  transit_gateway_id  = awscc_ec2_transit_gateway.example.id
  type                = "ipsec.1"

  vpn_tunnel_options_specifications = [{
    tunnel_inside_cidr = "169.254.10.0/30"
    pre_shared_key     = "example1"
    },
    {
      tunnel_inside_cidr = "169.254.11.0/30"
      pre_shared_key     = "example2"
    }
  ]
}