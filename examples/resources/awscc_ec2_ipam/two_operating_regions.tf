resource "awscc_ec2_ipam" "example" {
  description = "example IPAM"

  operating_regions = [
    {
      region_name = "us-east-1"
    },
    {
      region_name = "us-west-2"
    }
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}