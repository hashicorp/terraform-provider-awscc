# Create EventSchemas Registry
resource "awscc_eventschemas_registry" "example" {
  registry_name = "example-registry"
  description   = "Example schema registry created by AWSCC provider"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}