# Get account ID and region
data "aws_caller_identity" "current" {}
# Note: Using data.aws_region.current.region (AWS provider v6.0+)
# For AWS provider < v6.0, use data.aws_region.current.name instead
data "aws_region" "current" {}

# Create policy statement for entity resolution workflow
resource "awscc_entityresolution_policy_statement" "example" {
  arn          = "arn:aws:entityresolution:${data.aws_region.current.region}:${data.aws_caller_identity.current.account_id}:matchingworkflow/example-workflow"
  statement_id = "AllowEntityResolutionAccess"
  action       = ["entityresolution:StartMatching", "entityresolution:GetMatching"]
  effect       = "Allow"
  principal = [
    data.aws_caller_identity.current.account_id
  ]
}