# Get current AWS region
data "aws_region" "current" {}

# Create VPC
resource "awscc_ec2_vpc" "example" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create Internet Gateway
resource "awscc_ec2_internet_gateway" "example" {
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Attach Internet Gateway to VPC
resource "awscc_ec2_vpc_gateway_attachment" "example" {
  vpc_id              = awscc_ec2_vpc.example.id
  internet_gateway_id = awscc_ec2_internet_gateway.example.id
}

# Create Subnets
resource "awscc_ec2_subnet" "subnet1" {
  vpc_id            = awscc_ec2_vpc.example.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "${data.aws_region.current.name}a"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_ec2_subnet" "subnet2" {
  vpc_id            = awscc_ec2_vpc.example.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "${data.aws_region.current.name}b"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create Transit Gateway
resource "awscc_ec2_transit_gateway" "example" {
  description = "Example Transit Gateway"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create Transit Gateway VPC Attachment
resource "awscc_ec2_transit_gateway_vpc_attachment" "example" {
  subnet_ids         = [awscc_ec2_subnet.subnet1.id, awscc_ec2_subnet.subnet2.id]
  transit_gateway_id = awscc_ec2_transit_gateway.example.id
  vpc_id             = awscc_ec2_vpc.example.id

  options = {
    dns_support                        = "enable"
    ipv_6_support                      = "disable"
    appliance_mode_support             = "disable"
    security_group_referencing_support = "disable"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}