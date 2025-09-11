resource "awscc_sagemaker_cluster" "hyperpod_autoscaling" {
  cluster_name = "hyperpod-eks-autoscaling"
  
  instance_groups = [
    {
      execution_role      = awscc_iam_role.execution.arn
      instance_count      = 1
      instance_type       = "ml.c5.xlarge"
      instance_group_name = "system"
      life_cycle_config = {
        source_s3_uri = "s3://${aws_s3_bucket.lifecycle.id}/config/"
        on_create     = "on_create.sh"
      }
    },
    {
      execution_role      = awscc_iam_role.execution.arn
      instance_count      = 0
      instance_type       = "ml.c5.xlarge"
      instance_group_name = "auto-c5-az1"
      life_cycle_config = {
        source_s3_uri = "s3://${aws_s3_bucket.lifecycle.id}/config/"
        on_create     = "on_create.sh"
      }
    },
    {
      execution_role      = awscc_iam_role.execution.arn
      instance_count      = 0
      instance_type       = "ml.c5.4xlarge"
      instance_group_name = "auto-c5-4xaz2"
      life_cycle_config = {
        source_s3_uri = "s3://${aws_s3_bucket.lifecycle.id}/config/"
        on_create     = "on_create.sh"
      }
      override_vpc_config = {
        security_group_ids = [var.security_group_id]
        subnets           = [var.subnet_2]
      }
    }
  ]
  
  orchestrator = {
    eks = {
      cluster_arn = var.eks_cluster_arn
    }
  }
  
  vpc_config = {
    security_group_ids = [var.security_group_id]
    subnets           = [var.subnet_1]
  }
  
  cluster_role = awscc_iam_role.cluster.arn
  
  auto_scaling = {
    mode            = "Enable"
    auto_scaler_type = "Karpenter"
  }
  
  node_provisioning_mode = "Continuous"
  
  tags = [{
    key   = "Environment"
    value = "Development"
  }, {
    key   = "AutoScaling"
    value = "Enabled"
  }]
}

# IAM Role for Karpenter Autoscaling
resource "awscc_iam_role" "cluster" {
  role_name = "SageMakerHyperPodKarpenterRole"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Principal = {
          Service = ["hyperpod.sagemaker.amazonaws.com"]
        }
        Action = "sts:AssumeRole"
      }
    ]
  })

  policies = [
    {
      policy_name = "SageMakerHyperPodKarpenterPolicy"
      policy_document = jsonencode({
        Version = "2012-10-17"
        Statement = [
          {
            Effect = "Allow"
            Action = [
              "sagemaker:BatchAddClusterNodes",
              "sagemaker:BatchDeleteClusterNodes"
            ]
            Resource = "arn:aws:sagemaker:*:*:cluster/*"
            Condition = {
              StringEquals = {
                "aws:ResourceAccount" = "$${aws:PrincipalAccount}"
              }
            }
          },
          {
            Effect = "Allow"
            Action = [
              "kms:CreateGrant",
              "kms:DescribeKey"
            ]
            Resource = "arn:aws:kms:*:*:key/*"
            Condition = {
              StringLike = {
                "kms:ViaService" = "sagemaker.*.amazonaws.com"
              }
              Bool = {
                "kms:GrantIsForAWSResource" = "true"
              }
              "ForAllValues:StringEquals" = {
                "kms:GrantOperations" = [
                  "CreateGrant",
                  "Decrypt",
                  "DescribeKey",
                  "GenerateDataKeyWithoutPlaintext",
                  "ReEncryptTo",
                  "ReEncryptFrom",
                  "RetireGrant"
                ]
              }
            }
          }
        ]
      })
    }
  ]
}

resource "awscc_iam_role" "execution" {
  role_name = "SageMakerHyperPodExecutionRole"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Principal = {
          Service = ["sagemaker.amazonaws.com"]
        }
        Action = "sts:AssumeRole"
      }
    ]
  })
  
  managed_policy_arns = [
    "arn:aws:iam::aws:policy/AmazonSageMakerFullAccess"
  ]
}

resource "aws_s3_bucket" "lifecycle" {
  bucket = "sagemaker-hyperpod-lifecycle-${random_id.bucket_suffix.hex}"
}

resource "aws_s3_object" "script" {
  bucket = aws_s3_bucket.lifecycle.id
  key    = "config/on_create.sh"
  content = "#!/bin/bash\necho 'HyperPod node initialization complete'"
}

resource "random_id" "bucket_suffix" {
  byte_length = 4
}

variable "eks_cluster_arn" {
  description = "ARN of the EKS cluster"
  type        = string
}

variable "security_group_id" {
  description = "Security group ID for the cluster"
  type        = string
}

variable "subnet_1" {
  description = "First subnet ID"
  type        = string
}

variable "subnet_2" {
  description = "Second subnet ID"
  type        = string
}
