resource "awscc_iam_role" "example" {
  role_name = "example-scheduled-query-role"

  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "events.amazonaws.com"
        }
      }
    ]
  })

  tags = [
    {
      key   = "Name"
      value = "example-scheduled-query-role"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}

resource "awscc_iam_role_policy" "example" {
  policy_name = "example-scheduled-query-policy"
  role_name   = awscc_iam_role.example.role_name

  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "logs:StartQuery",
          "logs:StopQuery",
          "logs:GetQueryResults",
          "logs:DescribeLogGroups",
          "logs:DescribeLogStreams"
        ]
        Resource = "*"
      }
    ]
  })
}

resource "awscc_logs_log_group" "example" {
  log_group_name    = "/aws/example/scheduled-query"
  retention_in_days = 7

  tags = [
    {
      key   = "Name"
      value = "example-scheduled-query-logs"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}

resource "awscc_logs_scheduled_query" "example" {
  name                  = "example-scheduled-query"
  execution_role_arn    = awscc_iam_role.example.arn
  query_language        = "CWLI"
  query_string          = "fields @timestamp, @message | sort @timestamp desc | limit 10"
  schedule_expression   = "cron(0 * * * ? *)"
  start_time_offset     = 300
  log_group_identifiers = [awscc_logs_log_group.example.log_group_name]
  description           = "Example scheduled CloudWatch Logs query"
  state                 = "ENABLED"

  tags = [
    {
      key   = "Name"
      value = "example-scheduled-query"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}