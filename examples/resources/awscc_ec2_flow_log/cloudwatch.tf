resource "awscc_ec2_flow_log" "example" {
  deliver_logs_permission_arn = awscc_iam_role.example.arn
  log_destination_type        = "cloud-watch-logs"
  log_destination             = awscc_logs_log_group.example.arn
  traffic_type                = "ALL"
  resource_id                 = "vpc-07ddade55bee92f5f"
  resource_type               = "VPC"
  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}

resource "awscc_logs_log_group" "example" {
  log_group_name = "example"
}

resource "awscc_iam_role" "example" {
  role_name = "cloudwatch_flow_log_role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "vpc-flow-logs.amazonaws.com"
        }
      },
    ]
  })
}

resource "awscc_iam_role_policy" "example" {
  policy_name = "example"
  role_name   = awscc_iam_role.example.role_name
  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "logs:CreateLogStream",
          "logs:DescribeLogGroups",
          "logs:DescribeLogStreams",
          "logs:PutLogEvents",
        ]
        Effect   = "Allow"
        Resource = "${awscc_logs_log_group.example.arn}:*"
      },
    ]
  })
}
