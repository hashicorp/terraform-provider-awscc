resource "awscc_logs_metric_filter" "this" {
  filter_name    = "SamplePattern"
  filter_pattern = ""
  log_group_name = awscc_logs_log_group.this.id
  metric_transformations = [{
    metric_name      = "EventCount"
    metric_namespace = "YourNamespace"
    metric_value     = "1"
  }]
}

resource "awscc_logs_log_group" "this" {
  log_group_name    = "SampleLogGroup"
  retention_in_days = 90
  tags = [
    {
      key   = "Name"
      value = "SampleLogGroup"
    },
    {
      key   = "Environment"
      value = "Development"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}