# Get AWS Account ID
data "aws_caller_identity" "current" {}

# Get current AWS Region
data "aws_region" "current" {}

# Create a Cost anomaly monitor first
resource "awscc_ce_anomaly_monitor" "example" {
  monitor_name      = "example-monitor"
  monitor_type      = "DIMENSIONAL"
  monitor_dimension = "SERVICE"
  resource_tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the anomaly subscription
resource "awscc_ce_anomaly_subscription" "example" {
  subscription_name = "example-subscription"
  frequency         = "DAILY"
  monitor_arn_list  = [awscc_ce_anomaly_monitor.example.monitor_arn]
  threshold         = 100
  subscribers = [{
    type    = "EMAIL"
    address = "example@example.com"
  }]
  resource_tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}