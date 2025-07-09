data "aws_region" "current" {}

# Create an SNS topic
resource "awscc_sns_topic" "notification" {
  topic_name = "vpc-endpoint-notifications"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# SNS Topic Policy
data "aws_iam_policy_document" "notification" {
  statement {
    effect = "Allow"
    actions = [
      "SNS:Publish"
    ]
    principals {
      type        = "Service"
      identifiers = ["vpce.amazonaws.com"]
    }
    resources = [awscc_sns_topic.notification.topic_arn]
  }
}

resource "aws_sns_topic_policy" "notification" {
  arn    = awscc_sns_topic.notification.topic_arn
  policy = data.aws_iam_policy_document.notification.json
}

resource "awscc_ec2_vpc" "main" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_ec2_vpc_endpoint" "example" {
  service_name      = "com.amazonaws.${data.aws_region.current.name}.s3"
  vpc_id            = awscc_ec2_vpc.main.id
  vpc_endpoint_type = "Interface"
}

resource "awscc_ec2_vpc_endpoint_connection_notification" "example" {
  connection_notification_arn = awscc_sns_topic.notification.topic_arn
  vpc_endpoint_id             = awscc_ec2_vpc_endpoint.example.id
  connection_events           = ["Accept", "Reject"]
}