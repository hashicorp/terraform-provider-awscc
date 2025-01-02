data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create VPC 1 (Requester)
resource "awscc_ec2_vpc" "vpc1" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Name"
    value = "VPC1-Requester"
  }]
}

# Create VPC 2 (Accepter)
resource "awscc_ec2_vpc" "vpc2" {
  cidr_block = "172.16.0.0/16"
  tags = [{
    key   = "Name"
    value = "VPC2-Accepter"
  }]
}

# Create VPC Peering Connection
resource "awscc_ec2_vpc_peering_connection" "example" {
  vpc_id        = awscc_ec2_vpc.vpc1.id
  peer_vpc_id   = awscc_ec2_vpc.vpc2.id
  peer_owner_id = data.aws_caller_identity.current.account_id
  peer_region   = data.aws_region.current.name
  tags = [{
    key   = "Name"
    value = "VPC-Peering-Example"
  }]
}