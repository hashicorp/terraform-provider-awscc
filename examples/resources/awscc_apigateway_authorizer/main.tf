data "aws_region" "current" {}

# IAM role for Lambda function
data "aws_iam_policy_document" "lambda_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    effect  = "Allow"
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

resource "awscc_iam_role" "lambda_role" {
  role_name                   = "example-lambda-authorizer-role"
  assume_role_policy_document = data.aws_iam_policy_document.lambda_assume_role.json
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# IAM role for API Gateway authorizer
data "aws_iam_policy_document" "authorizer_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    effect  = "Allow"
    principals {
      type        = "Service"
      identifiers = ["apigateway.amazonaws.com"]
    }
  }
}

# Policy to allow API Gateway to invoke Lambda
data "aws_iam_policy_document" "invoke_lambda" {
  statement {
    actions   = ["lambda:InvokeFunction"]
    effect    = "Allow"
    resources = [aws_lambda_function.authorizer.arn]
  }
}

resource "awscc_iam_role" "authorizer_role" {
  role_name                   = "example-apigateway-authorizer-role"
  assume_role_policy_document = data.aws_iam_policy_document.authorizer_assume_role.json
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_iam_role_policy" "authorizer_lambda_invoke" {
  policy_name     = "lambda-invoke"
  policy_document = data.aws_iam_policy_document.invoke_lambda.json
  role_name       = awscc_iam_role.authorizer_role.role_name
}

# Lambda function for the authorizer
resource "aws_lambda_function" "authorizer" {
  filename      = "lambda_function.zip"
  function_name = "example-lambda-authorizer"
  role          = awscc_iam_role.lambda_role.arn
  handler       = "index.handler"
  runtime       = "nodejs18.x"

  source_code_hash = filebase64sha256("lambda_function.zip")
}

# API Gateway REST API
resource "awscc_apigateway_rest_api" "example" {
  name = "example-api"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# API Gateway Authorizer
resource "awscc_apigateway_authorizer" "example" {
  name                   = "example-authorizer"
  rest_api_id            = awscc_apigateway_rest_api.example.id
  type                   = "TOKEN"
  authorizer_uri         = "arn:aws:apigateway:${data.aws_region.current.name}:lambda:path/2015-03-31/functions/${aws_lambda_function.authorizer.arn}/invocations"
  identity_source        = "method.request.header.Authorization"
  authorizer_credentials = awscc_iam_role.authorizer_role.arn
}