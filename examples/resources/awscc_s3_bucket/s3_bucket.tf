resource "awscc_s3_bucket" "example" {
  bucket_name = "example-bucket"

  tags = [{
    key   = "Name"
    value = "My bucket"
  }]

}