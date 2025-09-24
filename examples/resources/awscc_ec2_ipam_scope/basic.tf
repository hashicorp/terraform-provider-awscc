resource "awscc_ec2_ipam_scope" "example" {
  ipam_id = awscc_ec2_ipam.example.id

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}