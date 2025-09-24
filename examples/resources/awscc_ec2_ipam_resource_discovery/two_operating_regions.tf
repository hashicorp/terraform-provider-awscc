resource "awscc_ec2_ipam_resource_discovery" "example" {
  description = "example resource discovery"

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