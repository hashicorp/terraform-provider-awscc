provider "awscc" {
  region = "us-east-1"
}

provider "awscc" {
  region = "us-east-2"
  alias  = "secondary"
}

resource "awscc_kms_replica_key" "example" {
  provider = awscc.secondary

  primary_key_arn        = awscc_kms_key.example.arn
  description            = "Example KMS replica key"
  enabled                = true
  pending_window_in_days = 7
  key_policy = jsonencode(
    {
      "Version" : "2012-10-17",
      "Id" : "key_policy",
      "Statement" : [
        {
          "Sid" : "Enable IAM User Permissions",
          "Effect" : "Allow",
          "Principal" : {
            "AWS" : "arn:aws:iam::0123456789012:root"
          },
          "Action" : "kms:*",
          "Resource" : "*"
        },
        {
          "Sid" : "Allow administration of the key",
          "Effect" : "Allow",
          "Principal" : {
            "AWS" : "arn:aws:iam::0123456789012:role/Admin"
          },
          "Action" : [
            "kms:Create*",
            "kms:Delete*",
            "kms:Disable*",
            "kms:Describe*",
            "kms:Enable*",
            "kms:Get*",
            "kms:List*",
            "kms:Put*",
            "kms:Revoke*",
            "kms:UpdateAlias",
            "kms:ScheduleKeyDeletion",
            "kms:CancelKeyDeletion"
          ],
          "Resource" : "*"
        }
      ]
    }
  )
}

resource "awscc_kms_key" "example" {
  provider     = awscc
  description  = "multi region primary key"
  multi_region = true
  key_policy = jsonencode({
    "Version" : "2012-10-17",
    "Id" : "example_policy",
    "Statement" : [
      {
        "Sid" : "Enable IAM User Permissions",
        "Effect" : "Allow",
        "Principal" : {
          "AWS" : "arn:aws:iam::0123456789012:root"
        },
        "Action" : "kms:*",
        "Resource" : "*"
      },
    ],
    }
  )
}
