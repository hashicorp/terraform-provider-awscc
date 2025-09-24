# Find the most recent ECS-optimized AMI
data "aws_ami" "ecs_ami" {
  most_recent = true
  owners      = ["amazon"]

  filter {
    name   = "name"
    values = ["amzn2-ami-ecs-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }
}

# Get default VPC and subnet for simplicity
data "aws_vpc" "default" {
  default = true
}

data "aws_subnets" "default" {
  filter {
    name   = "vpc-id"
    values = [data.aws_vpc.default.id]
  }
}

# Security Group for ASG instances
resource "aws_security_group" "ecs_instances" {
  name_prefix = "ecs-instances-"
  description = "Security group for ECS instances"
  vpc_id      = data.aws_vpc.default.id

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "ecs-capacity-provider-sg"
  }
}

# Create IAM role for ECS instances
data "aws_iam_policy_document" "ecs_instance_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["ec2.amazonaws.com"]
    }
  }
}

resource "awscc_iam_role" "ecs_instance_role" {
  assume_role_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.ecs_instance_assume_role.json))
  managed_policy_arns         = ["arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role"]
  tags = [{
    key   = "Name"
    value = "ecs-instance-role"
  }]
}

# Create instance profile for ECS instances
resource "aws_iam_instance_profile" "ecs_instance_profile" {
  name = "ecs-instance-profile"
  role = awscc_iam_role.ecs_instance_role.role_name
}

# Create Launch Template
resource "aws_launch_template" "ecs" {
  name_prefix   = "ecs-template"
  image_id      = data.aws_ami.ecs_ami.id
  instance_type = "t3.micro"

  user_data = base64encode(<<-EOF
              #!/bin/bash
              echo ECS_CLUSTER=${awscc_ecs_cluster.example.cluster_name} >> /etc/ecs/ecs.config
              EOF
  )

  iam_instance_profile {
    name = aws_iam_instance_profile.ecs_instance_profile.name
  }

  vpc_security_group_ids = [aws_security_group.ecs_instances.id]

  tag_specifications {
    resource_type = "instance"
    tags = {
      Name = "ecs-instance"
    }
  }
}

# Create Auto Scaling Group
resource "aws_autoscaling_group" "ecs" {
  name                = "ecs-asg"
  desired_capacity    = 1
  max_size            = 2
  min_size            = 0
  target_group_arns   = []
  vpc_zone_identifier = [data.aws_subnets.default.ids[0]]

  launch_template {
    id      = aws_launch_template.ecs.id
    version = "$Latest"
  }

  tag {
    key                 = "Name"
    value               = "ecs-asg"
    propagate_at_launch = true
  }

  lifecycle {
    ignore_changes = [desired_capacity]
  }
}

# Create ECS Cluster
resource "awscc_ecs_cluster" "example" {
  cluster_name = "example"
  tags = [{
    key   = "Name"
    value = "example-cluster"
  }]
}

# Create ECS Capacity Provider
resource "awscc_ecs_capacity_provider" "example" {
  name = "example-capacity-provider"

  auto_scaling_group_provider = {
    auto_scaling_group_arn         = aws_autoscaling_group.ecs.arn
    managed_termination_protection = "DISABLED"
    managed_scaling = {
      status                    = "ENABLED"
      target_capacity           = 100
      minimum_scaling_step_size = 1
      maximum_scaling_step_size = 100
      instance_warmup_period    = 300
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}