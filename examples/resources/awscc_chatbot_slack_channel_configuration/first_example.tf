resource "awscc_chatbot_slack_channel_configuration" "example" {
  configuration_name = "example-slack-channel-config"
  iam_role_arn       = awscc_iam_role.example.arn
  slack_channel_id   = var.channel_id
  slack_workspace_id = var.workspace_id
}

resource "awscc_iam_role" "example" {
  role_name = "ChatBot-Channel-Role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "chatbot.amazonaws.com"
        }
      },
    ]
  })
  managed_policy_arns = ["arn:aws:iam::aws:policy/AWSResourceExplorerReadOnlyAccess"]
}

