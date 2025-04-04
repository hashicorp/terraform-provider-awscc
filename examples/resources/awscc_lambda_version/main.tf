data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }

    actions = ["sts:AssumeRole"]
  }
}

# Required IAM role for Lambda
resource "awscc_iam_role" "lambda" {
  role_name                   = "example-lambda-version-role"
  assume_role_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.assume_role.json))
  managed_policy_arns         = ["arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Lambda function
resource "aws_lambda_function" "example" {
  filename         = "function.zip"
  function_name    = "example-lambda-version"
  role             = awscc_iam_role.lambda.arn
  handler          = "index.handler"
  runtime          = "nodejs18.x"
  source_code_hash = filebase64sha256("function.zip")

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Lambda version
resource "awscc_lambda_version" "example" {
  function_name = aws_lambda_function.example.function_name
  description   = "Example version created by AWSCC provider"
}