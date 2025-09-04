# Create S3 bucket in us-west-2
resource "awscc_s3_bucket" "west" {
  bucket_name = "example-bucket-west"
}

# Create the Multi-Region Access Point
resource "awscc_s3_multi_region_access_point" "example" {
  name = "example-mrap"

  regions = [
    {
      bucket = awscc_s3_bucket.west.bucket_name
    }
  ]

  public_access_block_configuration = {
    block_public_acls       = true
    block_public_policy     = true
    ignore_public_acls      = true
    restrict_public_buckets = true
  }
}
