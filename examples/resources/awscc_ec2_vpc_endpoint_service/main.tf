# Note: Using data.aws_region.current.region (AWS provider v6.0+)
# For AWS provider < v6.0, use data.aws_region.current.name instead
data "aws_region" "current" {}

# Create VPC and Subnet for the NLB
resource "awscc_ec2_vpc" "example" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_ec2_subnet" "example" {
  vpc_id            = awscc_ec2_vpc.example.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "${data.aws_region.current.region}a"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create Network Load Balancer
resource "awscc_elasticloadbalancingv2_load_balancer" "example" {
  name               = "example-nlb"
  scheme            = "internal"
  type              = "network"
  subnets           = [awscc_ec2_subnet.example.id]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create VPC Endpoint Service
resource "awscc_ec2_vpc_endpoint_service" "example" {
  acceptance_required = true
  network_load_balancer_arns = [
    awscc_elasticloadbalancingv2_load_balancer.example.id
  ]
}