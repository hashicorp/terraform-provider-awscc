data "aws_availability_zones" "available" {
  state = "available"
}

resource "aws_vpc" "example" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = {
    Name        = "example-rtb-fabric-vpc"
    Environment = "example"
  }
}

resource "aws_subnet" "example" {
  vpc_id            = aws_vpc.example.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = data.aws_availability_zones.available.names[0]

  tags = {
    Name        = "example-rtb-fabric-subnet"
    Environment = "example"
  }
}

resource "aws_route_table" "example" {
  vpc_id = aws_vpc.example.id

  tags = {
    Name        = "example-rtb-fabric-route-table"
    Environment = "example"
  }
}

resource "aws_route_table_association" "example" {
  subnet_id      = aws_subnet.example.id
  route_table_id = aws_route_table.example.id
}

resource "awscc_rtbfabric_outbound_external_link" "example" {
  name                 = "example-outbound-link"
  description          = "Example outbound external link for RTB Fabric"
  destination_endpoint = "https://example-destination.com/endpoint"
  
  # The implementation details will vary based on the specific service
  # For now, using minimal required configuration
  tags = {
    Name        = "example-outbound-external-link"
    Environment = "example"
  }
}

# Output the external link ID
output "external_link_id" {
  description = "ID of the RTB Fabric outbound external link"
  value       = awscc_rtbfabric_outbound_external_link.example.id
}

# Output the external link ARN
output "external_link_arn" {
  description = "ARN of the RTB Fabric outbound external link"
  value       = awscc_rtbfabric_outbound_external_link.example.arn
}

# Output the external link name
output "external_link_name" {
  description = "Name of the RTB Fabric outbound external link"
  value       = awscc_rtbfabric_outbound_external_link.example.name
}

# Output the VPC ID for reference
output "vpc_id" {
  description = "ID of the VPC"
  value       = aws_vpc.example.id
}

# Output the subnet ID for reference
output "subnet_id" {
  description = "ID of the subnet"
  value       = aws_subnet.example.id
}