# Data sources to get AWS Account ID and Region
data "aws_caller_identity" "current" {}

data "aws_region" "current" {}

# SecurityHub Insight resource
resource "awscc_securityhub_insight" "example" {
  name               = "High-severity-findings-last-30-days"
  group_by_attribute = "ResourceId"

  # Example filter configuration to look for high severity findings in the last 30 days
  filters = {
    severity_label = [{
      comparison = "EQUALS"
      value      = "HIGH"
    }]

    updated_at = [{
      date_range = {
        unit  = "DAYS"
        value = 30
      }
    }]

    aws_account_id = [{
      comparison = "EQUALS"
      value      = data.aws_caller_identity.current.account_id
    }]

    region = [{
      comparison = "EQUALS"
      value      = data.aws_region.current.name
    }]

    record_state = [{
      comparison = "EQUALS"
      value      = "ACTIVE"
    }]

    workflow_status = [{
      comparison = "EQUALS"
      value      = "NEW"
    }]
  }
}

# Output the Insight ARN
output "insight_arn" {
  value = awscc_securityhub_insight.example.insight_arn
}