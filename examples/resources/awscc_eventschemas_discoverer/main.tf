# Create an EventBridge Event Bus
resource "awscc_events_event_bus" "example" {
  name = "example-event-bus"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create an event schema discoverer
resource "awscc_eventschemas_discoverer" "example" {
  source_arn  = awscc_events_event_bus.example.arn
  description = "Example Event Schema Discoverer"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
  cross_account = false
}