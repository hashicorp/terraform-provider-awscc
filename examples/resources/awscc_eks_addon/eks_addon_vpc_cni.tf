# AWS IAM expects the OIDC provider URL without the `https://` prefix in the condition block. 
# This creates a local variable for it:
locals {
  oidc_provider = replace(awscc_eks_cluster.eks_cluster.open_id_connect_issuer_url, "https://", "")
}

# Create an IAM policy for EKS VPC CNI IPv6 support
# https://docs.aws.amazon.com/eks/latest/userguide/cni-iam-role.html
resource "awscc_iam_managed_policy" "eks_vpc_cni_ipv6_policy" {
  managed_policy_name = "AmazonEKS_CNI_IPv6_Policy"
  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "ec2:AssignIpv6Addresses",
          "ec2:DescribeInstances",
          "ec2:DescribeTags",
          "ec2:DescribeNetworkInterfaces",
          "ec2:DescribeInstanceTypes"
        ]
        Resource = "*"
      },
      {
        Effect   = "Allow"
        Action   = ["ec2:CreateTags"]
        Resource = "arn:aws:ec2:*:*:network-interface/*"
      }
    ]
  })
}

resource "awscc_iam_role" "eks_vpc_cni_role" {
  role_name = "AmazonEKSVPCCNIRole"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect    = "Allow"
        Principal = { Federated = awscc_iam_oidc_provider.eks.arn }
        Action    = "sts:AssumeRoleWithWebIdentity"
        Condition = {
          StringEquals = {
            "${local.oidc_provider}:aud" = "sts.amazonaws.com"
            "${local.oidc_provider}:sub" = "system:serviceaccount:kube-system:aws-node"
          }
        }
      }
    ]
  })
  managed_policy_arns = [
    "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy",
    awscc_iam_managed_policy.eks_vpc_cni_ipv6_policy.policy_arn
  ]
}

# Now that the IAM role is ready, create the VPC CNI plugin:
resource "awscc_eks_addon" "vpc_cni" {
  cluster_name             = var.cluster_name
  addon_name               = "vpc-cni"
  service_account_role_arn = awscc_iam_role.eks_vpc_cni_role.arn
  resolve_conflicts        = "OVERWRITE"
}

variable "cluster_name" {
  type = string
}
