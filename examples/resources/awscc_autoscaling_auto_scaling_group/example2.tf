resource "awscc_autoscaling_auto_scaling_group" "example" {
  max_size = "1"
  min_size = "0"

  launch_template = {
    version            = awscc_ec2_launch_template.example.latest_version_number
    launch_template_id = awscc_ec2_launch_template.example.id
  }

  desired_capacity = "1"

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