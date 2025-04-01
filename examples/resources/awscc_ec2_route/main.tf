# Create VPC
resource "awscc_ec2_vpc" "example" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true
  tags = [{
    key   = "Name"
    value = "example-vpc"
  }]
}

# Create Internet Gateway
resource "awscc_ec2_internet_gateway" "example" {
  tags = [{
    key   = "Name"
    value = "example-igw"
  }]
}

# Attach Internet Gateway to VPC
resource "awscc_ec2_vpc_gateway_attachment" "example" {
  vpc_id              = awscc_ec2_vpc.example.id
  internet_gateway_id = awscc_ec2_internet_gateway.example.id
}

# Create Route Table
resource "awscc_ec2_route_table" "example" {
  vpc_id = awscc_ec2_vpc.example.id
  tags = [{
    key   = "Name"
    value = "example-rt"
  }]
}

# Create Route
resource "awscc_ec2_route" "example" {
  route_table_id         = awscc_ec2_route_table.example.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = awscc_ec2_internet_gateway.example.id
}