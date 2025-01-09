# Data sources to dynamically fetch AWS account ID and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create an SES Mail Manager Addon Subscription
resource "awscc_ses_mail_manager_addon_subscription" "example" {
  addon_name = "SPAMHAUS_DBL"

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}

# Output the subscription ARN
output "addon_subscription_arn" {
  value = awscc_ses_mail_manager_addon_subscription.example.addon_subscription_arn
}