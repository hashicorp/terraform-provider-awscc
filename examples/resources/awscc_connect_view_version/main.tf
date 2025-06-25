data "aws_caller_identity" "current" {}
# Note: Using data.aws_region.current.region (AWS provider v6.0+)
# For AWS provider < v6.0, use data.aws_region.current.name instead
data "aws_region" "current" {}

resource "awscc_connect_view_version" "example" {
  view_arn            = "arn:aws:connect:${data.aws_region.current.region}:${data.aws_caller_identity.current.account_id}:instance/REPLACE_WITH_YOUR_INSTANCE_ID/view/REPLACE_WITH_YOUR_VIEW_ID"
  version_description = "Initial version"
}