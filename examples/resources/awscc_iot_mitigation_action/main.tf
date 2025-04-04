# Create IAM role for IoT mitigation action using awscc provider
resource "awscc_iam_role" "iot_mitigation_role" {
  role_name = "iot-mitigation-action-role"

  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "iot.amazonaws.com"
        }
      }
    ]
  })

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create SNS topic for findings
resource "awscc_sns_topic" "findings" {
  topic_name = "iot-findings-topic"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create IAM policy for publishing to SNS
resource "awscc_iam_role_policy" "iot_mitigation_policy" {
  policy_name = "iot-mitigation-sns-policy"
  role_name   = awscc_iam_role.iot_mitigation_role.role_name

  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "sns:Publish"
        ]
        Resource = [awscc_sns_topic.findings.topic_arn]
      }
    ]
  })
}

# Create IoT mitigation action
resource "awscc_iot_mitigation_action" "example" {
  action_name = "publish-finding-to-sns"
  role_arn    = awscc_iam_role.iot_mitigation_role.arn

  action_params = {
    publish_finding_to_sns_params = {
      topic_arn = awscc_sns_topic.findings.topic_arn
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}