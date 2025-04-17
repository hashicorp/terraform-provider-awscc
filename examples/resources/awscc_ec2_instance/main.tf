# VPC
resource "awscc_ec2_vpc" "example" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true
  tags = [{
    key   = "Name"
    value = "example-vpc"
  }]
}

# Subnet
resource "awscc_ec2_subnet" "example" {
  vpc_id                  = awscc_ec2_vpc.example.id
  cidr_block              = "10.0.1.0/24"
  map_public_ip_on_launch = true
  tags = [{
    key   = "Name"
    value = "example-subnet"
  }]
}

# Security Group
resource "awscc_ec2_security_group" "example" {
  group_description = "Example security group"
  vpc_id            = awscc_ec2_vpc.example.id
  security_group_ingress = [
    {
      ip_protocol = "tcp"
      from_port   = 22
      to_port     = 22
      cidr_ip     = "0.0.0.0/0"
    }
  ]
  tags = [{
    key   = "Name"
    value = "example-sg"
  }]
}

# IAM Role for the EC2 Instance
resource "awscc_iam_role" "example" {
  role_name = "example-ec2-role"
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
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# EC2 Instance Profile
resource "awscc_iam_instance_profile" "example" {
  instance_profile_name = "example-profile"
  roles                 = [awscc_iam_role.example.role_name]
}

# EC2 Instance
resource "awscc_ec2_instance" "example" {
  instance_type = "t2.micro"
  image_id      = "ami-0c55b159cbfafe1f0" # Amazon Linux 2 AMI ID
  subnet_id     = awscc_ec2_subnet.example.id

  iam_instance_profile = awscc_iam_instance_profile.example.instance_profile_name

  security_group_ids = [awscc_ec2_security_group.example.id]

  tags = [{
    key   = "Name"
    value = "example-instance"
    }, {
    key   = "Modified By"
    value = "AWSCC"
  }]

  block_device_mappings = [
    {
      device_name = "/dev/xvda"
      ebs = {
        volume_size           = 8
        volume_type           = "gp2"
        delete_on_termination = true
      }
    }
  ]
}