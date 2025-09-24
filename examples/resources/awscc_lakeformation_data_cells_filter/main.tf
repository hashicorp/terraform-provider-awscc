# Current AWS account data source
data "aws_caller_identity" "current" {}

# Set up Lake Formation resources
resource "aws_lakeformation_data_lake_settings" "example" {
  admins = [data.aws_caller_identity.current.arn]
}

resource "aws_glue_catalog_database" "example" {
  name        = "example_database"
  description = "Example database for data cells filter"
}

resource "aws_glue_catalog_table" "example" {
  name          = "example_table"
  database_name = aws_glue_catalog_database.example.name

  storage_descriptor {
    columns {
      name = "id"
      type = "int"
    }
    columns {
      name = "name"
      type = "string"
    }
    columns {
      name = "email"
      type = "string"
    }
    location = "s3://example-bucket/example-path/"
  }
}

resource "awscc_lakeformation_data_cells_filter" "example" {
  name             = "example-filter"
  table_catalog_id = data.aws_caller_identity.current.account_id
  database_name    = aws_glue_catalog_database.example.name
  table_name       = aws_glue_catalog_table.example.name

  column_names = [
    "id",
    "name"
  ]

  row_filter = {
    filter_expression = "id > 100"
  }
}