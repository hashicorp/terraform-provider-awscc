resource "awscc_s3_bucket" "example" {
  bucket_name = "example-bucket-lifecycle-rules"
  versioning_configuration = {
    status = "Enabled"
  }
  lifecycle_configuration = {
    rules = [
      {
        id = "expire_non_current_version"
        noncurrent_version_expiration = {
          newer_noncurrent_versions = 1
          noncurrent_days           = 1
        }
        status = "Enabled"
      }
    ]
  }

  tags = [
    {
      key   = "Name"
      value = "My bucket"
    }
  ]
}