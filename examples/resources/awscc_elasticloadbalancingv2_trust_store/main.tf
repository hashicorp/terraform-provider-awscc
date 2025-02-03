# Get account ID and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# First let's create a sample CA bundle in S3
resource "aws_s3_bucket" "ca_bundle" {
  bucket        = "trust-store-ca-bundle-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"
  force_destroy = true
}

# Block public access to the bucket
resource "aws_s3_bucket_public_access_block" "ca_bundle" {
  bucket = aws_s3_bucket.ca_bundle.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

# Enable versioning for the S3 bucket
resource "aws_s3_bucket_versioning" "ca_bundle" {
  bucket = aws_s3_bucket.ca_bundle.id
  versioning_configuration {
    status = "Enabled"
  }
}



# Create a sample CA bundle file
resource "aws_s3_object" "ca_bundle" {
  bucket       = aws_s3_bucket.ca_bundle.id
  key          = "ca-bundle.pem"
  source       = "ca-bundle2.pem"
  content_type = "application/x-pem-file"
}

# Create the trust store
resource "awscc_elasticloadbalancingv2_trust_store" "example" {
  name                             = "example-trust-store"
  ca_certificates_bundle_s3_bucket = aws_s3_bucket.ca_bundle.id
  ca_certificates_bundle_s3_key    = aws_s3_object.ca_bundle.key

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}