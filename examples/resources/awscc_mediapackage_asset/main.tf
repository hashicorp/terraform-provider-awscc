# IAM role for MediaPackage Asset to access S3
data "aws_iam_policy_document" "mediapackage_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["mediapackage.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "mediapackage_s3_access" {
  statement {
    actions = [
      "s3:GetObject",
      "s3:GetBucketLocation"
    ]
    resources = [
      "arn:aws:s3:::example-bucket/*",
      "arn:aws:s3:::example-bucket"
    ]
  }
}

resource "awscc_iam_role" "mediapackage_role" {
  role_name                   = "mediapackage-asset-access-role"
  assume_role_policy_document = data.aws_iam_policy_document.mediapackage_assume_role.json

  managed_policy_arns = []

  policies = [
    {
      policy_document = data.aws_iam_policy_document.mediapackage_s3_access.json
      policy_name     = "mediapackage-s3-access"
    }
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Required packaging group for the asset
resource "awscc_mediapackage_packaging_group" "example" {
  packaging_group_id = "example-packaging-group"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# MediaPackage Asset
resource "awscc_mediapackage_asset" "example" {
  asset_id           = "example-asset"
  packaging_group_id = awscc_mediapackage_packaging_group.example.id
  source_arn         = "arn:aws:s3:::example-bucket/example-video.mp4"
  source_role_arn    = awscc_iam_role.mediapackage_role.arn

  resource_id = "example-resource-id"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}