resource "awscc_wafv2_logging_configuration" "awscc_waf_logging_redacted_fields" {
  resource_arn            = aws_wafv2_web_acl.example.arn
  log_destination_configs = [aws_cloudwatch_log_group.example.arn]
  redacted_fields = [{
    single_header = {
      name = "authorization"
    }
  }]
}
