data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create a log group
resource "awscc_logs_log_group" "example" {
  log_group_name    = "/aws/example/subscription-filter-demo"
  retention_in_days = 7
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create IAM role for Lambda
data "aws_iam_policy_document" "lambda_assume_role" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

resource "awscc_iam_role" "lambda_role" {
  role_name                   = "SubscriptionFilterLambdaRole"
  assume_role_policy_document = data.aws_iam_policy_document.lambda_assume_role.json
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

data "aws_iam_policy_document" "lambda_policy" {
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

resource "awscc_iam_role_policy" "lambda_policy" {
  policy_name     = "LambdaBasicExecution"
  role_name       = awscc_iam_role.lambda_role.role_name
  policy_document = data.aws_iam_policy_document.lambda_policy.json
}

# Create Lambda function
resource "aws_lambda_function" "processor" {
  filename      = "function.zip"
  function_name = "log-processor"
  role          = awscc_iam_role.lambda_role.arn
  handler       = "index.handler"
  runtime       = "nodejs18.x"
  depends_on    = [awscc_iam_role_policy.lambda_policy]
}

# Create Lambda permission for CloudWatch Logs
resource "aws_lambda_permission" "cloudwatch_logs" {
  statement_id  = "AllowCloudWatchLogs"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.processor.function_name
  principal     = "logs.${data.aws_region.current.name}.amazonaws.com"
  source_arn    = "arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:log-group:/aws/example/subscription-filter-demo:*"
}

# Create subscription filter
resource "awscc_logs_subscription_filter" "example" {
  destination_arn = aws_lambda_function.processor.arn
  filter_pattern  = "[timestamp, requestid, field1, field2, field3, field4, field5, field6, field7]"
  log_group_name  = awscc_logs_log_group.example.log_group_name
  filter_name     = "ExampleSubscriptionFilter"
  depends_on      = [aws_lambda_permission.cloudwatch_logs]
}