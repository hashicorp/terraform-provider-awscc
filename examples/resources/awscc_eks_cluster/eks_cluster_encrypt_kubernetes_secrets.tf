data "aws_caller_identity" "current" {}

resource "awscc_iam_role" "main" {
  description = "IAM Role of EKS Cluster"
  role_name   = "example-role"
  assume_role_policy_document = jsonencode({
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Principal = {
        Service = "eks.amazonaws.com"
      }
    }]
    Version = "2012-10-17"
  })
  managed_policy_arns = [
    "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy",
    # Optionally, enable Security Groups for Pods
    # Reference: https://docs.aws.amazon.com/eks/latest/userguide/security-groups-for-pods.html
    "arn:aws:iam::aws:policy/AmazonEKSVPCResourceController"
  ]
  max_session_duration = 7200
  path                 = "/"
  tags = [
    {
      key   = "Name"
      value = "example-role"
    },
    {
      key   = "Environment"
      value = "Development"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}

resource "awscc_eks_cluster" "main" {
  name     = "example-cluster"
  role_arn = awscc_iam_role.main.arn
  resources_vpc_config = {
    subnet_ids = ["subnet-xxxx", "subnet-yyyy"] // EKS Cluster Subnet-IDs
  }
  encryption_config = [{
    provider = {
      key_arn = awscc_kms_key.main.arn
    }
    resources = ["secrets"]
  }]
  tags = [
    {
      key   = "Name"
      value = "example-cluster"
    },
    {
      key   = "Environment"
      value = "Development"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
  depends_on = [awscc_kms_key.main]
}

resource "awscc_kms_key" "main" {
  description            = "KMS Key for EKS Secrets Encryption"
  enabled                = "true"
  enable_key_rotation    = "false"
  pending_window_in_days = 30
  key_policy = jsonencode({
    "Version" : "2012-10-17",
    "Id" : "KMS-Key-Policy-For-Root",
    "Statement" : [
      {
        "Sid" : "Enable IAM User Permissions",
        "Effect" : "Allow",
        "Principal" : {
          "AWS" : "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
        },
        "Action" : "kms:*",
        "Resource" : "*"
      },
    ],
    },
  )
  tags = [
    {
      key   = "Name"
      value = "example-kms-key"
    },
    {
      key   = "Environment"
      value = "Development"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}