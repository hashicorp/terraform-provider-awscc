resource "awscc_s3outposts_bucket" "example" {
  bucket_name = "example-bucket"
  outpost_id  = "op-01ac5d28a6a232904"

  lifecycle_configuration = {
    rules = [
      {
        id                 = "rule1"
        expiration_in_days = 30
        status             = "Enabled"
      },
      {
        id = "rule2"
        abort_incomplete_multipart_upload = {
          days_after_initiation = 7
        }
        status = "Enabled"
        filter = {
          prefix = "documents/"
        }
      }
    ]
  }


  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}