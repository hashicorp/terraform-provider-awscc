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