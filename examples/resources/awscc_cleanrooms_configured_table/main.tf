data "aws_caller_identity" "current" {}

# First create a Glue database and table for the example
resource "awscc_glue_database" "example" {
  catalog_id = data.aws_caller_identity.current.account_id
  database_input = {
    name        = "cleanrooms_example_db"
    description = "Example database for AWS Clean Rooms"
  }
}

resource "aws_glue_catalog_table" "example" {
  name          = "sales_data"
  database_name = awscc_glue_database.example.database_input.name
  description   = "Example sales data table"
  table_type    = "EXTERNAL_TABLE"

  storage_descriptor {
    columns {
      name = "transaction_id"
      type = "string"
    }
    columns {
      name = "customer_id"
      type = "string"
    }
    columns {
      name = "purchase_amount"
      type = "double"
    }
    columns {
      name = "transaction_date"
      type = "date"
    }

    location      = "s3://example-bucket/sales-data/"
    input_format  = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"
    ser_de_info {
      serialization_library = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"
    }
  }
}

# Create the Clean Rooms configured table
resource "awscc_cleanrooms_configured_table" "example" {
  name        = "example-sales-table"
  description = "Example sales data configured table for Clean Rooms"

  # Specify which columns are allowed for analysis
  allowed_columns = [
    "customer_id",
    "purchase_amount",
    "transaction_date"
  ]

  # Specify the analysis method
  analysis_method = "DIRECT_QUERY"

  # Reference to the Glue table
  table_reference = {
    glue = {
      database_name = awscc_glue_database.example.database_input.name
      table_name    = aws_glue_catalog_table.example.name
    }
  }

  # Add analysis rules for aggregation
  analysis_rules = [
    {
      type = "AGGREGATION"
      policy = {
        v1 = {
          aggregation = {
            aggregate_columns = [
              {
                column_names = ["purchase_amount"]
                function     = "SUM"
              },
              {
                column_names = ["purchase_amount"]
                function     = "AVG"
              }
            ]
            dimension_columns = ["transaction_date"]
            output_constraints = [
              {
                type        = "COUNT_DISTINCT"
                column_name = "customer_id"
                minimum     = 5
              }
            ]
            join_columns     = ["customer_id"]
            scalar_functions = ["ABS", "CEILING"]
          }
        }
      }
    }
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}