# CloudWatch Log Group for Transfer Server
resource "aws_cloudwatch_log_group" "transfer" {
  name              = "/aws/transfer/server"
  retention_in_days = 7
}

# IAM role for Transfer Server logging
data "aws_iam_policy_document" "transfer_assume_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["transfer.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

data "aws_iam_policy_document" "transfer_logging" {
  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogStream",
      "logs:DescribeLogStreams",
      "logs:CreateLogGroup",
      "logs:PutLogEvents"
    ]
    resources = [
      "${aws_cloudwatch_log_group.transfer.arn}:*"
    ]
  }
}

resource "awscc_iam_role" "transfer_logging" {
  role_name                   = "transfer-server-logging-role"
  assume_role_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.transfer_assume_role.json))

  policies = [{
    policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.transfer_logging.json))
    policy_name     = "transfer-server-logging"
  }]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Transfer Server
resource "awscc_transfer_server" "sftp" {
  identity_provider_type = "SERVICE_MANAGED"
  protocols              = ["SFTP"]
  logging_role           = awscc_iam_role.transfer_logging.arn

  endpoint_type = "PUBLIC"

  protocol_details = {
    passive_ip = "0.0.0.0"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}