resource "awscc_s3_bucket" "example" {}

resource "awscc_omics_sequence_store" "example" {
  name        = "example"
  description = "example"

  fallback_location = "s3://${awscc_s3_bucket.example.bucket_name}"
}