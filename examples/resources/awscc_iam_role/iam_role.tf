resource "awscc_iam_role" "main" {
  role_name   = "sample_iam_role"
  description = "This is a sample IAM role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "ec2.amazonaws.com"
        }
      },
    ]
  })
  max_session_duration = 7200
  path                 = "/"
  tags = [
    {
      key   = "Name"
      value = "Sample IAM Role"
    },
    {
      key   = "Environment"
      value = "Development"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}