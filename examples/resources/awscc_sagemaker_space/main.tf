data "aws_region" "current" {}

# First, create a SageMaker domain
data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type = "Service"
      identifiers = [
        "sagemaker.amazonaws.com"
      ]
    }
  }
}

resource "awscc_iam_role" "sagemaker_execution_role" {
  role_name                   = "sagemaker-space-example-role"
  assume_role_policy_document = data.aws_iam_policy_document.assume_role.json
  managed_policy_arns         = ["arn:aws:iam::aws:policy/AmazonSageMakerFullAccess"]
}

# Create a SageMaker domain first
resource "awscc_sagemaker_domain" "example" {
  domain_name = "example-domain"
  auth_mode   = "IAM"
  vpc_id      = aws_default_vpc.default.id
  subnet_ids  = [aws_default_subnet.default.id]

  default_user_settings = {
    execution_role = awscc_iam_role.sagemaker_execution_role.arn
  }

  default_space_settings = {
    execution_role = awscc_iam_role.sagemaker_execution_role.arn
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Use default VPC and subnet for simplicity
resource "aws_default_vpc" "default" {}

resource "aws_default_subnet" "default" {
  availability_zone = "${data.aws_region.current.name}a"
}

# Create a SageMaker space
resource "awscc_sagemaker_space" "example" {
  domain_id  = awscc_sagemaker_domain.example.id
  space_name = "example-space"

  space_display_name = "Example Space"

  space_settings = {
    jupyter_server_app_settings = {
      default_resource_spec = {
        instance_type = "system"
      }
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}