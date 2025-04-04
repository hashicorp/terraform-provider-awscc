# Create IoT Thing for Greengrass V2 Core Device
resource "awscc_iot_thing" "greengrass_core" {
  thing_name = "sitewise-gateway-core-example-2024"
}

# Create IoT SiteWise Gateway
resource "awscc_iotsitewise_gateway" "example" {
  gateway_name = "example-gateway-2024"
  gateway_platform = {
    greengrass_v2 = {
      core_device_thing_name = awscc_iot_thing.greengrass_core.thing_name
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}