resource "awscc_ec2_dhcp_options" "example" {
  domain_name          = "example.internal"
  domain_name_servers  = ["AmazonProvidedDNS"]
  ntp_servers          = ["169.254.169.123"]
  netbios_name_servers = ["192.168.1.1", "192.168.1.2"]
  netbios_node_type    = 2 # H-node

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}