# SNS Topic for notifications
resource "awscc_sns_topic" "incident_notifications" {
  topic_name = "incident-notifications"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# IAM Role for SSM Automation
resource "awscc_iam_role" "ssm_automation" {
  role_name = "ssm-automation-role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Effect = "Allow"
      Principal = {
        Service = "ssm.amazonaws.com"
      }
      Action = "sts:AssumeRole"
    }]
  })

  policies = [{
    policy_document = jsonencode({
      Version = "2012-10-17"
      Statement = [{
        Effect = "Allow"
        Action = [
          "ssm:StartAutomationExecution",
          "ssm:GetAutomationExecution"
        ]
        Resource = ["*"]
      }]
    })
    policy_name = "SSMAutomationPolicy"
  }]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Response Plan
resource "awscc_ssmincidents_response_plan" "example" {
  name         = "example-response-plan"
  display_name = "Example Response Plan"

  incident_template = {
    impact  = 3
    title   = "Example Incident"
    summary = "This is an example incident response plan"
    notification_targets = [{
      sns_topic_arn = awscc_sns_topic.incident_notifications.arn
    }]
    incident_tags = [{
      key   = "Environment"
      value = "Production"
    }]
  }

  actions = [{
    ssm_automation = {
      document_name    = "AWS-RestartEC2Instance"
      document_version = "1"
      role_arn         = awscc_iam_role.ssm_automation.arn
      target_account   = "RESPONSE_PLAN_OWNER_ACCOUNT"
      parameters = [{
        key    = "AutomationAssumeRole"
        values = [awscc_iam_role.ssm_automation.arn]
      }]
    }
  }]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
