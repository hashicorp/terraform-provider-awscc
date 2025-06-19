# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Get current AWS region
# Note: Using data.aws_region.current.region (AWS provider v6.0+)
# For AWS provider < v6.0, use data.aws_region.current.name instead
data "aws_region" "current" {}

# Create an S3 bucket first
resource "awscc_s3_bucket" "example" {
  bucket_name = "example-access-point-bucket-${data.aws_caller_identity.current.account_id}"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Define the access point policy
data "aws_iam_policy_document" "access_point_policy" {
  statement {
    effect = "Allow"
    principals {
      type = "AWS"
      identifiers = [
        "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
      ]
    }
    actions = [
      "s3:GetObject",
      "s3:PutObject"
    ]
    resources = [
      "arn:aws:s3:${data.aws_region.current.region}:${data.aws_caller_identity.current.account_id}:accesspoint/example-access-point/object/*"
    ]
  }
}

# Create the S3 Access Point
resource "awscc_s3_access_point" "example" {
  name   = "example-access-point"
  bucket = awscc_s3_bucket.example.id
  policy = jsonencode(jsondecode(data.aws_iam_policy_document.access_point_policy.json))

  public_access_block_configuration = {
    block_public_acls       = true
    block_public_policy     = true
    ignore_public_acls      = true
    restrict_public_buckets = true
  }
}