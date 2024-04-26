resource "awscc_cloudfront_key_value_store" "example" {
  name    = "ExampleKeyValueStore"
  comment = "This is an example key value store"
  import_source = {
    source_arn  = "arn:aws:s3:::your-bucket-name/key-value-file"
    source_type = "S3"
  }
}