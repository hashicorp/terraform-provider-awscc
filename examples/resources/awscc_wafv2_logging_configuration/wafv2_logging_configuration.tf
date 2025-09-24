resource "awscc_wafv2_logging_configuration" "awscc_waf_logging" {
  resource_arn            = aws_wafv2_web_acl.example.arn
  log_destination_configs = [aws_cloudwatch_log_group.example.arn]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "aws_cloudwatch_log_group" "example" {
  name = "example"
}

resource "aws_wafv2_web_acl" "example" {
  name        = "managed-rule-example"
  description = "Example of a managed rule."
  scope       = "REGIONAL"

  default_action {
    block {}
  }

  rule {
    name     = "AWS-AWSManagedRulesCommonRuleSet"
    priority = 1

    override_action {
      none {}
    }

    statement {
      managed_rule_group_statement {
        name        = "AWSManagedRulesCommonRuleSet"
        vendor_name = "AWS"
      }
    }

    visibility_config {
      cloudwatch_metrics_enabled = true
      metric_name                = "AWS-AWSManagedRulesCommonRuleSet"
      sampled_requests_enabled   = true
    }
  }

  visibility_config {
    cloudwatch_metrics_enabled = true
    metric_name                = "ExternalACL"
    sampled_requests_enabled   = true
  }
}



