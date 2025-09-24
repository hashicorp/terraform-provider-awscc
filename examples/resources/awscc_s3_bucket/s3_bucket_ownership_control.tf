resource "awscc_s3_bucket" "example" {
  bucket_name = "example-bucket"
  ownership_controls = {
    rules = [{
      object_ownership = "BucketOwnerPreferred"
    }]
  }

  tags = [{
    key   = "Name"
    value = "My bucket"
  }]

}