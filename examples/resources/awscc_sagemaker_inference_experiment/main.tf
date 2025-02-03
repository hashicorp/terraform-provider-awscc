# Data sources to get current AWS account ID and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# IAM Role for SageMaker Inference Experiment
data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    effect  = "Allow"
    principals {
      type        = "Service"
      identifiers = ["sagemaker.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "sagemaker_inference_experiment" {
  statement {
    effect = "Allow"
    actions = [
      "sagemaker:*",
      "s3:GetObject",
      "s3:PutObject",
      "s3:ListBucket",
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents"
    ]
    resources = ["*"]
  }
}

resource "awscc_iam_role" "sagemaker_role" {
  assume_role_policy_document = data.aws_iam_policy_document.assume_role.json
  description                 = "Role for SageMaker Inference Experiment"
  managed_policy_arns         = ["arn:aws:iam::aws:policy/AmazonSageMakerFullAccess"]
  max_session_duration        = 3600
  path                        = "/"
  policies = [{
    policy_document = data.aws_iam_policy_document.sagemaker_inference_experiment.json
    policy_name     = "SageMakerInferenceExperimentPolicy"
  }]
  role_name = "SageMakerInferenceExperimentRole"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# S3 Bucket for inference data storage
resource "awscc_s3_bucket" "inference_data" {
  bucket_name = "inference-data-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# SageMaker Inference Experiment
resource "awscc_sagemaker_inference_experiment" "example" {
  name          = "sample-inference-experiment"
  endpoint_name = "existing-endpoint" # Replace with your existing endpoint name
  type          = "ShadowMode"
  role_arn      = awscc_iam_role.sagemaker_role.arn

  model_variants = [
    {
      model_name   = "existing-model" # Replace with your existing model name
      variant_name = "variant-1"
      infrastructure_config = {
        infrastructure_type = "RealTimeInference"
        real_time_inference_config = {
          instance_count = 1
          instance_type  = "ml.t2.medium"
        }
      }
    }
  ]

  shadow_mode_config = {
    source_model_variant_name = "variant-1"
    shadow_model_variants = [
      {
        shadow_model_variant_name = "shadow-variant-1"
        sampling_percentage       = 10
      }
    ]
  }

  data_storage_config = {
    destination = "s3://${awscc_s3_bucket.inference_data.bucket_name}/inference-data/"
    content_type = {
      json_content_types = ["application/json"]
      csv_content_types  = ["text/csv"]
    }
  }

  description = "Sample inference experiment for testing"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}