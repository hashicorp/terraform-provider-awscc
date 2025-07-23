data "aws_caller_identity" "current" {}

resource "awscc_ec2_vpc" "main" {
  cidr_block = "10.0.0.0/16"
}

resource "awscc_ec2_subnet" "main" {
  vpc_id            = awscc_ec2_vpc.main.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "us-west-1c"
}

resource "aws_internet_gateway" "ig" {
  vpc_id = awscc_ec2_vpc.main.id
}

resource "aws_lb" "test" {
  name               = "test-lb-tf"
  load_balancer_type = "gateway"
  subnets            = [awscc_ec2_subnet.main.id]
}

resource "aws_vpc_endpoint_service" "example" {
  acceptance_required        = false
  gateway_load_balancer_arns = [aws_lb.test.arn]
}

resource "awscc_ec2_vpc_endpoint" "example" {
  service_name      = aws_vpc_endpoint_service.example.service_name
  vpc_endpoint_type = aws_vpc_endpoint_service.example.service_type
  vpc_id            = awscc_ec2_vpc.main.id
}