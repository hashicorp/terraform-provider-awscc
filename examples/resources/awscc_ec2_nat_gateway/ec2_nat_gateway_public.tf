resource "awscc_ec2_nat_gateway" "main" {
  subnet_id         = awscc_ec2_subnet.main.subnet_id
  allocation_id     = awscc_ec2_eip.main.allocation_id
  connectivity_type = "public"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}