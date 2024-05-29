#Creates a transit gateway route table association
resource "awscc_ec2_transit_gateway_route_table_association" "transit_gateway_rt_association_example" {
  transit_gateway_attachment_id  = awscc_ec2_transit_gateway_attachment.transit_gateway_attachment_example.id
  transit_gateway_route_table_id = awscc_ec2_transit_gateway_route_table.transit_gateway_rt_example.id
}

#Create a transit gateway attachment
resource "awscc_ec2_transit_gateway_attachment" "transit_gateway_attachment_example" {
  subnet_ids         = [aws_subnet.example.id]
  transit_gateway_id = awscc_ec2_transit_gateway.transit_gateway_example.id
  vpc_id             = aws_vpc.example.id
}

# Create an VPC
resource "aws_vpc" "example" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_support   = true
  enable_dns_hostnames = true
  tags = {
    Name = "MainVPC"
  }
}

# Create a public subnet 
resource "aws_subnet" "example" {
  vpc_id     = aws_vpc.example.id
  cidr_block = "10.0.0.0/24"
  tags = {
    Name = "PublicSubnet"
  }

  depends_on = [aws_vpc.example]
}

# Create a transit gateway
resource "awscc_ec2_transit_gateway" "transit_gateway_example" {
  amazon_side_asn                 = 64512
  auto_accept_shared_attachments  = "enable"
  default_route_table_association = "disable"
  default_route_table_propagation = "disable"
  dns_support                     = "enable"
  vpn_ecmp_support                = "enable"
}

# Create a transit gateway route table
resource "awscc_ec2_transit_gateway_route_table" "transit_gateway_rt_example" {
  transit_gateway_id = awscc_ec2_transit_gateway.transit_gateway_example.id
}

