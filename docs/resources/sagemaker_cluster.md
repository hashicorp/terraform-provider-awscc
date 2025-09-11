---
page_title: "awscc_sagemaker_cluster Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::SageMaker::Cluster
---

# awscc_sagemaker_cluster (Resource)

Resource Type definition for AWS::SageMaker::Cluster

## Example Usage

### Basic usage
To create a SageMaker HyperPod Cluster resource. You can find some of the lifecycle scripts at https://github.com/aws-samples/awsome-distributed-training/tree/main/1.architectures/5.sagemaker-hyperpod/LifecycleScripts/base-config.

```terraform
resource "awscc_sagemaker_cluster" "example" {
  cluster_name = "example"
  instance_groups = [
    {
      execution_role      = awscc_iam_role.example.arn
      instance_count      = 1
      instance_type       = "ml.c5.2xlarge"
      instance_group_name = "example"
      life_cycle_config = {
        source_s3_uri = "s3://${aws_s3_bucket.example.id}/config/"
        on_create     = "on_create_noop.sh"
      }
      instance_storage_configs = [{
        ebs_volume_config = {
          volume_size_in_gb = 30
        }
      }]
    }
  ]

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]

}

resource "aws_s3_bucket" "example" {
  bucket = "example"
}

resource "aws_s3_object" "script" {
  bucket = aws_s3_bucket.example.id
  key    = "config/on_create_noop.sh"
  source = "on_create_noop.sh"
}

resource "aws_s3_object" "params" {
  bucket = aws_s3_bucket.example.id
  key    = "config/provisioning_parameters.json"
  source = "provisioning_parameters.json"
}
```

### EKS orchestrator
To create a SageMaker HyperPod Cluster resource with an existing EKS cluster as the orchestrator.

```terraform
resource "awscc_sagemaker_cluster" "this" {
  cluster_name = "example"
  instance_groups = [
    {
      execution_role      = awscc_iam_role.example.arn
      instance_count      = 1
      instance_type       = "ml.c5.2xlarge"
      instance_group_name = "example"
      life_cycle_config = {
        source_s3_uri = "s3://${aws_s3_bucket.this.id}/config/"
        on_create     = "on_create_noop.sh"
      }
    }
  ]
  orchestrator = {
    eks = {
      cluster_arn = "arn:${data.aws_partition.current.partition}:eks:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:cluster/hyperpod-eks-example"
    }
  }
  vpc_config = {
    security_group_ids = [var.sg_id]
    subnets            = [var.subnet_id]
  }

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]

}

resource "aws_s3_bucket" "example" {
  bucket = "example"
}

resource "aws_s3_object" "script" {
  bucket = aws_s3_bucket.example.id
  key    = "config/on_create_noop.sh"
  source = "on_create_noop.sh"
}

resource "aws_s3_object" "params" {
  bucket = aws_s3_bucket.example.id
  key    = "config/provisioning_parameters.json"
  source = "provisioning_parameters.json"
}

data "aws_caller_identity" "current" {}
data "aws_region" "current" {}
data "aws_partition" "current" {}
```

### EKS with Karpenter Autoscaling
To create a SageMaker HyperPod Cluster with EKS orchestration and Karpenter-based autoscaling enabled.

```terraform
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
    }
  ]
  
  orchestrator = {
    eks = {
      cluster_arn = var.eks_cluster_arn
    }
  }
  
  vpc_config = {
    security_group_ids = [var.security_group_id]
    subnets           = [var.subnet_id]
  }
  
  cluster_role = awscc_iam_role.cluster.arn
  
  auto_scaling = {
    mode            = "Enable"
    auto_scaler_type = "Karpenter"
  }
  
  node_provisioning_mode = "Continuous"
  
  tags = [{
    key   = "AutoScaling"
    value = "Enabled"
  }]
}

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
          }
        ]
      })
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `auto_scaling` (Attributes) Configuration for autoscaling the SageMaker HyperPod cluster. (see [below for nested schema](#nestedatt--auto_scaling))
- `cluster_name` (String) The name of the HyperPod Cluster.
- `cluster_role` (String) The IAM role ARN for the cluster.
- `instance_groups` (Attributes List) The instance groups of the SageMaker HyperPod cluster. (see [below for nested schema](#nestedatt--instance_groups))
- `node_provisioning_mode` (String) Determines the scaling strategy for the SageMaker HyperPod cluster. When set to 'Continuous', enables continuous scaling which dynamically manages node provisioning. If the parameter is omitted, uses the standard scaling approach in previous release.
- `node_recovery` (String) If node auto-recovery is set to true, faulty nodes will be replaced or rebooted when a failure is detected. If set to false, nodes will be labelled when a fault is detected.
- `orchestrator` (Attributes) Specifies parameter(s) specific to the orchestrator, e.g. specify the EKS cluster. (see [below for nested schema](#nestedatt--orchestrator))
- `restricted_instance_groups` (Attributes List) The restricted instance groups of the SageMaker HyperPod cluster. (see [below for nested schema](#nestedatt--restricted_instance_groups))
- `tags` (Attributes Set) Custom tags for managing the SageMaker HyperPod cluster as an AWS resource. You can add tags to your cluster in the same way you add them in other AWS services that support tagging. (see [below for nested schema](#nestedatt--tags))
- `vpc_config` (Attributes) Specifies an Amazon Virtual Private Cloud (VPC) that your SageMaker jobs, hosted models, and compute resources have access to. You can control access to and from your resources by configuring a VPC. (see [below for nested schema](#nestedatt--vpc_config))

### Read-Only

- `cluster_arn` (String) The Amazon Resource Name (ARN) of the HyperPod Cluster.
- `cluster_status` (String) The status of the HyperPod Cluster.
- `creation_time` (String) The time at which the HyperPod cluster was created.
- `failure_message` (String) The failure message of the HyperPod Cluster.
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--auto_scaling"></a>
### Nested Schema for `auto_scaling`

Optional:

- `auto_scaler_type` (String) The type of autoscaler to use. Valid values: `Karpenter`.
- `mode` (String) The autoscaling mode for the cluster. Valid values: `Enable`, `Disable`.

Read-Only:

- `status` (String) The status of the autoscaling configuration.


<a id="nestedatt--instance_groups"></a>
### Nested Schema for `instance_groups`

Optional:

- `current_count` (Number) The number of instances that are currently in the instance group of a SageMaker HyperPod cluster.
- `execution_role` (String) The execution role for the instance group to assume.
- `image_id` (String) AMI Id to be used for launching EC2 instances - HyperPodPublicAmiId or CustomAmiId
- `instance_count` (Number) The number of instances you specified to add to the instance group of a SageMaker HyperPod cluster.
- `instance_group_name` (String) The name of the instance group of a SageMaker HyperPod cluster.
- `instance_storage_configs` (Attributes List) The instance storage configuration for the instance group. (see [below for nested schema](#nestedatt--instance_groups--instance_storage_configs))
- `instance_type` (String) The instance type of the instance group of a SageMaker HyperPod cluster.
- `life_cycle_config` (Attributes) The lifecycle configuration for a SageMaker HyperPod cluster. (see [below for nested schema](#nestedatt--instance_groups--life_cycle_config))
- `on_start_deep_health_checks` (List of String) Nodes will undergo advanced stress test to detect and replace faulty instances, based on the type of deep health check(s) passed in.
- `override_vpc_config` (Attributes) Specifies an Amazon Virtual Private Cloud (VPC) that your SageMaker jobs, hosted models, and compute resources have access to. You can control access to and from your resources by configuring a VPC. (see [below for nested schema](#nestedatt--instance_groups--override_vpc_config))
- `scheduled_update_config` (Attributes) The configuration object of the schedule that SageMaker follows when updating the AMI. (see [below for nested schema](#nestedatt--instance_groups--scheduled_update_config))
- `threads_per_core` (Number) The number you specified to TreadsPerCore in CreateCluster for enabling or disabling multithreading. For instance types that support multithreading, you can specify 1 for disabling multithreading and 2 for enabling multithreading.
- `training_plan_arn` (String) The Amazon Resource Name (ARN) of the training plan to use for this cluster instance group. For more information about how to reserve GPU capacity for your SageMaker HyperPod clusters using Amazon SageMaker Training Plan, see CreateTrainingPlan.

<a id="nestedatt--instance_groups--instance_storage_configs"></a>
### Nested Schema for `instance_groups.instance_storage_configs`

Optional:

- `ebs_volume_config` (Attributes) Defines the configuration for attaching additional Amazon Elastic Block Store (EBS) volumes to the instances in the SageMaker HyperPod cluster instance group. The additional EBS volume is attached to each instance within the SageMaker HyperPod cluster instance group and mounted to /opt/sagemaker. (see [below for nested schema](#nestedatt--instance_groups--instance_storage_configs--ebs_volume_config))

<a id="nestedatt--instance_groups--instance_storage_configs--ebs_volume_config"></a>
### Nested Schema for `instance_groups.instance_storage_configs.ebs_volume_config`

Optional:

- `volume_size_in_gb` (Number) The size in gigabytes (GB) of the additional EBS volume to be attached to the instances in the SageMaker HyperPod cluster instance group. The additional EBS volume is attached to each instance within the SageMaker HyperPod cluster instance group and mounted to /opt/sagemaker.



<a id="nestedatt--instance_groups--life_cycle_config"></a>
### Nested Schema for `instance_groups.life_cycle_config`

Optional:

- `on_create` (String) The file name of the entrypoint script of lifecycle scripts under SourceS3Uri. This entrypoint script runs during cluster creation.
- `source_s3_uri` (String) An Amazon S3 bucket path where your lifecycle scripts are stored.


<a id="nestedatt--instance_groups--override_vpc_config"></a>
### Nested Schema for `instance_groups.override_vpc_config`

Optional:

- `security_group_ids` (List of String) The VPC security group IDs, in the form sg-xxxxxxxx. Specify the security groups for the VPC that is specified in the Subnets field.
- `subnets` (List of String) The ID of the subnets in the VPC to which you want to connect your training job or model.


<a id="nestedatt--instance_groups--scheduled_update_config"></a>
### Nested Schema for `instance_groups.scheduled_update_config`

Optional:

- `deployment_config` (Attributes) The configuration to use when updating the AMI versions. (see [below for nested schema](#nestedatt--instance_groups--scheduled_update_config--deployment_config))
- `schedule_expression` (String) A cron expression that specifies the schedule that SageMaker follows when updating the AMI.

<a id="nestedatt--instance_groups--scheduled_update_config--deployment_config"></a>
### Nested Schema for `instance_groups.scheduled_update_config.deployment_config`

Optional:

- `auto_rollback_configuration` (Attributes List) An array that contains the alarms that SageMaker monitors to know whether to roll back the AMI update. (see [below for nested schema](#nestedatt--instance_groups--scheduled_update_config--deployment_config--auto_rollback_configuration))
- `rolling_update_policy` (Attributes) The policy that SageMaker uses when updating the AMI versions of the cluster. (see [below for nested schema](#nestedatt--instance_groups--scheduled_update_config--deployment_config--rolling_update_policy))
- `wait_interval_in_seconds` (Number) The duration in seconds that SageMaker waits before updating more instances in the cluster.

<a id="nestedatt--instance_groups--scheduled_update_config--deployment_config--auto_rollback_configuration"></a>
### Nested Schema for `instance_groups.scheduled_update_config.deployment_config.auto_rollback_configuration`

Optional:

- `alarm_name` (String) The name of the alarm.


<a id="nestedatt--instance_groups--scheduled_update_config--deployment_config--rolling_update_policy"></a>
### Nested Schema for `instance_groups.scheduled_update_config.deployment_config.rolling_update_policy`

Optional:

- `maximum_batch_size` (Attributes) The configuration of the size measurements of the AMI update. Using this configuration, you can specify whether SageMaker should update your instance group by an amount or percentage of instances. (see [below for nested schema](#nestedatt--instance_groups--scheduled_update_config--deployment_config--rolling_update_policy--maximum_batch_size))
- `rollback_maximum_batch_size` (Attributes) The configuration of the size measurements of the AMI update. Using this configuration, you can specify whether SageMaker should update your instance group by an amount or percentage of instances. (see [below for nested schema](#nestedatt--instance_groups--scheduled_update_config--deployment_config--rolling_update_policy--rollback_maximum_batch_size))

<a id="nestedatt--instance_groups--scheduled_update_config--deployment_config--rolling_update_policy--maximum_batch_size"></a>
### Nested Schema for `instance_groups.scheduled_update_config.deployment_config.rolling_update_policy.maximum_batch_size`

Optional:

- `type` (String) Specifies whether SageMaker should process the update by amount or percentage of instances.
- `value` (Number) Specifies the amount or percentage of instances SageMaker updates at a time.


<a id="nestedatt--instance_groups--scheduled_update_config--deployment_config--rolling_update_policy--rollback_maximum_batch_size"></a>
### Nested Schema for `instance_groups.scheduled_update_config.deployment_config.rolling_update_policy.rollback_maximum_batch_size`

Optional:

- `type` (String) Specifies whether SageMaker should process the update by amount or percentage of instances.
- `value` (Number) Specifies the amount or percentage of instances SageMaker updates at a time.






<a id="nestedatt--orchestrator"></a>
### Nested Schema for `orchestrator`

Optional:

- `eks` (Attributes) Specifies parameter(s) related to EKS as orchestrator, e.g. the EKS cluster nodes will attach to, (see [below for nested schema](#nestedatt--orchestrator--eks))

<a id="nestedatt--orchestrator--eks"></a>
### Nested Schema for `orchestrator.eks`

Optional:

- `cluster_arn` (String) The ARN of the EKS cluster, such as arn:aws:eks:us-west-2:123456789012:cluster/my-eks-cluster



<a id="nestedatt--restricted_instance_groups"></a>
### Nested Schema for `restricted_instance_groups`

Optional:

- `current_count` (Number) The number of instances that are currently in the restricted instance group of a SageMaker HyperPod cluster.
- `environment_config` (Attributes) The configuration for the restricted instance groups (RIG) environment. (see [below for nested schema](#nestedatt--restricted_instance_groups--environment_config))
- `execution_role` (String) The execution role for the instance group to assume.
- `instance_count` (Number) The number of instances you specified to add to the restricted instance group of a SageMaker HyperPod cluster.
- `instance_group_name` (String) The name of the instance group of a SageMaker HyperPod cluster.
- `instance_storage_configs` (Attributes List) The instance storage configuration for the instance group. (see [below for nested schema](#nestedatt--restricted_instance_groups--instance_storage_configs))
- `instance_type` (String) The instance type of the instance group of a SageMaker HyperPod cluster.
- `on_start_deep_health_checks` (List of String) Nodes will undergo advanced stress test to detect and replace faulty instances, based on the type of deep health check(s) passed in.
- `override_vpc_config` (Attributes) Specifies an Amazon Virtual Private Cloud (VPC) that your SageMaker jobs, hosted models, and compute resources have access to. You can control access to and from your resources by configuring a VPC. (see [below for nested schema](#nestedatt--restricted_instance_groups--override_vpc_config))
- `threads_per_core` (Number) The number you specified to TreadsPerCore in CreateCluster for enabling or disabling multithreading. For instance types that support multithreading, you can specify 1 for disabling multithreading and 2 for enabling multithreading.
- `training_plan_arn` (String) The Amazon Resource Name (ARN) of the training plan to use for this cluster restricted instance group. For more information about how to reserve GPU capacity for your SageMaker HyperPod clusters using Amazon SageMaker Training Plan, see CreateTrainingPlan.

<a id="nestedatt--restricted_instance_groups--environment_config"></a>
### Nested Schema for `restricted_instance_groups.environment_config`

Optional:

- `fsx_lustre_config` (Attributes) Configuration settings for an Amazon FSx for Lustre file system to be used with the cluster. (see [below for nested schema](#nestedatt--restricted_instance_groups--environment_config--fsx_lustre_config))

<a id="nestedatt--restricted_instance_groups--environment_config--fsx_lustre_config"></a>
### Nested Schema for `restricted_instance_groups.environment_config.fsx_lustre_config`

Optional:

- `per_unit_storage_throughput` (Number) The throughput capacity of the FSx for Lustre file system, measured in MB/s per TiB of storage.
- `size_in_gi_b` (Number) The storage capacity of the FSx for Lustre file system, specified in gibibytes (GiB).



<a id="nestedatt--restricted_instance_groups--instance_storage_configs"></a>
### Nested Schema for `restricted_instance_groups.instance_storage_configs`

Optional:

- `ebs_volume_config` (Attributes) Defines the configuration for attaching additional Amazon Elastic Block Store (EBS) volumes to the instances in the SageMaker HyperPod cluster instance group. The additional EBS volume is attached to each instance within the SageMaker HyperPod cluster instance group and mounted to /opt/sagemaker. (see [below for nested schema](#nestedatt--restricted_instance_groups--instance_storage_configs--ebs_volume_config))

<a id="nestedatt--restricted_instance_groups--instance_storage_configs--ebs_volume_config"></a>
### Nested Schema for `restricted_instance_groups.instance_storage_configs.ebs_volume_config`

Optional:

- `volume_size_in_gb` (Number) The size in gigabytes (GB) of the additional EBS volume to be attached to the instances in the SageMaker HyperPod cluster instance group. The additional EBS volume is attached to each instance within the SageMaker HyperPod cluster instance group and mounted to /opt/sagemaker.



<a id="nestedatt--restricted_instance_groups--override_vpc_config"></a>
### Nested Schema for `restricted_instance_groups.override_vpc_config`

Optional:

- `security_group_ids` (List of String) The VPC security group IDs, in the form sg-xxxxxxxx. Specify the security groups for the VPC that is specified in the Subnets field.
- `subnets` (List of String) The ID of the subnets in the VPC to which you want to connect your training job or model.



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.


<a id="nestedatt--vpc_config"></a>
### Nested Schema for `vpc_config`

Optional:

- `security_group_ids` (List of String) The VPC security group IDs, in the form sg-xxxxxxxx. Specify the security groups for the VPC that is specified in the Subnets field.
- `subnets` (List of String) The ID of the subnets in the VPC to which you want to connect your training job or model.

## Import

Import is supported using the following syntax:

In Terraform v1.12.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `identity` attribute, for example:

```terraform
import {
  to = awscc_sagemaker_cluster.example
  identity = {
    cluster_arn = "cluster_arn"
  }
}
```

<!-- schema generated by tfplugindocs -->
### Identity Schema

#### Required

- `cluster_arn` (String) The Amazon Resource Name (ARN) of the HyperPod Cluster

#### Optional

- `account_id` (String) AWS Account where this resource is managed
- `region` (String) Region where this resource is managed

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_sagemaker_cluster.example
  id = "cluster_arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_sagemaker_cluster.example "cluster_arn"
```
