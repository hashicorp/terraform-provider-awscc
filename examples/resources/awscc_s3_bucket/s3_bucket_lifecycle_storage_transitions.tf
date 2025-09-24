resource "awscc_s3_bucket" "example" {
  bucket_name = "example-bucket-lifecycle-rules"
  versioning_configuration = {
    status = "Enabled"
  }
  lifecycle_configuration = {
    rules = [
      {
        id = "non_current_version_transitions"

        noncurrent_version_expiration_in_days = 90
        noncurrent_version_transitions = [
          {
            transition_in_days = 30
            storage_class      = "STANDARD_IA"
          },
          {
            transition_in_days = 60
            storage_class      = "INTELLIGENT_TIERING"
          }
        ]
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