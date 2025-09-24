resource "awscc_logs_log_group" "my_log_group" {
  log_group_name    = "my-log-group"
  retention_in_days = 7
}