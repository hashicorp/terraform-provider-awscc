# S3 bucket for SageMaker processing job
resource "awscc_s3_bucket" "example_bucket" {
  bucket_name = "sagemaker-example-bucket"
}

# IAM role for SageMaker processing job
resource "awscc_iam_role" "example_role" {
  role_name = "example-sagemaker-processing-role"

  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "sagemaker.amazonaws.com"
        }
      },
    ]
  })
}

# IAM managed policy for SageMaker
resource "awscc_iam_managed_policy" "example_policy" {
  managed_policy_name = "example-sagemaker-policy"
  path                = "/"

  policy_document = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect   = "Allow",
        Action   = "sagemaker:*",
        Resource = "*"
      },
      {
        Effect = "Allow",
        Action = [
          "s3:GetObject",
          "s3:PutObject",
          "s3:ListBucket",
          "s3:CreateBucket",
          "s3:GetBucketLocation"
        ],
        Resource = "*"
      }
    ]
  })
}

# IAM role policy for SageMaker access
resource "awscc_iam_role_policy" "example_role_policy" {
  role_name   = awscc_iam_role.example_role.role_name
  policy_name = "SageMakerAccess"

  policy_document = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect   = "Allow",
        Action   = "sagemaker:*",
        Resource = "*"
      },
      {
        Effect = "Allow",
        Action = [
          "s3:GetObject",
          "s3:PutObject",
          "s3:ListBucket",
          "s3:CreateBucket",
          "s3:GetBucketLocation"
        ],
        Resource = "*"
      }
    ]
  })
}

# SageMaker processing job
resource "awscc_sagemaker_processing_job" "example" {
  processing_job_name = "example-processing-job"

  app_specification = {
    image_uri            = "683313688378.dkr.ecr.us-west-2.amazonaws.com/sagemaker-scikit-learn:0.23-1-cpu-py3"
    container_entrypoint = ["python3", "-c", "import pandas as pd; print('Processing complete!')"]
  }

  processing_resources = {
    cluster_config = {
      instance_count    = 1
      instance_type     = "ml.m5.xlarge"
      volume_size_in_gb = 30
    }
  }

  role_arn = awscc_iam_role.example_role.arn

  processing_output_config = {
    outputs = [
      {
        output_name = "processed-data",
        s3_output = {
          s3_uri         = "s3://${awscc_s3_bucket.example_bucket.bucket_name}/output/",
          local_path     = "/opt/ml/processing/output",
          s3_upload_mode = "EndOfJob"
        }
      }
    ]
  }

  environment = {
    "PREPROCESSING_MODE" = "full"
  }

  stopping_condition = {
    max_runtime_in_seconds = 3600
  }

  tags = [
    {
      key   = "Environment"
      value = "Development"
    },
    {
      key   = "Name"
      value = "example-processing-job"
    }
  ]
}
