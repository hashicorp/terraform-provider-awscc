resource "awscc_kendra_index" "example" {
  edition     = "ENTERPRISE_EDITION"
  name        = "example-index"
  role_arn    = awscc_iam_role.example.arn
  description = "Example Kendra index"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}


resource "awscc_iam_role" "example" {
  role_name   = "kendra_index_role"
  description = "Role assigned to the Kendra index"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "kendra.amazonaws.com"
        }
      }
    ]
  })
  max_session_duration = 7200
  tags = [
    {
      key   = "Name"
      value = "Kendra index role"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}

resource "awscc_iam_role_policy" "example" {
  policy_name = "kendra_role_policy"
  role_name   = awscc_iam_role.example.id

  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect   = "Allow"
        Action   = "cloudwatch:PutMetricData"
        Resource = "*"
        Condition = {
          "StringEquals" : {
            "cloudwatch:namespace" : "AWS/Kendra"
          }
        }
      },
      {
        Effect   = "Allow"
        Action   = "logs:DescribeLogGroups"
        Resource = "*"
      },
      {
        Effect   = "Allow"
        Action   = "logs:CreateLogGroup",
        Resource = "arn:aws:logs:us-east-1:${data.aws_caller_identity.current.account_id}:log-group:/aws/kendra/*"
      },
      {
        Effect = "Allow"
        Action = [
          "logs:DescribeLogStreams",
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ],
        Resource = "arn:aws:logs:us-east-1:${data.aws_caller_identity.current.account_id}:log-group:/aws/kendra/*:log-stream:*"
      }
    ]
  })
}

data "aws_caller_identity" "current" {}
