# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Get current AWS region
data "aws_region" "current" {}

# Create EBS snapshot block public access
resource "awscc_ec2_snapshot_block_public_access" "example" {
  state = "block-all-sharing" # block all sharing of EBS snapshots
}