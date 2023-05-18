resource "awscc_ec2_vpc" "vpc" {
  cidr_block       = "10.0.0.0/16"
  instance_tenancy = "default"
  tags = [{
    key   = "Name"
    value = "demovpc"
  }]
}

resource "awscc_ec2_route_table" "custom_route_table" {
  vpc_id = awscc_ec2_vpc.vpc.id
  tags = [{
    key   = "Name"
    value = "Custom Route Table"
  }]
}
