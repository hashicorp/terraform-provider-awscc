resource "awscc_s3_bucket" "example" {
  bucket_name = "example-bucket-lifecycle-rules"
  versioning_configuration = {
    status = "Enabled"
  }
  lifecycle_configuration = {
    rules = [
      {
        id = "abort_incomplete_upload"
        abort_incomplete_multipart_upload = {
          days_after_initiation = 1
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

