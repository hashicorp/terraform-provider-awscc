data "aws_region" "current" {}

# Create a VPC
resource "awscc_ec2_vpc" "main" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create route server subnet
resource "awscc_ec2_subnet" "route_server" {
  vpc_id            = awscc_ec2_vpc.main.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "${data.aws_region.current.name}b"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create a route table
resource "awscc_ec2_route_table" "main" {
  vpc_id = awscc_ec2_vpc.main.id
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create Route Server Security Group
resource "awscc_ec2_security_group" "route_server" {
  group_description = "Security group for Route Server"
  vpc_id            = awscc_ec2_vpc.main.id
  security_group_ingress = [
    {
      cidr_ip     = "0.0.0.0/0"
      from_port   = -1
      ip_protocol = "-1"
      to_port     = -1
      description = "Allow all inbound traffic"
    }
  ]
  security_group_egress = [
    {
      cidr_ip     = "0.0.0.0/0"
      from_port   = -1
      ip_protocol = "-1"
      to_port     = -1
      description = "Allow all outbound traffic"
    }
  ]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create transit gateway
resource "awscc_ec2_transit_gateway" "main" {
  amazon_side_asn = 64512
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create VPC attachment
resource "awscc_ec2_transit_gateway_vpc_attachment" "main" {
  subnet_ids         = [awscc_ec2_subnet.route_server.id]
  transit_gateway_id = awscc_ec2_transit_gateway.main.id
  vpc_id             = awscc_ec2_vpc.main.id
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create transit gateway route table
resource "awscc_ec2_transit_gateway_route_table" "main" {
  transit_gateway_id = awscc_ec2_transit_gateway.main.id
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the route server propagation
resource "awscc_ec2_route_server_propagation" "main" {
  route_server_id = awscc_ec2_transit_gateway.main.id
  route_table_id  = awscc_ec2_route_table.main.id
}