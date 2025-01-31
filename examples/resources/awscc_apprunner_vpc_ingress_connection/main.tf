data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# VPC Resources
resource "aws_vpc" "main" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = {
    Name = "app-runner-vpc"
  }
}

resource "awscc_ec2_subnet" "private" {
  vpc_id     = aws_vpc.main.id
  cidr_block = "10.0.1.0/24"

  tags = [{
    key   = "Name"
    value = "app-runner-private"
  }]
}

# VPC Endpoint
resource "awscc_ec2_vpc_endpoint" "apprunner" {
  vpc_id             = aws_vpc.main.id
  service_name       = "com.amazonaws.${data.aws_region.current.name}.apprunner.requests"
  vpc_endpoint_type  = "Interface"
  subnet_ids         = [awscc_ec2_subnet.private.id]
  security_group_ids = [awscc_ec2_security_group.endpoint.id]

  tags = [{
    key   = "Name"
    value = "app-runner-endpoint"
  }]
}

# Security Group
resource "awscc_ec2_security_group" "endpoint" {
  group_name        = "app-runner-endpoint-sg"
  group_description = "Security group for App Runner VPC endpoint"
  vpc_id            = aws_vpc.main.id
  security_group_ingress = [{
    description = "HTTPS from VPC"
    from_port   = 443
    to_port     = 443
    ip_protocol = "tcp"
    cidr_ipv4   = aws_vpc.main.cidr_block
  }]

  tags = [{
    key   = "Name"
    value = "app-runner-endpoint-sg"
  }]
}

# IAM Role for App Runner Service
resource "awscc_iam_role" "apprunner_service" {
  role_name = "app-runner-service-role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "build.apprunner.amazonaws.com"
        }
      }
    ]
  })
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# App Runner Service
resource "awscc_apprunner_service" "example" {
  service_name = "example-service-${formatdate("YYYYMMDD-HHmm", timestamp())}"
  source_configuration = {
    auto_deployments_enabled = false
    image_repository = {
      image_configuration = {
        port = "80"
      }
      image_identifier      = "public.ecr.aws/docker/library/httpd:latest"
      image_repository_type = "ECR_PUBLIC"
    }
  }
  network_configuration = {
    egress_configuration = {
      egress_type = "DEFAULT"
    }
    ingress_configuration = {
      is_publicly_accessible = false
    }
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# App Runner VPC Ingress Connection
resource "awscc_apprunner_vpc_ingress_connection" "example" {
  service_arn                 = awscc_apprunner_service.example.service_arn
  vpc_ingress_connection_name = "example-vpc-ingress"
  ingress_vpc_configuration = {
    vpc_id          = aws_vpc.main.id
    vpc_endpoint_id = awscc_ec2_vpc_endpoint.apprunner.id
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}