# SNS Topic for successful invocation destination
resource "aws_sns_topic" "example_destination" {
  name = "example-destination"

  tags = {
    Name        = "example-destination"
    Environment = "test"
  }
}

# IAM role for the Lambda function
resource "aws_iam_role" "lambda_exec" {
  name = "example-lambda-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })
}

# IAM policy to allow Lambda to publish to SNS
resource "aws_iam_policy" "lambda_sns" {
  name        = "lambda-sns-policy"
  description = "IAM policy for Lambda to publish to SNS"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "sns:Publish",
          "sns:GetTopicAttributes",
        ]
        Effect   = "Allow"
        Resource = aws_sns_topic.example_destination.arn
      }
    ]
  })
}

# Attach the SNS policy to the Lambda role
resource "aws_iam_role_policy_attachment" "lambda_sns" {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = aws_iam_policy.lambda_sns.arn
}

# Basic Lambda execution policy
resource "aws_iam_role_policy_attachment" "lambda_basic" {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

# Lambda function for the event invoke config
resource "aws_lambda_function" "example" {
  function_name = "example-function"
  role          = aws_iam_role.lambda_exec.arn
  handler       = "index.handler"
  runtime       = "nodejs18.x"
  filename      = "function.zip"

  depends_on = [
    aws_iam_role_policy_attachment.lambda_basic,
    aws_iam_role_policy_attachment.lambda_sns
  ]

  tags = {
    Name        = "example-function"
    Environment = "test"
  }
}

# Event invoke configuration for the Lambda function
resource "awscc_lambda_event_invoke_config" "example" {
  function_name = aws_lambda_function.example.function_name
  qualifier     = "$LATEST"

  maximum_event_age_in_seconds = 3600
  maximum_retry_attempts       = 2

  destination_config = {
    on_success = {
      destination = aws_sns_topic.example_destination.arn
    }
  }

  depends_on = [
    aws_lambda_function.example,
    aws_iam_role_policy_attachment.lambda_sns
  ]
}
