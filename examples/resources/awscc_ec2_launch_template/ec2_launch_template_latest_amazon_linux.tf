resource "awscc_ec2_launch_template" "amazon_linux" {
  launch_template_data = {
    image_id      = data.aws_ami.amazon_linux.id
    instance_type = "t2.micro"
  }
  launch_template_name = "latest_amazon_linux"
}

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