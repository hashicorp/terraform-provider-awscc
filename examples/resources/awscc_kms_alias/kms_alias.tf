resource "awscc_kms_key" "this" {
  key_policy = jsonencode({
    Id = "example_kms_policy"
    Statement = [
      {
        Action = "kms:*"
        Effect = "Allow"
        Principal = {
          AWS = "*"
        }

        Resource = "*"
        Sid      = "Enable IAM User Permissions"
      },
    ]
    Version = "2012-10-17"
  })
}

resource "awscc_kms_alias" "this" {
  alias_name    = "alias/example-kms-alias"
  target_key_id = awscc_kms_key.this.key_id
}