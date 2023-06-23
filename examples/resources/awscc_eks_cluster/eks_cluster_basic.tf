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
}