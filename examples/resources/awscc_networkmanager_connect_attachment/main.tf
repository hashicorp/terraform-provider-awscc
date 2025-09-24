data "aws_region" "current" {}

# Example Connect Attachment - This is an example template, not intended for direct use
# since it requires existing core network, transport attachment, and VPC resources
resource "awscc_networkmanager_connect_attachment" "example" {
  # The ID of the core network to attach to - Replace with your actual core network ID
  core_network_id = "core-network-xxxx"

  # The ID of existing transport attachment (e.g., VPC attachment) - Replace with actual attachment ID
  transport_attachment_id = "attachment-xxxx"

  # The edge location for the attachment (AWS Region)
  edge_location = data.aws_region.current.name

  # Configuration options for the connect attachment
  options = {
    protocol = "GRE" # The tunnel protocol for the connect attachment
  }

  # Optional tags for the attachment
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}