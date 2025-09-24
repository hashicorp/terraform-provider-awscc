resource "awscc_wafv2_logging_configuration" "awscc_waf_logging_filter" {
  resource_arn            = aws_wafv2_web_acl.example.arn
  log_destination_configs = [aws_cloudwatch_log_group.example.arn]

  logging_filter = {
    default_behavior = "KEEP"

    filters = [{
      behavior = "DROP"
      conditions = [{
        action_condition = {
          action = "BLOCK"
        }
      }]

      requirement = "MEETS_ANY"
    }]
  }

}