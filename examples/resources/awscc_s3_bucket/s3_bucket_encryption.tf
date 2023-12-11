resource "awscc_s3_bucket" "example" {
  bucket_name = "wellsiau-example-bucket-2"

  bucket_encryption = {
    server_side_encryption_configuration = [{
      server_side_encryption_by_default = {
        sse_algorithm = "AES256"
      }
    }]
  }
}