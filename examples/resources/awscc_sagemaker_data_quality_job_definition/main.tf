data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create S3 bucket for outputs
resource "awscc_s3_bucket" "monitoring" {
  bucket_name = "sagemaker-dataquality-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# IAM role for SageMaker
data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["sagemaker.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "sagemaker_policy" {
  statement {
    effect = "Allow"
    actions = [
      "s3:GetObject",
      "s3:PutObject",
      "s3:ListBucket"
    ]
    resources = [
      "arn:aws:s3:::${awscc_s3_bucket.monitoring.bucket_name}",
      "arn:aws:s3:::${awscc_s3_bucket.monitoring.bucket_name}/*"
    ]
  }
  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents"
    ]
    resources = ["arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:log-group:/aws/sagemaker/*"]
  }
}

resource "awscc_iam_role" "sagemaker_role" {
  assume_role_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.assume_role.json))
  description                 = "IAM role for SageMaker Data Quality Job Definition"
  policies = [{
    policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.sagemaker_policy.json))
    policy_name     = "SageMakerDataQualityPolicy"
  }]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Data Quality Job Definition
resource "awscc_sagemaker_data_quality_job_definition" "example" {
  job_definition_name = "example-data-quality-job"

  data_quality_app_specification = {
    image_uri = "174368400705.dkr.ecr.${data.aws_region.current.name}.amazonaws.com/sagemaker-model-monitor-analyzer:latest"
  }

  data_quality_job_input = {
    batch_transform_input = {
      data_captured_destination_s3_uri = "s3://${awscc_s3_bucket.monitoring.bucket_name}/input"
      local_path                       = "/opt/ml/processing/input"
      s3_data_distribution_type        = "FullyReplicated"
      s3_input_mode                    = "File"
      dataset_format = {
        csv = {
          header = true
        }
      }
    }
  }

  data_quality_job_output_config = {
    monitoring_outputs = [{
      s3_output = {
        local_path     = "/opt/ml/processing/output"
        s3_uri         = "s3://${awscc_s3_bucket.monitoring.bucket_name}/output"
        s3_upload_mode = "EndOfJob"
      }
    }]
  }

  job_resources = {
    cluster_config = {
      instance_count    = 1
      instance_type     = "ml.m5.large"
      volume_size_in_gb = 20
    }
  }

  role_arn = awscc_iam_role.sagemaker_role.arn

  stopping_condition = {
    max_runtime_in_seconds = 3600
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}