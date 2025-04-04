# Create Certificate Manager Account resource
resource "awscc_certificatemanager_account" "example" {
  expiry_events_configuration = {
    days_before_expiry = 30
  }
}