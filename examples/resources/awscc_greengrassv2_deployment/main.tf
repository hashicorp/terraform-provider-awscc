# Data sources for dynamic values
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# IoT Thing
resource "awscc_iot_thing" "example" {
  thing_name = "example-greengrass-core"
  attribute_payload = {
    attributes = {
      type = "greengrass-core"
    }
  }
}

# IoT Thing Group for deployment target
resource "awscc_iot_thing_group" "example" {
  thing_group_name = "example-greengrass-group"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Main Greengrass V2 deployment
resource "awscc_greengrassv2_deployment" "example" {
  deployment_name = "example-deployment"
  target_arn      = awscc_iot_thing_group.example.arn

  deployment_policies = {
    configuration_validation_policy = {
      timeout_in_seconds = 60
    }
    component_update_policy = {
      timeout_in_seconds = 120
      action             = "NOTIFY_COMPONENTS"
    }
    failure_handling_policy = "ROLLBACK"
  }

  components = {
    "aws.greengrass.Nucleus" = {
      component_version = "2.12.4"
      configuration_update = {
        reset = [""]
        merge = jsonencode({
          jvmOptions = ["-Xmx64m"]
          mqtt = {
            spooler = {
              storageType    = "Disk"
              maxSizeInBytes = "10485760"
              keepQos0       = true
            }
          }
        })
      }
    }
    "aws.greengrass.LogManager" = {
      component_version = "2.3.7"
      configuration_update = {
        reset = [""]
        merge = jsonencode({
          logsUploaderConfiguration = {
            systemLogsConfiguration = {
              uploadToCloudWatch            = "true"
              minimumLogLevel               = "INFO"
              diskSpaceLimit                = "10"
              diskSpaceLimitUnit            = "MB"
              deleteLogFileAfterCloudUpload = "false"
            }
          }
          periodicUploadIntervalSec = "300"
        })
      }
    }
  }

  tags = {
    "Modified By" = "AWSCC"
  }
}