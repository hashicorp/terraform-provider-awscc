# Create the application
resource "awscc_servicecatalogappregistry_application" "example" {
  name        = "example-application"
  description = "Example application for testing attribute group association"

  tags = {
    ModifiedBy = "AWSCC"
  }
}

# Create the attribute group
resource "awscc_servicecatalogappregistry_attribute_group" "example" {
  name        = "example-attribute-group"
  description = "Example attribute group for testing association"
  attributes = jsonencode({
    stage = "dev"
    team  = "platform"
  })

  tags = {
    ModifiedBy = "AWSCC"
  }
}

# Create the association between application and attribute group
resource "awscc_servicecatalogappregistry_attribute_group_association" "example" {
  application     = awscc_servicecatalogappregistry_application.example.id
  attribute_group = awscc_servicecatalogappregistry_attribute_group.example.id
}