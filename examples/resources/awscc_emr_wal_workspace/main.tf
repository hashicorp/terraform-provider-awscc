# Get the current AWS account ID and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create an EMR WAL workspace
resource "awscc_emr_wal_workspace" "example" {
  wal_workspace_name = "examplewalworkspace"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}