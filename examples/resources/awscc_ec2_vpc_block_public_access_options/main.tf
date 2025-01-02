# Example resource for AWS VPC Block Public Access Options
resource "awscc_ec2_vpc_block_public_access_options" "example" {
  # Valid values are "block-bidirectional" or "block-ingress"
  internet_gateway_block_mode = "block-bidirectional"
}