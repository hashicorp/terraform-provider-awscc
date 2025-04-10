# Create an EventBridge bus first
resource "awscc_events_event_bus" "example" {
  name = "example-bus"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create an archive for the event bus
resource "awscc_events_archive" "example" {
  archive_name = "example-archive"
  description  = "Example EventBridge archive created with AWSCC provider"
  source_arn   = awscc_events_event_bus.example.arn
  event_pattern = jsonencode({
    "source" : ["aws.ec2"],
    "detail-type" : ["EC2 Instance State-change Notification"]
  })
  retention_days = 30
}