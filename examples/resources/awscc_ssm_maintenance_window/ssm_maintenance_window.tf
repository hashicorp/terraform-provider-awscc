resource "awscc_ssm_maintenance_window" "example" {
  name              = "example-maintenance-window"
  description       = "Example maintenance window for demonstration"
  schedule          = "cron(0 02 ? * SUN *)"
  duration          = 2
  cutoff            = 1
  allow_unassociated_targets = false

  tags = [
    {
      key   = "Environment"
      value = "Example"
    }
  ]
}