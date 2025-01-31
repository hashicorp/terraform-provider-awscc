data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# S3 bucket for datastore
resource "awscc_s3_bucket" "datastore" {
  bucket_name = "iot-analytics-datastore-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"
}

# IAM role for IoT Analytics
data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["iotanalytics.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "datastore" {
  statement {
    actions = [
      "s3:GetBucketLocation",
      "s3:GetObject",
      "s3:ListBucket",
      "s3:PutObject",
      "s3:DeleteObject"
    ]
    resources = [
      awscc_s3_bucket.datastore.arn,
      "${awscc_s3_bucket.datastore.arn}/*"
    ]
  }
}

resource "awscc_iam_role" "datastore" {
  role_name                   = "IoTAnalyticsDatastoreRole"
  assume_role_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.assume_role.json))
  policies = [{
    policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.datastore.json))
    policy_name     = "IoTAnalyticsDatastoreAccess"
  }]
}

# IoT Analytics Datastore
resource "awscc_iotanalytics_datastore" "example" {
  datastore_name = "example-datastore"

  datastore_storage = {
    customer_managed_s3 = {
      bucket     = awscc_s3_bucket.datastore.id
      key_prefix = "data/"
      role_arn   = awscc_iam_role.datastore.arn
    }
  }

  retention_period = {
    number_of_days = 90
    unlimited      = false
  }

  file_format_configuration = {
    parquet_configuration = {
      schema_definition = {
        columns = [
          {
            name = "device_id"
            type = "STRING"
          },
          {
            name = "timestamp"
            type = "STRING"
          },
          {
            name = "temperature"
            type = "DOUBLE"
          }
        ]
      }
    }
  }

  tags = [{
    key   = "Environment"
    value = "Production"
  }]
}