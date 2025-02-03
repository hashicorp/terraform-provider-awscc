# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Get current AWS region
data "aws_region" "current" {}

# IAM role for API Gateway CloudWatch logging
resource "awscc_iam_role" "apigw_cloudwatch" {
  role_name = "APIGatewayCloudWatchLogs"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "apigateway.amazonaws.com"
        }
      }
    ]
  })
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# IAM role policy for CloudWatch logs
resource "awscc_iam_role_policy" "apigw_cloudwatch" {
  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:DescribeLogGroups",
          "logs:DescribeLogStreams",
          "logs:PutLogEvents",
          "logs:GetLogEvents",
          "logs:FilterLogEvents"
        ]
        Resource = "arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:*"
      }
    ]
  })
  policy_name = "APIGatewayCloudWatchLogsPolicy"
  role_name   = awscc_iam_role.apigw_cloudwatch.role_name
}

# API Gateway Account settings
resource "awscc_apigateway_account" "demo" {
  cloudwatch_role_arn = awscc_iam_role.apigw_cloudwatch.arn

  depends_on = [awscc_iam_role_policy.apigw_cloudwatch]
}