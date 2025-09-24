# Global Network
resource "awscc_networkmanager_global_network" "example" {
  description = "Example Global Network for Link"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Site
resource "awscc_networkmanager_site" "example" {
  global_network_id = awscc_networkmanager_global_network.example.id
  description       = "Example Site for Link"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Link
resource "awscc_networkmanager_link" "example" {
  global_network_id = awscc_networkmanager_global_network.example.id
  site_id           = awscc_networkmanager_site.example.site_id
  description       = "Example Network Link"
  type              = "Broadband"
  provider_name     = "Example ISP"

  bandwidth = {
    download_speed = 100
    upload_speed   = 100
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}