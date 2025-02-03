# Get current AWS region
data "aws_region" "current" {}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Create the Deadline Farm
resource "awscc_deadline_farm" "example" {
  display_name = "ExampleFarm"
  description  = "Example Deadline Farm created with AWSCC provider"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}