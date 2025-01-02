# Get current region
data "aws_region" "current" {}

# Get current account ID
data "aws_caller_identity" "current" {}

# Create CodeDeploy Application
resource "awscc_codedeploy_application" "example" {
  application_name = "example-application"
  compute_platform = "Server" # Valid values: Server, Lambda, or ECS

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}