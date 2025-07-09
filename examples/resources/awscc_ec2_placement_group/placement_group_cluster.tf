resource "awscc_ec2_placement_group" "web" {
  strategy = "cluster"
  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}