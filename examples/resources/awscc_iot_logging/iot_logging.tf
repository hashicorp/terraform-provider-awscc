resource "awscc_iot_logging" "example" {
  account_id        = data.aws_caller_identity.current.account_id
  default_log_level = "INFO"
  role_arn          = awscc_iam_role.example.arn
}

resource "awscc_iam_role" "example" {
  role_name   = "example"
  description = "Role that allows IoT to write to Cloudwatch logs"
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
  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}

resource "awscc_iam_role_policy" "example" {
  policy_name = "example"
  role_name   = awscc_iam_role.example.role_name
  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents",
          "logs:PutMetricFilter",
          "logs:PutRetentionPolicy",
          "iot:GetLoggingOptions",
          "iot:SetLoggingOptions",
          "iot:SetV2LoggingOptions",
          "iot:GetV2LoggingOptions",
          "iot:SetV2LoggingLevel",
          "iot:ListV2LoggingLevels",
          "iot:DeleteV2LoggingLevel"
        ]
        Resource = [
          "arn:aws:logs:us-east-1:${data.aws_caller_identity.current.account_id}:log-group:AWSIotLogsV2:*"
        ]
      }
    ]
  })
}

data "aws_caller_identity" "current" {}