resource "awscc_cloudtrail_trail" "example" {
  trail_name                    = "example"
  is_logging                    = true
  s3_bucket_name                = awscc_s3_bucket.example.id
  cloudwatch_logs_log_group_arn = awscc_logs_log_group.example.arn
  cloudwatch_logs_role_arn      = awscc_iam_role.example.arn

  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}

resource "awscc_logs_log_group" "example" {
  log_group_name = "example"
}

resource "awscc_iam_role" "example" {
  role_name = "cloudtrail_logs_role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "cloudtrail.amazonaws.com"
        }
      },
    ]
  })
}
resource "awscc_iam_role_policy" "example" {
  policy_name = "cloudtrail_cloudwatch_logs_policy"
  role_name   = awscc_iam_role.example.role_name
  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ]
        Effect   = "Allow"
        Resource = "arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:log-group:${awscc_logs_log_group.example.id}:log-stream:${data.aws_caller_identity.current.account_id}_CloudTrail_${data.aws_region.current.name}*"
      }
    ]
  })
}

data "aws_caller_identity" "current" {}

data "aws_region" "current" {}
