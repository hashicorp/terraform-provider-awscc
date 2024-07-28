resource "awscc_s3outposts_bucket_policy" "example" {
  bucket = awscc_s3outposts_bucket.example.id
  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid      = "AllowAll"
        Action   = "s3-outposts:*"
        Effect   = "Allow"
        Resource = awscc_s3outposts_bucket.example.arn
        Principal = {
          AWS = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
        }
      }
    ]
  })
}
data "aws_caller_identity" "current" {}
