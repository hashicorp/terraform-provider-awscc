# Get current AWS account and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create Certificate Manager Account resource
resource "awscc_certificatemanager_account" "example" {
  expiry_events_configuration = {
    days_before_expiry = 30
  }
}