data "aws_caller_identity" "current" {}

resource "awscc_athena_data_catalog" "this" {
  name = "awscc-catalog"
  type = "GLUE"
  parameters = {
    catalog-id = data.aws_caller_identity.current.account_id
  }
}