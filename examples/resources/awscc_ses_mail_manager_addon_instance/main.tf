# Get current AWS account and region information
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Example awscc_ses_mail_manager_addon_instance resource
resource "awscc_ses_mail_manager_addon_instance" "example" {
  addon_subscription_id = "as-abc123def456" # This should be replaced with a valid subscription ID

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}