# Batch Consumable Resource Example
resource "awscc_batch_consumable_resource" "demo" {
  resource_type            = "REPLENISHABLE"
  total_quantity           = 10
  consumable_resource_name = "demo-license-resource"

  tags = [{
    key   = "Environment"
    value = "demo"
  }, {
    key   = "Modified By"
    value = "AWSCC"
  }]
}