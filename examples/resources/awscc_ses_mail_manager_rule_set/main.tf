data "aws_caller_identity" "current" {}

# S3 bucket for email archival
resource "awscc_s3_bucket" "email_archive" {
  bucket_name = "ses-email-archive-${data.aws_caller_identity.current.account_id}"
}

# IAM role for SES to write to S3
data "aws_iam_policy_document" "ses_assume_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["ses.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

data "aws_iam_policy_document" "ses_s3_access" {
  statement {
    effect = "Allow"
    actions = [
      "s3:PutObject",
      "s3:GetObject",
      "s3:ListBucket"
    ]
    resources = [
      awscc_s3_bucket.email_archive.arn,
      "${awscc_s3_bucket.email_archive.arn}/*"
    ]
  }
}

resource "awscc_iam_role" "ses_s3_role" {
  role_name = "SESMailManagerS3Role"
  assume_role_policy_document = jsonencode(
    jsondecode(data.aws_iam_policy_document.ses_assume_role.json)
  )
  policies = [{
    policy_document = jsonencode(
      jsondecode(data.aws_iam_policy_document.ses_s3_access.json)
    )
    policy_name = "SESMailManagerS3Access"
  }]
  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}

# SES Mail Manager Rule Set
resource "awscc_ses_mail_manager_rule_set" "example" {
  rule_set_name = "example-ruleset"
  rules = [
    {
      name = "archive-emails"
      actions = [
        {
          write_to_s3 = {
            action_failure_policy = "CONTINUE"
            role_arn              = awscc_iam_role.ses_s3_role.arn
            s3_bucket             = awscc_s3_bucket.email_archive.bucket_name
            s3_prefix             = "incoming/"
          }
        }
      ]
      conditions = [
        {
          string_expression = {
            evaluate = {
              attribute = "FROM"
            }
            operator = "CONTAINS"
            values   = ["@example.com"]
          }
        }
      ]
    }
  ]
  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}