resource "awscc_vpclattice_resource_gateway" "example" {
  name            = "example-gateway"
  vpc_identifier  = awscc_ec2_vpc.example.id
  subnet_ids      = [awscc_ec2_subnet.example.id]
  ip_address_type = "IPV4"
}