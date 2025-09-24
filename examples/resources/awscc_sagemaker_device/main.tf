# Create a SageMaker Device
resource "awscc_sagemaker_device" "example" {
  device_fleet_name = "my-device-fleet"

  device = {
    device_name    = "my-edge-device"
    description    = "Example edge device for SageMaker"
    iot_thing_name = "my-iot-thing"
  }

  tags = [{
    key   = "Environment"
    value = "Production"
    }, {
    key   = "Modified By"
    value = "AWSCC"
  }]
}