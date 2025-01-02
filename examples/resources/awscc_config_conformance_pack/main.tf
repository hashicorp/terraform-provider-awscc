# Get current AWS region
data "aws_region" "current" {}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# S3 bucket for storing conformance pack files
resource "awscc_s3_bucket" "config_bucket" {
  bucket_name = "config-conformance-pack-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Bucket policy to allow AWS Config to write to the bucket
data "aws_iam_policy_document" "bucket_policy" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["config.amazonaws.com"]
    }
    actions = [
      "s3:GetBucketAcl",
      "s3:PutObject"
    ]
    resources = [
      "arn:aws:s3:::config-conformance-pack-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}",
      "arn:aws:s3:::config-conformance-pack-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}/*"
    ]
  }
}

# Apply bucket policy
resource "awscc_s3_bucket_policy" "config_bucket_policy" {
  bucket          = awscc_s3_bucket.config_bucket.id
  policy_document = jsonencode(data.aws_iam_policy_document.bucket_policy.json)
}

# Create the conformance pack
resource "awscc_config_conformance_pack" "example" {
  conformance_pack_name  = "operational-best-practices"
  delivery_s3_bucket     = awscc_s3_bucket.config_bucket.id
  delivery_s3_key_prefix = "config"

  template_body = <<EOF
Resources:
  IAMPasswordPolicy:
    Type: AWS::Config::ConfigRule
    Properties:
      ConfigRuleName: iam-password-policy
      Description: Checks if IAM password policy meets security standards
      Source:
        Owner: AWS
        SourceIdentifier: IAM_PASSWORD_POLICY
      Scope:
        ComplianceResourceTypes:
          - AWS::IAM::User
  RootAccountMFAEnabled:
    Type: AWS::Config::ConfigRule
    Properties:
      ConfigRuleName: root-account-mfa-enabled
      Description: Checks if root account has MFA enabled
      Source:
        Owner: AWS
        SourceIdentifier: ROOT_ACCOUNT_MFA_ENABLED
EOF
}