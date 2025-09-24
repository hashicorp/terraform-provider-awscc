resource "awscc_location_geofence_collection" "example" {
  collection_name = "example"
  description     = "Example geofence collection"
  pricing_plan    = "RequestBasedUsage"
  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}
