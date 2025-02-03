data "aws_caller_identity" "current" {}

data "aws_region" "current" {}

# VPC and Network Resources
resource "aws_vpc" "emr_studio_vpc" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = {
    Name = "emr-studio-vpc"
  }
}

resource "aws_subnet" "emr_studio_subnet" {
  vpc_id     = aws_vpc.emr_studio_vpc.id
  cidr_block = "10.0.1.0/24"

  tags = {
    Name = "emr-studio-subnet"
  }
}

# Security Groups
resource "aws_security_group" "workspace_sg" {
  name_prefix = "emr-studio-workspace-"
  description = "EMR Studio Workspace Security Group"
  vpc_id      = aws_vpc.emr_studio_vpc.id

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "emr-studio-workspace-sg"
  }
}

resource "aws_security_group" "engine_sg" {
  name_prefix = "emr-studio-engine-"
  description = "EMR Studio Engine Security Group"
  vpc_id      = aws_vpc.emr_studio_vpc.id

  ingress {
    from_port       = 0
    to_port         = 0
    protocol        = "-1"
    security_groups = [aws_security_group.workspace_sg.id]
  }

  tags = {
    Name = "emr-studio-engine-sg"
  }
}

# IAM Role and Policy for EMR Studio Service Role
data "aws_iam_policy_document" "trust_policy" {
  statement {
    actions = ["sts:AssumeRole"]
    effect  = "Allow"
    principals {
      type        = "Service"
      identifiers = ["elasticmapreduce.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "service_policy" {
  statement {
    effect = "Allow"
    actions = [
      "elasticmapreduce:ListInstances",
      "elasticmapreduce:DescribeCluster",
      "elasticmapreduce:ListSteps"
    ]
    resources = ["arn:aws:elasticmapreduce:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:cluster/*"]
  }

  statement {
    effect = "Allow"
    actions = [
      "s3:GetObject",
      "s3:PutObject",
      "s3:ListBucket"
    ]
    resources = [
      "arn:aws:s3:::emr-studio-${data.aws_caller_identity.current.account_id}/*",
      "arn:aws:s3:::emr-studio-${data.aws_caller_identity.current.account_id}"
    ]
  }
}

resource "aws_iam_role" "emr_studio_service_role" {
  name               = "emr-studio-service-role"
  assume_role_policy = data.aws_iam_policy_document.trust_policy.json

  inline_policy {
    name   = "emr-studio-service-policy"
    policy = data.aws_iam_policy_document.service_policy.json
  }

  tags = {
    Modified_By = "AWSCC"
  }
}

# S3 Bucket for EMR Studio
resource "aws_s3_bucket" "emr_studio" {
  bucket = "emr-studio-${data.aws_caller_identity.current.account_id}"

  tags = {
    Modified_By = "AWSCC"
  }
}

# EMR Studio
resource "awscc_emr_studio" "example" {
  name                        = "example-emr-studio"
  auth_mode                   = "IAM"
  vpc_id                      = aws_vpc.emr_studio_vpc.id
  subnet_ids                  = [aws_subnet.emr_studio_subnet.id]
  service_role                = aws_iam_role.emr_studio_service_role.arn
  workspace_security_group_id = aws_security_group.workspace_sg.id
  engine_security_group_id    = aws_security_group.engine_sg.id
  default_s3_location         = "s3://${aws_s3_bucket.emr_studio.id}/notebooks/"
  description                 = "Example EMR Studio created with AWSCC provider"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}