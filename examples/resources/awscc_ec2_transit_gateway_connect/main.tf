# Get current region details
data "aws_region" "current" {}

# Create a Transit Gateway
resource "awscc_ec2_transit_gateway" "example" {
  description = "Example Transit Gateway for Connect attachment"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create VPC for attachment
resource "awscc_ec2_vpc" "example" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create a subnet
resource "awscc_ec2_subnet" "example" {
  vpc_id            = awscc_ec2_vpc.example.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "${data.aws_region.current.name}a"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create a VPC attachment
resource "awscc_ec2_transit_gateway_vpc_attachment" "example" {
  subnet_ids         = [awscc_ec2_subnet.example.id]
  transit_gateway_id = awscc_ec2_transit_gateway.example.id
  vpc_id             = awscc_ec2_vpc.example.id
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the Transit Gateway Connect attachment
resource "awscc_ec2_transit_gateway_connect" "example" {
  transport_transit_gateway_attachment_id = awscc_ec2_transit_gateway_vpc_attachment.example.id
  options = {
    protocol = "gre"
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}