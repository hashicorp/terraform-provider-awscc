# Get current AWS region
data "aws_region" "current" {}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Create the Route Calculator
resource "awscc_location_route_calculator" "example" {
  calculator_name = "example-route-calculator"
  data_source     = "Esri"
  description     = "Example Route Calculator created by AWSCC provider"
  pricing_plan    = "RequestBasedUsage"

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}