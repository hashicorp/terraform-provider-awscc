# Create a global network first
resource "awscc_networkmanager_global_network" "example" {
  description = "Example Global Network for Device"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create a network manager device
resource "awscc_networkmanager_device" "example" {
  global_network_id = awscc_networkmanager_global_network.example.id
  description       = "Example Network Manager Device"
  model             = "example-model"
  serial_number     = "123456789"
  type              = "HARDWARE"
  vendor            = "Example Vendor"

  location = {
    address   = "123 Example Street"
    latitude  = "47.6062"
    longitude = "-122.3321"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}