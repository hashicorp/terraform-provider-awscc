resource "awscc_kms_key" "this" {
  key_policy = jsonencode({
    "Version" : "2012-10-17",
    "Id" : "KMS-Key-Policy",
    "Statement" : [
      {
        "Sid" : "Enable IAM User Permissions",
        "Effect" : "Allow",
        "Principal" : {
          "AWS" : "arn:aws:iam::111122223333:root"
        },
        "Action" : "kms:*",
        "Resource" : "*"
      },
    ],
    }
  )
}

resource "awscc_kms_alias" "this" {
  alias_name    = "alias/example-kms-alias"
  target_key_id = awscc_kms_key.this.key_id
}