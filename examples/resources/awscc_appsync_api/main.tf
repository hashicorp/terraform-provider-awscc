
# Get current AWS region
data "aws_region" "current" {}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# IAM role for CloudWatch logging
resource "awscc_iam_role" "appsync_logs" {
  role_name = "appsync-logs-role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "appsync.amazonaws.com"
        }
      }
    ]
  })

  policies = [
    {
      policy_name = "appsync-cloudwatch-logs"
      policy_document = jsonencode({
        Version = "2012-10-17"
        Statement = [
          {
            Effect = "Allow"
            Action = [
              "logs:CreateLogGroup",
              "logs:CreateLogStream",
              "logs:PutLogEvents"
            ]
            Resource = [
              "arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:log-group:/aws/appsync/*:*"
            ]
          }
        ]
      })
    }
  ]
}

# AppSync API
resource "awscc_appsync_api" "example" {
  name = "example-event-api"

  event_config = {
    log_config = {
      cloudwatch_logs_role_arn = awscc_iam_role.appsync_logs.arn
      log_level                = "INFO"
    }
    auth_providers = [
      {
        auth_type = "API_KEY"
      }
    ]
    connection_auth_modes = [
      {
        auth_type = "API_KEY"
      }
    ]
    default_publish_auth_modes = [
      {
        auth_type = "API_KEY"
      }
    ]
    default_subscribe_auth_modes = [
      {
        auth_type = "API_KEY"
      }
    ]
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Output the AppSync API ID and ARN
output "api_id" {
  value = awscc_appsync_api.example.api_id
}

output "api_arn" {
  value = awscc_appsync_api.example.api_arn
}