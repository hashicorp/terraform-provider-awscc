data "aws_region" "current" {}

# VPC
resource "awscc_ec2_vpc" "main" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Subnet
resource "awscc_ec2_subnet" "main" {
  vpc_id            = awscc_ec2_vpc.main.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "${data.aws_region.current.name}a"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Transit Gateway
resource "awscc_ec2_transit_gateway" "main" {
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Route Server Endpoint
resource "awscc_ec2_route_server_endpoint" "main" {
  route_server_id = awscc_ec2_transit_gateway.main.id
  subnet_id       = awscc_ec2_subnet.main.id

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Output the Route Server Endpoint ID
output "route_server_endpoint_id" {
  value = awscc_ec2_route_server_endpoint.main.route_server_endpoint_id
}