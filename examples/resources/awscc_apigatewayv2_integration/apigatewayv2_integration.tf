resource "aws_iam_role" "lambda_role" {
  name = "example-lambda-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Principal = {
        Service = "lambda.amazonaws.com"
      }
    }]
  })
}

resource "aws_lambda_function" "example" {
  function_name = "example-function"
  filename      = "lambda_function.zip"
  handler       = "index.handler"
  runtime       = "nodejs18.x"
  role          = aws_iam_role.lambda_role.arn
}

resource "awscc_apigatewayv2_api" "example" {
  name          = "example-api"
  protocol_type = "HTTP"

  tags = {
    Environment = "dev"
    Name        = "example-api"
  }
}

resource "awscc_apigatewayv2_integration" "example" {
  api_id           = awscc_apigatewayv2_api.example.id
  integration_type = "AWS_PROXY"

  integration_uri        = aws_lambda_function.example.arn
  integration_method     = "POST"
  payload_format_version = "2.0"

  description = "Lambda integration example"
}
