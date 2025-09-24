resource "awscc_ec2_transit_gateway_route_table_propagation" "example" {
  transit_gateway_attachment_id  = awscc_ec2_transit_gateway_attachment.example.id
  transit_gateway_route_table_id = awscc_ec2_transit_gateway_route_table.example.id
}

#Create a transit gateway attachment
resource "awscc_ec2_transit_gateway_attachment" "example" {
  subnet_ids         = [awscc_ec2_subnet.example.id]
  transit_gateway_id = awscc_ec2_transit_gateway.example.id
  vpc_id             = awscc_ec2_vpc.example.id
}

# Create an VPC
resource "awscc_ec2_vpc" "example" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_support   = true
  enable_dns_hostnames = true
}

# Create a public subnet 
resource "awscc_ec2_subnet" "example" {
  vpc_id     = awscc_ec2_vpc.example.id
  cidr_block = "10.0.0.0/24"
}

# Create a transit gateway
resource "awscc_ec2_transit_gateway" "example" {
  amazon_side_asn                 = 64512
  auto_accept_shared_attachments  = "enable"
  default_route_table_association = "disable"
  default_route_table_propagation = "disable"
  dns_support                     = "enable"
  vpn_ecmp_support                = "enable"
}

# Create a transit gateway route table
resource "awscc_ec2_transit_gateway_route_table" "example" {
  transit_gateway_id = awscc_ec2_transit_gateway.example.id
}

