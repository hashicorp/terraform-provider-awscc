# Get the current AWS region and account ID
data "aws_region" "current" {}

data "aws_caller_identity" "current" {}

# Create a Secrets Manager Secret
resource "awscc_secretsmanager_secret" "example" {
  name        = "example-rotation-secret"
  description = "Example secret for rotation schedule"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create a Lambda function role
data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "lambda_permissions" {
  statement {
    effect = "Allow"
    actions = [
      "secretsmanager:GetSecretValue",
      "secretsmanager:DescribeSecret",
      "secretsmanager:PutSecretValue",
      "secretsmanager:UpdateSecretVersionStage"
    ]
    resources = [awscc_secretsmanager_secret.example.id]
  }

  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents"
    ]
    resources = ["arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:log-group:/aws/lambda/*"]
  }
}

resource "awscc_iam_role" "lambda_role" {
  role_name                   = "example-rotation-lambda-role"
  assume_role_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.assume_role.json))
  policies = [{
    policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.lambda_permissions.json))
    policy_name     = "example-rotation-lambda-policy"
  }]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create a Lambda function for rotation
resource "aws_lambda_function" "rotation" {
  filename      = "rotation.zip"
  function_name = "example-rotation-function"
  role          = awscc_iam_role.lambda_role.arn
  handler       = "index.handler"
  runtime       = "python3.9"

  # Dummy code for example purposes
  source_code_hash = filebase64sha256("rotation.zip")
}

# Allow Secrets Manager to invoke the Lambda function
resource "aws_lambda_permission" "allow_secrets_manager" {
  statement_id  = "AllowSecretsManagerInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.rotation.function_name
  principal     = "secretsmanager.amazonaws.com"
}

# Create the rotation schedule
resource "awscc_secretsmanager_rotation_schedule" "example" {
  secret_id           = awscc_secretsmanager_secret.example.id
  rotation_lambda_arn = aws_lambda_function.rotation.arn
  rotation_rules = {
    automatically_after_days = 30
    duration                 = "3h"
  }
  rotate_immediately_on_update = true
}