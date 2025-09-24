resource "awscc_logs_log_group" "first" {
  log_group_name    = "SampleLogGroup_1"
  retention_in_days = 90
  tags = [
    {
      key   = "Name"
      value = "SampleLogGroup_1"
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

resource "awscc_logs_log_group" "second" {
  log_group_name    = "SampleLogGroup_2"
  retention_in_days = 90
  tags = [
    {
      key   = "Name"
      value = "SampleLogGroup_2"
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

resource "awscc_logs_query_definition" "this" {
  name         = "custom_query"
  query_string = <<EOF
fields @timestamp, @message
| sort @timestamp desc
| limit 25
EOF

  log_group_names = [
    awscc_logs_log_group.first.id,
    awscc_logs_log_group.second.id
  ]
}