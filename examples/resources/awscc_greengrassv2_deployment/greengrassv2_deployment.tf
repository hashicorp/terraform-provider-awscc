resource "awscc_greengrassv2_deployment" "example" {
  deployment_name = "example"
  target_arn      = awscc_iot_thing.example.arn
  deployment_policies = {
    configuration_validation_policy = {
      timeout_in_seconds = 60
    }
    component_update_policy = {
      timeout_in_seconds = 55
    }
    failure_handling_policy = "ROLLBACK"
  }

  components = {
    "aws.greengrass.Nucleus" = {
      component_version = "2.12.4"
      configuration_update = {
        reset = [""]
        merge = jsonencode(
          {
            "spooler" : {
              "storageType" : "Disk"
              "maxSizeInBytes" : "2048"
            }
          }
        )
      }
    },
    "aws.greengrass.LogManager" = {
      component_version = "2.3.7"
      configuration_update = {
        reset = [""]
        merge = jsonencode(
          {
            "logsUploaderConfiguration" : {
              "systemLogsConfiguration" : {
                "uploadToCloudWatch" : "true",
                "minimumLogLevel" : "INFO",
                "diskSpaceLimit" : "10",
                "diskSpaceLimitUnit" : "MB",
                "deleteLogFileAfterCloudUpload" : "false"
              },
              "componentLogsConfigurationMap" : {
                "aws.greengrass.SystemsManagerAgent" : {
                  "minimumLogLevel" : "INFO",
                  "diskSpaceLimit" : "20",
                  "diskSpaceLimitUnit" : "MB",
                  "deleteLogFileAfterCloudUpload" : "false"
                }
              }
            },
            "periodicUploadIntervalSec" : "300",
            "deprecatedVersionSupport" : "false"
          }
        )
      }
    }
  }

}


resource "awscc_iot_thing" "example" {
  thing_name = "example"
  attribute_payload = {
    attributes = {
      name = "examplevalue"
    }
  }
}