# Get current AWS region
data "aws_region" "current" {}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Create S3 bucket for offline store
resource "awscc_s3_bucket" "feature_store" {
  bucket_name = "sagemaker-feature-store-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create IAM role for SageMaker Feature Group
data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    effect  = "Allow"

    principals {
      type        = "Service"
      identifiers = ["sagemaker.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "feature_group" {
  statement {
    effect = "Allow"
    actions = [
      "s3:PutObject",
      "s3:GetObject",
      "s3:DeleteObject",
      "s3:ListBucket",
      "s3:GetBucketAcl"
    ]
    resources = [
      awscc_s3_bucket.feature_store.arn,
      "${awscc_s3_bucket.feature_store.arn}/*"
    ]
  }

  statement {
    effect = "Allow"
    actions = [
      "glue:CreateDatabase",
      "glue:DeleteDatabase",
      "glue:GetDatabase",
      "glue:GetTable",
      "glue:CreateTable",
      "glue:DeleteTable",
      "glue:UpdateTable"
    ]
    resources = ["*"]
  }
}

resource "awscc_iam_role" "feature_group" {
  role_name = "sagemaker-feature-group-role"
  assume_role_policy_document = jsonencode(
    jsondecode(data.aws_iam_policy_document.assume_role.json)
  )
  policies = [
    {
      policy_name = "SageMakerFeatureGroupAccess"
      policy_document = jsonencode(
        jsondecode(data.aws_iam_policy_document.feature_group.json)
      )
    }
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create SageMaker Feature Group
resource "awscc_sagemaker_feature_group" "example" {
  feature_group_name             = "customer-churn-features"
  record_identifier_feature_name = "customer_id"
  event_time_feature_name        = "timestamp"
  description                    = "Customer churn prediction feature group"

  feature_definitions = [
    {
      feature_name = "customer_id"
      feature_type = "String"
    },
    {
      feature_name = "timestamp"
      feature_type = "String"
    },
    {
      feature_name = "age"
      feature_type = "Integral"
    },
    {
      feature_name = "subscription_type"
      feature_type = "String"
    }
  ]

  offline_store_config = {
    disable_glue_table_creation = false
    s3_storage_config = {
      s3_uri = "s3://${awscc_s3_bucket.feature_store.bucket_name}/offline-store"
    }
  }

  online_store_config = {
    enable_online_store = true
    security_config     = {}
  }

  role_arn = awscc_iam_role.feature_group.arn

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}