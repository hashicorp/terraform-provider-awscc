# S3 Bucket Origin with bucket policy to Origin Access Control
resource "aws_s3_bucket" "s3_origin" {
  bucket = "sampleawsccbucket345"
}

# Block public access to S3 bucket
resource "aws_s3_bucket_public_access_block" "s3_block_public_access" {
  bucket                  = aws_s3_bucket.s3_origin.id
  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

# Attach bucket policy with object access to cloudfront origin
resource "aws_s3_bucket_policy" "allow_access_from_cloudfront" {
  bucket = aws_s3_bucket.s3_origin.id
  policy = data.aws_iam_policy_document.bucket_policy.json
}

# IAM policy document to allow S3 bucket read access to cloudfront origin access control
data "aws_iam_policy_document" "bucket_policy" {
  statement {
    principals {
      type        = "Service"
      identifiers = ["cloudfront.amazonaws.com"]
    }
    effect = "Allow"
    actions = [
      "s3:GetObject",
    ]
    resources = [
      "arn:aws:s3:::${aws_s3_bucket.s3_origin.id}/*"
    ]
    condition {
      test     = "StringEquals"
      variable = "aws:SourceArn"
      values   = ["arn:aws:cloudfront::111111111111:distribution/${awscc_cloudfront_distribution.cloudfront_s3_origin.id}"]
    }
  }
}

# Cloudfront origin access control using AWSCC provider
resource "awscc_cloudfront_origin_access_control" "cf_oac" {
  origin_access_control_config = {
    name                              = "sample-oac"
    description                       = "Sample Origin Access Control Setting using AWSCC"
    origin_access_control_origin_type = "s3"
    signing_behavior                  = "always"
    signing_protocol                  = "sigv4"
  }
}

# Cloudfront distribution with S3 origin using AWSCC provider
resource "awscc_cloudfront_distribution" "cloudfront_s3_origin" {
  distribution_config = {
    enabled             = true
    compress            = true
    default_root_object = "index.html"
    comment             = "Sample Cloudfront Distribution using AWSCC provider"
    default_cache_behavior = {
      target_origin_id       = aws_s3_bucket.s3_origin.id
      viewer_protocol_policy = "redirect-to-https"
      allowed_methods        = ["GET", "HEAD", "OPTIONS"]
      cached_methods         = ["GET", "HEAD", "OPTIONS"]
      min_ttl                = 0
      default_ttl            = 5 * 60
      max_ttl                = 60 * 60
    }
    restrictions = {
      geo_restriction = {
        restriction_type = "none"
      }
    }
    viewer_certificate = {
      cloudfront_default_certificate = true
      minimum_protocol_version       = "TLSv1.2_2018"
    }
    s3_origin = {
      dns_name = aws_s3_bucket.s3_origin.bucket_regional_domain_name
    }
    origins = [{
      domain_name              = aws_s3_bucket.s3_origin.bucket_regional_domain_name
      id                       = "SampleCloudfrontOrigin"
      origin_access_control_id = awscc_cloudfront_origin_access_control.cf_oac.id
    }]
  }
  tags = [{
    key   = "Name"
    value = "Cloudfront Distribution with S3 Origin"
  }]
}