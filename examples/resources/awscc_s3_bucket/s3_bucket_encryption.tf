resource "awscc_s3_bucket" "example" {
  bucket_name = "example-bucket"

  bucket_encryption = {
    server_side_encryption_configuration = [{
      server_side_encryption_by_default = {
        sse_algorithm = "AES256"
      }
    }]
  }
}