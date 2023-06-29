resource "awscc_ec2_ipam" "example" {
  description = "example IPAM"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}