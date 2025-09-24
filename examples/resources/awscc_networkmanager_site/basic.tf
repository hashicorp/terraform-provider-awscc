resource "awscc_networkmanager_site" "example" {
  global_network_id = awscc_networkmanager_global_network.example.id
  description       = "example"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}