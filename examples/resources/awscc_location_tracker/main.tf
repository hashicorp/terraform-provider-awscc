# Location Tracker
resource "awscc_location_tracker" "example" {
  tracker_name = "example-tracker"
  description  = "Example Location Tracker"

  position_filtering = "TimeBased"
  pricing_plan       = "RequestBasedUsage"

  tags = [{
    key   = "Modified_By"
    value = "AWSCC"
  }]
}