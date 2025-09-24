# Get current AWS region details
data "aws_region" "current" {}

# VPC and subnet for the route server
resource "awscc_ec2_vpc" "main" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_ec2_subnet" "main" {
  vpc_id            = awscc_ec2_vpc.main.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "${data.aws_region.current.name}a"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create Route Server
resource "awscc_ec2_route_server" "example" {
  amazon_side_asn           = 65000
  sns_notifications_enabled = true
  persist_routes            = "enable"
  persist_routes_duration   = 5
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}