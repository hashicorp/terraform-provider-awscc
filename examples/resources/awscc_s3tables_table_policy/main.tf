data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

resource "awscc_s3tables_table_bucket" "example" {
  table_bucket_name = "example-table-bucket-${data.aws_caller_identity.current.account_id}"

  unreferenced_file_removal = {
    status            = "Enabled"
    unreferenced_days = 7
    noncurrent_days   = 30
  }


}

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
          name     = "id"
          type     = "int"
          required = true
        },
        {
          name = "name"
          type = "string"
        },
        {
          name = "timestamp"
          type = "timestamp"
        }
      ]
    }
  }

  snapshot_management = {
    status                 = "enabled"
    max_snapshot_age_hours = 168
    min_snapshots_to_keep  = 3
  }
}

# S3 Tables table policy
resource "awscc_s3tables_table_policy" "example" {
  table_arn = awscc_s3tables_table.example.table_arn

  resource_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid    = "AllowReadAccess"
        Effect = "Allow"
        Principal = {
          AWS = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
        }
        Action = [
          "s3tables:GetTable",
          "s3tables:GetTableMetadata",
          "s3tables:GetTablePolicy"
        ]
        Resource = awscc_s3tables_table.example.table_arn
      },
      {
        Sid    = "AllowWriteAccess"
        Effect = "Allow"
        Principal = {
          AWS = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
        }
        Action = [
          "s3tables:PutTableData",
          "s3tables:UpdateTableMetadata"
        ]
        Resource = awscc_s3tables_table.example.table_arn
      }
    ]
  })
}
