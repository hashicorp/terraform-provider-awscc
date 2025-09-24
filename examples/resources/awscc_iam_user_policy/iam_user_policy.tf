resource "awscc_iam_user_policy" "example" {
  policy_name = "sample_iam_user_policy"
  user_name   = awscc_iam_user.example.id

  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "s3:ListAllMyBuckets",
          "s3:GetBucketLocation",
        ]
        Effect   = "Allow"
        Resource = "arn:aws:s3:::*"
      },
    ]
  })
}

resource "awscc_iam_user" "example" {
  user_name = "sample_iam_user"
}