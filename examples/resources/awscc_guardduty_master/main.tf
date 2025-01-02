# Data sources for account and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Get existing detector
data "aws_guardduty_detector" "existing" {}

# Create GuardDuty Master in member account
resource "awscc_guardduty_master" "example" {
  detector_id   = data.aws_guardduty_detector.existing.id
  master_id     = data.aws_caller_identity.current.account_id
  invitation_id = "12345" # Replace with actual invitation ID
}