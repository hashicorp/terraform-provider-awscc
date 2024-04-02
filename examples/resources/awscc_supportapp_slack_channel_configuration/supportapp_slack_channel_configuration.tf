resource "aws_iam_role" "AWSSupportSlackAppTFRole" {
  name = "AWSSupportSlackAppTFRole"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Principal = {
          Service = "supportapp.amazonaws.com"
        }
        Action = "sts:AssumeRole"
      }
    ]
  })

  managed_policy_arns = [
    "arn:aws:iam::aws:policy/AWSSupportAppFullAccess"
  ]
}

resource "awscc_supportapp_slack_channel_configuration" "slack_channel_example" {
  team_id                              = "TXXXXXXXXX"
  channel_id                           = "C0XXXXXXXX"
  channel_name                         = "tftemplatechannel1"
  notify_on_create_or_reopen_case      = true
  notify_on_add_correspondence_to_case = false
  notify_on_resolve_case               = true
  notify_on_case_severity              = "high"
  channel_role_arn                     = aws_iam_role.AWSSupportSlackAppTFRole.arn
}