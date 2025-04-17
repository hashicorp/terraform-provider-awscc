# Kinesis stream for logging
resource "aws_kinesis_stream" "logging" {
  name             = "amazon-workspaces-web-logging"
  shard_count      = 1
  retention_period = 24

  tags = {
    Environment = "test"
  }
}

# IAM role that will be used by WorkSpaces Web to write to Kinesis
resource "aws_iam_role" "workspaces_web_logging" {
  name = "workspaces-web-logging"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "workspaces-web.amazonaws.com"
        }
      }
    ]
  })

  tags = {
    "Modified By" = "AWSCC"
  }
}

# IAM policy for the role
resource "aws_iam_role_policy" "workspaces_web_logging" {
  name = "workspaces-web-logging"
  role = aws_iam_role.workspaces_web_logging.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "kinesis:PutRecord",
          "kinesis:PutRecords"
        ]
        Resource = [aws_kinesis_stream.logging.arn]
      }
    ]
  })
}

# Create the user access logging settings
resource "awscc_workspacesweb_user_access_logging_settings" "example" {
  kinesis_stream_arn = aws_kinesis_stream.logging.arn

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

  depends_on = [
    aws_iam_role_policy.workspaces_web_logging
  ]
}

output "user_access_logging_settings_arn" {
  value = awscc_workspacesweb_user_access_logging_settings.example.user_access_logging_settings_arn
}