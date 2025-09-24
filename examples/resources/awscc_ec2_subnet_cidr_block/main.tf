# Create VPC
resource "awscc_ec2_vpc" "example" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true
  instance_tenancy     = "default"

  tags = [{
    key   = "Name"
    value = "example-vpc"
  }]
}

# Create subnet
resource "awscc_ec2_subnet" "example" {
  vpc_id     = awscc_ec2_vpc.example.id
  cidr_block = "10.0.1.0/24"

  tags = [{
    key   = "Name"
    value = "example-subnet"
  }]
}

# Enable IPv6 for VPC
resource "awscc_ec2_vpc_cidr_block" "example" {
  vpc_id                           = awscc_ec2_vpc.example.id
  amazon_provided_ipv_6_cidr_block = true
}

# Associate IPv6 CIDR with subnet
resource "awscc_ec2_subnet_cidr_block" "example" {
  subnet_id        = awscc_ec2_subnet.example.id
  ipv_6_cidr_block = cidrsubnet(awscc_ec2_vpc_cidr_block.example.ipv_6_cidr_block, 8, 1)
}