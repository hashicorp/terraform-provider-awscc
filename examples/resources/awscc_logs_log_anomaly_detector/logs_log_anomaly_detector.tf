resource "awscc_logs_log_anomaly_detector" "example" {
  account_id              = data.aws_caller_identity.current.account_id
  anomaly_visibility_time = 30
  detector_name           = "example"
  evaluation_frequency    = "ONE_HOUR"
  filter_pattern          = "%AUTHORIZED%"
  log_group_arn_list      = ["arn:${data.aws_partition.current.name}:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:log-group:${awscc_logs_log_group.example.id}"]
  kms_key_id              = awscc_kms_key.example.arn
}

resource "awscc_logs_log_group" "example" {
  log_group_name    = "example"
  retention_in_days = 7
}

resource "awscc_kms_key" "example" {
  description = "KMS Key for log anomaly detector"
  key_policy = jsonencode({
    "Version" : "2012-10-17",
    "Id" : "KMS-Key-Policy-For-Root",
    "Statement" : [
      {
        "Sid" : "Enable IAM User Permissions",
        "Effect" : "Allow",
        "Principal" : {
          "AWS" : "arn:${data.aws_partition.current.name}:iam::${data.aws_caller_identity.current.account_id}:root"
        },
        "Action" : "kms:*",
        "Resource" : "*"
      },
      {
        "Effect" : "Allow",
        "Principal" : {
          "Service" : "logs.${data.aws_region.current.name}.amazonaws.com"
        },
        "Action" : [
          "kms:Encrypt",
          "kms:Decrypt",
          "kms:GenerateDataKey*",
          "kms:DescribeKey"
        ],
        "Resource" : "*",
        "Condition" : {
          "ArnLike" : {
            "kms:EncryptionContext:aws:logs:arn" : "arn:${data.aws_partition.current.name}:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:anomaly-detector:*"
          }
        }
      },
      {
        "Effect" : "Allow",
        "Principal" : {
          "Service" : "logs.${data.aws_region.current.name}.amazonaws.com"
        },
        "Action" : [
          "kms:Encrypt",
          "kms:Decrypt",
          "kms:ReEncrypt*",
          "kms:GenerateDataKey*",
          "kms:DescribeKey"
        ],
        "Resource" : "*",
        "Condition" : {
          "ArnLike" : {
            "kms:EncryptionContext:aws-crypto-ec:aws:logs:arn" : "arn:${data.aws_partition.current.name}:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:anomaly-detector:*"
          }
        }
      }
    ],
    }
  )
}

data "aws_caller_identity" "current" {}
data "aws_region" "current" {}
data "aws_partition" "current" {}