data "aws_region" "current" {}

# VPC Resources
resource "aws_vpc" "main" {
  cidr_block = "10.0.0.0/16"
  tags = {
    "Modified By" = "AWSCC"
  }
}

resource "aws_subnet" "private" {
  vpc_id            = aws_vpc.main.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "${data.aws_region.current.name}a"
  tags = {
    "Modified By" = "AWSCC"
  }
}

resource "aws_subnet" "private2" {
  vpc_id            = aws_vpc.main.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "${data.aws_region.current.name}b"
  tags = {
    "Modified By" = "AWSCC"
  }
}

# Security Group
resource "aws_security_group" "app_runner" {
  name_prefix = "app-runner-"
  description = "Security group for App Runner VPC connector"
  vpc_id      = aws_vpc.main.id

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    "Modified By" = "AWSCC"
  }
}

# App Runner VPC Connector
resource "awscc_apprunner_vpc_connector" "example" {
  vpc_connector_name = "example-vpc-connector"
  subnets = [
    aws_subnet.private.id,
    aws_subnet.private2.id
  ]
  security_groups = [aws_security_group.app_runner.id]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}