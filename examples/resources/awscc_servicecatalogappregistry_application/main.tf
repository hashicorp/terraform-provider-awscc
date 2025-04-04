# Example of App Registry Application
resource "awscc_servicecatalogappregistry_application" "example" {
  name        = "demo-application"
  description = "Demo application for App Registry"

  tags = [{
    key   = "Environment"
    value = "Production"
    }, {
    key   = "Department"
    value = "Engineering"
    }, {
    key   = "Modified By"
    value = "AWSCC"
  }]
}