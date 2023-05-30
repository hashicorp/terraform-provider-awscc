resource "awscc_ec2_vpc" "vpc" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}

resource "aws_internet_gateway" "internet_gateway" {
  vpc_id = awscc_ec2_vpc.vpc.id
}

resource "aws_route_table" "public_route_table" {
  vpc_id = awscc_ec2_vpc.vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.internet_gateway.id
  }
}