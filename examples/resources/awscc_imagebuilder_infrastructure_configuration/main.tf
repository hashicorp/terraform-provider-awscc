data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create S3 bucket for logs
resource "awscc_s3_bucket" "image_builder_logs" {
  bucket_name = "image-builder-logs-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"
}

# Create IAM role for Image Builder
data "aws_iam_policy_document" "image_builder_assume_role" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["ec2.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "image_builder_permissions" {
  statement {
    effect = "Allow"
    actions = [
      "s3:PutObject",
      "s3:GetObject",
      "s3:ListBucket"
    ]
    resources = [
      awscc_s3_bucket.image_builder_logs.arn,
      "${awscc_s3_bucket.image_builder_logs.arn}/*"
    ]
  }
  statement {
    effect = "Allow"
    actions = [
      "imagebuilder:*",
      "ec2:CreateTags",
      "ec2:DescribeImages",
      "ec2:RunInstances",
      "ec2:DescribeInstances",
      "ec2:TerminateInstances"
    ]
    resources = ["*"]
  }
}

resource "awscc_iam_role" "image_builder_role" {
  assume_role_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.image_builder_assume_role.json))
  policies = [{
    policy_name     = "ImageBuilderServicePolicy"
    policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.image_builder_permissions.json))
  }]
  role_name = "ImageBuilderServiceRole"
}

# Create instance profile
resource "awscc_iam_instance_profile" "image_builder_profile" {
  instance_profile_name = "ImageBuilderInstanceProfile"
  roles                 = [awscc_iam_role.image_builder_role.role_name]
}

# Create VPC resources
resource "awscc_ec2_vpc" "main" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Name"
    value = "ImageBuilderVPC"
  }]
}

resource "awscc_ec2_subnet" "main" {
  vpc_id     = awscc_ec2_vpc.main.id
  cidr_block = "10.0.1.0/24"
  tags = [{
    key   = "Name"
    value = "ImageBuilderSubnet"
  }]
}

resource "awscc_ec2_security_group" "allow_build" {
  group_name        = "image-builder-sg"
  vpc_id            = awscc_ec2_vpc.main.id
  group_description = "Security group for Image Builder"
  security_group_ingress = [{
    ip_protocol = "-1"
    cidr_ipv4   = "0.0.0.0/0"
    from_port   = -1
    to_port     = -1
  }]
  tags = [{
    key   = "Name"
    value = "ImageBuilderSecurityGroup"
  }]
}

# Create the infrastructure configuration
resource "awscc_imagebuilder_infrastructure_configuration" "example" {
  name                          = "example-infra-config"
  description                   = "Example Image Builder infrastructure configuration"
  instance_profile_name         = awscc_iam_instance_profile.image_builder_profile.instance_profile_name
  instance_types                = ["t3.micro"]
  subnet_id                     = awscc_ec2_subnet.main.id
  security_group_ids            = [awscc_ec2_security_group.allow_build.id]
  terminate_instance_on_failure = true

  logging = {
    s3_logs = {
      s3_bucket_name = awscc_s3_bucket.image_builder_logs.bucket_name
      s3_key_prefix  = "image-builder-logs/"
    }
  }

  instance_metadata_options = {
    http_tokens                 = "required"
    http_put_response_hop_limit = 1
  }

  tags = [{
    key   = "Environment"
    value = "Test"
    }, {
    key   = "Modified By"
    value = "AWSCC"
  }]
}