# Create an example IoT Wireless Gateway
resource "awscc_iotwireless_wireless_gateway" "example" {
  description = "Example LoRaWAN Gateway"
  name        = "example-lorawan-gateway"

  lo_ra_wan = {
    gateway_eui = "A1B2C3D4E5F6A1B2"
    rf_region   = "US915"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}