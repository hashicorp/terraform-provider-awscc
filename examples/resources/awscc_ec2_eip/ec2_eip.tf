resource "awscc_ec2_eip" "main" {
  domain = "vpc"
  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}