# Data sources to get AWS account ID and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# CloudFront Origin Access Control
resource "awscc_cloudfront_origin_access_control" "example" {
  origin_access_control_config = {
    name                              = "example-oac"
    description                       = "Example Origin Access Control for S3"
    origin_access_control_origin_type = "s3"
    signing_behavior                  = "always"
    signing_protocol                  = "sigv4"
  }
}