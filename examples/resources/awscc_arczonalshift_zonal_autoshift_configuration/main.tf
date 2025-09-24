# Get current AWS region
data "aws_region" "current" {}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

resource "awscc_arczonalshift_zonal_autoshift_configuration" "example" {
  resource_identifier    = "example-autoshift-config"
  zonal_autoshift_status = "ENABLED"

  practice_run_configuration = {
    blocked_dates   = ["2025-01-01"]          # Example blocked date
    blocked_windows = ["Mon:23:00-Tue:01:00"] # Example blocked window

    blocking_alarms = [
      {
        alarm_identifier = "arn:aws:cloudwatch:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:alarm:example-blocking-alarm"
        type             = "CloudWatch"
      }
    ]

    outcome_alarms = [
      {
        alarm_identifier = "arn:aws:cloudwatch:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:alarm:example-outcome-alarm"
        type             = "CloudWatch"
      }
    ]
  }
}