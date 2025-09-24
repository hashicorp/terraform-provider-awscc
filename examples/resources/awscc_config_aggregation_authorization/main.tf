# Get current AWS region
data "aws_region" "current" {}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# AWS Config Aggregation Authorization
resource "awscc_config_aggregation_authorization" "example" {
  authorized_account_id = data.aws_caller_identity.current.account_id
  authorized_aws_region = data.aws_region.current.name

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}