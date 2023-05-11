resource "awscc_kms_key" "this" {
  description            = "KMS Key for root"
  enabled                = "true"
  enable_key_rotation    = "false"
  pending_window_in_days = 30
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
    },
  )
  tags = [{
    key   = "Name"
    value = "this"
  }]
}
