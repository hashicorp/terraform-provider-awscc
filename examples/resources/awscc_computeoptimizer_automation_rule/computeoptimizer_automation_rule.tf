resource "awscc_computeoptimizer_automation_rule" "example" {
  name   = "example-automation-rule"
  status = "Active"
  
  rule_type = "AccountRule"
  
  recommended_action_types = [
    "SnapshotAndDeleteUnattachedEbsVolume",
    "UpgradeEbsVolumeType"
  ]
  
  schedule = {
    schedule_expression          = "cron(0 9 ? * MON *)"
    schedule_expression_timezone = "America/Los_Angeles"
    execution_window_in_minutes  = 60
  }
  
  criteria = {
    estimated_monthly_savings = [{
      comparison_operator = "GreaterThanOrEqual"
      threshold_value     = 50
      unit                = "Usd"
    }]
  }
  
  resource_selection_config = {
    resource_selection_criteria = {
      resource_types = [
        "EbsVolume",
        "EbsSnapshot"
      ]
    }
  }

  tags = {
    Environment = "example"
    Purpose     = "automation-rule-demo"
  }
}