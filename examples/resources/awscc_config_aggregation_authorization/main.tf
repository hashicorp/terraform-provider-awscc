# Get current AWS region
# Note: Using data.aws_region.current.region (AWS provider v6.0+)
# For AWS provider < v6.0, use data.aws_region.current.name instead
data "aws_region" "current" {}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# AWS Config Aggregation Authorization
resource "awscc_config_aggregation_authorization" "example" {
  authorized_account_id = data.aws_caller_identity.current.account_id
  authorized_aws_region = data.aws_region.current.region

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}