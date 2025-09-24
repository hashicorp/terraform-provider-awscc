resource "awscc_lightsail_bucket" "example" {
  bucket_name = "example-bucket"
  bundle_id   = "small_1_0"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
