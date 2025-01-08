# Get the current AWS region
data "aws_region" "current" {}

# Get the current AWS account ID
data "aws_caller_identity" "current" {}

# Create a WorkSpaces ThinClient Environment
resource "awscc_workspacesthinclient_environment" "example" {
  name        = "example-thinclient-env"
  desktop_arn = "arn:aws:workspaces:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:workspace/ws-abcd1234"

  maintenance_window = {
    type              = "CUSTOM"
    days_of_the_week  = ["MONDAY"]
    start_time_hour   = 2
    start_time_minute = 0
    end_time_hour     = 4
    end_time_minute   = 0
    apply_time_of     = "UTC"
  }

  software_set_update_mode     = "USE_LATEST"
  software_set_update_schedule = "USE_MAINTENANCE_WINDOW"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

  device_creation_tags = [{
    key   = "Environment"
    value = "Production"
  }]
}