resource "awscc_ec2_ipam_resource_discovery_association" "example" {
  ipam_id                    = awscc_ec2_ipam.example.id
  ipam_resource_discovery_id = awscc_ec2_ipam_resource_discovery.example.id
}