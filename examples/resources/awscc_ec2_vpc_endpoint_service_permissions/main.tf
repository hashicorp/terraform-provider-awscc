# Get current AWS region
data "aws_region" "current" {}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Create VPC for testing
resource "awscc_ec2_vpc" "example" {
  cidr_block = "10.0.0.0/16"
}

# Create an internet gateway
resource "awscc_ec2_internet_gateway" "example" {
}

# Attach internet gateway to VPC
resource "awscc_ec2_vpc_gateway_attachment" "example" {
  vpc_id              = awscc_ec2_vpc.example.id
  internet_gateway_id = awscc_ec2_internet_gateway.example.id
}

# Create a route table
resource "awscc_ec2_route_table" "example" {
  vpc_id = awscc_ec2_vpc.example.id
}

# Create internet route
resource "awscc_ec2_route" "internet" {
  route_table_id         = awscc_ec2_route_table.example.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = awscc_ec2_internet_gateway.example.id
}

# Create a subnet for the NLB
resource "awscc_ec2_subnet" "example" {
  vpc_id                  = awscc_ec2_vpc.example.id
  cidr_block              = "10.0.1.0/24"
  availability_zone       = "${data.aws_region.current.name}a"
  map_public_ip_on_launch = true
}

# Associate route table with subnet
resource "awscc_ec2_subnet_route_table_association" "example" {
  route_table_id = awscc_ec2_route_table.example.id
  subnet_id      = awscc_ec2_subnet.example.id
}

# Create Network Load Balancer for the endpoint service
resource "awscc_elasticloadbalancingv2_load_balancer" "example" {
  name = "endpoint-service-nlb"
  subnets = [
    awscc_ec2_subnet.example.id
  ]
  type = "network"
}

# Create VPC endpoint service
resource "awscc_ec2_vpc_endpoint_service" "example" {
  network_load_balancer_arns = [awscc_elasticloadbalancingv2_load_balancer.example.id]
  acceptance_required        = false
}

# Create VPC endpoint service permissions
resource "awscc_ec2_vpc_endpoint_service_permissions" "example" {
  service_id = awscc_ec2_vpc_endpoint_service.example.id
  allowed_principals = [
    "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
  ]
}