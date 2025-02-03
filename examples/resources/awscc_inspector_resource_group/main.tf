# Get current region for reference
data "aws_region" "current" {}

# Get current account ID for reference
data "aws_caller_identity" "current" {}

# Create an Inspector Resource Group
resource "awscc_inspector_resource_group" "example" {
  resource_group_tags = [{
    key   = "Environment"
    value = "Production"
    }, {
    key   = "Modified By"
    value = "AWSCC"
  }]
}