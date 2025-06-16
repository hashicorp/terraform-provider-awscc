data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create a directory bucket first (requires S3 Express)
resource "awscc_s3express_directory_bucket" "example" {
  bucket_name     = "example-express-directory-bucket"
  data_redundancy = "SingleAvailabilityZone"
  location_name   = "${data.aws_region.current.name}a"
}

data "aws_iam_policy_document" "access_point_policy" {
  statement {
    effect = "Allow"
    principals {
      type = "AWS"
      identifiers = [
        data.aws_caller_identity.current.arn
      ]
    }
    actions = [
      "s3:GetObject",
      "s3:PutObject",
      "s3:ListBucket"
    ]
    resources = [
      "arn:aws:s3:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:accesspoint/*",
      "arn:aws:s3:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:accesspoint/*/object/*"
    ]
  }
}

resource "awscc_s3express_access_point" "example" {
  name   = "example-access-point"
  bucket = awscc_s3express_directory_bucket.example.id
  policy = jsonencode(data.aws_iam_policy_document.access_point_policy.json)

  public_access_block_configuration = {
    block_public_acls       = true
    block_public_policy     = true
    ignore_public_acls      = true
    restrict_public_buckets = true
  }

  scope = {
    permissions = ["GetObject", "PutObject", "ListBucket"]
    prefixes    = ["documents/", "images/"]
  }
}