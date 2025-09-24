# Create a schedule group
resource "awscc_scheduler_schedule_group" "example" {
  name = "my-schedule-group"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}