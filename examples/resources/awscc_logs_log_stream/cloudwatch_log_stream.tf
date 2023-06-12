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

resource "awscc_logs_log_stream" "this" {
  log_group_name  = awscc_logs_log_group.this.id
  log_stream_name = "SampleStream"
}