resource "awscc_ec2_vpc" "vpc" {
  cidr_block       = "10.0.0.0/16"
  instance_tenancy = "default"
  tags = [{
    key   = "Name"
    value = "demovpc"
  }]
}

resource "aws_internet_gateway" "internet_gateway" {
  vpc_id = awscc_ec2_vpc.vpc.id
  tags = {
    Name = "Demo Internet Gateway"
  }
}

resource "aws_route_table" "public_route_table" {
  vpc_id = awscc_ec2_vpc.vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.internet_gateway.id
  }

  tags = {
    Name = "PublicRouteTable"
  }
}

