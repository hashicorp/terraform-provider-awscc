resource "awscc_acmpca_certificate_authority" "example" {
  key_algorithm     = "RSA_4096"
  signing_algorithm = "SHA512WITHRSA"
  type              = "ROOT"
  subject = {
    common_name = "example.com"
  }
  usage_mode = "GENERAL_PURPOSE"
  revocation_configuration = {
    crl_configuration = {
      custom_name        = "crl.example.com"
      enabled            = true
      expiration_in_days = 7
      s3_bucket_name     = awscc_s3_bucket.example.id
      s3_object_acl      = "BUCKET_OWNER_FULL_CONTROL"
    }
  }
  depends_on = [awscc_s3_bucket_policy.example]
}

resource "awscc_s3_bucket_policy" "example" {
  bucket = awscc_s3_bucket.example.id
  policy_document = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Sid    = "ACMPCA CRLs Access",
        Effect = "Allow",
        Principal = {
          Service = "acm-pca.amazonaws.com"
        },
        Action = [
          "s3:GetBucketAcl",
          "s3:GetBucketLocation",
          "s3:PutObject",
          "s3:PutObjectAcl"
        ],
        Resource = [
          "${awscc_s3_bucket.example.arn}",
          "${awscc_s3_bucket.example.arn}/*"
        ]
        Condition = {
          StringEquals = {
            "aws:SourceAccount" = data.aws_caller_identity.current.account_id
          }
        }
      }
    ]
  })
}

resource "awscc_s3_bucket" "example" {
  bucket_name = "example-certificate-revocation-list"
}

data "aws_caller_identity" "current" {}
