# Get current AWS region
data "aws_region" "current" {}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Create EventSchemas Registry
resource "awscc_eventschemas_registry" "example" {
  registry_name = "example-registry"
  description   = "Example schema registry created by AWSCC provider"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}