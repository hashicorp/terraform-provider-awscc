# Create an SES Mail Manager Addon Subscription
resource "awscc_ses_mail_manager_addon_subscription" "example" {
  addon_name = "SPAMHAUS_DBL"

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}