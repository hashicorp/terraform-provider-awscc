# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Create an SNS topic for FMS notifications
resource "awscc_sns_topic" "fms_notifications" {
  topic_name = "fms-notifications-${data.aws_caller_identity.current.account_id}"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Trust policy for FMS service
data "aws_iam_policy_document" "fms_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    effect  = "Allow"

    principals {
      type        = "Service"
      identifiers = ["fms.amazonaws.com"]
    }
  }
}

# Permission policy for the role
data "aws_iam_policy_document" "fms_publish" {
  statement {
    effect = "Allow"
    actions = [
      "sns:Publish"
    ]
    resources = [awscc_sns_topic.fms_notifications.topic_arn]
  }
}

# Create IAM role for FMS
resource "awscc_iam_role" "fms_notification_role" {
  assume_role_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.fms_assume_role.json))
  description                 = "Role for FMS to publish notifications"
  role_name                   = "FMSNotificationRole"
  policies = [{
    policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.fms_publish.json))
    policy_name     = "FMSNotificationPolicy"
  }]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the FMS notification channel
resource "awscc_fms_notification_channel" "example" {
  sns_topic_arn = awscc_sns_topic.fms_notifications.topic_arn
  sns_role_name = awscc_iam_role.fms_notification_role.role_name
}