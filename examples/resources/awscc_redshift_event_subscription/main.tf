# IAM policy for SNS topic
data "aws_iam_policy_document" "redshift_sns" {
  statement {
    actions = ["sns:Publish"]
    effect  = "Allow"
    principals {
      type = "Service"
      identifiers = [
        "redshift.amazonaws.com"
      ]
    }
    resources = [awscc_sns_topic.redshift_events.topic_arn]
  }
}

# SNS topic for event subscription
resource "awscc_sns_topic" "redshift_events" {
  topic_name = "redshift-events"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "aws_sns_topic_policy" "redshift_events" {
  arn    = awscc_sns_topic.redshift_events.topic_arn
  policy = data.aws_iam_policy_document.redshift_sns.json
}

# Redshift event subscription
resource "awscc_redshift_event_subscription" "example" {
  subscription_name = "redshift-event-sub"
  sns_topic_arn     = awscc_sns_topic.redshift_events.topic_arn
  enabled           = true
  severity          = "INFO"
  source_type       = "cluster"
  event_categories  = ["monitoring"]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}