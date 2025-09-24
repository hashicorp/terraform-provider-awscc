resource "awscc_s3_bucket" "example" {
  bucket_name = "example-bucket-versioned"
  versioning_configuration = {
    status = "Enabled"
  }

  tags = [{
    key   = "Name"
    value = "My bucket"
  }]

}