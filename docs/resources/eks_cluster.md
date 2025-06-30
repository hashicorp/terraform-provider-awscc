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
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_eks_cluster" "main" {
  name     = "example-cluster"
  role_arn = awscc_iam_role.main.arn
  resources_vpc_config = {
    subnet_ids = ["subnet-xxxx", "subnet-yyyy"] // EKS Cluster Subnet-IDs
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

### Enable Control Plane Logging in Amazon EKS
To use awscc_eks_cluster for creating Amazon EKS Cluster with control plane logging enabled
```terraform
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
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_eks_cluster" "main" {
  name     = "example-cluster"
  role_arn = awscc_iam_role.main.arn
  resources_vpc_config = {
    subnet_ids = ["subnet-xxxx", "subnet-yyyy"] // EKS Cluster Subnet-IDs
  }
  logging = {
    cluster_logging = {
      enabled_types = [
        {
          type = "api"
        },
        {
          type = "audit"
        },
        {
          type = "authenticator"
        }
      ]
    }
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
  depends_on = [awscc_logs_log_group.main]
}

resource "awscc_logs_log_group" "main" {
  # The log group name format is /aws/eks/<cluster-name>/cluster
  # Reference: https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html
  log_group_name    = "/aws/eks/example-cluster/cluster"
  retention_in_days = 7
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

### Enable Secrets Encryption with KMS in Amazon EKS
To use awscc_eks_cluster for creating Amazon EKS Cluster with secrets encryption enabled using AWS KMS
```terraform
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
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
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
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
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
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `resources_vpc_config` (Attributes) An object representing the VPC configuration to use for an Amazon EKS cluster. (see [below for nested schema](#nestedatt--resources_vpc_config))
- `role_arn` (String) The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.

### Optional

- `access_config` (Attributes) An object representing the Access Config to use for the cluster. (see [below for nested schema](#nestedatt--access_config))
- `bootstrap_self_managed_addons` (Boolean) Set this value to false to avoid creating the default networking add-ons when the cluster is created.
- `compute_config` (Attributes) Todo: add description (see [below for nested schema](#nestedatt--compute_config))
- `encryption_config` (Attributes List) (see [below for nested schema](#nestedatt--encryption_config))
- `force` (Boolean) Force cluster version update
- `kubernetes_network_config` (Attributes) The Kubernetes network configuration for the cluster. (see [below for nested schema](#nestedatt--kubernetes_network_config))
- `logging` (Attributes) Enable exporting the Kubernetes control plane logs for your cluster to CloudWatch Logs based on log types. By default, cluster control plane logs aren't exported to CloudWatch Logs. (see [below for nested schema](#nestedatt--logging))
- `name` (String) The unique name to give to your cluster.
- `outpost_config` (Attributes) An object representing the Outpost configuration to use for AWS EKS outpost cluster. (see [below for nested schema](#nestedatt--outpost_config))
- `remote_network_config` (Attributes) Configuration fields for specifying on-premises node and pod CIDRs that are external to the VPC passed during cluster creation. (see [below for nested schema](#nestedatt--remote_network_config))
- `storage_config` (Attributes) Todo: add description (see [below for nested schema](#nestedatt--storage_config))
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))
- `upgrade_policy` (Attributes) An object representing the Upgrade Policy to use for the cluster. (see [below for nested schema](#nestedatt--upgrade_policy))
- `version` (String) The desired Kubernetes version for your cluster. If you don't specify a value here, the latest version available in Amazon EKS is used.
- `zonal_shift_config` (Attributes) The current zonal shift configuration to use for the cluster. (see [below for nested schema](#nestedatt--zonal_shift_config))

### Read-Only

- `arn` (String) The ARN of the cluster, such as arn:aws:eks:us-west-2:666666666666:cluster/prod.
- `certificate_authority_data` (String) The certificate-authority-data for your cluster.
- `cluster_id` (String) The unique ID given to your cluster.
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


<a id="nestedatt--access_config"></a>
### Nested Schema for `access_config`

Optional:

- `authentication_mode` (String) Specify the authentication mode that should be used to create your cluster.
- `bootstrap_cluster_creator_admin_permissions` (Boolean) Set this value to false to avoid creating a default cluster admin Access Entry using the IAM principal used to create the cluster.


<a id="nestedatt--compute_config"></a>
### Nested Schema for `compute_config`

Optional:

- `enabled` (Boolean) Todo: add description
- `node_pools` (List of String) Todo: add description
- `node_role_arn` (String) Todo: add description


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

- `elastic_load_balancing` (Attributes) Todo: add description (see [below for nested schema](#nestedatt--kubernetes_network_config--elastic_load_balancing))
- `ip_family` (String) Ipv4 or Ipv6. You can only specify ipv6 for 1.21 and later clusters that use version 1.10.1 or later of the Amazon VPC CNI add-on
- `service_ipv_4_cidr` (String) The CIDR block to assign Kubernetes service IP addresses from. If you don't specify a block, Kubernetes assigns addresses from either the 10.100.0.0/16 or 172.20.0.0/16 CIDR blocks. We recommend that you specify a block that does not overlap with resources in other networks that are peered or connected to your VPC.

Read-Only:

- `service_ipv_6_cidr` (String) The CIDR block to assign Kubernetes service IP addresses from.

<a id="nestedatt--kubernetes_network_config--elastic_load_balancing"></a>
### Nested Schema for `kubernetes_network_config.elastic_load_balancing`

Optional:

- `enabled` (Boolean) Todo: add description



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




<a id="nestedatt--outpost_config"></a>
### Nested Schema for `outpost_config`

Optional:

- `control_plane_instance_type` (String) Specify the Instance type of the machines that should be used to create your cluster.
- `control_plane_placement` (Attributes) Specify the placement group of the control plane machines for your cluster. (see [below for nested schema](#nestedatt--outpost_config--control_plane_placement))
- `outpost_arns` (List of String) Specify one or more Arn(s) of Outpost(s) on which you would like to create your cluster.

<a id="nestedatt--outpost_config--control_plane_placement"></a>
### Nested Schema for `outpost_config.control_plane_placement`

Optional:

- `group_name` (String) Specify the placement group name of the control place machines for your cluster.



<a id="nestedatt--remote_network_config"></a>
### Nested Schema for `remote_network_config`

Optional:

- `remote_node_networks` (Attributes List) Network configuration of nodes run on-premises with EKS Hybrid Nodes. (see [below for nested schema](#nestedatt--remote_network_config--remote_node_networks))
- `remote_pod_networks` (Attributes List) Network configuration of pods run on-premises with EKS Hybrid Nodes. (see [below for nested schema](#nestedatt--remote_network_config--remote_pod_networks))

<a id="nestedatt--remote_network_config--remote_node_networks"></a>
### Nested Schema for `remote_network_config.remote_node_networks`

Optional:

- `cidrs` (List of String) Specifies the list of remote node CIDRs.


<a id="nestedatt--remote_network_config--remote_pod_networks"></a>
### Nested Schema for `remote_network_config.remote_pod_networks`

Optional:

- `cidrs` (List of String) Specifies the list of remote pod CIDRs.



<a id="nestedatt--storage_config"></a>
### Nested Schema for `storage_config`

Optional:

- `block_storage` (Attributes) Todo: add description (see [below for nested schema](#nestedatt--storage_config--block_storage))

<a id="nestedatt--storage_config--block_storage"></a>
### Nested Schema for `storage_config.block_storage`

Optional:

- `enabled` (Boolean) Todo: add description



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.


<a id="nestedatt--upgrade_policy"></a>
### Nested Schema for `upgrade_policy`

Optional:

- `support_type` (String) Specify the support type for your cluster.


<a id="nestedatt--zonal_shift_config"></a>
### Nested Schema for `zonal_shift_config`

Optional:

- `enabled` (Boolean) Set this value to true to enable zonal shift for the cluster.

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_eks_cluster.example "name"
```