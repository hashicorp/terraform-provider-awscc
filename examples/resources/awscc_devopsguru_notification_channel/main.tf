# Get current AWS account ID and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create SNS Topic using AWS provider
resource "aws_sns_topic" "devopsguru" {
  name = "devopsguru-notifications"
  tags = {
    "Modified By" = "AWSCC"
  }
}

# SNS Topic Policy
data "aws_iam_policy_document" "devopsguru_sns" {
  statement {
    actions = [
      "sns:Publish"
    ]
    effect = "Allow"
    principals {
      type = "Service"
      identifiers = [
        "devops-guru.amazonaws.com"
      ]
    }
    resources = [aws_sns_topic.devopsguru.arn]
    condition {
      test     = "StringEquals"
      variable = "aws:SourceAccount"
      values   = [data.aws_caller_identity.current.account_id]
    }
    condition {
      test     = "StringEquals"
      variable = "aws:SourceArn"
      values   = ["arn:aws:devops-guru:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:*"]
    }
  }
}

# Attach policy to SNS topic using AWS provider
resource "aws_sns_topic_policy" "devopsguru" {
  arn    = aws_sns_topic.devopsguru.arn
  policy = data.aws_iam_policy_document.devopsguru_sns.json
}

# Create DevOps Guru notification channel
resource "awscc_devopsguru_notification_channel" "example" {
  config = {
    sns = {
      topic_arn = aws_sns_topic.devopsguru.arn
    }
    filters = {
      severities    = ["HIGH", "MEDIUM"]
      message_types = ["NEW_INSIGHT", "CLOSED_INSIGHT"]
    }
  }
}