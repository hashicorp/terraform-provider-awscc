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