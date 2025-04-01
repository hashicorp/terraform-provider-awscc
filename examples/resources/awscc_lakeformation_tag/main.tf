# Get the current AWS account ID
data "aws_caller_identity" "current" {}

# Create the Lake Formation Tag
resource "awscc_lakeformation_tag" "example" {
  catalog_id = data.aws_caller_identity.current.account_id
  tag_key    = "environment"
  tag_values = ["dev", "staging", "prod"]
}