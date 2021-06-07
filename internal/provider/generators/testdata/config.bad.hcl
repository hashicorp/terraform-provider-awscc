resource_schemer "aws_logs_log_group" {
  source {
    url = "https://raw.githubusercontent.com/aws-cloudformation/aws-cloudformation-resource-providers-logs/8b9229f78b832800b7fb6c1165bcb3893f44b856/aws-logs-loggroup/aws-logs-loggroup.json"
  }

  local = "aws_logs_log_group.cf-resource-schema.json"
}
