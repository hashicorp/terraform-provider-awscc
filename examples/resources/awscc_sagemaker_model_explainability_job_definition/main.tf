data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type = "Service"
      identifiers = [
        "sagemaker.amazonaws.com"
      ]
    }
  }
}

data "aws_iam_policy_document" "sagemaker" {
  statement {
    effect = "Allow"
    actions = [
      "s3:GetObject",
      "s3:PutObject",
      "s3:DeleteObject",
      "s3:ListBucket"
    ]
    resources = [
      "arn:aws:s3:::sagemaker-${data.aws_region.current.name}-${data.aws_caller_identity.current.account_id}/*",
      "arn:aws:s3:::sagemaker-${data.aws_region.current.name}-${data.aws_caller_identity.current.account_id}"
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

resource "awscc_iam_role" "explainability" {
  assume_role_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.assume_role.json))
  description                 = "Role for SageMaker Model Explainability Job Definition"
  path                        = "/service-role/"
  role_name                   = "AWSCCSageMakerExplainabilityRole"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_iam_role_policy" "explainability" {
  policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.sagemaker.json))
  policy_name     = "SageMakerExplainabilityAccess"
  role_name       = awscc_iam_role.explainability.role_name
}

resource "awscc_sagemaker_model_explainability_job_definition" "example" {
  job_definition_name = "example-explainability-job"
  role_arn            = awscc_iam_role.explainability.arn

  job_resources = {
    cluster_config = {
      instance_count    = 1
      instance_type     = "ml.m5.large"
      volume_size_in_gb = 20
    }
  }

  model_explainability_app_specification = {
    image_uri  = "123456789012.dkr.ecr.us-west-2.amazonaws.com/sagemaker-clarify-processing:1.0"
    config_uri = "s3://sagemaker-${data.aws_region.current.name}-${data.aws_caller_identity.current.account_id}/explainability/config.json"
  }

  model_explainability_job_input = {
    endpoint_input = {
      endpoint_name             = "example-endpoint"
      local_path                = "/opt/ml/processing/input"
      s3_data_distribution_type = "FullyReplicated"
      s3_input_mode             = "File"
      features_attribute        = "feature"
      inference_attribute       = "inference"
      probability_attribute     = "probability"
    }
  }

  model_explainability_job_output_config = {
    monitoring_outputs = [{
      s3_output = {
        local_path     = "/opt/ml/processing/output"
        s3_uri         = "s3://sagemaker-${data.aws_region.current.name}-${data.aws_caller_identity.current.account_id}/explainability/output"
        s3_upload_mode = "EndOfJob"
      }
    }]
  }

  network_config = {
    enable_network_isolation                  = true
    enable_inter_container_traffic_encryption = true
  }

  stopping_condition = {
    max_runtime_in_seconds = 3600
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}