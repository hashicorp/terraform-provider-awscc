# Get current AWS region
data "aws_region" "current" {}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Create an outcome for Fraud Detector
resource "awscc_frauddetector_outcome" "example" {
  name        = "test-outcome"
  description = "Test fraud detector outcome created by AWSCC"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}