
# Get current AWS region
data "aws_region" "current" {}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# CloudWatch logs role for AppSync
data "aws_iam_policy_document" "appsync_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    effect  = "Allow"

    principals {
      type        = "Service"
      identifiers = ["appsync.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "appsync_logs" {
  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents"
    ]
    resources = [
      "arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:*"
    ]
  }
}

# Create IAM role for AppSync logging
resource "awscc_iam_role" "appsync_logs" {
  role_name                   = "AppSyncLoggingRole"
  assume_role_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.appsync_assume_role.json))
  policies = [
    {
      policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.appsync_logs.json))
      policy_name     = "AppSyncLoggingPolicy"
    }
  ]
}

# Create AppSync GraphQL API
resource "aws_appsync_graphql_api" "example" {
  name                = "example-api"
  authentication_type = "API_KEY"
  xray_enabled        = true

  log_config {
    cloudwatch_logs_role_arn = awscc_iam_role.appsync_logs.arn
    field_log_level          = "ERROR"
    exclude_verbose_content  = false
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}