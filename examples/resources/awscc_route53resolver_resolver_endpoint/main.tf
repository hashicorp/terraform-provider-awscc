data "aws_region" "current" {}

# Create VPC and Subnets for the resolver endpoint
resource "awscc_ec2_vpc" "main" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_support   = true
  enable_dns_hostnames = true

  tags = [{
    key   = "Name"
    value = "resolver-endpoint-vpc"
  }]
}

resource "awscc_ec2_subnet" "subnet1" {
  vpc_id            = awscc_ec2_vpc.main.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "${data.aws_region.current.name}a"

  tags = [{
    key   = "Name"
    value = "resolver-endpoint-subnet1"
  }]
}

resource "awscc_ec2_subnet" "subnet2" {
  vpc_id            = awscc_ec2_vpc.main.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "${data.aws_region.current.name}b"

  tags = [{
    key   = "Name"
    value = "resolver-endpoint-subnet2"
  }]
}

# Security Group for Resolver Endpoint
resource "awscc_ec2_security_group" "resolver" {
  group_description = "Security group for Route 53 Resolver endpoint"
  vpc_id            = awscc_ec2_vpc.main.id

  security_group_ingress = [
    {
      ip_protocol = "tcp"
      from_port   = 53
      to_port     = 53
      cidr_ip     = "0.0.0.0/0"
    },
    {
      ip_protocol = "udp"
      from_port   = 53
      to_port     = 53
      cidr_ip     = "0.0.0.0/0"
    }
  ]

  security_group_egress = [
    {
      ip_protocol = "-1"
      from_port   = -1
      to_port     = -1
      cidr_ip     = "0.0.0.0/0"
    }
  ]

  tags = [{
    key   = "Name"
    value = "resolver-endpoint-sg"
  }]
}

# Route 53 Resolver Endpoint
resource "awscc_route53resolver_resolver_endpoint" "example" {
  name               = "example-resolver-endpoint"
  direction          = "INBOUND"
  security_group_ids = [awscc_ec2_security_group.resolver.id]

  ip_addresses = [
    {
      subnet_id = awscc_ec2_subnet.subnet1.id
    },
    {
      subnet_id = awscc_ec2_subnet.subnet2.id
    }
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}