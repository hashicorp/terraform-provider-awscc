data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create an S3 bucket to store job document
resource "awscc_s3_bucket" "job_document" {
  bucket_name = "iot-job-document-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"
}

# Create job document content in S3
resource "aws_s3_object" "job_document" {
  bucket = awscc_s3_bucket.job_document.id
  key    = "job-document.json"
  content = jsonencode({
    operation : "reboot"
    timestamp : "2024-12-31T12:00:00Z"
  })
}

# IAM role for pre-signed URL access
data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["iot.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "s3_access" {
  statement {
    actions = [
      "s3:GetObject"
    ]
    resources = [
      "${awscc_s3_bucket.job_document.arn}/*"
    ]
  }
}

resource "awscc_iam_role" "job_template" {
  assume_role_policy_document = data.aws_iam_policy_document.assume_role.json
  path                        = "/service-role/"
  role_name                   = "IoTJobTemplateRole"
  description                 = "Role for IoT Job Template to access S3"
}

resource "aws_iam_role_policy" "job_template" {
  name   = "IoTJobTemplatePolicy"
  role   = awscc_iam_role.job_template.id
  policy = data.aws_iam_policy_document.s3_access.json
}

# Create IoT Job Template
resource "awscc_iot_job_template" "example" {
  job_template_id = "ExampleJobTemplate"
  description     = "Example IoT job template using AWSCC provider"
  document_source = "s3://${awscc_s3_bucket.job_document.id}/${aws_s3_object.job_document.key}"

  job_executions_rollout_config = {
    maximum_per_minute = 100
    exponential_rollout_rate = {
      base_rate_per_minute = 50
      increment_factor     = 2
      rate_increase_criteria = {
        number_of_notified_things = 1000
      }
    }
  }

  timeout_config = {
    in_progress_timeout_in_minutes = 30
  }

  presigned_url_config = {
    role_arn       = awscc_iam_role.job_template.arn
    expires_in_sec = 3600
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}