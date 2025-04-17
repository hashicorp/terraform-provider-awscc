# Create a Launch Template first
resource "awscc_ec2_launch_template" "example" {
  launch_template_data = {
    instance_type = "t3.micro"
    image_id      = data.aws_ami.amazon_linux_2.id
  }

}

# Get the latest Amazon Linux 2 AMI
data "aws_ami" "amazon_linux_2" {
  most_recent = true
  owners      = ["amazon"]

  filter {
    name   = "name"
    values = ["amzn2-ami-hvm-*-x86_64-gp2"]
  }
}

# Data source to get default VPC subnets
data "aws_vpc" "default" {
  default = true
}

data "aws_subnets" "default" {
  filter {
    name   = "vpc-id"
    values = [data.aws_vpc.default.id]
  }
}

# Create an Auto Scaling Group
resource "awscc_autoscaling_auto_scaling_group" "example" {
  max_size            = 4
  min_size            = 1
  desired_capacity    = 1
  vpc_zone_identifier = [data.aws_subnets.default.ids[0]]

  launch_template = {
    launch_template_id = awscc_ec2_launch_template.example.id
    version            = 1
  }

  tags = [{
    key                 = "Modified By"
    value               = "AWSCC"
    propagate_at_launch = true
  }]
}

# Target Tracking Scaling Policy Example
resource "awscc_autoscaling_scaling_policy" "cpu_tracking" {
  auto_scaling_group_name   = awscc_autoscaling_auto_scaling_group.example.id
  policy_type               = "TargetTrackingScaling"
  estimated_instance_warmup = 300

  target_tracking_configuration = {
    target_value = 40.0
    predefined_metric_specification = {
      predefined_metric_type = "ASGAverageCPUUtilization"
    }
    disable_scale_in = false
  }
}

# Step Scaling Policy Example
resource "awscc_autoscaling_scaling_policy" "step_scaling" {
  auto_scaling_group_name   = awscc_autoscaling_auto_scaling_group.example.id
  policy_type               = "StepScaling"
  adjustment_type           = "ChangeInCapacity"
  metric_aggregation_type   = "Average"
  estimated_instance_warmup = 300

  step_adjustments = [
    {
      metric_interval_lower_bound = 0.0
      metric_interval_upper_bound = 20.0
      scaling_adjustment          = 1
    },
    {
      metric_interval_lower_bound = 20.0
      scaling_adjustment          = 2
    }
  ]
}

# Simple Scaling Policy Example
resource "awscc_autoscaling_scaling_policy" "simple_scaling" {
  auto_scaling_group_name = awscc_autoscaling_auto_scaling_group.example.id
  policy_type             = "SimpleScaling"
  adjustment_type         = "ChangeInCapacity"
  scaling_adjustment      = 1
  cooldown                = "300"
}