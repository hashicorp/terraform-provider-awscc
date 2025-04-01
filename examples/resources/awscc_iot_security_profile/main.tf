# SNS Topic for alerts
resource "awscc_sns_topic" "iot_alerts" {
  topic_name = "iot-security-alerts"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# IAM role for security profile alerts
resource "awscc_iam_role" "iot_security_alerts" {
  role_name = "iot-security-alerts-role"
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
  policies = [{
    policy_name = "IoTSecurityAlertsPolicy"
    policy_document = jsonencode({
      Version = "2012-10-17"
      Statement = [
        {
          Effect = "Allow"
          Action = [
            "sns:Publish"
          ]
          Resource = [awscc_sns_topic.iot_alerts.topic_arn]
        }
      ]
    })
  }]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# IoT Security Profile
resource "awscc_iot_security_profile" "example" {
  security_profile_name        = "example-security-profile"
  security_profile_description = "Example IoT Security Profile"

  behaviors = [
    {
      name   = "num-messages-sent"
      metric = "aws:message-byte-size"
      criteria = {
        consecutive_datapoints_to_alarm = 1
        consecutive_datapoints_to_clear = 1
        comparison_operator             = "greater-than"
        value = {
          count = "50"
        }
      }
    },
    {
      name   = "auth-failures"
      metric = "aws:num-authorization-failures"
      criteria = {
        consecutive_datapoints_to_alarm = 1
        consecutive_datapoints_to_clear = 1
        comparison_operator             = "greater-than"
        duration_seconds                = 300
        value = {
          count = "10"
        }
      }
    }
  ]

  alert_targets = {
    "SNS" = {
      alert_target_arn = awscc_sns_topic.iot_alerts.topic_arn
      role_arn         = awscc_iam_role.iot_security_alerts.arn
    }
  }

  additional_metrics_to_retain_v2 = [
    {
      metric = "aws:message-byte-size"
    }
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}