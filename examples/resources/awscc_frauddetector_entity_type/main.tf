# Data sources for AWS account ID and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create the Fraud Detector Entity Type
resource "awscc_frauddetector_entity_type" "example" {
  name        = "customer_account"
  description = "Entity type for customer accounts"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}