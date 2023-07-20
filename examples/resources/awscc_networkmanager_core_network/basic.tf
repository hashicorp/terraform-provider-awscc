resource "awscc_networkmanager_global_network" "example" {}

resource "awscc_networkmanager_core_network" "example" {
  global_network_id = awscc_networkmanager_global_network.example.id
  description       = "example"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}