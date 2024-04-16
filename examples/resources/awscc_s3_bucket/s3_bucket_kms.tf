resource "awscc_kms_key" "example" {
  description         = "S3 KMS key"
  enable_key_rotation = true
}

resource "awscc_s3_bucket" "example" {
  bucket_name = "example-bucket-kms"

  bucket_encryption = {
    server_side_encryption_configuration = [{
      server_side_encryption_by_default = {
        sse_algorithm     = "aws:kms"
        kms_master_key_id = awscc_kms_key.example.arn
      }
    }]
  }
}