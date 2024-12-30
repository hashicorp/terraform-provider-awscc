
# Get current AWS region
data "aws_region" "current" {}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Example of awscc_appintegrations_event_integration resource
resource "awscc_appintegrations_event_integration" "example" {
  name             = "example-event-integration"
  description      = "Example event integration using AWSCC provider"
  event_bridge_bus = "default"

  event_filter = {
    source = "aws.partner/example.com/source"
  }

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}