# Create CloudWatch Log Group for IVS Chat logging
resource "aws_cloudwatch_log_group" "ivs_chat_logs" {
  name              = "/aws/ivschat/example-logs"
  retention_in_days = 7

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Create IAM role for IVS Chat logging
data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    effect  = "Allow"
    principals {
      type        = "Service"
      identifiers = ["ivschat.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "ivs_chat_logging" {
  name               = "IVSChatLoggingRole"
  description        = "IAM role for IVS Chat logging"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Create IAM policy for IVS Chat logging
data "aws_iam_policy_document" "logging_policy" {
  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogStream",
      "logs:PutLogEvents"
    ]
    resources = [
      "${aws_cloudwatch_log_group.ivs_chat_logs.arn}:*"
    ]
  }
}

resource "aws_iam_role_policy" "ivs_chat_logging" {
  name   = "IVSChatLoggingPolicy"
  role   = aws_iam_role.ivs_chat_logging.name
  policy = data.aws_iam_policy_document.logging_policy.json
}

# Create IVS Chat Logging Configuration
resource "awscc_ivschat_logging_configuration" "example" {
  name = "example-ivschat-logging"

  destination_configuration = {
    cloudwatch_logs = {
      log_group_name = aws_cloudwatch_log_group.ivs_chat_logs.name
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

  depends_on = [aws_iam_role_policy.ivs_chat_logging]
}