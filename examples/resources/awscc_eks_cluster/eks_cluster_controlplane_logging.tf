data "aws_partition" "current" {}

locals {
  policy_arn_prefix = "arn:${data.aws_partition.current.partition}:iam::aws:policy"
}

variable "eks_default_tags" {
  description = "Default tags to be applied to EKS resources"
  default = [
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

variable "eks_cluster_name" {
  description = "EKS Cluster Name"
  type        = string
  default     = "example-cluster"
}

variable "eks_cluster_version" {
  description = "EKS Cluster Version"
  type        = string
  default     = "1.27"
}

variable "eks_cluster_subnets" {
  description = "Subnets for EKS Cluster"
  type        = list(string)
  default     = ["subnet-xxxx", "subnet-yyyy"] // Provide a list of subnet-ids for Amazon EKS Cluster
}

variable "enabled_cluster_log_types" {
  default     = ["api", "audit", "authenticator", "controllerManager", "scheduler"]
  description = "A list of the Logs to be enabled in control plane."
  type        = list(string)
}

resource "awscc_iam_role" "main" {
  description = "IAM Role of EKS Cluster"
  role_name   = "${var.eks_cluster_name}-role"
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
    "${local.policy_arn_prefix}/AmazonEKSClusterPolicy",
    # Optionally, enable Security Groups for Pods
    # Reference: https://docs.aws.amazon.com/eks/latest/userguide/security-groups-for-pods.html
    "${local.policy_arn_prefix}/AmazonEKSVPCResourceController"
  ]
  max_session_duration = 7200
  path                 = "/"
  tags = concat([
    {
      key   = "Name"
      value = "${var.eks_cluster_name}-role"
    }],
  var.eks_default_tags)
}

resource "awscc_eks_cluster" "main" {
  name     = var.eks_cluster_name
  role_arn = awscc_iam_role.main.arn
  version  = var.eks_cluster_version
  resources_vpc_config = {
    subnet_ids = var.eks_cluster_subnets
  }
  logging = {
    cluster_logging = {
      enabled_types = [
        for log_type in var.enabled_cluster_log_types : {
          type = log_type
        }
      ]
    }
  }
  tags = concat([
    {
      key   = "Name"
      value = var.eks_cluster_name
    }],
  var.eks_default_tags)
  depends_on = [awscc_logs_log_group.main]
}

resource "awscc_logs_log_group" "main" {
  # The log group name format is /aws/eks/<cluster-name>/cluster
  # Reference: https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html
  log_group_name    = "/aws/eks/${var.eks_cluster_name}/cluster"
  retention_in_days = 7
  tags = concat([
    {
      key   = "Name"
      value = "/aws/eks/${var.eks_cluster_name}/cluster"
    }],
  var.eks_default_tags)
}


output "eks_cluster_endpoint" {
  value = awscc_eks_cluster.main.endpoint
}

output "eks_cluster_arn" {
  value = awscc_eks_cluster.main.arn
}

output "eks_cluster_certificate_authority_data" {
  value = awscc_eks_cluster.main.certificate_authority_data
}

# Cluster Security Group ID created by Amazon EKS for cluster
output "eks_cluster_security_group_id" {
  value = awscc_eks_cluster.main.cluster_security_group_id
}

output "eks_cluster_log_group_arn" {
  value = awscc_logs_log_group.main.arn
}