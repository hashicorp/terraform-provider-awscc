resource "awscc_servicecatalogappregistry_attribute_group" "example" {
  name        = "example-attribute-group"
  description = "Example attribute group created via AWSCC provider"

  attributes = jsonencode({
    "environment" = "production"
    "team"        = "infrastructure"
    "cost-center" = "12345"
  })

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
    }, {
    key   = "Environment"
    value = "Example"
  }]
}