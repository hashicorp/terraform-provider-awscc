# Get current AWS region
data "aws_region" "current" {}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Create a Fraud Detector Label
resource "awscc_frauddetector_label" "example" {
  name        = "example-fraud-label"
  description = "Example fraud label for demonstration"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}