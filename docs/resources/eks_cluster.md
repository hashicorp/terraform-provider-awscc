---
page_title: "awscc_eks_cluster Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  An object representing an Amazon EKS cluster.
---

# awscc_eks_cluster (Resource)

An object representing an Amazon EKS cluster.

## Example Usage

### Basic usage with IAM Role and Tags
To use awscc_eks_cluster for creating Amazon EKS cluster with a IAM role and tags
```terraform
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
    } ],
    var.eks_default_tags)
}

resource "awscc_eks_cluster" "main" {
  name     = var.eks_cluster_name
  role_arn = awscc_iam_role.main.arn
  version  = var.eks_cluster_version
  resources_vpc_config = {
    subnet_ids = var.eks_cluster_subnets
  }
  tags = concat([ 
    {
       key   = "Name"
       value = var.eks_cluster_name
    } ],
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
```

### Enable Control Plane Logging in Amazon EKS
To use awscc_eks_cluster for creating Amazon EKS Cluster with control plane logging enabled
```terraform
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
```

### Enable Secrets Encryption with KMS in Amazon EKS
To use awscc_eks_cluster for creating Amazon EKS Cluster with secrets encryption enabled using AWS KMS
```terraform
data "aws_partition" "current" {}

data "aws_caller_identity" "current" {}

locals {
  policy_arn_prefix = "arn:${data.aws_partition.current.partition}:iam::aws:policy"
  account_id        = data.aws_caller_identity.current.account_id
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
  encryption_config = [{
    provider = {
      key_arn = awscc_kms_key.main.arn
    }
    resources = ["secrets"]
  }]
  tags = concat([
    {
      key   = "Name"
      value = var.eks_cluster_name
    }],
  var.eks_default_tags)
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
          "AWS" : "arn:aws:iam::${local.account_id}:root"
        },
        "Action" : "kms:*",
        "Resource" : "*"
      },
    ],
    },
  )
  tags = concat([
    {
      key   = "Name"
      value = "${var.eks_cluster_name}-kms-key"
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

output "eks_cluster_encryption_config_key_arn" {
  value = awscc_eks_cluster.main.encryption_config_key_arn
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `resources_vpc_config` (Attributes) An object representing the VPC configuration to use for an Amazon EKS cluster. (see [below for nested schema](#nestedatt--resources_vpc_config))
- `role_arn` (String) The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.

### Optional

- `encryption_config` (Attributes List) (see [below for nested schema](#nestedatt--encryption_config))
- `kubernetes_network_config` (Attributes) The Kubernetes network configuration for the cluster. (see [below for nested schema](#nestedatt--kubernetes_network_config))
- `logging` (Attributes) Enable exporting the Kubernetes control plane logs for your cluster to CloudWatch Logs based on log types. By default, cluster control plane logs aren't exported to CloudWatch Logs. (see [below for nested schema](#nestedatt--logging))
- `name` (String) The unique name to give to your cluster.
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))
- `version` (String) The desired Kubernetes version for your cluster. If you don't specify a value here, the latest version available in Amazon EKS is used.

### Read-Only

- `arn` (String) The ARN of the cluster, such as arn:aws:eks:us-west-2:666666666666:cluster/prod.
- `certificate_authority_data` (String) The certificate-authority-data for your cluster.
- `cluster_security_group_id` (String) The cluster security group that was created by Amazon EKS for the cluster. Managed node groups use this security group for control plane to data plane communication.
- `encryption_config_key_arn` (String) Amazon Resource Name (ARN) or alias of the customer master key (CMK).
- `endpoint` (String) The endpoint for your Kubernetes API server, such as https://5E1D0CEXAMPLEA591B746AFC5AB30262.yl4.us-west-2.eks.amazonaws.com.
- `id` (String) Uniquely identifies the resource.
- `open_id_connect_issuer_url` (String) The issuer URL for the cluster's OIDC identity provider, such as https://oidc.eks.us-west-2.amazonaws.com/id/EXAMPLED539D4633E53DE1B716D3041E. If you need to remove https:// from this output value, you can include the following code in your template.

<a id="nestedatt--resources_vpc_config"></a>
### Nested Schema for `resources_vpc_config`

Required:

- `subnet_ids` (List of String) Specify subnets for your Amazon EKS nodes. Amazon EKS creates cross-account elastic network interfaces in these subnets to allow communication between your nodes and the Kubernetes control plane.

Optional:

- `endpoint_private_access` (Boolean) Set this value to true to enable private access for your cluster's Kubernetes API server endpoint. If you enable private access, Kubernetes API requests from within your cluster's VPC use the private VPC endpoint. The default value for this parameter is false, which disables private access for your Kubernetes API server. If you disable private access and you have nodes or AWS Fargate pods in the cluster, then ensure that publicAccessCidrs includes the necessary CIDR blocks for communication with the nodes or Fargate pods.
- `endpoint_public_access` (Boolean) Set this value to false to disable public access to your cluster's Kubernetes API server endpoint. If you disable public access, your cluster's Kubernetes API server can only receive requests from within the cluster VPC. The default value for this parameter is true, which enables public access for your Kubernetes API server.
- `public_access_cidrs` (List of String) The CIDR blocks that are allowed access to your cluster's public Kubernetes API server endpoint. Communication to the endpoint from addresses outside of the CIDR blocks that you specify is denied. The default value is 0.0.0.0/0. If you've disabled private endpoint access and you have nodes or AWS Fargate pods in the cluster, then ensure that you specify the necessary CIDR blocks.
- `security_group_ids` (List of String) Specify one or more security groups for the cross-account elastic network interfaces that Amazon EKS creates to use to allow communication between your worker nodes and the Kubernetes control plane. If you don't specify a security group, the default security group for your VPC is used.


<a id="nestedatt--encryption_config"></a>
### Nested Schema for `encryption_config`

Optional:

- `provider` (Attributes) The encryption provider for the cluster. (see [below for nested schema](#nestedatt--encryption_config--provider))
- `resources` (List of String) Specifies the resources to be encrypted. The only supported value is "secrets".

<a id="nestedatt--encryption_config--provider"></a>
### Nested Schema for `encryption_config.provider`

Optional:

- `key_arn` (String) Amazon Resource Name (ARN) or alias of the KMS key. The KMS key must be symmetric, created in the same region as the cluster, and if the KMS key was created in a different account, the user must have access to the KMS key.



<a id="nestedatt--kubernetes_network_config"></a>
### Nested Schema for `kubernetes_network_config`

Optional:

- `ip_family` (String) Ipv4 or Ipv6. You can only specify ipv6 for 1.21 and later clusters that use version 1.10.1 or later of the Amazon VPC CNI add-on
- `service_ipv_4_cidr` (String) The CIDR block to assign Kubernetes service IP addresses from. If you don't specify a block, Kubernetes assigns addresses from either the 10.100.0.0/16 or 172.20.0.0/16 CIDR blocks. We recommend that you specify a block that does not overlap with resources in other networks that are peered or connected to your VPC.

Read-Only:

- `service_ipv_6_cidr` (String) The CIDR block to assign Kubernetes service IP addresses from.


<a id="nestedatt--logging"></a>
### Nested Schema for `logging`

Optional:

- `cluster_logging` (Attributes) The cluster control plane logging configuration for your cluster. (see [below for nested schema](#nestedatt--logging--cluster_logging))

<a id="nestedatt--logging--cluster_logging"></a>
### Nested Schema for `logging.cluster_logging`

Optional:

- `enabled_types` (Attributes List) Enable control plane logs for your cluster, all log types will be disabled if the array is empty (see [below for nested schema](#nestedatt--logging--cluster_logging--enabled_types))

<a id="nestedatt--logging--cluster_logging--enabled_types"></a>
### Nested Schema for `logging.cluster_logging.enabled_types`

Optional:

- `type` (String) name of the log type




<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_eks_cluster.example <resource ID>
```