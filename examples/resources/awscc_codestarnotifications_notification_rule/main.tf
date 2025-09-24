data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create an SNS topic for notifications
resource "awscc_sns_topic" "notifications" {
  topic_name = "codestar-notifications"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create SNS topic policy
data "aws_iam_policy_document" "notification_policy" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["codestar-notifications.amazonaws.com"]
    }
    actions = [
      "SNS:Publish"
    ]
    resources = [awscc_sns_topic.notifications.topic_arn]
  }
}

resource "aws_sns_topic_policy" "default" {
  arn    = awscc_sns_topic.notifications.topic_arn
  policy = data.aws_iam_policy_document.notification_policy.json
}

# Create CodeStar Notification Rule
resource "awscc_codestarnotifications_notification_rule" "example" {
  name        = "example-notification-rule"
  detail_type = "BASIC"

  resource = "arn:aws:codecommit:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:MyDemoRepo"

  event_type_ids = [
    "codecommit-repository-comments-on-commits",
    "codecommit-repository-pull-request-created",
    "codecommit-repository-pull-request-merged"
  ]

  targets = [
    {
      target_type    = "SNS"
      target_address = awscc_sns_topic.notifications.topic_arn
    }
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

  status = "ENABLED"
}