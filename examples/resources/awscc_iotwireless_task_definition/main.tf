# IAM role for the task definition
data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["iotwireless.amazonaws.com"]
    }
  }
}

resource "awscc_iam_role" "iotwireless_task_role" {
  assume_role_policy_document = data.aws_iam_policy_document.assume_role.json

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# IoT Wireless Task Definition
resource "awscc_iotwireless_task_definition" "example" {
  name                 = "example-task-definition"
  auto_create_tasks    = true
  task_definition_type = "UPDATE"

  update = {
    update_data_role   = awscc_iam_role.iotwireless_task_role.arn
    update_data_source = "s3://example-bucket/firmware.bin"
    lo_ra_wan = {
      current_version = {
        model           = "example-model"
        package_version = "1.0.0"
        station         = "EXAMPLE-STATION"
      }
      update_version = {
        model           = "example-model"
        package_version = "1.0.1"
        station         = "EXAMPLE-STATION"
      }
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}