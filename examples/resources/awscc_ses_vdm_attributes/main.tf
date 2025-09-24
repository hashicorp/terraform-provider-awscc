# Configure VDM (Virtual Deliverability Manager) attributes for SES
# This enables engagement tracking and optimized shared delivery features
# Create VDM attributes for SES
resource "awscc_ses_vdm_attributes" "example" {
  dashboard_attributes = {
    engagement_metrics = "ENABLED"
  }
  guardian_attributes = {
    optimized_shared_delivery = "ENABLED"
  }
}