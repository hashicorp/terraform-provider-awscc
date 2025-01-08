# Get current AWS region and account ID
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

# Create a schedule group
resource "awscc_scheduler_schedule_group" "example" {
  name = "my-schedule-group"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Show outputs
output "schedule_group_arn" {
  value = awscc_scheduler_schedule_group.example.arn
}