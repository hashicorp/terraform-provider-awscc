resource "awscc_iam_group_policy" "example" {
  group_name  = awscc_iam_group.example.id
  policy_name = "sample_group_policy"

  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
        "s3:ListBucket"]
        Resource = "arn:aws:s3:::my_bucket_name"
      }
    ]
  })
}

resource "awscc_iam_group" "example" {
  group_name = "sample_group"
}