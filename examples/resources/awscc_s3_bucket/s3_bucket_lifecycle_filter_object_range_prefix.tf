resource "awscc_s3_bucket" "example" {
  bucket_name = "example-bucket-lifecycle-rules"
  versioning_configuration = {
    status = "Enabled"
  }

  tags = [
    {
      key   = "Name"
      value = "My bucket"
    }
  ]

  lifecycle_configuration = {
    rules = [
      {
        id = "expire_non_current_version"
        noncurrent_version_expiration = {
          newer_noncurrent_versions = 1
          noncurrent_days           = 1
        }
        prefix                   = "logs/"
        object_size_greater_than = 500
        object_size_less_than    = 64000
        status                   = "Enabled"
      }

    ]
  }
}
