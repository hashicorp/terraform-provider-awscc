# Data sources to get the AWS account ID and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# B2BI Profile resource
resource "awscc_b2bi_profile" "example" {
  name          = "example-b2bi-profile"
  business_name = "Example Business"
  phone         = "+1-555-0123"
  logging       = "ENABLED"
  email         = "example@example.com" # Optional

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}