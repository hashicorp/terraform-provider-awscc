# Data sources for current AWS account and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Example resource group that queries EC2 instances with specific tags
resource "awscc_resourcegroups_group" "example" {
  name        = "example-resource-group"
  description = "Example resource group that finds EC2 instances with specific tags"

  resource_query = {
    query = {
      resource_type_filters = ["AWS::EC2::Instance"]
      tag_filters = [{
        key    = "Environment"
        values = ["Production"]
      }]
    }
    type = "TAG_FILTERS_1_0"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}