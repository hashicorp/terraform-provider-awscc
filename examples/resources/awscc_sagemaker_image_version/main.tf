# Data sources to get AWS region
data "aws_region" "current" {}

# Create IAM role for SageMaker using AWSCC
resource "awscc_iam_role" "sagemaker_image_role" {
  role_name = "sagemaker-image-role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "sagemaker.amazonaws.com"
        }
      }
    ]
  })

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create IAM policy for SageMaker (keeping AWS provider as no AWSCC equivalent)
resource "aws_iam_role_policy" "sagemaker_image_policy" {
  name = "sagemaker-image-policy"
  role = awscc_iam_role.sagemaker_image_role.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "ecr:BatchCheckLayerAvailability",
          "ecr:BatchGetImage",
          "ecr:GetDownloadUrlForLayer",
          "ecr:GetAuthorizationToken"
        ]
        Resource = ["arn:aws:ecr:${data.aws_region.current.name}:763104351884:repository/*"]
      },
      {
        Effect = "Allow"
        Action = [
          "ecr:GetAuthorizationToken"
        ]
        Resource = ["*"]
      }
    ]
  })
}

# First create a SageMaker Image
resource "awscc_sagemaker_image" "example" {
  image_name     = "example-sagemaker-image"
  image_role_arn = awscc_iam_role.sagemaker_image_role.arn
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the Image Version
resource "awscc_sagemaker_image_version" "example" {
  image_name       = awscc_sagemaker_image.example.image_name
  base_image       = "763104351884.dkr.ecr.${data.aws_region.current.name}.amazonaws.com/pytorch-training:1.8.1-gpu-py36-cu111-ubuntu18.04"
  alias            = "v1"
  aliases          = ["latest", "stable"]
  horovod          = true
  job_type         = "TRAINING"
  ml_framework     = "PyTorch1.8"
  processor        = "GPU"
  programming_lang = "Python3.6"
  release_notes    = "PyTorch 1.8.1 training image version"
  vendor_guidance  = "STABLE"
}