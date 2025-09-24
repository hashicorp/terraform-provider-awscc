data "aws_caller_identity" "current" {}

# Create an S3 bucket for CloudFront origin
resource "aws_s3_bucket" "example" {
  bucket = "cloudfront-monitoring-example-${data.aws_caller_identity.current.account_id}"
}

# Origin access identity for CloudFront
resource "aws_cloudfront_origin_access_identity" "example" {
  comment = "Example OAI for monitoring subscription"
}

# Bucket policy to allow CloudFront access
resource "aws_s3_bucket_policy" "example" {
  bucket = aws_s3_bucket.example.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid    = "AllowCloudFrontOAI"
        Effect = "Allow"
        Principal = {
          AWS = "arn:aws:iam::cloudfront:user/CloudFront Origin Access Identity ${aws_cloudfront_origin_access_identity.example.id}"
        }
        Action   = "s3:GetObject"
        Resource = "${aws_s3_bucket.example.arn}/*"
      }
    ]
  })
}

# Create CloudFront distribution
resource "aws_cloudfront_distribution" "example" {
  enabled             = true
  is_ipv6_enabled     = true
  comment             = "Example distribution for monitoring subscription"
  default_root_object = "index.html"

  origin {
    domain_name = aws_s3_bucket.example.bucket_regional_domain_name
    origin_id   = "myS3Origin"

    s3_origin_config {
      origin_access_identity = aws_cloudfront_origin_access_identity.example.cloudfront_access_identity_path
    }
  }

  default_cache_behavior {
    target_origin_id       = "myS3Origin"
    viewer_protocol_policy = "redirect-to-https"
    allowed_methods        = ["GET", "HEAD"]
    cached_methods         = ["GET", "HEAD"]
    min_ttl                = 0
    default_ttl            = 3600
    max_ttl                = 86400

    forwarded_values {
      query_string = false
      cookies {
        forward = "none"
      }
    }
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  viewer_certificate {
    cloudfront_default_certificate = true
  }

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Create CloudFront monitoring subscription
resource "awscc_cloudfront_monitoring_subscription" "example" {
  distribution_id = aws_cloudfront_distribution.example.id
  monitoring_subscription = {
    realtime_metrics_subscription_config = {
      realtime_metrics_subscription_status = "Enabled"
    }
  }
}