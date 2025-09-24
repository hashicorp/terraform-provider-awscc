# Data sources to get current AWS region and account ID
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

# IAM Role for DataZone subscription target
resource "awscc_iam_role" "datazone_subscription_target" {
  role_name = "DataZoneSubscriptionTargetRole"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Principal = {
          Service = "datazone.amazonaws.com"
        }
        Action = "sts:AssumeRole"
      }
    ]
  })
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "aws_iam_role_policy" "datazone_subscription_target" {
  name = "DataZoneSubscriptionTargetPolicy"
  role = awscc_iam_role.datazone_subscription_target.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "s3:GetObject",
          "s3:PutObject",
          "s3:DeleteObject",
          "s3:ListBucket"
        ]
        Resource = [
          "arn:aws:s3:::datazone-subscription-*",
          "arn:aws:s3:::datazone-subscription-*/*"
        ]
      }
    ]
  })
}

# DataZone Subscription Target
resource "awscc_datazone_subscription_target" "example" {
  name                   = "example-subscription-target"
  domain_identifier      = "dzd_example123456" # Replace with actual domain ID
  environment_identifier = "env-example123456" # Replace with actual environment ID
  type                   = "S3"
  provider_name          = "AWS"

  applicable_asset_types = ["DATA_CATALOG"]
  authorized_principals  = [data.aws_caller_identity.current.arn]
  manage_access_role     = awscc_iam_role.datazone_subscription_target.arn

  subscription_target_config = [
    {
      form_name = "AWS_S3_CONFIGURATION"
      content = jsonencode({
        bucket = "datazone-subscription-${data.aws_caller_identity.current.account_id}"
        region = data.aws_region.current.name
      })
    }
  ]
}