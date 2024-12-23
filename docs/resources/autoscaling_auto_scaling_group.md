---
page_title: "awscc_autoscaling_auto_scaling_group Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::AutoScaling::AutoScalingGroup resource defines an Amazon EC2 Auto Scaling group, which is a collection of Amazon EC2 instances that are treated as a logical grouping for the purposes of automatic scaling and management.
  For more information about Amazon EC2 Auto Scaling, see the Amazon EC2 Auto Scaling User Guide https://docs.aws.amazon.com/autoscaling/ec2/userguide/what-is-amazon-ec2-auto-scaling.html.
  Amazon EC2 Auto Scaling configures instances launched as part of an Auto Scaling group using either a launch template https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html or a launch configuration. We strongly recommend that you do not use launch configurations. For more information, see Launch configurations https://docs.aws.amazon.com/autoscaling/ec2/userguide/launch-configurations.html in the Amazon EC2 Auto Scaling User Guide.
  For help migrating from launch configurations to launch templates, see Migrate CloudFormation stacks from launch configurations to launch templates https://docs.aws.amazon.com/autoscaling/ec2/userguide/migrate-launch-configurations-with-cloudformation.html in the Amazon EC2 Auto Scaling User Guide.
---

# awscc_autoscaling_auto_scaling_group (Resource)

The ``AWS::AutoScaling::AutoScalingGroup`` resource defines an Amazon EC2 Auto Scaling group, which is a collection of Amazon EC2 instances that are treated as a logical grouping for the purposes of automatic scaling and management. 
 For more information about Amazon EC2 Auto Scaling, see the [Amazon EC2 Auto Scaling User Guide](https://docs.aws.amazon.com/autoscaling/ec2/userguide/what-is-amazon-ec2-auto-scaling.html). 
  Amazon EC2 Auto Scaling configures instances launched as part of an Auto Scaling group using either a [launch template](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html) or a launch configuration. We strongly recommend that you do not use launch configurations. For more information, see [Launch configurations](https://docs.aws.amazon.com/autoscaling/ec2/userguide/launch-configurations.html) in the *Amazon EC2 Auto Scaling User Guide*.
 For help migrating from launch configurations to launch templates, see [Migrate CloudFormation stacks from launch configurations to launch templates](https://docs.aws.amazon.com/autoscaling/ec2/userguide/migrate-launch-configurations-with-cloudformation.html) in the *Amazon EC2 Auto Scaling User Guide*.

## Example Usage

### Basic

An Auto Scaling group and a launch template.

```terraform
resource "awscc_autoscaling_auto_scaling_group" "example" {
  max_size = "1"
  min_size = "0"

  launch_template = {
    version            = awscc_ec2_launch_template.example.latest_version_number
    launch_template_id = awscc_ec2_launch_template.example.id
  }

  desired_capacity    = "1"
  vpc_zone_identifier = [
    "subnetIdAz1",
    "subnetIdAz2",
    "subnetIdAz3"
  ]

  depends_on = [
    awscc_ec2_launch_template.example
  ]
}

resource "awscc_ec2_launch_template" "example" {
  launch_template_data = {
    image_id      = data.aws_ami.amazon_linux.id
    instance_type = "t2.micro"
  }
  launch_template_name = "${data.aws_caller_identity.current.account_id}-launch-template"
}

data "aws_caller_identity" "current" {}

data "aws_ami" "amazon_linux" {
  most_recent = true
  owners      = ["amazon"]
  filter {
    name   = "name"
    values = ["amzn2-ami-hvm-*-gp2"]
  }

  filter {
    name   = "root-device-type"
    values = ["ebs"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  filter {
    name   = "architecture"
    values = ["x86_64"]
  }
}
```

### Metrics Collection

Auto Scaling group with CloudWatch monitoring enabled

```terraform
resource "awscc_autoscaling_auto_scaling_group" "example" {
  max_size = "1"
  min_size = "0"

  launch_template = {
    version            = awscc_ec2_launch_template.example.latest_version_number
    launch_template_id = awscc_ec2_launch_template.example.id
  }

  desired_capacity    = "1"
  
  vpc_zone_identifier = [
    "subnetIdAz1",
    "subnetIdAz2",
    "subnetIdAz3"
  ]

  metrics_collection = [{
    granularity = "1Minute"
    metrics = [
      "GroupMinSize",
      "GroupMaxSize"
    ]
  }]
}

resource "awscc_ec2_launch_template" "example" {
  launch_template_data = {
    image_id      = data.aws_ami.amazon_linux.id
    instance_type = "t2.micro"
  }
  launch_template_name = "${data.aws_caller_identity.current.account_id}-launch-template"
}

data "aws_caller_identity" "current" {}

data "aws_ami" "amazon_linux" {
  most_recent = true
  owners      = ["amazon"]
  filter {
    name   = "name"
    values = ["amzn2-ami-hvm-*-gp2"]
  }

  filter {
    name   = "root-device-type"
    values = ["ebs"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  filter {
    name   = "architecture"
    values = ["x86_64"]
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `max_size` (String) The maximum size of the group.
  With a mixed instances policy that uses instance weighting, Amazon EC2 Auto Scaling may need to go above ``MaxSize`` to meet your capacity requirements. In this event, Amazon EC2 Auto Scaling will never go above ``MaxSize`` by more than your largest instance weight (weights that define how many units each instance contributes to the desired capacity of the group).
- `min_size` (String) The minimum size of the group.

### Optional

- `auto_scaling_group_name` (String) The name of the Auto Scaling group. This name must be unique per Region per account.
 The name can contain any ASCII character 33 to 126 including most punctuation characters, digits, and upper and lowercased letters.
  You cannot use a colon (:) in the name.
- `availability_zone_distribution` (Attributes) The instance capacity distribution across Availability Zones. (see [below for nested schema](#nestedatt--availability_zone_distribution))
- `availability_zone_impairment_policy` (Attributes) The Availability Zone impairment policy. (see [below for nested schema](#nestedatt--availability_zone_impairment_policy))
- `availability_zones` (List of String) A list of Availability Zones where instances in the Auto Scaling group can be created. Used for launching into the default VPC subnet in each Availability Zone when not using the ``VPCZoneIdentifier`` property, or for attaching a network interface when an existing network interface ID is specified in a launch template.
- `capacity_rebalance` (Boolean) Indicates whether Capacity Rebalancing is enabled. Otherwise, Capacity Rebalancing is disabled. When you turn on Capacity Rebalancing, Amazon EC2 Auto Scaling attempts to launch a Spot Instance whenever Amazon EC2 notifies that a Spot Instance is at an elevated risk of interruption. After launching a new instance, it then terminates an old instance. For more information, see [Use Capacity Rebalancing to handle Amazon EC2 Spot Interruptions](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-capacity-rebalancing.html) in the in the *Amazon EC2 Auto Scaling User Guide*.
- `capacity_reservation_specification` (Attributes) (see [below for nested schema](#nestedatt--capacity_reservation_specification))
- `context` (String) Reserved.
- `cooldown` (String) *Only needed if you use simple scaling policies.* 
 The amount of time, in seconds, between one scaling activity ending and another one starting due to simple scaling policies. For more information, see [Scaling cooldowns for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-scaling-cooldowns.html) in the *Amazon EC2 Auto Scaling User Guide*.
 Default: ``300`` seconds
- `default_instance_warmup` (Number) The amount of time, in seconds, until a new instance is considered to have finished initializing and resource consumption to become stable after it enters the ``InService`` state. 
 During an instance refresh, Amazon EC2 Auto Scaling waits for the warm-up period after it replaces an instance before it moves on to replacing the next instance. Amazon EC2 Auto Scaling also waits for the warm-up period before aggregating the metrics for new instances with existing instances in the Amazon CloudWatch metrics that are used for scaling, resulting in more reliable usage data. For more information, see [Set the default instance warmup for an Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-default-instance-warmup.html) in the *Amazon EC2 Auto Scaling User Guide*.
  To manage various warm-up settings at the group level, we recommend that you set the default instance warmup, *even if it is set to 0 seconds*. To remove a value that you previously set, include the property but specify ``-1`` for the value. However, we strongly recommend keeping the default instance warmup enabled by specifying a value of ``0`` or other nominal value.
  Default: None
- `desired_capacity` (String) The desired capacity is the initial capacity of the Auto Scaling group at the time of its creation and the capacity it attempts to maintain. It can scale beyond this capacity if you configure automatic scaling.
 The number must be greater than or equal to the minimum size of the group and less than or equal to the maximum size of the group. If you do not specify a desired capacity when creating the stack, the default is the minimum size of the group.
 CloudFormation marks the Auto Scaling group as successful (by setting its status to CREATE_COMPLETE) when the desired capacity is reached. However, if a maximum Spot price is set in the launch template or launch configuration that you specified, then desired capacity is not used as a criteria for success. Whether your request is fulfilled depends on Spot Instance capacity and your maximum price.
- `desired_capacity_type` (String) The unit of measurement for the value specified for desired capacity. Amazon EC2 Auto Scaling supports ``DesiredCapacityType`` for attribute-based instance type selection only. For more information, see [Create a mixed instances group using attribute-based instance type selection](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-mixed-instances-group-attribute-based-instance-type-selection.html) in the *Amazon EC2 Auto Scaling User Guide*.
 By default, Amazon EC2 Auto Scaling specifies ``units``, which translates into number of instances.
 Valid values: ``units`` | ``vcpu`` | ``memory-mib``
- `health_check_grace_period` (Number) The amount of time, in seconds, that Amazon EC2 Auto Scaling waits before checking the health status of an EC2 instance that has come into service and marking it unhealthy due to a failed health check. This is useful if your instances do not immediately pass their health checks after they enter the ``InService`` state. For more information, see [Set the health check grace period for an Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/health-check-grace-period.html) in the *Amazon EC2 Auto Scaling User Guide*.
 Default: ``0`` seconds
- `health_check_type` (String) A comma-separated value string of one or more health check types.
 The valid values are ``EC2``, ``EBS``, ``ELB``, and ``VPC_LATTICE``. ``EC2`` is the default health check and cannot be disabled. For more information, see [Health checks for instances in an Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-health-checks.html) in the *Amazon EC2 Auto Scaling User Guide*.
 Only specify ``EC2`` if you must clear a value that was previously set.
- `instance_id` (String) The ID of the instance used to base the launch configuration on. For more information, see [Create an Auto Scaling group using an EC2 instance](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-asg-from-instance.html) in the *Amazon EC2 Auto Scaling User Guide*.
 If you specify ``LaunchTemplate``, ``MixedInstancesPolicy``, or ``LaunchConfigurationName``, don't specify ``InstanceId``.
- `instance_maintenance_policy` (Attributes) An instance maintenance policy. For more information, see [Set instance maintenance policy](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-instance-maintenance-policy.html) in the *Amazon EC2 Auto Scaling User Guide*. (see [below for nested schema](#nestedatt--instance_maintenance_policy))
- `launch_configuration_name` (String) The name of the launch configuration to use to launch instances.
 Required only if you don't specify ``LaunchTemplate``, ``MixedInstancesPolicy``, or ``InstanceId``.
- `launch_template` (Attributes) Information used to specify the launch template and version to use to launch instances. You can alternatively associate a launch template to the Auto Scaling group by specifying a ``MixedInstancesPolicy``. For more information about creating launch templates, see [Create a launch template for an Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-launch-template.html) in the *Amazon EC2 Auto Scaling User Guide*.
 If you omit this property, you must specify ``MixedInstancesPolicy``, ``LaunchConfigurationName``, or ``InstanceId``. (see [below for nested schema](#nestedatt--launch_template))
- `lifecycle_hook_specification_list` (Attributes List) One or more lifecycle hooks to add to the Auto Scaling group before instances are launched. (see [below for nested schema](#nestedatt--lifecycle_hook_specification_list))
- `load_balancer_names` (List of String) A list of Classic Load Balancers associated with this Auto Scaling group. For Application Load Balancers, Network Load Balancers, and Gateway Load Balancers, specify the ``TargetGroupARNs`` property instead.
- `max_instance_lifetime` (Number) The maximum amount of time, in seconds, that an instance can be in service. The default is null. If specified, the value must be either 0 or a number equal to or greater than 86,400 seconds (1 day). For more information, see [Replace Auto Scaling instances based on maximum instance lifetime](https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-max-instance-lifetime.html) in the *Amazon EC2 Auto Scaling User Guide*.
- `metrics_collection` (Attributes List) Enables the monitoring of group metrics of an Auto Scaling group. By default, these metrics are disabled. (see [below for nested schema](#nestedatt--metrics_collection))
- `mixed_instances_policy` (Attributes) An embedded object that specifies a mixed instances policy.
 The policy includes properties that not only define the distribution of On-Demand Instances and Spot Instances, the maximum price to pay for Spot Instances (optional), and how the Auto Scaling group allocates instance types to fulfill On-Demand and Spot capacities, but also the properties that specify the instance configuration information—the launch template and instance types. The policy can also include a weight for each instance type and different launch templates for individual instance types.
 For more information, see [Auto Scaling groups with multiple instance types and purchase options](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups.html) in the *Amazon EC2 Auto Scaling User Guide*. (see [below for nested schema](#nestedatt--mixed_instances_policy))
- `new_instances_protected_from_scale_in` (Boolean) Indicates whether newly launched instances are protected from termination by Amazon EC2 Auto Scaling when scaling in. For more information about preventing instances from terminating on scale in, see [Use instance scale-in protection](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-instance-protection.html) in the *Amazon EC2 Auto Scaling User Guide*.
- `notification_configuration` (Attributes) A structure that specifies an Amazon SNS notification configuration for the ``NotificationConfigurations`` property of the [AWS::AutoScaling::AutoScalingGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html) resource.
 For an example template snippet, see [Configure Amazon EC2 Auto Scaling resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-ec2-auto-scaling.html).
 For more information, see [Get Amazon SNS notifications when your Auto Scaling group scales](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ASGettingNotifications.html) in the *Amazon EC2 Auto Scaling User Guide*. (see [below for nested schema](#nestedatt--notification_configuration))
- `notification_configurations` (Attributes List) Configures an Auto Scaling group to send notifications when specified events take place. (see [below for nested schema](#nestedatt--notification_configurations))
- `placement_group` (String) The name of the placement group into which to launch your instances. For more information, see [Placement groups](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html) in the *Amazon EC2 User Guide for Linux Instances*.
  A *cluster* placement group is a logical grouping of instances within a single Availability Zone. You cannot specify multiple Availability Zones and a cluster placement group.
- `service_linked_role_arn` (String) The Amazon Resource Name (ARN) of the service-linked role that the Auto Scaling group uses to call other AWS service on your behalf. By default, Amazon EC2 Auto Scaling uses a service-linked role named ``AWSServiceRoleForAutoScaling``, which it creates if it does not exist. For more information, see [Service-linked roles](https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-service-linked-role.html) in the *Amazon EC2 Auto Scaling User Guide*.
- `skip_zonal_shift_validation` (Boolean)
- `tags` (Attributes List) One or more tags. You can tag your Auto Scaling group and propagate the tags to the Amazon EC2 instances it launches. Tags are not propagated to Amazon EBS volumes. To add tags to Amazon EBS volumes, specify the tags in a launch template but use caution. If the launch template specifies an instance tag with a key that is also specified for the Auto Scaling group, Amazon EC2 Auto Scaling overrides the value of that instance tag with the value specified by the Auto Scaling group. For more information, see [Tag Auto Scaling groups and instances](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-tagging.html) in the *Amazon EC2 Auto Scaling User Guide*. (see [below for nested schema](#nestedatt--tags))
- `target_group_ar_ns` (List of String) The Amazon Resource Names (ARN) of the Elastic Load Balancing target groups to associate with the Auto Scaling group. Instances are registered as targets with the target groups. The target groups receive incoming traffic and route requests to one or more registered targets. For more information, see [Use Elastic Load Balancing to distribute traffic across the instances in your Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-load-balancer.html) in the *Amazon EC2 Auto Scaling User Guide*.
- `termination_policies` (List of String) A policy or a list of policies that are used to select the instance to terminate. These policies are executed in the order that you list them. For more information, see [Configure termination policies for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-termination-policies.html) in the *Amazon EC2 Auto Scaling User Guide*.
 Valid values: ``Default`` | ``AllocationStrategy`` | ``ClosestToNextInstanceHour`` | ``NewestInstance`` | ``OldestInstance`` | ``OldestLaunchConfiguration`` | ``OldestLaunchTemplate`` | ``arn:aws:lambda:region:account-id:function:my-function:my-alias``
- `traffic_sources` (Attributes Set) The traffic sources associated with this Auto Scaling group. (see [below for nested schema](#nestedatt--traffic_sources))
- `vpc_zone_identifier` (List of String) A list of subnet IDs for a virtual private cloud (VPC) where instances in the Auto Scaling group can be created.
 If this resource specifies public subnets and is also in a VPC that is defined in the same stack template, you must use the [DependsOn attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) to declare a dependency on the [VPC-gateway attachment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-gateway-attachment.html).
  When you update ``VPCZoneIdentifier``, this retains the same Auto Scaling group and replaces old instances with new ones, according to the specified subnets. You can optionally specify how CloudFormation handles these updates by using an [UpdatePolicy attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html).
  Required to launch instances into a nondefault VPC. If you specify ``VPCZoneIdentifier`` with ``AvailabilityZones``, the subnets that you specify for this property must reside in those Availability Zones.

### Read-Only

- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--availability_zone_distribution"></a>
### Nested Schema for `availability_zone_distribution`

Optional:

- `capacity_distribution_strategy` (String) If launches fail in an Availability Zone, the following strategies are available. The default is ``balanced-best-effort``. 
  +   ``balanced-only`` - If launches fail in an Availability Zone, Auto Scaling will continue to attempt to launch in the unhealthy zone to preserve a balanced distribution.
  +   ``balanced-best-effort`` - If launches fail in an Availability Zone, Auto Scaling will attempt to launch in another healthy Availability Zone instead.


<a id="nestedatt--availability_zone_impairment_policy"></a>
### Nested Schema for `availability_zone_impairment_policy`

Optional:

- `impaired_zone_health_check_behavior` (String) Specifies the health check behavior for the impaired Availability Zone in an active zonal shift. If you select ``Replace unhealthy``, instances that appear unhealthy will be replaced in all Availability Zones. If you select ``Ignore unhealthy``, instances will not be replaced in the Availability Zone with the active zonal shift. For more information, see [Auto Scaling group zonal shift](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-zonal-shift.html) in the *Amazon EC2 Auto Scaling User Guide*.
- `zonal_shift_enabled` (Boolean) If ``true``, enable zonal shift for your Auto Scaling group.


<a id="nestedatt--capacity_reservation_specification"></a>
### Nested Schema for `capacity_reservation_specification`

Optional:

- `capacity_reservation_preference` (String)
- `capacity_reservation_target` (Attributes) (see [below for nested schema](#nestedatt--capacity_reservation_specification--capacity_reservation_target))

<a id="nestedatt--capacity_reservation_specification--capacity_reservation_target"></a>
### Nested Schema for `capacity_reservation_specification.capacity_reservation_target`

Optional:

- `capacity_reservation_ids` (List of String)
- `capacity_reservation_resource_group_arns` (List of String)



<a id="nestedatt--instance_maintenance_policy"></a>
### Nested Schema for `instance_maintenance_policy`

Optional:

- `max_healthy_percentage` (Number) Specifies the upper threshold as a percentage of the desired capacity of the Auto Scaling group. It represents the maximum percentage of the group that can be in service and healthy, or pending, to support your workload when replacing instances. Value range is 100 to 200. To clear a previously set value, specify a value of ``-1``.
 Both ``MinHealthyPercentage`` and ``MaxHealthyPercentage`` must be specified, and the difference between them cannot be greater than 100. A large range increases the number of instances that can be replaced at the same time.
- `min_healthy_percentage` (Number) Specifies the lower threshold as a percentage of the desired capacity of the Auto Scaling group. It represents the minimum percentage of the group to keep in service, healthy, and ready to use to support your workload when replacing instances. Value range is 0 to 100. To clear a previously set value, specify a value of ``-1``.


<a id="nestedatt--launch_template"></a>
### Nested Schema for `launch_template`

Optional:

- `launch_template_id` (String) The ID of the launch template.
 You must specify the ``LaunchTemplateID`` or the ``LaunchTemplateName``, but not both.
- `launch_template_name` (String) The name of the launch template.
 You must specify the ``LaunchTemplateName`` or the ``LaunchTemplateID``, but not both.
- `version` (String) The version number of the launch template.
 Specifying ``$Latest`` or ``$Default`` for the template version number is not supported. However, you can specify ``LatestVersionNumber`` or ``DefaultVersionNumber`` using the ``Fn::GetAtt`` intrinsic function. For more information, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).
  For an example of using the ``Fn::GetAtt`` function, see the [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#aws-resource-autoscaling-autoscalinggroup--examples) section of the ``AWS::AutoScaling::AutoScalingGroup`` resource.


<a id="nestedatt--lifecycle_hook_specification_list"></a>
### Nested Schema for `lifecycle_hook_specification_list`

Optional:

- `default_result` (String) The action the Auto Scaling group takes when the lifecycle hook timeout elapses or if an unexpected failure occurs. The default value is ``ABANDON``.
 Valid values: ``CONTINUE`` | ``ABANDON``
- `heartbeat_timeout` (Number) The maximum time, in seconds, that can elapse before the lifecycle hook times out. The range is from ``30`` to ``7200`` seconds. The default value is ``3600`` seconds (1 hour).
- `lifecycle_hook_name` (String) The name of the lifecycle hook.
- `lifecycle_transition` (String) The lifecycle transition. For Auto Scaling groups, there are two major lifecycle transitions.
  +  To create a lifecycle hook for scale-out events, specify ``autoscaling:EC2_INSTANCE_LAUNCHING``.
  +  To create a lifecycle hook for scale-in events, specify ``autoscaling:EC2_INSTANCE_TERMINATING``.
- `notification_metadata` (String) Additional information that you want to include any time Amazon EC2 Auto Scaling sends a message to the notification target.
- `notification_target_arn` (String) The Amazon Resource Name (ARN) of the notification target that Amazon EC2 Auto Scaling sends notifications to when an instance is in a wait state for the lifecycle hook. You can specify an Amazon SNS topic or an Amazon SQS queue.
- `role_arn` (String) The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target. For information about creating this role, see [Prepare to add a lifecycle hook to your Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/prepare-for-lifecycle-notifications.html) in the *Amazon EC2 Auto Scaling User Guide*.
 Valid only if the notification target is an Amazon SNS topic or an Amazon SQS queue.


<a id="nestedatt--metrics_collection"></a>
### Nested Schema for `metrics_collection`

Optional:

- `granularity` (String) The frequency at which Amazon EC2 Auto Scaling sends aggregated data to CloudWatch. The only valid value is ``1Minute``.
- `metrics` (List of String) Identifies the metrics to enable.
 You can specify one or more of the following metrics:
  +   ``GroupMinSize`` 
  +   ``GroupMaxSize`` 
  +   ``GroupDesiredCapacity`` 
  +   ``GroupInServiceInstances`` 
  +   ``GroupPendingInstances`` 
  +   ``GroupStandbyInstances`` 
  +   ``GroupTerminatingInstances`` 
  +   ``GroupTotalInstances`` 
  +   ``GroupInServiceCapacity`` 
  +   ``GroupPendingCapacity`` 
  +   ``GroupStandbyCapacity`` 
  +   ``GroupTerminatingCapacity`` 
  +   ``GroupTotalCapacity`` 
  +   ``WarmPoolDesiredCapacity`` 
  +   ``WarmPoolWarmedCapacity`` 
  +   ``WarmPoolPendingCapacity`` 
  +   ``WarmPoolTerminatingCapacity`` 
  +   ``WarmPoolTotalCapacity`` 
  +   ``GroupAndWarmPoolDesiredCapacity`` 
  +   ``GroupAndWarmPoolTotalCapacity`` 
  
 If you specify ``Granularity`` and don't specify any metrics, all metrics are enabled.
 For more information, see [Amazon CloudWatch metrics for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-metrics.html) in the *Amazon EC2 Auto Scaling User Guide*.


<a id="nestedatt--mixed_instances_policy"></a>
### Nested Schema for `mixed_instances_policy`

Optional:

- `instances_distribution` (Attributes) The instances distribution. (see [below for nested schema](#nestedatt--mixed_instances_policy--instances_distribution))
- `launch_template` (Attributes) One or more launch templates and the instance types (overrides) that are used to launch EC2 instances to fulfill On-Demand and Spot capacities. (see [below for nested schema](#nestedatt--mixed_instances_policy--launch_template))

<a id="nestedatt--mixed_instances_policy--instances_distribution"></a>
### Nested Schema for `mixed_instances_policy.instances_distribution`

Optional:

- `on_demand_allocation_strategy` (String) The allocation strategy to apply to your On-Demand Instances when they are launched. Possible instance types are determined by the launch template overrides that you specify.
 The following lists the valid values:
  + lowest-price Uses price to determine which instance types are the highest priority, launching the lowest priced instance types within an Availability Zone first. This is the default value for Auto Scaling groups that specify InstanceRequirements. + prioritized You set the order of instance types for the launch template overrides from highest to lowest priority (from first to last in the list). Amazon EC2 Auto Scaling launches your highest priority instance types first. If all your On-Demand capacity cannot be fulfilled using your highest priority instance type, then Amazon EC2 Auto Scaling launches the remaining capacity using the second priority instance type, and so on. This is the default value for Auto Scaling groups that don't specify InstanceRequirements and cannot be used for groups that do.
- `on_demand_base_capacity` (Number) The minimum amount of the Auto Scaling group's capacity that must be fulfilled by On-Demand Instances. This base portion is launched first as your group scales.
 This number has the same unit of measurement as the group's desired capacity. If you change the default unit of measurement (number of instances) by specifying weighted capacity values in your launch template overrides list, or by changing the default desired capacity type setting of the group, you must specify this number using the same unit of measurement.
 Default: 0
  An update to this setting means a gradual replacement of instances to adjust the current On-Demand Instance levels. When replacing instances, Amazon EC2 Auto Scaling launches new instances before terminating the previous ones.
- `on_demand_percentage_above_base_capacity` (Number) Controls the percentages of On-Demand Instances and Spot Instances for your additional capacity beyond ``OnDemandBaseCapacity``. Expressed as a number (for example, 20 specifies 20% On-Demand Instances, 80% Spot Instances). If set to 100, only On-Demand Instances are used.
 Default: 100
  An update to this setting means a gradual replacement of instances to adjust the current On-Demand and Spot Instance levels for your additional capacity higher than the base capacity. When replacing instances, Amazon EC2 Auto Scaling launches new instances before terminating the previous ones.
- `spot_allocation_strategy` (String) The allocation strategy to apply to your Spot Instances when they are launched. Possible instance types are determined by the launch template overrides that you specify.
 The following lists the valid values:
  + capacity-optimized Requests Spot Instances using pools that are optimally chosen based on the available Spot capacity. This strategy has the lowest risk of interruption. To give certain instance types a higher chance of launching first, use capacity-optimized-prioritized. + capacity-optimized-prioritized You set the order of instance types for the launch template overrides from highest to lowest priority (from first to last in the list). Amazon EC2 Auto Scaling honors the instance type priorities on a best effort basis but optimizes for capacity first. Note that if the On-Demand allocation strategy is set to prioritized, the same priority is applied when fulfilling On-Demand capacity. This is not a valid value for Auto Scaling groups that specify InstanceRequirements. + lowest-price Requests Spot Instances using the lowest priced pools within an Availability Zone, across the number of Spot pools that you specify for the SpotInstancePools property. To ensure that your desired capacity is met, you might receive Spot Instances from several pools. This is the default value, but it might lead to high interruption rates because this strategy only considers instance price and not available capacity. + price-capacity-optimized (recommended) The price and capacity optimized allocation strategy looks at both price and capacity to select the Spot Instance pools that are the least likely to be interrupted and have the lowest possible price.
- `spot_instance_pools` (Number) The number of Spot Instance pools across which to allocate your Spot Instances. The Spot pools are determined from the different instance types in the overrides. Valid only when the ``SpotAllocationStrategy`` is ``lowest-price``. Value must be in the range of 1–20.
 Default: 2
- `spot_max_price` (String) The maximum price per unit hour that you are willing to pay for a Spot Instance. If your maximum price is lower than the Spot price for the instance types that you selected, your Spot Instances are not launched. We do not recommend specifying a maximum price because it can lead to increased interruptions. When Spot Instances launch, you pay the current Spot price. To remove a maximum price that you previously set, include the property but specify an empty string ("") for the value.
  If you specify a maximum price, your instances will be interrupted more frequently than if you do not specify one.
  Valid Range: Minimum value of 0.001


<a id="nestedatt--mixed_instances_policy--launch_template"></a>
### Nested Schema for `mixed_instances_policy.launch_template`

Optional:

- `launch_template_specification` (Attributes) The launch template. (see [below for nested schema](#nestedatt--mixed_instances_policy--launch_template--launch_template_specification))
- `overrides` (Attributes List) Any properties that you specify override the same properties in the launch template. (see [below for nested schema](#nestedatt--mixed_instances_policy--launch_template--overrides))

<a id="nestedatt--mixed_instances_policy--launch_template--launch_template_specification"></a>
### Nested Schema for `mixed_instances_policy.launch_template.launch_template_specification`

Optional:

- `launch_template_id` (String) The ID of the launch template.
 You must specify the ``LaunchTemplateID`` or the ``LaunchTemplateName``, but not both.
- `launch_template_name` (String) The name of the launch template.
 You must specify the ``LaunchTemplateName`` or the ``LaunchTemplateID``, but not both.
- `version` (String) The version number of the launch template.
 Specifying ``$Latest`` or ``$Default`` for the template version number is not supported. However, you can specify ``LatestVersionNumber`` or ``DefaultVersionNumber`` using the ``Fn::GetAtt`` intrinsic function. For more information, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).
  For an example of using the ``Fn::GetAtt`` function, see the [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#aws-resource-autoscaling-autoscalinggroup--examples) section of the ``AWS::AutoScaling::AutoScalingGroup`` resource.


<a id="nestedatt--mixed_instances_policy--launch_template--overrides"></a>
### Nested Schema for `mixed_instances_policy.launch_template.overrides`

Optional:

- `instance_requirements` (Attributes) The instance requirements. Amazon EC2 Auto Scaling uses your specified requirements to identify instance types. Then, it uses your On-Demand and Spot allocation strategies to launch instances from these instance types.
 You can specify up to four separate sets of instance requirements per Auto Scaling group. This is useful for provisioning instances from different Amazon Machine Images (AMIs) in the same Auto Scaling group. To do this, create the AMIs and create a new launch template for each AMI. Then, create a compatible set of instance requirements for each launch template. 
  If you specify ``InstanceRequirements``, you can't specify ``InstanceType``. (see [below for nested schema](#nestedatt--mixed_instances_policy--launch_template--overrides--instance_requirements))
- `instance_type` (String) The instance type, such as ``m3.xlarge``. You must specify an instance type that is supported in your requested Region and Availability Zones. For more information, see [Instance types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the *Amazon EC2 User Guide for Linux Instances*.
 You can specify up to 40 instance types per Auto Scaling group.
- `launch_template_specification` (Attributes) Provides a launch template for the specified instance type or set of instance requirements. For example, some instance types might require a launch template with a different AMI. If not provided, Amazon EC2 Auto Scaling uses the launch template that's specified in the ``LaunchTemplate`` definition. For more information, see [Specifying a different launch template for an instance type](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups-launch-template-overrides.html) in the *Amazon EC2 Auto Scaling User Guide*. 
 You can specify up to 20 launch templates per Auto Scaling group. The launch templates specified in the overrides and in the ``LaunchTemplate`` definition count towards this limit. (see [below for nested schema](#nestedatt--mixed_instances_policy--launch_template--overrides--launch_template_specification))
- `weighted_capacity` (String) If you provide a list of instance types to use, you can specify the number of capacity units provided by each instance type in terms of virtual CPUs, memory, storage, throughput, or other relative performance characteristic. When a Spot or On-Demand Instance is launched, the capacity units count toward the desired capacity. Amazon EC2 Auto Scaling launches instances until the desired capacity is totally fulfilled, even if this results in an overage. For example, if there are two units remaining to fulfill capacity, and Amazon EC2 Auto Scaling can only launch an instance with a ``WeightedCapacity`` of five units, the instance is launched, and the desired capacity is exceeded by three units. For more information, see [Configure instance weighting for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups-instance-weighting.html) in the *Amazon EC2 Auto Scaling User Guide*. Value must be in the range of 1-999. 
 If you specify a value for ``WeightedCapacity`` for one instance type, you must specify a value for ``WeightedCapacity`` for all of them.
  Every Auto Scaling group has three size parameters (``DesiredCapacity``, ``MaxSize``, and ``MinSize``). Usually, you set these sizes based on a specific number of instances. However, if you configure a mixed instances policy that defines weights for the instance types, you must specify these sizes with the same units that you use for weighting instances.

<a id="nestedatt--mixed_instances_policy--launch_template--overrides--instance_requirements"></a>
### Nested Schema for `mixed_instances_policy.launch_template.overrides.instance_requirements`

Optional:

- `accelerator_count` (Attributes) The minimum and maximum number of accelerators (GPUs, FPGAs, or AWS Inferentia chips) for an instance type.
 To exclude accelerator-enabled instance types, set ``Max`` to ``0``.
 Default: No minimum or maximum limits (see [below for nested schema](#nestedatt--mixed_instances_policy--launch_template--overrides--instance_requirements--accelerator_count))
- `accelerator_manufacturers` (Set of String) Indicates whether instance types must have accelerators by specific manufacturers.
  +  For instance types with NVIDIA devices, specify ``nvidia``.
  +  For instance types with AMD devices, specify ``amd``.
  +  For instance types with AWS devices, specify ``amazon-web-services``.
  +  For instance types with Xilinx devices, specify ``xilinx``.
  
 Default: Any manufacturer
- `accelerator_names` (Set of String) Lists the accelerators that must be on an instance type.
  +  For instance types with NVIDIA A100 GPUs, specify ``a100``.
  +  For instance types with NVIDIA V100 GPUs, specify ``v100``.
  +  For instance types with NVIDIA K80 GPUs, specify ``k80``.
  +  For instance types with NVIDIA T4 GPUs, specify ``t4``.
  +  For instance types with NVIDIA M60 GPUs, specify ``m60``.
  +  For instance types with AMD Radeon Pro V520 GPUs, specify ``radeon-pro-v520``.
  +  For instance types with Xilinx VU9P FPGAs, specify ``vu9p``.
  
 Default: Any accelerator
- `accelerator_total_memory_mi_b` (Attributes) The minimum and maximum total memory size for the accelerators on an instance type, in MiB.
 Default: No minimum or maximum limits (see [below for nested schema](#nestedatt--mixed_instances_policy--launch_template--overrides--instance_requirements--accelerator_total_memory_mi_b))
- `accelerator_types` (Set of String) Lists the accelerator types that must be on an instance type.
  +  For instance types with GPU accelerators, specify ``gpu``.
  +  For instance types with FPGA accelerators, specify ``fpga``.
  +  For instance types with inference accelerators, specify ``inference``.
  
 Default: Any accelerator type
- `allowed_instance_types` (List of String) The instance types to apply your specified attributes against. All other instance types are ignored, even if they match your specified attributes.
 You can use strings with one or more wild cards, represented by an asterisk (``*``), to allow an instance type, size, or generation. The following are examples: ``m5.8xlarge``, ``c5*.*``, ``m5a.*``, ``r*``, ``*3*``.
 For example, if you specify ``c5*``, Amazon EC2 Auto Scaling will allow the entire C5 instance family, which includes all C5a and C5n instance types. If you specify ``m5a.*``, Amazon EC2 Auto Scaling will allow all the M5a instance types, but not the M5n instance types.
  If you specify ``AllowedInstanceTypes``, you can't specify ``ExcludedInstanceTypes``.
  Default: All instance types
- `bare_metal` (String) Indicates whether bare metal instance types are included, excluded, or required.
 Default: ``excluded``
- `baseline_ebs_bandwidth_mbps` (Attributes) The minimum and maximum baseline bandwidth performance for an instance type, in Mbps. For more information, see [Amazon EBS–optimized instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-optimized.html) in the *Amazon EC2 User Guide for Linux Instances*.
 Default: No minimum or maximum limits (see [below for nested schema](#nestedatt--mixed_instances_policy--launch_template--overrides--instance_requirements--baseline_ebs_bandwidth_mbps))
- `baseline_performance_factors` (Attributes) (see [below for nested schema](#nestedatt--mixed_instances_policy--launch_template--overrides--instance_requirements--baseline_performance_factors))
- `burstable_performance` (String) Indicates whether burstable performance instance types are included, excluded, or required. For more information, see [Burstable performance instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html) in the *Amazon EC2 User Guide for Linux Instances*.
 Default: ``excluded``
- `cpu_manufacturers` (Set of String) Lists which specific CPU manufacturers to include.
  +  For instance types with Intel CPUs, specify ``intel``.
  +  For instance types with AMD CPUs, specify ``amd``.
  +  For instance types with AWS CPUs, specify ``amazon-web-services``.
  
  Don't confuse the CPU hardware manufacturer with the CPU hardware architecture. Instances will be launched with a compatible CPU architecture based on the Amazon Machine Image (AMI) that you specify in your launch template. 
  Default: Any manufacturer
- `excluded_instance_types` (List of String) The instance types to exclude. You can use strings with one or more wild cards, represented by an asterisk (``*``), to exclude an instance family, type, size, or generation. The following are examples: ``m5.8xlarge``, ``c5*.*``, ``m5a.*``, ``r*``, ``*3*``. 
 For example, if you specify ``c5*``, you are excluding the entire C5 instance family, which includes all C5a and C5n instance types. If you specify ``m5a.*``, Amazon EC2 Auto Scaling will exclude all the M5a instance types, but not the M5n instance types.
  If you specify ``ExcludedInstanceTypes``, you can't specify ``AllowedInstanceTypes``.
  Default: No excluded instance types
- `instance_generations` (Set of String) Indicates whether current or previous generation instance types are included.
  +  For current generation instance types, specify ``current``. The current generation includes EC2 instance types currently recommended for use. This typically includes the latest two to three generations in each instance family. For more information, see [Instance types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the *Amazon EC2 User Guide for Linux Instances*.
  +  For previous generation instance types, specify ``previous``.
  
 Default: Any current or previous generation
- `local_storage` (String) Indicates whether instance types with instance store volumes are included, excluded, or required. For more information, see [Amazon EC2 instance store](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html) in the *Amazon EC2 User Guide for Linux Instances*.
 Default: ``included``
- `local_storage_types` (Set of String) Indicates the type of local storage that is required.
  +  For instance types with hard disk drive (HDD) storage, specify ``hdd``.
  +  For instance types with solid state drive (SSD) storage, specify ``ssd``.
  
 Default: Any local storage type
- `max_spot_price_as_percentage_of_optimal_on_demand_price` (Number) [Price protection] The price protection threshold for Spot Instances, as a percentage of an identified On-Demand price. The identified On-Demand price is the price of the lowest priced current generation C, M, or R instance type with your specified attributes. If no current generation C, M, or R instance type matches your attributes, then the identified price is from either the lowest priced current generation instance types or, failing that, the lowest priced previous generation instance types that match your attributes. When Amazon EC2 Auto Scaling selects instance types with your attributes, we will exclude instance types whose price exceeds your specified threshold.
 The parameter accepts an integer, which Amazon EC2 Auto Scaling interprets as a percentage.
 If you set ``DesiredCapacityType`` to ``vcpu`` or ``memory-mib``, the price protection threshold is based on the per-vCPU or per-memory price instead of the per instance price. 
  Only one of ``SpotMaxPricePercentageOverLowestPrice`` or ``MaxSpotPriceAsPercentageOfOptimalOnDemandPrice`` can be specified. If you don't specify either, Amazon EC2 Auto Scaling will automatically apply optimal price protection to consistently select from a wide range of instance types. To indicate no price protection threshold for Spot Instances, meaning you want to consider all instance types that match your attributes, include one of these parameters and specify a high value, such as ``999999``.
- `memory_gi_b_per_v_cpu` (Attributes) The minimum and maximum amount of memory per vCPU for an instance type, in GiB.
 Default: No minimum or maximum limits (see [below for nested schema](#nestedatt--mixed_instances_policy--launch_template--overrides--instance_requirements--memory_gi_b_per_v_cpu))
- `memory_mi_b` (Attributes) The minimum and maximum instance memory size for an instance type, in MiB. (see [below for nested schema](#nestedatt--mixed_instances_policy--launch_template--overrides--instance_requirements--memory_mi_b))
- `network_bandwidth_gbps` (Attributes) The minimum and maximum amount of network bandwidth, in gigabits per second (Gbps).
 Default: No minimum or maximum limits (see [below for nested schema](#nestedatt--mixed_instances_policy--launch_template--overrides--instance_requirements--network_bandwidth_gbps))
- `network_interface_count` (Attributes) The minimum and maximum number of network interfaces for an instance type.
 Default: No minimum or maximum limits (see [below for nested schema](#nestedatt--mixed_instances_policy--launch_template--overrides--instance_requirements--network_interface_count))
- `on_demand_max_price_percentage_over_lowest_price` (Number) [Price protection] The price protection threshold for On-Demand Instances, as a percentage higher than an identified On-Demand price. The identified On-Demand price is the price of the lowest priced current generation C, M, or R instance type with your specified attributes. If no current generation C, M, or R instance type matches your attributes, then the identified price is from either the lowest priced current generation instance types or, failing that, the lowest priced previous generation instance types that match your attributes. When Amazon EC2 Auto Scaling selects instance types with your attributes, we will exclude instance types whose price exceeds your specified threshold. 
 The parameter accepts an integer, which Amazon EC2 Auto Scaling interprets as a percentage.
 To turn off price protection, specify a high value, such as ``999999``. 
 If you set ``DesiredCapacityType`` to ``vcpu`` or ``memory-mib``, the price protection threshold is applied based on the per-vCPU or per-memory price instead of the per instance price. 
 Default: ``20``
- `require_hibernate_support` (Boolean) Indicates whether instance types must provide On-Demand Instance hibernation support.
 Default: ``false``
- `spot_max_price_percentage_over_lowest_price` (Number) [Price protection] The price protection threshold for Spot Instances, as a percentage higher than an identified Spot price. The identified Spot price is the price of the lowest priced current generation C, M, or R instance type with your specified attributes. If no current generation C, M, or R instance type matches your attributes, then the identified price is from either the lowest priced current generation instance types or, failing that, the lowest priced previous generation instance types that match your attributes. When Amazon EC2 Auto Scaling selects instance types with your attributes, we will exclude instance types whose price exceeds your specified threshold.
 The parameter accepts an integer, which Amazon EC2 Auto Scaling interprets as a percentage. 
 If you set ``DesiredCapacityType`` to ``vcpu`` or ``memory-mib``, the price protection threshold is based on the per-vCPU or per-memory price instead of the per instance price. 
  Only one of ``SpotMaxPricePercentageOverLowestPrice`` or ``MaxSpotPriceAsPercentageOfOptimalOnDemandPrice`` can be specified. If you don't specify either, Amazon EC2 Auto Scaling will automatically apply optimal price protection to consistently select from a wide range of instance types. To indicate no price protection threshold for Spot Instances, meaning you want to consider all instance types that match your attributes, include one of these parameters and specify a high value, such as ``999999``.
- `total_local_storage_gb` (Attributes) The minimum and maximum total local storage size for an instance type, in GB.
 Default: No minimum or maximum limits (see [below for nested schema](#nestedatt--mixed_instances_policy--launch_template--overrides--instance_requirements--total_local_storage_gb))
- `v_cpu_count` (Attributes) The minimum and maximum number of vCPUs for an instance type. (see [below for nested schema](#nestedatt--mixed_instances_policy--launch_template--overrides--instance_requirements--v_cpu_count))

<a id="nestedatt--mixed_instances_policy--launch_template--overrides--instance_requirements--accelerator_count"></a>
### Nested Schema for `mixed_instances_policy.launch_template.overrides.instance_requirements.accelerator_count`

Optional:

- `max` (Number) The maximum value.
- `min` (Number) The minimum value.


<a id="nestedatt--mixed_instances_policy--launch_template--overrides--instance_requirements--accelerator_total_memory_mi_b"></a>
### Nested Schema for `mixed_instances_policy.launch_template.overrides.instance_requirements.accelerator_total_memory_mi_b`

Optional:

- `max` (Number) The memory maximum in MiB.
- `min` (Number) The memory minimum in MiB.


<a id="nestedatt--mixed_instances_policy--launch_template--overrides--instance_requirements--baseline_ebs_bandwidth_mbps"></a>
### Nested Schema for `mixed_instances_policy.launch_template.overrides.instance_requirements.baseline_ebs_bandwidth_mbps`

Optional:

- `max` (Number) The maximum value in Mbps.
- `min` (Number) The minimum value in Mbps.


<a id="nestedatt--mixed_instances_policy--launch_template--overrides--instance_requirements--baseline_performance_factors"></a>
### Nested Schema for `mixed_instances_policy.launch_template.overrides.instance_requirements.baseline_performance_factors`

Optional:

- `cpu` (Attributes) (see [below for nested schema](#nestedatt--mixed_instances_policy--launch_template--overrides--instance_requirements--baseline_performance_factors--cpu))

<a id="nestedatt--mixed_instances_policy--launch_template--overrides--instance_requirements--baseline_performance_factors--cpu"></a>
### Nested Schema for `mixed_instances_policy.launch_template.overrides.instance_requirements.baseline_performance_factors.cpu`

Optional:

- `references` (Attributes List) (see [below for nested schema](#nestedatt--mixed_instances_policy--launch_template--overrides--instance_requirements--baseline_performance_factors--cpu--references))

<a id="nestedatt--mixed_instances_policy--launch_template--overrides--instance_requirements--baseline_performance_factors--cpu--references"></a>
### Nested Schema for `mixed_instances_policy.launch_template.overrides.instance_requirements.baseline_performance_factors.cpu.references`

Optional:

- `instance_family` (String)




<a id="nestedatt--mixed_instances_policy--launch_template--overrides--instance_requirements--memory_gi_b_per_v_cpu"></a>
### Nested Schema for `mixed_instances_policy.launch_template.overrides.instance_requirements.memory_gi_b_per_v_cpu`

Optional:

- `max` (Number) The memory maximum in GiB.
- `min` (Number) The memory minimum in GiB.


<a id="nestedatt--mixed_instances_policy--launch_template--overrides--instance_requirements--memory_mi_b"></a>
### Nested Schema for `mixed_instances_policy.launch_template.overrides.instance_requirements.memory_mi_b`

Optional:

- `max` (Number) The memory maximum in MiB.
- `min` (Number) The memory minimum in MiB.


<a id="nestedatt--mixed_instances_policy--launch_template--overrides--instance_requirements--network_bandwidth_gbps"></a>
### Nested Schema for `mixed_instances_policy.launch_template.overrides.instance_requirements.network_bandwidth_gbps`

Optional:

- `max` (Number) The maximum amount of network bandwidth, in gigabits per second (Gbps).
- `min` (Number) The minimum amount of network bandwidth, in gigabits per second (Gbps).


<a id="nestedatt--mixed_instances_policy--launch_template--overrides--instance_requirements--network_interface_count"></a>
### Nested Schema for `mixed_instances_policy.launch_template.overrides.instance_requirements.network_interface_count`

Optional:

- `max` (Number) The maximum number of network interfaces.
- `min` (Number) The minimum number of network interfaces.


<a id="nestedatt--mixed_instances_policy--launch_template--overrides--instance_requirements--total_local_storage_gb"></a>
### Nested Schema for `mixed_instances_policy.launch_template.overrides.instance_requirements.total_local_storage_gb`

Optional:

- `max` (Number) The storage maximum in GB.
- `min` (Number) The storage minimum in GB.


<a id="nestedatt--mixed_instances_policy--launch_template--overrides--instance_requirements--v_cpu_count"></a>
### Nested Schema for `mixed_instances_policy.launch_template.overrides.instance_requirements.v_cpu_count`

Optional:

- `max` (Number) The maximum number of vCPUs.
- `min` (Number) The minimum number of vCPUs.



<a id="nestedatt--mixed_instances_policy--launch_template--overrides--launch_template_specification"></a>
### Nested Schema for `mixed_instances_policy.launch_template.overrides.launch_template_specification`

Optional:

- `launch_template_id` (String) The ID of the launch template.
 You must specify the ``LaunchTemplateID`` or the ``LaunchTemplateName``, but not both.
- `launch_template_name` (String) The name of the launch template.
 You must specify the ``LaunchTemplateName`` or the ``LaunchTemplateID``, but not both.
- `version` (String) The version number of the launch template.
 Specifying ``$Latest`` or ``$Default`` for the template version number is not supported. However, you can specify ``LatestVersionNumber`` or ``DefaultVersionNumber`` using the ``Fn::GetAtt`` intrinsic function. For more information, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).
  For an example of using the ``Fn::GetAtt`` function, see the [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#aws-resource-autoscaling-autoscalinggroup--examples) section of the ``AWS::AutoScaling::AutoScalingGroup`` resource.





<a id="nestedatt--notification_configuration"></a>
### Nested Schema for `notification_configuration`

Optional:

- `notification_types` (List of String) A list of event types that send a notification. Event types can include any of the following types. 
  *Allowed values*:
  +   ``autoscaling:EC2_INSTANCE_LAUNCH`` 
  +   ``autoscaling:EC2_INSTANCE_LAUNCH_ERROR`` 
  +   ``autoscaling:EC2_INSTANCE_TERMINATE`` 
  +   ``autoscaling:EC2_INSTANCE_TERMINATE_ERROR`` 
  +   ``autoscaling:TEST_NOTIFICATION``
- `topic_arn` (String) The Amazon Resource Name (ARN) of the Amazon SNS topic.


<a id="nestedatt--notification_configurations"></a>
### Nested Schema for `notification_configurations`

Optional:

- `notification_types` (List of String) A list of event types that send a notification. Event types can include any of the following types. 
  *Allowed values*:
  +   ``autoscaling:EC2_INSTANCE_LAUNCH`` 
  +   ``autoscaling:EC2_INSTANCE_LAUNCH_ERROR`` 
  +   ``autoscaling:EC2_INSTANCE_TERMINATE`` 
  +   ``autoscaling:EC2_INSTANCE_TERMINATE_ERROR`` 
  +   ``autoscaling:TEST_NOTIFICATION``
- `topic_arn` (String) The Amazon Resource Name (ARN) of the Amazon SNS topic.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The tag key.
- `propagate_at_launch` (Boolean) Set to ``true`` if you want CloudFormation to copy the tag to EC2 instances that are launched as part of the Auto Scaling group. Set to ``false`` if you want the tag attached only to the Auto Scaling group and not copied to any instances launched as part of the Auto Scaling group.
- `value` (String) The tag value.


<a id="nestedatt--traffic_sources"></a>
### Nested Schema for `traffic_sources`

Optional:

- `identifier` (String) Identifies the traffic source.
 For Application Load Balancers, Gateway Load Balancers, Network Load Balancers, and VPC Lattice, this will be the Amazon Resource Name (ARN) for a target group in this account and Region. For Classic Load Balancers, this will be the name of the Classic Load Balancer in this account and Region.
 For example: 
  +  Application Load Balancer ARN: ``arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/my-targets/1234567890123456`` 
  +  Classic Load Balancer name: ``my-classic-load-balancer`` 
  +  VPC Lattice ARN: ``arn:aws:vpc-lattice:us-west-2:123456789012:targetgroup/tg-1234567890123456`` 
  
 To get the ARN of a target group for a Application Load Balancer, Gateway Load Balancer, or Network Load Balancer, or the name of a Classic Load Balancer, use the Elastic Load Balancing [DescribeTargetGroups](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_DescribeTargetGroups.html) and [DescribeLoadBalancers](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_DescribeLoadBalancers.html) API operations.
 To get the ARN of a target group for VPC Lattice, use the VPC Lattice [GetTargetGroup](https://docs.aws.amazon.com/vpc-lattice/latest/APIReference/API_GetTargetGroup.html) API operation.
- `type` (String) Provides additional context for the value of ``Identifier``.
 The following lists the valid values:
  +   ``elb`` if ``Identifier`` is the name of a Classic Load Balancer.
  +   ``elbv2`` if ``Identifier`` is the ARN of an Application Load Balancer, Gateway Load Balancer, or Network Load Balancer target group.
  +   ``vpc-lattice`` if ``Identifier`` is the ARN of a VPC Lattice target group.
  
 Required if the identifier is the name of a Classic Load Balancer.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_autoscaling_auto_scaling_group.example "auto_scaling_group_name"
```