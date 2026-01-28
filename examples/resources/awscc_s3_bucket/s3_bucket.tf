# S3 Bucket
resource "awscc_s3_bucket" "example" {
  bucket_name = "example-bucket"
  
  public_access_block_configuration = {
    block_public_acls       = true
    block_public_policy     = true
    ignore_public_acls      = true
    restrict_public_buckets = true
  }
  
  versioning_configuration = {
    status = "Enabled"
  }
  
  tags = [
    {
      key   = "Environment"
      value = "example"
    },
    {
      key   = "Name"
      value = "example-s3-bucket"
    }
  ]
}