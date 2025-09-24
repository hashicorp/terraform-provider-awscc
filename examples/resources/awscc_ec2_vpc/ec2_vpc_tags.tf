resource "awscc_ec2_vpc" "main" {
  cidr_block       = "10.0.0.0/16"
  instance_tenancy = "default"
  tags = [{
    key   = "Name"
    value = "main"
  }]
}