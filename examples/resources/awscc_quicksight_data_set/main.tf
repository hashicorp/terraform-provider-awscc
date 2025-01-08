# Get current account ID and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Example of an QuickSight dataset
resource "awscc_quicksight_data_set" "example" {
  aws_account_id = data.aws_caller_identity.current.account_id
  data_set_id    = "example-dataset"
  name           = "Example Dataset"
  import_mode    = "SPICE"

  physical_table_map = {
    "table1" = {
      relational_table = {
        data_source_arn = "arn:aws:quicksight:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:datasource/example-datasource"
        schema          = "default"
        name            = "example_table"
        input_columns = [
          {
            name = "id"
            type = "INTEGER"
          },
          {
            name = "name"
            type = "STRING"
          }
        ]
      }
    }
  }

  logical_table_map = {
    "logical1" = {
      alias = "Example Table"
      source = {
        physical_table_id = "table1"
      }
    }
  }

  permissions = [{
    actions = [
      "quicksight:UpdateDataSetPermissions",
      "quicksight:DescribeDataSet",
      "quicksight:DescribeDataSetPermissions",
      "quicksight:PassDataSet",
      "quicksight:DescribeIngestion",
      "quicksight:ListIngestions",
      "quicksight:UpdateDataSet",
      "quicksight:DeleteDataSet",
      "quicksight:CreateIngestion",
      "quicksight:CancelIngestion"
    ]
    principal = "arn:aws:quicksight:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:group/default/example-group"
  }]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}