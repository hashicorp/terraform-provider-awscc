data "aws_availability_zones" "available" {
  state = "available"
}

# Get the latest Amazon Linux 2 AMI
data "aws_ami" "amazon_linux" {
  most_recent = true
  owners      = ["amazon"]

  filter {
    name   = "name"
    values = ["amzn2-ami-hvm-*-x86_64-gp2"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }
}

# VPC for the Auto Scaling group
resource "awscc_ec2_vpc" "main" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = [{
    key   = "Name"
    value = "warm-pool-vpc"
    }, {
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Internet Gateway
resource "awscc_ec2_internet_gateway" "main" {
  tags = [{
    key   = "Name"
    value = "warm-pool-igw"
    }, {
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Attach Internet Gateway to VPC
resource "awscc_ec2_vpc_gateway_attachment" "main" {
  vpc_id              = awscc_ec2_vpc.main.id
  internet_gateway_id = awscc_ec2_internet_gateway.main.id
}

# Public subnets for Auto Scaling group
resource "awscc_ec2_subnet" "public" {
  count = 2

  vpc_id                  = awscc_ec2_vpc.main.id
  cidr_block              = "10.0.${count.index + 1}.0/24"
  availability_zone       = data.aws_availability_zones.available.names[count.index]
  map_public_ip_on_launch = true

  tags = [{
    key   = "Name"
    value = "warm-pool-public-subnet-${count.index + 1}"
    }, {
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Route table for public subnets
resource "awscc_ec2_route_table" "public" {
  vpc_id = awscc_ec2_vpc.main.id

  tags = [{
    key   = "Name"
    value = "warm-pool-public-rt"
    }, {
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Route to Internet Gateway
resource "awscc_ec2_route" "public_internet" {
  route_table_id         = awscc_ec2_route_table.public.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = awscc_ec2_internet_gateway.main.id

  depends_on = [awscc_ec2_vpc_gateway_attachment.main]
}

# Associate public subnets with route table
resource "awscc_ec2_subnet_route_table_association" "public" {
  count = 2

  subnet_id      = awscc_ec2_subnet.public[count.index].id
  route_table_id = awscc_ec2_route_table.public.id
}
# Security group for web servers
resource "awscc_ec2_security_group" "web" {
  group_name        = "warm-pool-web-sg"
  group_description = "Security group for web servers in warm pool"
  vpc_id            = awscc_ec2_vpc.main.id

  security_group_ingress = [
    {
      ip_protocol = "tcp"
      from_port   = 80
      to_port     = 80
      cidr_ip     = "0.0.0.0/0"
      description = "HTTP access"
    },
    {
      ip_protocol = "tcp"
      from_port   = 443
      to_port     = 443
      cidr_ip     = "0.0.0.0/0"
      description = "HTTPS access"
    },
    {
      ip_protocol = "tcp"
      from_port   = 22
      to_port     = 22
      cidr_ip     = "10.0.0.0/16"
      description = "SSH access from VPC"
    }
  ]

  security_group_egress = [
    {
      ip_protocol = "-1"
      cidr_ip     = "0.0.0.0/0"
      description = "All outbound traffic"
    }
  ]

  tags = [{
    key   = "Name"
    value = "warm-pool-web-sg"
    }, {
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# IAM role for EC2 instances
resource "awscc_iam_role" "ec2_role" {
  role_name = "warm-pool-ec2-role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "ec2.amazonaws.com"
        }
      }
    ]
  })

  managed_policy_arns = [
    "arn:aws:iam::aws:policy/CloudWatchAgentServerPolicy"
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Instance profile for EC2 instances
resource "awscc_iam_instance_profile" "ec2_profile" {
  instance_profile_name = "warm-pool-ec2-profile"
  roles                 = [awscc_iam_role.ec2_role.role_name]
}
# Launch template for Auto Scaling group
resource "aws_launch_template" "web" {
  name_prefix   = "warm-pool-web-"
  image_id      = data.aws_ami.amazon_linux.id
  instance_type = "t3.micro"

  vpc_security_group_ids = [awscc_ec2_security_group.web.id]

  iam_instance_profile {
    name = awscc_iam_instance_profile.ec2_profile.instance_profile_name
  }

  user_data = base64encode(<<-EOF
    #!/bin/bash
    yum update -y
    yum install -y httpd
    systemctl start httpd
    systemctl enable httpd
    echo "<h1>Hello from Warm Pool Instance</h1>" > /var/www/html/index.html
    echo "<p>Instance ID: $(curl -s http://169.254.169.254/latest/meta-data/instance-id)</p>" >> /var/www/html/index.html
    echo "<p>Availability Zone: $(curl -s http://169.254.169.254/latest/meta-data/placement/availability-zone)</p>" >> /var/www/html/index.html
  EOF
  )

  tag_specifications {
    resource_type = "instance"
    tags = {
      Name          = "warm-pool-web-instance"
      "Modified By" = "AWSCC"
    }
  }

  # Enable hibernation support for instances
  hibernation_options {
    configured = true
  }

  # Use EBS-optimized instances for better performance
  ebs_optimized = true

  # Configure root volume
  block_device_mappings {
    device_name = "/dev/xvda"
    ebs {
      volume_size           = 20
      volume_type           = "gp3"
      encrypted             = true
      delete_on_termination = true
    }
  }

  lifecycle {
    create_before_destroy = true
  }
}

# Auto Scaling group
resource "aws_autoscaling_group" "web" {
  name                      = "warm-pool-asg"
  vpc_zone_identifier       = [for subnet in awscc_ec2_subnet.public : subnet.id]
  target_group_arns         = []
  health_check_type         = "EC2"
  health_check_grace_period = 300

  min_size         = 1
  max_size         = 10
  desired_capacity = 2

  launch_template {
    id      = aws_launch_template.web.id
    version = "$Latest"
  }

  # Enable instance refresh for rolling updates
  instance_refresh {
    strategy = "Rolling"
    preferences {
      min_healthy_percentage = 50
    }
  }

  tag {
    key                 = "Name"
    value               = "warm-pool-asg-instance"
    propagate_at_launch = true
  }

  tag {
    key                 = "Modified By"
    value               = "AWSCC"
    propagate_at_launch = true
  }

  lifecycle {
    create_before_destroy = true
  }
}

# AWSCC Auto Scaling Warm Pool - Main Resource
resource "awscc_autoscaling_warm_pool" "web" {
  auto_scaling_group_name     = aws_autoscaling_group.web.name
  min_size                    = 2
  max_group_prepared_capacity = 8
  pool_state                  = "Stopped"
  instance_reuse_policy = {
    reuse_on_scale_in = true
  }
}
