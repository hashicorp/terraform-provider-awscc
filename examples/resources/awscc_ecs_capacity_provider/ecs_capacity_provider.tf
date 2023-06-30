resource "awscc_ecs_capacity_provider" "example" {
  name = "example"

  auto_scaling_group_provider = {
    auto_scaling_group_arn         = aws_autoscaling_group.example.arn
    managed_termination_protection = "ENABLED"

    managed_scaling = {
      maximum_scaling_step_size = 1000
      minimum_scaling_step_size = 1
      status                    = "ENABLED"
      target_capacity           = 10
    }
  }
  lifecycle {
    ignore_changes = [auto_scaling_group_provider.auto_scaling_group_arn]
  }
}

resource "aws_autoscaling_group" "example" {
  availability_zones    = ["us-east-1a"]
  desired_capacity      = 1
  max_size              = 5
  min_size              = 1
  protect_from_scale_in = true
  launch_template {
    id      = aws_launch_template.example.id
    version = "$Latest"
  }
  
  tag {
    key                 = "AmazonECSManaged"
    value               = true
    propagate_at_launch = true
  }
}

resource "aws_launch_template" "example" {
  name_prefix   = "example"
  image_id      = "ami-0ff8a91507f77f867"
  instance_type = "t2.micro"
}