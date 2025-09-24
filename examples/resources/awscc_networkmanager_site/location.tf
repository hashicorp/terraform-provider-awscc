resource "awscc_networkmanager_site" "example" {
  global_network_id = awscc_networkmanager_global_network.example.id
  description       = "example"

  location = {
    address   = "Antarctica"
    latitude  = "82.8628"
    longitude = "135.0000"
  }
}