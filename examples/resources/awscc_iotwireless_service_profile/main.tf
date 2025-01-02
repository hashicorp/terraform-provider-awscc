# Get current region
data "aws_region" "current" {}

# Get current account ID
data "aws_caller_identity" "current" {}

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