# Example of creating a Synthetics Group using AWSCC provider
# Fetches AWS region and account information
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

# Example Synthetics Group
resource "awscc_synthetics_group" "example" {
  name = "example-synthetics-group"

  # Optional tags
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}