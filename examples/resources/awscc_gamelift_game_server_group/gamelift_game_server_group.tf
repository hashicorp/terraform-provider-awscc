resource "awscc_gamelift_game_server_group" "example_server_group" {
  game_server_group_name = "example-game-server-group"
  instance_definitions = [
    { instance_type = "c5.large" },
    { instance_type = "c6g.xlarge" }
  ]
  role_arn = awscc_iam_role.example_iam_role.arn
  min_size = "1"
  max_size = "3"
  launch_template = {
    launch_template_id = awscc_ec2_launch_template.gamelift_launch_template.id
  }
}

resource "awscc_iam_role" "example_iam_role" {
  role_name                   = "example_gamelift_iam_role"
  description                 = "This IAM role grants Amazon GameLift GameServerGroup to manage GameLift EC2 Fleet."
  assume_role_policy_document = data.aws_iam_policy_document.instance_assume_role_policy.json
  managed_policy_arns         = ["arn:aws:iam::aws:policy/GameLiftGameServerGroupPolicy"]
  max_session_duration        = 7200
  path                        = "/"
  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}

resource "awscc_ec2_launch_template" "gamelift_launch_template" {
  launch_template_data = {
    image_id = data.aws_ami.amazon_linux.id
  }
  launch_template_name = "gamelift_launch_template"
}

data "aws_iam_policy_document" "instance_assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["gamelift.amazonaws.com"]
    }
  }
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