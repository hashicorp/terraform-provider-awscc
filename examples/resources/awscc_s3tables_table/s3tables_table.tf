# S3 Tables table bucket
resource "awscc_s3tables_table_bucket" "example" {
  table_bucket_name = "example-bucket"
}

# S3 Tables namespace
resource "awscc_s3tables_namespace" "example" {
  table_bucket_arn = awscc_s3tables_table_bucket.example.table_bucket_arn
  namespace        = "example_namespace"
}

# S3 Tables table
resource "awscc_s3tables_table" "example" {
  table_bucket_arn  = awscc_s3tables_table_bucket.example.table_bucket_arn
  namespace         = awscc_s3tables_namespace.example.namespace
  table_name        = "example_table"
  open_table_format = "ICEBERG"

  iceberg_metadata = {
    iceberg_schema = {
      schema_field_list = [
        {
          name        = "id"
          type        = "int"
          is_required = true
        },
        {
          name = "name"
          type = "string"
        },
        {
          name = "value"
          type = "int"
        }
      ]
    }
  }

  # Optional configuration for snapshot management
  snapshot_management = {
    snapshot_retention_period = 7
  }
}
