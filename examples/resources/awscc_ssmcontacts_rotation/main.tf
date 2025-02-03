# Get current AWS account ID and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Example SSM Contact Rotation
resource "awscc_ssmcontacts_rotation" "example" {
  name         = "weekly_oncall_rotation"
  contact_ids  = ["arn:aws:ssm-contacts:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:contact/example-contact"]
  time_zone_id = "America/Los_Angeles"
  start_time   = "2025-01-05T09:00:00"

  recurrence = {
    weekly_settings = [{
      day_of_week   = "MON"
      hand_off_time = "09:00"
    }]
    number_of_on_calls    = 1
    recurrence_multiplier = 1
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}