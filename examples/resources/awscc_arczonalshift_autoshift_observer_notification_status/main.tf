# Get current AWS account ID and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create the AutoshiftObserverNotificationStatus
resource "awscc_arczonalshift_autoshift_observer_notification_status" "example" {
  status = "ENABLED"
}