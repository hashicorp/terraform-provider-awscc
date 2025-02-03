# Get current AWS region
data "aws_region" "current" {}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Create a Cognito User Pool
resource "aws_cognito_user_pool" "example" {
  name = "example-user-pool"

  admin_create_user_config {
    allow_admin_create_user_only = false
  }

  auto_verified_attributes = ["email"]

  password_policy {
    minimum_length    = 8
    require_lowercase = true
    require_numbers   = true
    require_symbols   = true
    require_uppercase = true
  }

  schema {
    attribute_data_type = "String"
    mutable             = true
    name                = "email"
    required            = true

    string_attribute_constraints {
      min_length = 0
      max_length = 2048
    }
  }
}

# Create IAM role for the group
data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"
    principals {
      type = "Federated"
      identifiers = [
        "cognito-identity.amazonaws.com"
      ]
    }
    actions = ["sts:AssumeRoleWithWebIdentity"]
    condition {
      test     = "StringEquals"
      variable = "cognito-identity.amazonaws.com:aud"
      values   = ["${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:identitypool/example"]
    }
  }
}

# Create IAM role using AWSCC provider
resource "awscc_iam_role" "group_role" {
  role_name                   = "cognito-group-role"
  assume_role_policy_document = data.aws_iam_policy_document.assume_role.json
  managed_policy_arns         = ["arn:aws:iam::aws:policy/AWSLambdaExecute"]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create Cognito User Pool Group
resource "awscc_cognito_user_pool_group" "example" {
  user_pool_id = aws_cognito_user_pool.example.id
  group_name   = "example-group"
  description  = "Example Cognito User Pool Group"
  precedence   = 1
  role_arn     = awscc_iam_role.group_role.arn
}