# Example awscc_ses_mail_manager_addon_instance resource
resource "awscc_ses_mail_manager_addon_instance" "example" {
  addon_subscription_id = "as-abc123def456" # This should be replaced with a valid subscription ID

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}