resource "awscc_backup_backup_vault" "example" {
  backup_vault_name  = "example_backup_vault_kms"
  encryption_key_arn = awscc_kms_key.example.arn
}

resource "awscc_kms_key" "example" {
  description = "KMS Key for root"
  key_policy = jsonencode({
    "Version" : "2012-10-17",
    "Id" : "KMS-Key-Policy-For-Root",
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

resource "awscc_kms_alias" "example" {
  alias_name    = "alias/backup-kms-example"
  target_key_id = awscc_kms_key.example.key_id
}