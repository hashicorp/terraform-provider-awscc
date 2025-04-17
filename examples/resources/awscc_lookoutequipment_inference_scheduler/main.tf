data "aws_caller_identity" "current" {}

# S3 bucket for input data
resource "awscc_s3_bucket" "input_bucket" {
  bucket_name = "lookout-equipment-input-${data.aws_caller_identity.current.account_id}"
}

# S3 bucket for output data
resource "awscc_s3_bucket" "output_bucket" {
  bucket_name = "lookout-equipment-output-${data.aws_caller_identity.current.account_id}"
}

# IAM role for Lookout for Equipment
data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    effect  = "Allow"

    principals {
      type        = "Service"
      identifiers = ["lookoutequipment.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "lookout_equipment" {
  statement {
    effect = "Allow"
    actions = [
      "s3:ListBucket"
    ]
    resources = [
      "arn:aws:s3:::${awscc_s3_bucket.input_bucket.bucket_name}",
      "arn:aws:s3:::${awscc_s3_bucket.output_bucket.bucket_name}"
    ]
  }
  statement {
    effect = "Allow"
    actions = [
      "s3:GetObject",
      "s3:PutObject"
    ]
    resources = [
      "arn:aws:s3:::${awscc_s3_bucket.input_bucket.bucket_name}/*",
      "arn:aws:s3:::${awscc_s3_bucket.output_bucket.bucket_name}/*"
    ]
  }
}

resource "awscc_iam_role" "lookout_equipment_role" {
  role_name = "LookoutEquipmentInferenceRole"
  assume_role_policy_document = jsonencode(
    jsondecode(data.aws_iam_policy_document.assume_role.json)
  )
  policies = [
    {
      policy_document = jsonencode(
        jsondecode(data.aws_iam_policy_document.lookout_equipment.json)
      )
      policy_name = "LookoutEquipmentS3Access"
    }
  ]
}

# Lookout Equipment Inference Scheduler
resource "awscc_lookoutequipment_inference_scheduler" "example" {
  inference_scheduler_name     = "example-scheduler"
  model_name                   = "example-model"
  role_arn                     = awscc_iam_role.lookout_equipment_role.arn
  data_delay_offset_in_minutes = 15

  data_input_configuration = {
    s3_input_configuration = {
      bucket = awscc_s3_bucket.input_bucket.bucket_name
      prefix = "input/"
    }
    inference_input_name_configuration = {
      component_timestamp_delimiter = "-"
      timestamp_format              = "yyyyMMddHHmmss"
    }
    input_time_zone_offset = "+00:00"
  }

  data_output_configuration = {
    s3_output_configuration = {
      bucket = awscc_s3_bucket.output_bucket.bucket_name
      prefix = "output/"
    }
  }

  data_upload_frequency = "PT5M"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}