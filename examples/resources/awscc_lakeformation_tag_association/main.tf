# Get the current AWS account ID
data "aws_caller_identity" "current" {}

# Create example database
resource "awscc_glue_database" "example" {
  catalog_id = data.aws_caller_identity.current.account_id
  database_input = {
    name        = "example_db"
    description = "Example database for Lake Formation tag association"
  }
}

# Create example LF-Tag
resource "awscc_lakeformation_tag" "example" {
  catalog_id = data.aws_caller_identity.current.account_id
  tag_key    = "environment"
  tag_values = ["dev", "test", "prod"]
}

# Create the tag association
resource "awscc_lakeformation_tag_association" "example" {
  lf_tags = [{
    catalog_id = data.aws_caller_identity.current.account_id
    tag_key    = awscc_lakeformation_tag.example.tag_key
    tag_values = ["dev"]
  }]

  resource = {
    database = {
      catalog_id = data.aws_caller_identity.current.account_id
      name       = awscc_glue_database.example.database_input.name
    }
  }
}