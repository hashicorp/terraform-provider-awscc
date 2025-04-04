# Create IoT Wireless Service Profile
resource "awscc_iotwireless_service_profile" "example" {
  name = "example-service-profile"
  lo_ra_wan = {
    add_gw_metadata = true
    pr_allowed      = true
    ra_allowed      = false
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}