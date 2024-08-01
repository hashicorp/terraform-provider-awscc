resource "awscc_lightsail_bucket" "example" {
  bucket_name = "example-bucket"
  bundle_id   = "small_1_0"
}

resource "awscc_lightsail_distribution" "example" {
  bundle_id = "small_1_0"
  default_cache_behavior = {
    behavior = "cache"
  }
  distribution_name = "example-dist"
  origin = {
    name        = awscc_lightsail_bucket.example.bucket_name
    region_name = "us-east-1"
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
