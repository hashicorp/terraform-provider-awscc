data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Glue Database Example
resource "awscc_glue_database" "example" {
  catalog_id = data.aws_caller_identity.current.account_id
  database_input = {
    name        = "example_database"
    description = "Example AWS Glue Database created using AWSCC provider"
    parameters = jsonencode({
      "classification" = "example"
      "purpose"        = "testing"
    })
    create_table_default_permissions = [
      {
        permissions = ["SELECT", "ALTER", "DROP"]
        principal = {
          data_lake_principal_identifier = "IAM_ALLOWED_PRINCIPALS"
        }
      }
    ]
    location_uri = "s3://example-bucket-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}/databases/example_database"
  }
}