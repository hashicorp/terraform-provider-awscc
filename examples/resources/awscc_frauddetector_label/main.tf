# Create a Fraud Detector Label
resource "awscc_frauddetector_label" "example" {
  name        = "example-fraud-label"
  description = "Example fraud label for demonstration"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}