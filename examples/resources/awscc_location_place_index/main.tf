# Get current AWS account information
data "aws_caller_identity" "current" {}

# Get current AWS region
data "aws_region" "current" {}

# Create location place index
resource "awscc_location_place_index" "example" {
  data_source = "Esri"
  index_name  = "location-place-index-example"

  data_source_configuration = {
    intended_use = "SingleUse"
  }

  description  = "Example Location place index using AWSCC provider"
  pricing_plan = "RequestBasedUsage"

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}