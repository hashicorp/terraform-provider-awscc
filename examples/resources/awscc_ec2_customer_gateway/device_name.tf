resource "awscc_ec2_customer_gateway" "example" {
  bgp_asn     = 65000
  ip_address  = "12.1.2.3"
  type        = "ipsec.1"
  device_name = "example"
}