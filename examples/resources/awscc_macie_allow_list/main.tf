# Create an S3 bucket to store the allow list
resource "awscc_s3_bucket" "allow_list_bucket" {
  bucket_name = "macie-allow-list-example"
}

# Create a bucket object for the allow list words
resource "aws_s3_object" "allow_list_words" {
  bucket  = awscc_s3_bucket.allow_list_bucket.id
  key     = "allowlist.txt"
  content = "example.com\ntest.com"
}

# Create Macie allow list
resource "awscc_macie_allow_list" "example" {
  name        = "example-allow-list"
  description = "Example Macie allow list for demonstration"

  criteria = {
    s3_words_list = {
      bucket_name = awscc_s3_bucket.allow_list_bucket.id
      object_key  = aws_s3_object.allow_list_words.key
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}