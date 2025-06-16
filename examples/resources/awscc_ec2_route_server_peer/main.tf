# Data source for region
data "aws_region" "current" {}

# Create VPC
resource "awscc_ec2_vpc" "main" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = [{
    key   = "Name"
    value = "route-server-vpc"
  }, {
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create Subnet
resource "awscc_ec2_subnet" "main" {
  vpc_id            = awscc_ec2_vpc.main.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "${data.aws_region.current.name}a"

  tags = [{
    key   = "Name"
    value = "route-server-subnet"
  }, {
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create an internet gateway
resource "awscc_ec2_internet_gateway" "main" {
  tags = [{
    key   = "Name"
    value = "route-server-igw"
  }, {
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Attach Internet Gateway to VPC
resource "aws_internet_gateway_attachment" "main" {
  internet_gateway_id = awscc_ec2_internet_gateway.main.id
  vpc_id             = awscc_ec2_vpc.main.id
}

# Route Server
resource "awscc_ec2_transit_gateway" "main" {
  description = "Transit gateway for route server"
  tags = [{
    key   = "Name"
    value = "route-server-tgw"
  }, {
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create Route Server endpoint
resource "awscc_ec2_transit_gateway_vpc_attachment" "main" {
  vpc_id             = awscc_ec2_vpc.main.id
  subnet_ids         = [awscc_ec2_subnet.main.id]
  transit_gateway_id = awscc_ec2_transit_gateway.main.id

  tags = [{
    key   = "Name"
    value = "route-server-endpoint"
  }, {
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create Route Server Peer
resource "awscc_ec2_route_server_peer" "example" {
  route_server_endpoint_id = awscc_ec2_transit_gateway_vpc_attachment.main.id
  peer_address             = "10.0.1.100"

  bgp_options = {
    peer_asn                = 65000
    peer_liveness_detection = "bgp-keepalive"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}