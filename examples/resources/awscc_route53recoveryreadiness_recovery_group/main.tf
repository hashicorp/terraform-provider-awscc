# Get current account information
data "aws_caller_identity" "current" {}

# Get current region information
data "aws_region" "current" {}

# First create a cell that we'll reference in the recovery group
resource "awscc_route53recoveryreadiness_cell" "example" {
  cell_name = "example-cell-${data.aws_caller_identity.current.account_id}"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the recovery group
resource "awscc_route53recoveryreadiness_recovery_group" "example" {
  recovery_group_name = "example-recovery-group-${data.aws_caller_identity.current.account_id}"
  cells               = [awscc_route53recoveryreadiness_cell.example.cell_arn]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}