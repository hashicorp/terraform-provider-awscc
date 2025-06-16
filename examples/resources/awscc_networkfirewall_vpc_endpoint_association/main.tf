# Example of Network Firewall VPC Endpoint Association configuration
resource "awscc_networkfirewall_vpc_endpoint_association" "example" {
  # The ARN of an existing Network Firewall
  firewall_arn = "arn:aws:network-firewall:us-west-2:123456789012:firewall/example-firewall"

  # The ID of an existing VPC
  vpc_id = "vpc-1234567890abcdef0"

  # The subnet mapping configuration
  subnet_mapping = {
    subnet_id       = "subnet-1234567890abcdef0"
    ip_address_type = "IPV4"
  }

  description = "Example VPC endpoint association"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}