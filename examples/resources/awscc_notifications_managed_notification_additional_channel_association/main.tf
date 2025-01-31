data "aws_caller_identity" "current" {}

data "aws_region" "current" {}

# Create an AWS Chatbot Slack channel configuration first
resource "awscc_chatbot_slack_channel_configuration" "example" {
  configuration_name = "example-channel"
  slack_channel_id   = "EXAMPLE123" # Replace with actual Slack channel ID
  slack_workspace_id = "T0123456"   # Replace with actual Slack workspace ID

  iam_role_arn = awscc_iam_role.chatbot_role.arn

  guardrail_policies = [
    "arn:aws:iam::aws:policy/ReadOnlyAccess"
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create IAM role for AWS Chatbot
resource "awscc_iam_role" "chatbot_role" {
  role_name = "AWSChatbotRole"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "chatbot.amazonaws.com"
        }
      }
    ]
  })
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Associate the channel with managed notification
resource "awscc_notifications_managed_notification_additional_channel_association" "example" {
  channel_arn                            = awscc_chatbot_slack_channel_configuration.example.arn
  managed_notification_configuration_arn = "arn:aws:notifications::${data.aws_caller_identity.current.account_id}:managed-notification-configuration/category/AWS-Health/sub-category/Billing"
}