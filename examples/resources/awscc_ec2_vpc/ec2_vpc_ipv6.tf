resource "awscc_ec2_vpc_cidr_block" "main" {
  amazon_provided_ipv_6_cidr_block = true
  vpc_id                           = awscc_ec2_vpc.selected.id
}

resource "awscc_ec2_vpc" "main" {
  cidr_block = "10.0.0.0/16"
}