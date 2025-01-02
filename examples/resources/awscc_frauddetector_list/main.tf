data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create a Fraud Detector List
resource "awscc_frauddetector_list" "example" {
  name          = "example_fraud_list"
  description   = "Example fraud detection list"
  variable_type = "CATEGORICAL"
  elements = [
    "example1",
    "example2",
    "example3"
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}