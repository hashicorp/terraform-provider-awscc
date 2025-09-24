# Create global network
resource "awscc_networkmanager_global_network" "example" {
  description = "Example Global Network"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create site with required location information
resource "awscc_networkmanager_site" "example" {
  global_network_id = awscc_networkmanager_global_network.example.id
  description       = "Example Site"
  location = {
    address   = "123 Example Street"
    latitude  = "37.7749"   # San Francisco latitude
    longitude = "-122.4194" # San Francisco longitude
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Wait for site to be available
resource "time_sleep" "wait_30_seconds" {
  depends_on      = [awscc_networkmanager_site.example]
  create_duration = "30s"
}

# Create device
resource "awscc_networkmanager_device" "example" {
  depends_on        = [time_sleep.wait_30_seconds]
  global_network_id = awscc_networkmanager_global_network.example.id
  site_id           = awscc_networkmanager_site.example.site_id
  description       = "Example Device"
  model             = "Example Model"
  vendor            = "Example Vendor"
  type              = "GENERIC"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create link
resource "awscc_networkmanager_link" "example" {
  depends_on        = [time_sleep.wait_30_seconds]
  global_network_id = awscc_networkmanager_global_network.example.id
  site_id           = awscc_networkmanager_site.example.site_id
  bandwidth = {
    download_speed = 100
    upload_speed   = 100
  }
  description = "Example Link"
  type        = "BROADBAND"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create link association
resource "awscc_networkmanager_link_association" "example" {
  depends_on = [
    awscc_networkmanager_device.example,
    awscc_networkmanager_link.example
  ]
  global_network_id = awscc_networkmanager_global_network.example.id
  device_id         = awscc_networkmanager_device.example.device_id
  link_id           = awscc_networkmanager_link.example.link_id
}