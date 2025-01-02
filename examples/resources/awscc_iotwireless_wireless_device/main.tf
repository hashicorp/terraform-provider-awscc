# Data sources for AWS account ID and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create a Destination for IoT Wireless
resource "awscc_iotwireless_destination" "example" {
  name            = "example-destination"
  expression      = "arn:aws:iot:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:topic/my-topic"
  expression_type = "MqttTopic"
  description     = "IoT Wireless Destination for Example"
  role_arn        = awscc_iam_role.iot_wireless_role.arn

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create IAM role for IoT Wireless
resource "awscc_iam_role" "iot_wireless_role" {
  role_name = "iot-wireless-role"
  path      = "/"

  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Principal = {
          Service = "iotwireless.amazonaws.com"
        }
        Action = "sts:AssumeRole"
      }
    ]
  })

  policies = [{
    policy_name = "IoTWirelessDevicePolicy"
    policy_document = jsonencode({
      Version = "2012-10-17"
      Statement = [
        {
          Effect = "Allow"
          Action = [
            "iot:Publish",
            "iot:Subscribe",
            "iot:Connect",
            "iot:Receive"
          ]
          Resource = [
            "arn:aws:iot:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:topic/*",
            "arn:aws:iot:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:client/*"
          ]
        }
      ]
    })
  }]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create IoT Wireless Device
# Note: device_profile_id and service_profile_id must be obtained from existing AWS IoT Wireless profiles
resource "awscc_iotwireless_wireless_device" "example" {
  destination_name = awscc_iotwireless_destination.example.name
  name             = "example-wireless-device"
  type             = "LoRaWAN"

  lo_ra_wan = {
    otaa_v11 = {
      app_key  = "a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4"
      join_eui = "a1b2c3d4e5f6a1b2"
      dev_eui  = "a1b2c3d4e5f6a1b2"
      nwk_key  = "a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4"
    }
    device_profile_id  = "a1b2c3d4-5678-90ab-cdef-EXAMPLE11111" # Replace with actual Device Profile ID
    service_profile_id = "a1b2c3d4-5678-90ab-cdef-EXAMPLE22222" # Replace with actual Service Profile ID
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}