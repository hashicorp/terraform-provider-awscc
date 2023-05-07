resource "awscc_kms_key" "this3" {
  description = "KMS Key for users"
  key_policy = jsonencode({
    "Version" : "2012-10-17",
    "Id" : "KMS-Key-Policy-For-User",
    "Statement" : [
              {
            "Sid": "Enable IAM User Permissions",
            "Effect": "Allow",
            "Principal": {
                "AWS": "arn:aws:iam::111122223333:root"
            },
            "Action": "kms:*",
            "Resource": "*"
        },
      {
        "Sid" : "Allow use of the key",
        "Effect" : "Allow",
        "Principal" : {
          "AWS" : "arn:aws:iam::111122223333:user/ExampleUserRole"
        },
        "Action" : [
          "kms:Encrypt",
          "kms:Decrypt",
          "kms:ReEncrypt*",
          "kms:GenerateDataKey*",
          "kms:DescribeKey"
        ],
        "Resource" : "*"
      },
    ],
    }
  )
}
