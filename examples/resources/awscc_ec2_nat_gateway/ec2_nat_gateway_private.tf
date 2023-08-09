resource "awscc_ec2_nat_gateway" "main" {
  subnet_id         = awscc_ec2_subnet.main.subnet_id
  connectivity_type = "private"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}