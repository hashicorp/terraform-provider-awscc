# Get current AWS account ID and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# IAM role for data source
data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["cloudwatch.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "logs_integration" {
  statement {
    effect = "Allow"
    actions = [
      "logs:GetLogRecord",
      "logs:GetLogDelivery",
      "logs:GetLogEvents",
      "logs:DescribeLogStreams",
      "logs:DescribeLogGroups"
    ]
    resources = ["arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:*"]
  }
}

resource "awscc_iam_role" "logs_integration_role" {
  assume_role_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.assume_role.json))
  description                 = "Role for CloudWatch Logs integration"
  role_name                   = "CloudWatchLogsIntegrationRole"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_iam_role_policy" "logs_integration_policy" {
  policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.logs_integration.json))
  policy_name     = "CloudWatchLogsIntegrationPolicy"
  role_name       = awscc_iam_role.logs_integration_role.role_name
}

# Create the logs integration
resource "awscc_logs_integration" "example" {
  integration_name = "example-logs-integration"
  integration_type = "OPENSEARCH"
  resource_config = {
    open_search_resource_config = {
      dashboard_viewer_principals = ["arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"]
      data_source_role_arn        = awscc_iam_role.logs_integration_role.arn
      retention_days              = 7
    }
  }
}