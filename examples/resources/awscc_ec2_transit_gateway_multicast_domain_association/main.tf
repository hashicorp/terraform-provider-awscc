# Create a VPC for testing
resource "awscc_ec2_vpc" "example" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Name"
    value = "transit-gateway-multicast-example"
  }]
}

# Create a subnet in the VPC
resource "awscc_ec2_subnet" "example" {
  vpc_id     = awscc_ec2_vpc.example.id
  cidr_block = "10.0.1.0/24"
  tags = [{
    key   = "Name"
    value = "transit-gateway-multicast-subnet"
  }]
}

# Create a Transit Gateway
resource "awscc_ec2_transit_gateway" "example" {
  auto_accept_shared_attachments  = "enable"
  default_route_table_association = "enable"
  default_route_table_propagation = "enable"
  multicast_support               = "enable"
  tags = [{
    key   = "Name"
    value = "multicast-transit-gateway"
  }]
}

# Create a Transit Gateway VPC Attachment
resource "awscc_ec2_transit_gateway_vpc_attachment" "example" {
  vpc_id             = awscc_ec2_vpc.example.id
  transit_gateway_id = awscc_ec2_transit_gateway.example.id
  subnet_ids         = [awscc_ec2_subnet.example.id]
  tags = [{
    key   = "Name"
    value = "multicast-vpc-attachment"
  }]
}

# Create a Transit Gateway Multicast Domain
resource "awscc_ec2_transit_gateway_multicast_domain" "example" {
  transit_gateway_id = awscc_ec2_transit_gateway.example.id
  tags = [{
    key   = "Name"
    value = "multicast-domain"
  }]
}

# Create a Transit Gateway Multicast Domain Association
resource "awscc_ec2_transit_gateway_multicast_domain_association" "example" {
  subnet_id                           = awscc_ec2_subnet.example.id
  transit_gateway_attachment_id       = awscc_ec2_transit_gateway_vpc_attachment.example.id
  transit_gateway_multicast_domain_id = awscc_ec2_transit_gateway_multicast_domain.example.id
}