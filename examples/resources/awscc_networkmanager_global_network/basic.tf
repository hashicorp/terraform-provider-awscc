resource "awscc_networkmanager_global_network" "example" {
  description = "example"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}