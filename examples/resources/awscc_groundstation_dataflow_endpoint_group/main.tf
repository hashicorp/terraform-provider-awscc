# Get current AWS region
data "aws_region" "current" {}

# Create a VPC for the Ground Station
resource "aws_vpc" "ground_station" {
  cidr_block = "10.0.0.0/16"
  tags = {
    Name = "ground-station-vpc"
  }
}

# Create a subnet
resource "aws_subnet" "ground_station" {
  vpc_id            = aws_vpc.ground_station.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "${data.aws_region.current.name}a"
  tags = {
    Name = "ground-station-subnet"
  }
}

# Create security group
resource "aws_security_group" "ground_station" {
  name_prefix = "ground-station-"
  description = "Security group for Ground Station dataflow endpoint"
  vpc_id      = aws_vpc.ground_station.id

  ingress {
    from_port   = 55556
    to_port     = 55556
    protocol    = "tcp"
    cidr_blocks = ["192.0.2.0/24"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "ground-station-sg"
  }
}

# Create IAM role for Ground Station
data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["groundstation.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_role" "ground_station" {
  name               = "groundstation-role"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json

  tags = {
    ModifiedBy = "AWSCC"
  }
}

# Create the Ground Station Dataflow Endpoint Group
resource "awscc_groundstation_dataflow_endpoint_group" "example" {
  endpoint_details = [{
    endpoint = {
      name = "example-endpoint"
      address = {
        name = "192.0.2.0"
        port = 55556
      }
      mtu = 1500
    }
    security_details = {
      role_arn           = aws_iam_role.ground_station.arn
      security_group_ids = [aws_security_group.ground_station.id]
      subnet_ids         = [aws_subnet.ground_station.id]
    }
  }]

  contact_post_pass_duration_seconds = 60
  contact_pre_pass_duration_seconds  = 60

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}