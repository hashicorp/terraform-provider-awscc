---
page_title: "awscc_eks_addon Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Schema for AWS::EKS::Addon
---

# awscc_eks_addon (Resource)

Resource Schema for AWS::EKS::Addon

## Example Usage

### Basic usage to create coredns and kube_proxy addons
```terraform
resource "awscc_eks_addon" "coredns" {
  cluster_name = var.cluster_name
  addon_name   = "coredns"
  # Optional: addon_version = "v1.8.4-eksbuild.1"
  # Optional: resolve_conflicts = "OVERWRITE"
}

resource "awscc_eks_addon" "kube_proxy" {
  cluster_name = var.cluster_name
  addon_name   = "kube-proxy"
}

variable "cluster_name" {
  type = string
}
```

### Create EBS CSI addon
```terraform
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
```

### Create EFS CSI addon
```terraform
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
```

### Create VPC CNI addon
```terraform
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `addon_name` (String) Name of Addon
- `cluster_name` (String) Name of Cluster

### Optional

- `addon_version` (String) Version of Addon
- `configuration_values` (String) The configuration values to use with the add-on
- `pod_identity_associations` (Attributes Set) An array of pod identities to apply to this add-on. (see [below for nested schema](#nestedatt--pod_identity_associations))
- `preserve_on_delete` (Boolean) PreserveOnDelete parameter value
- `resolve_conflicts` (String) Resolve parameter value conflicts
- `service_account_role_arn` (String) IAM role to bind to the add-on's service account
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String) Amazon Resource Name (ARN) of the add-on
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--pod_identity_associations"></a>
### Nested Schema for `pod_identity_associations`

Optional:

- `role_arn` (String) The IAM role ARN that the pod identity association is created for.
- `service_account` (String) The Kubernetes service account that the pod identity association is created for.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_eks_addon.example "cluster_name|addon_name"
```
