resource "awscc_iam_role_policy" "example" {
  policy_name = "sample_iam_role_policy"
  role_name   = awscc_iam_role.example.id

  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect   = "Allow"
        Action   = "s3:ListBucket"
        Resource = "arn:aws:s3:::my_bucket_name"
      }
    ]
  })
}

resource "awscc_iam_role" "example" {
  role_name   = "sample_iam_role"
  description = "This is a sample IAM role"
  path        = "/"

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
}