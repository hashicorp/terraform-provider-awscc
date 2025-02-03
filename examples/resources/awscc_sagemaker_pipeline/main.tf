# Get current AWS account ID and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# IAM role for SageMaker Pipeline
data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["sagemaker.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "pipeline_policy" {
  statement {
    effect = "Allow"
    actions = [
      "sagemaker:CreateTrainingJob",
      "sagemaker:CreateProcessingJob",
      "sagemaker:CreateModelPackage",
      "sagemaker:StopProcessingJob",
      "sagemaker:StopTrainingJob"
    ]
    resources = [
      "arn:aws:sagemaker:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:*"
    ]
  }

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

resource "awscc_iam_role" "sagemaker_pipeline_role" {
  role_name                   = "sagemaker-pipeline-role"
  assume_role_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.assume_role.json))
  managed_policy_arns         = ["arn:aws:iam::aws:policy/AmazonSageMakerFullAccess"]
  policies = [{
    policy_name     = "sagemaker-pipeline-policy"
    policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.pipeline_policy.json))
  }]
}

# SageMaker Pipeline
resource "awscc_sagemaker_pipeline" "example" {
  pipeline_name = "example-pipeline"
  role_arn      = awscc_iam_role.sagemaker_pipeline_role.arn

  pipeline_definition = {
    pipeline_definition_body = jsonencode({
      Version = "2020-12-01"
      Parameters = [
        {
          Name = "InputDataUrl"
          Type = "String"
        }
      ]
      Steps = [
        {
          Name = "Processing"
          Type = "Processing"
          Arguments = {
            ProcessingResources = {
              ClusterConfig = {
                InstanceCount  = 1
                InstanceType   = "ml.m5.xlarge"
                VolumeSizeInGB = 30
              }
            }
            AppSpecification = {
              ImageUri = "${data.aws_caller_identity.current.account_id}.dkr.ecr.${data.aws_region.current.name}.amazonaws.com/sagemaker-scikit-learn:0.23-1-cpu-py3"
              ContainerArguments = [
                "--input-data", "InputDataUrl"
              ]
            }
            ProcessingInputs = []
            ProcessingOutputConfig = {
              Outputs = []
            }
            RoleArn = awscc_iam_role.sagemaker_pipeline_role.arn
          }
        }
      ]
    })
  }

  pipeline_description = "Example SageMaker Pipeline using Terraform AWSCC provider"

  parallelism_configuration = {
    max_parallel_execution_steps = 10
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}