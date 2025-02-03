data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["sagemaker.amazonaws.com"]
    }
    effect = "Allow"
  }
}

resource "awscc_iam_role" "model_bias" {
  role_name                   = "AWSCCSageMakerModelBiasRole"
  assume_role_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.assume_role.json))
  managed_policy_arns         = ["arn:aws:iam::aws:policy/service-role/AWSGlueServiceRole"]
}

resource "awscc_sagemaker_model_bias_job_definition" "example" {
  job_definition_name = "example-model-bias-job"

  model_bias_app_specification = {
    image_uri  = "174368400705.dkr.ecr.${data.aws_region.current.name}.amazonaws.com/sagemaker-clarify-processing:1.0"
    config_uri = "s3://sagemaker-${data.aws_region.current.name}-${data.aws_caller_identity.current.account_id}/bias/config.json"
  }

  model_bias_job_input = {
    ground_truth_s3_input = {
      s3_uri = "s3://sagemaker-${data.aws_region.current.name}-${data.aws_caller_identity.current.account_id}/ground-truth/labels.csv"
    }
    endpoint_input = {
      endpoint_name             = "example-endpoint"
      local_path                = "/opt/ml/processing/input"
      s3_data_distribution_type = "FullyReplicated"
      s3_input_mode             = "File"
      features_attribute        = "features"
      inference_attribute       = "predictions"
      probability_attribute     = "probability"
      start_time_offset         = "-PT1H"
      end_time_offset           = "PT0H"
    }
  }

  model_bias_job_output_config = {
    monitoring_outputs = [{
      s3_output = {
        local_path     = "/opt/ml/processing/output"
        s3_uri         = "s3://sagemaker-${data.aws_region.current.name}-${data.aws_caller_identity.current.account_id}/bias-output/"
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

  role_arn = awscc_iam_role.model_bias.arn

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}