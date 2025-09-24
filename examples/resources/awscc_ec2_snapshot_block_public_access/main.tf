# Create EBS snapshot block public access
resource "awscc_ec2_snapshot_block_public_access" "example" {
  state = "block-all-sharing" # block all sharing of EBS snapshots
}