data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

data "aws_iam_policy_document" "pipeline_role_policy" {
  statement {
    actions = [
      "iot:DescribeThingGroup",
      "iot:GetThingShadow",
      "iot:DescribeThing",
      "lambda:InvokeFunction"
    ]
    resources = ["*"]
    effect    = "Allow"
  }
}

data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["iotanalytics.amazonaws.com"]
    }
  }
}

resource "awscc_iam_role" "pipeline_role" {
  role_name                   = "iot-analytics-pipeline-role"
  assume_role_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.assume_role.json))
  policies = [{
    policy_name     = "pipeline-policy"
    policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.pipeline_role_policy.json))
  }]
}

resource "awscc_iotanalytics_channel" "test" {
  channel_name = "test_channel"
  retention_period = {
    unlimited      = false
    number_of_days = 14
  }
}

resource "awscc_iotanalytics_datastore" "test" {
  datastore_name = "test_datastore"
  retention_period = {
    unlimited      = false
    number_of_days = 14
  }
}

resource "awscc_iotanalytics_pipeline" "test" {
  pipeline_name = "test_pipeline"
  pipeline_activities = [
    {
      channel = {
        name         = "ChannelActivity"
        channel_name = awscc_iotanalytics_channel.test.channel_name
        next         = "FilterActivity"
      }
    },
    {
      filter = {
        name   = "FilterActivity"
        filter = "temperature > 25"
        next   = "DeviceRegistryEnrich"
      }
    },
    {
      device_registry_enrich = {
        name       = "DeviceRegistryEnrich"
        attribute  = "metadata"
        role_arn   = awscc_iam_role.pipeline_role.arn
        thing_name = "$${iot:deviceId}"
        next       = "DatastoreActivity"
      }
    },
    {
      datastore = {
        name           = "DatastoreActivity"
        datastore_name = awscc_iotanalytics_datastore.test.datastore_name
      }
    }
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}