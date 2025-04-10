# Create CloudWatch Log Group as destination
resource "awscc_logs_log_group" "resolver_query_logs" {
  log_group_name    = "/aws/route53resolver/query-logs"
  retention_in_days = 14

  tags = [{
    key   = "Name"
    value = "Route53 Resolver Query Logs"
  }]
}

# IAM Role for Route53 Resolver
resource "awscc_iam_role" "resolver_query_logging" {
  role_name = "route53-resolver-query-logging"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "route53resolver.amazonaws.com"
        }
      }
    ]
  })

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# IAM Policy for Route53 Resolver to write logs
data "aws_iam_policy_document" "resolver_logging" {
  statement {
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents"
    ]
    effect = "Allow"
    resources = [
      "${awscc_logs_log_group.resolver_query_logs.arn}:*"
    ]
  }
}

resource "awscc_iam_role_policy" "resolver_logging" {
  policy_name     = "route53-resolver-query-logging"
  role_name       = awscc_iam_role.resolver_query_logging.role_name
  policy_document = data.aws_iam_policy_document.resolver_logging.json
}

# Route53 Resolver Query Logging Config
resource "awscc_route53resolver_resolver_query_logging_config" "example" {
  name            = "example-query-logging"
  destination_arn = awscc_logs_log_group.resolver_query_logs.arn
}