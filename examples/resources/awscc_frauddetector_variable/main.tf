# Create a Fraud Detector Variable for email address
resource "awscc_frauddetector_variable" "email_variable" {
  name          = "email_address"
  data_source   = "EVENT"
  data_type     = "STRING"
  default_value = "unknown@example.com"
  description   = "Customer email address for fraud detection"
  variable_type = "EMAIL_ADDRESS"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}