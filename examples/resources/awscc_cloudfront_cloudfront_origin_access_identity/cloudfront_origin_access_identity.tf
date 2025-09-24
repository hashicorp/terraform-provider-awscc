# Cloudfront origin access identity
resource "awscc_cloudfront_cloudfront_origin_access_identity" "cf_oai" {
  cloudfront_origin_access_identity_config = {
    comment = "SampleAWSCCCloudFrontOAI"
  }
}