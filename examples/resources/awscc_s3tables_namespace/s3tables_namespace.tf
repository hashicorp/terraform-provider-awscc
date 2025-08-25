# S3 Tables table bucket
resource "awscc_s3tables_table_bucket" "example" {
  table_bucket_name = "example-table-bucket"
}

# S3 Tables namespace
resource "awscc_s3tables_namespace" "example" {
  namespace        = "examplenamespace"
  table_bucket_arn = awscc_s3tables_table_bucket.example.table_bucket_arn
}
