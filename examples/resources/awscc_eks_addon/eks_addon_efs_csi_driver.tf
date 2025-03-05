# AWS IAM expects the OIDC provider URL without the `https://` prefix in the condition block.This creates a local variable for it:

locals {
  oidc_provider = replace(awscc_eks_cluster.eks_cluster.open_id_connect_issuer_url, "https://", "")
}

# Optional Custom policy for KMS support and EBS CSI Driver Role

resource "awscc_iam_managed_policy" "efs_csi_kms_policy" {
  managed_policy_name = "AmazonEKS_EFS_CSI_KMS_Policy"
  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "kms:CreateGrant",
          "kms:ListGrants",
          "kms:RevokeGrant"
        ]
        Resource = awscc_kms_key.example.arn
        Condition = {
          Bool = {
            "kms:GrantIsForAWSResource" = "true"
          }
        }
      },
      {
        Effect = "Allow"
        Action = [
          "kms:Encrypt",
          "kms:Decrypt",
          "kms:ReEncrypt*",
          "kms:GenerateDataKey*",
          "kms:DescribeKey"
        ]
        Resource = awscc_kms_key.example.arn
      }
    ]
  })
}

# Create IAM role for EBS CSI Driver
resource "awscc_iam_role" "efs_csi_role" {
  role_name = "AmazonEKS_EFS_CSI_Driver_Role"
  assume_role_policy_document = jsonencode({
    Statement = [{
      Action = "sts:AssumeRoleWithWebIdentity"
      Effect = "Allow"
      Principal = {
        Federated = awscc_iam_oidc_provider.eks.arn
        # Example: "arn:aws:iam::111122223333:oidc-provider/oidc.eks.region-code.amazonaws.com/id/EXAMPLED539D4633E53DE1B71EXAMPLE"
      }
      Condition = {
        StringEquals = {
          "${local.oidc_provider}:sub" = "system:serviceaccount:kube-system:efs-csi-controller-sa"
          "${local.oidc_provider}:aud" = "sts.amazonaws.com"
        }
      }
    }]
    Version = "2012-10-17"
  })

  managed_policy_arns = [
    awscc_iam_managed_policy.efs_csi_kms_policy.arn,
    "arn:aws:iam::aws:policy/service-role/AmazonEFSCSIDriverPolicy"
  ]
}

# Once the IAM role is ready, create EFS CSI addon

resource "awscc_eks_addon" "efs_csi" {
  cluster_name             = awscc_eks_cluster.cluster_name
  addon_name               = "aws-efs-csi-driver"
  addon_version            = "v2.1.4-eksbuild.1" #Change version to required
  service_account_role_arn = awscc_iam_role.efs_csi_role.arn
  resolve_conflicts        = "OVERWRITE"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
