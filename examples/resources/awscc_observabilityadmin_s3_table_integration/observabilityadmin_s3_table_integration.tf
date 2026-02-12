resource "aws_iam_role" "example_role" {
  name = "example-observability-admin-s3-table-integration-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "observabilityadmin.amazonaws.com"
        }
      }
    ]
  })

  tags = {
    Name        = "example-observability-admin-s3-table-integration-role"
    Environment = "example"
  }
}

resource "aws_iam_policy" "example_policy" {
  name        = "example-observability-admin-s3-table-integration-policy"
  description = "Policy for Observability Admin S3 Table Integration"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "s3:GetObject",
          "s3:PutObject",
          "s3:DeleteObject",
          "s3:ListBucket",
          "s3:GetBucketLocation",
          "s3:CreateBucket"
        ]
        Resource = ["*"]
      },
      {
        Effect = "Allow"
        Action = [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents",
          "logs:DescribeLogGroups",
          "logs:DescribeLogStreams"
        ]
        Resource = ["*"]
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "example_policy" {
  policy_arn = aws_iam_policy.example_policy.arn
  role       = aws_iam_role.example_role.name
}

resource "awscc_observabilityadmin_s3_table_integration" "example" {
  role_arn = aws_iam_role.example_role.arn

  encryption = {
    sse_algorithm = "AES256"
  }

  tags = [
    {
      key   = "Name"
      value = "example-s3-table-integration"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}
