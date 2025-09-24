# Example configuration for awscc_elasticloadbalancingv2_trust_store_revocation

data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create an S3 bucket to store CA certificates and revocation lists
resource "aws_s3_bucket" "revocations" {
  bucket = "example-revocations-bucket-${data.aws_caller_identity.current.account_id}"
}

# Create S3 bucket policy to allow ELB to access objects
resource "aws_s3_bucket_policy" "allow_elb_access" {
  bucket = aws_s3_bucket.revocations.id
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid    = "AllowELBService"
        Effect = "Allow"
        Principal = {
          Service = "elasticloadbalancing.amazonaws.com"
        }
        Action = [
          "s3:GetObject",
          "s3:GetBucketLocation"
        ]
        Resource = [
          aws_s3_bucket.revocations.arn,
          "${aws_s3_bucket.revocations.arn}/*"
        ]
      }
    ]
  })
}

# Configure bucket to be private
resource "aws_s3_bucket_public_access_block" "revocations" {
  bucket = aws_s3_bucket.revocations.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

# Example CRL revocation object
resource "aws_s3_object" "revocation_list" {
  bucket  = aws_s3_bucket.revocations.id
  key     = "revocations/example.crl"
  content = "# This is a placeholder. Replace with real CRL content"
}

# Example awscc_elasticloadbalancingv2_trust_store_revocation resource
resource "awscc_elasticloadbalancingv2_trust_store_revocation" "example" {
  trust_store_arn = "arn:aws:elasticloadbalancing:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:truststore/example-trust-store"
  revocation_contents = [{
    revocation_type = "CRL"
    s3_bucket       = aws_s3_bucket.revocations.id
    s3_key          = aws_s3_object.revocation_list.key
  }]
}