resource "awscc_ec2_placement_group" "web" {
  strategy     = "spread"
  spread_level = "host"
  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}