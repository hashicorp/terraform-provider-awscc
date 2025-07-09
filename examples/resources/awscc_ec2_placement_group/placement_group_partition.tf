resource "awscc_ec2_placement_group" "web" {
  strategy        = "partition"
  partition_count = 2
  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}