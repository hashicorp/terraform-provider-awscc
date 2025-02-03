# Get current region and account ID
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

# IAM role for Lambda
data "aws_iam_policy_document" "lambda_assume_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

# Create IAM role for Lambda
resource "awscc_iam_role" "lambda_role" {
  assume_role_policy_document = data.aws_iam_policy_document.lambda_assume_role.json
  description                 = "Role for IoT custom authorizer Lambda function"
  role_name                   = "IoTCustomAuthorizerRole"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Basic Lambda execution policy
data "aws_iam_policy_document" "lambda_basic_execution" {
  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents"
    ]
    resources = [
      "arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:log-group:/aws/lambda/*"
    ]
  }
}

# Attach basic execution policy to Lambda role
resource "awscc_iam_role_policy" "lambda_execution" {
  policy_document = data.aws_iam_policy_document.lambda_basic_execution.json
  policy_name     = "lambda-basic-execution"
  role_name       = awscc_iam_role.lambda_role.role_name
}

# Lambda function for IoT custom authorizer
resource "aws_lambda_function" "iot_authorizer" {
  filename      = "authorizer.zip"
  function_name = "IoTCustomAuthorizer"
  role          = awscc_iam_role.lambda_role.arn
  handler       = "index.handler"
  runtime       = "nodejs16.x"
  description   = "IoT custom authorizer function"

  environment {
    variables = {
      IOT_ENDPOINT = "unused"
    }
  }
}

# IoT Custom Authorizer
resource "awscc_iot_authorizer" "example" {
  authorizer_name         = "ExampleIoTAuthorizer"
  authorizer_function_arn = aws_lambda_function.iot_authorizer.arn
  signing_disabled        = true
  status                  = "ACTIVE"
  enable_caching_for_http = false

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}