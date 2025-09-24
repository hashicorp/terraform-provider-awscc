data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create the VPC Lattice Service Network
resource "awscc_vpclattice_service_network" "example" {
  name      = "example-network"
  auth_type = "NONE"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create S3 bucket for logs
resource "awscc_s3_bucket" "logs" {
  bucket_name = "vpc-lattice-logs-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create bucket policy
data "aws_iam_policy_document" "bucket_policy" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["vpc-lattice.amazonaws.com"]
    }
    actions = [
      "s3:PutObject"
    ]
    resources = [
      "${awscc_s3_bucket.logs.arn}/*"
    ]
  }
}

resource "awscc_s3_bucket_policy" "logs" {
  bucket          = awscc_s3_bucket.logs.id
  policy_document = data.aws_iam_policy_document.bucket_policy.json
}

# Create the access log subscription
resource "awscc_vpclattice_access_log_subscription" "example" {
  destination_arn          = awscc_s3_bucket.logs.arn
  resource_identifier      = awscc_vpclattice_service_network.example.id
  service_network_log_type = "SERVICE"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}