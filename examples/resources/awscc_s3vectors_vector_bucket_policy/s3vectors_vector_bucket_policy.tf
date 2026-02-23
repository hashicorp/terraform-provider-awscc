resource "awscc_s3vectors_vector_bucket" "example" {
  vector_bucket_name = "example-vector-bucket"

  tags = [
    {
      key   = "Name"
      value = "example-vector-bucket"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}

resource "awscc_s3vectors_vector_bucket_policy" "example" {
  vector_bucket_name = awscc_s3vectors_vector_bucket.example.vector_bucket_name

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid    = "AllowVectorOperations"
        Effect = "Allow"
        Principal = {
          AWS = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
        }
        Action = [
          "s3vectors:GetVector",
          "s3vectors:PutVector",
          "s3vectors:DeleteVector",
          "s3vectors:ListVectors"
        ]
        Resource = [
          "${awscc_s3vectors_vector_bucket.example.vector_bucket_arn}",
          "${awscc_s3vectors_vector_bucket.example.vector_bucket_arn}/*"
        ]
      }
    ]
  })
}

data "aws_caller_identity" "current" {}
