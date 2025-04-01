# Prerequisites:
# 1. AWS Organizations must be enabled for your AWS account
# 2. AWS Firewall Manager must be enabled in your organization
# 3. AWS Firewall Manager administrator account must be designated

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Example FMS Security Policy
resource "awscc_fms_policy" "example" {
  policy_name           = "example-fms-policy"
  remediation_enabled   = true
  exclude_resource_tags = false

  security_service_policy_data = {
    type = "WAF"
    managed_service_data = jsonencode({
      type = "WAF"
      ruleGroups = [{
        id = "default"
        overrideAction = {
          type = "COUNT"
        }
      }]
      defaultAction = {
        type = "BLOCK"
      }
      overrideCustomerWebACLAssociation = false
    })
  }

  resource_type = "AWS::ElasticLoadBalancingV2::LoadBalancer"

  policy_description = "Example FMS WAF Policy"

  # Example include map - you can include specific accounts or OUs
  include_map = {
    account = [data.aws_caller_identity.current.account_id]
  }

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}