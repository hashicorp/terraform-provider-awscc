# Get current AWS region
data "aws_region" "current" {}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Lambda function for message review
resource "aws_lambda_function" "message_review" {
  filename         = "message_review.zip"
  function_name    = "example-ivs-chat-message-review"
  role             = aws_iam_role.lambda_message_review.arn
  handler          = "index.handler"
  source_code_hash = filebase64sha256("message_review.zip")
  runtime          = "nodejs16.x"
}

# IAM role for the Lambda function
resource "aws_iam_role" "lambda_message_review" {
  name = "example-ivs-chat-message-review-role"

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

# IAM policy for the Lambda function to log to CloudWatch
data "aws_iam_policy_document" "lambda_logging" {
  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents"
    ]
    resources = [
      "arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:log-group:/aws/lambda/${aws_lambda_function.message_review.function_name}:*"
    ]
  }
}

# Attach logging policy to Lambda role
resource "aws_iam_role_policy" "lambda_logging" {
  name   = "lambda-logging"
  role   = aws_iam_role.lambda_message_review.id
  policy = data.aws_iam_policy_document.lambda_logging.json
}

# Add permission for IVS Chat to invoke Lambda function
resource "aws_lambda_permission" "ivschat" {
  statement_id  = "AllowIVSChatInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.message_review.function_name
  principal     = "ivschat.amazonaws.com"
  source_arn    = "arn:aws:ivschat:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:room/*"
}

# Create an IVS Chat room
resource "awscc_ivschat_room" "example" {
  name                            = "example-chat-room"
  maximum_message_length          = 500
  maximum_message_rate_per_second = 5

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

  # Add a message review handler configuration
  message_review_handler = {
    fallback_result = "ALLOW"
    uri             = aws_lambda_function.message_review.arn
  }

  depends_on = [aws_lambda_permission.ivschat]
}