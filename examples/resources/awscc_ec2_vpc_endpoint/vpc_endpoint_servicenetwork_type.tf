resource "awscc_ec2_vpc_endpoint" "example" {
  vpc_id              = awscc_ec2_vpc.example.id
  vpc_endpoint_type   = "ServiceNetwork"
  subnet_ids          = [awscc_ec2_subnet.example.id]
  security_group_ids  = [awscc_ec2_security_group.example.id]
  service_network_arn = awscc_vpclattice_service_network.example.arn
  private_dns_enabled = true
}