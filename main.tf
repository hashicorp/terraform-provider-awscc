provider "aws" {}
provider "awscc" {}

resource "awscc_autoscaling_auto_scaling_group" "example" {
  count = 1

  max_size = "1"
  min_size = "0"

  launch_template = {
    version            = awscc_ec2_launch_template.example.latest_version_number
    launch_template_id = awscc_ec2_launch_template.example.id
  }

  desired_capacity    = "1"
  
  vpc_zone_identifier = [
    "subnet-0afab9555ada9f55c",
    "subnet-0d7b3145382362c4f",
    "subnet-0b83da4594361dcc0",

    # "subnet-0afab9555ada9f55c",
    # "subnet-0b83da4594361dcc0",
    # "subnet-0d7b3145382362c4f",
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
  launch_template_name = "ewbankkit-test"
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

