resource "awscc_s3vectors_vector_bucket" "example" {
  vector_bucket_name = "example-vectors-bucket"

  encryption_configuration = {
    sse_type = "AES256"
  }

  tags = [
    {
      key   = "Name"
      value = "example-vectors-bucket"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}

resource "awscc_s3vectors_index" "example" {
  index_name         = "example-vectors-index"
  vector_bucket_name = awscc_s3vectors_vector_bucket.example.vector_bucket_name
  data_type          = "float32"
  dimension          = 1536
  distance_metric    = "cosine"

  tags = [
    {
      key   = "Name"
      value = "example-vectors-index"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]

  depends_on = [awscc_s3vectors_vector_bucket.example]
}