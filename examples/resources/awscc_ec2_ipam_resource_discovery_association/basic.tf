resource "awscc_ec2_ipam_resource_discovery_association" "example" {
  ipam_id                    = awscc_ec2_ipam.example.id
  ipam_resource_discovery_id = awscc_ec2_ipam_resource_discovery.example.id
}

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
}

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
}