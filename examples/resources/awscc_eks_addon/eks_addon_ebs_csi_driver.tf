# Create custom policy for KMS support. It's optional, but recommended.
resource "awscc_iam_managed_policy" "ebs_csi_kms_policy" {
  managed_policy_name = "AmazonEKS_EBS_CSI_KMS_Policy"
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
resource "awscc_iam_role" "ebs_csi_role" {
  role_name = "AmazonEKS_EBS_CSI_DriverRole"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Effect = "Allow"
      Principal = {
        Federated = awscc_iam_oidc_provider.eks.arn
        # Example: "arn:aws:iam::111122223333:oidc-provider/oidc.eks.region-code.amazonaws.com/id/EXAMPLED539D4633E53DE1B71EXAMPLE"
      }
      Action = "sts:AssumeRoleWithWebIdentity"
      Condition = {
        StringEquals = {
          "${local.oidc_provider}:aud" = "sts.amazonaws.com"
          "${local.oidc_provider}:sub" = "system:serviceaccount:kube-system:ebs-csi-controller-sa"
        }
      }
    }]
  })
  managed_policy_arns = [
    "arn:aws:iam::aws:policy/service-role/AmazonEBSCSIDriverPolicy",
    awscc_iam_managed_policy.ebs_csi_kms_policy.policy_arn
  ]
}

# Now that the IAM role is ready, create EBS CSI addon
resource "awscc_eks_addon" "ebs_csi" {
  cluster_name             = awscc_eks_cluster.example.name
  addon_name               = "aws-ebs-csi-driver"
  service_account_role_arn = awscc_iam_role.ebs_csi_role.arn
  resolve_conflicts        = "OVERWRITE"
}
