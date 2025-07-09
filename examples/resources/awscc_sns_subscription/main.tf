# Get current AWS region
data "aws_region" "current" {}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Create a SNS topic
resource "awscc_sns_topic" "example" {
  topic_name = "example-topic"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create IAM role for Lambda
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

resource "aws_iam_role" "lambda_role" {
  name               = "example-lambda-role"
  assume_role_policy = data.aws_iam_policy_document.lambda_assume_role.json
  tags = {
    "Modified By" = "AWSCC"
  }
}

# Create IAM policy for Lambda logging
data "aws_iam_policy_document" "lambda_logging" {
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

resource "aws_iam_policy" "lambda_logging" {
  name   = "lambda-logging"
  policy = data.aws_iam_policy_document.lambda_logging.json
  tags = {
    "Modified By" = "AWSCC"
  }
}

resource "aws_iam_role_policy_attachment" "lambda_logs" {
  role       = aws_iam_role.lambda_role.name
  policy_arn = aws_iam_policy.lambda_logging.arn
}

resource "aws_lambda_function" "example" {
  filename      = "lambda_function.zip"
  function_name = "example-function"
  role          = aws_iam_role.lambda_role.arn
  handler       = "index.handler"
  runtime       = "nodejs18.x"

  lifecycle {
    ignore_changes = [filename]
  }
}

resource "aws_lambda_permission" "sns" {
  statement_id  = "AllowSNSInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.example.function_name
  principal     = "sns.amazonaws.com"
  source_arn    = "arn:aws:sns:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:${awscc_sns_topic.example.topic_name}"
}

resource "awscc_sns_subscription" "example" {
  protocol  = "lambda"
  topic_arn = "arn:aws:sns:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:${awscc_sns_topic.example.topic_name}"
  endpoint  = aws_lambda_function.example.arn
  filter_policy = jsonencode({
    "event_type" : ["order_placed", "order_cancelled"]
  })
  filter_policy_scope = "MessageAttributes"
  region              = data.aws_region.current.name

  depends_on = [aws_lambda_permission.sns]
}