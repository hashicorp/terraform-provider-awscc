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
        id = "expire_non_current_version_filtered_by_tags"
        noncurrent_version_expiration = {
          newer_noncurrent_versions = 1
          noncurrent_days           = 1
        }
        prefix = "logs/"
        tag_filters = [{
          key   = "key1"
          value = "value1"
          },
          {
            key   = "key2"
            value = "value2"
          }
        ]
        status = "Enabled"
      }
    ]
  }
}