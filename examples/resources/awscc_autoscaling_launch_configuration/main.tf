# Get latest Amazon Linux 2 AMI
data "aws_ami" "amazon_linux_2" {
  most_recent = true
  owners      = ["amazon"]

  filter {
    name   = "name"
    values = ["amzn2-ami-hvm-*-x86_64-gp2"]
  }
}

# Create IAM instance profile and role
data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"
    actions = [
      "sts:AssumeRole",
    ]
    principals {
      type = "Service"
      identifiers = [
        "ec2.amazonaws.com"
      ]
    }
  }
}

resource "awscc_iam_role" "ec2_role" {
  role_name                   = "example-asg-instance-role"
  assume_role_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.assume_role.json))
  managed_policy_arns         = ["arn:aws:iam::aws:policy/service-role/AmazonEC2RoleforSSM"]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_iam_instance_profile" "instance_profile" {
  instance_profile_name = "example-asg-instance-profile"
  roles                 = [awscc_iam_role.ec2_role.role_name]
}

# Create a security group
resource "awscc_ec2_security_group" "example" {
  group_description = "Example security group for launch configuration"
  group_name        = "example-lc-sg"
  security_group_ingress = [
    {
      ip_protocol = "tcp"
      from_port   = 80
      to_port     = 80
      cidr_ip     = "0.0.0.0/0"
    }
  ]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create launch configuration
resource "awscc_autoscaling_launch_configuration" "example" {
  image_id      = data.aws_ami.amazon_linux_2.id
  instance_type = "t3.micro"

  iam_instance_profile = awscc_iam_instance_profile.instance_profile.instance_profile_name
  security_groups      = [awscc_ec2_security_group.example.id]

  associate_public_ip_address = true
  instance_monitoring         = true

  user_data = base64encode(<<-EOF
              #!/bin/bash
              yum update -y
              yum install -y httpd
              systemctl start httpd
              systemctl enable httpd
              echo "Hello from EC2 instance!" > /var/www/html/index.html
              EOF
  )

  block_device_mappings = [{
    device_name = "/dev/xvda"
    ebs = {
      volume_size           = 20
      volume_type           = "gp3"
      delete_on_termination = true
      encrypted             = true
    }
  }]

  metadata_options = {
    http_endpoint               = "enabled"
    http_tokens                 = "required"
    http_put_response_hop_limit = 2
  }
}