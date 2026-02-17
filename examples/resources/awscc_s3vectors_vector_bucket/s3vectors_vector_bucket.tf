resource "awscc_s3vectors_vector_bucket" "example" {
  vector_bucket_name = "example-vector-bucket"

  encryption_configuration = {
    sse_type = "AES256"
  }

  tags = [
    {
      key   = "Name"
      value = "example-vector-bucket"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}