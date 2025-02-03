# Data sources for AWS region and account ID
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

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