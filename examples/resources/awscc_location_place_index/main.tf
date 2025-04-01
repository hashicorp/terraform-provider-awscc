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