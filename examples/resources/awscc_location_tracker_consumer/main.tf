# Create the location tracker first
resource "awscc_location_tracker" "example" {
  tracker_name = "example-tracker"
  description  = "Example tracker for tracker consumer demonstration"

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}

# Create a geofence collection
resource "awscc_location_geofence_collection" "example" {
  collection_name = "example-collection"
  description     = "Example geofence collection"

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}

# Create the tracker consumer that connects the tracker to the geofence collection
resource "awscc_location_tracker_consumer" "example" {
  tracker_name = awscc_location_tracker.example.tracker_name
  consumer_arn = awscc_location_geofence_collection.example.collection_arn
}