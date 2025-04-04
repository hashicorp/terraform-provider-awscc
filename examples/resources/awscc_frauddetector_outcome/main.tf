# Create an outcome for Fraud Detector
resource "awscc_frauddetector_outcome" "example" {
  name        = "test-outcome"
  description = "Test fraud detector outcome created by AWSCC"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}