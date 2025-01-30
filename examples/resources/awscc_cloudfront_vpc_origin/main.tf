# Get AWS region and account details
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

# Create VPC for the ALB
resource "awscc_ec2_vpc" "example" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Name"
    value = "vpc-origin-example"
  }]
}

# Create Internet Gateway
resource "aws_internet_gateway" "example" {
  vpc_id = awscc_ec2_vpc.example.id
  tags = {
    Name = "vpc-origin-igw"
  }
}

# Create public subnets
resource "awscc_ec2_subnet" "public1" {
  vpc_id                  = awscc_ec2_vpc.example.id
  cidr_block              = "10.0.1.0/24"
  availability_zone       = "${data.aws_region.current.name}a"
  map_public_ip_on_launch = true
  tags = [{
    key   = "Name"
    value = "vpc-origin-public1"
  }]
}

resource "awscc_ec2_subnet" "public2" {
  vpc_id                  = awscc_ec2_vpc.example.id
  cidr_block              = "10.0.2.0/24"
  availability_zone       = "${data.aws_region.current.name}b"
  map_public_ip_on_launch = true
  tags = [{
    key   = "Name"
    value = "vpc-origin-public2"
  }]
}

# Create route table for public subnets
resource "awscc_ec2_route_table" "public" {
  vpc_id = awscc_ec2_vpc.example.id
  tags = [{
    key   = "Name"
    value = "vpc-origin-public-rt"
  }]
}

# Create route to Internet Gateway
resource "aws_route" "public_internet_gateway" {
  route_table_id         = awscc_ec2_route_table.public.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = aws_internet_gateway.example.id
}

# Associate public subnets with public route table
resource "awscc_ec2_subnet_route_table_association" "public1" {
  subnet_id      = awscc_ec2_subnet.public1.id
  route_table_id = awscc_ec2_route_table.public.id
}

resource "awscc_ec2_subnet_route_table_association" "public2" {
  subnet_id      = awscc_ec2_subnet.public2.id
  route_table_id = awscc_ec2_route_table.public.id
}

# Create security group for ALB
resource "awscc_ec2_security_group" "alb" {
  group_description = "Security group for ALB"
  vpc_id            = awscc_ec2_vpc.example.id
  security_group_ingress = [
    {
      ip_protocol = "tcp"
      from_port   = 80
      to_port     = 80
      cidr_ip     = "0.0.0.0/0"
    },
    {
      ip_protocol = "tcp"
      from_port   = 443
      to_port     = 443
      cidr_ip     = "0.0.0.0/0"
    }
  ]
  tags = [{
    key   = "Name"
    value = "vpc-origin-alb-sg"
  }]
}

# Create ALB
resource "awscc_elasticloadbalancingv2_load_balancer" "example" {
  name = "vpc-origin-alb"
  security_groups = [
    awscc_ec2_security_group.alb.id
  ]
  subnets = [
    awscc_ec2_subnet.public1.id,
    awscc_ec2_subnet.public2.id
  ]
  scheme = "internet-facing"
  type   = "application"
  tags = [{
    key   = "Name"
    value = "vpc-origin-alb"
  }]
}

# Create VPC Origin for CloudFront
resource "awscc_cloudfront_vpc_origin" "example" {
  vpc_origin_endpoint_config = {
    arn                    = awscc_elasticloadbalancingv2_load_balancer.example.load_balancer_arn
    name                   = "example-vpc-origin"
    http_port              = 80
    https_port             = 443
    origin_protocol_policy = "https-only"
    origin_ssl_protocols   = ["TLSv1.2"]
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}